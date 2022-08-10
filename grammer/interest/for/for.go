package main

import (
	"github.com/oam-dev/kubevela/pkg/controller/core.oam.dev/v1alpha2/application"
	"github.com/oam-dev/kubevela/pkg/controller/core.oam.dev/v1alpha2/core/components/componentdefinition"
	"github.com/oam-dev/kubevela/pkg/controller/core.oam.dev/v1alpha2/core/policies/policydefinition"
	"github.com/oam-dev/kubevela/pkg/controller/core.oam.dev/v1alpha2/core/traits/traitdefinition"
	"github.com/oam-dev/kubevela/pkg/controller/core.oam.dev/v1alpha2/core/workflow/workflowstepdefinition"
)

func main() {
	for _, setup := range []func(ctrl.Manager, controller.Args) error{
		application.Setup, traitdefinition.Setup, componentdefinition.Setup, policydefinition.Setup, workflowstepdefinition.Setup,
	} {
		if err := setup(mgr, args); err != nil {
			return err
		}
	}
}
