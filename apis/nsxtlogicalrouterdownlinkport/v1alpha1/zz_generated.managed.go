/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicalRouterDownlinkPort.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicalRouterDownlinkPort) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicalRouterDownlinkPort.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicalRouterDownlinkPort) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LogicalRouterDownlinkPort.
func (mg *LogicalRouterDownlinkPort) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}