FROM alpine:latest AS loader

ARG PLANTUML_VERSION

RUN apk add --no-cache \
  wget \
  git \
  ca-certificates

RUN wget \
  "https://github.com/plantuml/plantuml/releases/download/v1.2024.4/plantuml-1.2024.4.jar" \
  -O /opt/plantuml.jar

RUN git clone --branch v3.0.0 https://github.com/tupadr3/plantuml-icon-font-sprites.git /opt/plantuml-icon-font-sprites

FROM eclipse-temurin:17-jre-jammy

ENV LANG en_US.UTF-8

RUN apt-get update \
  && apt-get install --no-install-recommends -y \
    graphviz \
    fonts-dejavu \
  && apt-get autoremove \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY --from=loader /opt/plantuml.jar /opt/plantuml.jar
COPY --from=loader /opt/plantuml-icon-font-sprites /opt/plantuml-icon-font-sprites

WORKDIR /data

ENTRYPOINT ["java", "-Djava.awt.headless=true", "-jar", "/opt/plantuml.jar"]
