// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"sync/atomic"

	"strconv" // Cybersecurity Lab
	"time"    // Cybersecurity Lab

	"github.com/holiman/uint256"
	"github.com/simewu/ethereum-researcher/common"
	"github.com/simewu/ethereum-researcher/core/types"
	"github.com/simewu/ethereum-researcher/params"
	"golang.org/x/crypto/sha3"
)

type bytecodeInfoType struct {
	addNanoseconds            int64 `json:"number"`
	addCount                  int   `json:"number"`
	subNanoseconds            int64 `json:"number"`
	subCount                  int   `json:"number"`
	mulNanoseconds            int64 `json:"number"`
	mulCount                  int   `json:"number"`
	divNanoseconds            int64 `json:"number"`
	divCount                  int   `json:"number"`
	sdivNanoseconds           int64 `json:"number"`
	sdivCount                 int   `json:"number"`
	modNanoseconds            int64 `json:"number"`
	modCount                  int   `json:"number"`
	smodNanoseconds           int64 `json:"number"`
	smodCount                 int   `json:"number"`
	expNanoseconds            int64 `json:"number"`
	expCount                  int   `json:"number"`
	signExtendNanoseconds     int64 `json:"number"`
	signExtendCount           int   `json:"number"`
	notNanoseconds            int64 `json:"number"`
	notCount                  int   `json:"number"`
	ltNanoseconds             int64 `json:"number"`
	ltCount                   int   `json:"number"`
	gtNanoseconds             int64 `json:"number"`
	gtCount                   int   `json:"number"`
	sltNanoseconds            int64 `json:"number"`
	sltCount                  int   `json:"number"`
	sgtNanoseconds            int64 `json:"number"`
	sgtCount                  int   `json:"number"`
	eqNanoseconds             int64 `json:"number"`
	eqCount                   int   `json:"number"`
	isZeroNanoseconds         int64 `json:"number"`
	isZeroCount               int   `json:"number"`
	andNanoseconds            int64 `json:"number"`
	andCount                  int   `json:"number"`
	orNanoseconds             int64 `json:"number"`
	orCount                   int   `json:"number"`
	xorNanoseconds            int64 `json:"number"`
	xorCount                  int   `json:"number"`
	byteNanoseconds           int64 `json:"number"`
	byteCount                 int   `json:"number"`
	addmodNanoseconds         int64 `json:"number"`
	addmodCount               int   `json:"number"`
	mulmodNanoseconds         int64 `json:"number"`
	mulmodCount               int   `json:"number"`
	shlNanoseconds            int64 `json:"number"`
	shlCount                  int   `json:"number"`
	shrNanoseconds            int64 `json:"number"`
	shrCount                  int   `json:"number"`
	sarNanoseconds            int64 `json:"number"`
	sarCount                  int   `json:"number"`
	keccak256Nanoseconds      int64 `json:"number"`
	keccak256Count            int   `json:"number"`
	addressNanoseconds        int64 `json:"number"`
	addressCount              int   `json:"number"`
	balanceNanoseconds        int64 `json:"number"`
	balanceCount              int   `json:"number"`
	originNanoseconds         int64 `json:"number"`
	originCount               int   `json:"number"`
	callerNanoseconds         int64 `json:"number"`
	callerCount               int   `json:"number"`
	callValueNanoseconds      int64 `json:"number"`
	callValueCount            int   `json:"number"`
	callDataLoadNanoseconds   int64 `json:"number"`
	callDataLoadCount         int   `json:"number"`
	callDataSizeNanoseconds   int64 `json:"number"`
	callDataSizeCount         int   `json:"number"`
	callDataCopyNanoseconds   int64 `json:"number"`
	callDataCopyCount         int   `json:"number"`
	returnDataSizeNanoseconds int64 `json:"number"`
	returnDataSizeCount       int   `json:"number"`
	returnDataCopyNanoseconds int64 `json:"number"`
	returnDataCopyCount       int   `json:"number"`
	extCodeSizeNanoseconds    int64 `json:"number"`
	extCodeSizeCount          int   `json:"number"`
	codeSizeNanoseconds       int64 `json:"number"`
	codeSizeCount             int   `json:"number"`
	codeCopyNanoseconds       int64 `json:"number"`
	codeCopyCount             int   `json:"number"`
	extCodeCopyNanoseconds    int64 `json:"number"`
	extCodeCopyCount          int   `json:"number"`
	extCodeHashNanoseconds    int64 `json:"number"`
	extCodeHashCount          int   `json:"number"`
	gaspriceNanoseconds       int64 `json:"number"`
	gaspriceCount             int   `json:"number"`
	blockhashNanoseconds      int64 `json:"number"`
	blockhashCount            int   `json:"number"`
	coinbaseNanoseconds       int64 `json:"number"`
	coinbaseCount             int   `json:"number"`
	timestampNanoseconds      int64 `json:"number"`
	timestampCount            int   `json:"number"`
	numberNanoseconds         int64 `json:"number"`
	numberCount               int   `json:"number"`
	difficultyNanoseconds     int64 `json:"number"`
	difficultyCount           int   `json:"number"`
	randomNanoseconds         int64 `json:"number"`
	randomCount               int   `json:"number"`
	gasLimitNanoseconds       int64 `json:"number"`
	gasLimitCount             int   `json:"number"`
	popNanoseconds            int64 `json:"number"`
	popCount                  int   `json:"number"`
	mloadNanoseconds          int64 `json:"number"`
	mloadCount                int   `json:"number"`
	mstoreNanoseconds         int64 `json:"number"`
	mstoreCount               int   `json:"number"`
	mstore8Nanoseconds        int64 `json:"number"`
	mstore8Count              int   `json:"number"`
	sloadNanoseconds          int64 `json:"number"`
	sloadCount                int   `json:"number"`
	sstoreNanoseconds         int64 `json:"number"`
	sstoreCount               int   `json:"number"`
	jumpNanoseconds           int64 `json:"number"`
	jumpCount                 int   `json:"number"`
	jumpiNanoseconds          int64 `json:"number"`
	jumpiCount                int   `json:"number"`
	jumpdestNanoseconds       int64 `json:"number"`
	jumpdestCount             int   `json:"number"`
	pcNanoseconds             int64 `json:"number"`
	pcCount                   int   `json:"number"`
	msizeNanoseconds          int64 `json:"number"`
	msizeCount                int   `json:"number"`
	gasNanoseconds            int64 `json:"number"`
	gasCount                  int   `json:"number"`
	createNanoseconds         int64 `json:"number"`
	createCount               int   `json:"number"`
	create2Nanoseconds        int64 `json:"number"`
	create2Count              int   `json:"number"`
	callNanoseconds           int64 `json:"number"`
	callCount                 int   `json:"number"`
	callCodeNanoseconds       int64 `json:"number"`
	callCodeCount             int   `json:"number"`
	delegateCallNanoseconds   int64 `json:"number"`
	delegateCallCount         int   `json:"number"`
	staticCallNanoseconds     int64 `json:"number"`
	staticCallCount           int   `json:"number"`
	returnNanoseconds         int64 `json:"number"`
	returnCount               int   `json:"number"`
	revertNanoseconds         int64 `json:"number"`
	revertCount               int   `json:"number"`
	undefinedNanoseconds      int64 `json:"number"`
	undefinedCount            int   `json:"number"`
	stopNanoseconds           int64 `json:"number"`
	stopCount                 int   `json:"number"`
	selfdestructNanoseconds   int64 `json:"number"`
	selfdestructCount         int   `json:"number"`
	makeLogNanoseconds        int64 `json:"number"`
	makeLogCount              int   `json:"number"`
	push1Nanoseconds          int64 `json:"number"`
	push1Count                int   `json:"number"`
	makePushNanoseconds       int64 `json:"number"`
	makePushCount             int   `json:"number"`
	makeDupNanoseconds        int64 `json:"number"`
	makeDupCount              int   `json:"number"`
	makeSwapNanoseconds       int64 `json:"number"`
	makeSwapCount             int   `json:"number"`
}

