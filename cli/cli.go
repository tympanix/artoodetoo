package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Tympanix/automato/config"
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Username string `cli:"u,username" usage:"username" prompt:"type username"`
	Password string `pw:"p,password" usage:"password" prompt:"type password"`
}

var main func()

// Run the cli application
func Run(fn func()) {
	main = fn

	prog := cli.Root(root, cli.Tree(help), cli.Tree(adduser))

	if err := prog.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

type rootT struct {
	cli.Helper
	Port int `cli:"port" usage:"the port the application should run on"`
}

var help = cli.HelpCommand("display help information")

var root = &cli.Command{
	Desc: "run the artoodetoo server",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		main()
		return nil
	},
}

var adduser = &cli.Command{
	Desc: "register a new user for the system",
	Name: "adduser",
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if err := config.AddUser(argv.Username, argv.Password); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		log.Printf("Added user %s\n", argv.Username)
		return nil
	},
}

// Validate implements cli.Validator interface
func (adduser *argT) Validate(ctx *cli.Context) error {
	if len(adduser.Password) < 5 {
		return fmt.Errorf("password is too short")
	}
	return nil
}
