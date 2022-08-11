package scope

import (
	"github.com/crossplane/terrajet/pkg/config"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

var shortGroup = "scope"

// Configure configures the null group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("boundary_scope", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Scope"
		r.References["scope_id"] = config.Reference{
			Type: "Scope",
		}
	})
}
