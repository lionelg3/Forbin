#!/usr/bin/env bash

VERSION=9.4.4

if [ -e tools/postgresql-$VERSION-bin.tar.gz ]; then
    echo "Postgresql $VERSION binary present"
    exit 0
fi

cd tools
if [ ! -e postgresql-$VERSION.tar.gz ]; then
    echo "Fetching Postgresql source $VERSION"
    wget https://ftp.postgresql.org/pub/source/v$VERSION/postgresql-$VERSION.tar.gz
    tar xvfz postgresql-$VERSION.tar.gz
fi

if [ ! -e postgresql-$VERSION/build ]; then
    mkdir -p postgresql-$VERSION/build
    cd postgresql-$VERSION/build
    ../configure --prefix=$PWD --exec-prefix=$PWD
    make
    make install
    #rm -rf GNUmakefile Makefile build config config.log config.status contrib doc src
    tar cvfZ ../../postgresql-$VERSION-bin.tar.gz bin include lib share
    cp ../../postgresql-$VERSION-bin.tar.gz ../../postgresql-bin.tar.gz
fi