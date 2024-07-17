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

type PolicyIpPoolBlockSubnetContextObservation struct {

	// Id of the project which the resource belongs to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type PolicyIpPoolBlockSubnetContextParameters struct {

	// Id of the project which the resource belongs to.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

type PolicyIpPoolBlockSubnetObservation struct {

	// If true, the first IP in the range will be reserved for gateway
	AutoAssignGateway *bool `json:"autoAssignGateway,omitempty" tf:"auto_assign_gateway,omitempty"`

	// Policy path to the IP Block
	BlockPath *string `json:"blockPath,omitempty" tf:"block_path,omitempty"`

	// Resource context
	Context []PolicyIpPoolBlockSubnetContextObservation `json:"context,omitempty" tf:"context,omitempty"`

	// Description for this resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for this resource
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// NSX ID for this resource
	NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

	// Policy path for this resource
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Policy path to the IP Pool for this Subnet
	PoolPath *string `json:"poolPath,omitempty" tf:"pool_path,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Number of addresses
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Set of opaque identifiers meaningful to the user
	Tag []PolicyIpPoolBlockSubnetTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyIpPoolBlockSubnetParameters struct {

	// If true, the first IP in the range will be reserved for gateway
	// +kubebuilder:validation:Optional
	AutoAssignGateway *bool `json:"autoAssignGateway,omitempty" tf:"auto_assign_gateway,omitempty"`

	// Policy path to the IP Block
	// +kubebuilder:validation:Optional
	BlockPath *string `json:"blockPath,omitempty" tf:"block_path,omitempty"`

	// Resource context
	// +kubebuilder:validation:Optional
	Context []PolicyIpPoolBlockSubnetContextParameters `json:"context,omitempty" tf:"context,omitempty"`

	// Description for this resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name for this resource
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// NSX ID for this resource
	// +kubebuilder:validation:Optional
	NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

	// Policy path to the IP Pool for this Subnet
	// +kubebuilder:validation:Optional
	PoolPath *string `json:"poolPath,omitempty" tf:"pool_path,omitempty"`

	// Number of addresses
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Set of opaque identifiers meaningful to the user
	// +kubebuilder:validation:Optional
	Tag []PolicyIpPoolBlockSubnetTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyIpPoolBlockSubnetTagObservation struct {
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type PolicyIpPoolBlockSubnetTagParameters struct {

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyIpPoolBlockSubnetSpec defines the desired state of PolicyIpPoolBlockSubnet
type PolicyIpPoolBlockSubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyIpPoolBlockSubnetParameters `json:"forProvider"`
}

// PolicyIpPoolBlockSubnetStatus defines the observed state of PolicyIpPoolBlockSubnet.
type PolicyIpPoolBlockSubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyIpPoolBlockSubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyIpPoolBlockSubnet is the Schema for the PolicyIpPoolBlockSubnets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyIpPoolBlockSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.blockPath)",message="blockPath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.poolPath)",message="poolPath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.size)",message="size is a required parameter"
	Spec   PolicyIpPoolBlockSubnetSpec   `json:"spec"`
	Status PolicyIpPoolBlockSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyIpPoolBlockSubnetList contains a list of PolicyIpPoolBlockSubnets
type PolicyIpPoolBlockSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyIpPoolBlockSubnet `json:"items"`
}

// Repository type metadata.
var (
	PolicyIpPoolBlockSubnet_Kind             = "PolicyIpPoolBlockSubnet"
	PolicyIpPoolBlockSubnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyIpPoolBlockSubnet_Kind}.String()
	PolicyIpPoolBlockSubnet_KindAPIVersion   = PolicyIpPoolBlockSubnet_Kind + "." + CRDGroupVersion.String()
	PolicyIpPoolBlockSubnet_GroupVersionKind = CRDGroupVersion.WithKind(PolicyIpPoolBlockSubnet_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyIpPoolBlockSubnet{}, &PolicyIpPoolBlockSubnetList{})
}