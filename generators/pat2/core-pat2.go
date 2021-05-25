package pat2

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var partIGroupsLong = []string{
	"Eine repräsentative Gruppe deutscher Land- und Bundestagspolitiker (Gruppe %v).",
	"Eine repräsentative Gruppe deutscher Bürger (Gruppe %v).",
	`Eine Gruppe deutscher Bürger, 
				die <i>keine Politiker</i> sind, 
				die aber die <i>gleichen demographischen Eigenschaften wie Politiker</i> haben 
				(Gruppe %v). 
				Das heißt, diese Gruppe besteht z. B. zu 70&nbsp;%% aus Männern, 
				3&nbsp;%% der Mitglieder sind unter 30&nbsp;Jahre alt, 
				87&nbsp;%% der Mitglieder haben einen Hochschulabschluss 
				und 17&nbsp;%% sind alleinstehend.`,
}

var partIGroupsShort = []string{
	"pol_gr1:Ein Politiker aus Gruppe %v <br>(deutsche Land- und Bundestagspolitiker)",
	"cit_gr2:Ein Bürger aus Gruppe %v    <br>(repräsentativer deutscher Bürger)",
	"cit_gr3:Ein Bürger aus Gruppe %v    <br>(deutsche Bürger mit gleichen demographischen Eigenschaften wie die Politiker)",
}

/*
	<ul>
		<li>
			Gruppe 1: Eine repräsentative Gruppe deutscher Bürger.
		</li>
		<li>
			Gruppe 2: Eine repräsentative Gruppe deutscher Land- und Bundestagspolitiker.
		</li>
		<li>
			Gruppe 3: Eine Gruppe deutscher Bürger, die <i>keine Politiker</i> sind,
			die aber die <i>gleichen demographischen Eigenschaften wie Politiker</i> haben.

			Das heißt, Gruppe 3 besteht z. B. zu 70&nbsp;% aus Männern,
			nur 3&nbsp;% der Mitglieder sind unter 30&nbsp;Jahre alt,
			87&nbsp;% der Mitglieder haben einen Hochschulabschluss
			und nur 17&nbsp;% sind alleinstehend.
		</li>
	<ul>

*/
var partIIQuestLabels = []string{
	`
	Als die Präferenzen wie oben gegeben waren: <br>
	Was glauben Sie, wie haben sich die 10&nbsp;deutschen Land- und Bundestagspolitiker 
	(Gruppe&nbsp;%v) entschieden?
	`,

	`
	Als die Präferenzen wie oben gegeben waren: <br>
	Was glauben Sie, wie haben sich die 10&nbsp;repräsentativen deutschen Bürger  
	(Gruppe&nbsp;%v) entschieden?
	`,

	`
	Als die Präferenzen wie oben gegeben waren: <br>
	Was glauben Sie, wie haben sich die 10&nbsp;deutschen Bürger  
	(Gruppe&nbsp;%v) entschieden 

	<br>(deutsche Bürger mit gleichen demographischen Eigenschaften wie die Politiker;
	<br>also 70&nbsp;%% Männer, 3&nbsp;%% unter 30 Jahre, halb so oft alleinstehend)? 
	`,
}

// https://cloford.com/resources/charcodes/utf-8_geometric.htm
var groupIDs = []string{
	// "◈",
	"▣",
	"◉",
	"◬",
}

