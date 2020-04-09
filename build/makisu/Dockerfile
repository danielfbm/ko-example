FROM gcr.io/uber-container-tools/makisu:latest as origin

FROM ubuntu

COPY --from=origin /makisu-internal /makisu-internal

ENTRYPOINT [ "/bin/bash", "-c" ]
CMD [ "/bin/bash"]