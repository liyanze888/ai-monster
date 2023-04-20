#!/bin/sh

set -e
mkdir -p internal/generated/dsl
MYSQL="${MYSQL:-admin:Admin123@tcp(dev.cluster-cvqute6nrbz7.ap-southeast-1.rds.amazonaws.com:3306)/monster_base_backend}"
sqlingo-gen-mysql $MYSQL > internal/generated/dsl/dsl.go

git config url."git@gitlab.wuren.com:".insteadOf https://gitlab.wuren.com/
