package qst

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var not09azHyphenUnderscore = regexp.MustCompile(`[^a-z0-9\_\-]+`)

// Mustaz09Underscore tests strings for a-z, 0-9, _
func Mustaz09Underscore(s string) bool {
	if not09azHyphenUnderscore.MatchString(s) {
		return false
	}
	return true
}

// Either no translation - or all lcs must be set
func plausibleTranslation(key string, s trl.S, lcs []string) error {

	if !s.Set() {
		// log.Printf("%-20v completely empty for %v", key, lcs)
		return nil
	}

	allElementsEmpty := true
	for _, lc := range lcs {
		if strings.TrimSpace(s[lc]) != "" {
			allElementsEmpty = false
			break
		}
	}

	if allElementsEmpty {
		// log.Printf("%-20v has only empty strings for %v", key, lcs)
		return nil
	}

	for _, lc := range lcs {
		if strings.TrimSpace(s[lc]) == "" {
			return fmt.Errorf("%-20v translation for %v is missing (%v)", key, lc, s.String())
		}
		// log.Printf("%-20v - %10v - %v", key, lc, strings.TrimSpace(s[lc]))
	}

	// log.Printf("%-20v - all translations set for %v", key, lcs)
	return nil

}

// TranslationCompleteness tests all multilanguage strings for completeness.
// Use only at JSON creation time, since dynamic elements have only one language.
func (q *QuestionnaireT) TranslationCompleteness() error {
	for i1 := 0; i1 < len(q.Pages); i1++ {
		if err := plausibleTranslation(fmt.Sprintf("page%v_sect", i1), q.Pages[i1].Section, q.LangCodes); err != nil {
			log.Print(err)
			return err
		}
		if err := plausibleTranslation(fmt.Sprintf("page%v_lbl", i1), q.Pages[i1].Label, q.LangCodes); err != nil {
			log.Print(err)
			return err
		}
		if err := plausibleTranslation(fmt.Sprintf("page%v_desc", i1), q.Pages[i1].Desc, q.LangCodes); err != nil {
			log.Print(err)
			return err
		}
		if err := plausibleTranslation(fmt.Sprintf("page%v_short", i1), q.Pages[i1].Short, q.LangCodes); err != nil {
			log.Print(err)
			return err
		}
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				if err := plausibleTranslation(fmt.Sprintf("page%v_grp%v_inp%v_lbl", i1, i2, i3), q.Pages[i1].Groups[i2].Inputs[i3].Label, q.LangCodes); err != nil {
					log.Print(err)
					return err
				}
			}
		}
	}
	return nil
}

func (q *QuestionnaireT) SetColspans() {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				inp := q.Pages[i1].Groups[i2].Inputs[i3]

				// textblock  =>  span at least 1
				if inp.Type == "textblock" {
					if inp.ColSpanLabel == 0 {
						q.Pages[i1].Groups[i2].Inputs[i3].ColSpanLabel = 1
					}
				}
				// input label or desc not empty  =>  span > 0
				if (!inp.Label.Empty() || !inp.Desc.Empty()) && inp.ColSpanLabel == 0 {
					q.Pages[i1].Groups[i2].Inputs[i3].ColSpanLabel = 1
					if inp.Type == "label-as-input" || inp.Type == "button" {
						q.Pages[i1].Groups[i2].Inputs[i3].ColSpanLabel = 0
					}
				}

				// button has label - but never colspanlabel
				// we should create a special label for button?
				if inp.Type == "button" {
					q.Pages[i1].Groups[i2].Inputs[i3].ColSpanLabel = 0
				}

				if !inp.IsLabelOnly() && !inp.IsHidden() {
					if inp.ColSpanControl == 0 {
						if inp.Label == nil {
							inp.ColSpanControl = 1
						}
					}
				}

			}
		}
	}
}

