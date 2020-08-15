#!/bin/bash

TMPDIR=$(mktemp -d)
BRANCH=$(date +%Y%m%d)
CANDIDATE=$(date +%Y%m%d-%H%M)

{
        git clone git@github.com:johnwtracy/personal.git ${TMPDIR}
        cd ${TMPDIR}
        git checkout -B $BRANCH
        git tag $CANDIDATE refs/heads/$BRANCH
        git push origin $BRANCH --tags
} > /dev/stderr

echo $CANDIDATE
