fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(true)
        .build_client(false)
        .out_dir("src/model/rpc")
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
