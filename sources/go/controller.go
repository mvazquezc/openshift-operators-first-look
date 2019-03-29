package pythonapihw

import (
	"context"
        "reflect"
	mariov1alpha1 "github.com/mvazquezc/python-api-hw/pkg/apis/mario/v1alpha1"
        appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_pythonapihw")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new PythonAPIHw Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePythonAPIHw{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("pythonapihw-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PythonAPIHw
	err = c.Watch(&source.Kind{Type: &mariov1alpha1.PythonAPIHw{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner PythonAPIHw
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &mariov1alpha1.PythonAPIHw{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcilePythonAPIHw{}

// ReconcilePythonAPIHw reconciles a PythonAPIHw object
type ReconcilePythonAPIHw struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a PythonAPIHw object and makes changes based on the state read
// and what is in the PythonAPIHw.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcilePythonAPIHw) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling PythonAPIHw")

	// Fetch the PythonAPIHw instance
	pythonApiHw := &mariov1alpha1.PythonAPIHw{}
	err := r.client.Get(context.TODO(), request.NamespacedName, pythonApiHw)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
                        reqLogger.Info("PythonApiHw resource not found. Ignoring since object must be deleted.")
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
                reqLogger.Error(err, "Failed to get PythonApiHw.")
		return reconcile.Result{}, err
	}

        // Check if this Deployment already exists, if not create a new one
        deploymentFound := &appsv1.Deployment{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: pythonApiHw.Name, Namespace: pythonApiHw.Namespace}, deploymentFound)
        if err != nil && errors.IsNotFound(err) {
                // Define a new Deployment object
                dep := r.deploymentForHelloApi(pythonApiHw)
                reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
                err = r.client.Create(context.TODO(), dep)
                if err != nil {
                        reqLogger.Error(err, "Failed to create new Deployment.", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
                        return reconcile.Result{}, err
                }
                // Deployment created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
        } else if err != nil {
                reqLogger.Error(err, "Failed to get Deployment.")
                return reconcile.Result{}, err
        } else {
                // Deployment already exists - don't requeue
                reqLogger.Info("Skip reconcile: Deployment already exists", "Deployment.Namespace", deploymentFound.Namespace, "Deployment.Name", deploymentFound.Name)
        }

        // Check if this Service already exists, if not create a new one
        serviceFound := &corev1.Service{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: pythonApiHw.Name, Namespace: pythonApiHw.Namespace}, serviceFound)
        if err != nil && errors.IsNotFound(err) {
                // Define a new Service object
                svc := r.serviceForHelloApi(pythonApiHw)
                reqLogger.Info("Creating a new Service", "Service.Namespace", svc.Namespace, "Service.Name", svc.Name)
                err = r.client.Create(context.TODO(), svc)
                if err != nil {
                        reqLogger.Error(err, "Failed to create new Service.", "Service.Namespace", svc.Namespace, "Svc.Name", svc.Name)
                        return reconcile.Result{}, err
                }
                // Service created successfully - return and requeue
                return reconcile.Result{Requeue: true}, nil
        } else if err != nil {
                reqLogger.Error(err, "Failed to get Service.")
                return reconcile.Result{}, err
        } else {
                // Service already exists - don't requeue
                reqLogger.Info("Skip reconcile: Service already exists", "Service.Namespace", serviceFound.Namespace, "Service.Name", serviceFound.Name)
        }

        // Ensure current state matches the desired state TL;DR deployment size is the same as spec size

        // Get PythonApiHw size
        size := pythonApiHw.Spec.Size

        // Check if deployment replicas and size are equal
        reqLogger.Info("Checking if current state matches desired state")
        if *deploymentFound.Spec.Replicas != size {
                reqLogger.Info("Current state does not match desired state, updating replicas")
                deploymentFound.Spec.Replicas = &size
                err = r.client.Update(context.TODO(), deploymentFound)
                if err != nil {
                        reqLogger.Error(err, "Failed to update Deployment.", "Deployment.Namespace", deploymentFound.Namespace, "Deployment.Name", deploymentFound.Name)
			return reconcile.Result{}, err
                }
                // Spec updated - return and requeue
		return reconcile.Result{Requeue: true}, nil
        }
        // Update the PythonApiHw status with the pod names
        // List the pods for this PythonApiHw Deployment
        podList := &corev1.PodList{}
        labelSelector := labels.SelectorFromSet(getLabelsForApiHw(pythonApiHw.Name))
	listOps := &client.ListOptions{
		Namespace:     pythonApiHw.Namespace,
		LabelSelector: labelSelector,
	}
        err = r.client.List(context.TODO(), listOps, podList)
        if err != nil {
		reqLogger.Error(err, "Failed to list pods.", "PythonApiHw.Namespace", pythonApiHw.Namespace, "PythonApiHw.Name", pythonApiHw.Name)
		return reconcile.Result{}, err
	}
	podNames := getPodNames(podList.Items)

        // Update status.apiPods if needed
        if !reflect.DeepEqual(podNames, pythonApiHw.Status.ApiPods) {
                pythonApiHw.Status.ApiPods = podNames
                err := r.client.Status().Update(context.TODO(), pythonApiHw)
                if err != nil {
			reqLogger.Error(err, "Failed to update PythonApiHw status.")
			return reconcile.Result{}, err
		}
        } else {
                reqLogger.Info("Status has not changed")
        }

	return reconcile.Result{}, nil
}

// serviceForHelloApi returns a Service Object
func (r *ReconcilePythonAPIHw) serviceForHelloApi(pythonApiHw *mariov1alpha1.PythonAPIHw) *corev1.Service {
	labels := getLabelsForApiHw(pythonApiHw.Name)
	svc := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      pythonApiHw.Name,
			Namespace: pythonApiHw.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeLoadBalancer,
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 5000,
				},
			},
		},
	}
        // Set HelloApi instance as the owner and controller
        controllerutil.SetControllerReference(pythonApiHw, svc, r.scheme)
	return svc
}

// deploymentForHelloApi returns a HelloApi Deployment Object
func (r *ReconcilePythonAPIHw) deploymentForHelloApi(pythonApiHw *mariov1alpha1.PythonAPIHw) *appsv1.Deployment {
	labels := getLabelsForApiHw(pythonApiHw.Name)
	replicas := pythonApiHw.Spec.Size

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      pythonApiHw.Name,
			Namespace: pythonApiHw.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "quay.io/mavazque/hello-api",
						Name:  "api-hello-world",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 5000,
							Name:          "api-hello-world",
						}},
					}},
				},
			},
		},
	}
        // Set HelloApi instance as the owner and controller
	controllerutil.SetControllerReference(pythonApiHw, dep, r.scheme)
	return dep
}

func getLabelsForApiHw(name string) map[string]string {
	return map[string]string{"app": "api-hello-world", "name": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}
