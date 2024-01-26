package keeper

import (
	"bytes"
	"github.com/pajicf/stanalizator/configs"
	keepercore "github.com/pajicf/stanalizator/internal/keeper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type BuildConfig struct {
	refreshRate        int
	emailJSApiURL      string
	emailJSUserID      string
	emailJSServiceID   string
	emailJSTemplateID  string
	emailJSAccessToken string
	haloApiURL         string
}

type CommandArgs struct {
	geoConfigPath string
	toMail        string
}

func parseBuildConfig() BuildConfig {
	data, err := configs.Embedded.ReadFile("keeper/config.toml")
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("toml")
	err = viper.ReadConfig(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	refreshRate := viper.GetInt("cron.refresh_rate_minutes")
	emailJSApiUrl := viper.GetString("emailjs.api_url")
	emailJSUserID := viper.GetString("emailjs.user_id")
	emailJSServiceID := viper.GetString("emailjs.service_id")
	emailJSTemplateID := viper.GetString("emailjs.template_id")
	emailJSAccessToken := viper.GetString("emailjs.access_token")
	haloApiURL := viper.GetString("halo.api_url")

	return BuildConfig{
		refreshRate,
		emailJSApiUrl,
		emailJSUserID,
		emailJSServiceID,
		emailJSTemplateID,
		emailJSAccessToken,
		haloApiURL,
	}
}

func parseCommandArgs(cmd *cobra.Command) CommandArgs {
	geoConfigPath, err := cmd.Flags().GetString(FlagNameGeoConfig)
	if err != nil {
		panic(err)
	}

	toMail, err := cmd.Flags().GetString(FlagNameToMail)
	if err != nil {
		panic(err)
	}

	return CommandArgs{
		geoConfigPath: geoConfigPath,
		toMail:        toMail,
	}
}

func generateKeeperConfig(bc *BuildConfig, cmdArgs *CommandArgs) keepercore.Config {
	return keepercore.Config{
		GeoConfigPath:      cmdArgs.geoConfigPath,
		ToMail:             cmdArgs.toMail,
		EmailJSUserID:      bc.emailJSUserID,
		EmailJSServiceID:   bc.emailJSServiceID,
		EmailJSTemplateID:  bc.emailJSTemplateID,
		EmailJSAccessToken: bc.emailJSAccessToken,
		HaloApiURL:         bc.haloApiURL,
	}
}
