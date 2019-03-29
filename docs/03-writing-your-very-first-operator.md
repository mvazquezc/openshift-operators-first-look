# Writing your very first operator

We will go through the creation of a very simple operator using the [Operator Framework SDK](https://github.com/operator-framework/operator-sdk).

The operator will deploy a [Python application](https://quay.io/mavazque/hello-api). Application sources can be found [here](../sources/python-app).

**You will need GoLang,GoLang Dep and OC client packages installed before running the following commands**

At the moment of this writing the following versions were used:

* golang-1.10.3
* dep-0.4.1
* oc v3.11.0+0cbc58b

1. Installing the Operator Framework SDK

    ~~~sh

    $ go get github.com/operator-framework/operator-sdk
    $ cd $GOPATH/src/github.com/operator-framework/operator-sdk
    $ git checkout tags/v0.6.6
    $ make dep
    $ make install

    ~~~

2. Initialize your Operator project

    As previously discussed, Operators extend the K8s API, the K8s API has different groups and is versioned. Our Operator must define a new group, a new object kind and its versioning.

    In the example below, we're creating the API group "`mario.lab`", a new object kind "`PythonAPIHw`" and its versioning "`v1alpha1`". So our operator, will take care of this combination and will act upon different events affecting the objects it is observing.

    ~~~sh

    $ mkdir -p $GOPATH/src/github.com/<user>/
    $ cd $_
    $ $GOPATH/bin/operator-sdk new <operator-name>
    $ cd <operator-name>
    e.g: $GOPATH/bin/operator-sdk new python-api-hw

    ~~~

3. Modify your Operator types (example [here](../sources/go/types.go))

    First we need to define our API endpoint and its version and kind.
    
    ~~~sh
    $ $GOPATH/bin/operator-sdk add api --api-version=<your-crd-api-group>/v1alpha1 --kind=<your-crd-object-kind>
    e.g: $GOPATH/bin/operator-sdk add api --api-version=mario.lab/v1alpha1 --kind=PythonAPIHw

    ~~~
    
    Then we need to define the structure of our new object kind, in the example `types.go` we are defining a spec property called `size` which will be used to define the number of replicas of our application and an `apiPods` status property which will be used to specify which pods are part of our application.

    ~~~sh
    $ cat $GOPATH/src/github.com/<operator-name>/pkg/apis/<api-group>/v1alpha1/<your-crd-object-kind>_types.go
    ~~~

    ~~~go
    package v1alpha1

    import (
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    )

    // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

    type PythonAPIHwList struct {
        metav1.TypeMeta `json:",inline"`
        metav1.ListMeta `json:"metadata"`
        Items           []PythonAPIHw `json:"items"`
    }

    // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

    type PythonAPIHw struct {
        metav1.TypeMeta   `json:",inline"`
        metav1.ObjectMeta `json:"metadata"`
        Spec              PythonAPIHwSpec   `json:"spec"`
        Status            PythonAPIHwStatus `json:"status,omitempty"`
    }

    type PythonAPIHwSpec struct {
        Size int32 `json:"size"`
    }
    type PythonAPIHwStatus struct {
        ApiPods []string `json:"apiPods"`
    }

    ~~~

4. Re-generate some code after modifying the Operator types

    Every time we make modifications on the operator's types, we must run the code generator as some code must be updated accordingly.

    ~~~sh
    $ $GOPATH/bin/operator-sdk generate k8s
    ~~~

5. Create and code your operator business logic (example [here](../sources/go/handler.go))

    ~~~sh
    $ $GOPATH/bin/operator-sdk add controller --api-version=<your-crd-api-group>/v1alpha1 --kind=<your-crd-object-kind>
    e.g: $GOPATH/bin/operator-sdk add controller --api-version=mario.lab/v1alpha1 --kind=PythonAPIHw
    $ vim $GOPATH/src/github.com/<operator-name>/pkg/controller/<your-crd-object-kind>/<your-crd-object-kind>_controller.go
    vim $GOPATH/src/github.com/<operator-name>/pkg/controller/pythonapihw/pythonapihw_controller.go
    
    ~~~

6. Build and Package your Operator

    ~~~sh

    $ $GOPATH/bin/operator-sdk build quay.io/<user>/<operator-image-name>:<operator-image-tag>
    e.g: $GOPATH/bin/operator-sdk build quay.io/mavazque/pythonapihw:test

    ~~~

7. Push your Operator to the Quay Registry

    ~~~sh

    $ docker push quay.io/<user>/<operator-image-name>:<operator-image-tag>

    ~~~

8. Deploy your Operator

    ~~~sh

    $ oc cluster up --enable=router,registry,web-console
    $ oc login -u system:admin
    $ oc new-project helloworld-operator
    $ oc create -f /project/path/deploy/rbac.yaml
    $ oc create -f /project/path/deploy/crd.yaml
    $ oc create -f /project/path/deploy/operator.yaml

    ~~~

9. Create a Python API HelloWorld definition

    If you have modified `types.go` and added spec properties, you must update the `cr.yaml` file accordingly. In our example, we added the following code to the existing yaml:

    ~~~yaml

    spec:
      size: 5

    ~~~

    ~~~sh

    $ oc create -f /project/path/deploy/cr.yaml

    ~~~

10. Verify the deployment

    ~~~sh

    $ oc get pods
    $ oc get svc
    $ oc get <your-crd-object-kind> -o yaml
    $ curl <svc-ip>:<svc-port>

    ~~~

**Back to [Controllers](02-controllers.md)**
