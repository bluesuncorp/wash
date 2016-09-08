package middleware

import "github.com/go-playground/lars"

const (
	xFrameOptions                = "X-Frame-Options"
	xFrameOptionsValue           = "DENY"
	xContentTypeOptions          = "X-Content-Type-Options"
	xContentTypeOptionsValue     = "nosniff"
	xssProtection                = "X-XSS-Protection"
	xssProtectionValue           = "1; mode=block"
	strictTransportSecurity      = "Strict-Transport-Security"                    // details https://blog.bracelab.com/achieving-perfect-ssl-labs-score-with-go + https://developer.mozilla.org/en-US/docs/Web/Security/HTTP_strict_transport_security
	strictTransportSecurityValue = "max-age=31536000; includeSubDomains; preload" // 31536000 = just shy of 12 months
	// also look at Content-Security-Policy in the future.
)

// Security Adds HTTP headers for XSS Protection and alike.
func Security(c lars.Context) {
	c.Response().Header().Add(xFrameOptions, xFrameOptionsValue)
	c.Response().Header().Add(xContentTypeOptions, xContentTypeOptionsValue)
	c.Response().Header().Add(xssProtection, xssProtectionValue)
	c.Response().Header().Add(strictTransportSecurity, strictTransportSecurityValue)

	c.Next()
}
