#!/bin/sh

echo "START DB"
./bin/postgres -D ./data -d2 2> postgresql.log
