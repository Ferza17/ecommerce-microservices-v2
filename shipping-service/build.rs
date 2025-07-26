fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Generate gRPC code with utoipa support
    tonic_build::configure()
        .file_descriptor_set_path("descriptor.bin")
        .build_server(true)
        .build_client(true)
        .build_transport(true)
        .out_dir("src/model/rpc")
        .type_attribute(".", "#[derive(utoipa::ToSchema)]")
        .type_attribute(".", "#[derive(serde::Serialize, serde::Deserialize)]")
        .protoc_arg("--plugin=protoc-gen-prost-utoipa")
        .protoc_arg("--prost-utoipa_opt=generate_schema=true")
        .extern_path(".google.protobuf.Timestamp", "::prost_wkt_types::Timestamp")
        .extern_path(".google.protobuf.Struct", "::prost_wkt_types::Struct")
        // Target common timestamp field names - use more specific patterns
        .field_attribute("created_at", "#[schema(value_type = String, format = \"date-time\")]")
        .field_attribute("updated_at", "#[schema(value_type = String, format = \"date-time\")]")
        .field_attribute("deleted_at", "#[schema(value_type = String, format = \"date-time\")]")
        .field_attribute("discarded_at", "#[schema(value_type = String, format = \"date-time\")]")
        .field_attribute("timestamp", "#[schema(value_type = String, format = \"date-time\")]")
        // Target common struct field names - use more specific patterns
        .field_attribute("data", "#[schema(value_type = String, format = \"object\")]")
        
        
        .compile_protos(
            &[
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
                // Payment
                "proto/v1/payment/service.proto",
                "proto/v1/payment/request.proto",
                "proto/v1/payment/response.proto",
                "proto/v1/payment/model.proto",
                "proto/v1/payment/enum.proto",
                // COMMON
                "proto/v1/common/response/response.proto",
            ],
            &["proto/"],
        )?;
    Ok(())
}