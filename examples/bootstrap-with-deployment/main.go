package main

import (
	"github.com/TylerGillson/vcluster-sdk/plugin"
	"github.com/loft-sh/vcluster-mydeployment-example/syncers"
)

func main() {
	ctx := plugin.MustInit()
	plugin.MustRegister(syncers.NewMydeploymentSyncer(ctx))
	plugin.MustStart()
}
