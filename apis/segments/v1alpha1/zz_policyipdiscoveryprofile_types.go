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




type PolicyIpDiscoveryProfileContextObservation struct {


// Id of the project which the resource belongs to.
ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}


type PolicyIpDiscoveryProfileContextParameters struct {


// Id of the project which the resource belongs to.
// +kubebuilder:validation:Required
ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}


type PolicyIpDiscoveryProfileObservation struct {


// Maximum number of ARP bindings
ArpBindingLimit *float64 `json:"arpBindingLimit,omitempty" tf:"arp_binding_limit,omitempty"`

// ARP and ND cache timeout (in minutes)
ArpNdBindingTimeout *float64 `json:"arpNdBindingTimeout,omitempty" tf:"arp_nd_binding_timeout,omitempty"`

// Is ARP snooping enabled or not
ArpSnoopingEnabled *bool `json:"arpSnoopingEnabled,omitempty" tf:"arp_snooping_enabled,omitempty"`

// Resource context
Context []PolicyIpDiscoveryProfileContextObservation `json:"context,omitempty" tf:"context,omitempty"`

// Is DHCP snooping enabled or not
DHCPSnoopingEnabled *bool `json:"dhcpSnoopingEnabled,omitempty" tf:"dhcp_snooping_enabled,omitempty"`

// Is DHCP snoping v6 enabled or not
DHCPSnoopingV6Enabled *bool `json:"dhcpSnoopingV6Enabled,omitempty" tf:"dhcp_snooping_v6_enabled,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Duplicate IP detection
DuplicateIPDetectionEnabled *bool `json:"duplicateIpDetectionEnabled,omitempty" tf:"duplicate_ip_detection_enabled,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Is ND snooping enabled or not
NdSnoopingEnabled *bool `json:"ndSnoopingEnabled,omitempty" tf:"nd_snooping_enabled,omitempty"`

// Maximum number of ND (Neighbor Discovery Protocol) bindings
NdSnoopingLimit *float64 `json:"ndSnoopingLimit,omitempty" tf:"nd_snooping_limit,omitempty"`

// NSX ID for this resource
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []PolicyIpDiscoveryProfileTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`

// Is TOFU enabled or not
TofuEnabled *bool `json:"tofuEnabled,omitempty" tf:"tofu_enabled,omitempty"`

// Is VM tools enabled or not
VmtoolsEnabled *bool `json:"vmtoolsEnabled,omitempty" tf:"vmtools_enabled,omitempty"`

// Is VM tools enabled or not
VmtoolsV6Enabled *bool `json:"vmtoolsV6Enabled,omitempty" tf:"vmtools_v6_enabled,omitempty"`
}


type PolicyIpDiscoveryProfileParameters struct {


// Maximum number of ARP bindings
// +kubebuilder:validation:Optional
ArpBindingLimit *float64 `json:"arpBindingLimit,omitempty" tf:"arp_binding_limit,omitempty"`

// ARP and ND cache timeout (in minutes)
// +kubebuilder:validation:Optional
ArpNdBindingTimeout *float64 `json:"arpNdBindingTimeout,omitempty" tf:"arp_nd_binding_timeout,omitempty"`

// Is ARP snooping enabled or not
// +kubebuilder:validation:Optional
ArpSnoopingEnabled *bool `json:"arpSnoopingEnabled,omitempty" tf:"arp_snooping_enabled,omitempty"`

// Resource context
// +kubebuilder:validation:Optional
Context []PolicyIpDiscoveryProfileContextParameters `json:"context,omitempty" tf:"context,omitempty"`

// Is DHCP snooping enabled or not
// +kubebuilder:validation:Optional
DHCPSnoopingEnabled *bool `json:"dhcpSnoopingEnabled,omitempty" tf:"dhcp_snooping_enabled,omitempty"`

// Is DHCP snoping v6 enabled or not
// +kubebuilder:validation:Optional
DHCPSnoopingV6Enabled *bool `json:"dhcpSnoopingV6Enabled,omitempty" tf:"dhcp_snooping_v6_enabled,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Duplicate IP detection
// +kubebuilder:validation:Optional
DuplicateIPDetectionEnabled *bool `json:"duplicateIpDetectionEnabled,omitempty" tf:"duplicate_ip_detection_enabled,omitempty"`

// Is ND snooping enabled or not
// +kubebuilder:validation:Optional
NdSnoopingEnabled *bool `json:"ndSnoopingEnabled,omitempty" tf:"nd_snooping_enabled,omitempty"`

// Maximum number of ND (Neighbor Discovery Protocol) bindings
// +kubebuilder:validation:Optional
NdSnoopingLimit *float64 `json:"ndSnoopingLimit,omitempty" tf:"nd_snooping_limit,omitempty"`

// NSX ID for this resource
// +kubebuilder:validation:Optional
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []PolicyIpDiscoveryProfileTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

// Is TOFU enabled or not
// +kubebuilder:validation:Optional
TofuEnabled *bool `json:"tofuEnabled,omitempty" tf:"tofu_enabled,omitempty"`

// Is VM tools enabled or not
// +kubebuilder:validation:Optional
VmtoolsEnabled *bool `json:"vmtoolsEnabled,omitempty" tf:"vmtools_enabled,omitempty"`

// Is VM tools enabled or not
// +kubebuilder:validation:Optional
VmtoolsV6Enabled *bool `json:"vmtoolsV6Enabled,omitempty" tf:"vmtools_v6_enabled,omitempty"`
}


type PolicyIpDiscoveryProfileTagObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyIpDiscoveryProfileTagParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyIpDiscoveryProfileSpec defines the desired state of PolicyIpDiscoveryProfile
type PolicyIpDiscoveryProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyIpDiscoveryProfileParameters `json:"forProvider"`
}

// PolicyIpDiscoveryProfileStatus defines the observed state of PolicyIpDiscoveryProfile.
type PolicyIpDiscoveryProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyIpDiscoveryProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyIpDiscoveryProfile is the Schema for the PolicyIpDiscoveryProfiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyIpDiscoveryProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec              PolicyIpDiscoveryProfileSpec   `json:"spec"`
	Status            PolicyIpDiscoveryProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyIpDiscoveryProfileList contains a list of PolicyIpDiscoveryProfiles
type PolicyIpDiscoveryProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyIpDiscoveryProfile `json:"items"`
}

// Repository type metadata.
var (
	PolicyIpDiscoveryProfile_Kind             = "PolicyIpDiscoveryProfile"
	PolicyIpDiscoveryProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyIpDiscoveryProfile_Kind}.String()
	PolicyIpDiscoveryProfile_KindAPIVersion   = PolicyIpDiscoveryProfile_Kind + "." + CRDGroupVersion.String()
	PolicyIpDiscoveryProfile_GroupVersionKind = CRDGroupVersion.WithKind(PolicyIpDiscoveryProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyIpDiscoveryProfile{}, &PolicyIpDiscoveryProfileList{})
}
