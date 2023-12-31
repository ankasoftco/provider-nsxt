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
func (in *L4PortSetNsService) DeepCopyInto(out *L4PortSetNsService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4PortSetNsService.
func (in *L4PortSetNsService) DeepCopy() *L4PortSetNsService {
	if in == nil {
		return nil
	}
	out := new(L4PortSetNsService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L4PortSetNsService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4PortSetNsServiceList) DeepCopyInto(out *L4PortSetNsServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]L4PortSetNsService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4PortSetNsServiceList.
func (in *L4PortSetNsServiceList) DeepCopy() *L4PortSetNsServiceList {
	if in == nil {
		return nil
	}
	out := new(L4PortSetNsServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *L4PortSetNsServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4PortSetNsServiceObservation) DeepCopyInto(out *L4PortSetNsServiceObservation) {
	*out = *in
	if in.DefaultService != nil {
		in, out := &in.DefaultService, &out.DefaultService
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DestinationPorts != nil {
		in, out := &in.DestinationPorts, &out.DestinationPorts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.SourcePorts != nil {
		in, out := &in.SourcePorts, &out.SourcePorts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4PortSetNsServiceObservation.
func (in *L4PortSetNsServiceObservation) DeepCopy() *L4PortSetNsServiceObservation {
	if in == nil {
		return nil
	}
	out := new(L4PortSetNsServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4PortSetNsServiceParameters) DeepCopyInto(out *L4PortSetNsServiceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DestinationPorts != nil {
		in, out := &in.DestinationPorts, &out.DestinationPorts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SourcePorts != nil {
		in, out := &in.SourcePorts, &out.SourcePorts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4PortSetNsServiceParameters.
func (in *L4PortSetNsServiceParameters) DeepCopy() *L4PortSetNsServiceParameters {
	if in == nil {
		return nil
	}
	out := new(L4PortSetNsServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4PortSetNsServiceSpec) DeepCopyInto(out *L4PortSetNsServiceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4PortSetNsServiceSpec.
func (in *L4PortSetNsServiceSpec) DeepCopy() *L4PortSetNsServiceSpec {
	if in == nil {
		return nil
	}
	out := new(L4PortSetNsServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4PortSetNsServiceStatus) DeepCopyInto(out *L4PortSetNsServiceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4PortSetNsServiceStatus.
func (in *L4PortSetNsServiceStatus) DeepCopy() *L4PortSetNsServiceStatus {
	if in == nil {
		return nil
	}
	out := new(L4PortSetNsServiceStatus)
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
