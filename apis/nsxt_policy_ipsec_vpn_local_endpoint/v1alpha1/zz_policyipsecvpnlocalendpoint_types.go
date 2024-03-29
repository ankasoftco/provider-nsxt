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




type PolicyIpsecVpnLocalEndpointObservation struct {


// Policy path referencing site certificate
CertificatePath *string `json:"certificatePath,omitempty" tf:"certificate_path,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

LocalAddress *string `json:"localAddress,omitempty" tf:"local_address,omitempty"`

LocalID *string `json:"localId,omitempty" tf:"local_id,omitempty"`

// NSX ID for this resource
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Policy path for IPSec VPN service
ServicePath *string `json:"servicePath,omitempty" tf:"service_path,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`

TrustCAPaths []*string `json:"trustCaPaths,omitempty" tf:"trust_ca_paths,omitempty"`

TrustCrlPaths []*string `json:"trustCrlPaths,omitempty" tf:"trust_crl_paths,omitempty"`
}


type PolicyIpsecVpnLocalEndpointParameters struct {


// Policy path referencing site certificate
// +kubebuilder:validation:Optional
CertificatePath *string `json:"certificatePath,omitempty" tf:"certificate_path,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// +kubebuilder:validation:Optional
LocalAddress *string `json:"localAddress,omitempty" tf:"local_address,omitempty"`

// +kubebuilder:validation:Optional
LocalID *string `json:"localId,omitempty" tf:"local_id,omitempty"`

// NSX ID for this resource
// +kubebuilder:validation:Optional
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for IPSec VPN service
// +kubebuilder:validation:Optional
ServicePath *string `json:"servicePath,omitempty" tf:"service_path,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

// +kubebuilder:validation:Optional
TrustCAPaths []*string `json:"trustCaPaths,omitempty" tf:"trust_ca_paths,omitempty"`

// +kubebuilder:validation:Optional
TrustCrlPaths []*string `json:"trustCrlPaths,omitempty" tf:"trust_crl_paths,omitempty"`
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

// PolicyIpsecVpnLocalEndpointSpec defines the desired state of PolicyIpsecVpnLocalEndpoint
type PolicyIpsecVpnLocalEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyIpsecVpnLocalEndpointParameters `json:"forProvider"`
}

// PolicyIpsecVpnLocalEndpointStatus defines the observed state of PolicyIpsecVpnLocalEndpoint.
type PolicyIpsecVpnLocalEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyIpsecVpnLocalEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyIpsecVpnLocalEndpoint is the Schema for the PolicyIpsecVpnLocalEndpoints API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyIpsecVpnLocalEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.localAddress)",message="localAddress is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.servicePath)",message="servicePath is a required parameter"
	Spec              PolicyIpsecVpnLocalEndpointSpec   `json:"spec"`
	Status            PolicyIpsecVpnLocalEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyIpsecVpnLocalEndpointList contains a list of PolicyIpsecVpnLocalEndpoints
type PolicyIpsecVpnLocalEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyIpsecVpnLocalEndpoint `json:"items"`
}

// Repository type metadata.
var (
	PolicyIpsecVpnLocalEndpoint_Kind             = "PolicyIpsecVpnLocalEndpoint"
	PolicyIpsecVpnLocalEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyIpsecVpnLocalEndpoint_Kind}.String()
	PolicyIpsecVpnLocalEndpoint_KindAPIVersion   = PolicyIpsecVpnLocalEndpoint_Kind + "." + CRDGroupVersion.String()
	PolicyIpsecVpnLocalEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(PolicyIpsecVpnLocalEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyIpsecVpnLocalEndpoint{}, &PolicyIpsecVpnLocalEndpointList{})
}
