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
func (in *ActiveObservation) DeepCopyInto(out *ActiveObservation) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveObservation.
func (in *ActiveObservation) DeepCopy() *ActiveObservation {
	if in == nil {
		return nil
	}
	out := new(ActiveObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveParameters) DeepCopyInto(out *ActiveParameters) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveParameters.
func (in *ActiveParameters) DeepCopy() *ActiveParameters {
	if in == nil {
		return nil
	}
	out := new(ActiveParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LagObservation) DeepCopyInto(out *LagObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LoadBalanceAlgorithm != nil {
		in, out := &in.LoadBalanceAlgorithm, &out.LoadBalanceAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumberOfUplinks != nil {
		in, out := &in.NumberOfUplinks, &out.NumberOfUplinks
		*out = new(float64)
		**out = **in
	}
	if in.TimeoutType != nil {
		in, out := &in.TimeoutType, &out.TimeoutType
		*out = new(string)
		**out = **in
	}
	if in.Uplink != nil {
		in, out := &in.Uplink, &out.Uplink
		*out = make([]UplinkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LagObservation.
func (in *LagObservation) DeepCopy() *LagObservation {
	if in == nil {
		return nil
	}
	out := new(LagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LagParameters) DeepCopyInto(out *LagParameters) {
	*out = *in
	if in.LoadBalanceAlgorithm != nil {
		in, out := &in.LoadBalanceAlgorithm, &out.LoadBalanceAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumberOfUplinks != nil {
		in, out := &in.NumberOfUplinks, &out.NumberOfUplinks
		*out = new(float64)
		**out = **in
	}
	if in.TimeoutType != nil {
		in, out := &in.TimeoutType, &out.TimeoutType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LagParameters.
func (in *LagParameters) DeepCopy() *LagParameters {
	if in == nil {
		return nil
	}
	out := new(LagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedTeamingObservation) DeepCopyInto(out *NamedTeamingObservation) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]ActiveObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Standby != nil {
		in, out := &in.Standby, &out.Standby
		*out = make([]StandbyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedTeamingObservation.
func (in *NamedTeamingObservation) DeepCopy() *NamedTeamingObservation {
	if in == nil {
		return nil
	}
	out := new(NamedTeamingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedTeamingParameters) DeepCopyInto(out *NamedTeamingParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]ActiveParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Standby != nil {
		in, out := &in.Standby, &out.Standby
		*out = make([]StandbyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedTeamingParameters.
func (in *NamedTeamingParameters) DeepCopy() *NamedTeamingParameters {
	if in == nil {
		return nil
	}
	out := new(NamedTeamingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyObservation) DeepCopyInto(out *StandbyObservation) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyObservation.
func (in *StandbyObservation) DeepCopy() *StandbyObservation {
	if in == nil {
		return nil
	}
	out := new(StandbyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyParameters) DeepCopyInto(out *StandbyParameters) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyParameters.
func (in *StandbyParameters) DeepCopy() *StandbyParameters {
	if in == nil {
		return nil
	}
	out := new(StandbyParameters)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamingActiveObservation) DeepCopyInto(out *TeamingActiveObservation) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamingActiveObservation.
func (in *TeamingActiveObservation) DeepCopy() *TeamingActiveObservation {
	if in == nil {
		return nil
	}
	out := new(TeamingActiveObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamingActiveParameters) DeepCopyInto(out *TeamingActiveParameters) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamingActiveParameters.
func (in *TeamingActiveParameters) DeepCopy() *TeamingActiveParameters {
	if in == nil {
		return nil
	}
	out := new(TeamingActiveParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamingObservation) DeepCopyInto(out *TeamingObservation) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]TeamingActiveObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Standby != nil {
		in, out := &in.Standby, &out.Standby
		*out = make([]TeamingStandbyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamingObservation.
func (in *TeamingObservation) DeepCopy() *TeamingObservation {
	if in == nil {
		return nil
	}
	out := new(TeamingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamingParameters) DeepCopyInto(out *TeamingParameters) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]TeamingActiveParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Standby != nil {
		in, out := &in.Standby, &out.Standby
		*out = make([]TeamingStandbyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamingParameters.
func (in *TeamingParameters) DeepCopy() *TeamingParameters {
	if in == nil {
		return nil
	}
	out := new(TeamingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamingStandbyObservation) DeepCopyInto(out *TeamingStandbyObservation) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamingStandbyObservation.
func (in *TeamingStandbyObservation) DeepCopy() *TeamingStandbyObservation {
	if in == nil {
		return nil
	}
	out := new(TeamingStandbyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamingStandbyParameters) DeepCopyInto(out *TeamingStandbyParameters) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamingStandbyParameters.
func (in *TeamingStandbyParameters) DeepCopy() *TeamingStandbyParameters {
	if in == nil {
		return nil
	}
	out := new(TeamingStandbyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkHostSwitchProfile) DeepCopyInto(out *UplinkHostSwitchProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkHostSwitchProfile.
func (in *UplinkHostSwitchProfile) DeepCopy() *UplinkHostSwitchProfile {
	if in == nil {
		return nil
	}
	out := new(UplinkHostSwitchProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UplinkHostSwitchProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkHostSwitchProfileList) DeepCopyInto(out *UplinkHostSwitchProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UplinkHostSwitchProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkHostSwitchProfileList.
func (in *UplinkHostSwitchProfileList) DeepCopy() *UplinkHostSwitchProfileList {
	if in == nil {
		return nil
	}
	out := new(UplinkHostSwitchProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UplinkHostSwitchProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkHostSwitchProfileObservation) DeepCopyInto(out *UplinkHostSwitchProfileObservation) {
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
	if in.Lag != nil {
		in, out := &in.Lag, &out.Lag
		*out = make([]LagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.NamedTeaming != nil {
		in, out := &in.NamedTeaming, &out.NamedTeaming
		*out = make([]NamedTeamingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NsxID != nil {
		in, out := &in.NsxID, &out.NsxID
		*out = new(string)
		**out = **in
	}
	if in.OverlayEncap != nil {
		in, out := &in.OverlayEncap, &out.OverlayEncap
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.RealizedID != nil {
		in, out := &in.RealizedID, &out.RealizedID
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
	if in.Teaming != nil {
		in, out := &in.Teaming, &out.Teaming
		*out = make([]TeamingObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TransportVlan != nil {
		in, out := &in.TransportVlan, &out.TransportVlan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkHostSwitchProfileObservation.
func (in *UplinkHostSwitchProfileObservation) DeepCopy() *UplinkHostSwitchProfileObservation {
	if in == nil {
		return nil
	}
	out := new(UplinkHostSwitchProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkHostSwitchProfileParameters) DeepCopyInto(out *UplinkHostSwitchProfileParameters) {
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
	if in.Lag != nil {
		in, out := &in.Lag, &out.Lag
		*out = make([]LagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.NamedTeaming != nil {
		in, out := &in.NamedTeaming, &out.NamedTeaming
		*out = make([]NamedTeamingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NsxID != nil {
		in, out := &in.NsxID, &out.NsxID
		*out = new(string)
		**out = **in
	}
	if in.OverlayEncap != nil {
		in, out := &in.OverlayEncap, &out.OverlayEncap
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
	if in.Teaming != nil {
		in, out := &in.Teaming, &out.Teaming
		*out = make([]TeamingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TransportVlan != nil {
		in, out := &in.TransportVlan, &out.TransportVlan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkHostSwitchProfileParameters.
func (in *UplinkHostSwitchProfileParameters) DeepCopy() *UplinkHostSwitchProfileParameters {
	if in == nil {
		return nil
	}
	out := new(UplinkHostSwitchProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkHostSwitchProfileSpec) DeepCopyInto(out *UplinkHostSwitchProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkHostSwitchProfileSpec.
func (in *UplinkHostSwitchProfileSpec) DeepCopy() *UplinkHostSwitchProfileSpec {
	if in == nil {
		return nil
	}
	out := new(UplinkHostSwitchProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkHostSwitchProfileStatus) DeepCopyInto(out *UplinkHostSwitchProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkHostSwitchProfileStatus.
func (in *UplinkHostSwitchProfileStatus) DeepCopy() *UplinkHostSwitchProfileStatus {
	if in == nil {
		return nil
	}
	out := new(UplinkHostSwitchProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkObservation) DeepCopyInto(out *UplinkObservation) {
	*out = *in
	if in.UplinkName != nil {
		in, out := &in.UplinkName, &out.UplinkName
		*out = new(string)
		**out = **in
	}
	if in.UplinkType != nil {
		in, out := &in.UplinkType, &out.UplinkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkObservation.
func (in *UplinkObservation) DeepCopy() *UplinkObservation {
	if in == nil {
		return nil
	}
	out := new(UplinkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkParameters) DeepCopyInto(out *UplinkParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkParameters.
func (in *UplinkParameters) DeepCopy() *UplinkParameters {
	if in == nil {
		return nil
	}
	out := new(UplinkParameters)
	in.DeepCopyInto(out)
	return out
}
