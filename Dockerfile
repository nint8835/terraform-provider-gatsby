FROM golang:1.14-alpine AS builder

WORKDIR /build
COPY . /build
RUN ["go", "build", "-o", "terraform-provider-gatsby"]

FROM hashicorp/terraform:light
COPY --from=builder /build/terraform-provider-gatsby /bin/terraform-provider-gatsby