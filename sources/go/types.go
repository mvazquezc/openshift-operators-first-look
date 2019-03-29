package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PythonAPIHwSpec defines the desired state of PythonAPIHw
// +k8s:openapi-gen=true
type PythonAPIHwSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
        Size int32 `json:"size"`
}

// PythonAPIHwStatus defines the observed state of PythonAPIHw
// +k8s:openapi-gen=true
type PythonAPIHwStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
        ApiPods []string `json:"apiPods"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PythonAPIHw is the Schema for the pythonapihws API
// +k8s:openapi-gen=true
type PythonAPIHw struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PythonAPIHwSpec   `json:"spec,omitempty"`
	Status PythonAPIHwStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PythonAPIHwList contains a list of PythonAPIHw
type PythonAPIHwList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PythonAPIHw `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PythonAPIHw{}, &PythonAPIHwList{})
}
