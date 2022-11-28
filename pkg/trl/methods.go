// Package trl is a central store for application environment translations;
// While specific objects such as QuestionnaireT contain translations for its contents,
// we need some global store too.
package trl

import (
	"fmt"
)

// Tr translates by key.
// Defaults           to english.
// Defaults otherwise to first lang code.
// Returns a 'signifiers string' if no translation exists,
// to help to uncover missing translations.
func (s S) Tr(langCode string) string {
	if val, ok := s[langCode]; ok {
		return val
	}
	if val, ok := s[LangCodes[0]]; ok {
		return val
	}
	if s == nil {
		return noTrans
	}
	return noTrans
}

// TrSilent gives no warning - if the translation is not set.
// Good if we do not require a translation string.
// Good for i.e. HTML title attribute - where errors would be overlooked anyway.
func (s S) TrSilent(langCode string) string {
	ret := s.Tr(langCode)
	if ret == noTrans {
		return ""
	}
	return ret
}

// String is the default "stringer" implementation
// Similar to Tr() - except third try: take whatever exists
func (s S) String() string {
	if val, ok := s["en"]; ok {
		return val
	}
	if val, ok := s[LangCodes[0]]; ok {
		return val
	}
	for _, val := range s {
		return val
	}
	return ""
}

// All returns all translations
// with their % placeholders filled in by param args;
// ordered by lang codes
// separated by '\n\n'
func (s S) All(args ...string) string {

	argsIntf := make([]interface{}, len(args))
	for i, v := range args {
		argsIntf[i] = v
	}

	ret := ""
	for _, key := range LangCodes {
		if val, ok := s[key]; ok {
			val = fmt.Sprintf(val, argsIntf...)
			ret += val
			ret += "\n\n"
		}
	}
	return ret
}

// Set checks whether s is nil or has no keys
func (s S) Set() bool {
	if len(s) < 1 { // also covers s == nil
		return false
	}
	return true
}

// Empty checks whether s has only empty translations;
// trl.S{} creates one;
func (s S) Empty() bool {
	for _, loc := range s {
		if len(loc) > 0 {
			return false
		}
	}
	return true
}

// Left trims
func (s S) Left(max int) S {
	ret := S{}
	for key, val := range s {
		if len(val) <= max {
			ret[key] = val
			continue
		}
		ret[key] = val[0:max]
	}
	return ret
}

// Bold  encloses in <b> ... </b>
func (s S) Bold() S {
	ret := S{}
	for key, val := range s {
		ret[key] = fmt.Sprintf("<b>%v</b>", val)
	}
	return ret
}

// Outline prepends <b> pref </b>
func (s S) Outline(pref string) S {
	ret := S{}
	for key, val := range s {
		ret[key] = fmt.Sprintf("<b>%v</b> &nbsp; %v", pref, val)
	}
	return ret
}

// Pad with &nbsp;
func (s S) Pad(spaces int) S {
	ret := S{}
	if spaces < 1 {
		return s
	}
	// cloning
	for k, v := range s {
		ret[k] = v
	}
	// padding
	for i := 0; i < spaces; i++ {
		for k, v := range ret {
			ret[k] = fmt.Sprintf("&nbsp;%v&nbsp;", v)
		}
	}
	return ret
}

// Fill fills in the %v placeholders;
// S is a map - thus always a pointer
// thus we dont change the receiver,
// but return a clone
func (s S) Fill(args ...interface{}) S {
	ret := S{}
	// cloning
	for k, v := range s {
		ret[k] = v
	}
	for k, v := range ret {
		ret[k] = fmt.Sprintf(v, args...)
	}
	return ret
}
