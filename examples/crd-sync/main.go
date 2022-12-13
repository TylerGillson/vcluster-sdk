package main

import (
	"github.com/TylerGillson/vcluster-sdk/plugin"
	"github.com/loft-sh/vcluster-example/syncers"
)

func main() {
	ctx := plugin.MustInit()
	plugin.MustRegister(syncers.NewCarSyncer(ctx))
	plugin.MustStart()
}
