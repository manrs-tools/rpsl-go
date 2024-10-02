package rpsl

import (
	"fmt"
	"strings"
)

type Attribute struct {
	Name  string
	Value string
}

func parseAttribute(line string) (*Attribute, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid attribute: %q", line)
	}

	name := strings.TrimSpace(parts[0])
	name = strings.ToLower(name)

	value := strings.SplitN(parts[1], "#", 2)[0]
	value = strings.TrimSpace(value)

	attr := &Attribute{
		Name:  name,
		Value: value,
	}

	return attr, nil
}

func (a *Attribute) String() string {
	var str strings.Builder
	str.WriteString(a.Name)
	str.WriteString(":")
	str.WriteString(a.Value)
	return str.String()
}
