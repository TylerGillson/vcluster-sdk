package main

import (
	"os"

	"github.com/TylerGillson/vcluster-sdk/plugin"
	"github.com/loft-sh/vcluster-pull-secret-sync/syncers"
)

const (
	DefaultDestinationNamespace = "default"
	DestinationNamespaceEnvVar  = "DESTINATION_NAMESPACE"
)

func main() {
	// resolve configuration from environment variables
	destinationNamespace := os.Getenv(DestinationNamespaceEnvVar)
	if destinationNamespace == "" {
		destinationNamespace = DefaultDestinationNamespace
	}

	ctx := plugin.MustInit()
	plugin.MustRegister(syncers.NewPullSecretSyncer(ctx, destinationNamespace))
	plugin.MustStart()
}
