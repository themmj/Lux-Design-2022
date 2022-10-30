#!/bin/bash

help() {
    echo "Compilation script for go agent."
    echo "NOTE: Script must be run from the directory it is located in!"
    echo "USAGE: ./compile.sh [OPTIONS]"
    echo "OPTIONS can be:"
    echo "  -b  --build-dir     : alternative build dir to use (default: build)"
    echo "  -h  --help          : print this help page"
}

abort() {
    echo "$1" 1>&2
    echo "Aborting..." 1>&2
    exit 1
}

[ -f "$PWD/compile.sh" ] || abort "script not running from within the project's directory"
[ -z "$(which go)" ] && abort "go must be installed"

build_dir="build"

while [[ $# -gt 0 ]]; do
    case $1 in
        -b|--build-dir)
            shift
            build_dir="$1"
            shift
            ;;
        -h|--help)
            help && exit 0
            ;;
    esac
done

mkdir -p $build_dir

cd src
go build -o ../$build_dir/agent.out
cd ..

