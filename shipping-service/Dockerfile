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

# Now copy full source
COPY . .

# Build actual binary
RUN cargo build --release --bin shipping-service

FROM debian:bookworm-slim AS runtime
RUN apt-get update && apt-get install -y \
    libpq-dev ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/bin

COPY --from=builder /usr/local/cargo/bin/diesel /usr/local/bin/diesel
COPY --from=builder /usr/src/app/target/release/shipping-service /usr/bin/shipping-service

COPY migrations /migrations
COPY config /config

ENV DATABASE_URL=postgres://postgres:1234@postgres-local:5432/shippings

ENTRYPOINT ["/bin/sh", "-c", "diesel migration run --migration-dir /migrations && exec /usr/bin/shipping-service"]
