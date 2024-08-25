package rpslgo

import (
	"strings"
)

type Object struct {
	Attributes []Attribute
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
