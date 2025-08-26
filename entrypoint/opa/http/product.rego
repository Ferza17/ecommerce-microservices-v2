package http.product
default allow = false

allow if {
    startswith(input.path, "/api/v1/products")
    input.method == "GET"
    input.user_role in ["CUSTOMER", "CUSTOMER_MEMBERSHIP", "ADMIN"]
}

