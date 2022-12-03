package example

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("example")
	q.Survey = s
	q.LangCodes = []string{"en", "de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Beispielumfrage", "en": "Example survey"}

	for i1 := 0; i1 < 3; i1++ {
		page := q.AddPage()
		gr := page.AddGroup()
		gr.Cols = 3
		inp := gr.AddInput()
		inp.Name = fmt.Sprintf("name%v", i1)
		inp.Type = "text"
		inp.ColSpanControl = 1
		inp.Label = trl.S{"de": "Vorname", "en": "first name"}
		inp.MaxChars = 10
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	q.SetColspans()

	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
