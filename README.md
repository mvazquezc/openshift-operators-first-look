# ETCD Operator DEMO

In this first demo, the etcd operator will be deployed on our OpenShift Cluster. Once the operator is deployed we can create etcd clusters, definitions, and the operator will take care of those and make the required actions to ensure the cluster state is what we have asked for.

1. Clone the etcd operator repository
    ~~~sh
    $ git clone https://github.com/coreos/etcd-operator
    ~~~
2. Create a new project where the operator will be deployed
    ~~~sh
    $ oc new-project etcdoperator
    ~~~
3. The operator needs some permissions, the create_role.sh script will take care of that
    ~~~sh
    $ etcd-operator/example/rbac/create_role.sh --role-name=etcd-operator --namespace=etcdoperator --role-binding-name=etcd-operator
    ~~~
4. Create the operator's deployment in order to deploy the etcd operator
    ~~~sh
    $ oc create -f etcd-operator/example/deployment.yaml
    ~~~
5. The operator has been programmed to create the CRD by itself
    ~~~sh
    $ oc get crd
    ~~~
6. Create an etcd cluster definition
    ~~~sh
    $ oc create -f etcd-operator/example/example-etcd-cluster.yaml
    ~~~
7. Check the etcd cluster creation
    ~~~sh
    $ oc get pods -w
    ~~~
8. Try to connect to the cluster and insert a key in etcd
    ~~~sh
    $ docker run -ti --rm -e ETCDCTL_API=3 -e ETCDCTL_ENDPOINTS=http://$(oc get svc example-etcd-cluster-client -o jsonpath="{.spec.clusterIP}" -n etcdoperator):$(oc get svc example-etcd-cluster-client -o jsonpath="{.spec.ports[0].port}") centos:7 /bin/bash
    $ yum install -y etcd
    $ etcdctl put ostack meetup
    ~~~
9. Edit the cluster definition to scale it up
    ~~~sh
    $ oc patch etcdclusters example-etcd-cluster -p '{"spec":{"size":5}}' --type='merge'
    $ oc get pods -w
    ~~~
10. That's it!


The operator has way more to offer, feel free to review the operator docs and continue trying features.
