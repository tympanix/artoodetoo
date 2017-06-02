package user

import "github.com/mkideal/cli"

// UserCmd is a command tree for cli user commands
var UserCmd = cli.Tree(userCmd,
	cli.Tree(addUserCmd),
	cli.Tree(delUserCmd),
	cli.Tree(lsUserCmd))

type userT struct {
	Help bool `cli:"h,help" usage:"show help"`
}

func (argv *userT) AutoHelp() bool {
	return true
}

var userCmd = &cli.Command{
	Desc: "edit user configuration for the application",
	Name: "user",
	Argv: func() interface{} { return new(userT) },
}
