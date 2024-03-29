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




type ConditionObservation struct {


// The resource key attribute to apply the condition to.
Key *string `json:"key,omitempty" tf:"key,omitempty"`

// The NSX member to apply the condition to. Can be one of; IPSet, LogicalPort, LogicalSwitch, Segment, SegmentPort or VirtualMachine
MemberType *string `json:"memberType,omitempty" tf:"member_type,omitempty"`

// The operator to use for the condition. Can be one of; CONTAINS, ENDSWITH, EQUALS, NOTEQUALS or STARTSWITH
Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

// The value to check for in the condition
Value *string `json:"value,omitempty" tf:"value,omitempty"`
}


type ConditionParameters struct {


// The resource key attribute to apply the condition to.
// +kubebuilder:validation:Required
Key *string `json:"key" tf:"key,omitempty"`

// The NSX member to apply the condition to. Can be one of; IPSet, LogicalPort, LogicalSwitch, Segment, SegmentPort or VirtualMachine
// +kubebuilder:validation:Required
MemberType *string `json:"memberType" tf:"member_type,omitempty"`

// The operator to use for the condition. Can be one of; CONTAINS, ENDSWITH, EQUALS, NOTEQUALS or STARTSWITH
// +kubebuilder:validation:Required
Operator *string `json:"operator" tf:"operator,omitempty"`

// The value to check for in the condition
// +kubebuilder:validation:Required
Value *string `json:"value" tf:"value,omitempty"`
}


type ConjunctionObservation struct {


// The conjunction operator; either OR or AND
Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}


type ConjunctionParameters struct {


// The conjunction operator; either OR or AND
// +kubebuilder:validation:Required
Operator *string `json:"operator" tf:"operator,omitempty"`
}


type ContextObservation struct {


// Id of the project which the resource belongs to.
ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}


type ContextParameters struct {


// Id of the project which the resource belongs to.
// +kubebuilder:validation:Required
ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}


type CriteriaObservation struct {


// A Condition querying resources for membership in the Group
Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

// External ID expression specifying additional members in the Group
ExternalIDExpression []ExternalIDExpressionObservation `json:"externalIdExpression,omitempty" tf:"external_id_expression,omitempty"`

// An IP Address expression specifying IP Address members in the Group
IpaddressExpression []IpaddressExpressionObservation `json:"ipaddressExpression,omitempty" tf:"ipaddress_expression,omitempty"`

// MAC address expression specifying MAC Address members in the Group
MacaddressExpression []MacaddressExpressionObservation `json:"macaddressExpression,omitempty" tf:"macaddress_expression,omitempty"`

// A list of object paths for members in the group
PathExpression []PathExpressionObservation `json:"pathExpression,omitempty" tf:"path_expression,omitempty"`
}


type CriteriaParameters struct {


// A Condition querying resources for membership in the Group
// +kubebuilder:validation:Optional
Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

// External ID expression specifying additional members in the Group
// +kubebuilder:validation:Optional
ExternalIDExpression []ExternalIDExpressionParameters `json:"externalIdExpression,omitempty" tf:"external_id_expression,omitempty"`

// An IP Address expression specifying IP Address members in the Group
// +kubebuilder:validation:Optional
IpaddressExpression []IpaddressExpressionParameters `json:"ipaddressExpression,omitempty" tf:"ipaddress_expression,omitempty"`

// MAC address expression specifying MAC Address members in the Group
// +kubebuilder:validation:Optional
MacaddressExpression []MacaddressExpressionParameters `json:"macaddressExpression,omitempty" tf:"macaddress_expression,omitempty"`

// A list of object paths for members in the group
// +kubebuilder:validation:Optional
PathExpression []PathExpressionParameters `json:"pathExpression,omitempty" tf:"path_expression,omitempty"`
}


type ExtendedCriteriaObservation struct {


// Identity Group expression
IdentityGroup []IdentityGroupObservation `json:"identityGroup,omitempty" tf:"identity_group,omitempty"`
}


type ExtendedCriteriaParameters struct {


// Identity Group expression
// +kubebuilder:validation:Optional
IdentityGroup []IdentityGroupParameters `json:"identityGroup,omitempty" tf:"identity_group,omitempty"`
}


type ExternalIDExpressionObservation struct {


// List of external IDs
ExternalIds []*string `json:"externalIds,omitempty" tf:"external_ids,omitempty"`

// External ID member type, default to virtual machine if not specified
MemberType *string `json:"memberType,omitempty" tf:"member_type,omitempty"`
}


type ExternalIDExpressionParameters struct {


// List of external IDs
// +kubebuilder:validation:Required
ExternalIds []*string `json:"externalIds" tf:"external_ids,omitempty"`

// External ID member type, default to virtual machine if not specified
// +kubebuilder:validation:Optional
MemberType *string `json:"memberType,omitempty" tf:"member_type,omitempty"`
}


