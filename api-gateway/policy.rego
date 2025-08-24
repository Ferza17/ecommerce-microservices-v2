package authz

default allow = false

## AUTH ROUTE
allow {
    input.path == "/api/v1/auth/login"
    input.method == "POST"
}

allow {
    input.path == "/api/v1/auth/register"
    input.method == "POST"
}

