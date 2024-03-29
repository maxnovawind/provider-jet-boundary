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

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/maxnovawind/provider-jet-boundary/apis/v1alpha1"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/terrajet/pkg/terraform"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal boundary credentials as JSON"
)

const (
	keyAddr                        = "addr"
	keyAuthMethodID                = "auth_method_id"
	keyPasswordAuthMethodLoginName = "password_auth_method_login_name"
	keyPasswordAuthMethodPassword  = "password_auth_method_password"
	keyTLSInsecureSkipVerify       = "tls_insecure"
)

type boundaryCreds struct {
	Addr                        string `json:"addr"`
	AuthMethodID                string `json:"auth_method_id"`
	PasswordAuthMethodLoginName string `json:"password_auth_method_login_name"`
	PasswordAuthMethodPassword  string `json:"password_auth_method_password"`
	TLSInsecureSkipVerify       bool   `json:"tls_insecure,omitempty"`
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		var keyCreds boundaryCreds
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		if err := json.Unmarshal(data, &keyCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration

		ps.Configuration = map[string]interface{}{}
		ps.Configuration = map[string]interface{}{
			keyAddr:                        keyCreds.Addr,
			keyAuthMethodID:                keyCreds.AuthMethodID,
			keyPasswordAuthMethodLoginName: keyCreds.PasswordAuthMethodLoginName,
			keyPasswordAuthMethodPassword:  keyCreds.PasswordAuthMethodPassword,
			keyTLSInsecureSkipVerify:       keyCreds.TLSInsecureSkipVerify,
		}
		return ps, nil
	}
}
