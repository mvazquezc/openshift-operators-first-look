# Operator Lifecycle Manager

OLM, hosted at <https://github.com/operator-framework/operator-lifecycle-manager> is an operator which takes care of maintaining 'other' operators.

One of the features included is the ability to use `subscriptions`, this enables for example to subscribe to a specific update channel, so when an updated version is published, OLM takes care of performing required steps to upgrade the subscription and requirements for the updated version automatically.

For example, (from above URL):

~~~yaml
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: etcd
  namespace: local
spec:
  channel: alpha
  name: etcd
  source: rh-operators
~~~

This example, subscribes to channel `alpha` to get latest available version of `etcd`.

There's a detailed explanation of the 'internals' available at the [architecture](https://github.com/operator-framework/operator-lifecycle-manager/blob/master/Documentation/design/architecture.md) document.

From that I would highlight that OLM is composed of two operators: OLM and Catalog operators.

OLM uses Custom Resource Definitions (CRD) that introduces descriptors to indicate properties. Later, dependency resolution is done via `group`, `version`, `kind` of the CRD's, so no updates are happening at all unless they are compatible.

### Catalog operator

It manages the installation and resolution of dependencies as defined by ClusterServiceVersions. It also checks catalog updates for relevant channels and packages within to upgrade them (optionally) to the latest version (by using above mentioned 'subscription')


**Back to [Hands-on Lab](03-writing-your-very-first-operator.md)**

## Sources

* <https://github.com/operator-framework/operator-lifecycle-manager>
* <https://github.com/operator-framework/operator-lifecycle-manager/blob/master/Documentation/design/architecture.md>