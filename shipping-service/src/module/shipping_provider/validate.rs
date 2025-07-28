use crate::model::rpc::shipping::{GetShippingProviderByIdRequest, ListShippingProvidersRequest};
use tonic::{Request, Status};

pub fn validate_get_shipping_provider_by_id(
    request: &Request<GetShippingProviderByIdRequest>,
) -> Option<Status> {
    if request.get_ref().id.is_empty() {
        return Some(Status::invalid_argument("id"));
    }
    None
}

pub fn validate_list_shipping_providers(
    request: &Request<ListShippingProvidersRequest>,
) -> Option<Status> {
    if request.get_ref().page <= 0 {
        return Some(Status::invalid_argument("invalid argument page"));
    }
    if request.get_ref().limit <= 0 {
        return Some(Status::invalid_argument("invalid argument limit"));
    }
    None
}
