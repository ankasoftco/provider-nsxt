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
func (in *RateLimitsObservation) DeepCopyInto(out *RateLimitsObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RxBroadcast != nil {
		in, out := &in.RxBroadcast, &out.RxBroadcast
		*out = new(float64)
		**out = **in
	}
	if in.RxMulticast != nil {
		in, out := &in.RxMulticast, &out.RxMulticast
		*out = new(float64)
		**out = **in
	}
	if in.TxBroadcast != nil {
		in, out := &in.TxBroadcast, &out.TxBroadcast
		*out = new(float64)
		**out = **in
	}
	if in.TxMulticast != nil {
		in, out := &in.TxMulticast, &out.TxMulticast
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitsObservation.
func (in *RateLimitsObservation) DeepCopy() *RateLimitsObservation {
	if in == nil {
		return nil
	}
	out := new(RateLimitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitsParameters) DeepCopyInto(out *RateLimitsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RxBroadcast != nil {
		in, out := &in.RxBroadcast, &out.RxBroadcast
		*out = new(float64)
		**out = **in
	}
	if in.RxMulticast != nil {
		in, out := &in.RxMulticast, &out.RxMulticast
		*out = new(float64)
		**out = **in
	}
	if in.TxBroadcast != nil {
		in, out := &in.TxBroadcast, &out.TxBroadcast
		*out = new(float64)
		**out = **in
	}
	if in.TxMulticast != nil {
		in, out := &in.TxMulticast, &out.TxMulticast
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitsParameters.
func (in *RateLimitsParameters) DeepCopy() *RateLimitsParameters {
	if in == nil {
		return nil
	}
	out := new(RateLimitsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSecuritySwitchingProfile) DeepCopyInto(out *SwitchSecuritySwitchingProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSecuritySwitchingProfile.
func (in *SwitchSecuritySwitchingProfile) DeepCopy() *SwitchSecuritySwitchingProfile {
	if in == nil {
		return nil
	}
	out := new(SwitchSecuritySwitchingProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchSecuritySwitchingProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSecuritySwitchingProfileList) DeepCopyInto(out *SwitchSecuritySwitchingProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwitchSecuritySwitchingProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSecuritySwitchingProfileList.
func (in *SwitchSecuritySwitchingProfileList) DeepCopy() *SwitchSecuritySwitchingProfileList {
	if in == nil {
		return nil
	}
	out := new(SwitchSecuritySwitchingProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchSecuritySwitchingProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSecuritySwitchingProfileObservation) DeepCopyInto(out *SwitchSecuritySwitchingProfileObservation) {
	*out = *in
	if in.BlockClientDHCP != nil {
		in, out := &in.BlockClientDHCP, &out.BlockClientDHCP
		*out = new(bool)
		**out = **in
	}
	if in.BlockNonIP != nil {
		in, out := &in.BlockNonIP, &out.BlockNonIP
		*out = new(bool)
		**out = **in
	}
	if in.BlockServerDHCP != nil {
		in, out := &in.BlockServerDHCP, &out.BlockServerDHCP
		*out = new(bool)
		**out = **in
	}
	if in.BpduFilterEnabled != nil {
		in, out := &in.BpduFilterEnabled, &out.BpduFilterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BpduFilterWhitelist != nil {
		in, out := &in.BpduFilterWhitelist, &out.BpduFilterWhitelist
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RateLimits != nil {
		in, out := &in.RateLimits, &out.RateLimits
		*out = make([]RateLimitsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSecuritySwitchingProfileObservation.
func (in *SwitchSecuritySwitchingProfileObservation) DeepCopy() *SwitchSecuritySwitchingProfileObservation {
	if in == nil {
		return nil
	}
	out := new(SwitchSecuritySwitchingProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSecuritySwitchingProfileParameters) DeepCopyInto(out *SwitchSecuritySwitchingProfileParameters) {
	*out = *in
	if in.BlockClientDHCP != nil {
		in, out := &in.BlockClientDHCP, &out.BlockClientDHCP
		*out = new(bool)
		**out = **in
	}
	if in.BlockNonIP != nil {
		in, out := &in.BlockNonIP, &out.BlockNonIP
		*out = new(bool)
		**out = **in
	}
	if in.BlockServerDHCP != nil {
		in, out := &in.BlockServerDHCP, &out.BlockServerDHCP
		*out = new(bool)
		**out = **in
	}
	if in.BpduFilterEnabled != nil {
		in, out := &in.BpduFilterEnabled, &out.BpduFilterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BpduFilterWhitelist != nil {
		in, out := &in.BpduFilterWhitelist, &out.BpduFilterWhitelist
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
	if in.RateLimits != nil {
		in, out := &in.RateLimits, &out.RateLimits
		*out = make([]RateLimitsParameters, len(*in))
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSecuritySwitchingProfileParameters.
func (in *SwitchSecuritySwitchingProfileParameters) DeepCopy() *SwitchSecuritySwitchingProfileParameters {
	if in == nil {
		return nil
	}
	out := new(SwitchSecuritySwitchingProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSecuritySwitchingProfileSpec) DeepCopyInto(out *SwitchSecuritySwitchingProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSecuritySwitchingProfileSpec.
func (in *SwitchSecuritySwitchingProfileSpec) DeepCopy() *SwitchSecuritySwitchingProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchSecuritySwitchingProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSecuritySwitchingProfileStatus) DeepCopyInto(out *SwitchSecuritySwitchingProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSecuritySwitchingProfileStatus.
func (in *SwitchSecuritySwitchingProfileStatus) DeepCopy() *SwitchSecuritySwitchingProfileStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchSecuritySwitchingProfileStatus)
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
