package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

const firstColLbl = float32(3)
const outline2Indent = "1.2rem"

var styleHeaderCols1 = css.NewStylesResponsive(nil)
var styleHeaderCols2 = css.NewStylesResponsive(nil)
var styleHeaderCols3 = css.NewStylesResponsive(nil)
var styleHeaderCols4 = css.NewStylesResponsive(nil)

func init() {

	//
	//
	styleHeaderCols1.Desktop.StyleText.FontSize = 85
	styleHeaderCols1.Desktop.StyleGridItem.JustifySelf = "start"
	styleHeaderCols1.Desktop.StyleGridItem.AlignSelf = "end"

	styleHeaderCols1.Desktop.StyleText.AlignHorizontal = "center" // good for Mezzanine..., but bad for Senior
	styleHeaderCols1.Desktop.StyleText.AlignHorizontal = "left"

	// left margin
	styleHeaderCols1.Desktop.StyleBox.Margin = "0 0 0 0.41rem"
	styleHeaderCols1.Mobile.StyleBox.Margin = " 0"
	styleHeaderCols1.Mobile.StyleBox.Padding = "0 0.25rem 0 0.25rem"
	styleHeaderCols1.Desktop.StyleBox.Width = "100%"
	styleHeaderCols1.Desktop.StyleBox.WidthMax = "6rem"

	//
	//
	styleHeaderCols2.Desktop.StyleText.FontSize = 85
	styleHeaderCols2.Desktop.StyleGridItem.JustifySelf = "center"
	styleHeaderCols2.Desktop.StyleGridItem.AlignSelf = "end"
	styleHeaderCols2.Desktop.StyleText.AlignHorizontal = "center"

	// right margin  - exclude range radio
	// bottom margin - dy from range display
	styleHeaderCols2.Desktop.StyleBox.Margin = "0 0.98rem 0.5rem 0"
	// styleHeaderCols2.Desktop.StyleBox.Margin = "0 4.18rem 0.5rem 0"
	// styleHeaderCols2.Desktop.StyleBox.Margin = "0 4.68rem 0.5rem 0"
	styleHeaderCols2.Desktop.StyleBox.Margin = "0 5.18rem 0.5rem 0"
	styleHeaderCols2.Mobile.StyleBox.Margin = "0 1.1rem 0.5rem 0"
	styleHeaderCols2.Desktop.StyleBox.Width = "100%"

	*styleHeaderCols3 = *styleHeaderCols2
	styleHeaderCols3.Desktop.StyleBox.Margin = ""
	styleHeaderCols3.Desktop.StyleBox.Position = "relative"
	styleHeaderCols3.Desktop.StyleBox.Top = "0.6rem"
	// styleHeaderCols3.Desktop.StyleBox.BackgroundColor = "lightgrey"

	*styleHeaderCols4 = *styleHeaderCols2
	styleHeaderCols4.Desktop.StyleGridItem.JustifySelf = "left"
	styleHeaderCols4.Desktop.StyleText.AlignHorizontal = "left"
	styleHeaderCols4.Desktop.StyleBox.Margin = "0 0 0.5rem 3.2rem"

}

// asset classes
type assetClass struct {
	NameUnused   string // unusued
	Prefix       string
	Lbl          trl.S
	Short        trl.S // Short label
	TrancheTypes []trancheType
}

// tranche types
// strategies
type trancheType struct {
	NameUnused, Prefix string
	Lbl                trl.S
}

var PDSAssetClasses = []assetClass{
	{
		NameUnused: "ac1_corplending",
		Prefix:     "ac1",
		Lbl: trl.S{
			"en": "Corporate / direct lending",
			"de": "Corporate / direct lending",
		},
		Short: trl.S{
			"en": "Corp. lend.",
			"de": "Corp. lend.",
		},
		TrancheTypes: []trancheType{
			{
				NameUnused: "tt1_senior",
				Prefix:     "tt1",
				Lbl: trl.S{
					"en": "Senior",
					"de": "Senior",
				},
			},
			{
				NameUnused: "tt2_unittranche",
				Prefix:     "tt2",
				Lbl: trl.S{
					"en": "Uni&shy;tranche&nbsp;&nbsp;",
					"de": "Uni&shy;tranche&nbsp;&nbsp;",
				},
			},
			{
				NameUnused: "tt3_subpikoth",
				Prefix:     "tt3",
				Lbl: trl.S{
					"en": "Subordinated / mezzanine / other",
					"de": "Subordinated / mezzanine / other",
					// "en": "Subordinated / PIK / Other",
					// "de": "Subordinated / PIK / Other",
				},
			},
		},
	},
	{
		NameUnused: "ac2_realestate",
		Prefix:     "ac2",
		Lbl: trl.S{
			"en": "Real estate debt",
			"de": "Real estate debt",
		},
		Short: trl.S{
			"en": "Real est.",
			"de": "Real est.",
		},
		TrancheTypes: []trancheType{
			{
				NameUnused: "tt1_wholeloan",
				Prefix:     "tt1",
				Lbl: trl.S{
					"en": "Whole loan",
					"de": "Whole loan",
				},
			},
			{
				NameUnused: "tt2_subordinated",
				Prefix:     "tt2",
				Lbl: trl.S{
					"en": "Subordinated",
					"de": "Subordinated",
				},
			},
		},
	},
	{
		NameUnused: "ac3_infrastruct",
		Prefix:     "ac3",
		Lbl: trl.S{
			"en": "Infrastructure debt",
			"de": "Infrastructure debt",
		},
		Short: trl.S{
			"en": "Infrastruct.",
			"de": "Infrastruct.",
		},
		TrancheTypes: []trancheType{
			{
				NameUnused: "tt1_senior",
				Prefix:     "tt1",
				Lbl: trl.S{
					"en": "Senior",
					"de": "Senior",
				},
			},
			{
				NameUnused: "tt2_subordinated",
				Prefix:     "tt2",
				Lbl: trl.S{
					"en": "Subordinated",
					"de": "Subordinated",
				},
			},
		},
	},
}

