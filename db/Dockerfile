FROM --platform=$BUILDPLATFORM alpine:3.20.1 AS builder
RUN apk --no-cache add sqlite

WORKDIR /db
COPY bank-db.sqlite /db

FROM build AS dev-envs
RUN <<EOF
apk update
apk add git
EOF

RUN <<EOF
addgroup -S docker
adduser -S --shell /bin/bash --ingroup docker vscode
EOF

COPY --from=gloursdocker/docker / /

CMD ["sqlite3", "/data/bank-db.sqlite"]
