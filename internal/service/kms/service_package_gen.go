// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package kms

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	kms_sdkv2 "github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAlias,
			TypeName: "aws_kms_alias",
			Name:     "Alias",
		},
		{
			Factory:  dataSourceCiphertext,
			TypeName: "aws_kms_ciphertext",
			Name:     "Ciphertext",
		},
		{
			Factory:  dataSourceCustomKeyStore,
			TypeName: "aws_kms_custom_key_store",
			Name:     "Custom Key Store",
		},
		{
			Factory:  dataSourceKey,
			TypeName: "aws_kms_key",
			Name:     "Key",
		},
		{
			Factory:  dataSourcePublicKey,
			TypeName: "aws_kms_public_key",
			Name:     "Public Key",
		},
		{
			Factory:  dataSourceSecret,
			TypeName: "aws_kms_secret",
			Name:     "Secret",
		},
		{
			Factory:  dataSourceSecrets,
			TypeName: "aws_kms_secrets",
			Name:     "Secrets",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAlias,
			TypeName: "aws_kms_alias",
			Name:     "Alias",
		},
		{
			Factory:  resourceCiphertext,
			TypeName: "aws_kms_ciphertext",
			Name:     "Ciphertext",
		},
		{
			Factory:  resourceCustomKeyStore,
			TypeName: "aws_kms_custom_key_store",
			Name:     "Custom Key Store",
		},
		{
			Factory:  resourceExternalKey,
			TypeName: "aws_kms_external_key",
			Name:     "External Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceGrant,
			TypeName: "aws_kms_grant",
			Name:     "Grant",
		},
		{
			Factory:  resourceKey,
			TypeName: "aws_kms_key",
			Name:     "Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceKeyPolicy,
			TypeName: "aws_kms_key_policy",
			Name:     "Key Policy",
		},
		{
			Factory:  resourceReplicaExternalKey,
			TypeName: "aws_kms_replica_external_key",
			Name:     "Replica External Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceReplicaKey,
			TypeName: "aws_kms_replica_key",
			Name:     "Replica Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.KMS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*kms_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return kms_sdkv2.NewFromConfig(cfg,
		kms_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
