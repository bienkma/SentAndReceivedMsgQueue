FROM ubuntu:16.04
MAINTAINER bienkma (bientd88@gmail.com)
RUN mkdir -p /usr/src
RUN \
  apt-get update && \
  apt-get install -y --no-install-recommends build-essential libfuse-dev libcurl4-openssl-dev automake libtool pkg-config libssl-dev wget tar git && \
  apt-get clean && \
  apt-get autoremove -y && \
  rm -rf /var/lib/{apt,dpkg,cache,log}/

ADD ./requirement/jdk-8u162-linux-x64.tar.gz /usr/src/jdk-8u162-linux-x64.tar.gz
RUN \
  echo "export JAVA_HOME='/usr/local/jdk-8u163-linux-x64'" >> /etc/profile && \
  echo "export PATH=$JAVA_HOME/bin:$JAVA_HOME/sbin:$PATH" >> /etc/profile && \
  rm /usr/local/jdk-8u163-linux-x64

RUN \
  wget http://mirrors.viethosting.com/apache/kafka/1.1.0/kafka_2.11-1.1.0.tgz -O /usr/src/kafka_2.11-1.1.0.tgz && \
  tar xvz -C /usr/local/kafka -f /usr/src/kafka_2.11-1.1.0.tgz && \
  wget http://mirrors.viethosting.com/apache/zookeeper/zookeeper-3.3.6/zookeeper-3.3.6.tar.gz -O /usr/src/zookeeper-3.3.6.tar.gz && \
  tar xvz -C /usr/local/zookeeper -f /usr/src/zookeeper-3.3.6.tar.gz && \
  rm /usr/src/kafka_2.11-1.1.0.tgz && \
  rm /usr/src/zookeeper-3.3.6.tar.gz

ADD ./requirement/kafka/server.properties /usr/local/kafka/config/server.properties
ADD ./requirement/zookeeper/zoo.conf /usr/local/zookeeper/conf/zoo.conf

RUN \
  /usr/local/zookeeper/bin/zkServer.sh start && \
  /usr/local/kafka/bin/kafka-server-start.sh -daemon /usr/local/kafka/conf/server.properties

RUN \
  source /etc/profile && \
  /usr/local/kafka/bin/kafka-topics.sh --create --zookeeper 127.0.0.1:2181 --replication-factor 1 --partitions 1 --topic chatQueue













