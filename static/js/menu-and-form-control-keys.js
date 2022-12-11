// non essentiell JS helpers
function keyControls(e) {

    // [enter] key opens  2nd level menu, just as space bar does
    if (e.key === "Enter") {
        var menuCheckbox = document.getElementById("mnu-1st-lvl-toggler");
        var isFocused = (document.activeElement === menuCheckbox);
        if (isFocused) {
            menuCheckbox.checked = true;
            console.log("key listener ENTER fired");
        }
    }

    // [esc]   key closes 2nd level menu, if its expanded
    if (e.key === "Escape") {
        document.getElementById("mnu-1st-lvl-toggler").checked = false;


        // ExcelDB: hide all control-menu-2
        // var mnu2s = document.getElementsByClassName("control-menu-2");
        // for (var i = 0; i < mnu2s.length; i++) {
        // 	mnu2s[i].style.display = 'none';
        // }
        // console.log("key listener ESC fired");
    }

    // [enter] on inputs transformed into focus next input.
    // Sending events to inputs is security forbidden.
    // We find the next element and focus() it.
    //
    // TEXTAREA: SHIFT+ENTER mode is impossible on mobile -
    // thus we cannot include TEXTAREA into the func
    //
    //	optionally restrict to certain user agents: && /Android/.test(navigator.userAgent)
    if (e.key === "Enter") {

        var isShift = !!e.shiftKey; // convert to boolean
        if (isShift) {
            console.log("let SHIFT ENTER pass");
            return;
        }

        var el = document.activeElement;

        // skip for <input type=submit>  and <button>...
        if ((el.tagName == "INPUT" && el.type != "submit") || el.tagName == "SELECT") {

            e.preventDefault();
            var nextEl = null;


            if (false) {
                // first method for finding next element:
                // adding succinct tab indize
                // then taking current tab index and incrementing it
                var elements = el.form.elements;
                var cntr = 1;
                for (var i = 0, lpEl; lpEl = elements[i++];) {
                    if (lpEl.type !== "hidden" && lpEl.type !== "fieldset") {
                        lpEl.tabIndex = cntr;
                        cntr++;
                        // console.log("tab index", element.name, " to ", i);
                    } else {
                        // console.log("SKIPPING tab index ", element.name, " - ", i);
                    }
                }
                var nextTabIndex = el.tabIndex + 1;
                nextEl = el.form.elements[nextTabIndex];
                if (nextEl && nextEl.focus) nextEl.focus();
            }


            // second method: simply follow the form elements order
            var found = false;
            if (el.form) {
                for (var i = 0, lpEl; lpEl = el.form.elements[i++];) {
                    if (lpEl.type !== "hidden" && lpEl.type !== "fieldset") {
                        if (found) {
                            nextEl = lpEl;
                            // console.log(`found next	   ${lpEl.name} type ${lpEl.type} at `, i);
                            break;
                        }
                        if (el === lpEl) {
                            // console.log(`found current ${lpEl.name} type ${lpEl.type} at `, i);
                            found = true;
                        }
                        // console.log("iterating form elements", element.name, " to ", i);
                    } else {
                        // console.log("iterating form elements - skipping ", element.name, " - ", i);
                    }
                }
            }
            if (nextEl && nextEl.focus) nextEl.focus();


            if (nextEl) {
                // console.log("key listener ENTER - transformed into TAB:", el.tagName, el.name, nextEl.tagName, nextEl.name );
            } else {
                // console.log("key listener ENTER - transformed into TAB:", el.tagName, el.name, " next element not found" );
            }

        } else {
            // console.log("key listener ENTER on tagname:", el.tagName, el.name );
        }
    }

}

// click outside menu closes it
function outsideMenu(event) {
    var elNav = document.getElementsByTagName('nav');
    var nav = elNav[0];
    // event.preventDefault();
    if (!nav.contains(event.target)) {
        // console.log('click outside menu');
        document.getElementById("mnu-1st-lvl-toggler").checked = false;
    }
}

// click on nde-2nd-lvl pulls up mnu-3rd-lvl
//
// we would love to change li.nde-2nd-lvl::before
// into an upward arrow too, but pseudo elements
// cannot be selected / styled via javascript
var closeLevel3 = function () {
    for (let i = 0; i < this.children.length; i++) {
        if (this.children[i].tagName == "UL") {
            var el = this.children[i];
            var style = window.getComputedStyle(el);
            if (style.opacity < 0.5) {
                el.classList.remove("mnu-3rd-lvl-pull-up");  // remove means *show* ;this is the show / init branch - opacity 0 and growing
            } else {
                el.classList.add("mnu-3rd-lvl-pull-up");	 // add	means *hide*
            }
            break;
        }
    }
};


