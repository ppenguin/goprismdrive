package main

import (
	"fmt"
	"log"
	"os"

	gpd "github.com/ppenguin/goprismdrive"
	"github.com/urfave/cli/v2"
)

var pclient *gpd.Client
var ptoken string

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "login",
				Usage:       "login with name and password (echos the (session) token)",
				Description: "You should set the env var PRISMDRIVE_TOKEN to the session token, it will be used for all other commands",
				Action:      login,
				// Flags: []cli.Flag{
				// 	&cli.StringFlag{Name: "user", Aliases: []string{"u"}},
				//     &cli.StringFlag{Name: "pass", Aliases: []string{"p"}},
				// },
			},
			{
				Name: "ls",
				// Aliases: []string{"c"},
				Usage:  "ls [remote path]",
				Action: ls,
			},
		},
	}

	ptoken = os.Getenv("PRISMDRIVE_TOKEN")
	pclient = gpd.NewClient(ptoken)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func checkToken() error {
	if ptoken == "" {
		err := fmt.Errorf("environment variable PRISMDRIVE_TOKEN not set")
		return err
	}
	return nil
}

func login(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 2 {
		return fmt.Errorf("login: arguments name and password required")
	}
	return gpd.Login(cCtx.Args().Get(0), cCtx.Args().Get(1))
}

func ls(cCtx *cli.Context) error {
	if err := checkToken(); err != nil {
		return err
	}
	return pclient.Ls()
}
