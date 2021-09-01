package clients

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	xpawsclient "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-tf-aws/apis/v1alpha1"
)

// ProviderConfigBuilder returns provider specific configuration like provider
// credentials used to connect to cloud APIs in the expected form of a Terraform
// provider.
func ProviderConfigBuilder(ctx context.Context, client client.Client, mg resource.Managed) ([]byte, error) { //nolint:gocyclo
	if mg.GetProviderConfigReference() == nil {
		return nil, errors.New("no providerConfigRef provided")
	}
	pc := &v1alpha1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get referenced Provider")
	}

	region, err := getRegion(mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get region")
	}

	t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, "cannot track ProviderConfig usage")
	}

	var cfg *aws.Config
	switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
	case xpv1.CredentialsSourceInjectedIdentity:
		if cfg, err = xpawsclient.UsePodServiceAccount(ctx, []byte{}, xpawsclient.DefaultSection, region); err != nil {
			return nil, errors.Wrap(err, "failed to use pod service account")
		}
	default:
		data, err := resource.CommonCredentialExtractor(ctx, s, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get credentials")
		}
		if cfg, err = xpawsclient.UseProviderSecret(ctx, data, xpawsclient.DefaultSection, region); err != nil {
			return nil, errors.Wrap(err, "failed to use provider secret")
		}
	}
	awsConf := xpawsclient.SetResolver(ctx, mg, cfg)
	creds, err := awsConf.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
	}

	// TODO(hasan): figure out what other values could be possible set here.
	//   e.g. what about setting an assume_role section: https://registry.terraform.io/providers/hashicorp/aws/latest/docs#argument-reference
	tfPc := map[string]string{}
	tfPc["region"] = awsConf.Region
	tfPc["access_key"] = creds.AccessKeyID
	tfPc["secret_key"] = creds.SecretAccessKey
	tfPc["token"] = creds.SessionToken

	return json.Marshal(tfPc)
}

func getRegion(obj runtime.Object) (string, error) {
	fromMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", errors.Wrap(err, "cannot convert to unstructured")
	}
	return fieldpath.Pave(fromMap).GetString("spec.forProvider.region")
}
