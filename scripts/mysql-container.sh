#!/bin/bash

docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=changeme -d mysql:8