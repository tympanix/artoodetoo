package user

import (
	"fmt"
	"log"
	"os"

	"github.com/Tympanix/artoodetoo/config"
	"github.com/mkideal/cli"
)

type delUserArgs struct {
	cli.Helper
	Username string `cli:"*u,username" usage:"the username for the new user"`
	Path     string `cli:"f,file" usage:"the file to write the user information to" dft:".htpasswd"`
}

var delUserCmd = &cli.Command{
	Desc: "delete an existing user from the application",
	Name: "del",
	Argv: func() interface{} { return new(delUserArgs) },
	Fn:   doDeleteUser,
}

var doDeleteUser = func(ctx *cli.Context) error {
	argv := ctx.Argv().(*delUserArgs)
	if err := config.DeleteUser(argv.Path, argv.Username); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	log.Printf("Deleted user %s\n", argv.Username)
	return nil
}
