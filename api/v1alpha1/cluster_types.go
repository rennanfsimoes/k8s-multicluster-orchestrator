package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	Region   string `json:"region"`
	Provider string `json:"provider"`
	Version  string `json:"version"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	Phase string `json:"phase"`
	Ready bool   `json:"ready"`
}
