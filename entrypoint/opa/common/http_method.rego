package common.http_method
default allow = false


# --- POLICY
allow if {
    input.method = "GET"
}

allow if {
    input.method = "POST"
}

allow if {
    input.method = "PUT"
}

allow if {
    input.method = "DELETE"
}