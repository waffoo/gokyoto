package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	//if len(os.Args) < 2 {
	//fmt.Println("usage: ./gokyoto [url]")
	//os.Exit(0)
	//}
	app := cli.NewApp()
	app.Name = "GoKyoTo"
	app.Usage = "CLI tools for competitive programming written in golang"
	app.Version = "0.0.0"

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:    "test",
			Aliases: []string{"te"},
			Usage:   "test command: test [url]",
			Action: func(c *cli.Context) error {
				if c.Args().Len() == 0 {
					cli.ShowAppHelp(c)
				} else {
					TestCmd(c.Args().Slice())
				}
				return nil
			},
		},
	}

	app.Run(os.Args)

}
