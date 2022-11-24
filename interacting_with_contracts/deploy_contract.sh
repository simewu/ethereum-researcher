echo "Make sure you're using node version 16.15.1 (node -v)"

if ! command -v node &> /dev/null; then
	# sudo apt-get update -y
	# curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.2/install.sh | bash
	# wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.2/install.sh | bash
	# export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
	# [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm
	# nvm install 16.15.1
	exit 1
fi

	
# 	sudo apt-get install -y nodejs
# fi
# if ! command -v npm &> /dev/null; then
#  	sudo apt-get install -y npm
# 	npm install -g npm@latest
# 	npm install web3
# fi



node deploy_contract.js | grep -oP "Contract deployed to: \K.*" > ~/Desktop/local-geth-8545-node/deployed_contract.txt

cat ~/Desktop/local-geth-8545-node/deployed_contract.txt