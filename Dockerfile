# SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

#############      builder                          #############
FROM golang:1.23.3 AS builder

WORKDIR /go/src/github.com/gardener/machine-controller-manager-provider-azure
COPY . .

RUN .ci/build

#############      machine-controller               #############
FROM gcr.io/distroless/static-debian12:nonroot AS machine-controller
WORKDIR /

COPY --from=builder /go/src/github.com/gardener/machine-controller-manager-provider-azure/bin/rel/machine-controller /machine-controller
ENTRYPOINT ["/machine-controller"]
