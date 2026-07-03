package alias_test

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"testing"

	gloo "github.com/gloo-foo/framework"

	jq "github.com/gloo-foo/cmd-jq/alias"
)

// The alias package re-exports the Jq constructor under an unprefixed name. A
// mis-wired re-export (Jq bound to some other constructor) compiles cleanly, so
// only behavior can prove the wiring. Executing against a real jq install would
// be non-hermetic, so the test empties PATH and asserts the exec failure names
// the jq binary: the re-export provably forks "jq" with the production runner.
func TestAlias_JqForksTheJqBinary(t *testing.T) {
	t.Setenv("PATH", t.TempDir()) // an empty PATH: no jq resolvable
	_, err := jq.Jq(".").Execute(context.Background(), gloo.StreamOf[[]byte]()).Collect()
	if err == nil {
		t.Fatal("expected an exec error with jq absent from PATH, got nil")
	}
	if !errors.Is(err, exec.ErrNotFound) {
		t.Fatalf("expected the not-found exec error to propagate, got %v", err)
	}
	if !strings.Contains(err.Error(), "jq") {
		t.Fatalf("expected the exec error to name the jq binary, got %v", err)
	}
}

// The re-exported constructor must still build a usable Command for any argument
// vector, including the no-argument (bare jq) case.
func TestAlias_JqBuildsCommand(t *testing.T) {
	if jq.Jq(".") == nil {
		t.Fatal("alias.Jq(\".\") returned a nil Command")
	}
	if jq.Jq() == nil {
		t.Fatal("alias.Jq() returned a nil Command")
	}
}
