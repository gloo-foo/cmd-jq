package jq_test

import (
	jq "github.com/gloo-foo/cmd-jq/alias"
)

// ExampleJq shows how to construct a jq filter as a composable Command. The
// subprocess is not executed here — running it would require a working jq
// install — so the constructed Command is simply discarded.
func ExampleJq() {
	cmd := jq.Jq("-c", ".items[]")
	_ = cmd
}
