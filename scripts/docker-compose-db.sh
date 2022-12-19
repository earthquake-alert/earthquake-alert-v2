#!/bin/bash

# MySQLをdocker-composeで起動するためのスクリプト
# 使用法（1例）:
#  # 起動する（バックグラウンド）
#  ./script/docker-compose-db.sh up -d
#
#  # ログを見る
#  ./script/docker-compose-db.sh logs
#
# 詳しくは https://docs.docker.com/engine/reference/commandline/compose/ を参照してください。

docker-compose -f docker-compose.db.yaml $@
