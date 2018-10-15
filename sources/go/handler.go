package stub

import (
	"context"
        "fmt"
        "github.com/python-api-hw/pkg/apis/mario/v1alpha1"
        "reflect"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
        "k8s.io/apimachinery/pkg/labels"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        appsv1 "k8s.io/api/apps/v1"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
        logrus.Infof("Inside Handler")
	switch o := event.Object.(type) {
	case *v1alpha1.PythonAPIHw:
                helloApiWorld := o
                dep := deploymentForHelloApi(helloApiWorld)
                err := sdk.Create(dep)
                logrus.Infof("Inside switch for my object")
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create deployment : %v", err)
			return err
		}
                svc := serviceForHelloApi(helloApiWorld)
                err = sdk.Create(svc)
                if err != nil && !errors.IsAlreadyExists(err) {
                        logrus.Errorf("Failed to create service : %v", err)
                        return err
                }
                // Ensure the deployment size is the same as the spec
                err = sdk.Get(dep)
                if err != nil {
                        logrus.Errorf("Failed to get deployment : %v", err)
                        return err
                }
                size := helloApiWorld.Spec.Size
                logrus.Infof("Size is set to %d, current replias %d", size, *dep.Spec.Replicas)
                if *dep.Spec.Replicas != size {
                        logrus.Infof("Need to update replicas from %d to %d", *dep.Spec.Replicas, size)
                        dep.Spec.Replicas = &size
                        err = sdk.Update(dep)
                        if err != nil {
                                logrus.Errorf("Failed to update deployment : %v", err)
                                return err
                        }
                }
                // Update Status
                apiHwStatus, err := getApiHwStatus(helloApiWorld)
                if err != nil {
                        logrus.Errorf("Failed to get ApiHwStatus : %v", err)
                }
                err = updateApiHwStatus(helloApiWorld, apiHwStatus)
                if err != nil {
                        logrus.Errorf("Failed to update status : %v", err)
                }
	}
	return nil
}

func getLabelsForApiHw(name string) map[string]string {
        return map[string]string{"app": "api-hello-world", "name": name}
}

// get the status for our object type
func getApiHwStatus(h *v1alpha1.PythonAPIHw) (*v1alpha1.PythonAPIHwStatus, error) {
        pods := &corev1.PodList{
                TypeMeta: metav1.TypeMeta{
                        Kind:       "Pod",
                        APIVersion: "v1",
                },
        }
        sel := getLabelsForApiHw(h.Name)
	opt := &metav1.ListOptions{LabelSelector: labels.SelectorFromSet(sel).String()}
	err := sdk.List(h.GetNamespace(), pods, sdk.WithListOptions(opt))
        if err != nil {
                logrus.Errorf("Failted to get pods : %v", err)
                return nil, fmt.Errorf("Failted to get pods : %v", err)
        }
        var apiPods []string
        for _, p := range pods.Items {
                logrus.Infof("Pod name is %s", p.GetName())
                if p.Status.Phase != corev1.PodRunning || p.DeletionTimestamp != nil {
                        logrus.Errorf("Pod %s is terminating, not adding it to pods list", p.GetName()) 
                } else {
                        logrus.Infof("Adding pod %s to running pods list", p.GetName())
                        apiPods = append(apiPods, p.GetName())
                }
        }
        return &v1alpha1.PythonAPIHwStatus{
                ApiPods: apiPods,
        }, nil
}


func updateApiHwStatus(h *v1alpha1.PythonAPIHw, status *v1alpha1.PythonAPIHwStatus) error {
        logrus.Infof("Updating PythonApiHw Status")
	// don't update the status if there aren't any changes.
	if reflect.DeepEqual(h.Status, *status) {
                logrus.Infof("Status has not changed")
		return nil
	}
	h.Status = *status
        logrus.Infof("Status has changed, we need to update it")
        err := sdk.Update(h)
	return err
}

// serviceForHelloApi returns a Service Object
func serviceForHelloApi(h *v1alpha1.PythonAPIHw) *corev1.Service {
        labels := getLabelsForApiHw(h.Name)
        svc := &corev1.Service{
                TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
                ObjectMeta: metav1.ObjectMeta{
			Name:      h.Name,
			Namespace: h.Namespace,
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
        return svc
}

// deploymentForHelloApi returns a HelloApi Deployment Object
func deploymentForHelloApi(h *v1alpha1.PythonAPIHw) *appsv1.Deployment {
        labels := getLabelsForApiHw(h.Name)
        replicas := h.Spec.Size
        dep := &appsv1.Deployment{
                TypeMeta: metav1.TypeMeta{
                        APIVersion: "apps/v1",
                        Kind:       "Deployment",
                },
                ObjectMeta: metav1.ObjectMeta{
                        Name:      h.Name,
                        Namespace: h.Namespace,
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
						Image:   "quay.io/mavazque/hello-api",
						Name:    "api-hello-world",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 5000,
							Name:          "api-hello-world",
						}},
					}},
				},
			},
		},
	}
        return dep
}

