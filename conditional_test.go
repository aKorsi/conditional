package conditional_test

import (
	"github.com/aKorsi/conditional"
	"testing"
)

func TestConditional(t *testing.T) {

	if conditional.Ternary(true, 5, 4) != 5 {
		t.Fatal("Error Int")
	}

	if conditional.Ternary(false, 5, 4) != 4 {
		t.Fatal("Error Int")
	}

	if conditional.Ternary(true, "ok", "not ok") != "ok" {
		t.Fatal("Error string")
	}

	if conditional.Ternary(false, "ok", "not ok") != "not ok" {
		t.Fatal("Error string")
	}

	var val1 *string
	val2 := "val2"
	if conditional.NilCoalescing(val1, &val2) != &val2 {
		t.Fatal("Error NilCoalescing")
	}

	val3 := "val3"
	if conditional.NilCoalescing(&val3, &val2) != &val3 {
		t.Fatal("Error NilCoalescing")
	}

	var st *struct {
		val2 string
	}
	if conditional.OptionalChaining(st, func() *string {
		return &st.val2
	}) != nil {
		t.Fatal("Error OptionalChaining")
	}

	st = &struct {
		val2 string
	}{val2: "val2"}
	if *conditional.OptionalChaining(st, func() *string {
		return &st.val2
	}) != "val2" {
		t.Fatal("Error OptionalChaining")
	}
}
