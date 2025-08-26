package common.role
import rego.v1

valid_roles := {
    "ADMIN",
    "SUPER_ADMIN",
    "CUSTOMER",
    "CUSTOMER_MEMBERSHIP"
}

access_allowed if {
    input.user_role in valid_roles
}

has(allowed_roles) if {
    input.user_role in allowed_roles
}