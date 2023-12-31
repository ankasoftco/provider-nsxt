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
func (in *BodyConditionObservation) DeepCopyInto(out *BodyConditionObservation) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BodyConditionObservation.
func (in *BodyConditionObservation) DeepCopy() *BodyConditionObservation {
	if in == nil {
		return nil
	}
	out := new(BodyConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BodyConditionParameters) DeepCopyInto(out *BodyConditionParameters) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BodyConditionParameters.
func (in *BodyConditionParameters) DeepCopy() *BodyConditionParameters {
	if in == nil {
		return nil
	}
	out := new(BodyConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CookieConditionObservation) DeepCopyInto(out *CookieConditionObservation) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CookieConditionObservation.
func (in *CookieConditionObservation) DeepCopy() *CookieConditionObservation {
	if in == nil {
		return nil
	}
	out := new(CookieConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CookieConditionParameters) DeepCopyInto(out *CookieConditionParameters) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CookieConditionParameters.
func (in *CookieConditionParameters) DeepCopy() *CookieConditionParameters {
	if in == nil {
		return nil
	}
	out := new(CookieConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRedirectActionObservation) DeepCopyInto(out *HTTPRedirectActionObservation) {
	*out = *in
	if in.RedirectStatus != nil {
		in, out := &in.RedirectStatus, &out.RedirectStatus
		*out = new(string)
		**out = **in
	}
	if in.RedirectURL != nil {
		in, out := &in.RedirectURL, &out.RedirectURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirectActionObservation.
func (in *HTTPRedirectActionObservation) DeepCopy() *HTTPRedirectActionObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPRedirectActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRedirectActionParameters) DeepCopyInto(out *HTTPRedirectActionParameters) {
	*out = *in
	if in.RedirectStatus != nil {
		in, out := &in.RedirectStatus, &out.RedirectStatus
		*out = new(string)
		**out = **in
	}
	if in.RedirectURL != nil {
		in, out := &in.RedirectURL, &out.RedirectURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirectActionParameters.
func (in *HTTPRedirectActionParameters) DeepCopy() *HTTPRedirectActionParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPRedirectActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRejectActionObservation) DeepCopyInto(out *HTTPRejectActionObservation) {
	*out = *in
	if in.ReplyMessage != nil {
		in, out := &in.ReplyMessage, &out.ReplyMessage
		*out = new(string)
		**out = **in
	}
	if in.ReplyStatus != nil {
		in, out := &in.ReplyStatus, &out.ReplyStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRejectActionObservation.
func (in *HTTPRejectActionObservation) DeepCopy() *HTTPRejectActionObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPRejectActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRejectActionParameters) DeepCopyInto(out *HTTPRejectActionParameters) {
	*out = *in
	if in.ReplyMessage != nil {
		in, out := &in.ReplyMessage, &out.ReplyMessage
		*out = new(string)
		**out = **in
	}
	if in.ReplyStatus != nil {
		in, out := &in.ReplyStatus, &out.ReplyStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRejectActionParameters.
func (in *HTTPRejectActionParameters) DeepCopy() *HTTPRejectActionParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPRejectActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderConditionObservation) DeepCopyInto(out *HeaderConditionObservation) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderConditionObservation.
func (in *HeaderConditionObservation) DeepCopy() *HeaderConditionObservation {
	if in == nil {
		return nil
	}
	out := new(HeaderConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderConditionParameters) DeepCopyInto(out *HeaderConditionParameters) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderConditionParameters.
func (in *HeaderConditionParameters) DeepCopy() *HeaderConditionParameters {
	if in == nil {
		return nil
	}
	out := new(HeaderConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPConditionObservation) DeepCopyInto(out *IPConditionObservation) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.SourceAddress != nil {
		in, out := &in.SourceAddress, &out.SourceAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPConditionObservation.
func (in *IPConditionObservation) DeepCopy() *IPConditionObservation {
	if in == nil {
		return nil
	}
	out := new(IPConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPConditionParameters) DeepCopyInto(out *IPConditionParameters) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.SourceAddress != nil {
		in, out := &in.SourceAddress, &out.SourceAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPConditionParameters.
func (in *IPConditionParameters) DeepCopy() *IPConditionParameters {
	if in == nil {
		return nil
	}
	out := new(IPConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpForwardingRule) DeepCopyInto(out *LbHttpForwardingRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpForwardingRule.
func (in *LbHttpForwardingRule) DeepCopy() *LbHttpForwardingRule {
	if in == nil {
		return nil
	}
	out := new(LbHttpForwardingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbHttpForwardingRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpForwardingRuleList) DeepCopyInto(out *LbHttpForwardingRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbHttpForwardingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpForwardingRuleList.
func (in *LbHttpForwardingRuleList) DeepCopy() *LbHttpForwardingRuleList {
	if in == nil {
		return nil
	}
	out := new(LbHttpForwardingRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbHttpForwardingRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpForwardingRuleObservation) DeepCopyInto(out *LbHttpForwardingRuleObservation) {
	*out = *in
	if in.BodyCondition != nil {
		in, out := &in.BodyCondition, &out.BodyCondition
		*out = make([]BodyConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CookieCondition != nil {
		in, out := &in.CookieCondition, &out.CookieCondition
		*out = make([]CookieConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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
	if in.HTTPRedirectAction != nil {
		in, out := &in.HTTPRedirectAction, &out.HTTPRedirectAction
		*out = make([]HTTPRedirectActionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPRejectAction != nil {
		in, out := &in.HTTPRejectAction, &out.HTTPRejectAction
		*out = make([]HTTPRejectActionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HeaderCondition != nil {
		in, out := &in.HeaderCondition, &out.HeaderCondition
		*out = make([]HeaderConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPCondition != nil {
		in, out := &in.IPCondition, &out.IPCondition
		*out = make([]IPConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchStrategy != nil {
		in, out := &in.MatchStrategy, &out.MatchStrategy
		*out = new(string)
		**out = **in
	}
	if in.MethodCondition != nil {
		in, out := &in.MethodCondition, &out.MethodCondition
		*out = make([]MethodConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.SelectPoolAction != nil {
		in, out := &in.SelectPoolAction, &out.SelectPoolAction
		*out = make([]SelectPoolActionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TCPCondition != nil {
		in, out := &in.TCPCondition, &out.TCPCondition
		*out = make([]TCPConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URICondition != nil {
		in, out := &in.URICondition, &out.URICondition
		*out = make([]URIConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VersionCondition != nil {
		in, out := &in.VersionCondition, &out.VersionCondition
		*out = make([]VersionConditionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpForwardingRuleObservation.
func (in *LbHttpForwardingRuleObservation) DeepCopy() *LbHttpForwardingRuleObservation {
	if in == nil {
		return nil
	}
	out := new(LbHttpForwardingRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpForwardingRuleParameters) DeepCopyInto(out *LbHttpForwardingRuleParameters) {
	*out = *in
	if in.BodyCondition != nil {
		in, out := &in.BodyCondition, &out.BodyCondition
		*out = make([]BodyConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CookieCondition != nil {
		in, out := &in.CookieCondition, &out.CookieCondition
		*out = make([]CookieConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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
	if in.HTTPRedirectAction != nil {
		in, out := &in.HTTPRedirectAction, &out.HTTPRedirectAction
		*out = make([]HTTPRedirectActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPRejectAction != nil {
		in, out := &in.HTTPRejectAction, &out.HTTPRejectAction
		*out = make([]HTTPRejectActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HeaderCondition != nil {
		in, out := &in.HeaderCondition, &out.HeaderCondition
		*out = make([]HeaderConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPCondition != nil {
		in, out := &in.IPCondition, &out.IPCondition
		*out = make([]IPConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchStrategy != nil {
		in, out := &in.MatchStrategy, &out.MatchStrategy
		*out = new(string)
		**out = **in
	}
	if in.MethodCondition != nil {
		in, out := &in.MethodCondition, &out.MethodCondition
		*out = make([]MethodConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelectPoolAction != nil {
		in, out := &in.SelectPoolAction, &out.SelectPoolAction
		*out = make([]SelectPoolActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TCPCondition != nil {
		in, out := &in.TCPCondition, &out.TCPCondition
		*out = make([]TCPConditionParameters, len(*in))
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
	if in.URICondition != nil {
		in, out := &in.URICondition, &out.URICondition
		*out = make([]URIConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VersionCondition != nil {
		in, out := &in.VersionCondition, &out.VersionCondition
		*out = make([]VersionConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpForwardingRuleParameters.
func (in *LbHttpForwardingRuleParameters) DeepCopy() *LbHttpForwardingRuleParameters {
	if in == nil {
		return nil
	}
	out := new(LbHttpForwardingRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpForwardingRuleSpec) DeepCopyInto(out *LbHttpForwardingRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpForwardingRuleSpec.
func (in *LbHttpForwardingRuleSpec) DeepCopy() *LbHttpForwardingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(LbHttpForwardingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpForwardingRuleStatus) DeepCopyInto(out *LbHttpForwardingRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpForwardingRuleStatus.
func (in *LbHttpForwardingRuleStatus) DeepCopy() *LbHttpForwardingRuleStatus {
	if in == nil {
		return nil
	}
	out := new(LbHttpForwardingRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MethodConditionObservation) DeepCopyInto(out *MethodConditionObservation) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MethodConditionObservation.
func (in *MethodConditionObservation) DeepCopy() *MethodConditionObservation {
	if in == nil {
		return nil
	}
	out := new(MethodConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MethodConditionParameters) DeepCopyInto(out *MethodConditionParameters) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MethodConditionParameters.
func (in *MethodConditionParameters) DeepCopy() *MethodConditionParameters {
	if in == nil {
		return nil
	}
	out := new(MethodConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectPoolActionObservation) DeepCopyInto(out *SelectPoolActionObservation) {
	*out = *in
	if in.PoolID != nil {
		in, out := &in.PoolID, &out.PoolID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectPoolActionObservation.
func (in *SelectPoolActionObservation) DeepCopy() *SelectPoolActionObservation {
	if in == nil {
		return nil
	}
	out := new(SelectPoolActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectPoolActionParameters) DeepCopyInto(out *SelectPoolActionParameters) {
	*out = *in
	if in.PoolID != nil {
		in, out := &in.PoolID, &out.PoolID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectPoolActionParameters.
func (in *SelectPoolActionParameters) DeepCopy() *SelectPoolActionParameters {
	if in == nil {
		return nil
	}
	out := new(SelectPoolActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPConditionObservation) DeepCopyInto(out *TCPConditionObservation) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.SourcePort != nil {
		in, out := &in.SourcePort, &out.SourcePort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPConditionObservation.
func (in *TCPConditionObservation) DeepCopy() *TCPConditionObservation {
	if in == nil {
		return nil
	}
	out := new(TCPConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPConditionParameters) DeepCopyInto(out *TCPConditionParameters) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.SourcePort != nil {
		in, out := &in.SourcePort, &out.SourcePort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPConditionParameters.
func (in *TCPConditionParameters) DeepCopy() *TCPConditionParameters {
	if in == nil {
		return nil
	}
	out := new(TCPConditionParameters)
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
func (in *URIConditionObservation) DeepCopyInto(out *URIConditionObservation) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URIConditionObservation.
func (in *URIConditionObservation) DeepCopy() *URIConditionObservation {
	if in == nil {
		return nil
	}
	out := new(URIConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URIConditionParameters) DeepCopyInto(out *URIConditionParameters) {
	*out = *in
	if in.CaseSensitive != nil {
		in, out := &in.CaseSensitive, &out.CaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.MatchType != nil {
		in, out := &in.MatchType, &out.MatchType
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URIConditionParameters.
func (in *URIConditionParameters) DeepCopy() *URIConditionParameters {
	if in == nil {
		return nil
	}
	out := new(URIConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionConditionObservation) DeepCopyInto(out *VersionConditionObservation) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionConditionObservation.
func (in *VersionConditionObservation) DeepCopy() *VersionConditionObservation {
	if in == nil {
		return nil
	}
	out := new(VersionConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionConditionParameters) DeepCopyInto(out *VersionConditionParameters) {
	*out = *in
	if in.Inverse != nil {
		in, out := &in.Inverse, &out.Inverse
		*out = new(bool)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionConditionParameters.
func (in *VersionConditionParameters) DeepCopy() *VersionConditionParameters {
	if in == nil {
		return nil
	}
	out := new(VersionConditionParameters)
	in.DeepCopyInto(out)
	return out
}
