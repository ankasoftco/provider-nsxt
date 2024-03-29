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




type PolicyOspfConfigObservation struct {


// Flag to enable/disable advertisement of default route into OSPF domain
DefaultOriginate *bool `json:"defaultOriginate,omitempty" tf:"default_originate,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Flag to enable ECMP
Ecmp *bool `json:"ecmp,omitempty" tf:"ecmp,omitempty"`

// Flag to enable OSPF configuration
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// NSX ID of associated Tier0 Gateway
GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

// Policy path for the Tier0 Gateway
GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

// Graceful Restart Mode
GracefulRestartMode *string `json:"gracefulRestartMode,omitempty" tf:"graceful_restart_mode,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// NSX ID of associated Gateway Locale Service
LocaleServiceID *string `json:"localeServiceId,omitempty" tf:"locale_service_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// List of addresses to summarize or filter external routes
SummaryAddress []SummaryAddressObservation `json:"summaryAddress,omitempty" tf:"summary_address,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []PolicyOspfConfigTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyOspfConfigParameters struct {


// Flag to enable/disable advertisement of default route into OSPF domain
// +kubebuilder:validation:Optional
DefaultOriginate *bool `json:"defaultOriginate,omitempty" tf:"default_originate,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Flag to enable ECMP
// +kubebuilder:validation:Optional
Ecmp *bool `json:"ecmp,omitempty" tf:"ecmp,omitempty"`

// Flag to enable OSPF configuration
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Policy path for the Tier0 Gateway
// +kubebuilder:validation:Optional
GatewayPath *string `json:"gatewayPath,omitempty" tf:"gateway_path,omitempty"`

// Graceful Restart Mode
// +kubebuilder:validation:Optional
GracefulRestartMode *string `json:"gracefulRestartMode,omitempty" tf:"graceful_restart_mode,omitempty"`

// List of addresses to summarize or filter external routes
// +kubebuilder:validation:Optional
SummaryAddress []SummaryAddressParameters `json:"summaryAddress,omitempty" tf:"summary_address,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []PolicyOspfConfigTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyOspfConfigTagObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyOspfConfigTagParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type SummaryAddressObservation struct {


// Used to filter the advertisement of external routes into the OSPF domain
Advertise *bool `json:"advertise,omitempty" tf:"advertise,omitempty"`

// OSPF Summary address in CIDR format
Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}


type SummaryAddressParameters struct {


// Used to filter the advertisement of external routes into the OSPF domain
// +kubebuilder:validation:Optional
Advertise *bool `json:"advertise,omitempty" tf:"advertise,omitempty"`

// OSPF Summary address in CIDR format
// +kubebuilder:validation:Optional
Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

// PolicyOspfConfigSpec defines the desired state of PolicyOspfConfig
type PolicyOspfConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyOspfConfigParameters `json:"forProvider"`
}

// PolicyOspfConfigStatus defines the observed state of PolicyOspfConfig.
type PolicyOspfConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyOspfConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyOspfConfig is the Schema for the PolicyOspfConfigs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyOspfConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.gatewayPath)",message="gatewayPath is a required parameter"
	Spec              PolicyOspfConfigSpec   `json:"spec"`
	Status            PolicyOspfConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyOspfConfigList contains a list of PolicyOspfConfigs
type PolicyOspfConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyOspfConfig `json:"items"`
}

// Repository type metadata.
var (
	PolicyOspfConfig_Kind             = "PolicyOspfConfig"
	PolicyOspfConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyOspfConfig_Kind}.String()
	PolicyOspfConfig_KindAPIVersion   = PolicyOspfConfig_Kind + "." + CRDGroupVersion.String()
	PolicyOspfConfig_GroupVersionKind = CRDGroupVersion.WithKind(PolicyOspfConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyOspfConfig{}, &PolicyOspfConfigList{})
}
