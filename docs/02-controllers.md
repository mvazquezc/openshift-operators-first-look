# Controllers

* In Kubernetes, a controller is a control loop that watches the shared state of the cluster through the API server and makes changes attempting to move the current state towards the desired one.

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

## How a Controller Works

### Control Loop

* Every controller has a Control Loop which basically does:

  * Processes every single item from the Queue
  * Picks the item an do whatever it needs to do with that item
  * Pushes the item back to the queue or ignores it
  * Updates the status to reflect the new changes
  * Starts over

**Code Examples**
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L180
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L187

### Queue

* Stuff is put into the queue
* Stuff is taken out from the queue in the Sync Loop
* Queue doesn't store objects, it stores MetaNamespaceKey
  * Key Value with namespace for the resource and name for the resource

**Code Examples**
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L111
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L188

### SharedInformer

* Is a shared data cache and distributes the data to al the Listeners that are interested in knowing about changes happening to those data.
* The most important part of the SharedInformer are the EventHandlers
  * This is how you register your interest in specific object updates (Addition, Creation, Updation, Deletion)
* The controller will look at what was sent in the EventListener and will put that object into the Queue
  * When dealing with updates sometimes the SyncLoop will actually verify if it is needed to process the data

#### Listers

* Important part of the SharedInformers, you want to use them
  * Listers vs ClientGo: Listers are designed specifically to be used within controllers, they have access to the cache while ClientGo will hit the API Server which is expensive when dealing with thousands of objects

**Code Examples**
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L252
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L274

## SyncHandler AKA Reconciliation Loop

* The first invocation of the SyncHandler will be always getting the MetaNamespaceKey to get the Namespace/Resource you want to work with
* Once the MetaNamespaceKey is ready the object is gathered from the cache
  * Because we are using a Shared Cache, the resources we're getting are not an object but a pointer to the cached object
  * If you only want to read the object, you're good to go
  * If you want to modify the object then you have to call DeepCopy on the object
  * DeepCopy is an expensive operation, make sure you will be modifying the object before DeepCopying it
* Now you will be coding your business logic

**Code Examples**
* https://github.com/kubernetes/sample-controller/blob/master/controller.go#L243

## K8s Controllers

* Cronjob controller is probably the smallest one out there
* [Sample controller](https://github.com/kubernetes/sample-controller) that will help you get started


**Continue to [Hands-on Lab](03-writing-your-very-first-operator.md)** OR **Back to [What is an Operator](01-what-is-an-operator.md)**

## Sources

* [https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html](https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html)
* [Writing Kube Controllers for Everyone](https://www.youtube.com/watch?v=AUNPLQVxvmw)
