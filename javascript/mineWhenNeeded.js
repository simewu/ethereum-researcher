// var Web3 = require('web3')
const web3 = new Web3('http://127.0.0.1:8545')

web3.eth.getBlockNumber(function (error, result) {
  console.log(result)
})

var mining_threads = 1

function checkWork() {
    if (web3.eth.getBlock('pending').transactions.length > 0) {
        if (web3.eth.mining) return;
        console.log('== Pending transactions! Mining...');
        miner.start(mining_threads);
    } else {
        miner.stop();
        console.log('== No transactions! Mining stopped.');
    }
}

web3.eth.filter('latest', function(err, block) { checkWork(); });
web3.eth.filter('pending', function(err, block) { checkWork(); });

checkWork();