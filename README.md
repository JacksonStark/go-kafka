# Go Kafka
go-kafka is a Go-based messaging queue playground, bootstrapped with a Kafka cluster of 3 brokers serving as the distributed pipeline.

## Getting Started

### Ensure latest version of Java is installed.
    https://www.oracle.com/java/technologies/javase-downloads.html

### Ensure Zookeeper server and broker servers 1 thru 3 are running. 
    bin/zookeeper-server-start.sh config/zookeeper.properties
    bin/kafka-server-start.sh config/server.1.properties
    bin/kafka-server-start.sh config/server.2.properties
    bin/kafka-server-start.sh config/server.3.properties

### Create the topic, point it to Zookeeper cluster, specify partitions, and set replication factor.
    bin/kafka-topics.sh --create --topic my-kafka-topic --zookeeper localhost:2181 --partitions 3 --replication-factor 2

## Technologies / Packages
- Go
- Apache Kafka
- Java
- kafka-go

## Contributions
  Playground created by:
- https://github.com/JacksonStark
