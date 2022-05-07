import datetime
import os 
import psutil
import time

datadir = os.path.expanduser(os.path.join('~', 'Desktop'))

# Scan for a "Blockchains" directory on linux systems, use it if it exists
if os.path.exists(os.path.join('.', 'media', os.getlogin(), 'Blockchains')[1:]):
	datadir = os.path.join('.', 'media', os.getlogin(), 'Blockchains')[1:]

portNumber = 8545
if os.path.exists(os.path.join('datadir', f'mainnet-geth-{portNumber}-node')):
	datadir = os.path.join('datadir', f'mainnet-geth-{portNumber}-node')


def header():
	line = 'Timestamp,'
	line += 'Timestamp (Seconds),'
	line += 'CPU %,'
	line += 'CPU Frequency,'
	line += 'Virtual Memory %,'
	line += 'Virtual Memory,'
	line += 'Swap Memory %,'
	line += 'Swap Memory,'
	line += 'Disk Usage %,'
	line += 'Disk Usage,'
	return line

def log(file):
	cpu = psutil.cpu_percent()
	cpu_f = 0
	try:
		cpu_f = psutil.cpu_freq().current
	except: pass
	v_mem_p = psutil.virtual_memory().percent
	v_mem = psutil.virtual_memory().used
	s_mem_p = psutil.swap_memory().percent
	s_mem = psutil.swap_memory().used
	disk_usage_p = psutil.disk_usage('/').percent
	disk_usage = psutil.disk_usage('/').used

	now = datetime.datetime.now()
	time_end = (now - datetime.datetime(1970, 1, 1)).total_seconds()

	line = f'{now},{time_end},{cpu},{cpu_f},{v_mem_p},{v_mem},{s_mem_p},{s_mem},{disk_usage_p},{disk_usage},'
	file.write(line + '\n')

def run(file):
	count = 0
	try:
		while True:
			count += 1
			log(file)
			if count % 3600 == 0:
				print(f'Logged {count / 3600} hours')
			time.sleep(1)
	except KeyboardInterrupt:
		pass


fileName = os.path.join(datadir, 'LOGGED_CPU.csv')
file = open(fileName, 'w+')
file.write(header() + '\n')
run(file)
