package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// Context ...
type Context struct {
	Debug bool
}

type zipcm struct {
	Abbrevation string   `arg:"" optional:"" help:"Abbrevation for zip command."`
	Shared      bool     `flag:"" help:"Make command shared."`
	Editable    bool     `flag:"" help:"Make command editable. Not editable commands can't be updated."`
	Commands    []string `flag:"" sep:"," help:"Command array for zip commands. It can be done dynamically. If no command given the program will ask input step by step."`
	All         bool     `flag:"" short:"a" help:"List all available zip commands."`
	Oneline     bool     `flag:"" short:"o" help:"List zip commands in oneline."`
	Detail      bool     `flag:"" help:"List zipp commands with all commands."`
	DetailOne   string   `flag:"" placeholder:"UUID" help:"List detailed zipp command with the given UUID."`
	D           string   `flag:"" placeholder:"UUID" name:"D" help:"Delete the zipped command even if it is not editable. If you are not the author you can't delete it."`
	Update      string   `flag:"" placeholder:"UUID" help:"Update the zipped command. It opens the update menu if there is a zip command exist with the UUID given found."`
	Delete      string   `flag:"" placeholder:"UUID" short:"d" help:"Delete the zipped command. If command is not editable you can't delete it. You need to use -D for delete it."`
}

type ziprn struct {
	Abbrevation string `arg:"" help:"Abbrevation of the command to run. "`
}

func (zpcm *zipcm) Run(ctx *Context) error {
	if len(zpcm.Abbrevation) > 0 {
		// create
	} else {

	}
	fmt.Println(zpcm)
	fmt.Println(len(zpcm.Abbrevation))
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	ZipCm  zipcm `cmd:"" help:"zipcmd"`
	ZipRun ziprn `cmd:"" help:"run zip command."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
