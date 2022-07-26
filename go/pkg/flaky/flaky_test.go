package flaky

import "testing"

func TestFlake(t *testing.T) {
	if Flake() {
		t.Fatal("flaky")
	}
}
