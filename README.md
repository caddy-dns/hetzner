# Hetzner Module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It manages DNS records for the Hetzner Console using the [Cloud DNS API](https://docs.hetzner.cloud/reference/cloud#dns).

## Caddy Module Name

```
dns.providers.hetzner
```

## Configuration

To use this module inside your Caddy JSON config, append the JSON object provided below to
the list
[`apps.tls.automation.policies.issuers`](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme):

```json
{
    "module": "acme",
    "challenges": {
        "dns": {
            "propagation_delay": 30000000000,
            "provider": {
                "auth_api_token": "YOUR_HETZNER_AUTH_API_TOKEN",
                "name": "hetzner"
            }
        }
    }
}
```

Or in the Caddyfile:

```Caddyfile
# For all apps
{
    acme_dns hetzner {env.YOUR_HETZNER_AUTH_API_TOKEN}
}

# For an individual app
your.domain.com {
    respond "Hello World" # Replace with whatever config you need...

    tls {
        dns hetzner {env.YOUR_HETZNER_AUTH_API_TOKEN}
        propagation_delay 30s
    }
}
```

If you prefer to put the actual auth token directly in your config instead of an environment variable, you can replace `{env.YOUR_HETZNER_AUTH_API_TOKEN}` with it.

Setting propagation_delay to 30s causes Caddy to wait 30 seconds before starting the DNS TXT record propagation checks.  This resolves an issue that was occurring due to slow DNS propagation (see issue [#11](https://github.com/caddy-dns/hetzner/issues/11) for details).

## Authentication

For important information about credentials, see the associated [README](https://github.com/libdns/hetzner#authenticating) in the libdns package.
