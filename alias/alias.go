// Package alias provides unprefixed names for the jq command's public API.
//
//	import jq "github.com/gloo-foo/cmd-jq/alias"
//	jq.Jq("-c", ".items[]")
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-jq"
)

// Jq re-exports the constructor.
func Jq(args ...command.JqArg) gloo.Command[[]byte, []byte] { return command.Jq(args...) }

// Arg is one verbatim jq command-line argument.
type Arg = command.JqArg
