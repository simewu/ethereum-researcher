sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update -y
sudo apt-get install -y solc

solc contract.sol --combined-json abi,asm,ast,bin,bin-runtime,devdoc,opcodes,srcmap,srcmap-runtime,userdoc > contract.json