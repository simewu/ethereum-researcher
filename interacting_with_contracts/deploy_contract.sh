node deploy_app.js | grep -oP "Contract deployed to: \K.*" > ~/Desktop/local-geth-8545-node/deployed_contract.txt

cat ~/Desktop/local-geth-8545-node/deployed_contract.txt