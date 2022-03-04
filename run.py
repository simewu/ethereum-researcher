#!/usr/bin/python

# python3 run.py [ARGS]

# ARGS:
# 	-ropsten	Run a light node in the Ropsten testnet
#	-local		Run a local/offline node [still under development]

import os
import sys
import getopt

global gethCmdHeader

gethCmdHeader = './build/bin/geth --syncmode "light"'



# Send commands to the terminal
def terminal(cmd):
	print(cmd)
	return os.popen(cmd).read()

# # Run a geth command
# def geth(cmd, nodeDataDir=''):
# 	if nodeDataDir == '':
# 		return terminal(f'./build/bin/geth {cmd}')
# 	else:
# 		return terminal(f'./build/bin/geth -datadir '{nodeDataDir}' {cmd}')
def geth(cmd):
	global gethCmdHeader
	return terminal(f'{gethCmdHeader} {cmd}')

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


def main(argv):
	global gethCmdHeader

	contractFileName = ''

	try:
		opts, args = getopt.getopt(argv, 'rl:c:', ['local=', 'contract='])
		assert len(opts) > 0, 'Invalid number of arguments'
	except:
		print ('python3 run.py [arguments]')
		print('\t-r, --ropsten\t\t\tUse the ropsten testnet')
		print('\t-l, --local X\t\t\tUse local blockchain (index=X)')
		print('\t-c, --contract <path_to.sol>\tExecute a solidity contract file')
		sys.exit(2)

	# Loop through arguments
	for opt, arg in opts:
		if opt in ('-r', '--ropsten'):
			gethCmdHeader += ' --ropsten'
		
		elif opt in ('-l', '--local'):
			print('STILL UNDER DEVELOPMENT')
			sys.exit()

		elif opt in ('-c', '--contract'):
			contractFileName = arg

	if contractFileName != '':
		print('Executing ' + contractFileName)
	else:
		geth('console')
# print(cmd)
# terminal(cmd)
# sys.exit()

# nodeIndex = input('Enter a node index: ')
# assert nodeIndex is int, 'NodeIndex is not an integer'
# assert nodeIndex > 0, 'NodeIndex must be a positive integer'
# nodeDataDir = getOrGenDataDir(nodeIndex)

# geth('account list')


if __name__ == '__main__':
	main(sys.argv[1:])