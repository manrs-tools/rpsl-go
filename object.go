// Copyright (c) The RPSL Go Authors.
// SPDX-License-Identifier: Apache-2.0

package rpsl

import (
	"fmt"
	"strings"
)

type Object struct {
	Attributes []Attribute
}

// Returns a slice of unique keys present in the Object.
// If a key appears multiple times in the Object, it will only be included once in the returned slice.
func (o *Object) Keys() []string {
	keyPresent := make(map[string]struct{})
	keyList := make([]string, 0)
	for _, attr := range o.Attributes {
		if _, exists := keyPresent[attr.Name]; !exists {
			keyPresent[attr.Name] = struct{}{}
			keyList = append(keyList, attr.Name)
		}
	}

	return keyList
}

// Returns the number of attributes in the Object.
func (o *Object) Len() int {
	return len(o.Attributes)
}

// Returns the first value for a given key in the Object.
// If the key is not present in the Object, an empty string will be returned.
// If a key appears multiple times in the Object, only the first value will be returned.
func (o *Object) GetFirst(key string) *string {
	key = strings.ToLower(key)
	for _, attr := range o.Attributes {
		if attr.Name == key {
			return &attr.Value
		}
	}

	return nil
}

// Returns a slice of values for a given key in the Object.
// If the key is not present in the Object, an empty slice will be returned.
// If a key appears multiple times in the Object, all values will be included in the returned slice.
func (o *Object) GetAll(key string) []string {
	attributes := make([]string, 0)
	for _, attr := range o.Attributes {
		if attr.Name == key {
			attributes = append(attributes, attr.Value)
		}
	}

	return attributes
}

// Returns true if the Object contains a given key.
func (o *Object) Exists(key string) bool {
	key = strings.ToLower(key)
	for _, attr := range o.Attributes {
		if attr.Name == key {
			return true
		}
	}

	return false
}

// Returns a string representation of the Object.
func (o *Object) String() string {
	var str strings.Builder

	var attributes []string
	for _, attr := range o.Attributes {
		attributes = append(attributes, attr.String())
	}

	str.WriteString(strings.Join(attributes, "\n"))
	return str.String()
}

// Ensures that the first attribute in the Object is of a given class.
func (o *Object) EnsureClass(class string) error {
	if len(o.Attributes) == 0 {
		return fmt.Errorf("object has no attributes")
	}

	if o.Attributes[0].Name != class {
		return fmt.Errorf("object class is not %s", class)
	}

	return nil
}

// Ensures that the Object has at least one attribute with a given key.
func (o *Object) EnsureAtLeastOne(key string) error {
	if !o.Exists(key) {
		return fmt.Errorf("object has no %s attribute", key)
	}

	return nil
}

// Ensures that the Object has at most one attribute with a given key.
func (o *Object) EnsureAtMostOne(key string) error {
	if len(o.GetAll(key)) > 1 {
		return fmt.Errorf("object has more than one %s attribute", key)
	}

	return nil
}

// Ensures that the Object has exactly one attribute with a given key.
func (o *Object) EnsureOne(key string) error {
	if err := o.EnsureAtLeastOne(key); err != nil {
		return err
	}

	if err := o.EnsureAtMostOne(key); err != nil {
		return err
	}

	return nil
}

func parseObjects(buf string) ([]Object, error) {
	objects := []Object{}

	if buf == "" {
		return objects, nil
	}

	lines := strings.Split(buf, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "%") || strings.HasPrefix(line, "#") {
			lines[i] = ""
		}
	}

	buf = strings.Join(lines, "\n")
	for _, part := range strings.Split(buf, "\n\n") {
		part = strings.TrimPrefix(part, "\n")
		part = strings.TrimSuffix(part, "\n")
		if part == "" {
			continue
		}

		attributes, err := parseAttributes(part)
		if err != nil {
			return nil, err
		}

		object := Object{Attributes: attributes}
		objects = append(objects, object)
	}

	return objects, nil
}
