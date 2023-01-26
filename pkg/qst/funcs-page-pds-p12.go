package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage12AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage12(q, page, 0)
}
func pdsPage12AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage12(q, page, 1)
}
func pdsPage12AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage12(q, page, 2)
}

func pdsPage12(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]
	ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: &nbsp;&nbsp;  Loans issued in %v (continued)", ac.Lbl["en"], q.Survey.Quarter(-1)),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp;  Loans issued in %v (continued)", ac.Lbl["de"], q.Survey.Quarter(-1)),
	}.Outline(fmt.Sprintf("%c1.", rn))

	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>Changes 2", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>Changes 2", ac.Short["de"]),
	}
	page.CounterProgress = fmt.Sprintf("%c1b", rn)

	// marker for naviFuncs pds_ac1-3
	page.CounterProgress = "page12"

	page.SuppressInProgressbar = true

	page.WidthMax("58rem")
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("42rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("38rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"en": "Unlevered returns",
				"de": "Unlevered returns",
			}.Outline("1.2")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	page12Types := []string{
		"range-pct",
		"range-pct",
		"range-pct",
		"range-pct",
		"range-pct",
	}
	page12Inputs := []string{
		"q12a_cash_margin",
		"q12b_interest_floor",
		"q12c_fixed_rate_coupon",
		"q12d_upfront_fee",
		"q12e_irr_expected",
	}
	page12Lbls := []trl.S{
		{
			"en": `Average cash margin over the relevant base rate`,
			"de": `Average cash margin over the relevant base rate`,
		},
		{
			"en": `Average interest rate floor`,
			"de": `Average interest rate floor`,
		},
		{
			"en": `Average fixed rate coupon`,
			"de": `Average fixed rate coupon`,
		},
		{
			"en": `Average upfront fee`,
			"de": `Average upfront fee`,
		},
		{
			"en": `Average expected Gross IRR`,
			"de": `Average expected Gross IRR`,
		},
	}

	page12LblsDescr := []trl.S{
		{
			"de": `Pleases state the average cash margin over the relevant base rate for transactions closed in [quarter-1]. Only relevant for <i>floating rate</i> loans.`,
			"en": `Pleases state the average cash margin over the relevant base rate for transactions closed in [quarter-1]. Only relevant for <i>floating rate</i> loans.`,
		},
		{
			"en": `Please state the average interest floor for transactions closed in [quarter-1]. Only relevant for <i>floating rate</i> loans.`,
			"de": `Please state the average interest floor for transactions closed in [quarter-1]. Only relevant for <i>floating rate</i> loans.`,
		},
		{
			"en": `Please state the average fixed rate copuon for transactions closed in [quarter-1]. Only relevant for <i>fixed rate</i> loans.`,
			"de": `Please state the average fixed rate copuon for transactions closed in [quarter-1]. Only relevant for <i>fixed rate</i> loans.`,
		},
		{
			"en": `Please state the average upfront fees charged to the borrower for transactions closed in [quarter-1].`,
			"de": `Please state the average upfront fees charged to the borrower for transactions closed in [quarter-1].`,
		},
		{
			"en": `Please state the average expected Gross Internal Rate of Return (IRR) for transactions closed in [quarter-1].`,
			"de": `Please state the average expected Gross Internal Rate of Return (IRR) for transactions closed in [quarter-1].`,
		},
	}

	for i := 0; i < len(page12Lbls); i++ {
		page12Lbls[i].Append90(page12LblsDescr[i])
	}

	for i := 0; i < len(page12Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page12Lbls[i] = page12Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	createRows(
		page,
		ac,
		page12Inputs,
		page12Types,
		page12Lbls,
		[]*rangeConf{
			&range2To10,
			&range0To2a,
			&range3To20,
			&range0To4,
			&range3To25,
		},
	)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"en": "Risk",
				"de": "Risk",
			}.Outline("1.3")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	// 1.3 switch, branch
	if acIdx == 0 {

		page13Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			// "restricted-text-int",
			// "restricted-text-int",
		}
		page13Inputs := []string{
			"q13a_number_covenants",
			"q13b_contracted_maturity",
			"q13c_opening_leverage",
			"q13d_ebitda_avg",
			"q13e_ev_avg",
		}
		page13Lbls := []trl.S{
			{
				"en": `Average number of covenants`,
				"de": `Average number of covenants`,
			},
			{
				"en": `Contracted maturity`,
				"de": `Contracted maturity`,
			},
			{
				"en": `Opening leverage`,
				"de": `Opening leverage`,
			},
			{
				"en": `Average EBITDA`,
				"de": `Average EBITDA`,
			},
			{
				"en": `Average EV`,
				"de": `Average EV`,
			},
		}

		page13LblsDescr := []trl.S{
			{
				"en": `Please state the average number of financial covenants per loan for transactions closed in [quarter-1].`,
				"de": `Please state the average number of financial covenants per loan for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average contracted maturity for transactions closed in [quarter-1].`,
				"de": `Please state the average contracted maturity for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average opening leverage for transactions closed in [quarter-1]. Opening leverage is measured as a multile of EBITDA.`,
				"de": `Please state the average opening leverage for transactions closed in [quarter-1]. Opening leverage is measured as a multile of EBITDA.`,
			},
			{
				"en": `Please state the average EBITDA of borrower companies for transactions closed in [quarter-1].`,
				"de": `Please state the average EBITDA of borrower companies for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average enterprise value of borrower companies for transactions closed in [quarter-1].`,
				"de": `Please state the average enterprise value of borrower companies for transactions closed in [quarter-1].`,
			},
		}

		for i := 0; i < len(page13Lbls); i++ {
			page13Lbls[i].Append90(page13LblsDescr[i])
		}

		for i := 0; i < len(page13Lbls); i++ {
			rn := rune(97 + i) // 97 is a
			page13Lbls[i] = page13Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
		}

		createRows(
			page,
			ac,
			page13Inputs,
			page13Types,
			page13Lbls,
			[]*rangeConf{
				nil, // unused
				&range0To10,
				&rangeEBITDA2x10x,
				&rangeEBITDAZero150,
				&rangeEV0To500,
				// nil,
				// nil,
			},
		)
	}

	if acIdx == 1 {

		page13Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",

			// real estate specific
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
		}
		page13Inputs := []string{
			"q13a_number_covenants",
			"q13b_contracted_maturity",
			"q13c_opening_leverage",

			// real estate specific
			"q13d_opening_dscr",
			"q13e_opening_icr",
			"q13f_opening_debt_yield",
			"q13g_exit_leverage",
			"q13h_exit_dscr",
			"q13i_exit_icr",
			"q13j_exit_yield",
		}
		page13Lbls := []trl.S{
			{
				"en": `Average number of covenants`,
				"de": `Average number of covenants`,
			},
			{
				"en": `Contracted maturity`,
				"de": `Contracted maturity`,
			},
			{
				"en": `Opening leverage`,
				"de": `Opening leverage`,
			},

			// real estate specific
			{
				"en": `Opening DSCR`,
				"de": `Opening DSCR`,
			},
			{
				"en": `Opening ICR`,
				"de": `Opening ICR`,
			},
			{
				"en": `Opening debt yield`,
				"de": `Opening debt yield`,
			},
			{
				"en": `Expected exit leverage`,
				"de": `Expected exit leverage`,
			},
			{
				"en": `Expected exit DSCR`,
				"de": `Expected exit DSCR`,
			},
			{
				"en": `Expected exit ICR`,
				"de": `Expected exit ICR`,
			},
			{
				"en": `Expected exit yield`,
				"de": `Expected exit yield`,
			},
		}

		page13LblsDescr := []trl.S{
			{
				"en": `Please state the average number of financial covenants per loan  for transactions closed in [quarter-1].`,
				"de": `Please state the average number of financial covenants per loan  for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average contracted maturity for transactions closed in [quarter-1].`,
				"de": `Please state the average contracted maturity for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average opening LTV/LTC for transactions closed in [quarter-1].`,
				"de": `Please state the average opening LTV/LTC for transactions closed in [quarter-1].`,
			},

			// real estate specific
			{
				"en": `Please state the average opening Debt-Service Coverage Ratio (DSCR) for transactions closed in [quarter-1].`,
				"de": `Please state the average opening Debt-Service Coverage Ratio (DSCR) for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average opening Interest Coverage Ratio (ICR) for transactions closed in [quarter-1].`,
				"de": `Please state the average opening Interest Coverage Ratio (ICR) for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average average opening debt yield in [quarter-1].`,
				"de": `Please state the average average opening debt yield in [quarter-1].`,
			},
			{
				"en": `Please state the average expected LTV/LTC for transactions closed in [quarter-1].`,
				"de": `Please state the average expected LTV/LTC for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average expected Debt-Service Coverage Ratio (DSCR) for transactions closed in [quarter-1].`,
				"de": `Please state the average expected Debt-Service Coverage Ratio (DSCR) for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average expected Interest Coverage Ratio (ICR) for transactions closed in [quarter-1].`,
				"de": `Please state the average expected Interest Coverage Ratio (ICR) for transactions closed in [quarter-1].`,
			},
			{
				"en": `Please state the average expected Interest Coverage Ratio (ICR) for transactions closed in [quarter-1].`,
				"de": `Please state the average expected Interest Coverage Ratio (ICR) for transactions closed in [quarter-1].`,
			},
		}

		for i := 0; i < len(page13Lbls); i++ {
			page13Lbls[i].Append90(page13LblsDescr[i])
		}

		for i := 0; i < len(page13Lbls); i++ {
			rn := rune(97 + i) // 97 is a
			page13Lbls[i] = page13Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
		}

		createRows(
			page,
			ac,
			page13Inputs,
			page13Types,
			page13Lbls,
			[]*rangeConf{
				nil,
				&range0To10,
				&range30To100,

				// real estate specific
				&range1To5A,
				&range1To5B,
				&range2To10Experimental,
				&range30To100,
				&range1To5A,
				&range1To5B,
				&range2To10Experimental,
			},
		)
	}

	if acIdx == 2 {

		page13Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",

			// infrastruct specific
			"range-pct",
			"range-pct",
			"range-pct",
		}
		page13Inputs := []string{
			"q13a_number_covenants",
			"q13b_contracted_maturity",
			"q13c_opening_leverage",

			// infrastruct specific
			"q13d_maximum_leverage",
			"q13e_average_dscr",
			"q13f_minimum_dscr",
		}
		page13Lbls := []trl.S{
			{
				"en": `Average number of covenants`,
				"de": `Average number of covenants`,
			},
			{
				"en": `Contracted maturity`,
				"de": `Contracted maturity`,
			},
			{
				"en": `Opening leverage`,
				"de": `Opening leverage`,
			},

			// infrastruct specific
			{
				"en": `Expected maximum Leverage`,
				"de": `Expected maximum Leverage`,
			},
			{
				"en": `Expected average DSCR`,
				"de": `Expected average DSCR`,
			},
			{
				"en": `Expected minimum DSCR`,
				"de": `Expected minimum DSCR`,
			},
		}

		page13LblsDescr := []trl.S{
			{
				"en": `What is the average number of financial covenants per loan?`,
				"de": `What is the average number of financial covenants per loan?`,
			},
			{
				"en": `What is the average contracted maturity?`,
				"de": `What is the average contracted maturity?`,
			},
			{
				"en": `What is the average opening leverage?`,
				"de": `What is the average opening leverage?`,
			},

			// infrastruct specific
			{
				"en": `What is the average expected maximum LTV?`,
				"de": `What is the average expected maximum LTV?`,
			},
			{
				"en": `What is the expected average DSCR?`,
				"de": `What is the expected average DSCR?`,
			},
			{
				"en": `What is the expected minimum DSCR?`,
				"de": `What is the expected minimum DSCR?`,
			},
		}

		for i := 0; i < len(page13Lbls); i++ {
			page13Lbls[i].Append90(page13LblsDescr[i])
		}

		for i := 0; i < len(page13Lbls); i++ {
			rn := rune(97 + i) // 97 is a
			page13Lbls[i] = page13Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
		}

		createRows(
			page,
			ac,
			page13Inputs,
			page13Types,
			page13Lbls,
			[]*rangeConf{
				nil,
				&range0To10,
				&range50To100,

				// infrastruct specific
				&range50To100,
				&range1To175,
				&range1To175,
			},
		)
	}

	// here not possible, because necessary values in other pages are not yet fully populated at this point
	// pdsSpecialDisableColumns(...)

	return nil
}
