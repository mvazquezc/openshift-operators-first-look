# Custom Operator Demo

In this demo an operator will be created using the Operator Framework SDK. We will be deploying a Python application (<https://quay.io/mavazque/hello-api>, the application itself is in files folder) that exposes an API that returns "Hello World".

**You will need GoLang,GoLang Dep and OC client packages installed before running the following commands**

1. Installing the Operator Framework SDK

    ~~~sh
    $ go get github.com/operator-framework/operator-sdk
    $ cd $GOPATH/src/github.com/operator-framework/operator-sdk
    $ git checkout tags/v0.6.0
    $ make dep
    $ make install
    ~~~

2. Initialize your Operator project

    ~~~sh
    $ mkdir -p $GOPATH/src/github.com/<user>/
    $ cd $_
    $ $GOPATH/bin/operator-sdk new <operator-name>
    $ cd <operator-name>
    ~~~

3. Modify your Operator types (example [here](./sources/go/types.go))

    ~~~sh
    $ $GOPATH/bin/operator-sdk add api --api-version=<your-crd-api-group>/v1alpha1 --kind=<your-crd-object-kind>
    $ vim $GOPATH/src/github.com/<operator-name>/pkg/apis/<api-group>/<api-version>/<your-crd-object-kind>_types.go
    ~~~

4. Re-generate some code after modifying the Operator types

    Every time we make modifications on the operator's types, we must run the code generator as some code must be updated accordingly.

    ~~~sh
    $ $GOPATH/bin/operator-sdk generate k8s
    ~~~

5. Create and code your operator business logic (example [here](./sources/go/controller.go))

    ~~~sh
    $ $GOPATH/bin/operator-sdk add controller --api-version=<your-crd-api-group>/v1alpha1 --kind=<your-crd-object-kind>
    $ vim $GOPATH/src/github.com/<operator-name>/pkg/controller/<your-crd-object-kind>/<your-crd-object-kind>_controller.go
    ~~~

6. Build and Package your Operator

    ~~~sh
    $ $GOPATH/bin/operator-sdk build quay.io/<user>/<operator-image-name>:<operator-image-tag>
    ~~~

7. Push your Operator to the Quay Registry

    ~~~sh
    $ docker push quay.io/<user>/<operator-image-name>:<operator-image-tag>
    ~~~

8. Create a Namespace for deploying your operator and deploy the required RBAC

    ~~~sh
    $ oc create ns helloworld-operator
    $ oc -n helloworld-operator create -f /path/to/operator/project/deploy/role.yaml
    $ oc -n helloworld-operator create -f /path/to/operator/project/deploy/role_binding.yaml
    $ oc -n helloworld-operator create -f /path/to/operator/project/deploy/service_account.yaml
    ~~~

9. Load the CustomResourceDefinition for your new type

   ~~~sh
   $ oc create -f /path/to/operator/project/deploy/crds/mario_v1alpha1_pythonapihw_crd.yaml   
   ~~~

10. Configure the operator deployment to use your operator's image and deploy it

    ~~~sh
    $ sed -i "s/REPLACE_IMAGE/<your_image>/g" /path/to/operator/project/deploy/operator.yaml
    $ oc -n helloworld-operator create -f /path/to/operator/project/deploy/operator.yaml
    ~~~

11. Create a Python API HelloWorld definition

    ~~~sh
    $ oc -n helloworld-operator create -f /path/to/operator/project/deploy/cr.yaml
    ~~~

12. Verify the deployment

    ~~~sh
    $ oc get pods
    $ oc get svc
    $ oc get <your-cr-object> -o yaml
    $ curl <svc-ip>:<svc-port>
    ~~~
    
13. Cleanup

    ~~~sh
    $ oc delete ns helloworld-operator
    $ oc delete -f /path/to/operator/project/deploy/crds/mario_v1alpha1_pythonapihw_crd.yaml
    ~~~
