use crate::model::rpc::shipping::{
    CreateShippingRequest, CreateShippingResponse, DeleteShippingResponse, GetShippingByIdResponse,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    ListShippingRequest, ListShippingResponse, UpdateShippingRequest, UpdateShippingResponse,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use utoipa::openapi::security::{HttpAuthScheme, HttpBuilder, SecurityScheme};
use utoipa::{Modify, OpenApi};

#[allow(unused_imports)]
#[derive(OpenApi)]
#[openapi(
    servers(
        (url = "http://127.0.0.1:40057", description = "HTTP"),
        (url = "https://127.0.0.1/v1/shipping/", description = "HTTPS"),
    ),
    paths(
        // SHIPPING
        crate::module::shipping::presenter_http::create_shipping,
        crate::module::shipping::presenter_http::get_shipping_provider_by_id,
        crate::module::shipping::presenter_http::list_shipping_providers,
        crate::module::shipping::presenter_http::update_shipping,
        crate::module::shipping::presenter_http::delete_shipping,

        // SHIPPING PROVIDER
        crate::module::shipping_provider::presenter_http::list_shipping_providers,
        crate::module::shipping_provider::presenter_http::get_shipping_provider_by_id,
    ),
    components(
        schemas(
            // SHIPPING
            CreateShippingRequest,
            CreateShippingResponse,
            GetShippingByIdResponse,
            ListShippingResponse,
            ListShippingRequest,
            UpdateShippingRequest,
            UpdateShippingResponse,
            DeleteShippingResponse,

            // SHIPPING PROVIDER
            ListShippingProvidersRequest,
            ListShippingProvidersResponse,
            GetShippingProviderByIdResponse,
        ),
    ),
    tags(
        (name = "Shipping Provider", description = "Shipping Provider API"),
        (name = "Shipping", description = "Shipping API")

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
