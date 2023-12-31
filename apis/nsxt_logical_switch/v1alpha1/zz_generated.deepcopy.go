//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressBindingObservation) DeepCopyInto(out *AddressBindingObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressBindingObservation.
func (in *AddressBindingObservation) DeepCopy() *AddressBindingObservation {
	if in == nil {
		return nil
	}
	out := new(AddressBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressBindingParameters) DeepCopyInto(out *AddressBindingParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressBindingParameters.
func (in *AddressBindingParameters) DeepCopy() *AddressBindingParameters {
	if in == nil {
		return nil
	}
	out := new(AddressBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalSwitch) DeepCopyInto(out *LogicalSwitch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalSwitch.
func (in *LogicalSwitch) DeepCopy() *LogicalSwitch {
	if in == nil {
		return nil
	}
	out := new(LogicalSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalSwitch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalSwitchList) DeepCopyInto(out *LogicalSwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogicalSwitch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalSwitchList.
func (in *LogicalSwitchList) DeepCopy() *LogicalSwitchList {
	if in == nil {
		return nil
	}
	out := new(LogicalSwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalSwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalSwitchObservation) DeepCopyInto(out *LogicalSwitchObservation) {
	*out = *in
	if in.AddressBinding != nil {
		in, out := &in.AddressBinding, &out.AddressBinding
		*out = make([]AddressBindingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdminState != nil {
		in, out := &in.AdminState, &out.AdminState
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPPoolID != nil {
		in, out := &in.IPPoolID, &out.IPPoolID
		*out = new(string)
		**out = **in
	}
	if in.MacPoolID != nil {
		in, out := &in.MacPoolID, &out.MacPoolID
		*out = new(string)
		**out = **in
	}
	if in.ReplicationMode != nil {
		in, out := &in.ReplicationMode, &out.ReplicationMode
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.SwitchingProfileID != nil {
		in, out := &in.SwitchingProfileID, &out.SwitchingProfileID
		*out = make([]SwitchingProfileIDObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TransportZoneID != nil {
		in, out := &in.TransportZoneID, &out.TransportZoneID
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(float64)
		**out = **in
	}
	if in.Vni != nil {
		in, out := &in.Vni, &out.Vni
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalSwitchObservation.
func (in *LogicalSwitchObservation) DeepCopy() *LogicalSwitchObservation {
	if in == nil {
		return nil
	}
	out := new(LogicalSwitchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalSwitchParameters) DeepCopyInto(out *LogicalSwitchParameters) {
	*out = *in
	if in.AddressBinding != nil {
		in, out := &in.AddressBinding, &out.AddressBinding
		*out = make([]AddressBindingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdminState != nil {
		in, out := &in.AdminState, &out.AdminState
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.IPPoolID != nil {
		in, out := &in.IPPoolID, &out.IPPoolID
		*out = new(string)
		**out = **in
	}
	if in.MacPoolID != nil {
		in, out := &in.MacPoolID, &out.MacPoolID
		*out = new(string)
		**out = **in
	}
	if in.ReplicationMode != nil {
		in, out := &in.ReplicationMode, &out.ReplicationMode
		*out = new(string)
		**out = **in
	}
	if in.SwitchingProfileID != nil {
		in, out := &in.SwitchingProfileID, &out.SwitchingProfileID
		*out = make([]SwitchingProfileIDParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TransportZoneID != nil {
		in, out := &in.TransportZoneID, &out.TransportZoneID
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(float64)
		**out = **in
	}
	if in.Vni != nil {
		in, out := &in.Vni, &out.Vni
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalSwitchParameters.
func (in *LogicalSwitchParameters) DeepCopy() *LogicalSwitchParameters {
	if in == nil {
		return nil
	}
	out := new(LogicalSwitchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalSwitchSpec) DeepCopyInto(out *LogicalSwitchSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalSwitchSpec.
func (in *LogicalSwitchSpec) DeepCopy() *LogicalSwitchSpec {
	if in == nil {
		return nil
	}
	out := new(LogicalSwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalSwitchStatus) DeepCopyInto(out *LogicalSwitchStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalSwitchStatus.
func (in *LogicalSwitchStatus) DeepCopy() *LogicalSwitchStatus {
	if in == nil {
		return nil
	}
	out := new(LogicalSwitchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchingProfileIDObservation) DeepCopyInto(out *SwitchingProfileIDObservation) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchingProfileIDObservation.
func (in *SwitchingProfileIDObservation) DeepCopy() *SwitchingProfileIDObservation {
	if in == nil {
		return nil
	}
	out := new(SwitchingProfileIDObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchingProfileIDParameters) DeepCopyInto(out *SwitchingProfileIDParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchingProfileIDParameters.
func (in *SwitchingProfileIDParameters) DeepCopy() *SwitchingProfileIDParameters {
	if in == nil {
		return nil
	}
	out := new(SwitchingProfileIDParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagObservation) DeepCopyInto(out *TagObservation) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagObservation.
func (in *TagObservation) DeepCopy() *TagObservation {
	if in == nil {
		return nil
	}
	out := new(TagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagParameters) DeepCopyInto(out *TagParameters) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagParameters.
func (in *TagParameters) DeepCopy() *TagParameters {
	if in == nil {
		return nil
	}
	out := new(TagParameters)
	in.DeepCopyInto(out)
	return out
}
