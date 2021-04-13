package provider

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	spiretoken "github.com/spiffe/spire/proto/spire/api/server/agent/v1"
	spireentry "github.com/spiffe/spire/proto/spire/api/server/entry/v1"
	"google.golang.org/grpc"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

const (
	provCfgAuthKubeExecAttr = "auth_kube_exec"
	// provCfgAuthSSHAttr      = "auth_ssh"
	provCfgAuthX509Attr = "auth_x509"

	provCfgAuthKubeExecNamespaceAttr     = "namespace"
	provCfgAuthKubeExecLabelSelectorAttr = "label_selectors"

	provCfgAuthX509ServerHostAttr = "server_host"
	provCfgAuthX509ServerPortAttr = "server_port"
)

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				provCfgAuthKubeExecAttr: {
					Description:   "Reaches the SPIRE server on its local unix socket via `kubectl exec`.",
					Type:          schema.TypeList,
					MaxItems:      1,
					ConflictsWith: []string{provCfgAuthX509Attr},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							provCfgAuthKubeExecNamespaceAttr: {
								Description: "Host where the SPIRE server can be reached.",
								Required:    true,
							},
							provCfgAuthKubeExecLabelSelectorAttr: {
								Description: "Port on which the SPIRE server can be reached.",
								Required:    true,
							},
						},
					},
				},
				provCfgAuthX509Attr: {
					Description:   "Uses an x509 SVID that has the admin flag enabled to reach the SPIRE server on its remote address.",
					Type:          schema.TypeList,
					MaxItems:      1,
					ConflictsWith: []string{provCfgAuthKubeExecAttr},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							provCfgAuthX509ServerHostAttr: {
								Description: "Host where the SPIRE server can be reached.",
								Required:    true,
							},
							provCfgAuthX509ServerPortAttr: {
								Description: "Port on which the SPIRE server can be reached.",
								Required:    true,
							},
						},
					},
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"spire_spiffe_id": dataSourceSpiffeID(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"spire_registration_entry": resourceRegistrationEntry(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	spireServer      *grpc.ClientConn
	spireEntryClient spireentry.EntryClient
	spireTokenClient spiretoken.AgentClient
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		host := d.Get(provCfgServerHostAttr).(string)
		port := d.Get(provCfgServerPortAttr).(string)
		serverAddr := net.JoinHostPort(host, port)

		spireServer, err := grpc.DialContext(ctx, serverAddr)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return &apiClient{
			spireServer:      spireServer,
			spireEntryClient: spireentry.NewEntryClient(spireServer),
			spireTokenClient: spiretoken.NewAgentClient(spireServer),
		}, nil
	}
}
