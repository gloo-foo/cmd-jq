package command

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// JqArg is one verbatim jq command-line argument: a flag (`-c`), a filter
// (`.items[]`), or any other token jq itself interprets.
type JqArg string

// runner forks an external process named name with args, returning the streaming
// Command that drives it. patterns.Subprocess is the production runner; tests
// inject a fake to assert the exact argument vector without forking jq.
//
// name and args are patterns.ProcessName/ProcessArg because the vector flows
// verbatim into patterns.Subprocess (and ultimately os/exec), whose signature
// this wrapper does not control.
type runner func(name patterns.ProcessName, args ...patterns.ProcessArg) gloo.Command[[]byte, []byte]

// Jq returns a Command that forks jq with the given filter and arguments,
// streaming pipeline input to jq's stdin and jq's stdout back into the pipeline.
// Every argument is passed through verbatim — jq, not this wrapper, interprets
// them.
func Jq(args ...JqArg) gloo.Command[[]byte, []byte] {
	return jqWith(patterns.Subprocess, args...)
}

// jqWith is Jq with an injectable runner. It prepends "jq" to the argument
// vector and hands it to run; the seam lets tests prove the exact vector without
// executing jq.
func jqWith(run runner, args ...JqArg) gloo.Command[[]byte, []byte] {
	return run("jq", processArgs(args)...)
}

// processArgs renders the typed jq argument vector into the patterns.ProcessArg
// vector os/exec consumes.
func processArgs(args []JqArg) []patterns.ProcessArg {
	out := make([]patterns.ProcessArg, len(args))
	for i, a := range args {
		out[i] = patterns.ProcessArg(a)
	}
	return out
}
