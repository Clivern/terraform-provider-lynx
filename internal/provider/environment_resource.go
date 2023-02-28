// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package provider

import (
	"context"
	"fmt"

	"github.com/clivern/terraform-provider-lynx/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EnvironmentResource{}
var _ resource.ResourceWithImportState = &EnvironmentResource{}

func NewEnvironmentResource() resource.Resource {
	return &EnvironmentResource{}
}

// EnvironmentResource defines the resource implementation.
type EnvironmentResource struct {
	client *sdk.Client
}

// EnvironmentResourceModel describes the resource data model.
type EnvironmentResourceModel struct {
	ID       types.String               `tfsdk:"id"`
	Name     types.String               `tfsdk:"name"`
	Slug     types.String               `tfsdk:"slug"`
	Username types.String               `tfsdk:"username"`
	Secret   types.String               `tfsdk:"secret"`
	Project  *ProjectResourceSmallModel `tfsdk:"project"`
}

// ProjectResourceSmallModel describes the nested project resource data model.
type ProjectResourceSmallModel struct {
	ID types.String `tfsdk:"id"`
}

func (r *EnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment"
}

func (r *EnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Environment resource",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Environment's name",
				Required:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "Environment's slug",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Environment's username",
				Required:            true,
				Sensitive:           true,
			},
			"secret": schema.StringAttribute{
				MarkdownDescription: "Environment's secret",
				Required:            true,
				Sensitive:           true,
			},
			"project": schema.SingleNestedAttribute{
				MarkdownDescription: "Environment's project",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Project identifier",
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Environment identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *EnvironmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EnvironmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newEnvironment := sdk.Environment{
		Name:     data.Name.ValueString(),
		Slug:     data.Slug.ValueString(),
		Username: data.Username.ValueString(),
		Secret:   data.Secret.ValueString(),
		Project: sdk.Project{
			ID: data.Project.ID.ValueString(),
		},
	}

	tflog.Info(ctx, fmt.Sprintf("Create an environment with name %s", newEnvironment.Name))

	createdEnvironment, err := r.client.CreateEnvironment(newEnvironment)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to create environment, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Environment with id %s got created", createdEnvironment.ID))

	// Set the created environment's ID in the Terraform state
	data.ID = types.StringValue(createdEnvironment.ID)
	data.Name = types.StringValue(createdEnvironment.Name)
	data.Slug = types.StringValue(createdEnvironment.Slug)
	data.Username = types.StringValue(createdEnvironment.Username)
	data.Secret = types.StringValue(createdEnvironment.Secret)
	data.Project.ID = types.StringValue(createdEnvironment.Project.ID)

	// Write logs using the tflog package
	tflog.Trace(ctx, "created an environment")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data EnvironmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Read an environment with id %s", data.ID.ValueString()))

	// Retrieve the environment using the GetEnvironment method
	environment, err := r.client.GetEnvironment(data.Project.ID.ValueString(), data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read environment, got error: %s", err.Error()),
		)
		return
	}

	// Update the data model with the retrieved environment information
	data.Name = types.StringValue(environment.Name)
	data.Slug = types.StringValue(environment.Slug)
	data.Username = types.StringValue(environment.Username)
	data.Secret = types.StringValue(environment.Secret)

	data.Project = &ProjectResourceSmallModel{
		ID: types.StringValue(environment.Project.ID),
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data EnvironmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update the environment using the UpdateEnvironment method
	updatedEnvironment := sdk.Environment{
		ID:       data.ID.ValueString(),
		Name:     data.Name.ValueString(),
		Slug:     data.Slug.ValueString(),
		Username: data.Username.ValueString(),
		Secret:   data.Secret.ValueString(),
		Project: sdk.Project{
			ID: data.Project.ID.ValueString(),
		},
	}

	tflog.Info(ctx, fmt.Sprintf("Update an environment with id %s", updatedEnvironment.ID))

	_, err := r.client.UpdateEnvironment(updatedEnvironment)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to update environment, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Environment with id %s got updated", updatedEnvironment.ID))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data EnvironmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Delete an environment with id %s", data.ID.ValueString()))

	err := r.client.DeleteEnvironment(data.Project.ID.ValueString(), data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to delete environment, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Environment with id %s got deleted", data.ID.ValueString()))
}

func (r *EnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
