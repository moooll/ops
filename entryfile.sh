#!/bin/bash

go get github.com/gobuffalo/pop/v5/...

sleep 10

soda migrate up -e docker

DB_ENV=docker /app/opsd
