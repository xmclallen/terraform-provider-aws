// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
)

type WithOverrideProviderModel struct {
	OverrideProvider fwtypes.ListNestedObjectValueOf[OverrideProviderModel] `tfsdk:"override_provider"`
}

type OverrideProviderModel struct {
	Region types.String `tfsdk:"region"`
}
