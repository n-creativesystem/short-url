#!/bin/bash

rm -f .run.env
ENV_FILES=$(find . -mindepth 1 -maxdepth 1 -name "*.env" -not -name ".test.env")
for f in $ENV_FILES; do
  cat $f >> .run.env
done
