use crate::model::rpc::shipping::{
    CreateShippingProviderRequest, CreateShippingProviderResponse, DeleteShippingProviderResponse,
    GetShippingProviderByIdResponse, ListShippingProvidersRequest, ListShippingProvidersResponse,
    UpdateShippingProviderRequest, UpdateShippingProviderResponse,
};
use crate::package::context::auth::AUTHORIZATION_HEADER;
use crate::package::context::request_id::X_REQUEST_ID_HEADER;
use HttpAuthScheme::Bearer;
use utoipa::openapi::security::{HttpAuthScheme, HttpBuilder, SecurityScheme};
use utoipa::openapi::{Header, Schema};
use utoipa::{Modify, OpenApi};

use utoipa::openapi::schema::SchemaType;

#[allow(unused_imports)]
#[derive(OpenApi)]
#[openapi(
    paths(
        // SHIPPING PROVIDER
        crate::module::shipping_provider::presenter_http::create_shipping_provider,
        crate::module::shipping_provider::presenter_http::list_shipping_providers,
        crate::module::shipping_provider::presenter_http::get_shipping_provider_by_id,
        crate::module::shipping_provider::presenter_http::update_shipping_provider,
        crate::module::shipping_provider::presenter_http::delete_shipping_provider
    ),
    components(
        schemas(
            // SHIPPING PROVIDER
            CreateShippingProviderRequest,
            CreateShippingProviderResponse,
            ListShippingProvidersRequest,
            ListShippingProvidersResponse,
            GetShippingProviderByIdResponse,
            UpdateShippingProviderRequest,
            UpdateShippingProviderResponse,
            DeleteShippingProviderResponse
        ),
    ),
    tags(
        (name = "Shipping Provider", description = "Shipping Provider API")
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
                    .scheme(Bearer)
                    .description(Some(
                        "Bearer Token (e.g., JWT) for authentication. Format: Bearer <token>",
                    ))
                    .bearer_format("JWT")
                    .build(),
            ),
        );
    }
}
