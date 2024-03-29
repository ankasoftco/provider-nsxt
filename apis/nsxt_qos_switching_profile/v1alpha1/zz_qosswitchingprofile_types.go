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




type EgressRateShaperObservation struct {


// Average Bandwidth in mbps
AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

// Burst size in bytes
BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

// Whether this rate shaper is enabled
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Peak Bandwidth in mbps
PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}


type EgressRateShaperParameters struct {


// Average Bandwidth in mbps
// +kubebuilder:validation:Optional
AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

// Burst size in bytes
// +kubebuilder:validation:Optional
BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

// Whether this rate shaper is enabled
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Peak Bandwidth in mbps
// +kubebuilder:validation:Optional
PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}


type IngressBroadcastRateShaperObservation struct {


// Average Bandwidth in kbps
AverageBwKbps *float64 `json:"averageBwKbps,omitempty" tf:"average_bw_kbps,omitempty"`

// Burst size in bytes
BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

// Whether this rate shaper is enabled
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Peak Bandwidth in kbps
PeakBwKbps *float64 `json:"peakBwKbps,omitempty" tf:"peak_bw_kbps,omitempty"`
}


type IngressBroadcastRateShaperParameters struct {


// Average Bandwidth in kbps
// +kubebuilder:validation:Optional
AverageBwKbps *float64 `json:"averageBwKbps,omitempty" tf:"average_bw_kbps,omitempty"`

// Burst size in bytes
// +kubebuilder:validation:Optional
BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

// Whether this rate shaper is enabled
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Peak Bandwidth in kbps
// +kubebuilder:validation:Optional
PeakBwKbps *float64 `json:"peakBwKbps,omitempty" tf:"peak_bw_kbps,omitempty"`
}


type IngressRateShaperObservation struct {


// Average Bandwidth in mbps
AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

// Burst size in bytes
BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

// Whether this rate shaper is enabled
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Peak Bandwidth in mbps
PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}


type IngressRateShaperParameters struct {


// Average Bandwidth in mbps
// +kubebuilder:validation:Optional
AverageBwMbps *float64 `json:"averageBwMbps,omitempty" tf:"average_bw_mbps,omitempty"`

// Burst size in bytes
// +kubebuilder:validation:Optional
BurstSize *float64 `json:"burstSize,omitempty" tf:"burst_size,omitempty"`

// Whether this rate shaper is enabled
// +kubebuilder:validation:Optional
Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

// Peak Bandwidth in mbps
// +kubebuilder:validation:Optional
PeakBwMbps *float64 `json:"peakBwMbps,omitempty" tf:"peak_bw_mbps,omitempty"`
}


type QosSwitchingProfileObservation struct {


// Class of service
ClassOfService *float64 `json:"classOfService,omitempty" tf:"class_of_service,omitempty"`

// Description of this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The display name of this resource. Defaults to ID if not set
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// DSCP Priority
DscpPriority *float64 `json:"dscpPriority,omitempty" tf:"dscp_priority,omitempty"`

// Trust mode for DSCP
DscpTrusted *bool `json:"dscpTrusted,omitempty" tf:"dscp_trusted,omitempty"`

EgressRateShaper []EgressRateShaperObservation `json:"egressRateShaper,omitempty" tf:"egress_rate_shaper,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

IngressBroadcastRateShaper []IngressBroadcastRateShaperObservation `json:"ingressBroadcastRateShaper,omitempty" tf:"ingress_broadcast_rate_shaper,omitempty"`

IngressRateShaper []IngressRateShaperObservation `json:"ingressRateShaper,omitempty" tf:"ingress_rate_shaper,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type QosSwitchingProfileParameters struct {


// Class of service
// +kubebuilder:validation:Optional
ClassOfService *float64 `json:"classOfService,omitempty" tf:"class_of_service,omitempty"`

// Description of this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The display name of this resource. Defaults to ID if not set
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// DSCP Priority
// +kubebuilder:validation:Optional
DscpPriority *float64 `json:"dscpPriority,omitempty" tf:"dscp_priority,omitempty"`

// Trust mode for DSCP
// +kubebuilder:validation:Optional
DscpTrusted *bool `json:"dscpTrusted,omitempty" tf:"dscp_trusted,omitempty"`

// +kubebuilder:validation:Optional
EgressRateShaper []EgressRateShaperParameters `json:"egressRateShaper,omitempty" tf:"egress_rate_shaper,omitempty"`

// +kubebuilder:validation:Optional
IngressBroadcastRateShaper []IngressBroadcastRateShaperParameters `json:"ingressBroadcastRateShaper,omitempty" tf:"ingress_broadcast_rate_shaper,omitempty"`

// +kubebuilder:validation:Optional
IngressRateShaper []IngressRateShaperParameters `json:"ingressRateShaper,omitempty" tf:"ingress_rate_shaper,omitempty"`

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

// QosSwitchingProfileSpec defines the desired state of QosSwitchingProfile
type QosSwitchingProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       QosSwitchingProfileParameters `json:"forProvider"`
}

// QosSwitchingProfileStatus defines the observed state of QosSwitchingProfile.
type QosSwitchingProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          QosSwitchingProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QosSwitchingProfile is the Schema for the QosSwitchingProfiles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type QosSwitchingProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QosSwitchingProfileSpec   `json:"spec"`
	Status            QosSwitchingProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QosSwitchingProfileList contains a list of QosSwitchingProfiles
type QosSwitchingProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QosSwitchingProfile `json:"items"`
}

// Repository type metadata.
var (
	QosSwitchingProfile_Kind             = "QosSwitchingProfile"
	QosSwitchingProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QosSwitchingProfile_Kind}.String()
	QosSwitchingProfile_KindAPIVersion   = QosSwitchingProfile_Kind + "." + CRDGroupVersion.String()
	QosSwitchingProfile_GroupVersionKind = CRDGroupVersion.WithKind(QosSwitchingProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&QosSwitchingProfile{}, &QosSwitchingProfileList{})
}
