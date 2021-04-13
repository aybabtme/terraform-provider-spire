package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	spiffeIDTrustDomainAttr = "trust_domain"
	spiffeIDPathAttr        = "path"
	spiffeIDPathAttr        = "spiffe_id"

	selectorTypeAttr  = "type"
	selectorValueAttr = "value"

	regEntrySelectorsAttr     = "selectors"
	regEntryParentIDAttr      = "parent_id"
	regEntrySpiffeIDAttr      = "spiffe_id"
	regEntryTTLAttr           = "ttl"
	regEntryFederatesWithAttr = "federates_with"
	regEntryID                = "entry_id"
	regEntryAdmin             = "admin"
	regEntryDownstream        = "downstream"
	regEntryEntryExpiry       = "entry_expiry"
	regEntryDNSNames          = "dns_names"
	regEntryRevisionNumber    = "revision_number"
)

func dataSourceSpiffeID() *schema.Resource {
	return &schema.Resource{
		Description: "SPIFFE ID made of a trust domain and a path",
		Schema: map[string]*schema.Schema{
			spiffeIDTrustDomainAttr: {
				Description: "Trust domain of the SPIFFE ID.",
				Required:    true,
				Type:        schema.TypeString,
			},
			spiffeIDPathAttr: {
				Description: "Path in the trust domain for the SPIFFE ID.",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceSelector() *schema.Resource {
	return &schema.Resource{
		Description: "A type which describes the conditions under which a registration entry is matched.",
		Schema: map[string]*schema.Schema{
			selectorTypeAttr: {
				Description: "A selector type represents the type of attestation used in attesting the entity (Eg: AWS, K8).",
				Type:        schema.TypeString,
				Required:    true,
			},
			selectorValueAttr: {
				Description: "The value to be attested.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

// https://github.com/spiffe/spire/blob/48d560617bb28b9cd3fa958b3deb3505a78b9983/proto/spire/common/common.proto#L58
func resourceRegistrationEntry() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "This is a curated record that the Server uses to set up and manage the various registered nodes and workloads that are controlled by it.",

		CreateContext: resourceRegistrationEntryCreate,
		ReadContext:   resourceRegistrationEntryRead,
		UpdateContext: resourceRegistrationEntryUpdate,
		DeleteContext: resourceRegistrationEntryDelete,

		Schema: map[string]*schema.Schema{
			regEntrySelectorsAttr: {
				Description: "A list of selectors.",
				Type:        schema.TypeList,
				Elem:        resourceSelector(),
			},
			regEntryParentIDAttr: {
				Description: "The SPIFFE ID of an entity that is authorized to attest the validity of a selector",
				Type:        schema.TypeString,
			},
			regEntrySpiffeIDAttr: {
				Description: "The SPIFFE ID is a structured string used to identify a resource caller. It is defined as a URI comprising a “trust domain” and an associated path",
				Type:        schema.TypeString,
			},
			regEntryTTLAttr: {
				Description: "Time to live.",
				Type:        schema.TypeInt,
			},
			regEntryFederatesWithAttr: {
				Description: "A list of federated trust domain SPIFFE IDs.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			regEntryID: {
				Description: "Entry ID",
				Type:        schema.TypeString,
			},
			regEntryAdmin: {
				Description: "Whether or not the workload is an admin workload. Admin workloads can use their SVID's to authenticate with the Server APIs, for example.",
				Type:        schema.TypeBool,
			},
			regEntryDownstream: {
				Description: "To enable signing CA CSR in upstream spire server.",
				Type:        schema.TypeBool,
			},
			regEntryEntryExpiry: {
				Description: "Expiration of this entry, in seconds from epoch",
				Type:        schema.TypeInt,
			},
			regEntryDNSNames: {
				Description: "DNS entries",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			regEntryRevisionNumber: {
				Description: "Revision number is bumped every time the entry is updated",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func resourceRegistrationEntryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}

func resourceRegistrationEntryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceRegistrationEntryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceRegistrationEntryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
