FROM rust:1.88.0-trixie AS builder

# Install system dependencies
RUN apt-get update && apt-get install -y \
    libpq-dev pkg-config protobuf-compiler libssl-dev \
    && rm -rf /var/lib/apt/lists/*

RUN cargo install diesel_cli --no-default-features --features postgres

WORKDIR /usr/src/app

# Copy manifest first for dependency caching
COPY Cargo.toml Cargo.lock ./
COPY build.rs ./

# Create dummy src to cache deps
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo build --release || true
RUN rm -rf src

# copy full source
COPY . .

# Build actual binary
RUN cargo build --release --bin shipping-service

FROM debian:bookworm-slim AS runtime
RUN apt-get update && apt-get install -y \
    libpq-dev ca-certificates curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/bin

# Copy diesel binary
COPY --from=builder /usr/local/cargo/bin/diesel /usr/local/bin/diesel
# Copy the built binary
COPY --from=builder /usr/src/app/target/release/shipping-service /usr/bin/shipping-service

# Copy configuration and other files to the working directory
COPY migrations ./migrations
COPY config ./config
COPY proto ./proto

ENTRYPOINT ["/bin/sh", "-c", "diesel migration run --migration-dir ./migrations && exec /usr/bin/shipping-service run insert-mock && exec /usr/bin/shipping-service run --direction production"]