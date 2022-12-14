// Adding up columns - number of transactions, consisting of addend-a, addend-b and -c.
// Addends must be equal to sum.
// Or addends must be lighter equal to sum.

function funcInner{{.InpMain}}(){

    // let inp1 = document.forms.frmMain.xx_main.value;
    // let totalInp = document.forms.frmMain.{{.InpMain}};
    let nameTotal = "{{.InpMain}}" + "_main";
    let totalInp = document.getElementById(nameTotal);
    if (totalInp) {
        // 
    } else {
        console.log(nameTotal + " does not exist");
        return;
    }

    let totalInpVal = totalInp.value;
    let totalInpFloat = 0.0;
    let virtual = false;
    if (totalInpVal != "") {
        totalInpFloat = parseInt(totalInpVal, 10);
        totalInpFloat = parseFloat(totalInpVal);
        // alert(nameMain + " value " + iSumStr + "; " + iS);
    } else {
        let ph = totalInp.getAttribute('placeholder');
        if (ph == "100") {
            totalInpFloat = 100.0;
            virtual = true;
        }
    }

    // let summandNames = ["name1", "name2"];
    let summandNames = [{{.SummandNames}}];
    // let summandVals  = [1, 2];
    let summandValsStr  = [];
    let summandValsInt  = [];
    let sum = 0;


    for (let i1 = 0; i1 < summandNames.length; i1++) {
        const inpLp = document.getElementById( summandNames[i1] );
        summandValsStr.push(inpLp.value);        
        if (inpLp.value != "") {
            // let iVal = parseInt(inpLp.value, 10);
            let iVal = parseFloat(inpLp.value);
            summandValsInt.push( iVal);
            sum += iVal;
        } else {
            summandValsInt.push(0);
        }
    }
    
    // prevent 0.30000000004
    sum = Math.round(sum * 10000) / 10000;


    let suspicious = false;
    
    // parts adding up
    if (sum != 0 || totalInpFloat != 0) {
        if (totalInpFloat == 100 && sum == 0.0 && virtual) {
            // not suspicious

        } else if (sum {{.Operator}} totalInpFloat) {

            console.log("total:    ", nameTotal, totalInpFloat);
            console.log("summands str: ", summandValsStr);
            console.log("summands int: ", summandValsInt, " = " , sum);
            
            totalInp.focus();

            suspicious = true;
        }
    }

    return suspicious;

}

function funcOuter{{.InpMain}}(event) {

    if (funcInner{{.InpMain}}()) {
        // alert("{{.msg}}");
        // let doContinue = window.confirm("{{.msg}} {{.InpMain}}");
        let doContinue = window.confirm("{{.msg}}");
        if (doContinue) {
            return true;
        }
        event.preventDefault(); // not only return false - but also preventDefault()

        return false;
    }

    return true;

}

// non global block
{
    let frm = document.forms.frmMain;
    if (frm) {
        frm.addEventListener('submit', funcOuter{{.InpMain}});
    }
    console.log("   funcOuter{{.inp_1 }} registered")
}
