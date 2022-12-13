<br>
<a href="https://www.vcluster.com"><img src="hack/vcluster-logo-dark.svg"></a>

### **[Website](https://www.vcluster.com)** • **[Documentation](https://www.vcluster.com/docs/what-are-virtual-clusters)** • **[Blog](https://loft.sh/blog)** • **[Twitter](https://twitter.com/loft_sh)** • **[Slack](https://slack.loft.sh/)**

This project holds the official vcluster SDK, which can be used to develop plugins for [vcluster](https://github.com/loft-sh/vcluster). Plugins are a feature to extend the capabilities of vcluster. They allow you to add custom functionality, such as:

1. Syncing specific resources from or to the virtual clusters including cluster scoped resources like cluster roles.
2. Syncing custom resources from or to the virtual cluster.
3. Deploying resources on virtual cluster startup, such as CRDs, applications etc.
4. Manage resources and applications inside the host or virtual cluster.
5. Enforcing certain restrictions on synced resources or extending the existing syncers of vcluster.
6. Any other operator use case that could benefit from having access to the virtual cluster as well the host cluster at the same time.

To learn more about plugins, please take a look at the [vcluster documentation](https://www.vcluster.com/docs/plugins/overview) or [our plugin examples](https://github.com/TylerGillson/vcluster-sdk/tree/main/examples).
