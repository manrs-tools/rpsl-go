package rpsl

import (
	"testing"
)

func TestSingleLineNoValue(t *testing.T) {
	raw := "dry-run:"
	attr, err := parseAttribute(raw)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "dry-run" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "dry-run")
	}

	if attr.Value != "" {
		t.Fatalf(`parseAttribute.Value => %v, want %v`, attr.Value, "")
	}
}

func TestSingleLine(t *testing.T) {
	raw := "description:       CERN"
	attr, err := parseAttribute(raw)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if attr.Value != "CERN" {
		t.Fatalf(`parseAttribute.Value => %v, want %v`, attr.Value, "CERN")
	}
}

func TestSingleLineComment(t *testing.T) {
	raw := "description:       CERN # This is a test"
	attr, err := parseAttribute(raw)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if attr.Value != "CERN" {
		t.Fatalf(`parseAttribute.Value => %v, want %v`, attr.Value, "CERN")
	}
}
