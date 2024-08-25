package rpslgo

import "strings"

type Attribute struct {
	Name  string
	Value []string
}

func ParseAttribute(i *int, lines []string) (*Attribute, *error) {
	attribute := Attribute{}
	line := lines[*i]
	data := strings.SplitN(line, ":", 2)
	attribute.Name = data[0]

	if len(data) < 2 {
		return &attribute, nil
	}

	value := strings.SplitN(data[1], "#", 2)[0]
	value = strings.TrimSpace(value)
	if value == "" {
		return &attribute, nil
	}

	attribute.Value = append(attribute.Value, value)
	for *i = *i + 1; *i < len(lines); *i++ {
		line := lines[*i]
		if len(line) == 0 {
			break
		}

		c := line[0]
		if c != ' ' && c != '\t' && c != '+' && c != '#' && c != '%' {
			break
		}

		if c == '#' || c == '%' {
			continue
		}

		value := strings.SplitN(line[1:], "#", 2)[0]
		value = strings.TrimSpace(value)
		if value == "" && c != '+' {
			break
		}

		if value != "" && (value[0] == '#' || value[0] == '%') {
			continue
		}

		attribute.Value = append(attribute.Value, value)
	}

	return &attribute, nil
}
