#!/bin/bash
cd "$(dirname "$0")"
find . -path ./Godeps -prune -o -name "*.go" -print | xargs goimports -w