var bytecodeInfoLog bytecodeInfoType

// Cybersecurity Lab
type Export struct{}

func (export Export) GetBytecodeInfoStats() string {
	var output string = "{"
	output += "\"addNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.addNanoseconds, 10) + "\""
	output += "\"addCount\":\"" + strconv.Itoa(bytecodeInfoLog.addCount) + "\""
	output += "\"subNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.subNanoseconds, 10) + "\""
	output += "\"subCount\":\"" + strconv.Itoa(bytecodeInfoLog.subCount) + "\""
	output += "\"mulNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.mulNanoseconds, 10) + "\""
	output += "\"mulCount\":\"" + strconv.Itoa(bytecodeInfoLog.mulCount) + "\""
	output += "\"divNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.divNanoseconds, 10) + "\""
	output += "\"divCount\":\"" + strconv.Itoa(bytecodeInfoLog.divCount) + "\""
	output += "\"sdivNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.sdivNanoseconds, 10) + "\""
	output += "\"sdivCount\":\"" + strconv.Itoa(bytecodeInfoLog.sdivCount) + "\""
	output += "\"modNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.modNanoseconds, 10) + "\""
	output += "\"modCount\":\"" + strconv.Itoa(bytecodeInfoLog.modCount) + "\""
	output += "\"smodNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.smodNanoseconds, 10) + "\""
	output += "\"smodCount\":\"" + strconv.Itoa(bytecodeInfoLog.smodCount) + "\""
	output += "\"expNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.expNanoseconds, 10) + "\""
	output += "\"expCount\":\"" + strconv.Itoa(bytecodeInfoLog.expCount) + "\""
	output += "\"signExtendNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.signExtendNanoseconds, 10) + "\""
	output += "\"signExtendCount\":\"" + strconv.Itoa(bytecodeInfoLog.signExtendCount) + "\""
	output += "\"notNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.notNanoseconds, 10) + "\""
	output += "\"notCount\":\"" + strconv.Itoa(bytecodeInfoLog.notCount) + "\""
	output += "\"ltNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.ltNanoseconds, 10) + "\""
	output += "\"ltCount\":\"" + strconv.Itoa(bytecodeInfoLog.ltCount) + "\""
	output += "\"gtNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.gtNanoseconds, 10) + "\""
	output += "\"gtCount\":\"" + strconv.Itoa(bytecodeInfoLog.gtCount) + "\""
	output += "\"sltNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.sltNanoseconds, 10) + "\""
	output += "\"sltCount\":\"" + strconv.Itoa(bytecodeInfoLog.sltCount) + "\""
	output += "\"sgtNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.sgtNanoseconds, 10) + "\""
	output += "\"sgtCount\":\"" + strconv.Itoa(bytecodeInfoLog.sgtCount) + "\""
	output += "\"eqNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.eqNanoseconds, 10) + "\""
	output += "\"eqCount\":\"" + strconv.Itoa(bytecodeInfoLog.eqCount) + "\""
	output += "\"isZeroNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.isZeroNanoseconds, 10) + "\""
	output += "\"isZeroCount\":\"" + strconv.Itoa(bytecodeInfoLog.isZeroCount) + "\""
	output += "\"andNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.andNanoseconds, 10) + "\""
	output += "\"andCount\":\"" + strconv.Itoa(bytecodeInfoLog.andCount) + "\""
	output += "\"orNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.orNanoseconds, 10) + "\""
	output += "\"orCount\":\"" + strconv.Itoa(bytecodeInfoLog.orCount) + "\""
	output += "\"xorNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.xorNanoseconds, 10) + "\""
	output += "\"xorCount\":\"" + strconv.Itoa(bytecodeInfoLog.xorCount) + "\""
	output += "\"byteNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.byteNanoseconds, 10) + "\""
	output += "\"byteCount\":\"" + strconv.Itoa(bytecodeInfoLog.byteCount) + "\""
	output += "\"addmodNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.addmodNanoseconds, 10) + "\""
	output += "\"addmodCount\":\"" + strconv.Itoa(bytecodeInfoLog.addmodCount) + "\""
	output += "\"mulmodNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.mulmodNanoseconds, 10) + "\""
	output += "\"mulmodCount\":\"" + strconv.Itoa(bytecodeInfoLog.mulmodCount) + "\""
	output += "\"shlNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.shlNanoseconds, 10) + "\""
	output += "\"shlCount\":\"" + strconv.Itoa(bytecodeInfoLog.shlCount) + "\""
	output += "\"shrNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.shrNanoseconds, 10) + "\""
	output += "\"shrCount\":\"" + strconv.Itoa(bytecodeInfoLog.shrCount) + "\""
	output += "\"sarNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.sarNanoseconds, 10) + "\""
	output += "\"sarCount\":\"" + strconv.Itoa(bytecodeInfoLog.sarCount) + "\""
	output += "\"keccak256Nanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.keccak256Nanoseconds, 10) + "\""
	output += "\"keccak256Count\":\"" + strconv.Itoa(bytecodeInfoLog.keccak256Count) + "\""
	output += "\"addressNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.addressNanoseconds, 10) + "\""
	output += "\"addressCount\":\"" + strconv.Itoa(bytecodeInfoLog.addressCount) + "\""
	output += "\"balanceNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.balanceNanoseconds, 10) + "\""
	output += "\"balanceCount\":\"" + strconv.Itoa(bytecodeInfoLog.balanceCount) + "\""
	output += "\"originNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.originNanoseconds, 10) + "\""
	output += "\"originCount\":\"" + strconv.Itoa(bytecodeInfoLog.originCount) + "\""
	output += "\"callerNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callerNanoseconds, 10) + "\""
	output += "\"callerCount\":\"" + strconv.Itoa(bytecodeInfoLog.callerCount) + "\""
	output += "\"callValueNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callValueNanoseconds, 10) + "\""
	output += "\"callValueCount\":\"" + strconv.Itoa(bytecodeInfoLog.callValueCount) + "\""
	output += "\"callDataLoadNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callDataLoadNanoseconds, 10) + "\""
	output += "\"callDataLoadCount\":\"" + strconv.Itoa(bytecodeInfoLog.callDataLoadCount) + "\""
	output += "\"callDataSizeNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callDataSizeNanoseconds, 10) + "\""
	output += "\"callDataSizeCount\":\"" + strconv.Itoa(bytecodeInfoLog.callDataSizeCount) + "\""
	output += "\"callDataCopyNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callDataCopyNanoseconds, 10) + "\""
	output += "\"callDataCopyCount\":\"" + strconv.Itoa(bytecodeInfoLog.callDataCopyCount) + "\""
	output += "\"returnDataSizeNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.returnDataSizeNanoseconds, 10) + "\""
	output += "\"returnDataSizeCount\":\"" + strconv.Itoa(bytecodeInfoLog.returnDataSizeCount) + "\""
	output += "\"returnDataCopyNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.returnDataCopyNanoseconds, 10) + "\""
	output += "\"returnDataCopyCount\":\"" + strconv.Itoa(bytecodeInfoLog.returnDataCopyCount) + "\""
	output += "\"extCodeSizeNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.extCodeSizeNanoseconds, 10) + "\""
	output += "\"extCodeSizeCount\":\"" + strconv.Itoa(bytecodeInfoLog.extCodeSizeCount) + "\""
	output += "\"codeSizeNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.codeSizeNanoseconds, 10) + "\""
	output += "\"codeSizeCount\":\"" + strconv.Itoa(bytecodeInfoLog.codeSizeCount) + "\""
	output += "\"codeCopyNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.codeCopyNanoseconds, 10) + "\""
	output += "\"codeCopyCount\":\"" + strconv.Itoa(bytecodeInfoLog.codeCopyCount) + "\""
	output += "\"extCodeCopyNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.extCodeCopyNanoseconds, 10) + "\""
	output += "\"extCodeCopyCount\":\"" + strconv.Itoa(bytecodeInfoLog.extCodeCopyCount) + "\""
	output += "\"extCodeHashNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.extCodeHashNanoseconds, 10) + "\""
	output += "\"extCodeHashCount\":\"" + strconv.Itoa(bytecodeInfoLog.extCodeHashCount) + "\""
	output += "\"gaspriceNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.gaspriceNanoseconds, 10) + "\""
	output += "\"gaspriceCount\":\"" + strconv.Itoa(bytecodeInfoLog.gaspriceCount) + "\""
	output += "\"blockhashNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.blockhashNanoseconds, 10) + "\""
	output += "\"blockhashCount\":\"" + strconv.Itoa(bytecodeInfoLog.blockhashCount) + "\""
	output += "\"coinbaseNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.coinbaseNanoseconds, 10) + "\""
	output += "\"coinbaseCount\":\"" + strconv.Itoa(bytecodeInfoLog.coinbaseCount) + "\""
	output += "\"timestampNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.timestampNanoseconds, 10) + "\""
	output += "\"timestampCount\":\"" + strconv.Itoa(bytecodeInfoLog.timestampCount) + "\""
	output += "\"numberNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.numberNanoseconds, 10) + "\""
	output += "\"numberCount\":\"" + strconv.Itoa(bytecodeInfoLog.numberCount) + "\""
	output += "\"difficultyNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.difficultyNanoseconds, 10) + "\""
	output += "\"difficultyCount\":\"" + strconv.Itoa(bytecodeInfoLog.difficultyCount) + "\""
	output += "\"randomNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.randomNanoseconds, 10) + "\""
	output += "\"randomCount\":\"" + strconv.Itoa(bytecodeInfoLog.randomCount) + "\""
	output += "\"gasLimitNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.gasLimitNanoseconds, 10) + "\""
	output += "\"gasLimitCount\":\"" + strconv.Itoa(bytecodeInfoLog.gasLimitCount) + "\""
	output += "\"popNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.popNanoseconds, 10) + "\""
	output += "\"popCount\":\"" + strconv.Itoa(bytecodeInfoLog.popCount) + "\""
	output += "\"mloadNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.mloadNanoseconds, 10) + "\""
	output += "\"mloadCount\":\"" + strconv.Itoa(bytecodeInfoLog.mloadCount) + "\""
	output += "\"mstoreNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.mstoreNanoseconds, 10) + "\""
	output += "\"mstoreCount\":\"" + strconv.Itoa(bytecodeInfoLog.mstoreCount) + "\""
	output += "\"mstore8Nanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.mstore8Nanoseconds, 10) + "\""
	output += "\"mstore8Count\":\"" + strconv.Itoa(bytecodeInfoLog.mstore8Count) + "\""
	output += "\"sloadNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.sloadNanoseconds, 10) + "\""
	output += "\"sloadCount\":\"" + strconv.Itoa(bytecodeInfoLog.sloadCount) + "\""
	output += "\"sstoreNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.sstoreNanoseconds, 10) + "\""
	output += "\"sstoreCount\":\"" + strconv.Itoa(bytecodeInfoLog.sstoreCount) + "\""
	output += "\"jumpNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.jumpNanoseconds, 10) + "\""
	output += "\"jumpCount\":\"" + strconv.Itoa(bytecodeInfoLog.jumpCount) + "\""
	output += "\"jumpiNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.jumpiNanoseconds, 10) + "\""
	output += "\"jumpiCount\":\"" + strconv.Itoa(bytecodeInfoLog.jumpiCount) + "\""
	output += "\"jumpdestNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.jumpdestNanoseconds, 10) + "\""
	output += "\"jumpdestCount\":\"" + strconv.Itoa(bytecodeInfoLog.jumpdestCount) + "\""
	output += "\"pcNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.pcNanoseconds, 10) + "\""
	output += "\"pcCount\":\"" + strconv.Itoa(bytecodeInfoLog.pcCount) + "\""
	output += "\"msizeNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.msizeNanoseconds, 10) + "\""
	output += "\"msizeCount\":\"" + strconv.Itoa(bytecodeInfoLog.msizeCount) + "\""
	output += "\"gasNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.gasNanoseconds, 10) + "\""
	output += "\"gasCount\":\"" + strconv.Itoa(bytecodeInfoLog.gasCount) + "\""
	output += "\"createNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.createNanoseconds, 10) + "\""
	output += "\"createCount\":\"" + strconv.Itoa(bytecodeInfoLog.createCount) + "\""
	output += "\"create2Nanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.create2Nanoseconds, 10) + "\""
	output += "\"create2Count\":\"" + strconv.Itoa(bytecodeInfoLog.create2Count) + "\""
	output += "\"callNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callNanoseconds, 10) + "\""
	output += "\"callCount\":\"" + strconv.Itoa(bytecodeInfoLog.callCount) + "\""
	output += "\"callCodeNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.callCodeNanoseconds, 10) + "\""
	output += "\"callCodeCount\":\"" + strconv.Itoa(bytecodeInfoLog.callCodeCount) + "\""
	output += "\"delegateCallNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.delegateCallNanoseconds, 10) + "\""
	output += "\"delegateCallCount\":\"" + strconv.Itoa(bytecodeInfoLog.delegateCallCount) + "\""
	output += "\"staticCallNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.staticCallNanoseconds, 10) + "\""
	output += "\"staticCallCount\":\"" + strconv.Itoa(bytecodeInfoLog.staticCallCount) + "\""
	output += "\"returnNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.returnNanoseconds, 10) + "\""
	output += "\"returnCount\":\"" + strconv.Itoa(bytecodeInfoLog.returnCount) + "\""
	output += "\"revertNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.revertNanoseconds, 10) + "\""
	output += "\"revertCount\":\"" + strconv.Itoa(bytecodeInfoLog.revertCount) + "\""
	output += "\"undefinedNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.undefinedNanoseconds, 10) + "\""
	output += "\"undefinedCount\":\"" + strconv.Itoa(bytecodeInfoLog.undefinedCount) + "\""
	output += "\"stopNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.stopNanoseconds, 10) + "\""
	output += "\"stopCount\":\"" + strconv.Itoa(bytecodeInfoLog.stopCount) + "\""
	output += "\"selfdestructNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.selfdestructNanoseconds, 10) + "\""
	output += "\"selfdestructCount\":\"" + strconv.Itoa(bytecodeInfoLog.selfdestructCount) + "\""
	output += "\"makeLogNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.makeLogNanoseconds, 10) + "\""
	output += "\"makeLogCount\":\"" + strconv.Itoa(bytecodeInfoLog.makeLogCount) + "\""
	output += "\"push1Nanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.push1Nanoseconds, 10) + "\""
	output += "\"push1Count\":\"" + strconv.Itoa(bytecodeInfoLog.push1Count) + "\""
	output += "\"makePushNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.makePushNanoseconds, 10) + "\""
	output += "\"makePushCount\":\"" + strconv.Itoa(bytecodeInfoLog.makePushCount) + "\""
	output += "\"makeDupNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.makeDupNanoseconds, 10) + "\""
	output += "\"makeDupCount\":\"" + strconv.Itoa(bytecodeInfoLog.makeDupCount) + "\""
	output += "\"makeSwapNanoseconds\":\"" + strconv.FormatInt(bytecodeInfoLog.makeSwapNanoseconds, 10) + "\""
	output += "\"makeSwapCount\":\"" + strconv.Itoa(bytecodeInfoLog.makeSwapCount) + "\""
	output += "}"

	return output
}

