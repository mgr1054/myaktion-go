#!/usr/bin/env bash
export DB_CONNECT=mariadb:3306
docker run -d -p 3306:3306 --name database -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=myaktion mariadb:10.5
