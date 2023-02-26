// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package provider

import (
	"context"
	"os"

	"github.com/clivern/terraform-provider-lynx/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure lynxProvider satisfies various provider interfaces.
var _ provider.Provider = &lynxProvider{}
var _ provider.ProviderWithFunctions = &lynxProvider{}

// lynxProvider defines the provider implementation.
type lynxProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// LynxProviderModel describes the provider data model.
type LynxProviderModel struct {
	ApiURL types.String `tfsdk:"api_url"`
	ApiKey types.String `tfsdk:"api_key"`
}

func (p *lynxProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "lynx"
	resp.Version = p.version
}

func (p *lynxProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_url": schema.StringAttribute{
				MarkdownDescription: "Lynx API URL",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Lynx API Key",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *lynxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data LynxProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.ApiURL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_url"),
			"Unknown Lynx API URL",
			"The provider cannot create the Lynx API client as there is an unknown configuration value for the Lynx API URL. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the LYNX_API_URL environment variable.",
		)
	}

	if data.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Lynx API Key",
			"The provider cannot create the Lynx API client as there is an unknown configuration value for the Lynx API Key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the LYNX_API_KEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	api_url := os.Getenv("LYNX_API_URL")
	api_key := os.Getenv("LYNX_API_KEY")

	if !data.ApiURL.IsNull() {
		api_url = data.ApiURL.ValueString()
	}

	if !data.ApiKey.IsNull() {
		api_key = data.ApiKey.ValueString()
	}

	client := sdk.NewClient(api_url, api_key)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *lynxProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource,
		NewTeamResource,
		NewSnapshotResource,
		NewProjectResource,
		NewEnvironmentResource,
	}
}

func (p *lynxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *lynxProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &lynxProvider{
			version: version,
		}
	}
}
