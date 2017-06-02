package cli

import (
	"log"

	"github.com/Tympanix/artoodetoo/config"
	"github.com/mkideal/cli"
)

type genSecretArgs struct {
	cli.Helper
	Gen   bool   `cli:"g,generate" usage:"generate a new application secret"`
	Force bool   `cli:"force" usage:"force overwrite of existing app secret"`
	Size  int    `cli:"s,size" usage:"the size of the secret in bytes" dft:"2048"`
	Path  string `cli:"f,file" usage:"the file to write the secret to" dft:".appsecret"`
}

func (g *genSecretArgs) AutoHelp() bool {
	return !g.Gen || g.Help
}

var genSecretCmd = &cli.Command{
	Desc: "application secret manegement",
	Name: "secret",
	Argv: func() interface{} { return new(genSecretArgs) },
	Fn:   doGenSecret,
}

var doGenSecret = func(ctx *cli.Context) error {
	argv := ctx.Argv().(*genSecretArgs)
	if argv.Gen {
		if err := config.GenerateSecret(argv.Path, argv.Size, argv.Force); err != nil {
			return err
		}
		log.Println("Generated new application secret")
	} else {
		return nil
	}
	return nil
}
