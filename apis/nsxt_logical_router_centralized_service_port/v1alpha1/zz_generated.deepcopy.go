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
func (in *LogicalRouterCentralizedServicePort) DeepCopyInto(out *LogicalRouterCentralizedServicePort) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalRouterCentralizedServicePort.
func (in *LogicalRouterCentralizedServicePort) DeepCopy() *LogicalRouterCentralizedServicePort {
	if in == nil {
		return nil
	}
	out := new(LogicalRouterCentralizedServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalRouterCentralizedServicePort) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalRouterCentralizedServicePortList) DeepCopyInto(out *LogicalRouterCentralizedServicePortList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogicalRouterCentralizedServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalRouterCentralizedServicePortList.
func (in *LogicalRouterCentralizedServicePortList) DeepCopy() *LogicalRouterCentralizedServicePortList {
	if in == nil {
		return nil
	}
	out := new(LogicalRouterCentralizedServicePortList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalRouterCentralizedServicePortList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalRouterCentralizedServicePortObservation) DeepCopyInto(out *LogicalRouterCentralizedServicePortObservation) {
	*out = *in
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
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.LinkedLogicalSwitchPortID != nil {
		in, out := &in.LinkedLogicalSwitchPortID, &out.LinkedLogicalSwitchPortID
		*out = new(string)
		**out = **in
	}
	if in.LogicalRouterID != nil {
		in, out := &in.LogicalRouterID, &out.LogicalRouterID
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UrpfMode != nil {
		in, out := &in.UrpfMode, &out.UrpfMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalRouterCentralizedServicePortObservation.
func (in *LogicalRouterCentralizedServicePortObservation) DeepCopy() *LogicalRouterCentralizedServicePortObservation {
	if in == nil {
		return nil
	}
	out := new(LogicalRouterCentralizedServicePortObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalRouterCentralizedServicePortParameters) DeepCopyInto(out *LogicalRouterCentralizedServicePortParameters) {
	*out = *in
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
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.LinkedLogicalSwitchPortID != nil {
		in, out := &in.LinkedLogicalSwitchPortID, &out.LinkedLogicalSwitchPortID
		*out = new(string)
		**out = **in
	}
	if in.LogicalRouterID != nil {
		in, out := &in.LogicalRouterID, &out.LogicalRouterID
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UrpfMode != nil {
		in, out := &in.UrpfMode, &out.UrpfMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalRouterCentralizedServicePortParameters.
func (in *LogicalRouterCentralizedServicePortParameters) DeepCopy() *LogicalRouterCentralizedServicePortParameters {
	if in == nil {
		return nil
	}
	out := new(LogicalRouterCentralizedServicePortParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalRouterCentralizedServicePortSpec) DeepCopyInto(out *LogicalRouterCentralizedServicePortSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalRouterCentralizedServicePortSpec.
func (in *LogicalRouterCentralizedServicePortSpec) DeepCopy() *LogicalRouterCentralizedServicePortSpec {
	if in == nil {
		return nil
	}
	out := new(LogicalRouterCentralizedServicePortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalRouterCentralizedServicePortStatus) DeepCopyInto(out *LogicalRouterCentralizedServicePortStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalRouterCentralizedServicePortStatus.
func (in *LogicalRouterCentralizedServicePortStatus) DeepCopy() *LogicalRouterCentralizedServicePortStatus {
	if in == nil {
		return nil
	}
	out := new(LogicalRouterCentralizedServicePortStatus)
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
