/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PolicyDhcpRelay.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PolicyDhcpRelay) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PolicyDhcpRelay.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PolicyDhcpRelay) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PolicyDhcpRelay.
func (mg *PolicyDhcpRelay) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PolicyDhcpServer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PolicyDhcpServer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PolicyDhcpServer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PolicyDhcpServer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PolicyDhcpServer.
func (mg *PolicyDhcpServer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PolicyDhcpV4StaticBinding.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PolicyDhcpV4StaticBinding) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PolicyDhcpV4StaticBinding.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PolicyDhcpV4StaticBinding) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PolicyDhcpV4StaticBinding.
func (mg *PolicyDhcpV4StaticBinding) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PolicyDhcpV6StaticBinding.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PolicyDhcpV6StaticBinding) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PolicyDhcpV6StaticBinding.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PolicyDhcpV6StaticBinding) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PolicyDhcpV6StaticBinding.
func (mg *PolicyDhcpV6StaticBinding) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
