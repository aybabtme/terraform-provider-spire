package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/proto/spire/types"
)

const (
	spiffeIDTrustDomainAttr = "trust_domain"
	spiffeIDSegmentsAttr    = "segments"
	spiffeIDAttr            = "spiffe_id"
)

func dataSourceSpiffeID() *schema.Resource {
	return &schema.Resource{
		Description: "SPIFFE ID made of a trust domain and a path",
		ReadContext: dataSourceSpiffeIDRead,
		Schema: map[string]*schema.Schema{
			spiffeIDTrustDomainAttr: {
				Description: "Trust domain of the SPIFFE ID.",
				Required:    true,
				Type:        schema.TypeString,
			},
			spiffeIDSegmentsAttr: {
				Description: "Segments of the ID that form the path.",
				Required:    true,
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			spiffeIDAttr: {
				Description: "SPIFFE ID that results from the trust domain and segments.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceSpiffeIDRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	trustDomainRaw := d.Get(spiffeIDTrustDomainAttr).(string)
	path := d.Get(spiffeIDPathAttr).(string)

	trustDomain, err := spiffeid.TrustDomainFromString(trustDomainRaw)
	if err != nil {
		return diag.FromErr(err)
	}

	id := &types.SPIFFEID{
		TrustDomain: trustDomain.String(),
		Path:        path,
	}
	id.GetPath()

	return diag.Errorf("not implemented")
}
