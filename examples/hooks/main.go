package main

import (
	"github.com/TylerGillson/vcluster-sdk/plugin"
	"github.com/loft-sh/vcluster-pod-hooks/hooks"
)

func main() {
	_ = plugin.MustInit()
	plugin.MustRegister(hooks.NewPodHook())
	plugin.MustRegister(hooks.NewServiceHook())
	plugin.MustRegister(hooks.NewSecretHook())
	plugin.MustStart()
}
