pub const X_USER_ID_HEADER: &str = "x-user-id";

pub fn get_request_x_user_id_header_from_header(
    header: axum::http::HeaderMap<axum::http::HeaderValue>,
) -> String {
    header
        .get(X_USER_ID_HEADER)
        .unwrap()
        .to_str()
        .expect("Invalid header value")
        .to_string()
}
