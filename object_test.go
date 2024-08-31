package rpsl_test

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestObject(t *testing.T) {
	lines := []string{
		"organisation:      ORG-CEOf1-RIPE",
		"description:       CERN",
	}

	i := 0
	obj, err := rpsl.ParseObject(&i, lines)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	if len(obj.Attributes) != 2 {
		t.Fatalf(`parseObject.Attributes => length of %v, want %v`, len(obj.Attributes), 2)
	}

	if obj.Attributes[0].Name != "organisation" {
		t.Fatalf(`parseObject.Attributes[0].Name => %v, want %v`, obj.Attributes[0].Name, "organisation")
	}

	if len(obj.Attributes[0].Value) != 1 {
		t.Fatalf(`parseObject.Attributes[0].Attributes => length of %v, want %v`, len(obj.Attributes[0].Value), 1)
	}

	if obj.Attributes[0].Value[0] != "ORG-CEOf1-RIPE" {
		t.Fatalf(`parseObject.Attributes[0].Value[0] => %v, want %v`, obj.Attributes[0].Value[0], " ORG-CEOf1-RIPE")
	}

	if obj.Attributes[1].Name != "description" {
		t.Fatalf(`parseObject.Attributes[1].Name => %v, want %v`, obj.Attributes[1].Name, "description")
	}

	if len(obj.Attributes[1].Value) != 1 {
		t.Fatalf(`parseObject.Attributes[1].Attributes => length of %v, want %v`, len(obj.Attributes[1].Value), 1)
	}

	if obj.Attributes[1].Value[0] != "CERN" {
		t.Fatalf(`parseObject.Attributes[1].Value[0] => %v, want %v`, obj.Attributes[1].Value[0], " CERN")
	}
}

func TestObjectComments(t *testing.T) {
	lines := []string{
		"organisation:      ORG-CEOf1-RIPE",
		"# This is a comment",
		"description:       CERN",
	}

	i := 0
	obj, err := rpsl.ParseObject(&i, lines)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	if len(obj.Attributes) != 2 {
		t.Fatalf(`parseObject.Attributes => length of %v, want %v`, len(obj.Attributes), 2)
	}

	if obj.Attributes[0].Name != "organisation" {
		t.Fatalf(`parseObject.Attributes[0].Name => %v, want %v`, obj.Attributes[0].Name, "organisation")
	}

	if len(obj.Attributes[0].Value) != 1 {
		t.Fatalf(`parseObject.Attributes[0].Attributes => length of %v, want %v`, len(obj.Attributes[0].Value), 1)
	}

	if obj.Attributes[0].Value[0] != "ORG-CEOf1-RIPE" {
		t.Fatalf(`parseObject.Attributes[0].Value[0] => %v, want %v`, obj.Attributes[0].Value[0], " ORG-CEOf1-RIPE")
	}

	if obj.Attributes[1].Name != "description" {
		t.Fatalf(`parseObject.Attributes[1].Name => %v, want %v`, obj.Attributes[1].Name, "description")
	}

	if len(obj.Attributes[1].Value) != 1 {
		t.Fatalf(`parseObject.Attributes[1].Attributes => length of %v, want %v`, len(obj.Attributes[1].Value), 1)
	}

	if obj.Attributes[1].Value[0] != "CERN" {
		t.Fatalf(`parseObject.Attributes[1].Value[0] => %v, want %v`, obj.Attributes[1].Value[0], " CERN")
	}
}
