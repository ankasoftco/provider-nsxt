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
func (in *PolicyIpsecVpnLocalEndpoint) DeepCopyInto(out *PolicyIpsecVpnLocalEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnLocalEndpoint.
func (in *PolicyIpsecVpnLocalEndpoint) DeepCopy() *PolicyIpsecVpnLocalEndpoint {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnLocalEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyIpsecVpnLocalEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnLocalEndpointList) DeepCopyInto(out *PolicyIpsecVpnLocalEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyIpsecVpnLocalEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnLocalEndpointList.
func (in *PolicyIpsecVpnLocalEndpointList) DeepCopy() *PolicyIpsecVpnLocalEndpointList {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnLocalEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyIpsecVpnLocalEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnLocalEndpointObservation) DeepCopyInto(out *PolicyIpsecVpnLocalEndpointObservation) {
	*out = *in
	if in.CertificatePath != nil {
		in, out := &in.CertificatePath, &out.CertificatePath
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
	if in.LocalAddress != nil {
		in, out := &in.LocalAddress, &out.LocalAddress
		*out = new(string)
		**out = **in
	}
	if in.LocalID != nil {
		in, out := &in.LocalID, &out.LocalID
		*out = new(string)
		**out = **in
	}
	if in.NsxID != nil {
		in, out := &in.NsxID, &out.NsxID
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.ServicePath != nil {
		in, out := &in.ServicePath, &out.ServicePath
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
	if in.TrustCAPaths != nil {
		in, out := &in.TrustCAPaths, &out.TrustCAPaths
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TrustCrlPaths != nil {
		in, out := &in.TrustCrlPaths, &out.TrustCrlPaths
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnLocalEndpointObservation.
func (in *PolicyIpsecVpnLocalEndpointObservation) DeepCopy() *PolicyIpsecVpnLocalEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnLocalEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnLocalEndpointParameters) DeepCopyInto(out *PolicyIpsecVpnLocalEndpointParameters) {
	*out = *in
	if in.CertificatePath != nil {
		in, out := &in.CertificatePath, &out.CertificatePath
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
	if in.LocalAddress != nil {
		in, out := &in.LocalAddress, &out.LocalAddress
		*out = new(string)
		**out = **in
	}
	if in.LocalID != nil {
		in, out := &in.LocalID, &out.LocalID
		*out = new(string)
		**out = **in
	}
	if in.NsxID != nil {
		in, out := &in.NsxID, &out.NsxID
		*out = new(string)
		**out = **in
	}
	if in.ServicePath != nil {
		in, out := &in.ServicePath, &out.ServicePath
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
	if in.TrustCAPaths != nil {
		in, out := &in.TrustCAPaths, &out.TrustCAPaths
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TrustCrlPaths != nil {
		in, out := &in.TrustCrlPaths, &out.TrustCrlPaths
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnLocalEndpointParameters.
func (in *PolicyIpsecVpnLocalEndpointParameters) DeepCopy() *PolicyIpsecVpnLocalEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnLocalEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnLocalEndpointSpec) DeepCopyInto(out *PolicyIpsecVpnLocalEndpointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnLocalEndpointSpec.
func (in *PolicyIpsecVpnLocalEndpointSpec) DeepCopy() *PolicyIpsecVpnLocalEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnLocalEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnLocalEndpointStatus) DeepCopyInto(out *PolicyIpsecVpnLocalEndpointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnLocalEndpointStatus.
func (in *PolicyIpsecVpnLocalEndpointStatus) DeepCopy() *PolicyIpsecVpnLocalEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnLocalEndpointStatus)
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
