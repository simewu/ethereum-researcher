#!/usr/bin/python

# python3 run.py [ARGS]

# ARGS:
# 	-ropsten	Run a light node in the Ropsten testnet
#	-local		Run a local/offline node [still under development]

import getopt
import os
import psutil    
import re
import subprocess
import sys
import time
import webbrowser


# GETH COMMANDS:
#	https://geth.ethereum.org/docs/interface/command-line-options
#	txpool			Show unconfirmed transactions
#	miner.start		Start mining
#	miner.stop		Stop mining
#	--jspath 'home directory/javascript'
#	loadScript('mineWhenNeeded.js')

global gethCmdHeader
gethCmdHeader = './build/bin/geth'

global portNumber
portNumber = 8545

global accountAddress, accountKeystorePath
accountAddress = ''

# Send commands to the terminal
def terminal(cmd):
	print(cmd)
	return os.popen(cmd).read()

def terminal_newwindow(cmd):
	subprocess.call(f'{cmd}', shell=True)

# # Run a geth command
# def geth(cmd, nodeDataDir=''):
# 	if nodeDataDir == '':
# 		return terminal(f'./build/bin/geth {cmd}')
# 	else:
# 		return terminal(f'./build/bin/geth -datadir '{nodeDataDir}' {cmd}')
def geth(cmd):
	global gethCmdHeader
	return terminal(f'{gethCmdHeader} {cmd}')

def geth_newwindow(cmd):
	global gethCmdHeader
	return terminal_newwindow(f'{gethCmdHeader} {cmd}')


# def isGethRunning():
# 	return len(terminal('ps faux | grep geth').strip()) > 0

# def stopGethNode():
# 	terminal('pkill -f geth')


# def getAccount():
# 	output = geth('account list')
# 	if output == '':
# 		print()
# 		print('Please first compile ethereum-researcher using:')
# 		print('\t./compile.sh')
# 		sys.exit()
# 	return output

# def getOrGenDataDir(nodeIndex):
# 	path = os.path.expanduser(f'~/Desktop/eth_node_{nodeIndex}')
# 	if not os.path.exists(path):
# 		os.makedirs(path)
# 		genesisFile = open(os.path.join(path, 'genesis.json'), 'w')
# 		genesisFile.write('')
# 	return path

# Returns address, keystorePath, or None if no address exists
def getAccount(datadir):
	try:
		response = terminal(f'./build/bin/geth -datadir="{datadir}" account list')
		match = re.match(r'Account #0: \{([0-9a-fA-F]+)\} keystore://(.*)', response)
		address = match.group(1)
		keystorePath = match.group(2).strip()
		return address, keystorePath
	except:
		return None, None


def createLocalGethDirectory(datadir):
	global accountAddress, accountKeystorePath

	if not os.path.exists(datadir):
		print('Creating datadir directory "datadir"...')
		os.makedirs(datadir)

	accountAddress, accountKeystorePath = getAccount(datadir)
	if accountAddress is None:
		print('Account does not exist, creating account...')
		passwordPath = os.path.expanduser(os.path.join('~', 'tempPassword.txt'))
		# Just make the default password "" for simplicity, security is not important for testing environments:
		terminal(f'echo "" > {passwordPath}')
		terminal(f'./build/bin/geth -datadir="{datadir}" account new --password "{passwordPath}"')
		terminal(f'rm -rf {passwordPath}')
		accountAddress, accountKeystorePath = getAccount(datadir)

	# Check if it worked...
	if accountAddress is None:
		print('Attempted to create an account, but failed.')
		print('Terminating program.')
		sys.exit()

	print('Account address:', accountAddress)

	genesisPath = os.path.join(datadir, 'genesis.json')
	genesisFile = open(genesisPath, 'w')
	genesisFile.write('''{
	"config": {
		"chainID": 1234,
		"homesteadBlock": 0,
		"eip150Block": 0,
		"eip155Block": 0,
		"eip158Block": 0
	},
	"alloc": {
		"0x''' + accountAddress + '''": {
			"balance": "100000000000000000000000000000"
		}
	},
	"difficulty": "0x4000",
	"gasLimit": "0xffffffff",
	"nonce": "0x0000000000000000",
	"coinbase": "0x0000000000000000000000000000000000000000",
	"mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"extraData": "0x123458db4e347b1234537c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa",
	"timestamp": "0x00"
}''')
	genesisFile.close()
	print('Genesis written! Initializing...')
	terminal(f'./build/bin/geth -datadir="{datadir}" init "{genesisPath}"')


def main(argv):
	global gethCmdHeader
	global portNumber
	global accountAddress, accountKeystorePath

	contractFileName = ''

	try:
		opts, args = getopt.getopt(argv, 'rlc:p:', ['ropsten', 'local', 'contract', 'port'])
		assert len(opts) >= 0, 'Invalid number of arguments'
	except Exception as e:
		print('ERROR:', e)
		print()
		print ('python3 run.py [arguments]')
		print('\t-r, --ropsten\t\t\tUse the ropsten testnet')
		print('\t-l, --local\t\t\tUse local blockchain')
		print('\t-c, --contract <path_to.sol>\tExecute a solidity contract file')
		print('\t-p, --port 8545\t\tSet the RPC port / geth instance number')
		sys.exit(2)

	# Default to local blockchain instance
	if len(opts) == 0:
		opts.append(('--local', ''))

	# if isGethRunning():
	# 	print('Stopping current geth instance...')
	# 	stopGethNode()
	# 	time.sleep(1)
	# 	if isGethRunning():
	# 		print('Stopping geth failed! Please make sure geth is not running.')
	# 		sys.exit()

	# Round 1: loop through arguments
	for opt, arg in opts:
		if opt in ('-p', '--port'):
			portNumber = arg

	# portNumber has been updated, update the geth command
	gethCmdHeader += f' -http --http.port {portNumber}'

	# Round 2: loop through arguments
	for opt, arg in opts:

		if opt in ('-r', '--ropsten'):
			gethCmdHeader += ' --syncmode "light"'
			gethCmdHeader += ' --ropsten'
		
		elif opt in ('-l', '--local'):
			datadir = os.path.expanduser(os.path.join('~', 'Desktop', f'local-geth-{portNumber}-node'))
			if not os.path.exists(os.path.join(datadir, 'genesis.json')):
				print(f'Creating directory "{datadir}"...')
				createLocalGethDirectory(datadir)

			gethCmdHeader += f' -datadir "{datadir}"'

		elif opt in ('-c', '--contract'):
			contractFileName = arg

	accountAddress, accountKeystorePath = getAccount(datadir)

	if contractFileName != '':
		print('Executing ' + contractFileName)
		print(geth('attach geth.ipc --exec "eth.blockNumber"'))
		#print(geth('attach --exec "eth.blockNumber"'))
	else:

		print('Starting console...')
		print()
		print(f'Address: {accountAddress}')
		print('Hosted at:')
		print()
		print(f'\t\t127.0.0.1:{portNumber}')
		print()
		time.sleep(3)


		if 'firefox' not in (p.name() for p in psutil.process_iter()):
			webbrowser.open('https://remix.ethereum.org/#optimize=false&runs=200&evmVersion=null&version=soljson-v0.8.7+commit.e28d00a7.js')


		# geth --http --http.corsdomain="https://remix.ethereum.org" --http.api web3,eth,debug,personal,net --vmdebug --datadir <path/to/local/folder/for/test/chain> --dev console
		geth_newwindow(' --http.corsdomain="https://remix.ethereum.org" --http.api web3,eth,debug,personal,net --vmdebug --dev console')





if __name__ == '__main__':
	main(sys.argv[1:])