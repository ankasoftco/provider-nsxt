package firewall

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_context_profile", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyContextProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_context_profile_custom_attribute", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyContextProfileCustomAttribute"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyGatewayPolicy"
		r.Version = "v1alpha1"
	})
	
	p.AddResourceConfigurator("nsxt_policy_intrusion_service_policy", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyInstrusionServicePolicy"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_intrusion_service_profile", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyInstrusionServiceProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_predefined_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyPredefinedGatewayPolicy"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_predefined_security_policy", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyPredefinedSecurityPolicy"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_security_policy", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicySecurityPolicy"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_service", func(r *config.Resource) {
		r.ShortGroup = "firewall"
		r.Kind = "PolicyService"
		r.Version = "v1alpha1"
	})
}
