# syntax=docker/dockerfile:1

FROM openjdk:11
WORKDIR /app

RUN apt-get update -qq && \
    apt-get install -qq -y apt-transport-https curl gnupg && \
    echo "deb https://repo.scala-sbt.org/scalasbt/debian all main" | tee /etc/apt/sources.list.d/sbt.list && \
    echo "deb https://repo.scala-sbt.org/scalasbt/debian /" | tee /etc/apt/sources.list.d/sbt_old.list && \
    curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | apt-key add && \
    apt-get update -qq && \
    apt-get install sbt

COPY . /app

RUN sbt compile

EXPOSE 9000

CMD ["sbt", "run"]