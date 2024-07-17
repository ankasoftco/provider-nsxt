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

type IgmpTypeNsServiceObservation struct {

	// A boolean flag which reflects whether this is a default NSServices which can't be modified/deleted
	DefaultService *bool `json:"defaultService,omitempty" tf:"default_service,omitempty"`

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

type IgmpTypeNsServiceParameters struct {

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

// IgmpTypeNsServiceSpec defines the desired state of IgmpTypeNsService
type IgmpTypeNsServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IgmpTypeNsServiceParameters `json:"forProvider"`
}

// IgmpTypeNsServiceStatus defines the observed state of IgmpTypeNsService.
type IgmpTypeNsServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IgmpTypeNsServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IgmpTypeNsService is the Schema for the IgmpTypeNsServices API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type IgmpTypeNsService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IgmpTypeNsServiceSpec   `json:"spec"`
	Status            IgmpTypeNsServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IgmpTypeNsServiceList contains a list of IgmpTypeNsServices
type IgmpTypeNsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IgmpTypeNsService `json:"items"`
}

// Repository type metadata.
var (
	IgmpTypeNsService_Kind             = "IgmpTypeNsService"
	IgmpTypeNsService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IgmpTypeNsService_Kind}.String()
	IgmpTypeNsService_KindAPIVersion   = IgmpTypeNsService_Kind + "." + CRDGroupVersion.String()
	IgmpTypeNsService_GroupVersionKind = CRDGroupVersion.WithKind(IgmpTypeNsService_Kind)
)

func init() {
	SchemeBuilder.Register(&IgmpTypeNsService{}, &IgmpTypeNsServiceList{})
}