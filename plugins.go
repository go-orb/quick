package quick

import (
	_ "github.com/go-orb/plugins/config/source/cli/urfave"
	_ "github.com/go-orb/plugins/config/source/file"
	_ "github.com/go-orb/plugins/config/source/http"
	_ "github.com/go-orb/go-orb/config/source/cli"

	_ "github.com/go-orb/plugins/registry/mdns"
)
