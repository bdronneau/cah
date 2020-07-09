FROM golang:1.13 as builder

ENV APP_NAME cah
ENV GO111MODULE on
ENV WORKDIR ${GOPATH}/src/github.com/bdronneau/cah/api

WORKDIR ${WORKDIR}
COPY ./ ${WORKDIR}/

RUN make ${APP_NAME} \
 && mkdir -p /app \
 && curl -s -o /app/cacert.pem https://curl.haxx.se/ca/cacert.pem \
 && cp bin/${APP_NAME} /app/

FROM scratch

ENV APP_NAME cah
EXPOSE 1323

# HEALTHCHECK --retries=10 CMD [ "/api", "-url", "https://localhost:1080/health" ]
ENTRYPOINT [ "/cah" ]

COPY --from=builder /app/${APP_NAME} /