// Part1Intro renders
func Part1Intro(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		//
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<h3>Teil 1</h3>

				<p>
				In diesem Teil der Studie treffen Sie acht Entscheidungen 
				(und beantworten einige Fragen). 
				Nach der Erhebung werden 10&nbsp;% aller Teilnehmer*innen 
				zufällig ausgewählt. 
				Von jedem*r ausgewählten Teilnehmer*in wird eine der acht Entscheidungen zufällig bestimmt 
				und genau wie unten beschrieben umgesetzt 
				(alle unten erwähnten Personen existieren wirklich und alle Auszahlungen 
					werden wie beschrieben getätigt).				
				</p>

				<br>
				<br>
				`,
			}
		}

	}

	return nil
}

// part2Entscheidung78TwoTimesThree - helper to Part1Entscheidung78()
func part2Entscheidung78TwoTimesThree(q *qst.QuestionnaireT, pageIdx int, inpName string) error {

	page := q.EditPage(pageIdx)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<br>
				<p>
					Wenn die Präferenzen der fünf Personen wie oben gegeben sind: <br>
					Wer soll entscheiden, ob Stiftung A, B oder C die 30 € erhält? 
				</p>
			`}
		}
	}
	for idx, kv := range partIGroupsShort {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			val = fmt.Sprintf(val, groupIDs[idx])

			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q1"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.Label = lbl
			rad.ControlFirst()
			rad.ControlTop()
		}
	}

	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<p>
					Wenn die Präferenzen der fünf Personen wie oben gegeben sind: <br>
					Wer soll möglichst <i>nicht</i> entscheiden, ob Stiftung A, B oder C die 30 € erhält? 
				</p>
			`}
		}
	}
	for idx, kv := range partIGroupsShort {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 2
			gr.RandomizationSeed = 1
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			val = fmt.Sprintf(val, groupIDs[idx])

			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q2"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.Label = lbl

			rad.ControlFirst()
			rad.ControlTop()
		}
	}

	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<p style="font-size:86%">
				<b>Erläuterung:</b>


				Falls dieser Teil umgesetzt wird und somit bestimmt, 
				welche Stiftung 30&nbsp;€ erhält, 
				werden zufällig zwei der drei Gruppen ausgewählt, 
				die tatsächlich festlegen können, welche Stiftung das Geld erhält. 
				
				Die dritte Gruppe wird die Entscheidung definitiv nicht treffen. 
				
				Von den zwei Gruppen, die die Entscheidung treffen können, 
				wird jene die Entscheidung treffen, 
				die Sie gemäß Ihrer Antworten auf die letzten 
				beiden Fragen als besser erachten.	
				</p>
			`}
		}
	}

	return nil
}

