FROM ljfranklin/terraform-resource:latest

RUN mkdir -p .terraform.d/plugins
COPY oci-provider/linux_amd64/ .terraform.d/plugins/linux_amd64/
COPY null-provider/linux_amd64/ .terraform.d/plugins/linux_amd64/
