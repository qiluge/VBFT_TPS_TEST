# VBFT_TPS_TEST

this project is used to test ontology vbft tps.

filename|description
------|--------|
jmeter-support-jar/*.jar|support jar in jmeter
result/ontology.log|ontology vbft log file, log level is 2
result/VBFT-test report | vbft test report
Test.jmx|jmeter test configure file
run.sh|test start file
main.go|generate transaction data

### test steps
1. install jmeter
2. clone ontology and ontology-go-sdk
3. cp jmeter-support-jar/*.jar to jmeter-install-path/lib/ext/
4. cp Test.jmx, run.sh to jmeter-install-path/bin/
5. modify Test.jmx server attribute, set ip to yourself ontology ip
6. copy a bookkeeper's wallet.dat to jmeter-install-path/bin/, named wallet-admin.dat
7. generate a new wallet.dat, copy it to jmeter-install-path/bin/, named wallet-account.dat, record its address
8. modify main.go line 37, used address in step 9 replace the address in main.go
9. go build -o gen-transfer main.go, copy generated file 'gen-transfer' to jmeter-install-path/bin/
10. start ontology vbft network
11. in jmeter-install-path/bin/, execute './run.sh 30'
12. there will generate a transfer.txt, it contains 3,000,000 transactions; the jmeter will send them to ontology network
13. record test result.
