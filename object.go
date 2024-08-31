package rpsl

import (
	"strings"
)

type Object struct {
	Attributes []Attribute
}

// Returns a slice of unique keys present in the Object.
// If a key appears multiple times in the Object, it will only be included once in the returned slice.
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

// Returns the number of attributes in the Object.
func (o *Object) Len() int {
	return len(o.Attributes)
}

// Returns a slice of values for a given key in the Object.
// If the key is not present in the Object, an empty slice will be returned.
// If a key appears multiple times in the Object, all values will be included in the returned slice.
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

func (o *Object) Exists(key string) bool {
	key = strings.ToLower(key)
	for _, attr := range o.Attributes {
		if attr.Name == key {
			return true
		}
	}

	return false
}

func (o *Object) String() string {
	var str strings.Builder
	for _, attr := range o.Attributes {
		str.WriteString(attr.String())
		str.WriteString("\n")
	}

	return str.String()
}

func parseObjectLines(i *int, lines []string) (*Object, *error) {
	object := Object{}

	for {
		if *i >= len(lines) {
			break
		}

		line := lines[*i]
		if len(line) == 0 || strings.TrimSpace(line) == "" {
			break
		}

		attr, err := parseAttributeLines(i, lines)
		if err != nil {
			return nil, err
		}

		object.Attributes = append(object.Attributes, *attr)
	}

	return &object, nil
}
