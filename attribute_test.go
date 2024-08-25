package rpslgo_test

import (
	"testing"

	rpslgo "github.com/frederic-arr/rpsl-go"
)

func TestSingleLineNoValue(t *testing.T) {
	lines := []string{"dry-run:"}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "dry-run" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "dry-run")
	}

	if len(attr.Value) != 0 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 0)
	}
}

func TestSingleLine(t *testing.T) {
	lines := []string{"description:       CERN"}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 1 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 1)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}
}

func TestSingleLineComment(t *testing.T) {
	lines := []string{"description:       CERN # This is a test"}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 1 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 1)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}
}

func TestMultiLines(t *testing.T) {
	lines := []string{
		"description:       CERN",
		"                   European Organization for Nuclear Research",
	}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 2 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 2)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}

	if attr.Value[1] != "European Organization for Nuclear Research" {
		t.Fatalf(`parseAttribute.Value[1] => %v, want %v`, attr.Value[1], "European Organization for Nuclear Research")
	}
}

func TestMultiLinesComment(t *testing.T) {
	lines := []string{
		"description:       CERN",
		"                   European Organization for Nuclear Research # This is a test",
	}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 2 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 2)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}

	if attr.Value[1] != "European Organization for Nuclear Research" {
		t.Fatalf(`parseAttribute.Value[1] => %v, want %v`, attr.Value[1], "European Organization for Nuclear Research")
	}
}

func TestBlankLines(t *testing.T) {
	lines := []string{
		"description:       CERN",
		"+",
		"                   European Organization for Nuclear Research",
	}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 3 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 3)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}

	if attr.Value[1] != "" {
		t.Fatalf(`parseAttribute.Value[1] => %v, want %v`, attr.Value[1], "")
	}

	if attr.Value[2] != "European Organization for Nuclear Research" {
		t.Fatalf(`parseAttribute.Value[2] => %v, want %v`, attr.Value[2], "European Organization for Nuclear Research")
	}
}

func TestCommentLine(t *testing.T) {
	lines := []string{
		"description:       CERN",
		"# This is a test",
		"                   European Organization for Nuclear Research",
	}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 2 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 2)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}

	if attr.Value[1] != "European Organization for Nuclear Research" {
		t.Fatalf(`parseAttribute.Value[1] => %v, want %v`, attr.Value[1], "European Organization for Nuclear Research")
	}
}

func TestBlankLineEnd(t *testing.T) {
	lines := []string{
		"description:       CERN",
		"                   ",
		"                   European Organization for Nuclear Research",
	}

	i := 0
	attr, err := rpslgo.ParseAttribute(&i, lines)
	if err != nil {
		t.Fatalf(`parseAttribute => %v`, err)
	}

	if attr.Name != "description" {
		t.Fatalf(`parseAttribute.Name => %v, want %v`, attr.Name, "description")
	}

	if len(attr.Value) != 1 {
		t.Fatalf(`parseAttribute.Value => length of %v, want %v`, len(attr.Value), 1)
	}

	if attr.Value[0] != "CERN" {
		t.Fatalf(`parseAttribute.Value[0] => %v, want %v`, attr.Value[0], "CERN")
	}
}
