package pds

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

const outline2Indent = "1.2rem"

var trancheNameStyle = css.NewStylesResponsive(nil)

func init() {
	// align tranchenames
	trancheNameStyle.Desktop.StyleGridItem.JustifySelf = "start"
	trancheNameStyle.Desktop.StyleText.AlignHorizontal = "left"
	trancheNameStyle.Desktop.StyleText.FontSize = 85
	// this needs to be differentiated
	trancheNameStyle.Desktop.StyleBox.Margin = "0 0 0 1.2rem"

	// restrTextRowLabelsTop - 2.1
	// 3.1
}

// config multiple choice
type configMC struct {
	KeyLabels          string // key to a map of labels
	Cols               float32
	InpColspan         float32
	LabelBottom        bool
	DontKnow           bool
	GroupBottomSpacers int

	GroupLeftIndent string

	XDisplacements []string
}

var (
	mCh2 = configMC{
		KeyLabels:          "teamsize",
		Cols:               16,
		InpColspan:         16 / 4,
		LabelBottom:        true,
		DontKnow:           false,
		GroupBottomSpacers: 3,
	}

	mCh2a = configMC{
		KeyLabels:          "covenants-per-credit",
		Cols:               4,
		InpColspan:         1,
		LabelBottom:        false,
		DontKnow:           false,
		GroupBottomSpacers: 3,
		GroupLeftIndent:    outline2Indent,

		XDisplacements: []string{
			"1.6rem",
			"0.62rem",
			"0.62rem",
			"1.6rem",
		},
	}

	mCh3 = configMC{
		KeyLabels:   "relevance1-5",
		Cols:        10,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,
	}

	mCh4 = configMC{
		KeyLabels:       "improveDecline1-5",
		Cols:            10,
		InpColspan:      2,
		LabelBottom:     false,
		DontKnow:        false,
		GroupLeftIndent: outline2Indent,
		XDisplacements: []string{
			"1.6rem",
			"0.79rem",
			"",
			"0.79rem",
			"1.6rem",
		},
	}
	mCh5 = configMC{
		KeyLabels:   "closing-time-weeks",
		Cols:        14,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,

		// not yet
		// GroupLeftIndent: outline2Indent,

		XDisplacements: []string{
			"1.46rem",
			"1.27rem",
			"0.64rem",
			"",
			"0.64rem",
			"1.27rem",
			"1.46rem",
		},
	}
)

var assetClassesInputs = []string{
	"ac1_corplending",
	"ac2_realestate",
	"ac3_infrastruct",
}

var assetClassesLabels = []trl.S{
	{
		"en": "Corporate / direct lending",
		"de": "Corporate / direct lending",
	},
	{
		"en": "Real estate debt",
		"de": "Real estate debt",
	},
	{
		"en": "Infrastructure debt",
		"de": "Infrastructure debt",
	},
}

// strategy, strategies
var trancheTypeNamesAC1 = []string{
	"st1_senior",
	"st2_unittranche",
	"st3_subordinated",
	"st4_mezzanine", // "mezzanine_pik_other",
}
var trancheTypeNamesAC2 = []string{
	"st1_wholeloan",
	"st2_subordinated",
}
var trancheTypeNamesAC3 = []string{
	"st1_senior",
	"st2_subordinated",
}

var lblDont = trl.S{
	"de": "Don´t know",
	"en": "Don´t know",
}

var allLbls = map[string][]trl.S{
	"ac1-tranche-types": {
		{
			"en": "Senior",
			"de": "Senior",
		},
		{
			"en": "Unitranche",
			"de": "Unitranche",
		},
		{
			"en": "Subordinated",
			"de": "Subordinated",
		},
		{
			"en": "Mezzanine / PIK / other",
			"de": "Mezzanine / PIK / other",
		},
	},
	"ac2-tranche-types": {
		{
			"en": "Whole loan",
			"de": "Whole loan",
		},
		{
			"en": "Subordinated",
			"de": "Subordinated",
		},
	},
	"ac3-tranche-types": {
		{
			"en": "Senior",
			"de": "Senior",
		},
		{
			"en": "Subordinated",
			"de": "Subordinated",
		},
	},
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
			"en": "Not relevant<br>(1)",
			"de": "Not relevant<br>(1)",
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
			"en": "Potential dealbreaker<br>(5)",
			"de": "Potential dealbreaker<br>(5)",
		},
	},
	"improveDecline1-5": {
		{
			"en": "Improved",
			"de": "Improved",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "Same",
			"de": "Same",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "Declined",
			"de": "Declined",
		},
	},

	"improveWorsen1-5": {
		{
			"en": "Improve significantly",
			"de": "Improve significantly",
		},
		{
			"en": "Improve slightly",
			"de": "Improve slightly",
		},
		{
			"en": "Stay constant",
			"de": "Stay constant",
		},
		{
			"en": "Worsen slightly",
			"de": "Worsen slightly",
		},
		{
			"en": "Worsen significantly",
			"de": "Worsen significantly",
		},
	},
	"closing-time-weeks": {
		// {
		// 	"en": "below 6&nbsp;m",
		// 	"de": "below 6&nbsp;m",
		// },
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

		// {
		// 	"en": "over 18&nbsp;m",
		// 	"de": "over 18&nbsp;m",
		// },
		{
			"en": "><br>18",
			"de": "><br>18",
		},
	},

	"covenants-per-credit": {
		{
			"en": "0-1",
			"de": "0-1",
		},
		{
			// "en": "1-3",
			// "de": "1-3",
			"en": "2-3",
			"de": "2-3",
		},
		{
			// "en": "3-5",
			// "de": "3-5",
			"en": "4-5",
			"de": "4-5",
		},
		{
			"en": ">5",
			"de": ">5",
		},
	},
}

// const firstColLbl = float32(4)
// const firstColLbl = float32(2)
const firstColLbl = float32(3)

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
var suffixPercent = trl.S{
	"en": "%",
	"de": "%",
}

var suffixMillionEuro = trl.S{
	// capitalizemytitle.com/how-to-abbreviate-million/
	// "en": "million €",
	// "en": "MM €",
	"en": "mn €",
	"de": "Mio €",
}

var placeHolderNum = trl.S{
	"en": "#",
	"de": "#",
}

var placeHolderMillion = trl.S{
	"en": "million Euro",
	"de": "Millionen Euro",
}
