use crate::model::rpc::{
    payment::{FindPaymentProvidersRequest, FindPaymentProvidersResponse, CreatePaymentRequest, CreatePaymentResponse},
    product::{FindProductsWithPaginationRequest, FindProductsWithPaginationResponse},
    response::Response,
    user::{
        AuthUserRegisterRequest, AuthUserRegisterResponse, AuthUserVerifyOtpRequest,
        AuthUserVerifyOtpResponse,
    },
    shipping::{
        ListShippingProvidersRequest, ListShippingProvidersResponse,
    }
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use utoipa::{
    Modify, OpenApi,
    openapi::security::{HttpAuthScheme, HttpBuilder, SecurityScheme},
};

#[allow(unused_imports)]
#[derive(OpenApi)]
#[openapi(
    servers(
        (url = "http://127.0.0.1:4000", description = "HTTP"),
        (url = "https://127.0.0.1:4000", description = "HTTPS"),
    ),
    paths(
        // AUTH
        crate::module::user::http_presenter::auth_register,
        crate::module::user::http_presenter::auth_user_login_by_email_and_password,
        crate::module::user::http_presenter::auth_user_verify_otp,

        // PRODUCT
        crate::module::product::http_presenter::find_products_with_pagination,

        // PAYMENT PROVIDERS
        crate::module::payment::http_presenter::find_payment_providers,

        // PAYMENT
        crate::module::payment::http_presenter::create_payment,
    
        // SHIPPING PROVIDERS
        crate::module::shipping::http_presenter::list_shipping_providers,
    ),
    components(
        schemas(
            // COMMON
            Response,
            // AUTH
            AuthUserRegisterRequest,
            AuthUserRegisterResponse,
            AuthUserVerifyOtpRequest,
            AuthUserVerifyOtpResponse,
            // PRODUCT
            FindProductsWithPaginationRequest,
            FindProductsWithPaginationResponse,
            // PAYMENT PROVIDERS
            FindPaymentProvidersRequest,
            FindPaymentProvidersResponse,
            // PAYMENT
            CreatePaymentRequest,
            CreatePaymentResponse,
            // SHIPPING PROVIDERS
            ListShippingProvidersRequest,
            ListShippingProvidersResponse,
        ),
    ),
    tags(
        (name = "Auth", description = "Authentication route API"),
        (name = "Product", description = "Product route API"),
        (name = "ShippingProviders", description = "Shipping Providers route API"),
        (name = "Shipping", description = "Shipping route API"),
        (name = "PaymentProviders", description = "Payment Providers route API"),
        (name = "Payment", description = "Payment route API"),
    ),
    modifiers(&SecurityAddon)
)]
pub struct ApiDocs;

struct SecurityAddon;
impl Modify for SecurityAddon {
    fn modify(&self, openapi: &mut utoipa::openapi::OpenApi) {
        let components = openapi.components.get_or_insert_with(Default::default);
        components.security_schemes.insert(
            AUTHORIZATION_HEADER.to_string(),
            SecurityScheme::Http(
                HttpBuilder::new()
                    .scheme(HttpAuthScheme::Bearer)
                    .description(Some(
                        "Bearer Token (e.g., JWT) for authentication. Format: Bearer <token>",
                    ))
                    .bearer_format("JWT")
                    .build(),
            ),
        );
    }
}
