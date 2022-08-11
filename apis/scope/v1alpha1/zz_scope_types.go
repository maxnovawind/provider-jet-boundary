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

type ScopeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScopeParameters struct {

	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler HCL but results in role resources that are unmanaged by Terraform.
	// +kubebuilder:validation:Optional
	AutoCreateAdminRole *bool `json:"autoCreateAdminRole,omitempty" tf:"auto_create_admin_role,omitempty"`

	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources that are unmanaged by Terraform.
	// +kubebuilder:validation:Optional
	AutoCreateDefaultRole *bool `json:"autoCreateDefaultRole,omitempty" tf:"auto_create_default_role,omitempty"`

	// The scope description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	// +kubebuilder:validation:Optional
	GlobalScope *bool `json:"globalScope,omitempty" tf:"global_scope,omitempty"`

	// The scope name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The scope ID containing the sub scope resource.
	// +crossplane:generate:reference:type=Scope
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScopeIDRef *v1.Reference `json:"scopeIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScopeIDSelector *v1.Selector `json:"scopeIdSelector,omitempty" tf:"-"`
}

// ScopeSpec defines the desired state of Scope
type ScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScopeParameters `json:"forProvider"`
}

// ScopeStatus defines the observed state of Scope.
type ScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Scope is the Schema for the Scopes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundaryjet}
type Scope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScopeSpec   `json:"spec"`
	Status            ScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScopeList contains a list of Scopes
type ScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scope `json:"items"`
}

// Repository type metadata.
var (
	Scope_Kind             = "Scope"
	Scope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Scope_Kind}.String()
	Scope_KindAPIVersion   = Scope_Kind + "." + CRDGroupVersion.String()
	Scope_GroupVersionKind = CRDGroupVersion.WithKind(Scope_Kind)
)

func init() {
	SchemeBuilder.Register(&Scope{}, &ScopeList{})
}