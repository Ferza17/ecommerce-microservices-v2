package http.auth

import rego.v1
import data.common.http_method


allow if {
    input.path == "/api/v1/auth/register"
    http_method.has({"POST"})
}

allow if {
    input.path == "/api/v1/auth/login"
    http_method.has({"POST"})
}

allow if {
    input.path == "/api/v1/auth/verify-otp"
    http_method.has({"POST"})
}
