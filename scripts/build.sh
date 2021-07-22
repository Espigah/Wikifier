#!/usr/bin/env bash
set -e

if [ ! -z $1 ]; then
    ACTION=$1
fi

# If its dev mode, only build for ourself
if [ "$ACTION" == "dev" ]; then
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)
fi

# If its docker mode, only build linux-amd64
if [ "$ACTION" == "docker" ]; then
    XC_OS=linux
    XC_ARCH=amd64
fi

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

echo -n "## Recreate directory... "
rm -rf build && mkdir build
echo "OK"

# instruct to build statically linked binaries
export CGO_ENABLED=0

# build cmds
echo "## Build..."
for OS in $XC_OS; do
    for ARCH in $XC_ARCH; do
        if ([ $OS == "darwin" ] && ([ $ARCH == "386" ] || [ $ARCH == "arm" ])) ||
        ([ $OS == "windows" ] && [ $ARCH == "arm" ])
        then
            echo "-- Skipping to $OS-$ARCH "
            continue
        fi
        echo "## Building to $OS-$ARCH "
        GOOS=$OS GOARCH=$ARCH go build -ldflags "$LD_FLAGS" -o build/$OS-$ARCH/wikifier main.go
        upx -9 build/$OS-$ARCH/wikifier
        echo "### $OS-$ARCH finished"
    done
done

# If docker copy the binary to the root folder for dockerfile
if [ "$ACTION" == "docker" ]; then
    cp build/$XC_OS-$XC_ARCH/* ./
fi