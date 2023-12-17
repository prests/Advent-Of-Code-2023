package testingutils

import (
	"reflect"
	"testing"
)

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertDeepEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error: %v", err)
	}
}
