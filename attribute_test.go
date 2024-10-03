// Copyright (c) The RPSL Go Authors.
// SPDX-License-Identifier: Apache-2.0

package rpsl

import "testing"

func TestSingleObject(t *testing.T) {
	data := "" +
		"mntner:          DEV-MNT  # Comment \n" +
		"descr:           DEV maintainer\n" +
		"admin-c:         VM1-DEV\n" +
		"tech-c:          VM1-DEV\n" +
		"upd-to:          v.m@example.net\n" +
		"mnt-nfy:         auto@example.net\n" +
		"auth:            MD5-PW $1$q8Su3Hq/$rJt5M3TNLeRE4UoCh5bSH/\n" +
		"remarks:         password: secret\n" +
		"mnt-by:          DEV-MNT\n" +
		"source:          DEV\n"

	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 10 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 10)
	}

	if attr[0].Name != "mntner" {
		t.Fatalf(`mntner: got %v, want %v`, attr[0].Name, "mntner")
	}

	if attr[0].Value != "DEV-MNT" {
		t.Fatalf(`mntner: got %v, want %v`, attr[0].Value, "DEV-MNT")
	}

	if attr[9].Name != "source" {
		t.Fatalf(`source: got %v, want %v`, attr[9].Name, "source")
	}

	if attr[9].Value != "DEV" {
		t.Fatalf(`source: got %v, want %v`, attr[9].Value, "DEV")
	}
}

func TestAllowedCharactersInKey(t *testing.T) {
	data := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-:"
	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 1 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 2)
	}
}

func TestSpecialCharacters(t *testing.T) {
	data := "" +
		"person:  New Test Person\n" +
		"address: Flughafenstraße 120\n" +
		"address: D - 40474 Düsseldorf\n" +
		"nic-hdl: ABC-RIPE\n"

	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 4 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 4)
	}

	if attr[1].Value != "Flughafenstraße 120" {
		t.Fatalf(`address: got %v, want %v`, attr[1].Value, "Flughafenstraße 120")
	}

	if attr[2].Value != "D - 40474 Düsseldorf" {
		t.Fatalf(`address: got %v, want %v`, attr[2].Value, "D - 40474 Düsseldorf")
	}
}

func TestGarbageValue(t *testing.T) {
	data := "mntner::!@$%^&*()_+~![]{};':<>,./?\\"
	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 1 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 2)
	}

	if attr[0].Name != "mntner" {
		t.Fatalf(`(0.name): got %v, want %v`, attr[0].Name, "mntner")
	}

	if attr[0].Value != ":!@$%^&*()_+~![]{};':<>,./?\\" {
		t.Fatalf(`(0.value): got %v, want %v`, attr[0].Value, ":!@$%^&*()_+~![]{};':<>,./?\\")
	}
}

func TestContinuationLines(t *testing.T) {
	data := "" +
		"mntner:   DEV-MNT\n" +
		"descr:    \n" +
		"+1\n" +
		" 2\n" +
		"\t3"

	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 2 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 2)
	}

	if attr[1].Value != "1 2 3" {
		t.Fatalf(`descr: got %v, want %v`, attr[1].Value, "1 2 3")
	}
}

func TestNumberInKey(t *testing.T) {
	data := "" +
		"route6:         2001:0000::/32\n" +
		"origin:AS10"

	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 2 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 2)
	}

	if attr[0].Name != "route6" {
		t.Fatalf(`route6: got %v, want %v`, attr[0].Name, "route6")
	}
}

func TestCasingAndNumberInKey(t *testing.T) {
	data := "" +
		"roUte6:         2001:0000::/32\n" +
		"origin:AS10"

	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 2 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 2)
	}

	if attr[0].Name != "route6" {
		t.Fatalf(`route6: got %v, want %v`, attr[0].Name, "route6")
	}
}

func TestComment(t *testing.T) {
	data := "" +
		"person: foo\n" +
		"nic-hdl: VM1-DEV  # SOME COMMENT \n"

	attr, err := parseAttributes(data)
	if err != nil {
		t.Fatalf(`(error): %v`, err)
	}

	if len(attr) != 2 {
		t.Fatalf(`(length): got %v, want %v`, len(attr), 2)
	}

	if attr[1].Name != "nic-hdl" {
		t.Fatalf(`nic-hdl: got %v, want %v`, attr[1].Name, "nic-hdl")
	}

	if attr[1].Value != "VM1-DEV" {
		t.Fatalf(`nic-hdl: got %v, want %v`, attr[1].Value, "VM1-DEV")
	}
}

// func TestSingleLineNoValue(t *testing.T) {
// 	raw := "dry-run:"
// 	attr, err := parseAttribute(raw)
// 	if err != nil {
// 		t.Fatalf(`parseAttribute => %v`, err)
// 	}

// 	if attr.Name != "dry-run" {
// 		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "dry-run")
// 	}

// 	if attr.Value != "" {
// 		t.Fatalf(`parseAttribute.Value => %v, want %v`, attr.Value, "")
// 	}
// }

// func TestSingleLine(t *testing.T) {
// 	raw := "description:       CERN"
// 	attr, err := parseAttribute(raw)
// 	if err != nil {
// 		t.Fatalf(`parseAttribute => %v`, err)
// 	}

// 	if attr.Name != "description" {
// 		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
// 	}

// 	if attr.Value != "CERN" {
// 		t.Fatalf(`parseAttribute.Value => %v, want %v`, attr.Value, "CERN")
// 	}
// }

// func TestSingleLineComment(t *testing.T) {
// 	raw := "description:       CERN # This is a test"
// 	attr, err := parseAttribute(raw)
// 	if err != nil {
// 		t.Fatalf(`parseAttribute => %v`, err)
// 	}

// 	if attr.Name != "description" {
// 		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
// 	}

// 	if attr.Value != "CERN" {
// 		t.Fatalf(`parseAttribute.Value => %v, want %v`, attr.Value, "CERN")
// 	}
// }
