// Package generators contains packages creating particular questionnaires.
// 2025-02: since go workspaces - the package "generators" and "fmt" were not recognized.
// "fmt" was renamed to "fmtest".
// I am hesitating to rename "generators"
package generators

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/generators/biii"
	"github.com/zew/go-questionnaire/pkg/generators/example"
	"github.com/zew/go-questionnaire/pkg/generators/flit"
	"github.com/zew/go-questionnaire/pkg/generators/fmtest"
	"github.com/zew/go-questionnaire/pkg/generators/kneb1"
	"github.com/zew/go-questionnaire/pkg/generators/pat"
	"github.com/zew/go-questionnaire/pkg/generators/pat1"
	"github.com/zew/go-questionnaire/pkg/generators/pat2"
	"github.com/zew/go-questionnaire/pkg/generators/pat3"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/tpl"
)

type genT func(s qst.SurveyT) (*qst.QuestionnaireT, error)

var gens = map[string]genT{

	"fmt": fmtest.Create, // package renamed due to conflict with standard package fmt under modules

	"example": example.Create,
	"flit":    flit.Create,
	"biii":    biii.Create,

	"kneb1": kneb1.Create,

	"pat":  pat.Create,
	"pat1": pat1.Create,
	"pat2": pat2.Create,
	"pat3": pat3.Create,

	// disabled to reduce compile times
	// "pds":     pds.Create,

	// disabled, because not migrated to Version 2.0
	// "peu2018": peu2018.Create,
	// "mul":     mul.Create,
	// "euref":   euref.Create,
	// "lt2020":  lt2020.Create,
}

