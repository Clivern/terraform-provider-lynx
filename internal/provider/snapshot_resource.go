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
var _ resource.Resource = &SnapshotResource{}
var _ resource.ResourceWithImportState = &SnapshotResource{}

func NewSnapshotResource() resource.Resource {
	return &SnapshotResource{}
}

// SnapshotResource defines the resource implementation.
type SnapshotResource struct {
	client *sdk.Client
}

// SnapshotResourceModel describes the resource data model.
type SnapshotResourceModel struct {
	ID          types.String            `tfsdk:"id"`
	Title       types.String            `tfsdk:"title"`
	Description types.String            `tfsdk:"description"`
	RecordType  types.String            `tfsdk:"record_type"`
	RecordID    types.String            `tfsdk:"record_id"`
	Team        *TeamResourceSmallModel `tfsdk:"team"`
}

func (r *SnapshotResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snapshot"
}

func (r *SnapshotResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Snapshot resource",
		Attributes: map[string]schema.Attribute{
			"title": schema.StringAttribute{
				MarkdownDescription: "Snapshot's title",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Snapshot's description",
				Required:            true,
			},
			"record_type": schema.StringAttribute{
				MarkdownDescription: "Snapshot's record_type",
				Required:            true,
			},
			"record_id": schema.StringAttribute{
				MarkdownDescription: "Snapshot's record_id",
				Required:            true,
			},
			"team": schema.SingleNestedAttribute{
				MarkdownDescription: "Snapshot's team",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Team identifier",
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Snapshot identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *SnapshotResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SnapshotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnapshotResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newSnapshot := sdk.Snapshot{
		Title:       data.Title.ValueString(),
		Description: data.Description.ValueString(),
		RecordType:  data.RecordType.ValueString(),
		RecordID:    data.RecordID.ValueString(),
		Team: sdk.Team{
			ID: data.Team.ID.ValueString(),
		},
	}

	tflog.Info(ctx, fmt.Sprintf("Create a snapshot with title %s", newSnapshot.Title))

	createdSnapshot, err := r.client.CreateSnapshot(newSnapshot)

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to create snapshot, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Snapshot with title %s got created", newSnapshot.Title))

	// Set the created snapshot's ID in the Terraform state
	data.ID = types.StringValue(createdSnapshot.ID)
	data.Title = types.StringValue(createdSnapshot.Title)
	data.RecordType = types.StringValue(createdSnapshot.RecordType)
	data.RecordID = types.StringValue(createdSnapshot.RecordID)
	data.Team.ID = types.StringValue(createdSnapshot.Team.ID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnapshotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnapshotResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Read a snapshot with id %s", data.ID.ValueString()))

	// Retrieve the snapshot using the GetSnapshot method
	snapshot, err := r.client.GetSnapshot(data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read snapshot, got error: %s", err.Error()),
		)
		return
	}

	// Update the data model with the retrieved snapshot information
	data.Title = types.StringValue(snapshot.Title)
	data.RecordType = types.StringValue(snapshot.RecordType)
	data.RecordID = types.StringValue(snapshot.RecordID)
	data.Team.ID = types.StringValue(snapshot.Team.ID)

	// Update the nested team data
	data.Team = &TeamResourceSmallModel{
		ID: types.StringValue(snapshot.Team.ID),
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnapshotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnapshotResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnapshotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnapshotResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Delete a snapshot with id %s", data.ID.ValueString()))

	err := r.client.DeleteSnapshot(data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to delete snapshot, got error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Snapshot with id %s got deleted", data.ID.ValueString()))
}

func (r *SnapshotResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
