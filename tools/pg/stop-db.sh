#!/bin/sh

if [ -e data/postmaster.pid ]; then
    echo "STOP DB"
    PID=`head -n 1 data/postmaster.pid`
    echo $PID >> "TRACE.log"
    kill -INT $PID
fi

