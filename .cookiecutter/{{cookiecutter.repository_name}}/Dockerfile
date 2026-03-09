# ------------------------------------------------------------------------------
# {{cookiecutter.__copyright}}
# ------------------------------------------------------------------------------
ARG REGISTRY_PROXY=cds-harbor.rtplab.nimblestorage.com/docker_proxy/library/

# ------------------------------------------------------------------------------
# Builder Image
# ------------------------------------------------------------------------------
FROM ${REGISTRY_PROXY}golang:1.20-bullseye AS build

WORKDIR /build

# Copy the source files needed to build the project.
COPY Makefile Makefile
COPY go.mod go.sum ./
COPY vendor vendor
COPY internal internal
COPY cmd cmd

# Build the project.
ARG VERSION
RUN set -eux; \
    CGO_ENABLED=0 make build VERSION=${VERSION}

# ------------------------------------------------------------------------------
# Assets Retrieval Image
# ------------------------------------------------------------------------------
FROM ${REGISTRY_PROXY}debian:bullseye-slim AS assets

WORKDIR /build

RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        # Used to download gRPC health probe
        curl \
    ;

# Install pre-saved root CA for Amazon RDS (postgres).
# -- Use pre-saved AWS cert to avoid catch22 for zscalar building.
COPY build/certs/rds-ca-2019-root.pem /etc/ssl/certs/rds-ca-2019-root.pem

ARG GRPC_HEALTH_PROBE_VERSION=0.4.16
ARG GRPC_HEALTH_PROBE_CHECKSUM=c72704c9cd49fb18f8df28de26e29c2563b6515e38efee65a6c3a29ec7368a91

RUN set -eux; \
    # gRPC health probe for gRPC readiness checks
    curl \
        --location \
        --output grpc_health_probe \
        https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
    ; \
    echo "${GRPC_HEALTH_PROBE_CHECKSUM}  grpc_health_probe" | sha256sum -c; \
    chmod 755 grpc_health_probe;

# ------------------------------------------------------------------------------
# Base runtime image
# ------------------------------------------------------------------------------

# TODO: use `nonroot` over `debug-nonroot` when we no longer use env vars from vault
FROM gcr.io/distroless/static-debian11:debug-nonroot AS runtime

# Copy CA used for RDS certs
COPY --from=assets /etc/ssl/certs/rds-ca-2019-root.pem /etc/ssl/certs/rds-ca-2019-root.pem

# ------------------------------------------------------------------------------
# Base gRPC runtime image
# ------------------------------------------------------------------------------
FROM runtime AS runtime-grpc

COPY --from=assets /build/grpc_health_probe /usr/local/bin/grpc_health_probe

# ------------------------------------------------------------------------------
# Service REST Runtime Image
# ------------------------------------------------------------------------------

FROM runtime AS {{cookiecutter.service_name}}-rest

COPY --from=build /build/dist/{{cookiecutter.service_name}}-rest /usr/local/bin/{{cookiecutter.service_name}}-rest

ENTRYPOINT ["/usr/local/bin/{{cookiecutter.service_name}}-rest"]

# ------------------------------------------------------------------------------
# Service gRPC Runtime Image
# ------------------------------------------------------------------------------

FROM runtime AS {{cookiecutter.service_name}}-grpc

COPY --from=build /build/dist/{{cookiecutter.service_name}}-grpc /usr/local/bin/{{cookiecutter.service_name}}-grpc

ENTRYPOINT ["/usr/local/bin/{{cookiecutter.service_name}}-grpc"]

# ------------------------------------------------------------------------------
# DB Migration Runtime Image
# ------------------------------------------------------------------------------

FROM runtime AS {{cookiecutter.service_name}}-db-migrator

COPY --from=builder /build/dist/{{cookiecutter.service_name}}-db-migrator /usr/local/bin/{{cookiecutter.service_name}}-db-migrator

ENTRYPOINT ["/usr/local/bin/{{cookiecutter.service_name}}-db-migrator"]
