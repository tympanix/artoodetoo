package cli

import "github.com/mkideal/cli"

// AppArgs are the application arguments to run the application
type AppArgs struct {
	cli.Helper
	PortNumber int    `cli:"p,port" usage:"the port the application will listen on" dft:"2800"`
	Htpasswd   string `cli:"w,password" usage:"the htpasswd file to use for authentication" dft:".htpasswd"`
	Secret     string `cli:"s,secret" usage:"the secret to use for the application" dft:".appsecret"`
}

// Port is the port number for the application
func (a *AppArgs) Port() int {
	return a.PortNumber
}

// HtpasswdPath is the path for the htpasswd file
func (a *AppArgs) HtpasswdPath() string {
	return a.Htpasswd
}

// SecretPath is the path for the application secret
func (a *AppArgs) SecretPath() string {
	return a.Secret
}

var helpCmd = cli.HelpCommand("display help information")

var appCmd = &cli.Command{
	Desc: "starts the artoodetoo application server",
	Argv: func() interface{} { return new(AppArgs) },
	Fn: func(ctx *cli.Context) error {
		args := ctx.Argv().(*AppArgs)
		main(args)
		return nil
	},
}