type IdentityGroupObservation struct {


// LDAP distinguished name
DistinguishedName *string `json:"distinguishedName,omitempty" tf:"distinguished_name,omitempty"`

// Identity (Directory) domain base distinguished name
DomainBaseDistinguishedName *string `json:"domainBaseDistinguishedName,omitempty" tf:"domain_base_distinguished_name,omitempty"`

// Identity (Directory) Group SID (security identifier)
Sid *string `json:"sid,omitempty" tf:"sid,omitempty"`
}


type IdentityGroupParameters struct {


// LDAP distinguished name
// +kubebuilder:validation:Optional
DistinguishedName *string `json:"distinguishedName,omitempty" tf:"distinguished_name,omitempty"`

// Identity (Directory) domain base distinguished name
// +kubebuilder:validation:Optional
DomainBaseDistinguishedName *string `json:"domainBaseDistinguishedName,omitempty" tf:"domain_base_distinguished_name,omitempty"`

// Identity (Directory) Group SID (security identifier)
// +kubebuilder:validation:Optional
Sid *string `json:"sid,omitempty" tf:"sid,omitempty"`
}


type IpaddressExpressionObservation struct {


// List of single IP addresses, IP address ranges or Subnets. Cannot mix IPv4 and IPv6 in a single list
IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
}


type IpaddressExpressionParameters struct {


// List of single IP addresses, IP address ranges or Subnets. Cannot mix IPv4 and IPv6 in a single list
// +kubebuilder:validation:Required
IPAddresses []*string `json:"ipAddresses" tf:"ip_addresses,omitempty"`
}


type MacaddressExpressionObservation struct {


// List of Mac Addresses
MacAddresses []*string `json:"macAddresses,omitempty" tf:"mac_addresses,omitempty"`
}


type MacaddressExpressionParameters struct {


// List of Mac Addresses
// +kubebuilder:validation:Required
MacAddresses []*string `json:"macAddresses" tf:"mac_addresses,omitempty"`
}


type PathExpressionObservation struct {


// List of policy paths of direct group members
MemberPaths []*string `json:"memberPaths,omitempty" tf:"member_paths,omitempty"`
}


type PathExpressionParameters struct {


// List of policy paths of direct group members
// +kubebuilder:validation:Required
MemberPaths []*string `json:"memberPaths" tf:"member_paths,omitempty"`
}


type PolicyGroupObservation struct {


// A conjunction applied to 2 sets of criteria.
Conjunction []ConjunctionObservation `json:"conjunction,omitempty" tf:"conjunction,omitempty"`

// Resource context
Context []ContextObservation `json:"context,omitempty" tf:"context,omitempty"`

// Criteria to determine Group membership
Criteria []CriteriaObservation `json:"criteria,omitempty" tf:"criteria,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// The domain name to use for resources. If not specified 'default' is used
Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

// Extended criteria to determine group membership. extended_criteria is implicitly "AND" with criteria
ExtendedCriteria []ExtendedCriteriaObservation `json:"extendedCriteria,omitempty" tf:"extended_criteria,omitempty"`

// Indicates the group type
GroupType *string `json:"groupType,omitempty" tf:"group_type,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// NSX ID for this resource
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []PolicyGroupTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyGroupParameters struct {


// A conjunction applied to 2 sets of criteria.
// +kubebuilder:validation:Optional
Conjunction []ConjunctionParameters `json:"conjunction,omitempty" tf:"conjunction,omitempty"`

// Resource context
// +kubebuilder:validation:Optional
Context []ContextParameters `json:"context,omitempty" tf:"context,omitempty"`

// Criteria to determine Group membership
// +kubebuilder:validation:Optional
Criteria []CriteriaParameters `json:"criteria,omitempty" tf:"criteria,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// The domain name to use for resources. If not specified 'default' is used
// +kubebuilder:validation:Optional
Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

// Extended criteria to determine group membership. extended_criteria is implicitly "AND" with criteria
// +kubebuilder:validation:Optional
ExtendedCriteria []ExtendedCriteriaParameters `json:"extendedCriteria,omitempty" tf:"extended_criteria,omitempty"`

// Indicates the group type
// +kubebuilder:validation:Optional
GroupType *string `json:"groupType,omitempty" tf:"group_type,omitempty"`

// NSX ID for this resource
// +kubebuilder:validation:Optional
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []PolicyGroupTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyGroupTagObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyGroupTagParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyGroupSpec defines the desired state of PolicyGroup
type PolicyGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyGroupParameters `json:"forProvider"`
}

// PolicyGroupStatus defines the observed state of PolicyGroup.
type PolicyGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGroup is the Schema for the PolicyGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec              PolicyGroupSpec   `json:"spec"`
	Status            PolicyGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyGroupList contains a list of PolicyGroups
type PolicyGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyGroup `json:"items"`
}

// Repository type metadata.
var (
	PolicyGroup_Kind             = "PolicyGroup"
	PolicyGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyGroup_Kind}.String()
	PolicyGroup_KindAPIVersion   = PolicyGroup_Kind + "." + CRDGroupVersion.String()
	PolicyGroup_GroupVersionKind = CRDGroupVersion.WithKind(PolicyGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyGroup{}, &PolicyGroupList{})
}
