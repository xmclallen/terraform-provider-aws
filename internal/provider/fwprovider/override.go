// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fwprovider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

var (
	overrideProviderBlock = schema.ListNestedBlock{
		Validators: []validator.List{
			listvalidator.SizeAtMost(1),
		},
		NestedObject: schema.NestedBlockObject{
			Attributes: map[string]schema.Attribute{
				names.AttrRegion: schema.StringAttribute{
					Optional: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					},
					Validators: []validator.String{
						// TODO: Valid for partition
					},
					Description: `Per-resource Region override.`,
				},
			},
			Blocks: map[string]schema.Block{
				"default_tags": schema.ListNestedBlock{ // TODO: Only for aws_s3_object?
					Validators: []validator.List{
						listvalidator.SizeAtMost(1),
					},
					NestedObject: schema.NestedBlockObject{
						Attributes: map[string]schema.Attribute{
							"tags": schema.MapAttribute{
								ElementType: types.StringType,
								Optional:    true,
								Description: `Resource tags.`,
							},
						},
					},
					Description: `Per-resource default tags override.`,
				},
			},
		},
		Description: `Configuration block to override the provider configuration.`,
	}
)
