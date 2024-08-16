# Source for qemu-user-static >7.0
FROM --platform=$BUILDPLATFORM debian:bookworm-slim AS qemu_binaries

# hadolint ignore=DL3008
RUN apt-get update -qq \
 && apt-get install -qqy --no-install-recommends qemu-user-static

RUN mkdir -p /binaries && \
    if [ "$(uname -m)" = "aarch64" ]; then \
        cp /usr/bin/qemu-arm-static /binaries/; \
        cp /usr/bin/qemu-x86_64-static /binaries/; \
    else \
        cp /usr/bin/qemu-aarch64-static /binaries/; \
        cp /usr/bin/qemu-arm-static /binaries/; \
    fi

FROM --platform=$BUILDPLATFORM golang:1.23-bookworm AS builder
ARG ARMFLASH_PLUGIN_VERSION=1.1.1

# hadolint ignore=DL3008
RUN apt-get update -qq \
 && apt-get install -qqy --no-install-recommends \
  ca-certificates \
  git \
  unzip \
  xz-utils \
  curl \
  jq \
  libarchive-tools \
 && mkdir /build

WORKDIR /build

#install upx optimzer (latest)
RUN if [ "$(uname -m)" = "aarch64" ]; then \
        UPX_PACKAGE_URL=$(curl -s "https://api.github.com/repos/upx/upx/releases/latest" | jq -r '.assets[] | select(.name | contains("arm64_linux.tar.xz")) | .browser_download_url'); \
    else \
        UPX_PACKAGE_URL=$(curl -s "https://api.github.com/repos/upx/upx/releases/latest" | jq -r '.assets[] | select(.name | contains("amd64_linux.tar.xz")) | .browser_download_url'); \
    fi && \
    curl -sL "$UPX_PACKAGE_URL" | xz -d - | tar -x --strip-components=1 --wildcards "*upx"

# install packer (fixed version)
ENV PACKER_VERSION=1.11.2
RUN if [ "$(uname -m)" = "aarch64" ]; then PACKER_ARCH="arm64"; else PACKER_ARCH="amd64"; fi && \
    PACKER_PACKAGE_URL="https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_${PACKER_ARCH}.zip" && \ 
    curl -sL "$PACKER_PACKAGE_URL" -o /dev/stdout | bsdtar -xf- --to-stdout '*packer' > packer && \
    chmod +x packer && \
    ./upx packer

COPY . .

RUN go build -ldflags="-s -w -X 'version.Version=$ARMFLASH_PLUGIN_VERSION'" -o packer-plugin-armflash && \
    ./upx packer-plugin-armflash

FROM --platform=$BUILDPLATFORM debian:bookworm-slim

# Enable detailed logging
ENV PACKER_LOG=1

# Enable plugin / download caching for consecutive runs
ENV PACKER_ENV="/.packer/config"
ENV PACKER_PLUGIN_PATH="/.packer/plugins"
ENV PACKER_CACHE_DIR="/.packer/cache"

# hadolint ignore=DL3008
RUN apt-get update -qq \
 && apt-get install -qqy --no-install-recommends \
  ca-certificates \
  dosfstools \
  fdisk \
  gdisk \
  kpartx \
  libarchive-tools \
  parted \
  psmisc \
  qemu-utils \
  sudo \
  xz-utils \
 && mkdir -p /.packer ${PACKER_ENV} ${PACKER_PLUGIN_PATH} ${PACKER_CACHE_DIR} \
 && rm -rf /var/lib/apt/lists/*


#install plugin
WORKDIR /.packer

COPY --from=builder /build/packer /bin/packer
RUN --mount=type=bind,from=builder,source=/build,target=/build \
  /bin/packer plugins install --path /build/packer-plugin-armflash github.com/brederle/armflash
COPY --from=qemu_binaries /binaries/* /usr/bin/
COPY docker/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/bin/packer", "version"]
