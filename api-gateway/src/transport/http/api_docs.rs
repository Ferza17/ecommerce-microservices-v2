use crate::model::rpc::{
    payment::{
        CreatePaymentRequest, CreatePaymentResponse, FindPaymentProvidersRequest,
        FindPaymentProvidersResponse,
    },
    product::{FindProductsWithPaginationRequest, FindProductsWithPaginationResponse},
    response::{Response, ResponseCommand},
    shipping::{ListShippingProvidersRequest, ListShippingProvidersResponse},
    user::{
        AuthUserRegisterRequest, AuthUserRegisterResponse, AuthUserVerifyOtpRequest,
        AuthUserVerifyOtpResponse,
    },
};
use utoipa::openapi::Required;
use utoipa::openapi::path::{ParameterBuilder, ParameterIn};
use utoipa::openapi::security::{ApiKey, ApiKeyValue};
use utoipa::{
    Modify, OpenApi,
    openapi::security::{HttpAuthScheme, HttpBuilder, SecurityScheme},
};

pub const ROUTE_PREFIX: &str = "/api/v1/docs";

#[allow(unused_imports)]
#[derive(OpenApi)]
#[openapi(
    servers(
        (url = "http://127.0.0.1:4000", description = "HTTP"),
        (url = "https://127.0.0.1:4000", description = "HTTPS"),
    ),
    paths(
        // AUTH
        crate::module::auth::http_presenter::auth_register,
        crate::module::auth::http_presenter::auth_user_login_by_email_and_password,
        crate::module::auth::http_presenter::auth_user_verify_otp,

        // NOTIFICATION
        crate::module::notification::http_presenter::get_notification_with_request_id,

        // PRODUCT
        crate::module::product::http_presenter::find_products_with_pagination,

        // PAYMENT PROVIDERS
        crate::module::payment_providers::http_presenter::find_payment_providers,

        // PAYMENT
        crate::module::payment::http_presenter::create_payment,

        // SHIPPING PROVIDERS
        crate::module::shipping_provider::http_presenter::list_shipping_providers,
    ),
    components(
        schemas(
            // COMMON
            Response,
            ResponseCommand,
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
        (name = crate::module::auth::http_presenter::TAG, description = "Authentication route API"),
        (name = crate::module::notification::http_presenter::TAG, description = "Notification route API"),
        (name = crate::module::product::http_presenter::TAG, description = "Product route API"),
        (name = crate::module::shipping_provider::http_presenter::TAG, description = "Shipping Providers route API"),
        (name = crate::module::shipping::http_presenter::TAG, description = "Shipping route API"),
        (name = crate::module::payment_providers::http_presenter::TAG, description = "Payment Providers route API"),
        (name = crate::module::payment::http_presenter::TAG, description = "Payment route API"),
    ),
    modifiers(&SecurityAddon)
)]
pub struct ApiDocs;

struct SecurityAddon;
impl Modify for SecurityAddon {
    fn modify(&self, openapi: &mut utoipa::openapi::OpenApi) {
        let components = openapi.components.get_or_insert_with(Default::default);
        components.security_schemes.insert(
            crate::package::context::auth::AUTHORIZATION_HEADER.to_string(),
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

        components.security_schemes.insert(
            crate::package::context::request_id::X_REQUEST_ID_HEADER.to_string(),
            SecurityScheme::ApiKey(ApiKey::Header(ApiKeyValue::new(
                crate::package::context::request_id::X_REQUEST_ID_HEADER.to_string(),
            ))),
        );
    }
}
