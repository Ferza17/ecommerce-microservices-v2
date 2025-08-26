package http.payment_provider

import rego.v1
import data.common.role
import data.common.http_method

allow if {
    startswith(input.path, "/api/v1/payment-providers")
    http_method.has({"GET"})
    role.has({"CUSTOMER", "CUSTOMER_MEMBERSHIP", "ADMIN"})
    role.access_allowed
}