package hetzner

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/hetzner"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *hetzner.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.hetzner",
		New: func() caddy.Module { return &Provider{new(hetzner.Provider)} },
	}
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// hetzner [<api_token>] {
//     api_token <api_token>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	repl := caddy.NewReplacer()
	for d.Next() {
		if d.NextArg() {
			p.Provider.AuthAPIToken = repl.ReplaceAll(d.Val(), "")
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.AuthAPIToken != "" {
					return d.Err("API token already set")
				}
				p.Provider.AuthAPIToken = repl.ReplaceAll(d.Val(), "")
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.AuthAPIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guard
var _ caddyfile.Unmarshaler = (*Provider)(nil)
