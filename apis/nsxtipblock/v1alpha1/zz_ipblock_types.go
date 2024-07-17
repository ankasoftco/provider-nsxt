/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IpBlockObservation struct {

	// Represents network address and the prefix length which will be associated with a layer-2 broadcast domain
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of this resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of this resource. Defaults to ID if not set
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type IpBlockParameters struct {

	// Represents network address and the prefix length which will be associated with a layer-2 broadcast domain
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of this resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of this resource. Defaults to ID if not set
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type TagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type TagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// IpBlockSpec defines the desired state of IpBlock
type IpBlockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpBlockParameters `json:"forProvider"`
}

// IpBlockStatus defines the observed state of IpBlock.
type IpBlockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpBlockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpBlock is the Schema for the IpBlocks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type IpBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.cidr)",message="cidr is a required parameter"
	Spec   IpBlockSpec   `json:"spec"`
	Status IpBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpBlockList contains a list of IpBlocks
type IpBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpBlock `json:"items"`
}

// Repository type metadata.
var (
	IpBlock_Kind             = "IpBlock"
	IpBlock_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpBlock_Kind}.String()
	IpBlock_KindAPIVersion   = IpBlock_Kind + "." + CRDGroupVersion.String()
	IpBlock_GroupVersionKind = CRDGroupVersion.WithKind(IpBlock_Kind)
)

func init() {
	SchemeBuilder.Register(&IpBlock{}, &IpBlockList{})
}