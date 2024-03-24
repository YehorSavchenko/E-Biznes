# syntax=docker/dockerfile:1
FROM ubuntu:22.04

#install app dependecies
RUN apt-get update && \
    apt-get install -y openjdk-8-jdk wget unzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64

RUN cd /tmp && \
    wget https://github.com/JetBrains/kotlin/releases/download/v1.9.23/kotlin-compiler-1.9.23.zip && \
    unzip kotlin-compiler-*.zip -d /opt && \
    rm kotlin-compiler-*.zip

ENV PATH $PATH:/opt/kotlinc/bin

RUN wget https://services.gradle.org/distributions/gradle-8.7-bin.zip -P /tmp && \
    unzip -d /opt/gradle /tmp/gradle-*.zip && \
    rm /tmp/gradle-*.zip

ENV PATH $PATH:/opt/gradle/gradle-8.7/bin

COPY build.gradle /home/gradle/project/

WORKDIR /home/gradle/project

RUN gradle build

CMD ["sh", "-c", "java -version && kotlinc -version && gradle --version"]