#!/bin/sh

if [ -e data/postmaster.pid ]; then
    echo "STOP DB"
    kill -HUP `head -n 1 data/postmaster.pid`
fi

