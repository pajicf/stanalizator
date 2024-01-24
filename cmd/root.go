package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type RootCommand struct {
	cmd cobra.Command
}

func NewRootCommand() RootCommand {
	rootCmd := cobra.Command{
		Use:   "stanalizator",
		Short: "Real estate scraping toolbox",
		Long:  "Stanalizator is a real estate scraping toolbox for Serbian territory. Search for apartments, land, houses etc.",
	}

	return RootCommand{
		cmd: rootCmd,
	}
}

func (rc *RootCommand) initRootCommand() {
	return
}

func (rc *RootCommand) Execute() {
	rc.initRootCommand()
	err := rc.cmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
