package http.payment_providers
default allow = false

# --- PAYMENT PROVIDER
allow if {
    startswith(input.path, "/api/v1/payment-providers")
    input.method == "GET"
    input.user_role in ["CUSTOMER","CUSTOMER_MEMBERSHIP","ADMIN"]
}