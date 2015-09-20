#!/bin/sh

if [ -e data/postmaster.pid ]; then
    touch data/MASTER
fi

