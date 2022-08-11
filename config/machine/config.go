/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package machine

import (
	"github.com/crossplane/terrajet/pkg/config"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

var shortGroup = "machine"

// Configure configures the null group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("boundary_host", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Host"
		r.References["host_catalog_id"] = config.Reference{
			Type: "HostCatalog",
		}
	})
	p.AddResourceConfigurator("boundary_host_set", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "HostSet"
		r.References["host_catalog_id"] = config.Reference{
			Type: "HostCatalog",
		}
		r.References["host_ids"] = config.Reference{
			Type: "Host",
		}
	})
	p.AddResourceConfigurator("boundary_host_catalog", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "HostCatalog"
		r.References["scope_id"] = config.Reference{
			Type: "github.com/maxnovawind/provider-jet-boundary/apis/domain/v1alpha1.Scope",
		}
	})
	p.AddResourceConfigurator("boundary_target", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Target"
		r.References["scope_id"] = config.Reference{
			Type: "github.com/maxnovawind/provider-jet-boundary/apis/domain/v1alpha1.Scope",
		}
		r.References["host_source_ids"] = config.Reference{
			Type: "HostSet",
		}
	})
}
