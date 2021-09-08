package clients

import (
	"context"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ProviderConfigBuilder builds the provider configuration blob that Terraform
// uses to configure the Terraform Provider.
func ProviderConfigBuilder(ctx context.Context, client client.Client, mg xpresource.Managed) ([]byte, error) {
	// TODO(hasan): Read ProviderConfig here and prepare a Terraform provider
	//  config content with credentials accordingly.

	// TODO(hasan/muvaf) Terraform expects region or project information to be
	//  in provider configuration different than Crossplane. Figure out how to
	//  handle that.
	return nil, nil
}
