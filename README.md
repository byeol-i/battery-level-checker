# battery-level-checker

# Purpose

for check device's battery level in anywhere

# System diagram

WIP

That System is constructed in several kafka brokers and middleware which use HTTP transport

<img width="530" alt="image" src="https://user-images.githubusercontent.com/35767154/185879467-d1d8eb77-135a-46c5-a467-b3544597dabb.png">

# rest API

> https://aglide100.github.io/battery-level-checker/

# golang dependency

In Swaggo part

> https://github.com/swaggo/echo-swagger

> "github.com/swaggo/swag"

> github.com/labstack/echo/v4

In Kakfa part

> "gopkg.in/alecthomas/kingpin.v2"

> "github.com/Shopify/sarama"

Logger

> "go.uber.org/zap"
