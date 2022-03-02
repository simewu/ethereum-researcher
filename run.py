run.py
import os
import sys


# Send commands to the terminal
def terminal(cmd):
	return os.popen(cmd).read()

# Run a geth command
def geth(cmd, nodeDataDir=''):
	if nodeDataDir == '':
		return terminal(f'./build/bin/geth {cmd}')
	else:
		return terminal(f'./build/bin/geth -datadir "{nodeDataDir}" {cmd}')

def getAccount():
	output = geth('account list')
	if output == '':
		print()
		print('Please first compile ethereum-researcher using:')
		print('\t./compile.sh')
		sys.exit()
	return output

def getOrGenDataDir(nodeIndex):
	path = os.path.expanduser(f'~/Desktop/eth_node_{nodeIndex}')
	if not os.path.exists(path):
		os.makedirs(path)
		genesisFile = open(os.path.join(path, 'genesis.json'), 'w')
		genesisFile.write('')
	return path

print(getAccount())
sys.exit()

nodeIndex = input('Enter a node index: ')
assert nodeIndex is int, 'NodeIndex is not an integer'
assert nodeIndex > 0, 'NodeIndex must be a positive integer'
nodeDataDir = getOrGenDataDir(nodeIndex)

geth('account list')
