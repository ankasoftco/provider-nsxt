/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AlgorithmTypeNsService.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AlgorithmTypeNsService) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AlgorithmTypeNsService.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AlgorithmTypeNsService) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AlgorithmTypeNsService.
func (mg *AlgorithmTypeNsService) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}