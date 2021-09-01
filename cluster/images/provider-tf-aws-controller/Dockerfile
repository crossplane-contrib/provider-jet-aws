FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH
ARG TINI_VERSION
ENV USER_ID=1001

# Setup Terraform environment
ENV TERRAFORM_VERSION 1.0.5
ENV TERRAFORM_PROVIDER_AWS_VERSION 3.56.0
ENV PLUGIN_DIR /terraform/provider-mirror/registry.terraform.io/hashicorp/aws/${TERRAFORM_PROVIDER_AWS_VERSION}/linux_${ARCH}
ENV TF_CLI_CONFIG_FILE /terraform/.terraformrc
ENV TF_FORK 0

RUN mkdir -p ${PLUGIN_DIR}

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip /tmp
ADD https://releases.hashicorp.com/terraform-provider-aws/${TERRAFORM_PROVIDER_AWS_VERSION}/terraform-provider-aws_${TERRAFORM_PROVIDER_AWS_VERSION}_linux_${ARCH}.zip /tmp
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

RUN unzip /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip \
  && unzip /tmp/terraform-provider-aws_${TERRAFORM_PROVIDER_AWS_VERSION}_linux_${ARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/terraform-provider-aws_${TERRAFORM_PROVIDER_AWS_VERSION}_linux_${ARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

ADD provider /usr/local/bin/crossplane-tf-aws-provider

USER ${USER_ID}
EXPOSE 8080
ENTRYPOINT ["crossplane-tf-aws-provider"]
