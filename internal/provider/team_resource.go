// Copyright 2024 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package provider

import (
	"context"
	"fmt"
	"strings"

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
var _ resource.Resource = &TeamResource{}
var _ resource.ResourceWithImportState = &TeamResource{}

func NewTeamResource() resource.Resource {
	return &TeamResource{}
}

// TeamResource defines the resource implementation.
type TeamResource struct {
	client *sdk.Client
}

// TeamResourceModel describes the resource data model.
type TeamResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Slug        types.String `tfsdk:"slug"`
	Description types.String `tfsdk:"description"`
	Members     types.List   `tfsdk:"members"`
}

func (r *TeamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (r *TeamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Team resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Team's name",
				Required:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "Team's slug",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Team's description",
				Optional:            true,
			},
			"members": schema.ListAttribute{
				MarkdownDescription: "Team's members",
				ElementType:         types.StringType,
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Team identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *TeamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TeamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TeamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	membersList := data.Members.Elements()

	members := make([]string, 0, len(membersList))

	for i := 0; i < len(membersList); i++ {
		member := membersList[i]

		members = append(members, strings.Trim(member.String(), "\""))
	}

	newTeam := sdk.Team{
		Name:        data.Name.ValueString(),
		Slug:        data.Slug.ValueString(),
		Description: data.Description.ValueString(),
		Members:     members,
	}

	tflog.Info(ctx, fmt.Sprintf("Create a team with name %s", newTeam.Name))

	createdTeam, err := r.client.CreateTeam(newTeam)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to create team, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Team with id %s got created", createdTeam.ID))

	// Set the created team's ID in the Terraform state
	data.ID = types.StringValue(createdTeam.ID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TeamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Read a team with id %s", data.ID.ValueString()))

	// Retrieve the team using the GetTeam method
	team, err := r.client.GetTeam(data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read team, got error: %s", err.Error()),
		)
		return
	}

	members, _ := types.ListValueFrom(context.TODO(), types.StringType, team.Members)

	// Update the data model with the retrieved team information
	data.Name = types.StringValue(team.Name)
	data.Slug = types.StringValue(team.Slug)
	data.Description = types.StringValue(team.Description)
	data.Members = members

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TeamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	membersList := data.Members.Elements()

	members := make([]string, 0, len(membersList))

	for i := 0; i < len(membersList); i++ {
		member := membersList[i]

		members = append(members, member.String())
	}

	// Update the team using the UpdateTeam method
	updatedTeam := sdk.Team{
		ID:          data.ID.ValueString(),
		Name:        data.Name.ValueString(),
		Slug:        data.Slug.ValueString(),
		Description: data.Description.ValueString(),
		Members:     members,
	}

	tflog.Info(ctx, fmt.Sprintf("Update a team with id %s", updatedTeam.ID))

	_, err := r.client.UpdateTeam(updatedTeam)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to update team, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Team with id %s got updated", updatedTeam.ID))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TeamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Delete a team with id %s", data.ID.ValueString()))

	err := r.client.DeleteTeam(data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to delete team, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Team with id %s got deleted", data.ID.ValueString()))
}

func (r *TeamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
