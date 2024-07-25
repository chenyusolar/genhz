#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=address_server
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}