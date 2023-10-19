/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policydhcprelay "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcprelay"
	policydhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpserver"
	providerconfig "github.com/ankasoftco/provider-nsxt/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policydhcprelay.Setup,
		policydhcpserver.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
