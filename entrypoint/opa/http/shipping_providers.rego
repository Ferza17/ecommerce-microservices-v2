package http.shipping_providers
default allow = false

allow if {
    startswith(input.path, "/api/v1/shipping-providers")
    input.method == "GET"
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
}