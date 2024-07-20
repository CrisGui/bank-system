FROM --platform=$BUILDPLATFORM python:3.12.4-alpine AS builder

WORKDIR /code

COPY requirements.txt /code
RUN python -m pip install -r requirements.txt

COPY . ./core

CMD [ "python", "-m", "core" ]

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

CMD [ "python", "-m", "core" ]
