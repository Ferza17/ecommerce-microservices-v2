package common.http_method
import rego.v1

valid_http_method := {
    "GET", "POST", "PUT", "DELETE"
}

access_allowed if {
    input.method in valid_http_method
}

has(allowed_methods) if {
    input.method in allowed_methods
}