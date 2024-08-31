package rpsl

import (
	"strings"
)

// Parses a string containing a single RPSL object
// and returns a representation of the parsed data.
//
// Parameters:
//
//	raw - A string containing the RPSL object to be parsed.
//
// Returns:
//
//	*Object - A pointer to the Object representation of the parsed RPSL data.
//	*error  - A pointer to an error if the parsing fails, otherwise nil.
//
// Example:
//
//	raw := "person:	John Doe\n"+
//	    "address:	1234 Elm Street\n"+
//	    "phone:		+1 555 123456\n"+
//	    "nic-hdl:	JD1234-RIPE\n"+
//	    "mnt-by:	EXAMPLE-MNT\n"+
//	    "source:	RIPE\n"
//	obj, err := Parse(raw)
//
//	if err != nil {
//	    log.Fatalf("Failed to parse RPSL object: %v", err)
//	}
//
//	fmt.Printf("Parsed Object: %+v\n", obj)
func Parse(raw string) (*Object, *error) {
	lines := strings.Split(raw, "\n")
	i := 0

	obj, err := parseObjectLines(&i, lines)
	return obj, err
}

// Parses a string containing a multiple RPSL object
// and returns a representation of the parsed data.
//
// Parameters:
//
//	raw - A string containing the RPSL objects to be parsed.
//
// Returns:
//
//	*[]Object - A pointer to a list of Object representation of the parsed RPSL data.
//	*error  - A pointer to an error if the parsing fails, otherwise nil.
//
// Example:
//
//	raw := "person:	John Doe\n"+
//		"address:	1234 Elm Street\n"+
//		"phone:		+1 555 123456\n"+
//		"nic-hdl:	JD1234-RIPE\n"+
//		"mnt-by:	EXAMPLE-MNT\n"+
//		"source:	RIPE\n"+
//		"\n"+ // Objects are delimited by one or more empty lines
//		"person:	Jane Smith\n"+
//		"address:	5678 Oak Street\n"+
//		"phone:		+1 555 654321\n"+
//		"nic-hdl:	JS5678-RIPE\n"+
//		"mnt-by:	EXAMPLE-MNT\n"+
//		"source:	RIPE"
//	obj, err := ParseMany(raw)
//
//	if err != nil {
//		log.Fatalf("Failed to parse RPSL object: %v", err)
//	}
//
//	for _, obj := range *objs {
//		fmt.Printf("Parsed Object: %+v\n", obj)
//	}
func ParseMany(raw string) (*[]Object, *error) {
	lines := strings.Split(raw, "\n")
	objects := []Object{}
	i := 0

	for {
		if i >= len(lines) {
			break
		}

		obj, err := parseObjectLines(&i, lines)
		if err != nil {
			return nil, err
		}

		if obj.Len() == 0 {
			i++
			continue
		}

		objects = append(objects, *obj)
	}

	return &objects, nil
}
