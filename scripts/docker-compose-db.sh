#!/bin/bash

# MySQLをdocker-composeで起動するためのスクリプト
# 使用法:
#  ./script/docker-compose-db.sh up -d

docker-compose -f docker-compose.db.yaml $@
