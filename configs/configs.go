package configs

import "embed"

//go:embed **/*.toml
var Embedded embed.FS
