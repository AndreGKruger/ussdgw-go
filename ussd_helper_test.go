package ussdgwapp

import (
	"testing"
)

func TestTransformToSnakeCase(t *testing.T) {
	want := "this_is_a_test"
	got := transformToSnakeCase("this-is-a-test")
	if got != want {
		t.Errorf("transformToSnakeCase() = %v, want %v", got, want)
	}

	want = "this_is_a_test"
	got = transformToSnakeCase("this_is_a_test")
	if got != want {
		t.Errorf("transformToSnakeCase() = %v, want %v", got, want)
	}

	want = "this_is_a_test"
	got = transformToSnakeCase("thisIsATest")
	if got != want {
		t.Errorf("transformToSnakeCase() = %v, want %v", got, want)
	}
}
