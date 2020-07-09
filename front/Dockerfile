FROM node:current-alpine as node

RUN apk add --update  build-base

COPY . /app
WORKDIR /app

RUN make deps \
  && make build

FROM vibioh/viws

COPY --from=node /app/dist/ /www/
