package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Tympanix/automato/config"
	"github.com/mkideal/cli"
)

type addUserArgs struct {
	cli.Helper
	Username string `cli:"u,username" usage:"the username for the new user" prompt:"enter username"`
	Password string `pw:"p,password" usage:"the password for the new user" prompt:"enter password"`
	Path     string `cli:"f,file" usage:"the file to write the user information to" dft:".htpasswd"`
}

var addUserCmd = &cli.Command{
	Desc: "register a new user for the system",
	Name: "adduser",
	Argv: func() interface{} { return new(addUserArgs) },
	Fn:   doAddUser,
}

var doAddUser = func(ctx *cli.Context) error {
	argv := ctx.Argv().(*addUserArgs)
	if err := config.AddUser(argv.Path, argv.Username, argv.Password); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	log.Printf("Added user %s\n", argv.Username)
	return nil
}

// Validate implements cli.Validator interface
func (adduser *addUserArgs) Validate(ctx *cli.Context) error {
	if len(adduser.Password) < 5 {
		return fmt.Errorf("password is too short")
	}
	return nil
}
