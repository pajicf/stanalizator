package configs

import "embed"

//go:embed **/*.toml
var EmbeddedTomlConfigs embed.FS
