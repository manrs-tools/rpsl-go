// Copyright (c) The RPSL Go Authors.
// SPDX-License-Identifier: Apache-2.0

package rpsl

import (
	"testing"
)

func TestObject(t *testing.T) {
	raw := "organisation:      ORG-CEOf1-RIPE\n" +
		"description:       CERN"

	objects, err := parseObjects(raw)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	if len(objects) != 1 {
		t.Fatalf(`parseObject => length of %v, want %v`, len(objects), 1)
	}

	obj := objects[0]
	if len(obj.Attributes) != 2 {
		t.Fatalf(`object.Attributes => length of %v, want %v`, len(obj.Attributes), 2)
	}
}

func TestObjectLen(t *testing.T) {
	raw := "organisation:      ORG-CEOf1-RIPE\n" +
		"remarks:           This is a comment\n" +
		"description:       CERN\n" +
		"remarks:           This is another comment"

	objects, err := parseObjects(raw)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	if len(objects) != 1 {
		t.Fatalf(`parseObject => length of %v, want %v`, len(objects), 1)
	}

	obj := objects[0]
	if obj.Len() != 4 {
		t.Fatalf(`object.Len() => %v, want %v`, obj.Len(), 4)
	}
}

func TestObjectKeys(t *testing.T) {
	raw := "organisation:      ORG-CEOf1-RIPE\n" +
		"remarks:           This is a comment\n" +
		"description:       CERN\n" +
		"remarks:           This is another comment"

	objects, err := parseObjects(raw)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	if len(objects) != 1 {
		t.Fatalf(`parseObject => length of %v, want %v`, len(objects), 1)
	}

	obj := objects[0]
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

func TestObjectGetAll(t *testing.T) {
	raw := "organisation:      ORG-CEOf1-RIPE\n" +
		"remarks:           This is a comment\n" +
		"description:       CERN\n" +
		"remarks:           This is another comment"

	objects, err := parseObjects(raw)
	if err != nil {
		t.Fatalf(`parseObject => %v`, err)
	}

	if len(objects) != 1 {
		t.Fatalf(`parseObject => length of %v, want %v`, len(objects), 1)
	}

	obj := objects[0]
	attrs := obj.GetAll("organisation")
	if len(attrs) != 1 {
		t.Fatalf(`object.GetAll("organisation") => length of %v, want %v`, len(attrs), 1)
	}

	if attrs[0].Value != "ORG-CEOf1-RIPE" {
		t.Fatalf(`object.GetAll("organisation").Value => %v, want %v`, attrs[0].Value, "ORG-CEOf1-RIPE")
	}

	attrs = obj.GetAll("description")
	if len(attrs) != 1 {
		t.Fatalf(`object.GetAll("description") => length of %v, want %v`, len(attrs), 1)
	}

	if attrs[0].Value != "CERN" {
		t.Fatalf(`object.GetAll("description").Value => %v, want %v`, attrs[0].Value, "CERN")
	}

	attrs = obj.GetAll("remarks")
	if len(attrs) != 2 {
		t.Fatalf(`object.GetAll("remarks") => length of %v, want %v`, len(attrs), 2)
	}

	if attrs[0].Value != "This is a comment" {
		t.Fatalf(`object.GetAll("remarks")[0].Value => %v, want %v`, attrs[0].Value, "This is a comment")
	}

	if attrs[1].Value != "This is another comment" {
		t.Fatalf(`object.GetAll("remarks")[1].Value => %v, want %v`, attrs[1].Value, "This is another comment")
	}
}
