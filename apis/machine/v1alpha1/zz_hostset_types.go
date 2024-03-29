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

type HostSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostSetParameters struct {

	// The host set description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The catalog for the host set.
	// +crossplane:generate:reference:type=HostCatalog
	// +kubebuilder:validation:Optional
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// +kubebuilder:validation:Optional
	HostCatalogIDRef *v1.Reference `json:"hostCatalogIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HostCatalogIDSelector *v1.Selector `json:"hostCatalogIdSelector,omitempty" tf:"-"`

	// The list of host IDs contained in this set.
	// +crossplane:generate:reference:type=Host
	// +kubebuilder:validation:Optional
	HostIds []*string `json:"hostIds,omitempty" tf:"host_ids,omitempty"`

	// +kubebuilder:validation:Optional
	HostIdsRefs []v1.Reference `json:"hostIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HostIdsSelector *v1.Selector `json:"hostIdsSelector,omitempty" tf:"-"`

	// The host set name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of host set
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// HostSetSpec defines the desired state of HostSet
type HostSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostSetParameters `json:"forProvider"`
}

// HostSetStatus defines the observed state of HostSet.
type HostSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostSet is the Schema for the HostSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundaryjet}
type HostSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSetSpec   `json:"spec"`
	Status            HostSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostSetList contains a list of HostSets
type HostSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostSet `json:"items"`
}

// Repository type metadata.
var (
	HostSet_Kind             = "HostSet"
	HostSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostSet_Kind}.String()
	HostSet_KindAPIVersion   = HostSet_Kind + "." + CRDGroupVersion.String()
	HostSet_GroupVersionKind = CRDGroupVersion.WithKind(HostSet_Kind)
)

func init() {
	SchemeBuilder.Register(&HostSet{}, &HostSetList{})
}
