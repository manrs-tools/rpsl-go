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

// Returns the first value for a given key in the Object.
// If the key is not present in the Object, an empty string will be returned.
// If a key appears multiple times in the Object, only the first value will be returned.
func (o *Object) GetFirst(key string) *Attribute {
	key = strings.ToLower(key)
	for _, attr := range o.Attributes {
		if attr.Name == key {
			return &attr
		}
	}

	return nil
}

// Returns a slice of values for a given key in the Object.
// If the key is not present in the Object, an empty slice will be returned.
// If a key appears multiple times in the Object, all values will be included in the returned slice.
func (o *Object) GetAll(key string) []Attribute {
	var attributes []Attribute
	for _, attr := range o.Attributes {
		if attr.Name == key {
			attributes = append(attributes, attr)
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
	for _, attr := range o.Attributes {
		str.WriteString(attr.String())
		str.WriteString("\n")
	}

	return str.String()
}

func parseObjects(raw string) ([]Object, error) {
	objects := []Object{}
	for _, part := range strings.Split(raw, "\n\n") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		lines := strings.Split(part, "\n")
		object := Object{}
		hasAttributes := false
		for _, line := range lines {
			if strings.HasPrefix(lines[0], "%") {
				continue
			}

			attr, err := parseAttribute(line)
			if err != nil {
				return nil, err
			}

			hasAttributes = true
			object.Attributes = append(object.Attributes, *attr)
		}

		if !hasAttributes {
			continue
		}

		objects = append(objects, object)
	}

	return objects, nil
}
