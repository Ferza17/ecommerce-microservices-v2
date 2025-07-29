use std::env;
use std::path::PathBuf;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let out_dir = PathBuf::from("src/model/rpc");

    // List of proto files
    let proto_files = &[
        // SHIPPING
        "proto/v1/shipping/service.proto",
        "proto/v1/shipping/request.proto",
        "proto/v1/shipping/response.proto",
        "proto/v1/shipping/model.proto",
        // USER
        "proto/v1/user/enum.proto",
        "proto/v1/user/option.proto",
        "proto/v1/user/service.proto",
        "proto/v1/user/request.proto",
        "proto/v1/user/response.proto",
        "proto/v1/user/model.proto",
        // PAYMENT
        "proto/v1/payment/service.proto",
        "proto/v1/payment/request.proto",
        "proto/v1/payment/response.proto",
        "proto/v1/payment/model.proto",
        "proto/v1/payment/enum.proto",
        // COMMON
        "proto/v1/common/response/response.proto",
    ];
    let proto_include = "proto/";

    let mut tonic_config = tonic_build::configure()
        .file_descriptor_set_path("descriptor.bin")
        .build_server(true)
        .build_client(true)
        .build_transport(true)
        .out_dir(&out_dir)
        .type_attribute(".", "#[derive(utoipa::ToSchema)]")
        .type_attribute(".", "#[derive(serde::Serialize, serde::Deserialize)]")
        .extern_path(".google.protobuf.Timestamp", "::prost_wkt_types::Timestamp")
        .extern_path(".google.protobuf.Struct", "::prost_wkt_types::Struct")
        .field_attribute(
            "created_at",
            "#[schema(value_type = String, format = \"date-time\")]",
        )
        .field_attribute(
            "updated_at",
            "#[schema(value_type = String, format = \"date-time\")]",
        )
        .field_attribute(
            "deleted_at",
            "#[schema(value_type = String, format = \"date-time\")]",
        )
        .field_attribute(
            "discarded_at",
            "#[schema(value_type = String, format = \"date-time\")]",
        )
        .field_attribute(
            "timestamp",
            "#[schema(value_type = String, format = \"date-time\")]",
        )
        .field_attribute("Response.data", "#[schema(value_type = serde_json::Value)]");

    // let mut config = {
    //     let mut c = prost_build::Config::new();
    //     c.service_generator(tonic_config.clone().service_generator());
    //     c
    // };

    // prost_validate_build::Builder::new().compile_protos(proto_files, &[proto_include])?;

    tonic_config
        .clone()
        .out_dir(&out_dir)
        .compile_protos(proto_files, &[proto_include])?;

    Ok(())
}
