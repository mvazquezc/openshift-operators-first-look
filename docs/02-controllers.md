# Controllers

* In Kubernetes, a controller is a control loop that watches the shared state of the cluster through the API server and makes changes attempting to move the current state towards the desired state.

* Examples of controllers:

  * Replication Controller
  * Namespace Controller
  * ServiceAccount Controller

## Controller Components

* There are two main components of a controller: `Informer/SharedInformer` and `Workerqueue`.

* **Informer**

  * In order to retrieve an object’s information, the controller sends a request to Kubernetes API server. However, querying the API repeatedly can become expensive.
  * Additionally, the controller doesn’t really want to send requests continuously. It only cares about events when the object has been created, modified or deleted.
  * Not much used in the current Kubernetes (instead SharedInformer is used)

* **SharedInformer**

  * The informer creates a local cache of a set of resources only used by itself. But, in Kubernetes, there is a bundle of controllers running and caring about multiple kinds of resources.
  * In this case, the `SharedInformer` helps to create a single shared cache among controllers.

* **Workqueue**
 
  * The `SharedInformer` can’t track what each controller is up to (because it’s shared), so the controller must provider its own queuing and retrying mechanism (if required).
  * Whenever a resource changes, the Resource Event Handler puts a key to the `Workqueue`.


**Continue to [Hands-on Lab](03-writing-your-very-first-operator.md)** OR **Back to [What is an Operator](01-what-is-an-operator.md)**

## Sources

* [https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html](https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html)
