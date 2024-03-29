/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TargetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TargetParameters struct {

	// A list of application credential source ID's.
	// +kubebuilder:validation:Optional
	ApplicationCredentialSourceIds []*string `json:"applicationCredentialSourceIds,omitempty" tf:"application_credential_source_ids,omitempty"`

	// The default port for this target.
	// +kubebuilder:validation:Optional
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// The target description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of host source ID's.
	// +crossplane:generate:reference:type=HostSet
	// +kubebuilder:validation:Optional
	HostSourceIds []*string `json:"hostSourceIds,omitempty" tf:"host_source_ids,omitempty"`

	// +kubebuilder:validation:Optional
	HostSourceIdsRefs []v1.Reference `json:"hostSourceIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HostSourceIdsSelector *v1.Selector `json:"hostSourceIdsSelector,omitempty" tf:"-"`

	// The target name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	// +crossplane:generate:reference:type=github.com/maxnovawind/provider-jet-boundary/apis/domain/v1alpha1.Scope
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScopeIDRef *v1.Reference `json:"scopeIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScopeIDSelector *v1.Selector `json:"scopeIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SessionConnectionLimit *float64 `json:"sessionConnectionLimit,omitempty" tf:"session_connection_limit,omitempty"`

	// +kubebuilder:validation:Optional
	SessionMaxSeconds *float64 `json:"sessionMaxSeconds,omitempty" tf:"session_max_seconds,omitempty"`

	// The target resource type.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Boolean expression to filter the workers for this target
	// +kubebuilder:validation:Optional
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

// TargetSpec defines the desired state of Target
type TargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetParameters `json:"forProvider"`
}

// TargetStatus defines the observed state of Target.
type TargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Target is the Schema for the Targets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundaryjet}
type Target struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetSpec   `json:"spec"`
	Status            TargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetList contains a list of Targets
type TargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Target `json:"items"`
}

// Repository type metadata.
var (
	Target_Kind             = "Target"
	Target_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Target_Kind}.String()
	Target_KindAPIVersion   = Target_Kind + "." + CRDGroupVersion.String()
	Target_GroupVersionKind = CRDGroupVersion.WithKind(Target_Kind)
)

func init() {
	SchemeBuilder.Register(&Target{}, &TargetList{})
}
