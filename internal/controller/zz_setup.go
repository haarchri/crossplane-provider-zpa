/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/zscaler/provider-zpa/internal/controller/providerconfig"
	accessrule "github.com/zscaler/provider-zpa/internal/controller/zpaaccessrule/accessrule"
	server "github.com/zscaler/provider-zpa/internal/controller/zpaapplicationserver/server"
	forwardingrule "github.com/zscaler/provider-zpa/internal/controller/zpaforwardingrule/forwardingrule"
	inspectionrule "github.com/zscaler/provider-zpa/internal/controller/zpainspectionrule/inspectionrule"
	group "github.com/zscaler/provider-zpa/internal/controller/zpasegmentgroup/group"
	timeoutrule "github.com/zscaler/provider-zpa/internal/controller/zpatimeoutrule/timeoutrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		accessrule.Setup,
		server.Setup,
		forwardingrule.Setup,
		inspectionrule.Setup,
		group.Setup,
		timeoutrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
