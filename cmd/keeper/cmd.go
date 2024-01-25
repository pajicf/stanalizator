package keeper

import (
	keepercore "github.com/pajicf/stanalizator/internal/keeper"
	"github.com/spf13/cobra"
)

func NewCommand() cobra.Command {
	cmd := cobra.Command{
		Use:   "keeper",
		Short: "Sets up a keeper bot",
		Long:  "Sets up a cron job that searches for new real estate on a given time loop",
		Run:   Run,
	}

	registerFlags(&cmd)

	return cmd
}

func registerFlags(cmd *cobra.Command) {

}

func Run(cmd *cobra.Command, args []string) {
	// Read cmd args config and parse it

	bc := ParseBuildConfig()

	cfg := &keepercore.Config{}

	kc := keepercore.NewKeeper(cfg)

	kc.Run()
}
