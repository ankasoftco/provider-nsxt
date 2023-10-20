package loadbalancer

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_lb_pool", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "PolicyLbPool"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_lb_service", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "PolicyLbService"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_lb_virtual_server", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.Kind = "PolicyLbVirtualServer"
		r.Version = "v1alpha1"
	})
	
}
