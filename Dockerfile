FROM gcr.io/google.com/cloudsdktool/cloud-sdk:latest
COPY ./wli_start /wli_start
ENTRYPOINT ["/wli_start"]