package meroxa

import "testing"

func TestDumpTransport_Obfuscate(t *testing.T) {
	testCases := map[string]string{
		"t":                      "*",
		"te":                     "**",
		"tes":                    "***",
		"test":                   "****",
		"testa":                  "t...a",
		"testau":                 "t...u",
		"testaut":                "te...ut",
		"testauth":               "te...th",
		"testautho":              "tes...tho",
		"testauthor":             "tes...hor",
		"testauthori":            "test...hori",
		"testauthoriz":           "test...oriz",
		"testauthoriza":          "testa...oriza",
		"testauthorizat":         "testa...rizat",
		"testauthorizati":        "testau...rizati",
		"testauthorizatio":       "testau...izatio",
		"testauthorization":      "testaut...ization",
		"testauthorizationt":     "testaut...zationt",
		"testauthorizationto":    "testaut...ationto",
		"testauthorizationtok":   "testaut...tiontok",
		"testauthorizationtoke":  "testaut...iontoke",
		"testauthorizationtoken": "testaut...ontoken",
	}
	dt := &dumpTransport{}
	for have, want := range testCases {
		got := dt.obfuscate(have)
		if got != want {
			t.Fatalf("have %q, expected %q, got %q", have, want, got)
		}
	}
}
