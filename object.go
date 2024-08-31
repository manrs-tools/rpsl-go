package rpsl

import (
	"strings"
)

type Object struct {
	Attributes []Attribute
}

func (o *Object) Keys() []string {
	var keyList []string
	keyPresent := make(map[string]bool)
	for _, attr := range o.Attributes {
		if _, ok := keyPresent[attr.Name]; !ok {
			keyPresent[attr.Name] = true
			keyList = append(keyList, attr.Name)
		}
	}

	return keyList
}

func (o *Object) Len() int {
	return len(o.Attributes)
}

func (o *Object) Get(key string) []string {
	key = strings.ToLower(key)
	var values []string
	for _, attr := range o.Attributes {
		if attr.Name == key {
			values = append(values, attr.Value...)
		}
	}

	return values
}

func ParseObject(i *int, lines []string) (*Object, *error) {
	object := Object{}

	for {
		if *i >= len(lines) {
			break
		}

		line := lines[*i]
		if len(line) == 0 || strings.TrimSpace(line) == "" {
			break
		}

		attr, err := ParseAttribute(i, lines)
		if err != nil {
			return nil, err
		}

		object.Attributes = append(object.Attributes, *attr)
	}

	return &object, nil
}
