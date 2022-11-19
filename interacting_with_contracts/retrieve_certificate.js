const myArgs = process.argv.slice(2);
if(myArgs.length != 1) {
	console.error('Need one argument, the ID.');
	process.exit(1);
}
const argID = myArgs[0];

const os = require("os");
const fs = require('fs');
//const solc = require('solc');
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));

fs.readFile('contract.json', 'utf8', (error, data) => {
	if(error) {
		console.error('Failed to read contract json', error);
		return;
	}
	let json = JSON.parse(data);
	let contractJSON = json['contracts']['contract.sol:Storage'];
	let abi = contractJSON['abi'];
	let code = '0x' + contractJSON.bin;

	let contractAddressPath = os.homedir() + '/Desktop/local-geth-8545-node/deployed_contract.txt';
	fs.readFile(contractAddressPath, 'utf8', (error, data) => {
		if(error) {
			console.error('Failed to read contract address at ', contractAddressPath, error);
			return;
		}
		let contractAddress = data.trim();
		let contract = new web3.eth.Contract(abi, contractAddress);
		let account = '';

		web3.eth.getAccounts().then((accounts) => {
			account = accounts[0];
			try {
				let password = '';
				web3.eth.personal.unlockAccount(account, password);
			} catch(e) {
				console.error(e);
			}

			contract.methods.retrieveKey(argID).call({
				from: account
			}).then((response) => {
				console.log(response)
			});

		});
	});
});

