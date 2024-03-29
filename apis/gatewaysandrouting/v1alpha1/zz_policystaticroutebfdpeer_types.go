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




type PolicyStaticRouteBfdPeerObservation struct {


// Policy path for BFD Profile
BfdProfilePath *string `json:"bfdProfilePath,omitempty" tf:"bfd_profile_path,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Flag to enable/disable this peer
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Policy path for Tier0 gateway
GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// NSX ID for this resource
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// IPv4 Address of the peer
PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Array of Tier0 external interface IP addresses
SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []PolicyStaticRouteBfdPeerTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyStaticRouteBfdPeerParameters struct {


// Policy path for BFD Profile
// +kubebuilder:validation:Optional
BfdProfilePath *string `json:"bfdProfilePath,omitempty" tf:"bfd_profile_path,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Flag to enable/disable this peer
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Policy path for Tier0 gateway
// +kubebuilder:validation:Optional
GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

// NSX ID for this resource
// +kubebuilder:validation:Optional
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// IPv4 Address of the peer
// +kubebuilder:validation:Optional
PeerAddress *string `json:"peerAddress,omitempty" tf:"peer_address,omitempty"`

// Array of Tier0 external interface IP addresses
// +kubebuilder:validation:Optional
SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []PolicyStaticRouteBfdPeerTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyStaticRouteBfdPeerTagObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyStaticRouteBfdPeerTagParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyStaticRouteBfdPeerSpec defines the desired state of PolicyStaticRouteBfdPeer
type PolicyStaticRouteBfdPeerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyStaticRouteBfdPeerParameters `json:"forProvider"`
}

// PolicyStaticRouteBfdPeerStatus defines the observed state of PolicyStaticRouteBfdPeer.
type PolicyStaticRouteBfdPeerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyStaticRouteBfdPeerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyStaticRouteBfdPeer is the Schema for the PolicyStaticRouteBfdPeers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyStaticRouteBfdPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bfdProfilePath)",message="bfdProfilePath is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.gatewayPath)",message="gatewayPath is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.peerAddress)",message="peerAddress is a required parameter"
	Spec              PolicyStaticRouteBfdPeerSpec   `json:"spec"`
	Status            PolicyStaticRouteBfdPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyStaticRouteBfdPeerList contains a list of PolicyStaticRouteBfdPeers
type PolicyStaticRouteBfdPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyStaticRouteBfdPeer `json:"items"`
}

// Repository type metadata.
var (
	PolicyStaticRouteBfdPeer_Kind             = "PolicyStaticRouteBfdPeer"
	PolicyStaticRouteBfdPeer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyStaticRouteBfdPeer_Kind}.String()
	PolicyStaticRouteBfdPeer_KindAPIVersion   = PolicyStaticRouteBfdPeer_Kind + "." + CRDGroupVersion.String()
	PolicyStaticRouteBfdPeer_GroupVersionKind = CRDGroupVersion.WithKind(PolicyStaticRouteBfdPeer_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyStaticRouteBfdPeer{}, &PolicyStaticRouteBfdPeerList{})
}