// Validate performs integrity tests - suitable for every request
//
//	waveId, langCodes valid?
//	input type valid?
//	submit button jump page exists
//	validator func exists?
//	input names uniqueness?
//
// Validate also does some initialization stuff - needed only at JSON creation time
//
//	Setting page and group width to 100
//	Setting values for radiogroups
//	Setting navigation sequence enumeration values
func (q *QuestionnaireT) Validate() error {

	if q.Survey.Type == "" || !Mustaz09Underscore(q.Survey.Type) {
		s := fmt.Sprintf("WaveID must contain a SurveyID string consisting of lower case letters: %v", q.Survey.Type)
		log.Print(s)
		return fmt.Errorf(s)
	}

	for _, lc := range q.LangCodes {
		if _, ok := cfg.Get().Mp["lang_"+lc]; !ok {
			s := fmt.Sprintf("LangCodes val %v is not a key in cfg.Get().Mp['lang_...']", lc)
			log.Print(s)
			return fmt.Errorf(s)
		}
	}

	// Check inputs
	// Set page and group width to 100
	// Set values for radiogroups
	// Enumerate pages being in navigation sequence
	for i1 := 0; i1 < len(q.Pages); i1++ {

		// navigation function exists?
		naviKey := q.Pages[i1].NavigationCondition
		if naviKey != "" {
			if _, ok := naviFuncs[naviKey]; !ok {
				return fmt.Errorf("navigator func '%v' is not in %v ", naviKey, naviFuncs)
			}
		}

		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {

			// a number of columns per group must be set
			if q.Pages[i1].Groups[i2].Cols < 1 {
				return fmt.Errorf("page %v - group %v - number of columns must be greater 0: ", i1, i2)
			}

			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				s := fmt.Sprintf("Page %v - Group %v - Input %v - %8v: ", i1, i2, i3, inp.Name)

				if inp.Type == "textblock" {
					if inp.ColSpanControl > 0 {
						return fmt.Errorf("%v: textblock should not have ColSpanControl > 0", s)
					}
					if inp.Name != "" {
						return fmt.Errorf("%v: Type '%v' - no 'name' for textblock inputs ", s, inp.Type)
					}
				}

				// check input type
				if _, ok := implementedTypes[inp.Type]; !ok {
					return fmt.Errorf("%v: Type '%v' is not in %v ", s, inp.Type, implementedTypes)
				}

				// number inputs
				if inp.Type == "number" {
					if inp.Max-inp.Min <= 0 {
						return fmt.Errorf("%v: max - min needs to be positive", s)
					}
				}

				if inp.Type == "text" || inp.Type == "number" || inp.Type == "textarea" || inp.Type == "dropdown" {
					if inp.MaxChars < 1 {
						return fmt.Errorf("%v: MaxChars setting required", s)
					}
				}

				// jump to page exists?
				if inp.Type == "button" && inp.Response != "" {
					pgIdx, err := strconv.Atoi(inp.Response)
					if err != nil {
						return fmt.Errorf(s+" %w", err)
					}
					if pgIdx < 0 || pgIdx > len(q.Pages)-1 {
						return fmt.Errorf("%v points to page index non existent %v out of %v", s, inp.Response, len(q.Pages))
					}
				}

				// validator function exists
				if inp.Validator != "" {
					valiKeys := strings.Split(inp.Validator, ";")
					for _, valiKey := range valiKeys {
						valiKey = strings.TrimSpace(valiKey)
						if _, ok := validators[valiKey]; !ok {
							return fmt.Errorf(s + fmt.Sprintf("%v - validator '%v' is not in %v ", s, valiKey, validators))
						}
					}
				}

				if inp.Type == "radio" {
					if inp.ValueRadio == "" {
						// missing ValueRadio should be caught by non-unique inputs
						return fmt.Errorf(s + fmt.Sprintf("%v - must have a distinct ValueRadio", s))
					}
				}

			}
		}
	}

	// preflight for composite funcs
	// make sure, input names are unique
	// 		multiple radios are distinguished by their radioval
	// 		multiple javascript-blocks are distinguished by JSBlockStrings
	names := map[string]int{}

	// prevent doubles of radios and non-radios
	namesRadio := map[string]int{}

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {

			if _, ok, err := q.HasComposit(i1, i2); ok {
				if err != nil {
					return fmt.Errorf(
						`page %v - group %v - composite func %v
						err %w
						`,
						i1, i2,
						q.Pages[i1].Groups[i2].Inputs[0].DynamicFunc,
						err,
					)
				}
			}

			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				s := fmt.Sprintf("page %v - group %v - input %v: ", i1, i2, i3)
				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				// grp := q.Pages[i1].Elements[i2].Name

				if inp.IsLayout() {
					continue
				}

				if inp.Type == "javascript-block" {
					continue
				}

				if inp.IsReserved() {
					return fmt.Errorf(s+"Name '%v' is reserved", inp.Name)
				}

				nm := inp.Name

				// validity of nm
				if nm == "" {
					return fmt.Errorf(s+"Name %v is empty", nm)
				}
				if not09azHyphenUnderscore.MatchString(nm) {
					return fmt.Errorf(s+"Name %v must consist of [a-z0-9_-]", nm)
				}

				// additional checks of ValueRadio
				if inp.Type == "radio" {
					if inp.ValueRadio == "" {
						return fmt.Errorf(s+"Name %v has empty radioval", nm)
					}
					if not09azHyphenUnderscore.MatchString(inp.ValueRadio) {
						return fmt.Errorf(s+"Name %v - radioval %v must consist of [a-z0-9_-]", nm, inp.ValueRadio)
					}
					nmRadio := inp.Name + "__radioval__" + inp.ValueRadio // distinguish...
					names[nmRadio]++
				}

				if inp.Type != "radio" {
					names[nm]++
				} else {
					namesRadio[nm]++ // multiple
				}

			}
		}
	}

	for k, v := range names {
		if v > 1 {
			s := fmt.Sprintf("page element '%v' is not unique  (%v)", k, v)
			log.Print(s)
			return fmt.Errorf(s)
		}
		if k != strings.ToLower(k) {
			s := fmt.Sprintf("page element '%v' is not lower case  (%v)", k, v)
			log.Print(s)
			return fmt.Errorf(s)
		}
		if _, ok := namesRadio[k]; ok {
			s := fmt.Sprintf("page element '%v' input as radio and non-radio (%v)", k, v)
			log.Print(s)
			return fmt.Errorf(s)
		}
	}
	return nil
}

