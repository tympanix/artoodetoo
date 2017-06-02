package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Tympanix/automato/config"
	"github.com/mkideal/cli"
)

type genSecretArgs struct {
	cli.Helper
	Yes  bool   `cli:"y,yes" usage:"do not ask for permission" prompt:"this will replace any existing application secrets. Are you sure?"`
	Size int    `cli:"s,size" usage:"the size of the secret in bytes" dft:"2048"`
	Path string `cli:"f,file" usage:"the file to write the secret to" dft:".appsecret"`
}

var genSecretCmd = &cli.Command{
	Desc: "generate a new application secret",
	Name: "gensecret",
	Argv: func() interface{} { return new(genSecretArgs) },
	Fn:   doGenSecret,
}

var doGenSecret = func(ctx *cli.Context) error {
	argv := ctx.Argv().(*genSecretArgs)
	if argv.Yes {
		if err := config.GenerateSecret(argv.Path, argv.Size); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		log.Println("Generated new application secret")
	} else {
		return nil
	}
	return nil
}
