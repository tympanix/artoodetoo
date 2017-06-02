package cli

import (
	"fmt"
	"os"

	"github.com/Tympanix/artoodetoo/cli/user"
	"github.com/Tympanix/artoodetoo/types"
	"github.com/mkideal/cli"
)

var main func(types.AppArgs)

// Run the cli application
func Run(fn func(types.AppArgs)) {
	main = fn

	prog := cli.Root(appCmd, cli.Tree(helpCmd),
		user.UserCmd,
		cli.Tree(genSecretCmd))

	if err := prog.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
