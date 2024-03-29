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

type HostCatalogObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostCatalogParameters struct {

	// The host catalog description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The host catalog name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope ID in which the resource is created.
	// +crossplane:generate:reference:type=github.com/maxnovawind/provider-jet-boundary/apis/domain/v1alpha1.Scope
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScopeIDRef *v1.Reference `json:"scopeIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScopeIDSelector *v1.Selector `json:"scopeIdSelector,omitempty" tf:"-"`

	// The host catalog type. Only `static` is supported.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// HostCatalogSpec defines the desired state of HostCatalog
type HostCatalogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostCatalogParameters `json:"forProvider"`
}

// HostCatalogStatus defines the observed state of HostCatalog.
type HostCatalogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostCatalogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostCatalog is the Schema for the HostCatalogs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundaryjet}
type HostCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostCatalogSpec   `json:"spec"`
	Status            HostCatalogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostCatalogList contains a list of HostCatalogs
type HostCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostCatalog `json:"items"`
}

// Repository type metadata.
var (
	HostCatalog_Kind             = "HostCatalog"
	HostCatalog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostCatalog_Kind}.String()
	HostCatalog_KindAPIVersion   = HostCatalog_Kind + "." + CRDGroupVersion.String()
	HostCatalog_GroupVersionKind = CRDGroupVersion.WithKind(HostCatalog_Kind)
)

func init() {
	SchemeBuilder.Register(&HostCatalog{}, &HostCatalogList{})
}
