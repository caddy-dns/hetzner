package hetzner

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/hetzner/v2"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *hetzner.Provider }

func init() {
	caddy.RegisterModule(&Provider{})
}

// CaddyModule returns the Caddy module information.
func (*Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.hetzner",
		New: func() caddy.Module { return &Provider{new(hetzner.Provider)} },
	}
}

// Provision implements the caddy.Provisioner interface.
func (p *Provider) Provision(_ caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.APIToken = repl.ReplaceAll(p.Provider.APIToken, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	hetzner [<api_token>] {
//	    api_token <api_token>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				p.Provider.APIToken = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
