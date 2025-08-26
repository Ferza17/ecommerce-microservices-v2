package http.auth
default allow = false

allow if {
    input.path == "/api/v1/auth/register"
    input.method == "POST"
}

allow if {
    input.path == "/api/v1/auth/login"
    input.method == "POST"
}

allow if {
    input.path == "/api/v1/auth/verify-otp"
    input.method == "POST"
}