func sortedKeys() []string {
	ret := []string{}
	for key := range gens {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret
}

type frmT struct {
	Type     string `json:"type"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Deadline string `json:"deadline"`
	// Params    []qst.ParamT `json:"params"`
	ParamKeys []string `json:"param_keys,omitempty"`
	ParamVals []string `json:"param_vals,omitempty"`
	Submit    string   `json:"submit,omitempty"`
}

// GenerateQuestionnaireTemplates generates a questionnaire for a bespoke survey
func GenerateQuestionnaireTemplates(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s := qst.NewSurvey("placeholder") // type is modified later
	errStr := ""
	if r.Method == "POST" {
		// fmt.Fprint(w, "is POST<br>\n")
		frm := frmT{}
		dec := form.NewDecoder()
		dec.SetTagName("json") // recognizes and ignores ,omitempty
		err := dec.Decode(&frm, r.Form)
		if err != nil {
			errStr += fmt.Sprint(err.Error() + "<br>\n")
		}

		s.Type = frm.Type
		s.Year = frm.Year
		s.Month = time.Month(frm.Month)

		t, err := time.ParseInLocation("02.01.2006 15:04 CEST", frm.Deadline, cfg.Get().Loc)
		if err != nil {
			errStr += fmt.Sprint(err.Error() + "<br>\n")
		}
		wavePeriod := time.Date(s.Year, s.Month, 1, 0, 0, 0, 0, cfg.Get().Loc)
		if t.Sub(wavePeriod) > (30*24)*time.Hour ||
			t.Sub(wavePeriod) < -(10*24)*time.Hour {
			errStr += "Should the deadline not be close to the Year-Month?<br>\n"
		}
		s.Deadline = t

		newParams := []qst.ParamT{}
		for i := 0; i < len(frm.ParamKeys); i++ {
			p := qst.ParamT{}
			p.Name = frm.ParamKeys[i]
			p.Val = frm.ParamVals[i]
			newParams = append(newParams, p)
		}
		s.Params = newParams

	}

	html := s.HTMLForm(sortedKeys(), errStr)
	fmt.Fprint(w, html) // not Fprintf
	fmt.Fprintf(w, "<br>")
	//

	if r.Method != "POST" {
		fmt.Fprintf(w, "Not a POST request. Won't generate any questionnaire<br>\n")
		return
	}

	// previously generating all questionnaires
	for _, key := range sortedKeys() {
		if key != s.Type {
			continue
		}
	}

	fnc := gens[s.Type]
	q, err := fnc(s)
	if err != nil {
		fmt.Fprintf(w, "Error creating %v: %v<br>\n", s.Type, err)
		return
	}

	fn := path.Join(qst.BasePath(), s.Filename()+".json")
	err = q.Save1(fn)
	if err != nil {
		fmt.Fprintf(w, "Error saving %v: %v<br>\n", fn, err)
		return
	}
	fmt.Fprintf(w, "%v generated<br>\n", fn)

	if cfg.Get().AnonymousSurveyID == s.Type {
		fmt.Fprint(
			w,
			`<a  
				accesskey='c'  
				target='_check'
				tabindex=2 
				href='https://localhost:8083/survey/a' 
			><b>c</b>heck </a><br>`)
	}

	//
	// create empty styles-quest-[surveytype].css"
	// if it does not yet exist
	fcCreate := func(desktopOrMobile string) (bool, error) {
		siteCore, _ := tpl.SiteCore(q.Survey.Type)
		fileNameBody := desktopOrMobile + siteCore
		pth := path.Join(".", "templates", fileNameBody+".css")
		_, err := cloudio.ReadFile(pth)
		if err != nil {
			if cloudio.IsNotExist(err) {
				rdr := &bytes.Buffer{}
				err := cloudio.WriteFile(pth, rdr, 0755)
				if err != nil {
					return false, fmt.Errorf("could not create %v: %v <br>\n", pth, err)
				}
				fmt.Fprintf(w, "Done creating template %v<br>\n", pth)
				return true, nil
			}
			return false, fmt.Errorf("other error while checking for %v: %v <br>\n", pth, err)
		}
		return false, nil
	}

	// add to parsed templates
	for _, bt := range []string{"styles-quest-"} {
		ok, err := fcCreate(bt)
		if err != nil {
			fmt.Fprintf(w, "Could not generate template %v for %v<br>\n", bt, err)
			continue
		}
		if ok {
			// parse new and previous templates
			dummyReq, err := http.NewRequest("GET", "", nil)
			if err != nil {
				log.Fatalf("failed to create request for pre-loading assets %v", err)
			}
			respRec := httptest.NewRecorder()
			tpl.TemplatesPreparse(respRec, dummyReq)
			log.Printf("\n%v", respRec.Body.String())
		}
	}

	//
}

// GenerateLandtagsVariations creates 16 questionnaire templates
func GenerateLandtagsVariations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	key := "lt2020"

	for i := 0; i < 32; i++ {

		form := url.Values{}
		form.Add("type", key)
		form.Add("year", "2020")
		form.Add("month", "5")
		form.Add("deadline", "01.01.2030 00:00")
		form.Add("params[0].name", "varianten")
		form.Add("params[0].val", fmt.Sprintf("%04b", i%16))
		form.Add("params[1].name", "aboveOrBelowMedian")
		if i < 16 {
			form.Add("params[1].val", "besseren")
		} else {
			form.Add("params[1].val", "schlechteren")
		}
		form.Add("Submit", "any")

		var resp *http.Response
		var err error

		if true {
			req, err := http.NewRequest(
				"POST",
				"https://localhost:8083"+cfg.PrefTS("generate-questionnaire-templates"),
				bytes.NewBufferString(form.Encode()),
			)
			if err != nil {
				fmt.Fprintf(w, "Request creation error %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
			client := http.DefaultClient
			resp, err = client.Do(req)
			if err != nil {
				fmt.Fprintf(w, "Request execution error %v", err)
				return
			}
		} else {
			resp, err = http.PostForm(
				"https://localhost:8083"+cfg.PrefTS("generate-questionnaire-templates"),
				form,
			)
			if err != nil {
				fmt.Fprintf(w, "Request execution error %v", err)
				return
			}
		}

		defer resp.Body.Close()
		respBts, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading response body %v", err)
			return
		}

		fmt.Fprintf(w, "%s\n", respBts)

		fn := path.Join(qst.BasePath(), key+".json")
		qst, err := qst.Load1(fn)
		if err != nil {
			fmt.Fprintf(w, "Error re-loading qst for %v: %v", fn, err)
			return
		}

		fnNew := strings.ReplaceAll(fn, ".json", fmt.Sprintf("-%02v.json", i))
		qst.Save1(fnNew)

		fmt.Fprintf(w, "Iter %v - stop; resp status %v<br><br>\n", i, resp.Status)
		fmt.Fprintf(w, "<hr>\n")

	}

}