// Cybersecurity Lab: Measure duration
func opAdd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opAdd(pc, interpreter, scope)
	bytecodeInfoLog.addNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.addCount += 1
	return a, b
}
func _opAdd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Add(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSub(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSub(pc, interpreter, scope)
	bytecodeInfoLog.subNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.subCount += 1
	return a, b
}
func _opSub(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Sub(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMul(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMul(pc, interpreter, scope)
	bytecodeInfoLog.mulNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.mulCount += 1
	return a, b
}
func _opMul(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Mul(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opDiv(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opDiv(pc, interpreter, scope)
	bytecodeInfoLog.divNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.divCount += 1
	return a, b
}
func _opDiv(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Div(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSdiv(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSdiv(pc, interpreter, scope)
	bytecodeInfoLog.sdivNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.sdivCount += 1
	return a, b
}
func _opSdiv(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.SDiv(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMod(pc, interpreter, scope)
	bytecodeInfoLog.modNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.modCount += 1
	return a, b
}
func _opMod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Mod(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSmod(pc, interpreter, scope)
	bytecodeInfoLog.smodNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.smodCount += 1
	return a, b
}
func _opSmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.SMod(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opExp(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opExp(pc, interpreter, scope)
	bytecodeInfoLog.expNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.expCount += 1
	return a, b
}
func _opExp(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	base, exponent := scope.Stack.pop(), scope.Stack.peek()
	exponent.Exp(&base, exponent)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSignExtend(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSignExtend(pc, interpreter, scope)
	bytecodeInfoLog.signExtendNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.signExtendCount += 1
	return a, b
}
func _opSignExtend(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	back, num := scope.Stack.pop(), scope.Stack.peek()
	num.ExtendSign(num, &back)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opNot(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opNot(pc, interpreter, scope)
	bytecodeInfoLog.notNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.notCount += 1
	return a, b
}
func _opNot(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x := scope.Stack.peek()
	x.Not(x)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opLt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opLt(pc, interpreter, scope)
	bytecodeInfoLog.ltNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.ltCount += 1
	return a, b
}
func _opLt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Lt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opGt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opGt(pc, interpreter, scope)
	bytecodeInfoLog.gtNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.gtCount += 1
	return a, b
}
func _opGt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Gt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSlt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSlt(pc, interpreter, scope)
	bytecodeInfoLog.sltNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.sltCount += 1
	return a, b
}
func _opSlt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Slt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSgt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSgt(pc, interpreter, scope)
	bytecodeInfoLog.sgtNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.sgtCount += 1
	return a, b
}
func _opSgt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Sgt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opEq(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opEq(pc, interpreter, scope)
	bytecodeInfoLog.eqNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.eqCount += 1
	return a, b
}
func _opEq(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Eq(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opIszero(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opIszero(pc, interpreter, scope)
	bytecodeInfoLog.isZeroNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.isZeroCount += 1
	return a, b
}
func _opIszero(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x := scope.Stack.peek()
	if x.IsZero() {
		x.SetOne()
	} else {
		x.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opAnd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opAnd(pc, interpreter, scope)
	bytecodeInfoLog.andNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.andCount += 1
	return a, b
}
func _opAnd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.And(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opOr(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opOr(pc, interpreter, scope)
	bytecodeInfoLog.orNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.orCount += 1
	return a, b
}
func _opOr(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Or(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opXor(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opXor(pc, interpreter, scope)
	bytecodeInfoLog.xorNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.xorCount += 1
	return a, b
}
func _opXor(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Xor(&x, y)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opByte(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opByte(pc, interpreter, scope)
	bytecodeInfoLog.byteNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.byteCount += 1
	return a, b
}
func _opByte(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	th, val := scope.Stack.pop(), scope.Stack.peek()
	val.Byte(&th)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opAddmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opAddmod(pc, interpreter, scope)
	bytecodeInfoLog.addmodNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.addmodCount += 1
	return a, b
}
func _opAddmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y, z := scope.Stack.pop(), scope.Stack.pop(), scope.Stack.peek()
	if z.IsZero() {
		z.Clear()
	} else {
		z.AddMod(&x, &y, z)
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMulmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMulmod(pc, interpreter, scope)
	bytecodeInfoLog.mulmodNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.mulmodCount += 1
	return a, b
}
func _opMulmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x, y, z := scope.Stack.pop(), scope.Stack.pop(), scope.Stack.peek()
	z.MulMod(&x, &y, z)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSHL(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSHL(pc, interpreter, scope)
	bytecodeInfoLog.shlNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.shlCount += 1
	return a, b
}

// opSHL implements Shift Left
// The SHL instruction (shift left) pops 2 values from the stack, first arg1 and then arg2,
// and pushes on the stack arg2 shifted to the left by arg1 number of bits.
func _opSHL(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Note, second operand is left in the stack; accumulate result into it, and no need to push it afterwards
	shift, value := scope.Stack.pop(), scope.Stack.peek()
	if shift.LtUint64(256) {
		value.Lsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSHR(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSHR(pc, interpreter, scope)
	bytecodeInfoLog.shrNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.shrCount += 1
	return a, b
}

// opSHR implements Logical Shift Right
// The SHR instruction (logical shift right) pops 2 values from the stack, first arg1 and then arg2,
// and pushes on the stack arg2 shifted to the right by arg1 number of bits with zero fill.
func _opSHR(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Note, second operand is left in the stack; accumulate result into it, and no need to push it afterwards
	shift, value := scope.Stack.pop(), scope.Stack.peek()
	if shift.LtUint64(256) {
		value.Rsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSAR(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSAR(pc, interpreter, scope)
	bytecodeInfoLog.sarNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.sarCount += 1
	return a, b
}

// opSAR implements Arithmetic Shift Right
// The SAR instruction (arithmetic shift right) pops 2 values from the stack, first arg1 and then arg2,
// and pushes on the stack arg2 shifted to the right by arg1 number of bits with sign extension.
func _opSAR(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	shift, value := scope.Stack.pop(), scope.Stack.peek()
	if shift.GtUint64(256) {
		if value.Sign() >= 0 {
			value.Clear()
		} else {
			// Max negative shift: all bits set
			value.SetAllOne()
		}
		return nil, nil
	}
	n := uint(shift.Uint64())
	value.SRsh(value, n)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opKeccak256(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opKeccak256(pc, interpreter, scope)
	bytecodeInfoLog.keccak256Nanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.keccak256Count += 1
	return a, b
}
func _opKeccak256(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	offset, size := scope.Stack.pop(), scope.Stack.peek()
	data := scope.Memory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	if interpreter.hasher == nil {
		interpreter.hasher = sha3.NewLegacyKeccak256().(keccakState)
	} else {
		interpreter.hasher.Reset()
	}
	interpreter.hasher.Write(data)
	interpreter.hasher.Read(interpreter.hasherBuf[:])

	evm := interpreter.evm
	if evm.Config.EnablePreimageRecording {
		evm.StateDB.AddPreimage(interpreter.hasherBuf, data)
	}

	size.SetBytes(interpreter.hasherBuf[:])
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opAddress(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opAddress(pc, interpreter, scope)
	bytecodeInfoLog.addressNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.addressCount += 1
	return a, b
}
func _opAddress(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetBytes(scope.Contract.Address().Bytes()))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opBalance(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opBalance(pc, interpreter, scope)
	bytecodeInfoLog.balanceNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.balanceCount += 1
	return a, b
}
func _opBalance(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	slot := scope.Stack.peek()
	address := common.Address(slot.Bytes20())
	slot.SetFromBig(interpreter.evm.StateDB.GetBalance(address))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opOrigin(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opOrigin(pc, interpreter, scope)
	bytecodeInfoLog.originNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.originCount += 1
	return a, b
}
func _opOrigin(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetBytes(interpreter.evm.Origin.Bytes()))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCaller(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCaller(pc, interpreter, scope)
	bytecodeInfoLog.callerNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callerCount += 1
	return a, b
}
func _opCaller(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetBytes(scope.Contract.Caller().Bytes()))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCallValue(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCallValue(pc, interpreter, scope)
	bytecodeInfoLog.callValueNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callValueCount += 1
	return a, b
}
func _opCallValue(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v, _ := uint256.FromBig(scope.Contract.value)
	scope.Stack.push(v)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCallDataLoad(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCallDataLoad(pc, interpreter, scope)
	bytecodeInfoLog.callDataLoadNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callDataLoadCount += 1
	return a, b
}
func _opCallDataLoad(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	x := scope.Stack.peek()
	if offset, overflow := x.Uint64WithOverflow(); !overflow {
		data := getData(scope.Contract.Input, offset, 32)
		x.SetBytes(data)
	} else {
		x.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCallDataSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCallDataSize(pc, interpreter, scope)
	bytecodeInfoLog.callDataSizeNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callDataSizeCount += 1
	return a, b
}
func _opCallDataSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetUint64(uint64(len(scope.Contract.Input))))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCallDataCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCallDataCopy(pc, interpreter, scope)
	bytecodeInfoLog.callDataCopyNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callDataCopyCount += 1
	return a, b
}
func _opCallDataCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var (
		memOffset  = scope.Stack.pop()
		dataOffset = scope.Stack.pop()
		length     = scope.Stack.pop()
	)
	dataOffset64, overflow := dataOffset.Uint64WithOverflow()
	if overflow {
		dataOffset64 = 0xffffffffffffffff
	}
	// These values are checked for overflow during gas cost calculation
	memOffset64 := memOffset.Uint64()
	length64 := length.Uint64()
	scope.Memory.Set(memOffset64, length64, getData(scope.Contract.Input, dataOffset64, length64))

	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opReturnDataSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opReturnDataSize(pc, interpreter, scope)
	bytecodeInfoLog.returnDataSizeNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.returnDataSizeCount += 1
	return a, b
}
func _opReturnDataSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetUint64(uint64(len(interpreter.returnData))))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opReturnDataCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opReturnDataCopy(pc, interpreter, scope)
	bytecodeInfoLog.returnDataCopyNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.returnDataCopyCount += 1
	return a, b
}
func _opReturnDataCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var (
		memOffset  = scope.Stack.pop()
		dataOffset = scope.Stack.pop()
		length     = scope.Stack.pop()
	)

	offset64, overflow := dataOffset.Uint64WithOverflow()
	if overflow {
		return nil, ErrReturnDataOutOfBounds
	}
	// we can reuse dataOffset now (aliasing it for clarity)
	var end = dataOffset
	end.Add(&dataOffset, &length)
	end64, overflow := end.Uint64WithOverflow()
	if overflow || uint64(len(interpreter.returnData)) < end64 {
		return nil, ErrReturnDataOutOfBounds
	}
	scope.Memory.Set(memOffset.Uint64(), length.Uint64(), interpreter.returnData[offset64:end64])
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opExtCodeSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opExtCodeSize(pc, interpreter, scope)
	bytecodeInfoLog.extCodeSizeNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.extCodeSizeCount += 1
	return a, b
}
func _opExtCodeSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	slot := scope.Stack.peek()
	slot.SetUint64(uint64(interpreter.evm.StateDB.GetCodeSize(slot.Bytes20())))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCodeSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCodeSize(pc, interpreter, scope)
	bytecodeInfoLog.codeSizeNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.codeSizeCount += 1
	return a, b
}
func _opCodeSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	l := new(uint256.Int)
	l.SetUint64(uint64(len(scope.Contract.Code)))
	scope.Stack.push(l)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCodeCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCodeCopy(pc, interpreter, scope)
	bytecodeInfoLog.codeCopyNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.codeCopyCount += 1
	return a, b
}
func _opCodeCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var (
		memOffset  = scope.Stack.pop()
		codeOffset = scope.Stack.pop()
		length     = scope.Stack.pop()
	)
	uint64CodeOffset, overflow := codeOffset.Uint64WithOverflow()
	if overflow {
		uint64CodeOffset = 0xffffffffffffffff
	}
	codeCopy := getData(scope.Contract.Code, uint64CodeOffset, length.Uint64())
	scope.Memory.Set(memOffset.Uint64(), length.Uint64(), codeCopy)

	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opExtCodeCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opExtCodeCopy(pc, interpreter, scope)
	bytecodeInfoLog.extCodeCopyNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.extCodeCopyCount += 1
	return a, b
}
func _opExtCodeCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var (
		stack      = scope.Stack
		a          = stack.pop()
		memOffset  = stack.pop()
		codeOffset = stack.pop()
		length     = stack.pop()
	)
	uint64CodeOffset, overflow := codeOffset.Uint64WithOverflow()
	if overflow {
		uint64CodeOffset = 0xffffffffffffffff
	}
	addr := common.Address(a.Bytes20())
	codeCopy := getData(interpreter.evm.StateDB.GetCode(addr), uint64CodeOffset, length.Uint64())
	scope.Memory.Set(memOffset.Uint64(), length.Uint64(), codeCopy)

	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opExtCodeHash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opExtCodeHash(pc, interpreter, scope)
	bytecodeInfoLog.extCodeHashNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.extCodeHashCount += 1
	return a, b
}

// opExtCodeHash returns the code hash of a specified account.
// There are several cases when the function is called, while we can relay everything
// to `state.GetCodeHash` function to ensure the correctness.
//   (1) Caller tries to get the code hash of a normal contract account, state
// should return the relative code hash and set it as the result.
//
//   (2) Caller tries to get the code hash of a non-existent account, state should
// return common.Hash{} and zero will be set as the result.
//
//   (3) Caller tries to get the code hash for an account without contract code,
// state should return emptyCodeHash(0xc5d246...) as the result.
//
//   (4) Caller tries to get the code hash of a precompiled account, the result
// should be zero or emptyCodeHash.
//
// It is worth noting that in order to avoid unnecessary create and clean,
// all precompile accounts on mainnet have been transferred 1 wei, so the return
// here should be emptyCodeHash.
// If the precompile account is not transferred any amount on a private or
// customized chain, the return value will be zero.
//
//   (5) Caller tries to get the code hash for an account which is marked as suicided
// in the current transaction, the code hash of this account should be returned.
//
//   (6) Caller tries to get the code hash for an account which is marked as deleted,
// this account should be regarded as a non-existent account and zero should be returned.
func _opExtCodeHash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	slot := scope.Stack.peek()
	address := common.Address(slot.Bytes20())
	if interpreter.evm.StateDB.Empty(address) {
		slot.Clear()
	} else {
		slot.SetBytes(interpreter.evm.StateDB.GetCodeHash(address).Bytes())
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opGasprice(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opGasprice(pc, interpreter, scope)
	bytecodeInfoLog.gaspriceNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.gaspriceCount += 1
	return a, b
}
func _opGasprice(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v, _ := uint256.FromBig(interpreter.evm.GasPrice)
	scope.Stack.push(v)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opBlockhash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opBlockhash(pc, interpreter, scope)
	bytecodeInfoLog.blockhashNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.blockhashCount += 1
	return a, b
}
func _opBlockhash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	num := scope.Stack.peek()
	num64, overflow := num.Uint64WithOverflow()
	if overflow {
		num.Clear()
		return nil, nil
	}
	var upper, lower uint64
	upper = interpreter.evm.Context.BlockNumber.Uint64()
	if upper < 257 {
		lower = 0
	} else {
		lower = upper - 256
	}
	if num64 >= lower && num64 < upper {
		num.SetBytes(interpreter.evm.Context.GetHash(num64).Bytes())
	} else {
		num.Clear()
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCoinbase(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCoinbase(pc, interpreter, scope)
	bytecodeInfoLog.coinbaseNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.coinbaseCount += 1
	return a, b
}
func _opCoinbase(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetBytes(interpreter.evm.Context.Coinbase.Bytes()))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opTimestamp(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opTimestamp(pc, interpreter, scope)
	bytecodeInfoLog.timestampNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.timestampCount += 1
	return a, b
}
func _opTimestamp(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v, _ := uint256.FromBig(interpreter.evm.Context.Time)
	scope.Stack.push(v)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opNumber(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opNumber(pc, interpreter, scope)
	bytecodeInfoLog.numberNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.numberCount += 1
	return a, b
}
func _opNumber(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v, _ := uint256.FromBig(interpreter.evm.Context.BlockNumber)
	scope.Stack.push(v)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opDifficulty(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opDifficulty(pc, interpreter, scope)
	bytecodeInfoLog.difficultyNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.difficultyCount += 1
	return a, b
}
func _opDifficulty(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v, _ := uint256.FromBig(interpreter.evm.Context.Difficulty)
	scope.Stack.push(v)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opRandom(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opRandom(pc, interpreter, scope)
	bytecodeInfoLog.randomNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.randomCount += 1
	return a, b
}
func _opRandom(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v := new(uint256.Int).SetBytes((interpreter.evm.Context.Random.Bytes()))
	scope.Stack.push(v)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opGasLimit(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opGasLimit(pc, interpreter, scope)
	bytecodeInfoLog.gasLimitNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.gasLimitCount += 1
	return a, b
}
func _opGasLimit(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetUint64(interpreter.evm.Context.GasLimit))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opPop(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opPop(pc, interpreter, scope)
	bytecodeInfoLog.popNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.popCount += 1
	return a, b
}
func _opPop(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.pop()
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMload(pc, interpreter, scope)
	bytecodeInfoLog.mloadNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.mloadCount += 1
	return a, b
}
func _opMload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	v := scope.Stack.peek()
	offset := int64(v.Uint64())
	v.SetBytes(scope.Memory.GetPtr(offset, 32))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMstore(pc, interpreter, scope)
	bytecodeInfoLog.mstoreNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.mstoreCount += 1
	return a, b
}
func _opMstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// pop value of the stack
	mStart, val := scope.Stack.pop(), scope.Stack.pop()
	scope.Memory.Set32(mStart.Uint64(), &val)
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMstore8(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMstore8(pc, interpreter, scope)
	bytecodeInfoLog.mstore8Nanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.mstore8Count += 1
	return a, b
}
func _opMstore8(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	off, val := scope.Stack.pop(), scope.Stack.pop()
	scope.Memory.store[off.Uint64()] = byte(val.Uint64())
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSload(pc, interpreter, scope)
	bytecodeInfoLog.sloadNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.sloadCount += 1
	return a, b
}
func _opSload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	loc := scope.Stack.peek()
	hash := common.Hash(loc.Bytes32())
	val := interpreter.evm.StateDB.GetState(scope.Contract.Address(), hash)
	loc.SetBytes(val.Bytes())
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opSstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSstore(pc, interpreter, scope)
	bytecodeInfoLog.sstoreNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.sstoreCount += 1
	return a, b
}
func _opSstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	loc := scope.Stack.pop()
	val := scope.Stack.pop()
	interpreter.evm.StateDB.SetState(scope.Contract.Address(),
		loc.Bytes32(), val.Bytes32())
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opJump(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opJump(pc, interpreter, scope)
	bytecodeInfoLog.jumpNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.jumpCount += 1
	return a, b
}
func _opJump(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if atomic.LoadInt32(&interpreter.evm.abort) != 0 {
		return nil, errStopToken
	}
	pos := scope.Stack.pop()
	if !scope.Contract.validJumpdest(&pos) {
		return nil, ErrInvalidJump
	}
	*pc = pos.Uint64() - 1 // pc will be increased by the interpreter loop
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opJumpi(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opJumpi(pc, interpreter, scope)
	bytecodeInfoLog.jumpiNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.jumpiCount += 1
	return a, b
}
func _opJumpi(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if atomic.LoadInt32(&interpreter.evm.abort) != 0 {
		return nil, errStopToken
	}
	pos, cond := scope.Stack.pop(), scope.Stack.pop()
	if !cond.IsZero() {
		if !scope.Contract.validJumpdest(&pos) {
			return nil, ErrInvalidJump
		}
		*pc = pos.Uint64() - 1 // pc will be increased by the interpreter loop
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opJumpdest(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opJumpdest(pc, interpreter, scope)
	bytecodeInfoLog.jumpdestNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.jumpdestCount += 1
	return a, b
}
func _opJumpdest(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opPc(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opPc(pc, interpreter, scope)
	bytecodeInfoLog.pcNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.pcCount += 1
	return a, b
}
func _opPc(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetUint64(*pc))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opMsize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opMsize(pc, interpreter, scope)
	bytecodeInfoLog.msizeNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.msizeCount += 1
	return a, b
}
func _opMsize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetUint64(uint64(scope.Memory.Len())))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opGas(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opGas(pc, interpreter, scope)
	bytecodeInfoLog.gasNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.gasCount += 1
	return a, b
}
func _opGas(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int).SetUint64(scope.Contract.Gas))
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCreate(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCreate(pc, interpreter, scope)
	bytecodeInfoLog.createNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.createCount += 1
	return a, b
}
func _opCreate(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	var (
		value        = scope.Stack.pop()
		offset, size = scope.Stack.pop(), scope.Stack.pop()
		input        = scope.Memory.GetCopy(int64(offset.Uint64()), int64(size.Uint64()))
		gas          = scope.Contract.Gas
	)
	if interpreter.evm.chainRules.IsEIP150 {
		gas -= gas / 64
	}
	// reuse size int for stackvalue
	stackvalue := size

	scope.Contract.UseGas(gas)
	//TODO: use uint256.Int instead of converting with toBig()
	var bigVal = big0
	if !value.IsZero() {
		bigVal = value.ToBig()
	}

	res, addr, returnGas, suberr := interpreter.evm.Create(scope.Contract, input, gas, bigVal)
	// Push item on the stack based on the returned error. If the ruleset is
	// homestead we must check for CodeStoreOutOfGasError (homestead only
	// rule) and treat as an error, if the ruleset is frontier we must
	// ignore this error and pretend the operation was successful.
	if interpreter.evm.chainRules.IsHomestead && suberr == ErrCodeStoreOutOfGas {
		stackvalue.Clear()
	} else if suberr != nil && suberr != ErrCodeStoreOutOfGas {
		stackvalue.Clear()
	} else {
		stackvalue.SetBytes(addr.Bytes())
	}
	scope.Stack.push(&stackvalue)
	scope.Contract.Gas += returnGas

	if suberr == ErrExecutionReverted {
		interpreter.returnData = res // set REVERT data to return data buffer
		return res, nil
	}
	interpreter.returnData = nil // clear dirty return data buffer
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCreate2(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCreate2(pc, interpreter, scope)
	bytecodeInfoLog.create2Nanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.create2Count += 1
	return a, b
}
func _opCreate2(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	var (
		endowment    = scope.Stack.pop()
		offset, size = scope.Stack.pop(), scope.Stack.pop()
		salt         = scope.Stack.pop()
		input        = scope.Memory.GetCopy(int64(offset.Uint64()), int64(size.Uint64()))
		gas          = scope.Contract.Gas
	)

	// Apply EIP150
	gas -= gas / 64
	scope.Contract.UseGas(gas)
	// reuse size int for stackvalue
	stackvalue := size
	//TODO: use uint256.Int instead of converting with toBig()
	bigEndowment := big0
	if !endowment.IsZero() {
		bigEndowment = endowment.ToBig()
	}
	res, addr, returnGas, suberr := interpreter.evm.Create2(scope.Contract, input, gas,
		bigEndowment, &salt)
	// Push item on the stack based on the returned error.
	if suberr != nil {
		stackvalue.Clear()
	} else {
		stackvalue.SetBytes(addr.Bytes())
	}
	scope.Stack.push(&stackvalue)
	scope.Contract.Gas += returnGas

	if suberr == ErrExecutionReverted {
		interpreter.returnData = res // set REVERT data to return data buffer
		return res, nil
	}
	interpreter.returnData = nil // clear dirty return data buffer
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func opCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCall(pc, interpreter, scope)
	bytecodeInfoLog.callNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callCount += 1
	return a, b
}
func _opCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	stack := scope.Stack
	// Pop gas. The actual gas in interpreter.evm.callGasTemp.
	// We can use this as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, value, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get the arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	if interpreter.readOnly && !value.IsZero() {
		return nil, ErrWriteProtection
	}
	var bigVal = big0
	//TODO: use uint256.Int instead of converting with toBig()
	// By using big0 here, we save an alloc for the most common case (non-ether-transferring contract calls),
	// but it would make more sense to extend the usage of uint256.Int
	if !value.IsZero() {
		gas += params.CallStipend
		bigVal = value.ToBig()
	}

	ret, returnGas, err := interpreter.evm.Call(scope.Contract, toAddr, args, gas, bigVal)

	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		ret = common.CopyBytes(ret)
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	scope.Contract.Gas += returnGas

	interpreter.returnData = ret
	return ret, nil
}

// Cybersecurity Lab: Measure duration
func opCallCode(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opCallCode(pc, interpreter, scope)
	bytecodeInfoLog.callCodeNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.callCodeCount += 1
	return a, b
}
func _opCallCode(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	stack := scope.Stack
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, value, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	//TODO: use uint256.Int instead of converting with toBig()
	var bigVal = big0
	if !value.IsZero() {
		gas += params.CallStipend
		bigVal = value.ToBig()
	}

	ret, returnGas, err := interpreter.evm.CallCode(scope.Contract, toAddr, args, gas, bigVal)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		ret = common.CopyBytes(ret)
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	scope.Contract.Gas += returnGas

	interpreter.returnData = ret
	return ret, nil
}

// Cybersecurity Lab: Measure duration
func opDelegateCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opDelegateCall(pc, interpreter, scope)
	bytecodeInfoLog.delegateCallNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.delegateCallCount += 1
	return a, b
}
func _opDelegateCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	stack := scope.Stack
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	ret, returnGas, err := interpreter.evm.DelegateCall(scope.Contract, toAddr, args, gas)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		ret = common.CopyBytes(ret)
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	scope.Contract.Gas += returnGas

	interpreter.returnData = ret
	return ret, nil
}

// Cybersecurity Lab: Measure duration
func opStaticCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opStaticCall(pc, interpreter, scope)
	bytecodeInfoLog.staticCallNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.staticCallCount += 1
	return a, b
}
func _opStaticCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	stack := scope.Stack
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	ret, returnGas, err := interpreter.evm.StaticCall(scope.Contract, toAddr, args, gas)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		ret = common.CopyBytes(ret)
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	scope.Contract.Gas += returnGas

	interpreter.returnData = ret
	return ret, nil
}

// Cybersecurity Lab: Measure duration
func opReturn(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opReturn(pc, interpreter, scope)
	bytecodeInfoLog.returnNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.returnCount += 1
	return a, b
}
func _opReturn(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	offset, size := scope.Stack.pop(), scope.Stack.pop()
	ret := scope.Memory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	return ret, errStopToken
}

// Cybersecurity Lab: Measure duration
func opRevert(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opRevert(pc, interpreter, scope)
	bytecodeInfoLog.revertNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.revertCount += 1
	return a, b
}
func _opRevert(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	offset, size := scope.Stack.pop(), scope.Stack.pop()
	ret := scope.Memory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	interpreter.returnData = ret
	return ret, ErrExecutionReverted
}

// Cybersecurity Lab: Measure duration
func opUndefined(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opUndefined(pc, interpreter, scope)
	bytecodeInfoLog.undefinedNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.undefinedCount += 1
	return a, b
}
func _opUndefined(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	return nil, &ErrInvalidOpCode{opcode: OpCode(scope.Contract.Code[*pc])}
}

// Cybersecurity Lab: Measure duration
func opStop(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opStop(pc, interpreter, scope)
	bytecodeInfoLog.stopNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.stopCount += 1
	return a, b
}
func _opStop(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	return nil, errStopToken
}

// Cybersecurity Lab: Measure duration
func opSelfdestruct(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opSelfdestruct(pc, interpreter, scope)
	bytecodeInfoLog.selfdestructNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.selfdestructCount += 1
	return a, b
}
func _opSelfdestruct(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	beneficiary := scope.Stack.pop()
	balance := interpreter.evm.StateDB.GetBalance(scope.Contract.Address())
	interpreter.evm.StateDB.AddBalance(beneficiary.Bytes20(), balance)
	interpreter.evm.StateDB.Suicide(scope.Contract.Address())
	if interpreter.cfg.Debug {
		interpreter.cfg.Tracer.CaptureEnter(SELFDESTRUCT, scope.Contract.Address(), beneficiary.Bytes20(), []byte{}, 0, balance)
		interpreter.cfg.Tracer.CaptureExit([]byte{}, 0, nil)
	}
	return nil, errStopToken
}

// following functions are used by the instruction jump  table

// Cybersecurity Lab: Measure duration
func makeLog(size int) executionFunc {
	var a executionFunc
	start := time.Now()
	a = _makeLog(size)
	bytecodeInfoLog.makeLogNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.makeLogCount += 1
	return a
}

// make log instruction function
func _makeLog(size int) executionFunc {
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		if interpreter.readOnly {
			return nil, ErrWriteProtection
		}
		topics := make([]common.Hash, size)
		stack := scope.Stack
		mStart, mSize := stack.pop(), stack.pop()
		for i := 0; i < size; i++ {
			addr := stack.pop()
			topics[i] = addr.Bytes32()
		}

		d := scope.Memory.GetCopy(int64(mStart.Uint64()), int64(mSize.Uint64()))
		interpreter.evm.StateDB.AddLog(&types.Log{
			Address: scope.Contract.Address(),
			Topics:  topics,
			Data:    d,
			// This is a non-consensus field, but assigned here because
			// core/state doesn't know the current block number.
			BlockNumber: interpreter.evm.Context.BlockNumber.Uint64(),
		})

		return nil, nil
	}
}

// Cybersecurity Lab: Measure duration
func opPush1(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opPush1(pc, interpreter, scope)
	bytecodeInfoLog.push1Nanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.push1Count += 1
	return a, b
}

// opPush1 is a specialized version of pushN
func _opPush1(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var (
		codeLen = uint64(len(scope.Contract.Code))
		integer = new(uint256.Int)
	)
	*pc += 1
	if *pc < codeLen {
		scope.Stack.push(integer.SetUint64(uint64(scope.Contract.Code[*pc])))
	} else {
		scope.Stack.push(integer.Clear())
	}
	return nil, nil
}

// Cybersecurity Lab: Measure duration
func makePush(size uint64, pushByteSize int) executionFunc {
	var a executionFunc
	start := time.Now()
	a = _makePush(size, pushByteSize)
	bytecodeInfoLog.makePushNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.makePushCount += 1
	return a
}

// make push instruction function
func _makePush(size uint64, pushByteSize int) executionFunc {
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		codeLen := len(scope.Contract.Code)

		startMin := codeLen
		if int(*pc+1) < startMin {
			startMin = int(*pc + 1)
		}

		endMin := codeLen
		if startMin+pushByteSize < endMin {
			endMin = startMin + pushByteSize
		}

		integer := new(uint256.Int)
		scope.Stack.push(integer.SetBytes(common.RightPadBytes(
			scope.Contract.Code[startMin:endMin], pushByteSize)))

		*pc += size
		return nil, nil
	}
}

// Cybersecurity Lab: Measure duration
func makeDup(size int64) executionFunc {
	var a executionFunc
	start := time.Now()
	a = _makeDup(size)
	bytecodeInfoLog.makeDupNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.makeDupCount += 1
	return a
}

// make dup instruction function
func _makeDup(size int64) executionFunc {
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		scope.Stack.dup(int(size))
		return nil, nil
	}
}

// Cybersecurity Lab: Measure duration
func makeSwap(size int64) executionFunc {
	var a executionFunc
	start := time.Now()
	a = _makeSwap(size)
	bytecodeInfoLog.makeSwapNanoseconds += int64(time.Since(start) / time.Nanosecond)
	bytecodeInfoLog.makeSwapCount += 1
	return a
}

// make swap instruction function
func _makeSwap(size int64) executionFunc {
	// switch n + 1 otherwise n would be swapped with n
	size++
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		scope.Stack.swap(int(size))
		return nil, nil
	}
}
