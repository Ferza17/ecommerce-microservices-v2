package authz

default allow = false

# --- audit log as a set of messages
audit_log[msg] if {
    msg := sprintf("User %s (%s) attempted %s %s", [
        input.user_id,
        input.user_role,
        input.method,
        input.path,
    ])
}

# --- suspicious path checks
deny[msg] if {
    contains(input.path, "..")
    msg := "path contains .."
}

deny[msg] if {
    contains(input.path, "//")
    msg := "path contains //"
}

# --- helper role checks
access_allowed if {
    input.user_role == "ADMIN"
}

access_allowed if {
    input.user_role == "SUPER_ADMIN"
}

access_allowed if {
    input.user_role == "CUSTOMER"
}

access_allowed if {
    input.user_role == "CUSTOMER_MEMBERSHIP"
}

# --- AUTH
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

# --- PAYMENT
allow if {
    startswith(input.path, "/api/v1/payments")
    input.method in ["POST"]
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
    access_allowed
}

# --- PAYMENT PROVIDER
allow if {
    startswith(input.path, "/api/v1/payment-providers")
    input.method in ["GET"]
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
    access_allowed
}

# --- PRODUCT
allow if {
    startswith(input.path, "/api/v1/products")
    input.method in ["GET"]
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
    access_allowed
}

# --- SHIPPING PROVIDER
allow if {
    startswith(input.path, "/api/v1/shipping-providers")
    input.method in ["GET"]
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
    access_allowed
}
