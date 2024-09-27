// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conns

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// ServicePackage is the minimal interface exported from each AWS service package.
// Its methods return the Plugin SDK and Framework resources and data sources implemented in the package.
type ServicePackage interface {
	FrameworkDataSources(context.Context) []*types.ServicePackageFrameworkDataSource
	FrameworkResources(context.Context) []*types.ServicePackageFrameworkResource
	SDKDataSources(context.Context) []*types.ServicePackageSDKDataSource
	SDKResources(context.Context) []*types.ServicePackageSDKResource
	ServicePackageName() string
}

type (
	contextKeyType int
)

var (
	contextKey contextKeyType
)

// InContext represents the resource information kept in Context.
type InContext struct {
	isDataSource       bool   // Data source?
	resourceName       string // Friendly resource name, e.g. "Subnet"
	servicePackageName string // Canonical name defined as a constant in names package
}

// IsDataSource returns true if the resource is a data source.
func (c *InContext) IsDataSource() bool {
	return c.isDataSource
}

// ResourceName returns the friendly resource name, e.g. "Subnet".
func (c *InContext) ResourceName() string {
	return c.resourceName
}

// ServicePackageName returns the canonical service name defined as a constant in the `names` package.
func (c *InContext) ServicePackageName() string {
	return c.servicePackageName
}

func NewDataSourceContext(ctx context.Context, servicePackageName, resourceName string) context.Context {
	v := InContext{
		isDataSource:       true,
		resourceName:       resourceName,
		servicePackageName: servicePackageName,
	}

	return context.WithValue(ctx, contextKey, &v)
}

func NewResourceContext(ctx context.Context, servicePackageName, resourceName string) context.Context {
	v := InContext{
		resourceName:       resourceName,
		servicePackageName: servicePackageName,
	}

	return context.WithValue(ctx, contextKey, &v)
}

func FromContext(ctx context.Context) (*InContext, bool) {
	v, ok := ctx.Value(contextKey).(*InContext)
	return v, ok
}
