#!/bin/bash

echo "== start generating txns =="
./gen-transfer $1 > transfer.txt
echo `wc -l transfer.txt`
echo "== start testing =="
date
./jmeter -n -t Test.jmx
