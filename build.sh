#!/usr/bin/env bash
set -e

MODE="release"
USE_DOCKER=0
IMAGE="steamrt-go-builder"

for arg in "$@"; do
  case "$arg" in
    debug|release) MODE="$arg" ;;
    --docker) USE_DOCKER=1 ;;
  esac
done

if [ "$USE_DOCKER" = "1" ]; then
  if ! docker image inspect "$IMAGE" >/dev/null 2>&1; then
    echo "Image $IMAGE not found, building it..."
    docker build -t "$IMAGE" .
  fi

  docker run --rm \
    --platform linux/amd64 \
    --user "$(id -u):$(id -g)" \
    -v "$(pwd):/app" \
    -w /app \
    "$IMAGE" "$MODE"
  exit 0
fi

MANIFEST="assets/resetscore.pplugin"
TRANSLATIONS="assets/resetscore.yml"
RELEASE_MANIFEST=".github/release-please-manifest.json"

NAME=$(jq -r '.name' "$MANIFEST")
ENTRY=$(jq -r '.entry' "$MANIFEST")
FILENAME=$(basename "$ENTRY")

if [ -f "$RELEASE_MANIFEST" ]; then
  VERSION=$(jq -r '."."' "$RELEASE_MANIFEST")
else
  BASE_VERSION=$(jq -r '.version // "0.0.0"' "$MANIFEST")

  GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || true)
  if [ -n "$GIT_HASH" ]; then
    BUILD="+git.$GIT_HASH"
  else
    BUILD="+$(date +%Y%m%d)"
  fi

  VERSION="${BASE_VERSION}-dev${BUILD}"
fi

OUTDIR="build/$MODE"
mkdir -p "$OUTDIR"

echo "Building $NAME v$VERSION ($MODE)"

jq --arg version "$VERSION" \
'.version = $version' "$MANIFEST" > "$OUTDIR/$FILENAME.pplugin"

cp "$TRANSLATIONS" "$OUTDIR/resetscore.yml"

GOFLAGS=()
LDFLAGS=""

if [ "$MODE" = "debug" ]; then
  GOFLAGS+=(-tags=debug -gcflags=all='-N -l')
else
  LDFLAGS="-s -w"
  GOFLAGS+=(-trimpath)
fi

go build \
  -buildmode=c-shared \
  "${GOFLAGS[@]}" \
  -ldflags="$LDFLAGS" \
  -o "$OUTDIR/$FILENAME.so"

echo "Output:"
echo "$OUTDIR/$FILENAME.so"
echo "$OUTDIR/$FILENAME.pplugin"
echo "$OUTDIR/resetscore.yml"