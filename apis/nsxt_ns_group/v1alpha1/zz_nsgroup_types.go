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




type MemberObservation struct {


// Type of the resource on which this expression is evaluated
TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

// Value that satisfies this expression
Value *string `json:"value,omitempty" tf:"value,omitempty"`
}


type MemberParameters struct {


// Type of the resource on which this expression is evaluated
// +kubebuilder:validation:Required
TargetType *string `json:"targetType" tf:"target_type,omitempty"`

// Value that satisfies this expression
// +kubebuilder:validation:Required
Value *string `json:"value" tf:"value,omitempty"`
}


type MembershipCriteriaObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

ScopeOp *string `json:"scopeOp,omitempty" tf:"scope_op,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

TagOp *string `json:"tagOp,omitempty" tf:"tag_op,omitempty"`

TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`
}


type MembershipCriteriaParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
ScopeOp *string `json:"scopeOp,omitempty" tf:"scope_op,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

// +kubebuilder:validation:Optional
TagOp *string `json:"tagOp,omitempty" tf:"tag_op,omitempty"`

// +kubebuilder:validation:Required
TargetType *string `json:"targetType" tf:"target_type,omitempty"`
}


type NsGroupObservation struct {


// Description of this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The display name of this resource. Defaults to ID if not set
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Reference to the direct/static members of the NSGroup.
Member []MemberObservation `json:"member,omitempty" tf:"member,omitempty"`

// List of tag expressions which define the membership criteria for this NSGroup.
MembershipCriteria []MembershipCriteriaObservation `json:"membershipCriteria,omitempty" tf:"membership_criteria,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []TagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type NsGroupParameters struct {


// Description of this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The display name of this resource. Defaults to ID if not set
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Reference to the direct/static members of the NSGroup.
// +kubebuilder:validation:Optional
Member []MemberParameters `json:"member,omitempty" tf:"member,omitempty"`

// List of tag expressions which define the membership criteria for this NSGroup.
// +kubebuilder:validation:Optional
MembershipCriteria []MembershipCriteriaParameters `json:"membershipCriteria,omitempty" tf:"membership_criteria,omitempty"`

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

// NsGroupSpec defines the desired state of NsGroup
type NsGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       NsGroupParameters `json:"forProvider"`
}

// NsGroupStatus defines the observed state of NsGroup.
type NsGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          NsGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsGroup is the Schema for the NsGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type NsGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NsGroupSpec   `json:"spec"`
	Status            NsGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsGroupList contains a list of NsGroups
type NsGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsGroup `json:"items"`
}

// Repository type metadata.
var (
	NsGroup_Kind             = "NsGroup"
	NsGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsGroup_Kind}.String()
	NsGroup_KindAPIVersion   = NsGroup_Kind + "." + CRDGroupVersion.String()
	NsGroup_GroupVersionKind = CRDGroupVersion.WithKind(NsGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NsGroup{}, &NsGroupList{})
}
