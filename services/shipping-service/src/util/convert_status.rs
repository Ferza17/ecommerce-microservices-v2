pub fn tonic_to_http_status(status: tonic::Code) -> axum::http::StatusCode {
    match status {
        tonic::Code::Ok => axum::http::StatusCode::OK,
        tonic::Code::Cancelled => axum::http::StatusCode::REQUEST_TIMEOUT,
        tonic::Code::Unknown => axum::http::StatusCode::INTERNAL_SERVER_ERROR,
        tonic::Code::InvalidArgument => axum::http::StatusCode::BAD_REQUEST,
        tonic::Code::DeadlineExceeded => axum::http::StatusCode::GATEWAY_TIMEOUT,
        tonic::Code::NotFound => axum::http::StatusCode::NOT_FOUND,
        tonic::Code::AlreadyExists => axum::http::StatusCode::CONFLICT,
        tonic::Code::PermissionDenied => axum::http::StatusCode::FORBIDDEN,
        tonic::Code::Unauthenticated => axum::http::StatusCode::UNAUTHORIZED,
        tonic::Code::ResourceExhausted => axum::http::StatusCode::TOO_MANY_REQUESTS,
        tonic::Code::FailedPrecondition => axum::http::StatusCode::PRECONDITION_FAILED,
        tonic::Code::Aborted => axum::http::StatusCode::CONFLICT,
        tonic::Code::OutOfRange => axum::http::StatusCode::BAD_REQUEST,
        tonic::Code::Unimplemented => axum::http::StatusCode::NOT_IMPLEMENTED,
        tonic::Code::Internal => axum::http::StatusCode::INTERNAL_SERVER_ERROR,
        tonic::Code::Unavailable => axum::http::StatusCode::SERVICE_UNAVAILABLE,
        tonic::Code::DataLoss => axum::http::StatusCode::INTERNAL_SERVER_ERROR,
    }
}
