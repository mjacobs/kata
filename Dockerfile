# syntax=docker/dockerfile:1

# Build stage. Runs on the native build platform and cross-compiles to the
# target platform via buildx's TARGETOS/TARGETARCH, so multi-arch builds need
# no QEMU. CGO is off, yielding a static binary that runs on distroless static.
FROM --platform=$BUILDPLATFORM golang:1.26.3-bookworm@sha256:386d475a660466863d9f8c766fec64d7fdad3edac2c6a05020c09534d71edb4b AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
ARG VERSION=dev
ARG COMMIT=unknown
ARG DATE=unknown

# -buildvcs=false because .dockerignore omits .git, so there is no VCS metadata
# to stamp; the version is injected via ldflags instead.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS="$TARGETOS" GOARCH="$TARGETARCH" \
    go build -trimpath -buildvcs=false \
    -ldflags="-s -w \
    -X go.kenn.io/kata/internal/version.Version=${VERSION} \
    -X go.kenn.io/kata/internal/version.Commit=${COMMIT} \
    -X go.kenn.io/kata/internal/version.BuildDate=${DATE}" \
    -o /out/kata ./cmd/kata

# Stage the data dir with a .keep file so the final-stage COPY reliably
# preserves the directory and its ownership across all BuildKit versions; a
# fresh volume mounted at KATA_HOME then inherits writable ownership for the
# nonroot uid.
RUN mkdir -p /data && touch /data/.keep

# Final stage. Distroless static nonroot: no shell, no package manager, runs
# as uid 65532. Pinned by digest.
FROM gcr.io/distroless/static-debian12:nonroot@sha256:d093aa3e30dbadd3efe1310db061a14da60299baff8450a17fe0ccc514a16639

COPY --from=build /out/kata /kata
COPY --from=build --chown=65532:65532 /data /data

ENV KATA_HOME=/data
# Conventional daemon TCP port. Informational only — the daemon binds it only
# when told to via --listen or config.toml (see README "Container image").
EXPOSE 7777
VOLUME ["/data"]

USER 65532:65532

ENTRYPOINT ["/kata"]
CMD ["daemon", "start"]
