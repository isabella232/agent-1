package main

import (
	"github.com/buildkite/agent/buildkite"
	"github.com/buildkite/agent/command"
	"github.com/codegangsta/cli"
	"os"
)

var AppHelpTemplate = `Usage:

  {{.Name}} <command> [arguments...]

Available commands are:

  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Use "{{.Name}} <command> --help" for more information about a command.

`

var SubcommandHelpTemplate = `Usage:

  {{.Name}} {{if .Flags}}<command>{{end}} [arguments...]

Available commands are:

   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
Options:

   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

var CommandHelpTemplate = `{{.Description}}

Options:

   {{range .Flags}}{{.}}
   {{end}}
`

func main() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.SubcommandHelpTemplate = SubcommandHelpTemplate

	app := cli.NewApp()
	app.Name = "buildkite-agent"
	app.Version = buildkite.Version()
	app.Commands = []cli.Command{
		command.AgentStartCommand,
		{
			Name:  "artifact",
			Usage: "Upload/download artifacts from Buildkite jobs",
			Subcommands: []cli.Command{
				command.ArtifactUploadCommand,
				command.ArtifactDownloadCommand,
				command.ArtifactShasumCommand,
			},
		},
		{
			Name:  "meta-data",
			Usage: "Get/set data from Buildkite jobs",
			Subcommands: []cli.Command{
				command.MetaDataSetCommand,
				command.MetaDataGetCommand,
			},
		},
	}

	// When no sub command is used
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	// When a sub command can't be found
	app.CommandNotFound = func(c *cli.Context, command string) {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	app.Run(os.Args)
}
