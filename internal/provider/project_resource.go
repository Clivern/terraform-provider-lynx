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
var _ resource.Resource = &ProjectResource{}
var _ resource.ResourceWithImportState = &ProjectResource{}

func NewProjectResource() resource.Resource {
	return &ProjectResource{}
}

// ProjectResource defines the resource implementation.
type ProjectResource struct {
	client *sdk.Client
}

// ProjectResourceModel describes the resource data model.
type ProjectResourceModel struct {
	ID          types.String            `tfsdk:"id"`
	Name        types.String            `tfsdk:"name"`
	Slug        types.String            `tfsdk:"slug"`
	Description types.String            `tfsdk:"description"`
	Team        *TeamResourceSmallModel `tfsdk:"team"`
}

// TeamResourceSmallModel describes the nested team resource data model.
type TeamResourceSmallModel struct {
	ID types.String `tfsdk:"id"`
}

func (r *ProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Project resource",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Project's name",
				Required:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "Project's slug",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Project's description",
				Required:            true,
			},
			"team": schema.SingleNestedAttribute{
				MarkdownDescription: "Project's team",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						MarkdownDescription: "Team identifier",
						Required:            true,
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Project identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newProject := sdk.Project{
		Name:        data.Name.ValueString(),
		Slug:        data.Slug.ValueString(),
		Description: data.Description.ValueString(),
		Team: sdk.Team{
			ID: data.Team.ID.ValueString(),
		},
	}

	createdProject, err := r.client.CreateProject(newProject)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to create project, got error: %s", err.Error()),
		)
		return
	}

	// Set the created project's ID in the Terraform state
	data.ID = types.StringValue(createdProject.ID)
	data.Name = types.StringValue(createdProject.Name)
	data.Slug = types.StringValue(createdProject.Slug)
	data.Description = types.StringValue(createdProject.Description)
	data.Team.ID = types.StringValue(createdProject.Team.ID)

	// Write logs using the tflog package
	tflog.Trace(ctx, "created a project")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve the project using the GetProject method
	project, err := r.client.GetProject(data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read project, got error: %s", err.Error()),
		)
		return
	}

	// Update the data model with the retrieved project information
	data.Name = types.StringValue(project.Name)
	data.Slug = types.StringValue(project.Slug)
	data.Description = types.StringValue(project.Description)

	// Update the nested team data
	data.Team = &TeamResourceSmallModel{
		ID: types.StringValue(project.Team.ID),
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update the project using the UpdateProject method
	updatedProject := sdk.Project{
		ID:          data.ID.ValueString(),
		Name:        data.Name.ValueString(),
		Slug:        data.Slug.ValueString(),
		Description: data.Description.ValueString(),
		Team: sdk.Team{
			ID: data.Team.ID.ValueString(),
		},
	}

	_, err := r.client.UpdateProject(updatedProject)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to update project, got error: %s", err.Error()),
		)
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteProject(data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to delete project, got error: %s", err.Error()),
		)
		return
	}
}

func (r *ProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
