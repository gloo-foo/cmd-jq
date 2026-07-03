package command

import (
	"context"
	"slices"
	"testing"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// captured records the argument vector a fake runner was handed, so a test can
// assert exactly what Jq would have forked — proving the contract without
// executing jq (which would be non-hermetic and depend on a working install).
type captured struct {
	name patterns.ProcessName
	args []patterns.ProcessArg
}

// fakeRunner returns a runner that records its inputs into got and yields a
// trivial pass-through Command, so no real process is ever spawned.
func fakeRunner(got *captured) runner {
	return func(name patterns.ProcessName, args ...patterns.ProcessArg) gloo.Command[[]byte, []byte] {
		got.name = name
		got.args = args
		return gloo.FuncCommand[[]byte, []byte](
			func(_ context.Context, in gloo.Stream[[]byte]) gloo.Stream[[]byte] { return in },
		)
	}
}

func TestJqWith_PrependsJqToArgs(t *testing.T) {
	var got captured
	jqWith(fakeRunner(&got), "-c", ".items[]")
	if got.name != "jq" {
		t.Errorf("forked %q, want \"jq\"", got.name)
	}
	if !slices.Equal(got.args, []patterns.ProcessArg{"-c", ".items[]"}) {
		t.Errorf("got args %q, want [-c .items[]]", got.args)
	}
}

func TestJqWith_NoArgsForksBareJq(t *testing.T) {
	var got captured
	jqWith(fakeRunner(&got))
	if got.name != "jq" {
		t.Errorf("forked %q, want \"jq\"", got.name)
	}
	if len(got.args) != 0 {
		t.Errorf("got args %q, want none", got.args)
	}
}

// Jq wires the production runner (patterns.Subprocess). Constructing the Command
// must not fork jq, so this asserts only that a usable Command is returned; the
// argument-vector contract is proven against the fake runner.
func TestJq_ReturnsCommand(t *testing.T) {
	if Jq(".") == nil {
		t.Fatal("Jq returned a nil Command")
	}
}
