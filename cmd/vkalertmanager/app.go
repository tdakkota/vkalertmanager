package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tdakkota/vkalertmanager/pkg/server"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type App struct {
	server server.HookServer
	logger zerolog.Logger
}

func NewApp() *App {
	return &App{}
}

func (app *App) version(c *cli.Context) error {
	_, err := fmt.Printf("vkalertmanager %s, commit %s, built at %s by %s", version, commit, date, builtBy)
	return err
}

func (app *App) run(c *cli.Context) error {
	err := app.setup(c)
	if err != nil {
		return err
	}

	return app.server.Run(c.Context)
}

func (app *App) commands() []*cli.Command {
	commands := []*cli.Command{
		{
			Name:        "run",
			Description: "runs vkalertmanager",
			Flags:       app.flags(),
			Action:      app.run,
			Before:      app.setup,
		},
		{
			Name:        "version",
			Description: "prints version",
			Action:      app.version,
		},
	}

	app.addFileConfig("config.file", commands[0])
	return commands
}

func (app *App) cli() *cli.App {
	cliApp := &cli.App{
		Name:     "vkautoreader",
		Usage:    "vkautoreader reads vk messages",
		Commands: app.commands(),
	}

	return cliApp
}

func (app *App) Run(args []string) error {
	return app.cli().Run(args)
}
