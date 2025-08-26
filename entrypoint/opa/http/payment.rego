package http.payment
default allow = false

allow if {
    startswith(input.path, "/api/v1/payments")
    input.method == "POST"
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
}

