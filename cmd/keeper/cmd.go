package keeper

import (
	keepercore "github.com/pajicf/stanalizator/internal/keeper"
	"github.com/spf13/cobra"
)

func NewCommand() cobra.Command {
	return cobra.Command{
		Use:   "keeper",
		Short: "Sets up a keeper bot",
		Long:  "Sets up a cron job that searches for new real estate on a given time loop",
		Run:   Run,
	}
}

func Run(cmd *cobra.Command, args []string) {
	cfg := &keepercore.Config{}

	kc := keepercore.NewKeeper(cfg)

	kc.Run()
}
