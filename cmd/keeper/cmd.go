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
		Run:   run,
	}

	registerFlags(&cmd)

	return cmd
}

const (
	FlagNameGeoConfig = "geo-config-path"
	FlagNameToMail    = "to-mail"
)

func registerFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(FlagNameGeoConfig, "g", "PATH", "Path to the geo configuration json file")
	cmd.MarkFlagRequired(FlagNameGeoConfig)

	cmd.Flags().StringP(FlagNameToMail, "m", "example@example.com", "The mail to which the new estate will be mailed to")
	cmd.MarkFlagRequired(FlagNameToMail)
}

func run(cmd *cobra.Command, args []string) {
	bc := parseBuildConfig()
	cmdArgs := parseCommandArgs(cmd)

	cfg := generateKeeperConfig(&bc, &cmdArgs)

	kc := keepercore.NewKeeper(cfg)

	kc.Run()
}
