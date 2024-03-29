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




type AlgorithmEntryObservation struct {


// Algorithm
Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// A single destination port
DestinationPort *string `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Set of source ports or ranges
SourcePorts []*string `json:"sourcePorts,omitempty" tf:"source_ports,omitempty"`
}


type AlgorithmEntryParameters struct {


// Algorithm
// +kubebuilder:validation:Required
Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// A single destination port
// +kubebuilder:validation:Required
DestinationPort *string `json:"destinationPort" tf:"destination_port,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Set of source ports or ranges
// +kubebuilder:validation:Optional
SourcePorts []*string `json:"sourcePorts,omitempty" tf:"source_ports,omitempty"`
}


type EtherTypeEntryObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Type of the encapsulated protocol
EtherType *float64 `json:"etherType,omitempty" tf:"ether_type,omitempty"`
}


type EtherTypeEntryParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Type of the encapsulated protocol
// +kubebuilder:validation:Required
EtherType *float64 `json:"etherType" tf:"ether_type,omitempty"`
}


type IPProtocolEntryObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// IP protocol number
Protocol *float64 `json:"protocol,omitempty" tf:"protocol,omitempty"`
}


type IPProtocolEntryParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// IP protocol number
// +kubebuilder:validation:Required
Protocol *float64 `json:"protocol" tf:"protocol,omitempty"`
}


type IcmpEntryObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// ICMP message code
IcmpCode *string `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

// ICMP message type
IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

// Version of ICMP protocol (ICMPv4/ICMPv6)
Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}


type IcmpEntryParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// ICMP message code
// +kubebuilder:validation:Optional
IcmpCode *string `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

// ICMP message type
// +kubebuilder:validation:Optional
IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

// Version of ICMP protocol (ICMPv4/ICMPv6)
// +kubebuilder:validation:Required
Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}


type IgmpEntryObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}


type IgmpEntryParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}


type L4PortSetEntryObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Set of destination ports
DestinationPorts []*string `json:"destinationPorts,omitempty" tf:"destination_ports,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// L4 Protocol
Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

// Set of source ports
SourcePorts []*string `json:"sourcePorts,omitempty" tf:"source_ports,omitempty"`
}


type L4PortSetEntryParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Set of destination ports
// +kubebuilder:validation:Optional
DestinationPorts []*string `json:"destinationPorts,omitempty" tf:"destination_ports,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// L4 Protocol
// +kubebuilder:validation:Required
Protocol *string `json:"protocol" tf:"protocol,omitempty"`

// Set of source ports
// +kubebuilder:validation:Optional
SourcePorts []*string `json:"sourcePorts,omitempty" tf:"source_ports,omitempty"`
}


type NestedServiceEntryObservation struct {


// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Nested Service Path
NestedServicePath *string `json:"nestedServicePath,omitempty" tf:"nested_service_path,omitempty"`
}


type NestedServiceEntryParameters struct {


// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Nested Service Path
// +kubebuilder:validation:Required
NestedServicePath *string `json:"nestedServicePath" tf:"nested_service_path,omitempty"`
}


type PolicyServiceContextObservation struct {


// Id of the project which the resource belongs to.
ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}


type PolicyServiceContextParameters struct {


// Id of the project which the resource belongs to.
// +kubebuilder:validation:Required
ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}