// ComprehensionCheck - single question
func ComprehensionCheck(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		// loop over matrix questions

		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
				<p>
					<b>Frage</b>: <br>
					Nehmen Sie an, die Praeferenzen der Gruppenmitglieder 
					sind wie folgt gegeben:
				</p>
				`,
				}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsComprehensionCheck__0__0")
			}

		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2

			// q2
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q_comp_a"
				inp.MaxChars = 3
				inp.Min = 0
				inp.Max = 5
				inp.ColSpan = 1
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				// inp.Placeholder = trl.S{"de": "0-5"}
				inp.Label = trl.S{"de": "<b>1.</b> Wieviele Leute stufen Stiftung A als mittel ein? "}
				inp.Suffix = trl.S{"de": "[0, 1, 2, 3, 4, 5]"}
				// inp.Validator = "must"
				inp.Validator = "must;inRange10"
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q_comp_b"
				inp.MaxChars = 3
				inp.ColSpan = 1
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				// inp.Placeholder = trl.S{"de": "A,B oder C"}
				inp.Label = trl.S{"de": "<b>2.</b> Welche Stiftung wird von drei Leuten als am besten eingestuft? "}
				inp.Suffix = trl.S{"de": "[A, B, C]"}
				// inp.Validator = "inRange1000"
				inp.Validator = "must"
			}
		}

	}

	return nil
}

// Part2IntroUndEntscheidung78 module - calls Part1Entscheidung78TwoTimesThree
func Part2IntroUndEntscheidung78(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `

					<!-- Delegation  -->
				<p>
					<b>Teil 2</b>
				</p>

				<p>
					Zuletzt haben Sie entschieden, 
					wie die Präferenzen von fünf Personen 
					in eine gemeinsame Entscheidung zusammengefasst werden sollen. 

					Dadurch haben Sie festgelegt, welche politische Stiftung 
					eine Spende von 30&nbsp;€ erhält. 
				</p>

				<p>
					Im Folgenden entscheiden Sie, 
					an wen Sie diese Entscheidung delegieren möchten 
					statt selber zu entscheiden. 
					
					Die von Ihnen ausgewählte Person sieht dabei ebenfalls 
					die Präferenzkonstellation der fünf deutschen Staatsangehörigen 
					aus der Vorstudie und entscheidet darauf basierend, 
					welche Stiftung die 30&nbsp;€ erhalten soll.					
				</p>

				<p>
					Sie können die Entscheidung an eine zufällig ausgewählte Person 
					aus einer der folgenden drei Gruppen delegieren: 
				</p>


				`,
				}
			}

		}

		for idx, txt := range partIGroupsLong {
			gr := page.AddGroup()
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
			gr.BottomVSpacers = 0

			gr.Cols = 1
			{
				txt = fmt.Sprintf(txt, groupIDs[idx])
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
				<ul>
				<li>
				%v
				</li>
				</ul>
				
				`, txt),
				}
			}
		}

		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
				<p>
					<b>Erläuterung</b>: 
					
					Wir haben Mitgliedern dieser drei Gruppen die gleichen Fragen gestellt
					 wie Ihnen im letzten Teil. 
					 
					 Auch die Mitglieder der drei Gruppen haben Präferenzkonstellationen 
					 von Teilnehmer*innen aus der Vorstudie gesehen 
					 und entschieden welche Organisation das Geld 
					 gegeben der Präferenzkonstellation erhalten soll. 
					 
					 Wir haben aus jeder dieser drei Gruppen jeweils ein Mitglied 
					 zufällig für Sie ausgewählt. 
					 
					 Falls dieser Teil der Studie umgesetzt wird, 
					 wird die Entscheidung der für Sie zufällig ausgewählten 
					 Personen aus der von Ihnen gewählten Gruppe bestimmen, 
					 welche Stiftung die 30&nbsp;€ erhält.
				</p>
				`,
				}
			}

		}

	}

	//
	//
	//
	// Entscheidung 7
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60
		page.ValidationFuncMsg = trl.S{"de": "Erste und zweite Antwort schließen sich aus. Wirklich fortfahren?"}
		page.ValidationFuncName = "pat2-part1-q7-8"

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<p><b>Entscheidung 7.</b></p>
					<p>
						In dieser Entscheidung sind die Präferenzen der 
						fünf Personen aus der Vorstudie wie folgt:
					</p>
				`}
			}
		}

		// loop over matrix questions
		for i := 0; i < 1; i++ {

			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1

				{
					inp := gr.AddInput()
					inp.Type = "dyn-composite"
					inp.ColSpanControl = 1
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsStatic__%v__%v", i, i)
				}

			}
		}

		pageIdx := len(q.Pages) - 1
		part2Entscheidung78TwoTimesThree(q, pageIdx, "dec7")

	}

	//
	// Entscheidung 8
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60
		page.ValidationFuncMsg = trl.S{"de": "Erste und zweite Antwort schließen sich aus. Wirklich fortfahren?"}
		page.ValidationFuncName = "pat2-part1-q7-8"

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<p><b>Entscheidung 8.</b></p>
					<p>
						Nun sind die Präferenzen der fünf Personen aus der Vorstudie wie folgt:
					</p>
				`}
			}
		}

		// loop over matrix questions
		for i := 1; i < 2; i++ {

			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1

				{
					inp := gr.AddInput()
					inp.Type = "dyn-composite"
					inp.ColSpanControl = 1
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsStatic__%v__%v", i, i)
				}

			}
		}

		pageIdx := len(q.Pages) - 1
		part2Entscheidung78TwoTimesThree(q, pageIdx, "dec8")

	}

	return nil
}

// Part3Intro renders
func Part3Intro(q *qst.QuestionnaireT) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<p>
					<!-- beliefs -->
					<b>Teil 3</b>
				</p> 

				<p>
					In diesem Teil der Studie beantworten Sie sechs Fragen. 
					Nach der Erhebung werden 10&nbsp;% aller Teilnehmer*innen zufällig ausgewählt. 
					Jede*r ausgewählte Teilnehmer*in wird in Abhängigkeit der Genauigkeit 
					seiner Antworten eine Bonuszahlung von bis zu 50&nbsp;Norstat&nbsp;Coins erhalten (Wert: 5&nbsp;Euro).
				</p>

				<p style='padding-bottom: 0; padding-top: 0.5rem'>
					Es geht in diesem Teil wieder um die drei Gruppen aus dem letzten Teil:
				</p>

				`,
			}
		}
	}
	for idx, txt := range partIGroupsLong {
		gr := page.AddGroup()
		gr.RandomizationGroup = 1
		gr.RandomizationSeed = 1
		gr.BottomVSpacers = 0

		gr.Cols = 1
		{
			txt = fmt.Sprintf(txt, groupIDs[idx])
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
				<ul>
				<li>
				%v
				</li>
				</ul>
				
				`, txt),
			}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<br>
				<p>
					Wir bitten Sie zu schätzen, welche Stiftung 10 zufällig ausgewählte Mitglieder
					 dieser Gruppen als Empfänger der 30&nbsp;€ bestimmt haben, 
					 wenn sie bestimmte Präferenzkonstellationen aus der Vorstudie gesehen haben.

				</p>

				<p>
					Falls Sie für die Bonuszahlung ausgewählt werden, 
					werden wir eine der sechs folgenden Fragen zufällig auswählen 
					und Ihre Schätzung bei dieser Frage mit den echten Entscheidungen 
					der Gruppenmitglieder bei der jeweiligen Präferenzkonstellation abgleichen. 
					
					Ihre Bonuszahlung ist umso höher, je genauer Ihre Schätzung ist. 
					Bitte überlegen Sie sich Ihre Antworten daher sehr genau!
				</p>

				<br>

				<p style="font-size:86%">
					<b>Erläuterung:</b>
					Falls Sie in der ausgewählten Frage eine 100 % richtige Antwort geben, 
					werden Sie 50&nbsp;Norstat&nbsp;coins erhalten. 
					Für jede Person, die Sie bei Ihren folgenden Schätzungen zu viel oder zu wenig angeben, 
					werden Sie 2.5&nbsp;Norstat&nbsp;coins verlieren. 
					Falls beispielsweise alle 10&nbsp;Gruppenmitglieder Stiftung&nbsp;C wählten, 
					Sie aber angeben, dass 5&nbsp;Gruppenmitglieder Stiftung&nbsp;B wählen, 
					und weitere 5 Stiftung&nbsp;C wählten, 
					dann haben Sie für Stiftung&nbsp;C fünf Gruppenmitglieder zu wenig angegeben, 
					und für Stiftung&nbsp;B fünf zu viel. 
					Entsprechend wird Ihre Bezahlung auf 50-2.5× 5 -2.5× 5=25&nbsp;Norstat&nbsp;coins gesenkt.
				</p>

				`,
			}
		}

	}
	return nil

}

