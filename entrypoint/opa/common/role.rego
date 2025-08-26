package common.roles
default allow = false

# --- helper role checks
allow if {
    input.user_role == "ADMIN"
}

allow if {
    input.user_role == "SUPER_ADMIN"
}

allow if {
    input.user_role == "CUSTOMER"
}

allow if {
    input.user_role == "CUSTOMER_MEMBERSHIP"
}
