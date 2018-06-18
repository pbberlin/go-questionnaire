package fmt

import (
	"fmt"
	"time"

	"github.com/zew/go-questionaire/trl"
)

func nextWaveID() string {
	t := time.Now()
	if t.Day() > 20 {
		t = t.AddDate(0, 1, 0)
	}
	return t.Format("2006-01")
}

// Yields current quarter plus one
func nextQ() string {
	t := time.Now()
	m := t.Month() // 1 - january
	y := t.Year()
	qNow := int((m-1)/3) + 1 // jan: int(0/3)+1 == 1   feb: int(1/3)+1 == 1    mar: int(2/3)+1 == 1     apr: int(3/3)+1 == 2
	qNext := qNow + 1
	if qNext > 4 {
		qNext = 1
		y++
	}
	return fmt.Sprintf("Q%v %v", qNext, y)
}

func nextY() string {
	t := time.Now()
	y := t.Year()
	y++
	return fmt.Sprintf("%v", y)
}

func labelsGoodBad() []trl.S {

	tm := []trl.S{
		{
			"de": "gut",
			"en": "good",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "schlecht",
			"en": "bad",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsImproveDeteriorate() []trl.S {

	tm := []trl.S{
		{
			"de": "verbessern",
			"en": "improve",
		},
		{
			"de": "nicht verändern",
			"en": "not change",
		},
		{
			"de": "verschlechtern",
			"en": "deteriorate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsIncreaseDecrease() []trl.S {

	tm := []trl.S{
		{
			"de": "steigen",
			"en": "increase",
		},
		{
			"de": "gleich bleiben",
			"en": "remain unchanged",
		},
		{
			"de": "sinken",
			"en": "decrease",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsVeryHighVeryLow() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr hoch",
			"en": "very high",
		},
		{
			"de": "hoch",
			"en": "high",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "niedrig",
			"en": "low",
		},
		{
			"de": "sehr niedrig",
			"en": "very low",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsStrongIncreaseStrongDecrease() []trl.S {

	tm := []trl.S{
		{
			"de": "stark steigen",
			"en": "strongly increase",
		},
		{
			"de": "steigen",
			"en": "increase",
		},
		{
			"de": "gleich bleiben",
			"en": "remain unchanged",
		},
		{
			"de": "sinken",
			"en": "decrease",
		},
		{
			"de": "stark sinken",
			"en": "strongly decrease",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsVeryPositiveVeryNegative() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr positiv",
			"en": "very positive",
		},
		{
			"de": "positiv",
			"en": "positive",
		},
		{
			"de": "neutral",
			"en": "neutral",
		},
		{
			"de": "negativ",
			"en": "negative",
		},
		{
			"de": "sehr negativ",
			"en": "very negative",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsOvervaluedFairUndervalued() []trl.S {

	tm := []trl.S{
		{
			"de": "überbewertet",
			"en": "overvalued",
		},
		{
			"de": "fair bewertet",
			"en": "fair valued",
		},
		{
			"de": "unterbewertet",
			"en": "undervalued",
		},
	}

	return tm

}