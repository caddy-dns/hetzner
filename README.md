# Hetzner module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Hetzner accounts.

Version 2 of this module (`github.com/caddy-dns/hetzner/v2`) is designed for use with DNS zones managed in the Hetzner Console via the [Cloud DNS API](https://docs.hetzner.cloud/reference/cloud#dns).
If your zone is still managed in the old DNS Console and has not yet been [migrated](https://docs.hetzner.com/networking/dns/migration-to-hetzner-console/process) to the new Hetzner Console, please use version 1 of the module (`github.com/caddy-dns/hetzner`).

## Caddy module name

```
dns.providers.hetzner
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "hetzner",
        "api_token": "YOUR_HETZNER_AUTH_API_TOKEN"
      }
    }
  }
}
```

or with the Caddyfile:

```
your.domain.com {
  respond "Hello World"	# replace with whatever config you need...
  tls {
    dns hetzner {env.YOUR_HETZNER_AUTH_API_TOKEN}
    propagation_delay 30s
  }
}
```

You can replace `{env.YOUR_HETZNER_AUTH_API_TOKEN}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.

Setting `propagation_delay` to `30s` causes Caddy to wait for 30 seconds before starting the DNS TXT records propagation checks.
This fixes a known issue caused by slow DNS propagation (see [#11](https://github.com/caddy-dns/hetzner/issues/11) for details).

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/hetzner#authenticating) for important information about credentials.
