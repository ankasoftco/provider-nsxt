/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LogicalRouterCentralizedServicePort.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LogicalRouterCentralizedServicePort) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LogicalRouterCentralizedServicePort.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LogicalRouterCentralizedServicePort) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LogicalRouterCentralizedServicePort.
func (mg *LogicalRouterCentralizedServicePort) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
