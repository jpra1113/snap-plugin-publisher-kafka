FROM java:8-jre
MAINTAINER Intelsdi-x

RUN  wget -q -O - http://www.eu.apache.org/dist/kafka/0.8.2.1/kafka_2.10-0.8.2.1.tgz | tar -xzf - -C /opt

ENV PATH /opt/kafka_2.10-0.8.2.1/bin:$PATH

COPY ./conf /opt/kafka_2.10-0.8.2.1/config

EXPOSE 9092

CMD  ["kafka-server-start.sh", "/opt/kafka_2.10-0.8.2.1/config/server.properties"]
