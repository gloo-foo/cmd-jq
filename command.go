package command

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// runner forks an external process named name with args, returning the streaming
// Command that drives it. patterns.Subprocess is the production runner; tests
// inject a fake to assert the exact argument vector without forking jq.
//
// args is []string because it flows verbatim into patterns.Subprocess (and
// ultimately os/exec), whose signature this wrapper does not control.
type runner func(name string, args ...string) gloo.Command[[]byte, []byte]

// Jq returns a Command that forks jq with the given filter and arguments,
// streaming pipeline input to jq's stdin and jq's stdout back into the pipeline.
// Every argument is passed through verbatim — jq, not this wrapper, interprets
// them.
func Jq(args ...string) gloo.Command[[]byte, []byte] {
	return jqWith(patterns.Subprocess, args...)
}

// jqWith is Jq with an injectable runner. It prepends "jq" to the argument
// vector and hands it to run; the seam lets tests prove the exact vector without
// executing jq.
func jqWith(run runner, args ...string) gloo.Command[[]byte, []byte] {
	return run("jq", args...)
}