// Part3Block12 renders
// blockStart is either 0 - or 3
func Part3Block12(q *qst.QuestionnaireT, blockStart int) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	page.ValidationFuncName = "pat2-add-to-10"
	page.ValidationFuncMsg = trl.S{"de": "Wollen Sie wirklich weiterfahren, ohne dass sich Ihre Eintraege auf 10 summieren?"}

	//
	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<p>
					Schätzen Sie zunächst für die folgende Präferenzkonstellation der fünf Personen:
				</p>
			`}
			if blockStart > 0 {
				inp.Desc = trl.S{"de": `
				<p>
					Schätzen Sie als nächstes für die folgende Präferenzkonstellation der fünf Personen:
				</p>
			`}
			}
		}
	}

	// loop over matrix questions
	// blockStart is either 0 or 3
	zeroOrOne := blockStart / 3
	for i := zeroOrOne; i < zeroOrOne+1; i++ {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsStatic__%v__%v", i, i)
			}

		}
	}

	for i1 := blockStart; i1 < blockStart+3; i1++ {

		gr := page.AddGroup()
		gr.Cols = 24
		gr.BottomVSpacers = 3
		gr.RandomizationGroup = 1
		gr.RandomizationSeed = 1
		lbls := []string{"A", "B", "C"}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 24
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
					<p style='padding: 0;position: relative; top: 0.2rem;'>
						<!-- %v -->
						<!-- <b> Frage [groupID].</b> <br> -->
						<!-- <b> Frage:</b> <br> -->
						%v 
					</p>
					`, i1+1,
					fmt.Sprintf(partIIQuestLabels[i1%3], groupIDs[i1%3]),
				),
			}
		}
		for i2 := 0; i2 < 3; i2++ {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("part2_q%v_q%v", i1+1, i2+1)
			inp.MaxChars = 2
			inp.Min = 0
			inp.Max = 10
			inp.ColSpan = 8
			inp.Label = trl.S{"de": fmt.Sprintf("von 10 wählten Stiftung&nbsp;%v", lbls[i2])}
			inp.Validator = "inRange10"
			inp.ControlFirst()
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 24
			inp.Desc = trl.S{
				"de": `
					<p style='font-size:90%'>
					Ihre Antworten müssen sich auf 10 summieren.	
					</p>
					`,
			}
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "center"
		}

	}

	if blockStart > 0 {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<br>
					<p>
					<b>
						Dies ist das Ende dieser Studie. 
						Wir bedanken uns ganz herzlich für Ihre Teilnahme. 
						Falls Sie zu den zufällig ausgewählten 10% gehören, 
						werden Sie Ihre Bonuszahlung wie versprochen in den nächsten Tagen erhalten. 
					</b>
					</p>
				`}
			}
		}
	}

	return nil
}
