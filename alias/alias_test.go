package alias_test

import (
	"reflect"
	"testing"

	command "github.com/gloo-foo/cmd-jq"
	jq "github.com/gloo-foo/cmd-jq/alias"
)

// The alias package re-exports the Jq constructor under an unprefixed name. A
// mis-wired re-export (Jq bound to some other constructor) compiles cleanly, so
// only behavior can prove the wiring. Executing the returned Command would fork
// real jq — non-hermetic and dependent on a working install — so instead the
// test proves the re-export points at the exact same constructor: same function
// identity means identical forking behavior.
func TestAlias_JqReExportsConstructor(t *testing.T) {
	got := reflect.ValueOf(jq.Jq).Pointer()
	want := reflect.ValueOf(command.Jq).Pointer()
	if got != want {
		t.Fatalf("alias.Jq is not wired to command.Jq (%v != %v)", got, want)
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
