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




type AllocationRangesObservation struct {


End *string `json:"end,omitempty" tf:"end,omitempty"`

Start *string `json:"start,omitempty" tf:"start,omitempty"`
}


type AllocationRangesParameters struct {

}


type IpBlockSubnetObservation struct {


// A collection of IPv4 Pool Ranges
AllocationRanges []AllocationRangesObservation `json:"allocationRanges,omitempty" tf:"allocation_ranges,omitempty"`

// Block id for which the subnet is created
BlockID *string `json:"blockId,omitempty" tf:"block_id,omitempty"`

// Represents network address and the prefix length which will be associated with a layer-2 broadcast domain
Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

// Description of this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The display name of this resource. Defaults to ID if not set
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Represents the size or number of ip addresses in the subnet
Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type IpBlockSubnetParameters struct {


// Block id for which the subnet is created
// +kubebuilder:validation:Optional
BlockID *string `json:"blockId,omitempty" tf:"block_id,omitempty"`

// Description of this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The display name of this resource. Defaults to ID if not set
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Represents the size or number of ip addresses in the subnet
// +kubebuilder:validation:Optional
Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

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

// IpBlockSubnetSpec defines the desired state of IpBlockSubnet
type IpBlockSubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       IpBlockSubnetParameters `json:"forProvider"`
}

// IpBlockSubnetStatus defines the observed state of IpBlockSubnet.
type IpBlockSubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          IpBlockSubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpBlockSubnet is the Schema for the IpBlockSubnets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type IpBlockSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.blockId)",message="blockId is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.size)",message="size is a required parameter"
	Spec              IpBlockSubnetSpec   `json:"spec"`
	Status            IpBlockSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpBlockSubnetList contains a list of IpBlockSubnets
type IpBlockSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpBlockSubnet `json:"items"`
}

// Repository type metadata.
var (
	IpBlockSubnet_Kind             = "IpBlockSubnet"
	IpBlockSubnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpBlockSubnet_Kind}.String()
	IpBlockSubnet_KindAPIVersion   = IpBlockSubnet_Kind + "." + CRDGroupVersion.String()
	IpBlockSubnet_GroupVersionKind = CRDGroupVersion.WithKind(IpBlockSubnet_Kind)
)

func init() {
	SchemeBuilder.Register(&IpBlockSubnet{}, &IpBlockSubnetList{})
}
