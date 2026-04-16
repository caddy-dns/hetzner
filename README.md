# Hetzner Module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It manages DNS records for the Hetzner Console using the [Cloud DNS API](https://docs.hetzner.cloud/reference/cloud#dns).

## Caddy Module Name

```
dns.providers.hetzner
```

## Configuration

To use this module for the ACME DNS challenge, configure the [ACME issuer](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme) in your Caddy JSON as follows:

```json
{
    "module": "acme",
    "challenges": {
        "dns": {
            "propagation_delay": "30s",
            "provider": {
                "name": "hetzner",
                "api_token": "YOUR_HETZNER_AUTH_API_TOKEN"
            }
        }
    }
}
```

Or in the Caddyfile:

```Caddyfile
your.domain.com {
    respond "Hello World"

    tls {
        dns hetzner {env.YOUR_HETZNER_AUTH_API_TOKEN}
        propagation_delay 30s
    }
}
```

The Hetzner plugin can also be configured globally:

```Caddyfile
{
    acme_dns hetzner {env.YOUR_HETZNER_AUTH_API_TOKEN}
}
```

If you prefer to put the actual auth token directly in your config instead of an environment variable, you can replace `{env.YOUR_HETZNER_AUTH_API_TOKEN}` with it.

Setting propagation_delay to 30s causes Caddy to wait 30 seconds before starting the DNS TXT record propagation checks.  This resolves an issue that was occurring due to slow DNS propagation (see issue [#11](https://github.com/caddy-dns/hetzner/issues/11) for details).

## Authentication

For information about credentials, see the associated [README](https://github.com/libdns/hetzner#authenticating) in the libdns package.

## Building

To add this module to Caddy, go to the download page and add the `hetzner` module or build Caddy using `xcaddy` or `caddy add-package`.

```
# Using xcaddy
xcaddy build --with github.com/caddy-dns/hetzner/v2

# Using caddy add-package
caddy add-package github.com/caddy-dns/hetzner@v2.0.0
```
