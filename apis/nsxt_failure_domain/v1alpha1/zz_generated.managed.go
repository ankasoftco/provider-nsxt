/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this FailureDomain.
func (mg *FailureDomain) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FailureDomain.
func (mg *FailureDomain) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this FailureDomain.
func (mg *FailureDomain) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this FailureDomain.
func (mg *FailureDomain) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FailureDomain.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FailureDomain) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FailureDomain.
func (mg *FailureDomain) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FailureDomain.
func (mg *FailureDomain) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FailureDomain.
func (mg *FailureDomain) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FailureDomain.
func (mg *FailureDomain) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this FailureDomain.
func (mg *FailureDomain) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this FailureDomain.
func (mg *FailureDomain) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FailureDomain.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FailureDomain) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FailureDomain.
func (mg *FailureDomain) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FailureDomain.
func (mg *FailureDomain) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
