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




type PolicyGatewayPrefixListObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Policy path for Tier0 gateway
GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// NSX ID for this resource
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// Ordered list of network prefixes
Prefix []PrefixObservation `json:"prefix,omitempty" tf:"prefix,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []PolicyGatewayPrefixListTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyGatewayPrefixListParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Policy path for Tier0 gateway
// +kubebuilder:validation:Optional
GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

// NSX ID for this resource
// +kubebuilder:validation:Optional
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Ordered list of network prefixes
// +kubebuilder:validation:Optional
Prefix []PrefixParameters `json:"prefix,omitempty" tf:"prefix,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []PolicyGatewayPrefixListTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyGatewayPrefixListTagObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyGatewayPrefixListTagParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PrefixObservation struct {


// Action for the prefix list
Action *string `json:"action,omitempty" tf:"action,omitempty"`

// Prefix length greater than or equal to
Ge *float64 `json:"ge,omitempty" tf:"ge,omitempty"`

// Prefix length less than or equal to
Le *float64 `json:"le,omitempty" tf:"le,omitempty"`

// Network prefix in CIDR format. If not set it will match ANY network
Network *string `json:"network,omitempty" tf:"network,omitempty"`
}


type PrefixParameters struct {


// Action for the prefix list
// +kubebuilder:validation:Optional
Action *string `json:"action,omitempty" tf:"action,omitempty"`

// Prefix length greater than or equal to
// +kubebuilder:validation:Optional
Ge *float64 `json:"ge,omitempty" tf:"ge,omitempty"`

// Prefix length less than or equal to
// +kubebuilder:validation:Optional
Le *float64 `json:"le,omitempty" tf:"le,omitempty"`

// Network prefix in CIDR format. If not set it will match ANY network
// +kubebuilder:validation:Optional
Network *string `json:"network,omitempty" tf:"network,omitempty"`
}

// PolicyGatewayPrefixListSpec defines the desired state of PolicyGatewayPrefixList
type PolicyGatewayPrefixListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyGatewayPrefixListParameters `json:"forProvider"`
}

// PolicyGatewayPrefixListStatus defines the observed state of PolicyGatewayPrefixList.
type PolicyGatewayPrefixListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyGatewayPrefixListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGatewayPrefixList is the Schema for the PolicyGatewayPrefixLists API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyGatewayPrefixList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.gatewayPath)",message="gatewayPath is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.prefix)",message="prefix is a required parameter"
	Spec              PolicyGatewayPrefixListSpec   `json:"spec"`
	Status            PolicyGatewayPrefixListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGatewayPrefixListList contains a list of PolicyGatewayPrefixLists
type PolicyGatewayPrefixListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyGatewayPrefixList `json:"items"`
}

// Repository type metadata.
var (
	PolicyGatewayPrefixList_Kind             = "PolicyGatewayPrefixList"
	PolicyGatewayPrefixList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyGatewayPrefixList_Kind}.String()
	PolicyGatewayPrefixList_KindAPIVersion   = PolicyGatewayPrefixList_Kind + "." + CRDGroupVersion.String()
	PolicyGatewayPrefixList_GroupVersionKind = CRDGroupVersion.WithKind(PolicyGatewayPrefixList_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyGatewayPrefixList{}, &PolicyGatewayPrefixListList{})
}
