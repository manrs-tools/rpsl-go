package rpsl_test

import (
	"testing"

	"github.com/frederic-arr/rpsl-go"
)

func TestSingleObject(t *testing.T) {
	raw := "person:        John Doe\n" +
		"address:       1234 Elm Street\n" +
		"               Iceland\n" +
		"phone:         +1 555 123456\n" +
		"nic-hdl:       JD1234-RIPE\n" +
		"mnt-by:        FOO-MNT\n" +
		"mnt-by:        BAR-MNT\n" +
		"source:        RIPE"

	obj, err := rpsl.Parse(raw)
	if err != nil {
		t.Fatalf("parseObject => %v", err)
	}

	if obj.Len() != 7 {
		t.Fatalf(`parseObject.Len => %v, want %v`, obj.Len(), 7)
	}

	keys := obj.Keys()
	if len(keys) != 6 {
		t.Fatalf(`parseObject.Keys => length of %v, want %v`, len(keys), 6)
	}
}

func TestMultipleObjectsSingleParse(t *testing.T) {
	raw := "person:        John Doe\n" +
		"address:       1234 Elm Street\n" +
		"               Iceland\n" +
		"phone:         +1 555 123456\n" +
		"nic-hdl:       JD1234-RIPE\n" +
		"mnt-by:        FOO-MNT\n" +
		"mnt-by:        BAR-MNT\n" +
		"source:        RIPE\n" +
		"\n" +
		"person:        Jane Doe\n" +
		"address:       5678 Oak Street\n" +
		"               Greenland\n" +
		"phone:         +1 555 654321\n" +
		"nic-hdl:       JD5678-RIPE\n" +
		"mnt-by:        FOO-MNT\n" +
		"mnt-by:        BAR-MNT\n" +
		"source:        RIPE"

	obj, err := rpsl.Parse(raw)
	if err != nil {
		t.Fatalf("parseObject => %v", err)
	}

	if obj.Len() != 7 {
		t.Fatalf(`parseObject.Len => %v, want %v`, obj.Len(), 7)
	}

	keys := obj.Keys()
	if len(keys) != 6 {
		t.Fatalf(`parseObject.Keys => length of %v, want %v`, len(keys), 6)
	}
}

func TestMultipleObjectsMultipleParse(t *testing.T) {
	raw := "person:        John Doe\n" +
		"address:       1234 Elm Street\n" +
		"               Iceland\n" +
		"phone:         +1 555 123456\n" +
		"nic-hdl:       JD1234-RIPE\n" +
		"mnt-by:        FOO-MNT\n" +
		"mnt-by:        BAR-MNT\n" +
		"source:        RIPE\n" +
		"\n" +
		"person:        Jane Doe\n" +
		"address:       5678 Oak Street\n" +
		"               Greenland\n" +
		"phone:         +1 555 654321\n" +
		"nic-hdl:       JD5678-RIPE\n" +
		"mnt-by:        FOO-MNT\n" +
		"mnt-by:        BAR-MNT\n" +
		"source:        RIPE"

	objs, err := rpsl.ParseMany(raw)
	if err != nil {
		t.Fatalf("parseObject => %v", err)
	}

	if len(*objs) != 2 {
		t.Fatalf(`parseObject.Len => %v, want %v`, len(*objs), 2)
	}

	for _, obj := range *objs {
		if obj.Len() != 7 {
			t.Fatalf(`parseObject.Len => %v, want %v`, obj.Len(), 7)
		}

		keys := obj.Keys()
		if len(keys) != 6 {
			t.Fatalf(`parseObject.Keys => length of %v, want %v`, len(keys), 6)
		}
	}
}
