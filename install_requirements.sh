#!/bin/bash

# Install some required packages.
sudo apt-get install pkg-config libgit2

echo "Follow the instructions in https://github.com/libgit2/git2go#from-next:"
echo "1) go get -d github.com/libgit2/git2go"
echo "2) cd $GOPATH/src/github.com/libgit2/git2go"
echo "3) git checkout next"
echo "3) git submodule update --init # get libgit2"
echo "4) make install"
