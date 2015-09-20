#!/bin/sh

echo "INIT DB"
rm -rf ./data/
./bin/initdb --username lionel --locale=C --pgdata=./data
