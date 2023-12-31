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
func (in *LbTcpVirtualServer) DeepCopyInto(out *LbTcpVirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbTcpVirtualServer.
func (in *LbTcpVirtualServer) DeepCopy() *LbTcpVirtualServer {
	if in == nil {
		return nil
	}
	out := new(LbTcpVirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbTcpVirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbTcpVirtualServerList) DeepCopyInto(out *LbTcpVirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbTcpVirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbTcpVirtualServerList.
func (in *LbTcpVirtualServerList) DeepCopy() *LbTcpVirtualServerList {
	if in == nil {
		return nil
	}
	out := new(LbTcpVirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbTcpVirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbTcpVirtualServerObservation) DeepCopyInto(out *LbTcpVirtualServerObservation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbTcpVirtualServerObservation.
func (in *LbTcpVirtualServerObservation) DeepCopy() *LbTcpVirtualServerObservation {
	if in == nil {
		return nil
	}
	out := new(LbTcpVirtualServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbTcpVirtualServerParameters) DeepCopyInto(out *LbTcpVirtualServerParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbTcpVirtualServerParameters.
func (in *LbTcpVirtualServerParameters) DeepCopy() *LbTcpVirtualServerParameters {
	if in == nil {
		return nil
	}
	out := new(LbTcpVirtualServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbTcpVirtualServerSpec) DeepCopyInto(out *LbTcpVirtualServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbTcpVirtualServerSpec.
func (in *LbTcpVirtualServerSpec) DeepCopy() *LbTcpVirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(LbTcpVirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbTcpVirtualServerStatus) DeepCopyInto(out *LbTcpVirtualServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbTcpVirtualServerStatus.
func (in *LbTcpVirtualServerStatus) DeepCopy() *LbTcpVirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(LbTcpVirtualServerStatus)
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
