package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "APP_LANG",
		},
	}
	app.Action = func(c *cli.Context) {
		name := "someone"
		if len(c.Args()) > 0 {
			name = c.Args()[0]
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
	}
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) {
				fmt.Println("added task: ", c.Args().First())
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) {
				fmt.Println("completed task: ", c.Args().First())
			},
		},
	}

	app.RunAndExitOnError()
}
