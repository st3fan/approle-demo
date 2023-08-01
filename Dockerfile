# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/

FROM golang:1.20 AS build
WORKDIR /usr/local/src/approle-demo
COPY . .
RUN CGO_ENABLED=0 go build -v

FROM scratch
COPY --from=build /usr/local/src/approle-demo/approle-demo /usr/local/bin/approle-demo
ENTRYPOINT ["/usr/local/bin/approle-demo"]
