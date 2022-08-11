package role

import (
	"github.com/crossplane/terrajet/pkg/config"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

var shortGroup = "role"

// Configure configures the null group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("boundary_role", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Role"
		r.References["scope_id"] = config.Reference{
			Type: "github.com/maxnovawind/provider-jet-boundary/apis/scope/v1alpha1.Scope",
		}
		r.References["grant_scope_id"] = config.Reference{
			Type: "github.com/maxnovawind/provider-jet-boundary/apis/scope/v1alpha1.Scope",
		}

	})
}
