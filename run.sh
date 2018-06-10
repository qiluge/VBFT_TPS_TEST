#!/bin/bash

echo "== start generating txns =="
./gen-transfer > transfer.txt
for((i=0;i<$(($1-1));i++));do
./gen-transfer >> transfer.txt
done
echo `wc -l transfer.txt`
echo "== start testing =="
date
./jmeter -n -t Test.jmx