// ComputeDynamicContent computes elements of type dynamic func
func (q *QuestionnaireT) ComputeDynamicContent(idx int) error {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if i1 != idx {
			continue
		}
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				if q.Pages[i1].Groups[i2].Inputs[i3].Type == "dyn-textblock" {
					inp := q.Pages[i1].Groups[i2].Inputs[i3]
					if _, ok := dynFuncs[inp.DynamicFunc]; !ok {
						return fmt.Errorf("'%v' points to dynamic func '%v()' - which does not exist or is not registered", inp.Name, inp.DynamicFunc)
					}
					str, err := dynFuncs[inp.DynamicFunc](q, inp, inp.DynamicFuncParamset)
					if err != nil {
						return fmt.Errorf("'%v' points to dynamic func '%v()' - which returned error %v", inp.Name, inp.DynamicFunc, err)
					}
					q.Pages[i1].Groups[i2].Inputs[i3].Label = trl.S{q.LangCode: str}
					// log.Printf("'%v' points to dynamic func '%v()' - which returned '%v'", i.Name, i.DynamicFunc, str)
				}
			}
		}
	}
	return nil

}

// Hyphenize replaces "mittelfristig" with "mittel&shy;fristig"
// for all labels and descriptions
func (q *QuestionnaireT) Hyphenize() {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				i := q.Pages[i1].Groups[i2].Inputs[i3]
				// s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)
				// log.Printf("Hyphenize: %v", s)
				for lc, v := range i.Label {
					v = trl.HyphenizeText(v)
					q.Pages[i1].Groups[i2].Inputs[i3].Label[lc] = v
				}
				for lc, v := range i.Desc {
					v := trl.HyphenizeText(v)
					q.Pages[i1].Groups[i2].Inputs[i3].Desc[lc] = v
				}
				for lc, v := range i.Suffix {
					v := trl.HyphenizeText(v)
					q.Pages[i1].Groups[i2].Inputs[i3].Suffix[lc] = v
				}
			}
		}
	}
}

// ComputeMaxGroups computes the maximum number of groups
// and puts them into q.MaxGroups
func (q *QuestionnaireT) ComputeMaxGroups() {
	mG := 0
	for i1 := 0; i1 < len(q.Pages); i1++ {
		if len(q.Pages[i1].Groups) > mG {
			mG = len(q.Pages[i1].Groups)
		}
	}
	q.MaxGroups = mG
}
