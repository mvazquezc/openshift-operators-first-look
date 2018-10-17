# What is an Operator

* An operator aims to automate actions usually performed manually, reducing the chances for errors and simplify complexity.
* An Operator is a method of packaging, deploying and managing a Kubernetes application. A Kubernetes application is an application that is both deployed on Kubernetes and managed using the Kubernetes APIs and kubectl tooling.
* To be able to make the most of Kubernetes, you need a set of cohesive APIs to extend in order to service and manage your applications that run on Kubernetes. You can think of Operators as the runtime that manages this type of application on Kubernetes.
* For more information you can read the [Operator FAQ](https://coreos.com/operators/) by CoreOS.
* Simplest operator would be one that defines how to deploy an application, and advanced one, will also automate recovery, maintenance tasks, upgrades, etc
* As of this writing, there are some operators that showcase some of the functionalities:
    * [OLM: Operator Lifecycle Manager](https://github.com/operator-framework/operator-lifecycle-manager): it can be seen as an 'operator' for 'operators', allowing operators to get defined a policy for upgrades, installation and management.
    * [etcd operator](https://github.com/coreos/etcd-operator):  considered 'beta', manages etc clusters deployed on K8s and automates tasks like create/destroy, resize, failover, rolling upgrades, backup/restore.

## What about Controllers?

* Controllers take care of routine tasks to ensure the observed state matches the desired state of the cluster.
* Each controller is responsible for a particular resource in the Kubernetes world.
* Operators use the controller pattern, but not all controllers are Operators. It’s only an Operator if it’s got:
  * Controller Pattern
  * API Extension
  * Single-App Focus

**Continue to [Controllers](02-controllers.md)**

## Sources

* [https://coreos.com/operators/](https://coreos.com/operators/)
* [https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html](https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html)
* [https://github.com/kubeflow/tf-operator/issues/300#issuecomment-357319596](https://github.com/kubeflow/tf-operator/issues/300#issuecomment-357319596)
