package keeper

import (
	"bytes"
	"github.com/pajicf/stanalizator/configs"
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

func ParseBuildConfig() BuildConfig {
	data, err := configs.Embedded.ReadFile("keeper/config.toml")
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewReader(data))

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
