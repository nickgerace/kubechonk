/*
KUBECHONK
https://github.com/nickgerace/kubechonk

Apache 2.0 License
See "LICENSE" file for more information
*/

package main

import (
	"os"

	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/nickgerace/kubechonk/pkg/cmd"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-chonk", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdChonk(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
