# Golang Service Template

Golang back-end service template. Using this template, you can get started with back-end projects quickly.  

|  Web Framework   |     ORM      |   Database Driver    | Configuration Manager |   Log Manager   |  API Documentation  |
|:----------------:|:------------:|:--------------------:|:---------------------:|:---------------:|:-------------------:|
| labstack/echo/v4 | gorm.io/gorm | gorm.io/driver/postgres |      spf13/viper      | sirupsen/logrus | swaggo/echo-swagger |

## Configuration

1. Create database and user in **postgresql**, and grant privileges to user.  

```MySQL
CREATE DATABASE buz;
CREATE USER 'foo'@'localhost' IDENTIFIED BY 'bar';
GRANT ALL PRIVILEGES ON buz.* TO 'foo'@'localhost';
```

2. Create a file named `conf.yaml` in the project's root directory (refer to `template.yaml`).


3. Build a docker container named `todoqueue`
