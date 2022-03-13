#!/bin/bash

find . -type d -maxdepth 1 | while read line; do
  (cd $line; go build)
done
