ARG ARMFLASH_VERSION=1.1.0

# Source for qemu-user-static >7.0
FROM public.ecr.aws/ubuntu/ubuntu:lunar as qemu_binaries

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

FROM golang:1.21-bullseye AS builder

# hadolint ignore=DL3008
RUN apt-get update -qq \
 && apt-get install -qqy --no-install-recommends \
  ca-certificates \
  git \
  unzip \
  upx-ucl \
  wget \
 && mkdir /build

WORKDIR /build

COPY . .

RUN go build -o packer-plugin-armflash_v${ARMFLASH_VERSION}

ENV PACKER_VERSION 1.11.2

RUN if [ "$(uname -m)" = "aarch64" ]; then PACKER_ARCH="arm64"; else PACKER_ARCH="amd64"; fi && \
  wget https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_${PACKER_ARCH}.zip -q -O /tmp/packer.zip && \
  unzip /tmp/packer.zip && \
  chmod +x packer && \
  rm /tmp/packer.zip && \
  upx-ucl --lzma packer

# COMPRESS WITH UPX
RUN 

FROM public.ecr.aws/lts/ubuntu:jammy

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
 && mkdir -p /.packer \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /.packer

COPY docker/entrypoint.sh /entrypoint.sh
COPY --from=builder /build/packer /bin/
COPY --from=qemu_binaries /binaries/* /usr/bin/

RUN --mount=type=bind,from=builder,source=/build \
  packer plugins install /build/packer-plugin-armflash_v${ARMFLASH_VERSION} github.com/brederle/armflash

# Enable detailed logging
ENV PACKER_LOG=1

# Enable plugin / download caching for consecutive runs
ENV PACKER_ENV="/.packer/config"
ENV PACKER_PLUGIN_PATH="/.packer/plugins"
ENV PACKER_CACHE_DIR="/.packer/cache"

ENTRYPOINT ["/entrypoint.sh"]
