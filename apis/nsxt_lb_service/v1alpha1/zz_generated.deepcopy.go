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
func (in *LbService) DeepCopyInto(out *LbService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbService.
func (in *LbService) DeepCopy() *LbService {
	if in == nil {
		return nil
	}
	out := new(LbService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbServiceList) DeepCopyInto(out *LbServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbServiceList.
func (in *LbServiceList) DeepCopy() *LbServiceList {
	if in == nil {
		return nil
	}
	out := new(LbServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbServiceObservation) DeepCopyInto(out *LbServiceObservation) {
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ErrorLogLevel != nil {
		in, out := &in.ErrorLogLevel, &out.ErrorLogLevel
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
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
	if in.Size != nil {
		in, out := &in.Size, &out.Size
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
	if in.VirtualServerIds != nil {
		in, out := &in.VirtualServerIds, &out.VirtualServerIds
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbServiceObservation.
func (in *LbServiceObservation) DeepCopy() *LbServiceObservation {
	if in == nil {
		return nil
	}
	out := new(LbServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbServiceParameters) DeepCopyInto(out *LbServiceParameters) {
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ErrorLogLevel != nil {
		in, out := &in.ErrorLogLevel, &out.ErrorLogLevel
		*out = new(string)
		**out = **in
	}
	if in.LogicalRouterID != nil {
		in, out := &in.LogicalRouterID, &out.LogicalRouterID
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
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
	if in.VirtualServerIds != nil {
		in, out := &in.VirtualServerIds, &out.VirtualServerIds
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbServiceParameters.
func (in *LbServiceParameters) DeepCopy() *LbServiceParameters {
	if in == nil {
		return nil
	}
	out := new(LbServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbServiceSpec) DeepCopyInto(out *LbServiceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbServiceSpec.
func (in *LbServiceSpec) DeepCopy() *LbServiceSpec {
	if in == nil {
		return nil
	}
	out := new(LbServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbServiceStatus) DeepCopyInto(out *LbServiceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbServiceStatus.
func (in *LbServiceStatus) DeepCopy() *LbServiceStatus {
	if in == nil {
		return nil
	}
	out := new(LbServiceStatus)
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
