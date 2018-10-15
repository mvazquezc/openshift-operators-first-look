# What is an Operator

* An Operator is a method of packaging, deploying and managing a Kubernetes application. A Kubernetes application is an application that is both deployed on Kubernetes and managed using the Kubernetes APIs and kubectl tooling.

* To be able to make the most of Kubernetes, you need a set of cohesives APIs to extend in order to service and manage your applications that run on Kubernetes. You can think of Operators as the runtime that manages this type of application on Kubernetes.

* For more information you can read the [Operator FAQ](https://coreos.com/operators/) by CoreOS.

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
