package http.payment

import rego.v1
import data.common.role
import data.common.http_method

allow if {
    startswith(input.path, "/api/v1/payments")
    http_method.has({"POST"})
    role.has({"CUSTOMER", "CUSTOMER_MEMBERSHIP"})
    role.access_allowed
}