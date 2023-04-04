#!/bin/bash
set -eou pipefail

cur=$PWD
for item in "$cur"/*/
do
    echo "$item"
    cd "$item"
    go test tinyGorm/... 2>&1 | grep -v warning
done