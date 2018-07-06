#!/bin/bash

echo "== withdraw ont to one peer, input 7 bookkkeepers password =="
cd transfer-from-init-addr
./main -t TransferOntMultiSign
cd ..
echo "== start generating txns =="
./gen-transfer $1 > transfer.txt
echo `wc -l transfer.txt`
echo "== start testing =="
date
./jmeter -n -t Test.jmx
