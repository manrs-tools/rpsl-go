package rpsl

import (
	"errors"
	"regexp"
	"strings"
)

type Attribute struct {
	Name    string
	Value   string
	Comment *string
}

var attr_re = regexp.MustCompile(`^(?P<Name>[a-z0-9-]+): *(?P<Value>[^#]*?) *(?:# *(?P<Comment>.*?) *)?$`)

func parseAttribute(line string) (*Attribute, error) {
	matches := attr_re.FindStringSubmatch(line)
	if matches == nil {
		return nil, errors.New("invalid attribute: " + line)
	}

	var name string
	var value string
	var comment *string

	name = strings.TrimSpace(matches[attr_re.SubexpIndex("Name")])
	value = strings.TrimSpace(matches[attr_re.SubexpIndex("Value")])
	if strings.Contains(line, "#") {
		c := strings.TrimSpace(matches[attr_re.SubexpIndex("Comment")])
		comment = &c
	}

	attr := &Attribute{
		Name:    name,
		Value:   value,
		Comment: comment,
	}

	return attr, nil
}

func (a *Attribute) String() string {
	var str strings.Builder
	str.WriteString(a.Name)
	str.WriteString(":")
	str.WriteString(a.Value)
	if a.Comment != nil {
		str.WriteString(" # ")
		str.WriteString(*a.Comment)
	}

	return str.String()
}
