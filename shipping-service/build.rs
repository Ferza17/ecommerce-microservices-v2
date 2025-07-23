fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .file_descriptor_set_path("descriptor.bin")
        .build_server(true)
        .build_client(false)
        .out_dir("src/model/rpc")
        .type_attribute(".", "#[derive(serde::Serialize, serde::Deserialize)]")
        .extern_path(".google.protobuf.Timestamp", "::prost_wkt_types::Timestamp")
        .compile_protos(
            &[
                "proto/v1/shipping/service.proto",
                "proto/v1/shipping/request.proto",
                "proto/v1/shipping/response.proto",
                "proto/v1/shipping/model.proto",
            ],
            &["proto/"],
        )?;
    Ok(())
}
