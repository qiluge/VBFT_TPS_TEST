# VBFT_TPS_TEST

this project is used to test ontology vbft tps.

filename|description
------|--------|
jmeter-support-jar/*.jar|support jar in jmeter
result/ontology.log|ontology vbft log file, log level is 2
result/VBFT-test report | vbft test report
Test.jmx|jmeter test configure file
run.sh|test start file
transfer-from-init-addr/| ontology init ont is saved by multi-signature address from init 7 wallet, we should withdraw ont to a specific account
transfer-from-init-addr/wallets/| all bookkeeper's wallet 
transfer-from-init-addr/params/TransferOntMultiSign.json| config init transfer, PATH1 is all bookkeeper's wallet path, PATH2 is transfer to account, amount is transfer amount
main.go|generate transaction data

### test steps
1. install jmeter
2. clone ontology and ontology-go-sdk
3. cp jmeter-support-jar/*.jar to jmeter-install-path/lib/ext/
4. cp Test.jmx, run.sh to jmeter-install-path/bin/
5. modify Test.jmx server attribute, set ip to yourself ontology ip
6. copy a bookkeeper's wallet.dat to jmeter-install-path/bin/, named wallet-admin.dat
7. generate a new wallet.dat, copy it to jmeter-install-path/bin/, named wallet-account.dat, record its’ address
8. modify main.go line 27, used address in step 7 replace the address in main.go
9. go build -o gen-transfer main.go, copy generated file 'gen-transfer' to jmeter-install-path/bin/
10. modify ip in transfer-from-init-addr/config_test.json to yourself ontology node ip
11. config init transfer configuration, wallet-admin.dat should be configed to transfer-from-init-addr/params/TransferOntMultiSign.json PATH2
12. start ontology vbft network
13. in jmeter-install-path/bin/, execute './run.sh 30'，input 7 bookkkeepers password
14. there will generate a transfer.txt, it contains 3,000,000 transactions; the jmeter will send them to ontology network
15. record test result.
