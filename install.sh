#!/bin/sh
set -e

REPO="MitchTheStonky/pear"
INSTALL_DIR="/usr/local/bin"

# Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  darwin) OS="darwin" ;;
  linux)  OS="linux" ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

# Detect arch
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64) ARCH="arm64" ;;
  arm64)   ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# Get latest release tag
LATEST="$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"//;s/".*//')"
if [ -z "$LATEST" ]; then
  echo "Failed to fetch latest release"
  exit 1
fi

TARBALL="pear_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${LATEST}/${TARBALL}"

echo "Downloading pear ${LATEST} for ${OS}/${ARCH}..."
TMPDIR="$(mktemp -d)"
curl -sL "$URL" -o "${TMPDIR}/${TARBALL}"
tar -xzf "${TMPDIR}/${TARBALL}" -C "$TMPDIR"

echo "Installing to ${INSTALL_DIR}/pear..."
sudo install -m 755 "${TMPDIR}/pear" "${INSTALL_DIR}/pear"
rm -rf "$TMPDIR"

echo ""
echo "pear installed successfully!"
echo ""
echo "Get started:"
echo "  pear init"