// window.onload = ...   is *not* cumulative
// window.onload = function () {
//     //
// };
//
// addEventListener is cumulative
window.addEventListener("load", function (event) {

    document.addEventListener("keydown", keyControls, false);
    console.log("global key listener registered");


    var html = document.body.parentNode;
    html.addEventListener("touchstart", outsideMenu, false);
    html.addEventListener('click', outsideMenu, false);

    var nodesLvl2 = document.getElementsByClassName("nde-2nd-lvl");
    for (var i = 0; i < nodesLvl2.length; i++) {
        nodesLvl2[i].addEventListener('click', closeLevel3, false);
    }
    console.log("outsideMenu and closeLevel3 registered");


    var invalidInputs = false; // invalid by HTML5
    var invalidFields = document.querySelectorAll("form :invalid");  // excluding invalid form itself
    for (var i = 0; i < invalidFields.length; i++) {
        /*  first pages with first element after long text
                => scrolls down
            preventScroll supported only since 2018
         */
        try {
            invalidFields[i].focus({
                preventScroll: true
            });
        } catch (error) {
            // forgoing initial focussing
        }
        // console.log(`focus on first invalid input ${invalidFields[i].name}`);
        invalidInputs = true;
        break;
    }

    var invalidServerFields = document.querySelectorAll(".error-block-input"); // invalid by server rules
    var firstErrMsgTop = 0;
    if (invalidServerFields.length > 0) {
        firstErrMsgTop = invalidServerFields[0].getBoundingClientRect().y
        // console.log(`.error-block-input found at ${topPosOfErr}`);
    }


    if (!invalidInputs) {
        // focus on first visible input
        var elements = document.forms.frmMain.elements;
        for (var i = 0, el; el = elements[i++];) {
            if (el.type === "hidden") {
                continue;
            }

            if (firstErrMsgTop > 0 && el.getBoundingClientRect().y < firstErrMsgTop) {
                // console.log(`.error-block-input found ${element.getBoundingClientRect().y} < ${topPosOfErr}`);
                continue;
            }

            if (el.type === "range") {
                // if the first input is a range; 
                // then we dont want to focus - because otherwise hitting PageDown 
                // would not scroll down the page, but change the range-slider thumb
                break;
            }

            /*  first pages with first element after long text
                    => scrolls down
                preventScroll supported only since 2018
                */
            try {
                el.focus({
                    preventScroll: true
                });
            } catch (error) {
                // forgoing initial focussing
            }
            // console.log(`focus on ${i}th input ${element.name} of form main`);
            break;

        }
    }


});



// change or input for a input[type=range]
// updating   the corresponding 'display'
// unchecking the corresponding 'noanswer'
// updating   the corresponding 'hidden' input;
// if the range-input is dirty;
// otherwise we just set the 'display' to '--'
function pdsRangeInput(src){
    // console.log("rangeInput()");

    src.style.backgroundColor = "transparent";

    if (!src.parentNode) {
        return true
    }
    if (!src.parentNode.childNodes) {
        return true
    }

    let chn = src.parentNode.childNodes;
    // console.log("child nodes num", chn.length);

    let noAnsw = null;
    let label = null;
    for (i = 0; i < chn.length; i++) {

        let el = chn[i];

        if (el.nodeType == Node.TEXT_NODE) {
            // console.log("   ch #",i , " - textnode");
        } else {
            // console.log("   ch #", i, el.nodeType, el.type, el.name);
            if (el.type == "radio") {
                noAnsw = el;
            }
            if (el.tagName == "LABEL") {
                label = el;
            }
        }
    }

    if (noAnsw) {
        noAnsw.checked = false;
        // rg.disabled = true;
    }
    if (label) {

        let chn = label.childNodes;
        // console.log("  label child nodes num", chn.length);

        let display = null;
        for (i = 0; i < chn.length; i++) {

            let el = chn[i];

            if (el.nodeType == Node.TEXT_NODE) {
                // console.log("   ch #",i , " - textnode");
            } else {
                // nodeType is an integer
                // console.log("       ch #", i, el.nodeType, el.tagName, el.type, el.name);
                if (el.type == "text") {
                    display = el;
                }
            }
        }

        if (display) {

            // console.log("src.value=", src.value, "data-dirty:", src.dataset.dirty);

            if (src.dataset.dirty === "false") {
                display.value = " -- ";
                src.classList.add("hidethumb");
            } else {
                let incr = parseFloat(src.value) + parseFloat(src.step);
                let out = ""
                if (src.value) {
                    out += src.value;
                }
                out += " - ";
                out += incr;
                display.value = out;

                let catcher = document.getElementById(src.id + "_hidd");
                if (catcher) {
                    catcher.value = src.value;
                }
            }

        }

    }
    return true;
}

// activate an input[type=range] from de-activated visual state 
function pdsRangeClick(src){
    // console.log("rangeClick()");
    src.style.backgroundColor = "transparent";
    src.classList.remove("hidethumb");
    src.dataset.dirty = "";

    // src.disabled = false;
    return true;
}

// 'no answer' radio puts corresponding range-input into de-activated state
function pdsRangeRadioInput(src){
    // console.log("rangeRadioInput()");
    // console.log("rangeRadioInput() - src id", src.id);

    if (!src.parentNode) {
        return true
    }
    if (!src.parentNode.childNodes) {
        return true
    }

    const chn = src.parentNode.childNodes;
    // console.log("child nodes num", chn.length);

    let rg = null;
    let label = null;
    for (i = 0; i < chn.length; i++) {

        let el = chn[i];

        if (el.nodeType == Node.TEXT_NODE) {
            // console.log("   ch #",i , " - textnode");
        } else {
            // console.log("   ch #", i, el.nodeType, el.type, el.name);
            if (el.type == "range") {
                rg = el;
            }
            if (el.tagName == "LABEL") {
                label = el;
            }
        }
    }

    if (rg) {
        rg.style.backgroundColor = "darkgray";
        // rg.disabled = true;
        let catcherID = src.id
        catcherID = catcherID.replace("_noanswer","_hidd");
        let catcher = document.getElementById(catcherID);
        if (catcher) {
            catcher.value = "noanswer";
        } else {
            console.log("catcher not found", catcherID);
        }

    }

    if (label) {

        let chn = label.childNodes;
        // console.log("  label child nodes num", chn.length);

        let display = null;
        for (i = 0; i < chn.length; i++) {

            let el = chn[i];

            if (el.nodeType == Node.TEXT_NODE) {
                // console.log("   ch #",i , " - textnode");
            } else {
                // console.log("       ch #", i, el.nodeType, el.type, el.name);
                // input[type=text]
                if (el.type == "text") {
                    display = el;
                }
            }
        }

        if (display) {
            display.value = "no a.";
        }

    }


    return true;
}
