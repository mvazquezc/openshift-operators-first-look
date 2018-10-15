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
