use base64::Engine;
use rand::Rng;
use crate::config::config::AppConfig;

pub const CONNECTION_HEADER: &str = "connection";
const CONNECTION_HEADER_VALUE_UPGRADE: &str = "Upgrade";
pub fn default_connection_header() -> String {
    CONNECTION_HEADER_VALUE_UPGRADE.to_string()
}

pub const UPGRADE_HEADER: &str = "upgrade";
const UPGRADE_HEADER_VALUE_WEBSOCKET: &str = "websocket";
pub fn default_upgrade_header() -> String {
    UPGRADE_HEADER_VALUE_WEBSOCKET.to_string()
}

pub const SEC_WEBSOCKET_KEY: &str = "Sec-WebSocket-Key";
pub fn default_sec_websocket_key() -> String {
    let mut key = [0u8; 16];
    rand::thread_rng().fill(&mut key);
    base64::engine::general_purpose::STANDARD.encode(key)
}

pub const SEC_WEBSOCKET_VERSION: &str = "Sec-WebSocket-Version";
const SEC_WEBSOCKET_VERSION_VALUE: &str = "13";
pub fn default_sec_websocket_version() -> String {
    SEC_WEBSOCKET_VERSION_VALUE.to_string()
}

pub const SEC_WEBSOCKET_EXTENSIONS: &str = "Sec-WebSocket-Extensions";
const SEC_WEBSOCKET_EXTENSIONS_VALUE: &str = "permessage-deflate; client_max_window_bits";
pub fn default_sec_websocket_extensions() -> String {
    SEC_WEBSOCKET_EXTENSIONS_VALUE.to_string()
}


pub const HOST_HEADER: &str = "host";