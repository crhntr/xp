package xp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func MatchingStrings(t *testing.T, actual, expect string) {
	t.Helper()
	if expect != actual {
		t.Errorf("actual does not match expected")
		t.Logf("\texpected: %s", expect)
		t.Logf("\tactual: %s", actual)
	}
}

func MatchingInt(t *testing.T, actual, expect int) {
	t.Helper()
	if expect != actual {
		t.Errorf("actual does not match expected")
		t.Logf("\texpected: %d", expect)
		t.Logf("\tactual: %d", actual)
	}
}

func MatchingFloat64(t *testing.T, actual, expect, tollerance float64) {
	t.Helper()
	if (expect - actual) >= tollerance {
		t.Errorf("actual does not match expected")
		t.Logf("\texpected: %f", expect)
		t.Logf("\tactual: %f", actual)
		t.Logf("\ttollerance: %f", tollerance)
	}
}

func NoErrorOccured(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("an error occured when none was expected")
		t.Log(err)
	}
}

func MatchingJSON(t *testing.T, actual interface{}, expect string) {
	t.Helper()
	var (
		actbuf   []byte
		exp, act map[string]interface{}
	)
	switch exi := actual.(type) {
	case io.Reader:
		actbuf, _ = ioutil.ReadAll(exi)
	case []byte:
		actbuf = exi
	case string:
		actbuf = []byte(exi)
	}
	if err := json.Unmarshal(actbuf, &act); err != nil {
		t.Errorf("could not unmarshal actual json: %s", err)
	}
	if err := json.Unmarshal([]byte(expect), &exp); err != nil {
		t.Errorf("could not unmarshal expected json: %s", err)
	}
	if !reflect.DeepEqual(exp, act) {
		t.Errorf("actual json does not match expected")
		t.Logf("\texpected: %s", string(expect))
		t.Logf("\tactual: %s", string(actbuf))
	}
}
