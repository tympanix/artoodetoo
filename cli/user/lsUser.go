package user

import (
	"fmt"
	"os"

	"github.com/Tympanix/artoodetoo/config"
	"github.com/mkideal/cli"
)

type lsUserArgs struct {
	cli.Helper
	Path string `cli:"f,file" usage:"the file to write the user information to" dft:".htpasswd"`
}

var lsUserCmd = &cli.Command{
	Desc: "list users of the application",
	Name: "ls",
	Argv: func() interface{} { return new(lsUserArgs) },
	Fn:   doListUsers,
}

var doListUsers = func(ctx *cli.Context) error {
	argv := ctx.Argv().(*lsUserArgs)
	users, err := config.GetUsers(argv.Path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for username := range users {
		fmt.Println(username)
	}
	return nil
}
