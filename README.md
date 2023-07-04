# battery-level-checker

# Purpose

The purpose of this project is to hands on learning experience about Kafka while impl a battery level checking application

# System diagram

WIP

That System is constructed in several kafka brokers and middleware which use HTTP transport

<img width="330" alt="image" src="https://user-images.githubusercontent.com/35767154/185879467-d1d8eb77-135a-46c5-a467-b3544597dabb.png">

# Simple scenario

### 1. Register user

-   Getting token by firebase login

-   Create custom token (Using for generate access token)

-   Post to user date to server

### 2. Get devices

-   Added token at header

-   Get {domain}/api/v1/device

...

# rest API

For detailed information about the API, Please read this [documentation](https://byeol-i.github.io/battery-level-checker/)

# Running

To running this application, need to setup docker swarm system.

## 1. Running zookeeper & kafka

> docker stack deploy -c kafka.yml {stack name}

or

> cd shell && ./deployZoo.sh && ./deployKafka.sh

## 2. Running auth, apid ...etc

> cd shell && ./deployApp.sh

# Environment

## 1. labeling worker at worker

> docker node update --label-add kafka=1 {hostname}

> docker node update --label-add kafka=2 {hostname}

> ...

## 2. Creating firebase api

> locate firebase key at ./conf/firebase/key.json

# Dependency

The following dependencies are used in the project:

> [Echo swagger](https://github.com/swaggo/echo-swagger)

> [swag]("github.com/swaggo/swag")

> [echo]("github.com/labstack/echo/v4")

> [sarama]("github.com/Shopify/sarama")

> [zap]("go.uber.org/zap")
