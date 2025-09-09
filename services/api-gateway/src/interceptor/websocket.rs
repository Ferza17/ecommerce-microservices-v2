use axum::extract::OriginalUri;
use axum::http::header::HOST;

#[derive(Clone, Debug)]
pub struct WebsocketLayer;

impl<S> tower::Layer<S> for WebsocketLayer {
    type Service = WebsocketLayerService<S>;
    fn layer(&self, inner: S) -> Self::Service {
        WebsocketLayerService { inner }
    }
}

#[derive(Clone, Debug)]
pub struct WebsocketLayerService<S> {
    pub inner: S,
}

impl<S> tower::Service<axum::http::Request<axum::body::Body>> for WebsocketLayerService<S>
where
    S: tower::Service<
            axum::http::Request<axum::body::Body>,
            Error = std::convert::Infallible,
            Response = axum::http::Response<axum::body::Body>,
        > + Clone
        + Send
        + std::fmt::Debug
        + 'static,
    S::Response: Send + 'static,
    S::Future: Send + 'static,
{
    type Response = axum::http::Response<axum::body::Body>;
    type Error = std::convert::Infallible;
    type Future = futures::future::Either<
        S::Future,
        futures::future::Ready<
            Result<axum::http::Response<axum::body::Body>, std::convert::Infallible>,
        >,
    >;

    #[tracing::instrument("WebsocketLayer.poll_ready")]
    fn poll_ready(
        &mut self,
        cx: &mut std::task::Context<'_>,
    ) -> std::task::Poll<Result<(), Self::Error>> {
        self.inner.poll_ready(cx)
    }

    #[tracing::instrument("WebsocketLayer.call")]
    fn call(&mut self, mut req: axum::http::Request<axum::body::Body>) -> Self::Future {
        
        // SET DEFAULT CONNECTION HEADER
        req.headers_mut().insert(
            crate::package::context::websocket::CONNECTION_HEADER,
            crate::package::context::websocket::default_connection_header()
                .parse()
                .unwrap(),
        );

        // SET DEFAULT UPGRADE HEADER
        req.headers_mut().insert(
            crate::package::context::websocket::UPGRADE_HEADER,
            crate::package::context::websocket::default_upgrade_header()
                .parse()
                .unwrap(),
        );

        // SET DEFAULT SEC_WEBSOCKET_VERSION HEADER
        req.headers_mut().insert(
            crate::package::context::websocket::SEC_WEBSOCKET_VERSION,
            crate::package::context::websocket::default_sec_websocket_version()
                .parse()
                .unwrap(),
        );
        

        // SET DEFAULT SEC_WEBSOCKET_KEY HEADER
        // req.headers_mut().insert(
        //     crate::package::context::websocket::SEC_WEBSOCKET_KEY,
        //     crate::package::context::websocket::default_sec_websocket_key()
        //         .parse()
        //         .unwrap(),
        // );
        
        // SET DEFAULT SEC_WEBSOCKET_EXTENSIONS HEADER
        req.headers_mut().insert(
            crate::package::context::websocket::SEC_WEBSOCKET_EXTENSIONS,
            crate::package::context::websocket::default_sec_websocket_extensions()
                .parse()
                .unwrap(),
        );

        futures::future::Either::Left(self.inner.call(req))
    }
}
