FROM registry.access.redhat.com/ubi8/ubi:8.0

# Generic labels
LABEL Component="go" \
      Name="rabbitmq-go-poc" \
      Version="1.0" \
      Release="1"

# Labels consumed by OpenShift
LABEL io.k8s.description="A Go programming language image" \
      io.k8s.display-name="Go programming language image for DO288" \
      io.openshift.expose-services="8080:http"

# This label is used to categorize this image as a builder image in the
# OpenShift web console.
LABEL io.openshift.tags="go, golang"

# Sets the application folder and disables golang cache
ENV APP_ROOT=/opt/app-root \
    GOLANG_PACKAGE=/usr/lib/golang/src
# Install golang package
RUN   yum install -y --nodocs --disableplugin=subscription-manager golang git && \
      yum clean all --disableplugin=subscription-manager -y && \
      mkdir -p ${APP_ROOT} 

WORKDIR ${APP_ROOT}
COPY ./*.go .
# Permissions to run on OpenShift
RUN chgrp -R 0 ${APP_ROOT} ${GOLANG_PACKAGE} && \
    chmod -R 777 ${APP_ROOT} ${GOLANG_PACKAGE} && \
    git -C ${GOLANG_PACKAGE} clone http://github.com/streadway/amqp && \
    go build -o ./out/go-sample-app app.go && \
    go build -o ./out/publisher publisher.go && \
    go build -o ./out/consumer1-queue1 consumer_queue1.go && \
    go build -o ./out/consumer1-queue2 consumer_queue2.go && \
    go build -o ./out/consumer2-queue2 consumer_queue2.go && \
    go build -o ./out/nack-consumer-queue2 nack_consumer_queue2.go && \
    go build -o ./out/retry-consumer-queue2 retry_consumer_queue2.go

USER 0

EXPOSE 8080

CMD ["./out/go-sample-app"]