var lblDont = trl.S{
	"de": "Don??t know",
	"en": "Don??t know",
}

var PDSLbls = map[string][]trl.S{
	"teamsize": {
		{
			"en": "<5",
			"de": "<5",
		},
		{
			"en": "5-10",
			"de": "5-10",
		},
		{
			"en": "11-20",
			"de": "11-20",
		},
		{
			"en": ">20",
			"de": ">20",
		},
	},
	"relevance1-5": {
		{
			"en": "not relevant<br>(1)",
			"de": "not relevant<br>(1)",
		},
		{
			"en": "(2)",
			"de": "(2)",
		},
		{
			"en": "(3)",
			"de": "(3)",
		},
		{
			"en": "(4)",
			"de": "(4)",
		},
		{
			"en": "potential deal&shy;breaker<br>(5)",
			"de": "potential deal&shy;breaker<br>(5)",
		},
	},
	"improveDecline1-5": {
		// {
		// 	"en": "Im&shy;prov&shy;ed",
		// 	"de": "Im&shy;prov&shy;ed",
		// },
		{
			"en": "Im&shy;prove",
			"de": "Im&shy;prove",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			// yes - its terrible
			"en": "Sa&shy;me",
			"de": "Sa&shy;me",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "De&shy;cline",
			"de": "De&shy;cline",
		},
	},
	"improveworsen1-5": {
		{
			"en": "improve significantly",
			"de": "improve significantly",
		},
		{
			"en": "improve slightly",
			"de": "improve slightly",
		},
		{
			"en": "stay constant",
			"de": "stay constant",
		},
		{
			"en": "worsen slightly",
			"de": "worsen slightly",
		},
		{
			"en": "worsen significantly",
			"de": "worsen significantly",
		},
	},
	"closing-time-weeks": {
		{
			"en": "<<br>6",
			"de": "<<br>6",
		},

		{
			"en": "&nbsp;<br>6",
			"de": "&nbsp;<br>6",
		},
		{
			"en": "&nbsp;<br>9",
			"de": "&nbsp;<br>9",
		},
		{
			"en": "weeks<br>12",
			"de": "weeks<br>12",
		},
		{
			"en": "&nbsp;<br>15",
			"de": "&nbsp;<br>15",
		},
		{
			"en": "&nbsp;<br>18",
			"de": "&nbsp;<br>18",
		},
		{
			"en": "><br>18",
			"de": "><br>18",
		},
	},
	"covenants-per-credit": {
		// &#8209; - non breaking dash,
		//   to prevent
		// 			0-
		// 			1
		{
			"en": "0&#8209;1",
			"de": "0&#8209;1",
		},
		{
			"en": "2&#8209;3",
			"de": "2&#8209;3",
		},
		{
			"en": "4&#8209;5",
			"de": "4&#8209;5",
		},
		{
			"en": "&nbsp;>5",
			"de": "&nbsp;>5",
		},
	},
}

var suffixWeeks = trl.S{
	"en": "weeks",
	"de": "Wochen",
}

var suffixYears = trl.S{
	"en": "years",
	"de": "Jahre",
}

var suffixEBITDA = trl.S{
	"en": "x EBITDA",
	"de": "x EBITDA",
}

var suffixInvestedCapital = trl.S{
	"en": "x invested cap.",
	"de": "x invested cap.",
}
var suffixPercent = trl.S{
	"en": "%",
	"de": "%",
}

var suffixNumDeals = trl.S{
	"en": "deals",
	"de": "Deals",
	// "en": "transactions",
	// "de": "Transaktionen",
}

var suffixMillionEuro = trl.S{
	// capitalizemytitle.com/how-to-abbreviate-million/
	// "en": "million ???",
	// "en": "MM ???",
	"en": "mn&nbsp;???",
	"de": "Mio&nbsp;???",
}

var placeHolderNum = trl.S{
	"en": "#",
	"de": "#",
}

var placeHolderMillion = trl.S{
	"en": "million Euro",
	"de": "Millionen Euro",
}

var placeHolder100percent = trl.S{
	"en": "100",
	"de": "100",
}

func onlySelectedTranchTypes(q *QuestionnaireT, ac assetClass) assetClass {

	ln := len(ac.TrancheTypes)

	// iterate over all
	names := make([]string, 0, ln)
	for i := 0; i < ln; i++ {
		//                               ("ac1_tt1_q031")
		names = append(names, fmt.Sprintf("%v_%v_q031", ac.Prefix, ac.TrancheTypes[i].Prefix))
	}

	newTTs := make([]trancheType, 0, ln)
	for i, name := range names {
		inp := q.ByName(name)
		if inp.Response != "" && inp.Response != "0" {
			newTTs = append(newTTs, ac.TrancheTypes[i])
		}
	}

	acRet := ac
	acRet.TrancheTypes = newTTs

	// if len(acRet.TrancheTypes) == 0 {
	// 	log.Printf("  %v => %v tt(s)", ac.Prefix, len(acRet.TrancheTypes))
	// 	log.Print(util.IndentedDump(acRet))
	// }

	return acRet
}
