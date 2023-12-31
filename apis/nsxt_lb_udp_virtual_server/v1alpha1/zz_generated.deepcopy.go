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
func (in *LbUdpVirtualServer) DeepCopyInto(out *LbUdpVirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbUdpVirtualServer.
func (in *LbUdpVirtualServer) DeepCopy() *LbUdpVirtualServer {
	if in == nil {
		return nil
	}
	out := new(LbUdpVirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbUdpVirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbUdpVirtualServerList) DeepCopyInto(out *LbUdpVirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbUdpVirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbUdpVirtualServerList.
func (in *LbUdpVirtualServerList) DeepCopy() *LbUdpVirtualServerList {
	if in == nil {
		return nil
	}
	out := new(LbUdpVirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbUdpVirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbUdpVirtualServerObservation) DeepCopyInto(out *LbUdpVirtualServerObservation) {
	*out = *in
	if in.AccessLogEnabled != nil {
		in, out := &in.AccessLogEnabled, &out.AccessLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ApplicationProfileID != nil {
		in, out := &in.ApplicationProfileID, &out.ApplicationProfileID
		*out = new(string)
		**out = **in
	}
	if in.DefaultPoolMemberPorts != nil {
		in, out := &in.DefaultPoolMemberPorts, &out.DefaultPoolMemberPorts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
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
	if in.MaxConcurrentConnections != nil {
		in, out := &in.MaxConcurrentConnections, &out.MaxConcurrentConnections
		*out = new(float64)
		**out = **in
	}
	if in.MaxNewConnectionRate != nil {
		in, out := &in.MaxNewConnectionRate, &out.MaxNewConnectionRate
		*out = new(float64)
		**out = **in
	}
	if in.PersistenceProfileID != nil {
		in, out := &in.PersistenceProfileID, &out.PersistenceProfileID
		*out = new(string)
		**out = **in
	}
	if in.PoolID != nil {
		in, out := &in.PoolID, &out.PoolID
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.SorryPoolID != nil {
		in, out := &in.SorryPoolID, &out.SorryPoolID
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbUdpVirtualServerObservation.
func (in *LbUdpVirtualServerObservation) DeepCopy() *LbUdpVirtualServerObservation {
	if in == nil {
		return nil
	}
	out := new(LbUdpVirtualServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbUdpVirtualServerParameters) DeepCopyInto(out *LbUdpVirtualServerParameters) {
	*out = *in
	if in.AccessLogEnabled != nil {
		in, out := &in.AccessLogEnabled, &out.AccessLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ApplicationProfileID != nil {
		in, out := &in.ApplicationProfileID, &out.ApplicationProfileID
		*out = new(string)
		**out = **in
	}
	if in.DefaultPoolMemberPorts != nil {
		in, out := &in.DefaultPoolMemberPorts, &out.DefaultPoolMemberPorts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.MaxConcurrentConnections != nil {
		in, out := &in.MaxConcurrentConnections, &out.MaxConcurrentConnections
		*out = new(float64)
		**out = **in
	}
	if in.MaxNewConnectionRate != nil {
		in, out := &in.MaxNewConnectionRate, &out.MaxNewConnectionRate
		*out = new(float64)
		**out = **in
	}
	if in.PersistenceProfileID != nil {
		in, out := &in.PersistenceProfileID, &out.PersistenceProfileID
		*out = new(string)
		**out = **in
	}
	if in.PoolID != nil {
		in, out := &in.PoolID, &out.PoolID
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SorryPoolID != nil {
		in, out := &in.SorryPoolID, &out.SorryPoolID
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbUdpVirtualServerParameters.
func (in *LbUdpVirtualServerParameters) DeepCopy() *LbUdpVirtualServerParameters {
	if in == nil {
		return nil
	}
	out := new(LbUdpVirtualServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbUdpVirtualServerSpec) DeepCopyInto(out *LbUdpVirtualServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbUdpVirtualServerSpec.
func (in *LbUdpVirtualServerSpec) DeepCopy() *LbUdpVirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(LbUdpVirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbUdpVirtualServerStatus) DeepCopyInto(out *LbUdpVirtualServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbUdpVirtualServerStatus.
func (in *LbUdpVirtualServerStatus) DeepCopy() *LbUdpVirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(LbUdpVirtualServerStatus)
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
