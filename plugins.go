package quick

import (
	_ "github.com/go-micro/plugins/config/source/cli/urfave"
	_ "github.com/go-micro/plugins/config/source/file"
	_ "github.com/go-micro/plugins/config/source/http"
	_ "go-micro.dev/v5/config/source/cli"

	_ "github.com/go-micro/plugins/registry/mdns"
)
