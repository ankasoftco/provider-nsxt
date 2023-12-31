//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnSession) DeepCopyInto(out *PolicyIpsecVpnSession) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnSession.
func (in *PolicyIpsecVpnSession) DeepCopy() *PolicyIpsecVpnSession {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnSession)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyIpsecVpnSession) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnSessionList) DeepCopyInto(out *PolicyIpsecVpnSessionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyIpsecVpnSession, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnSessionList.
func (in *PolicyIpsecVpnSessionList) DeepCopy() *PolicyIpsecVpnSessionList {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnSessionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyIpsecVpnSessionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnSessionObservation) DeepCopyInto(out *PolicyIpsecVpnSessionObservation) {
	*out = *in
	if in.AuthenticationMode != nil {
		in, out := &in.AuthenticationMode, &out.AuthenticationMode
		*out = new(string)
		**out = **in
	}
	if in.ComplianceSuite != nil {
		in, out := &in.ComplianceSuite, &out.ComplianceSuite
		*out = new(string)
		**out = **in
	}
	if in.ConnectionInitiationMode != nil {
		in, out := &in.ConnectionInitiationMode, &out.ConnectionInitiationMode
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DpdProfilePath != nil {
		in, out := &in.DpdProfilePath, &out.DpdProfilePath
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
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IkeProfilePath != nil {
		in, out := &in.IkeProfilePath, &out.IkeProfilePath
		*out = new(string)
		**out = **in
	}
	if in.LocalEndpointPath != nil {
		in, out := &in.LocalEndpointPath, &out.LocalEndpointPath
		*out = new(string)
		**out = **in
	}
	if in.MaxSegmentSize != nil {
		in, out := &in.MaxSegmentSize, &out.MaxSegmentSize
		*out = new(float64)
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
	if in.PeerAddress != nil {
		in, out := &in.PeerAddress, &out.PeerAddress
		*out = new(string)
		**out = **in
	}
	if in.PeerID != nil {
		in, out := &in.PeerID, &out.PeerID
		*out = new(string)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.TunnelProfilePath != nil {
		in, out := &in.TunnelProfilePath, &out.TunnelProfilePath
		*out = new(string)
		**out = **in
	}
	if in.VPNType != nil {
		in, out := &in.VPNType, &out.VPNType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnSessionObservation.
func (in *PolicyIpsecVpnSessionObservation) DeepCopy() *PolicyIpsecVpnSessionObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnSessionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnSessionParameters) DeepCopyInto(out *PolicyIpsecVpnSessionParameters) {
	*out = *in
	if in.AuthenticationMode != nil {
		in, out := &in.AuthenticationMode, &out.AuthenticationMode
		*out = new(string)
		**out = **in
	}
	if in.ComplianceSuite != nil {
		in, out := &in.ComplianceSuite, &out.ComplianceSuite
		*out = new(string)
		**out = **in
	}
	if in.ConnectionInitiationMode != nil {
		in, out := &in.ConnectionInitiationMode, &out.ConnectionInitiationMode
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DpdProfilePath != nil {
		in, out := &in.DpdProfilePath, &out.DpdProfilePath
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IkeProfilePath != nil {
		in, out := &in.IkeProfilePath, &out.IkeProfilePath
		*out = new(string)
		**out = **in
	}
	if in.LocalEndpointPath != nil {
		in, out := &in.LocalEndpointPath, &out.LocalEndpointPath
		*out = new(string)
		**out = **in
	}
	if in.MaxSegmentSize != nil {
		in, out := &in.MaxSegmentSize, &out.MaxSegmentSize
		*out = new(float64)
		**out = **in
	}
	if in.NsxID != nil {
		in, out := &in.NsxID, &out.NsxID
		*out = new(string)
		**out = **in
	}
	if in.PeerAddress != nil {
		in, out := &in.PeerAddress, &out.PeerAddress
		*out = new(string)
		**out = **in
	}
	if in.PeerID != nil {
		in, out := &in.PeerID, &out.PeerID
		*out = new(string)
		**out = **in
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.PskSecretRef != nil {
		in, out := &in.PskSecretRef, &out.PskSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.TunnelProfilePath != nil {
		in, out := &in.TunnelProfilePath, &out.TunnelProfilePath
		*out = new(string)
		**out = **in
	}
	if in.VPNType != nil {
		in, out := &in.VPNType, &out.VPNType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnSessionParameters.
func (in *PolicyIpsecVpnSessionParameters) DeepCopy() *PolicyIpsecVpnSessionParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnSessionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnSessionSpec) DeepCopyInto(out *PolicyIpsecVpnSessionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnSessionSpec.
func (in *PolicyIpsecVpnSessionSpec) DeepCopy() *PolicyIpsecVpnSessionSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnSessionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIpsecVpnSessionStatus) DeepCopyInto(out *PolicyIpsecVpnSessionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIpsecVpnSessionStatus.
func (in *PolicyIpsecVpnSessionStatus) DeepCopy() *PolicyIpsecVpnSessionStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyIpsecVpnSessionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObservation) DeepCopyInto(out *RuleObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NsxID != nil {
		in, out := &in.NsxID, &out.NsxID
		*out = new(string)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObservation.
func (in *RuleObservation) DeepCopy() *RuleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleParameters) DeepCopyInto(out *RuleParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleParameters.
func (in *RuleParameters) DeepCopy() *RuleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleParameters)
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
