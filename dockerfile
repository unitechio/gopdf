# ===== STAGE 1: PDF TOOLS BUILDER (Build once, reuse forever) =====
FROM ubuntu:22.04 AS pdf-tools-builder

ENV DEBIAN_FRONTEND=noninteractive

# Install all PDF tools and dependencies
RUN apt-get update && apt-get install -y \
    poppler-utils \
    mupdf-tools \
    wkhtmltopdf \
    xvfb \
    fonts-noto \
    fonts-noto-cjk \
    fonts-dejavu-core \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create tools directory structure for volume mounting
RUN mkdir -p /pdf-tools/bin /pdf-tools/lib /pdf-tools/fonts /pdf-tools/share

# Copy binaries
RUN cp /usr/bin/pdftotext /pdf-tools/bin/ && \
    cp /usr/bin/mutool /pdf-tools/bin/ && \
    cp /usr/bin/wkhtmltopdf /pdf-tools/bin/ && \
    cp /usr/bin/xvfb-run /pdf-tools/bin/

# Copy required libraries
RUN cp -r /lib/x86_64-linux-gnu/* /pdf-tools/lib/ 2>/dev/null || true && \
    cp -r /usr/lib/x86_64-linux-gnu/* /pdf-tools/lib/ 2>/dev/null || true

# Copy fonts
RUN cp -r /usr/share/fonts/* /pdf-tools/fonts/ && \
    cp -r /etc/fonts /pdf-tools/share/

# Create wkhtmltopdf wrapper
RUN echo '#!/bin/bash\nLD_LIBRARY_PATH=/pdf-tools/lib xvfb-run -a --server-args="-screen 0 1280x1024x24" /pdf-tools/bin/wkhtmltopdf "$@"' > /pdf-tools/bin/wkhtmltopdf-headless && \
    chmod +x /pdf-tools/bin/wkhtmltopdf-headless

# ===== STAGE 2: GO APP BUILDER =====
FROM golang:1.23-alpine AS go-builder

WORKDIR /app
COPY go.mod go.sum ./
ENV GOTOOLCHAIN=auto
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s' \
    -o app ./cmd

# ===== STAGE 3: ULTRA-MINIMAL RUNTIME =====
FROM gcr.io/distroless/static-debian12:nonroot

# Copy only the Go binary
COPY --from=go-builder /app/app /app

# Environment variables for volume-mounted tools
ENV PATH="/pdf-tools/bin:$PATH"
ENV LD_LIBRARY_PATH="/pdf-tools/lib"
ENV FONTCONFIG_PATH="/pdf-tools/share/fonts"
ENV DISPLAY=:99
ENV WKHTMLTOPDF_PATH="/pdf-tools/bin/wkhtmltopdf"
ENV PDFTOTEXT_PATH="/pdf-tools/bin/pdftotext"
ENV MUTOOL_PATH="/pdf-tools/bin/mutool"

# Create mount points
USER 0
RUN mkdir -p /pdf-tools
USER nonroot

EXPOSE 8080
ENTRYPOINT ["/app"]

# Build instructions in comments:
# docker build -t pdf-tools-volume --target pdf-tools-builder .
# docker build -t gopdf-minimal .
# docker create --name pdf-tools-container pdf-tools-volume
# docker cp pdf-tools-container:/pdf-tools ./pdf-tools-volume
# docker rm pdf-tools-container