package main

import (
	"github.com/codegangsta/cli"
	"os"
)

//
func main() {
	app := cli.NewApp()
	app.Name = "hygo"
	app.Usage = "A CLI tool for HYFN project tasks"

	app.Commands = []cli.Command{
		{
			Name:      "init",
			ShortName: "i",
			Usage:     "Create or update your service credentials",
			Action: func(c *cli.Context) {
				// TODO: have it reuse existing
				p := Project{}
				p.PromptForConfig()
				p.WriteConfig()
			},
		}, {
			Name:      "config",
			ShortName: "c",
			Usage:     "View your service credentials",
			Action: func(c *cli.Context) {
				p := Project{}
				p.ReadConfig()

				b, _ := p.Conf.ToJSON()
				println(string(b))
			},
		}, {
			Name:  "hipchat_rooms",
			Usage: "List Hipchat rooms",
			Action: func(c *cli.Context) {
				p := Project{IO: ConsoleIO{}}
				p.ListHipchatRooms()

			},
		}, {
			Name:  "github_repos",
			Usage: "List Github repos",
			Action: func(c *cli.Context) {
				p := Project{IO: ConsoleIO{}}
				p.ListGithubRepos()
			},
		}, {
			Name:  "hipchat_hook",
			Usage: "Add a Hipchat+Github hook",
			Action: func(c *cli.Context) {
				p := Project{IO: ConsoleIO{}}
				p.AddGithubHipchatHook()
			},
		}, {
			Name:  "create_repo",
			Usage: "Create a new Github Repo",
			Action: func(c *cli.Context) {
				p := Project{IO: ConsoleIO{}}
				p.CreateGithubRepo()
			},
		},
	}

	app.Run(os.Args)
}
