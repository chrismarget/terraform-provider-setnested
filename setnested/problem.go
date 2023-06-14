package setnested

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = &resourceSetNested{}

type resourceSetNested struct {
}

func (o *resourceSetNested) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_problem"
}

func (o *resourceSetNested) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"sna": schema.SetNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"required_string": schema.StringAttribute{
							Required: true,
						},
						"computed_int": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (o *resourceSetNested) Create(_ context.Context, _ resource.CreateRequest, _ *resource.CreateResponse) {
}

func (o *resourceSetNested) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

func (o *resourceSetNested) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
}

func (o *resourceSetNested) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}
