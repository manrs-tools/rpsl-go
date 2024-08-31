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

func TestObjectLen(t *testing.T) {
	lines := []string{
		"organisation:      ORG-CEOf1-RIPE",
		"remarks:           This is a comment",
		"description:       CERN",
		"remarks:           This is another comment",
	}

	i := 0
	obj, err := rpsl.ParseObject(&i, lines)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	l := obj.Len()
	if l != 4 {
		t.Fatalf(`object.Len() => %v, want %v`, l, 4)
	}
}

func TestObjectKeys(t *testing.T) {
	lines := []string{
		"organisation:      ORG-CEOf1-RIPE",
		"remarks:           This is a comment",
		"description:       CERN",
		"remarks:           This is another comment",
	}

	i := 0
	obj, err := rpsl.ParseObject(&i, lines)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	keys := obj.Keys()
	if len(keys) != 3 {
		t.Fatalf(`object.Keys() => length of %v, want %v`, len(keys), 3)
	}

	if keys[0] != "organisation" {
		t.Fatalf(`object.Keys()[0] => %v, want %v`, keys[0], "organisation")
	}

	if keys[1] != "remarks" {
		t.Fatalf(`object.Keys()[1] => %v, want %v`, keys[1], "remarks")
	}

	if keys[2] != "description" {
		t.Fatalf(`object.Keys()[2] => %v, want %v`, keys[2], "description")
	}
}

func TestObjectGet(t *testing.T) {
	lines := []string{
		"organisation:      ORG-CEOf1-RIPE",
		"remarks:           This is a comment",
		"description:       CERN",
		"remarks:           This is another comment",
	}

	i := 0
	obj, err := rpsl.ParseObject(&i, lines)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	keys := obj.Get("remarks")
	if len(keys) != 2 {
		t.Fatalf(`object.Get("remarks") => length of %v, want %v`, len(keys), 2)
	}

	if keys[0] != "This is a comment" {
		t.Fatalf(`object.Get("remarks")[0] => %v, want %v`, keys[0], "This is a comment")
	}

	if keys[1] != "This is another comment" {
		t.Fatalf(`object.Get("remarks")[1] => %v, want %v`, keys[1], "This is another comment")
	}
}
