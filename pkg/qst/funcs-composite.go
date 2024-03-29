package qst

import (
	"github.com/zew/go-questionnaire/pkg/qstcp/cpbiii"
	"github.com/zew/go-questionnaire/pkg/qstcp/cpfmt"
	"github.com/zew/go-questionnaire/pkg/qstcp/cppat"
	"github.com/zew/go-questionnaire/pkg/qstif"
)

// CompositeFuncT inputs combine challenging HTML and
// multiple inputs in complicated ways
//
// # Compare dynFuncT, validatorT
//
// A composite func returns dynamic HTML
// with session values inserted from the questionnaire
// A composite func also returns the input *names* for json generation of the questionnaire template
//
// Matching is required for
//
//	returned input names
//	input names in HTML form
//	input names to query session values from *q argument
//
// Parameters
//
//	dynamic questionnaire - filled with response values
//	sequence  idx  -  usually a visible page sequence number
//	param set idx  -  statically determined - from a slice of param sets
//	preflight      -  only return the input names for questionnaire generation; dont render HTML output
//
// Returns
//
//	rendered HTML of the group
//	slice of input names
//	error
type CompositeFuncT func(qstif.Q, int, int, bool) (string, []string, error)

// CompositeFuncs is a lookup map
var CompositeFuncs = map[string]CompositeFuncT{
	"PoliticalFoundationsPretext":            cppat.PoliticalFoundationsPretext,            // belongs to pat
	"PoliticalFoundations":                   cppat.PoliticalFoundations,                   //   ...
	"PoliticalFoundationsStatic":             cppat.PoliticalFoundationsStatic,             //   ... no input
	"PoliticalFoundationsComprehensionCheck": cppat.PoliticalFoundationsComprehensionCheck, //   ... no input
	"TimePreferenceSelf":                     cppat.TimePreferenceSelf,                     // belongs to pat
	"TimePreferenceSelfStatic":               cppat.TimePreferenceSelfStatic,               //   ... disabled
	"TimePreferenceSelfComprehensionCheck":   cppat.TimePreferenceSelfComprehensionCheck,   //   ... disabled
	"GroupPreferences":                       cppat.GroupPreferences,                       // belongs to pat
	"GroupPreferencesPOP3":                   cppat.GroupPreferencesPOP3,                   //
	"QuestForOrg":                            cpbiii.QuestForOrg,                           //
	"Special202212Q3":                        cpfmt.Special202212Q3,                        //
	"Special202303":                          cpfmt.Special202303,                          //
	"Special202403QS1":                       cpfmt.Special202403QS1,                       //
	"Special202403QS2":                       cpfmt.Special202403QS2,                       //
}
