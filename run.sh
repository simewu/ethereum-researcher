datadir=~/eth_node1
mkdir -p $datadir

./build/bin/geth --datadir="$datadir" -verbosity 6 --ipcdisable --port 30301 --http.port 8101 console 2>> $datadir/01.log
