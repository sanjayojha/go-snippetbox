package assert

import "testing"

func Equal[T comparable](t *testing.T, actaul, expected T) {
	t.Helper()
	if actaul != expected {
		t.Errorf("expected: %v; but got %v", expected, actaul)
	}
}
