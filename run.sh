if ! wmctrl -l | grep -q "Custom Geth Console" ; then
	python3 run.py
else
	echo "Go Ethereum is already running!"
fi