package regnet_test

import (
	"github.com/anaray/regnet"
	"testing"
	//"fmt"
)

func TestNewRegnet(t *testing.T) {
	r := regnet.New()
	if r == nil {
		t.Errorf("regnet initialization failed")
	}

	key := "%{DAY}"
	pattern := "(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)"
	r.AddPattern(key, pattern)
	value, _ := r.GetPattern(key)

	if value != "(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)" {
		t.Errorf("expected pattern %s but received %s", pattern, value)
	}
}

func TestRegnetCompile(t *testing.T) {
	r := regnet.New()
	if r == nil {
		t.Errorf("regnet initialization failed")
	}

	//Compile a pattern without adding it. Throw pattern not found error
	pattern := "%{DAY}"
	_, err := r.Compile(pattern)
	if err == nil {
    	t.Errorf("expected error as the pattern is missing in regnet !")
	}
	

	day := "%{DAY}"
	dayPattern := "(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)"
	r.AddPattern(day, dayPattern)
}
