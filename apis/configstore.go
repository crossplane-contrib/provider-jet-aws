package apis

import (
	"github.com/crossplane-contrib/terrajet/pkg/terraform/resource"

	iamv1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1"
	rdsv1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/rds/v1alpha1"
	vpcv1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1"
)

func init() {
	AddToConfigStores = append(AddToConfigStores,
		rdsv1alpha1.ConfigStoreBuilder.AddToConfigStore,
		vpcv1alpha1.ConfigStoreBuilder.AddToConfigStore,
		iamv1alpha1.ConfigStoreBuilder.AddToConfigStore,
	)
}

var AddToConfigStores resource.ConfigStoreBuilder

func AddToConfigStore(c *resource.ConfigStore) error {
	return AddToConfigStores.AddToConfigStore(c)
}
