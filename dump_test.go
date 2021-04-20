package meroxa

import "testing"

func TestDumpTransport_Obfuscate(t *testing.T) {
	testCases := map[string]string{
		"t":             "*",
		"te":            "**",
		"tes":           "***",
		"test":          "****",
		"testa":         "****a",
		"testau":        "****au",
		"testaut":       "****aut",
		"testauth":      "****auth",
		"testautht":     "*****utht",
		"testauthto":    "******thto",
		"testauthtok":   "*******htok",
		"testauthtoke":  "********toke",
		"testauthtoken": "*********oken",
	}
	dt := &dumpTransport{}
	for have, want := range testCases {
		got := dt.obfuscate(have)
		if got != want {
			t.Fatalf("have %q, expected %q, got %q", have, want, got)
		}
	}
}
