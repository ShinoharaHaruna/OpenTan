#!/usr/bin/env bash
set -eo pipefail

# Set the build variables
PROJECT_MODULE="OpenTan"
BINARY_NAME="OpenTan"
VERSION=${BUILD_VERSION:-"dev"}
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
SCRIPT_DIR=$(cd "$SCRIPT_DIR" && pwd)
if [[ -f "$SCRIPT_DIR/go.mod" ]]; then
    ROOT_DIR="$SCRIPT_DIR"
elif [[ -f "$SCRIPT_DIR/../go.mod" ]]; then
    ROOT_DIR=$(dirname "$SCRIPT_DIR")
else
    echo "❌ Error: go.mod not found in current or parent directories. Current ROOT_DIR: $ROOT_DIR"
    exit 1
fi
echo "▸ Setting the project root: ${ROOT_DIR}"
RELEASE_DIR="${ROOT_DIR}/release"
echo "▸ Setting the release directory: ${RELEASE_DIR}"
TARGET="${RELEASE_DIR}/${BINARY_NAME}"

# Create the release directory
mkdir -p "${RELEASE_DIR}"

# Get module path from go.mod
if [ -f "${ROOT_DIR}/go.mod" ]; then
    MODULE_PATH=$(head -n1 "${ROOT_DIR}/go.mod" | awk '{print $2}')
else
    echo "❌ Error: go.mod not found in project root. Detected ROOT_DIR: $ROOT_DIR"
    exit 1
fi

# Construct the parameter settings using local module path
LDFLAGS="-X '${MODULE_PATH}/cmd.version=${VERSION}' \
         -X '${MODULE_PATH}/cmd.commit=${COMMIT}' \
         -X '${MODULE_PATH}/cmd.buildDate=${BUILD_TIME}'"

# Start building
echo "▸ Building ${BINARY_NAME}..."
(cd "${ROOT_DIR}" && \
  go build -ldflags "${LDFLAGS}" -o "${TARGET}" main.go)

# Check if the binary exists
if [ ! -f "${TARGET}" ]; then
    echo "❌ Error: Build failed, binary not found"
    exit 1
fi

# Display build information
echo "✔ Build successful"
echo "→ Binary path: ${TARGET}"
echo "→ Version:     ${VERSION} (commit: ${COMMIT})"
echo "→ Build time:  ${BUILD_TIME}"
echo "→ Go version:  $(go version)"
