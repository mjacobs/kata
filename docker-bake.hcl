# Build definition for the kata daemon image.
#
#   docker buildx bake image-local --load   # native single-arch, for dev
#   docker buildx bake image                # multi-arch, used by CI/release
#
# Version metadata is injected via build args; override with
#   VERSION=v1.2.3 COMMIT=abc1234 DATE=2026-05-27T00:00:00Z docker buildx bake image

variable "VERSION" {
  default = "dev"
}

variable "COMMIT" {
  default = "unknown"
}

variable "DATE" {
  default = "unknown"
}

# Populated by docker/metadata-action's generated bake file in CI (tags,
# labels, annotations). The empty stub keeps local bake runs working.
target "docker-metadata-action" {}

# Shared build inputs. Hidden (underscore prefix): never built directly.
target "_common" {
  context    = "."
  dockerfile = "Dockerfile"
  args = {
    VERSION = VERSION
    COMMIT  = COMMIT
    DATE    = DATE
  }
}

# Release image: multi-arch, attested, tags/labels from metadata-action.
target "image" {
  inherits  = ["_common", "docker-metadata-action"]
  platforms = ["linux/amd64", "linux/arm64"]
  attest = [
    "type=provenance,mode=max",
    "type=sbom",
  ]
}

# Dev/CI smoke build: the host's native platform, no attestations, a fixed
# local tag so it can be loaded into the engine and run.
target "image-local" {
  inherits = ["_common"]
  tags     = ["kata:local"]
}
