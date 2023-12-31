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
func (in *LbHttpMonitor) DeepCopyInto(out *LbHttpMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpMonitor.
func (in *LbHttpMonitor) DeepCopy() *LbHttpMonitor {
	if in == nil {
		return nil
	}
	out := new(LbHttpMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbHttpMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpMonitorList) DeepCopyInto(out *LbHttpMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbHttpMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpMonitorList.
func (in *LbHttpMonitorList) DeepCopy() *LbHttpMonitorList {
	if in == nil {
		return nil
	}
	out := new(LbHttpMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbHttpMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpMonitorObservation) DeepCopyInto(out *LbHttpMonitorObservation) {
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
	if in.FallCount != nil {
		in, out := &in.FallCount, &out.FallCount
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(float64)
		**out = **in
	}
	if in.MonitorPort != nil {
		in, out := &in.MonitorPort, &out.MonitorPort
		*out = new(string)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestHeader != nil {
		in, out := &in.RequestHeader, &out.RequestHeader
		*out = make([]RequestHeaderObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequestMethod != nil {
		in, out := &in.RequestMethod, &out.RequestMethod
		*out = new(string)
		**out = **in
	}
	if in.RequestURL != nil {
		in, out := &in.RequestURL, &out.RequestURL
		*out = new(string)
		**out = **in
	}
	if in.RequestVersion != nil {
		in, out := &in.RequestVersion, &out.RequestVersion
		*out = new(string)
		**out = **in
	}
	if in.ResponseBody != nil {
		in, out := &in.ResponseBody, &out.ResponseBody
		*out = new(string)
		**out = **in
	}
	if in.ResponseStatusCodes != nil {
		in, out := &in.ResponseStatusCodes, &out.ResponseStatusCodes
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.RiseCount != nil {
		in, out := &in.RiseCount, &out.RiseCount
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
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpMonitorObservation.
func (in *LbHttpMonitorObservation) DeepCopy() *LbHttpMonitorObservation {
	if in == nil {
		return nil
	}
	out := new(LbHttpMonitorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpMonitorParameters) DeepCopyInto(out *LbHttpMonitorParameters) {
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
	if in.FallCount != nil {
		in, out := &in.FallCount, &out.FallCount
		*out = new(float64)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(float64)
		**out = **in
	}
	if in.MonitorPort != nil {
		in, out := &in.MonitorPort, &out.MonitorPort
		*out = new(string)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestHeader != nil {
		in, out := &in.RequestHeader, &out.RequestHeader
		*out = make([]RequestHeaderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequestMethod != nil {
		in, out := &in.RequestMethod, &out.RequestMethod
		*out = new(string)
		**out = **in
	}
	if in.RequestURL != nil {
		in, out := &in.RequestURL, &out.RequestURL
		*out = new(string)
		**out = **in
	}
	if in.RequestVersion != nil {
		in, out := &in.RequestVersion, &out.RequestVersion
		*out = new(string)
		**out = **in
	}
	if in.ResponseBody != nil {
		in, out := &in.ResponseBody, &out.ResponseBody
		*out = new(string)
		**out = **in
	}
	if in.ResponseStatusCodes != nil {
		in, out := &in.ResponseStatusCodes, &out.ResponseStatusCodes
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.RiseCount != nil {
		in, out := &in.RiseCount, &out.RiseCount
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpMonitorParameters.
func (in *LbHttpMonitorParameters) DeepCopy() *LbHttpMonitorParameters {
	if in == nil {
		return nil
	}
	out := new(LbHttpMonitorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpMonitorSpec) DeepCopyInto(out *LbHttpMonitorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpMonitorSpec.
func (in *LbHttpMonitorSpec) DeepCopy() *LbHttpMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(LbHttpMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpMonitorStatus) DeepCopyInto(out *LbHttpMonitorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpMonitorStatus.
func (in *LbHttpMonitorStatus) DeepCopy() *LbHttpMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(LbHttpMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestHeaderObservation) DeepCopyInto(out *RequestHeaderObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestHeaderObservation.
func (in *RequestHeaderObservation) DeepCopy() *RequestHeaderObservation {
	if in == nil {
		return nil
	}
	out := new(RequestHeaderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestHeaderParameters) DeepCopyInto(out *RequestHeaderParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestHeaderParameters.
func (in *RequestHeaderParameters) DeepCopy() *RequestHeaderParameters {
	if in == nil {
		return nil
	}
	out := new(RequestHeaderParameters)
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
