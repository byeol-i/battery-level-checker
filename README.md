# battery-level-checker

1. [System diagram](#System-diagram)
2. [User scenario](#User-scenario)
3. [rest API](#rest-API)
4. [Table Relationship](#Table-Relationship)
5. [Running](#Running)
6. [Performance Test](#Performance-Test)
7. [Dependency](#Dependency)

# Purpose

The purpose of this project is to hands on learning experience about Kafka while impl a battery level checking application

# System diagram

WIP

That System is constructed in several kafka brokers and middleware which use HTTP transport

<img width="663" alt="image" src="https://github.com/byeol-i/battery-level-checker/assets/35767154/ec27ca4e-f7d7-4c40-a1ba-0a8a17528137">

# User scenario

### 1. Register user

-   Getting token by firebase login

-   Create custom token (Using for generate access token)

-   Post to user date to server

### 2. Get devices

-   Added token at header

-   [HTTP]Get {domain}/api/v1/device

...

# rest API

For detailed information about the API, Please read this [documentation](https://byeol-i.github.io/battery-level-checker/)

# Table Relationship

<img width="497" alt="image" src="https://github.com/byeol-i/battery-level-checker/assets/35767154/2426a95a-f626-4d73-bc58-20c705234452">

# Running

To running this application, need to setup docker swarm system.

### 1. labeling worker at worker

> docker node update --label-add kafka=1 {hostname}

> docker node update --label-add kafka=2 {hostname}

> ...

### 2. Creating firebase api

> locate firebase key at ./conf/firebase/key.json

### 3. Running zookeeper & kafka

> docker stack deploy -c kafka.yml {stack name}

or

> cd shell && ./deployZoo.sh && ./deployKafka.sh

### 4. Running auth, apid ...etc

> cd shell && ./deployApp.sh

### or just cd shell && ./deploy.sh

<img src="https://github.com/byeol-i/battery-level-checker/assets/35767154/700e6836-a271-4927-aeed-62d3e318095c" width="1200">

# Performance Test

### Test environment

-   tested on rpi4 x 4 cluster
-   Used Docker swarm
-   Each service(apid, auth, consumer) scaled to 3 instance

<img width="1879" alt="Screenshot 2023-10-01 at 12 21 56 PM" src="https://github.com/byeol-i/battery-level-checker/assets/35767154/4223cbcd-08b6-4a6b-8a8b-1092c21f1f24">

### Test tool

-   test tool : jmeter
-   HTTP method : post
-   Thread : 2000
-   Ramp-up period : 1
-   Loop : 10

> Summary report

<img width="1143" alt="Screenshot 2023-10-01 at 12 20 26 PM" src="https://github.com/byeol-i/battery-level-checker/assets/35767154/6730f59d-e239-47d7-8ae8-7afa3530542d">

> Response time graph

<img width="1143" alt="Screenshot 2023-10-01 at 12 20 49 PM" src="https://github.com/byeol-i/battery-level-checker/assets/35767154/277fb5af-6c9d-40c7-a8eb-3d7494afa412">

# Dependency

The following dependencies are used in the project:

> [Echo swagger](https://github.com/swaggo/echo-swagger)

> [swag]("github.com/swaggo/swag")

> [echo]("github.com/labstack/echo/v4")

> [sarama]("github.com/Shopify/sarama")

> [zap]("go.uber.org/zap")

> [go-cache]("github.com/patrickmn/go-cache")

> [dateparse]("github.com/araddon/dateparse")
