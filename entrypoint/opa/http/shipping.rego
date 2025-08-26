package http.shipping
default allow = false

allow if {
    startswith(input.path, "/api/v1/shippings")
    input.method == "GET"
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
}

