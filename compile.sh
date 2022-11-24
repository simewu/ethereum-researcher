command -v go version >/dev/null 2>&1 || {
	echo ""
	echo "Installing go..."
	git clone https://github.com/canha/golang-tools-install-script.git
	cd golang-tools-install-script/
	./goinstall.sh
	source ~/.bashrc
	cd ..
	rm -rf golang-tools-install-script
	echo
	echo "Please reboot, then re-run the script"
	exit 1
}

make all

echo
echo
echo "Go Ethereum compilation finished, starting...\n"

python3 run.py
