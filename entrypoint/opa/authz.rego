package authz
default allow = false

# --- COMMON
import data.common.base
import data.common.http_method
import data.common.role

# Collect denies from base
deny[msg] if {
    base.deny[msg]
}

audit_log[msg] if {
    base.audit_log[msg]
}

# --- HTTP
import data.http.auth
import data.http.payment
import data.http.payment_providers
import data.http.product
import data.http.shipping
import data.http.shipping_providers


allow if {
    http_method.allow
    role.allow
    auth.allow
}

allow if {
    http_method.allow
    role.allow
    payment.allow
}

allow if {
    http_method.allow
    role.allow
    payment_providers.allow
}

allow if {
    http_method.allow
    role.allow
    product.allow
}

allow if {
    http_method.allow
    role.allow
    shipping.allow
}

allow if {
    http_method.allow
    role.allow
    shipping_providers.allow
}