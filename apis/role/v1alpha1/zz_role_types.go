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

type RoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleParameters struct {

	// Indicates that the role containing this value is the default role (that is, has the id 'r_default'), which triggers some specialized behavior to allow it to be imported and managed.
	// +kubebuilder:validation:Optional
	DefaultRole *bool `json:"defaultRole,omitempty" tf:"default_role,omitempty"`

	// The role description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/maxnovawind/provider-jet-boundary/apis/scope/v1alpha1.Scope
	// +kubebuilder:validation:Optional
	GrantScopeID *string `json:"grantScopeId,omitempty" tf:"grant_scope_id,omitempty"`

	// +kubebuilder:validation:Optional
	GrantScopeIDRef *v1.Reference `json:"grantScopeIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GrantScopeIDSelector *v1.Selector `json:"grantScopeIdSelector,omitempty" tf:"-"`

	// A list of stringified grants for the role.
	// +kubebuilder:validation:Optional
	GrantStrings []*string `json:"grantStrings,omitempty" tf:"grant_strings,omitempty"`

	// The role name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of principal (user or group) IDs to add as principals on the role.
	// +kubebuilder:validation:Optional
	PrincipalIds []*string `json:"principalIds,omitempty" tf:"principal_ids,omitempty"`

	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	// +crossplane:generate:reference:type=github.com/maxnovawind/provider-jet-boundary/apis/scope/v1alpha1.Scope
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScopeIDRef *v1.Reference `json:"scopeIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScopeIDSelector *v1.Selector `json:"scopeIdSelector,omitempty" tf:"-"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Role is the Schema for the Roles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundaryjet}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleSpec   `json:"spec"`
	Status            RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