type PolicyServiceObservation struct {


// Algorithm type service entry
AlgorithmEntry []AlgorithmEntryObservation `json:"algorithmEntry,omitempty" tf:"algorithm_entry,omitempty"`

// Resource context
Context []PolicyServiceContextObservation `json:"context,omitempty" tf:"context,omitempty"`

// Description for this resource
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Ether type service entry
EtherTypeEntry []EtherTypeEntryObservation `json:"etherTypeEntry,omitempty" tf:"ether_type_entry,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// IP Protocol type service entry
IPProtocolEntry []IPProtocolEntryObservation `json:"ipProtocolEntry,omitempty" tf:"ip_protocol_entry,omitempty"`

// ICMP type service entry
IcmpEntry []IcmpEntryObservation `json:"icmpEntry,omitempty" tf:"icmp_entry,omitempty"`

// IGMP type service entry
IgmpEntry []IgmpEntryObservation `json:"igmpEntry,omitempty" tf:"igmp_entry,omitempty"`

// L4 port set type service entry
L4PortSetEntry []L4PortSetEntryObservation `json:"l4PortSetEntry,omitempty" tf:"l4_port_set_entry,omitempty"`

// Nested service service entry
NestedServiceEntry []NestedServiceEntryObservation `json:"nestedServiceEntry,omitempty" tf:"nested_service_entry,omitempty"`

// NSX ID for this resource
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Policy path for this resource
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected
Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

// Set of opaque identifiers meaningful to the user
Tag []PolicyServiceTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyServiceParameters struct {


// Algorithm type service entry
// +kubebuilder:validation:Optional
AlgorithmEntry []AlgorithmEntryParameters `json:"algorithmEntry,omitempty" tf:"algorithm_entry,omitempty"`

// Resource context
// +kubebuilder:validation:Optional
Context []PolicyServiceContextParameters `json:"context,omitempty" tf:"context,omitempty"`

// Description for this resource
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Display name for this resource
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// Ether type service entry
// +kubebuilder:validation:Optional
EtherTypeEntry []EtherTypeEntryParameters `json:"etherTypeEntry,omitempty" tf:"ether_type_entry,omitempty"`

// IP Protocol type service entry
// +kubebuilder:validation:Optional
IPProtocolEntry []IPProtocolEntryParameters `json:"ipProtocolEntry,omitempty" tf:"ip_protocol_entry,omitempty"`

// ICMP type service entry
// +kubebuilder:validation:Optional
IcmpEntry []IcmpEntryParameters `json:"icmpEntry,omitempty" tf:"icmp_entry,omitempty"`

// IGMP type service entry
// +kubebuilder:validation:Optional
IgmpEntry []IgmpEntryParameters `json:"igmpEntry,omitempty" tf:"igmp_entry,omitempty"`

// L4 port set type service entry
// +kubebuilder:validation:Optional
L4PortSetEntry []L4PortSetEntryParameters `json:"l4PortSetEntry,omitempty" tf:"l4_port_set_entry,omitempty"`

// Nested service service entry
// +kubebuilder:validation:Optional
NestedServiceEntry []NestedServiceEntryParameters `json:"nestedServiceEntry,omitempty" tf:"nested_service_entry,omitempty"`

// NSX ID for this resource
// +kubebuilder:validation:Optional
NsxID *string `json:"nsxId,omitempty" tf:"nsx_id,omitempty"`

// Set of opaque identifiers meaningful to the user
// +kubebuilder:validation:Optional
Tag []PolicyServiceTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyServiceTagObservation struct {


Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}


type PolicyServiceTagParameters struct {


// +kubebuilder:validation:Optional
Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

// +kubebuilder:validation:Optional
Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// PolicyServiceSpec defines the desired state of PolicyService
type PolicyServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       PolicyServiceParameters `json:"forProvider"`
}

// PolicyServiceStatus defines the observed state of PolicyService.
type PolicyServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          PolicyServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyService is the Schema for the PolicyServices API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nsxt}
type PolicyService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec              PolicyServiceSpec   `json:"spec"`
	Status            PolicyServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyServiceList contains a list of PolicyServices
type PolicyServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyService `json:"items"`
}

// Repository type metadata.
var (
	PolicyService_Kind             = "PolicyService"
	PolicyService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyService_Kind}.String()
	PolicyService_KindAPIVersion   = PolicyService_Kind + "." + CRDGroupVersion.String()
	PolicyService_GroupVersionKind = CRDGroupVersion.WithKind(PolicyService_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyService{}, &PolicyServiceList{})
}
