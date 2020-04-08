FROM gcr.io/uber-container-tools/makisu:latest as origin

FROM alpine

COPY --from=origin /makisu-internal /makisu-internal

ENTRYPOINT [ "/makisu-internal/makisu" ]