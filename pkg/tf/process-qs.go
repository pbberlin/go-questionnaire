package tf

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"path"
	"sort"
	"strings"

	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/qst"
)

// ProcessQs iterates over qs
// and extracts columns and values;
// it is independent of the structure of the questionaires in qs
func ProcessQs(cfgRem *RemoteConnConfigT, qs []*qst.QuestionnaireT, saveQSFilesToDownloadDir bool) (string, error) {

	if cfgRem.DownloadDir == "" {
		log.Fatal("download dir cannot be empty")
	}

	fnCSV := path.Join(cfgRem.DownloadDir, fmt.Sprintf("%v-%v.csv", cfgRem.SurveyType, cfgRem.WaveID))

	dirFull := path.Join(cfgRem.DownloadDir, cfgRem.SurveyType, cfgRem.WaveID)
	dirEmpty := path.Join(dirFull, "empty")

	//
	//
	//
	maxPages := 0
	for _, q := range qs {
		if maxPages < len(q.Pages) {
			maxPages = len(q.Pages)
		}
	}

	keysByQ := [][]string{} // keys per questionnaire, separate slice for every response
	valsByQ := [][]string{} // vals per questionnaire, separate slice for every response

	// CSV header stuff:
	staticCols := []string{ // across all questionnaires
		"user_id",
		"lang_code",
		"closing_time",
		"status",
		"remote_ip",
		"user_agent",
		"version",
		"version_max",
	}
	for iPg := 0; iPg < maxPages; iPg++ {
		staticCols = append(staticCols, fmt.Sprintf("page_%v", iPg+1))
	}

	nonEmpty := 0
	empty := 0

	//
	//
	// Process questionnaires
	for i, q := range qs {

		// log.Printf("  ")

		serverSideMD5 := q.MD5

		pthEmpty := path.Join(dirEmpty, q.UserID+".json") // delete empty questionnaires and save them elsewhere
		pthFull := path.Join(dirFull, q.UserID+".json")

		//
		if q.MD5 != serverSideMD5 {
			// log.Printf("%3v: MD5 does not match: %v\nwnt %v\ngot %v", i, pth2, md5BeforeSave, q.MD5)
			log.Printf("%3v: Server side and new client side MD5 hashes do not match %v - %v", i, q.Survey.String(), pthFull)
		}

		if cfgRem.MinUserID != 0 {
			if q.UserIDInt() < cfgRem.MinUserID {
				log.Printf("%3v: Skipping UserID %v < %v", i, q.UserID, cfgRem.MinUserID)
				continue
			}
		}

		if cfgRem.MaxUserID != 0 {
			if q.UserIDInt() > cfgRem.MaxUserID {
				log.Printf("%3v: Skipping UserID %v > %v", i, q.UserID, cfgRem.MaxUserID)
				continue
			}
		}

		//
		//
		//
		// current run: move empty to dir empty
		realEntries, _, _ := q.Statistics()

		if realEntries > 0 {

			if saveQSFilesToDownloadDir {
				// save non empty
				err := q.Save1(pthFull)
				if err != nil {
					log.Printf("%3v: Error saving %v: %v", i, pthFull, err)
					continue
				}
				// quest file may have been empty during previous runs;
				// but may be filled now; delete previously empty
				err = cloudio.Delete(pthEmpty)
				if err != nil && !cloudio.IsNotExist(err) {
					log.Printf("%3v: Error removing previously empty %v - %v", i, pthEmpty, err)
				}

			}
			nonEmpty++

		} else {

			// realEntries == 0
			if saveQSFilesToDownloadDir {
				log.Printf("%3v: %v. No answers, moving to %v.", i, pthFull, "empty")
				err := cloudio.Delete(pthFull)
				if err != nil && !cloudio.IsNotExist(err) {
					log.Printf("%3v: Error removing empty %v - %v", i, pthFull, err)
				}
				err = q.Save1(pthEmpty)
				if err != nil {
					log.Printf("%3v: Error saving  to empty %v: %v", i, pthEmpty, err)
				}
			}
			empty++
			continue

		}

		// Prepare columns...
		finishes, ks, vs := q.KeysValues(true)

		ks = append(staticCols, ks...)
		keysByQ = append(keysByQ, ks) // appending to _all_ responses

		formattedClosingTime := ""
		status := "0"
		if qs[i].ClosingTime.IsZero() {
			for i2 := len(qs[i].Pages) - 1; i2 > -1; i2-- {
				if !qs[i].Pages[i2].Finished.IsZero() {
					formattedClosingTime = fmt.Sprintf("%v", qs[i].Pages[i2].Finished.Unix())
					status = "1"
					break
				}
			}
		} else {
			formattedClosingTime = fmt.Sprintf("%v", qs[i].ClosingTime.Unix())
			status = "2"
		}

		// Collecting the values for staticCols...
		prepend := []string{
			qs[i].UserID,                      // user_id
			qs[i].LangCode,                    // lang_code
			formattedClosingTime,              // closing_time
			status,                            // status
			q.RemoteIP,                        // remote_ip
			qst.CleanseUserAgent(q.UserAgent), // user_agent
			fmt.Sprint(q.VersionEffective),    // version effective
			fmt.Sprint(q.VersionMax),          // number of versions
		}
		for iPg := 0; iPg < maxPages; iPg++ {
			if iPg < len(finishes) {
				prepend = append(prepend, finishes[iPg])
			} else {
				prepend = append(prepend, "n.a.") // response had less than max pages - not finishing time
			}
		}
		vs = append(prepend, vs...)
		valsByQ = append(valsByQ, vs) // appending to _all_ responses

	} // forr questionnaires
	// all responses have been collected
	//

	//
	// Post-processing of collected keys, vals
	//
	// Keys...
	allKeysSuperset := Superset(keysByQ)

	// Since the ordering from sparse data of Superset() is not perfect
	// we resort to alphanumeric ordering - relying on the good sequential naming of the HTML inputs
	lnStatic := len(staticCols)
	pStatic, pSorted := allKeysSuperset[:lnStatic], allKeysSuperset[lnStatic:]
	log.Print("pStatic ends   with ", pStatic[len(pStatic)-1])
	log.Print("pSorted starts with ", pSorted[0])

	if cfgRem.SurveyType == "pds" {
		// custom sort columns for PDS survey

		// standard "pre-sort" - necessary for custom sort below to work
		sort.Strings(pSorted)
		// actual custom sort
		sort.SliceStable(
			pSorted,
			func(i, j int) bool {
				// q always before other field names
				if pSorted[i][0] == 'q' && pSorted[j][0] != 'q' {
					return true
				}
				// *_main always first among equally prefixed names
				splI := strings.Split(pSorted[i], "_")
				splJ := strings.Split(pSorted[j], "_")
				lstI := len(splI) - 1
				lstJ := len(splJ) - 1
				if lstI > 0 && lstJ > 0 {
					prefI := strings.Join(splI[:lstI], "_")
					prefJ := strings.Join(splJ[:lstJ], "_")
					samePfx := prefI == prefJ
					distSfx := splI[lstI] == "main" && splJ[lstJ] != "main"
					if samePfx && distSfx {
						return true
					}
				}
				return pSorted[i] < pSorted[j]
			},
		)
	}

	allKeysSuperset = append(pStatic, pSorted...)

	positionByName := map[string]int{} // looking up the ordering/sequence number of a key by its name
	for idx, v := range allKeysSuperset {
		positionByName[v] = idx
	}
	for colIdx, colName := range allKeysSuperset {
		log.Printf("\tcol %2v  %v", colIdx, colName)
	}

	// Values...
	valsBySuperset := [][]string{}
	for i1 := 0; i1 < len(valsByQ); i1++ {
		keys := keysByQ[i1]
		vals := valsByQ[i1]
		valsBySuperset = append(valsBySuperset, make([]string, len(allKeysSuperset)))
		for i2 := 0; i2 < len(keys); i2++ {
			v := vals[i2]
			k := keys[i2]
			pos := positionByName[k]
			valsBySuperset[i1][pos] = v
		}
	}

	// Data into CSV matrix...
	var wtr = new(bytes.Buffer)
	csvWtr := csv.NewWriter(wtr)
	csvWtr.Comma = ';'
	if err := csvWtr.Write(allKeysSuperset); err != nil {
		return fnCSV, fmt.Errorf("error writing header line to csv: %w", err)
	}
	for _, record := range valsBySuperset {
		if err := csvWtr.Write(record); err != nil {
			return fnCSV, fmt.Errorf("error writing record to csv: %w", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	csvWtr.Flush()
	if err := csvWtr.Error(); err != nil {
		return fnCSV, fmt.Errorf("error flushing csv to response writer: %w", err)
	}

	err := cloudio.WriteFile(fnCSV, wtr, 0644)
	if err != nil {
		return fnCSV, fmt.Errorf("could not write CSV file %v: %v", fnCSV, err)
	}

	//
	//
	//
	// fnCreateLabels creates a separate CSV file with labels for input fields.
	// Implemented as a closure, in order to break processing with least nested conditions
	fnCreateLabels := func() {

		if len(qs) < 1 {
			return
		}

		nams := []string{} // input names
		lbls := []string{} // input labels

		fnCore := cfgRem.SurveyType + "-" + cfgRem.WaveID    // fmt-2023-01
		pthBase := path.Join(qst.BasePath(), fnCore+".json") // ./responses/fmt-2023-01.json
		qBase, err := qst.Load1(pthBase)
		if err != nil {
			log.Printf("loading base questionnaire error %v", err)
			return
		}

		// enclosing every cell value in double quotes allows to include newlines
		// excelWindowsNewline is the inside cell newlince character for Excel under Windows
		// excel newline for windows - inside cells
		const excelNL = string(rune(int32(10)))

		// copy(staticLabels, staticCols)
		byNames, _, _ := qBase.LabelsByInputNames()
		for _, name := range allKeysSuperset {
			nams = append(nams, name)
			if lbl, ok := byNames[name]; ok {
				if !strings.HasPrefix(lbl, excelNL) {
					lbl += excelNL
				}
				lbl = "\"" + strings.ReplaceAll(lbl, " -- ", excelNL) + "\""
				lbls = append(lbls, lbl)
			} else {
				lbls = append(lbls, name)
			}
		}

		buf := &bytes.Buffer{}
		buf.WriteString(strings.Join(nams, ";"))
		buf.WriteString("\n")
		buf.WriteString(strings.Join(lbls, ";"))

		fnLabels := strings.ReplaceAll(fnCSV, ".csv", "-labels.csv")
		err = cloudio.WriteFile(fnLabels, buf, 0644)
		if err != nil {
			log.Printf("writing labels file failed: %v - error %v", fnLabels, err)
		}

	}
	fnCreateLabels()

	//
	log.Printf(
		"\n\nRegular finish. %v questionnaire(s) processed\n%v non empty - %v empty\nresults in %v\n\n", len(qs),
		nonEmpty, empty, fnCSV,
	)

	return fnCSV, nil

}
