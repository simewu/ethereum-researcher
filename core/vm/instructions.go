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
	addNanoseconds               int64 `json:"number"`
	addMinNanoseconds            int64 `json:"number"`
	addMaxNanoseconds            int64 `json:"number"`
	addGas                       int   `json:"number"`
	addMinGas                    int   `json:"number"`
	addMaxGas                    int   `json:"number"`
	addCount                     int   `json:"number"`
	subNanoseconds               int64 `json:"number"`
	subMinNanoseconds            int64 `json:"number"`
	subMaxNanoseconds            int64 `json:"number"`
	subGas                       int   `json:"number"`
	subMinGas                    int   `json:"number"`
	subMaxGas                    int   `json:"number"`
	subCount                     int   `json:"number"`
	mulNanoseconds               int64 `json:"number"`
	mulMinNanoseconds            int64 `json:"number"`
	mulMaxNanoseconds            int64 `json:"number"`
	mulGas                       int   `json:"number"`
	mulMinGas                    int   `json:"number"`
	mulMaxGas                    int   `json:"number"`
	mulCount                     int   `json:"number"`
	divNanoseconds               int64 `json:"number"`
	divMinNanoseconds            int64 `json:"number"`
	divMaxNanoseconds            int64 `json:"number"`
	divGas                       int   `json:"number"`
	divMinGas                    int   `json:"number"`
	divMaxGas                    int   `json:"number"`
	divCount                     int   `json:"number"`
	sdivNanoseconds              int64 `json:"number"`
	sdivMinNanoseconds           int64 `json:"number"`
	sdivMaxNanoseconds           int64 `json:"number"`
	sdivGas                      int   `json:"number"`
	sdivMinGas                   int   `json:"number"`
	sdivMaxGas                   int   `json:"number"`
	sdivCount                    int   `json:"number"`
	modNanoseconds               int64 `json:"number"`
	modMinNanoseconds            int64 `json:"number"`
	modMaxNanoseconds            int64 `json:"number"`
	modGas                       int   `json:"number"`
	modMinGas                    int   `json:"number"`
	modMaxGas                    int   `json:"number"`
	modCount                     int   `json:"number"`
	smodNanoseconds              int64 `json:"number"`
	smodMinNanoseconds           int64 `json:"number"`
	smodMaxNanoseconds           int64 `json:"number"`
	smodGas                      int   `json:"number"`
	smodMinGas                   int   `json:"number"`
	smodMaxGas                   int   `json:"number"`
	smodCount                    int   `json:"number"`
	expNanoseconds               int64 `json:"number"`
	expMinNanoseconds            int64 `json:"number"`
	expMaxNanoseconds            int64 `json:"number"`
	expGas                       int   `json:"number"`
	expMinGas                    int   `json:"number"`
	expMaxGas                    int   `json:"number"`
	expCount                     int   `json:"number"`
	signExtendNanoseconds        int64 `json:"number"`
	signExtendMinNanoseconds     int64 `json:"number"`
	signExtendMaxNanoseconds     int64 `json:"number"`
	signExtendGas                int   `json:"number"`
	signExtendMinGas             int   `json:"number"`
	signExtendMaxGas             int   `json:"number"`
	signExtendCount              int   `json:"number"`
	notNanoseconds               int64 `json:"number"`
	notMinNanoseconds            int64 `json:"number"`
	notMaxNanoseconds            int64 `json:"number"`
	notGas                       int   `json:"number"`
	notMinGas                    int   `json:"number"`
	notMaxGas                    int   `json:"number"`
	notCount                     int   `json:"number"`
	ltNanoseconds                int64 `json:"number"`
	ltMinNanoseconds             int64 `json:"number"`
	ltMaxNanoseconds             int64 `json:"number"`
	ltGas                        int   `json:"number"`
	ltMinGas                     int   `json:"number"`
	ltMaxGas                     int   `json:"number"`
	ltCount                      int   `json:"number"`
	gtNanoseconds                int64 `json:"number"`
	gtMinNanoseconds             int64 `json:"number"`
	gtMaxNanoseconds             int64 `json:"number"`
	gtGas                        int   `json:"number"`
	gtMinGas                     int   `json:"number"`
	gtMaxGas                     int   `json:"number"`
	gtCount                      int   `json:"number"`
	sltNanoseconds               int64 `json:"number"`
	sltMinNanoseconds            int64 `json:"number"`
	sltMaxNanoseconds            int64 `json:"number"`
	sltGas                       int   `json:"number"`
	sltMinGas                    int   `json:"number"`
	sltMaxGas                    int   `json:"number"`
	sltCount                     int   `json:"number"`
	sgtNanoseconds               int64 `json:"number"`
	sgtMinNanoseconds            int64 `json:"number"`
	sgtMaxNanoseconds            int64 `json:"number"`
	sgtGas                       int   `json:"number"`
	sgtMinGas                    int   `json:"number"`
	sgtMaxGas                    int   `json:"number"`
	sgtCount                     int   `json:"number"`
	eqNanoseconds                int64 `json:"number"`
	eqMinNanoseconds             int64 `json:"number"`
	eqMaxNanoseconds             int64 `json:"number"`
	eqGas                        int   `json:"number"`
	eqMinGas                     int   `json:"number"`
	eqMaxGas                     int   `json:"number"`
	eqCount                      int   `json:"number"`
	isZeroNanoseconds            int64 `json:"number"`
	isZeroMinNanoseconds         int64 `json:"number"`
	isZeroMaxNanoseconds         int64 `json:"number"`
	isZeroGas                    int   `json:"number"`
	isZeroMinGas                 int   `json:"number"`
	isZeroMaxGas                 int   `json:"number"`
	isZeroCount                  int   `json:"number"`
	andNanoseconds               int64 `json:"number"`
	andMinNanoseconds            int64 `json:"number"`
	andMaxNanoseconds            int64 `json:"number"`
	andGas                       int   `json:"number"`
	andMinGas                    int   `json:"number"`
	andMaxGas                    int   `json:"number"`
	andCount                     int   `json:"number"`
	orNanoseconds                int64 `json:"number"`
	orMinNanoseconds             int64 `json:"number"`
	orMaxNanoseconds             int64 `json:"number"`
	orGas                        int   `json:"number"`
	orMinGas                     int   `json:"number"`
	orMaxGas                     int   `json:"number"`
	orCount                      int   `json:"number"`
	xorNanoseconds               int64 `json:"number"`
	xorMinNanoseconds            int64 `json:"number"`
	xorMaxNanoseconds            int64 `json:"number"`
	xorGas                       int   `json:"number"`
	xorMinGas                    int   `json:"number"`
	xorMaxGas                    int   `json:"number"`
	xorCount                     int   `json:"number"`
	byteNanoseconds              int64 `json:"number"`
	byteMinNanoseconds           int64 `json:"number"`
	byteMaxNanoseconds           int64 `json:"number"`
	byteGas                      int   `json:"number"`
	byteMinGas                   int   `json:"number"`
	byteMaxGas                   int   `json:"number"`
	byteCount                    int   `json:"number"`
	addmodNanoseconds            int64 `json:"number"`
	addmodMinNanoseconds         int64 `json:"number"`
	addmodMaxNanoseconds         int64 `json:"number"`
	addmodGas                    int   `json:"number"`
	addmodMinGas                 int   `json:"number"`
	addmodMaxGas                 int   `json:"number"`
	addmodCount                  int   `json:"number"`
	mulmodNanoseconds            int64 `json:"number"`
	mulmodMinNanoseconds         int64 `json:"number"`
	mulmodMaxNanoseconds         int64 `json:"number"`
	mulmodGas                    int   `json:"number"`
	mulmodMinGas                 int   `json:"number"`
	mulmodMaxGas                 int   `json:"number"`
	mulmodCount                  int   `json:"number"`
	shlNanoseconds               int64 `json:"number"`
	shlMinNanoseconds            int64 `json:"number"`
	shlMaxNanoseconds            int64 `json:"number"`
	shlGas                       int   `json:"number"`
	shlMinGas                    int   `json:"number"`
	shlMaxGas                    int   `json:"number"`
	shlCount                     int   `json:"number"`
	shrNanoseconds               int64 `json:"number"`
	shrMinNanoseconds            int64 `json:"number"`
	shrMaxNanoseconds            int64 `json:"number"`
	shrGas                       int   `json:"number"`
	shrMinGas                    int   `json:"number"`
	shrMaxGas                    int   `json:"number"`
	shrCount                     int   `json:"number"`
	sarNanoseconds               int64 `json:"number"`
	sarMinNanoseconds            int64 `json:"number"`
	sarMaxNanoseconds            int64 `json:"number"`
	sarGas                       int   `json:"number"`
	sarMinGas                    int   `json:"number"`
	sarMaxGas                    int   `json:"number"`
	sarCount                     int   `json:"number"`
	keccak256Nanoseconds         int64 `json:"number"`
	keccak256MinNanoseconds      int64 `json:"number"`
	keccak256MaxNanoseconds      int64 `json:"number"`
	keccak256Gas                 int   `json:"number"`
	keccak256MinGas              int   `json:"number"`
	keccak256MaxGas              int   `json:"number"`
	keccak256Count               int   `json:"number"`
	addressNanoseconds           int64 `json:"number"`
	addressMinNanoseconds        int64 `json:"number"`
	addressMaxNanoseconds        int64 `json:"number"`
	addressGas                   int   `json:"number"`
	addressMinGas                int   `json:"number"`
	addressMaxGas                int   `json:"number"`
	addressCount                 int   `json:"number"`
	balanceNanoseconds           int64 `json:"number"`
	balanceMinNanoseconds        int64 `json:"number"`
	balanceMaxNanoseconds        int64 `json:"number"`
	balanceGas                   int   `json:"number"`
	balanceMinGas                int   `json:"number"`
	balanceMaxGas                int   `json:"number"`
	balanceCount                 int   `json:"number"`
	originNanoseconds            int64 `json:"number"`
	originMinNanoseconds         int64 `json:"number"`
	originMaxNanoseconds         int64 `json:"number"`
	originGas                    int   `json:"number"`
	originMinGas                 int   `json:"number"`
	originMaxGas                 int   `json:"number"`
	originCount                  int   `json:"number"`
	callerNanoseconds            int64 `json:"number"`
	callerMinNanoseconds         int64 `json:"number"`
	callerMaxNanoseconds         int64 `json:"number"`
	callerGas                    int   `json:"number"`
	callerMinGas                 int   `json:"number"`
	callerMaxGas                 int   `json:"number"`
	callerCount                  int   `json:"number"`
	callValueNanoseconds         int64 `json:"number"`
	callValueMinNanoseconds      int64 `json:"number"`
	callValueMaxNanoseconds      int64 `json:"number"`
	callValueGas                 int   `json:"number"`
	callValueMinGas              int   `json:"number"`
	callValueMaxGas              int   `json:"number"`
	callValueCount               int   `json:"number"`
	callDataLoadNanoseconds      int64 `json:"number"`
	callDataLoadMinNanoseconds   int64 `json:"number"`
	callDataLoadMaxNanoseconds   int64 `json:"number"`
	callDataLoadGas              int   `json:"number"`
	callDataLoadMinGas           int   `json:"number"`
	callDataLoadMaxGas           int   `json:"number"`
	callDataLoadCount            int   `json:"number"`
	callDataSizeNanoseconds      int64 `json:"number"`
	callDataSizeMinNanoseconds   int64 `json:"number"`
	callDataSizeMaxNanoseconds   int64 `json:"number"`
	callDataSizeGas              int   `json:"number"`
	callDataSizeMinGas           int   `json:"number"`
	callDataSizeMaxGas           int   `json:"number"`
	callDataSizeCount            int   `json:"number"`
	callDataCopyNanoseconds      int64 `json:"number"`
	callDataCopyMinNanoseconds   int64 `json:"number"`
	callDataCopyMaxNanoseconds   int64 `json:"number"`
	callDataCopyGas              int   `json:"number"`
	callDataCopyMinGas           int   `json:"number"`
	callDataCopyMaxGas           int   `json:"number"`
	callDataCopyCount            int   `json:"number"`
	returnDataSizeNanoseconds    int64 `json:"number"`
	returnDataSizeMinNanoseconds int64 `json:"number"`
	returnDataSizeMaxNanoseconds int64 `json:"number"`
	returnDataSizeGas            int   `json:"number"`
	returnDataSizeMinGas         int   `json:"number"`
	returnDataSizeMaxGas         int   `json:"number"`
	returnDataSizeCount          int   `json:"number"`
	returnDataCopyNanoseconds    int64 `json:"number"`
	returnDataCopyMinNanoseconds int64 `json:"number"`
	returnDataCopyMaxNanoseconds int64 `json:"number"`
	returnDataCopyGas            int   `json:"number"`
	returnDataCopyMinGas         int   `json:"number"`
	returnDataCopyMaxGas         int   `json:"number"`
	returnDataCopyCount          int   `json:"number"`
	extCodeSizeNanoseconds       int64 `json:"number"`
	extCodeSizeMinNanoseconds    int64 `json:"number"`
	extCodeSizeMaxNanoseconds    int64 `json:"number"`
	extCodeSizeGas               int   `json:"number"`
	extCodeSizeMinGas            int   `json:"number"`
	extCodeSizeMaxGas            int   `json:"number"`
	extCodeSizeCount             int   `json:"number"`
	codeSizeNanoseconds          int64 `json:"number"`
	codeSizeMinNanoseconds       int64 `json:"number"`
	codeSizeMaxNanoseconds       int64 `json:"number"`
	codeSizeGas                  int   `json:"number"`
	codeSizeMinGas               int   `json:"number"`
	codeSizeMaxGas               int   `json:"number"`
	codeSizeCount                int   `json:"number"`
	codeCopyNanoseconds          int64 `json:"number"`
	codeCopyMinNanoseconds       int64 `json:"number"`
	codeCopyMaxNanoseconds       int64 `json:"number"`
	codeCopyGas                  int   `json:"number"`
	codeCopyMinGas               int   `json:"number"`
	codeCopyMaxGas               int   `json:"number"`
	codeCopyCount                int   `json:"number"`
	extCodeCopyNanoseconds       int64 `json:"number"`
	extCodeCopyMinNanoseconds    int64 `json:"number"`
	extCodeCopyMaxNanoseconds    int64 `json:"number"`
	extCodeCopyGas               int   `json:"number"`
	extCodeCopyMinGas            int   `json:"number"`
	extCodeCopyMaxGas            int   `json:"number"`
	extCodeCopyCount             int   `json:"number"`
	extCodeHashNanoseconds       int64 `json:"number"`
	extCodeHashMinNanoseconds    int64 `json:"number"`
	extCodeHashMaxNanoseconds    int64 `json:"number"`
	extCodeHashGas               int   `json:"number"`
	extCodeHashMinGas            int   `json:"number"`
	extCodeHashMaxGas            int   `json:"number"`
	extCodeHashCount             int   `json:"number"`
	gaspriceNanoseconds          int64 `json:"number"`
	gaspriceMinNanoseconds       int64 `json:"number"`
	gaspriceMaxNanoseconds       int64 `json:"number"`
	gaspriceGas                  int   `json:"number"`
	gaspriceMinGas               int   `json:"number"`
	gaspriceMaxGas               int   `json:"number"`
	gaspriceCount                int   `json:"number"`
	blockhashNanoseconds         int64 `json:"number"`
	blockhashMinNanoseconds      int64 `json:"number"`
	blockhashMaxNanoseconds      int64 `json:"number"`
	blockhashGas                 int   `json:"number"`
	blockhashMinGas              int   `json:"number"`
	blockhashMaxGas              int   `json:"number"`
	blockhashCount               int   `json:"number"`
	coinbaseNanoseconds          int64 `json:"number"`
	coinbaseMinNanoseconds       int64 `json:"number"`
	coinbaseMaxNanoseconds       int64 `json:"number"`
	coinbaseGas                  int   `json:"number"`
	coinbaseMinGas               int   `json:"number"`
	coinbaseMaxGas               int   `json:"number"`
	coinbaseCount                int   `json:"number"`
	timestampNanoseconds         int64 `json:"number"`
	timestampMinNanoseconds      int64 `json:"number"`
	timestampMaxNanoseconds      int64 `json:"number"`
	timestampGas                 int   `json:"number"`
	timestampMinGas              int   `json:"number"`
	timestampMaxGas              int   `json:"number"`
	timestampCount               int   `json:"number"`
	numberNanoseconds            int64 `json:"number"`
	numberMinNanoseconds         int64 `json:"number"`
	numberMaxNanoseconds         int64 `json:"number"`
	numberGas                    int   `json:"number"`
	numberMinGas                 int   `json:"number"`
	numberMaxGas                 int   `json:"number"`
	numberCount                  int   `json:"number"`
	difficultyNanoseconds        int64 `json:"number"`
	difficultyMinNanoseconds     int64 `json:"number"`
	difficultyMaxNanoseconds     int64 `json:"number"`
	difficultyGas                int   `json:"number"`
	difficultyMinGas             int   `json:"number"`
	difficultyMaxGas             int   `json:"number"`
	difficultyCount              int   `json:"number"`
	randomNanoseconds            int64 `json:"number"`
	randomMinNanoseconds         int64 `json:"number"`
	randomMaxNanoseconds         int64 `json:"number"`
	randomGas                    int   `json:"number"`
	randomMinGas                 int   `json:"number"`
	randomMaxGas                 int   `json:"number"`
	randomCount                  int   `json:"number"`
	gasLimitNanoseconds          int64 `json:"number"`
	gasLimitMinNanoseconds       int64 `json:"number"`
	gasLimitMaxNanoseconds       int64 `json:"number"`
	gasLimitGas                  int   `json:"number"`
	gasLimitMinGas               int   `json:"number"`
	gasLimitMaxGas               int   `json:"number"`
	gasLimitCount                int   `json:"number"`
	popNanoseconds               int64 `json:"number"`
	popMinNanoseconds            int64 `json:"number"`
	popMaxNanoseconds            int64 `json:"number"`
	popGas                       int   `json:"number"`
	popMinGas                    int   `json:"number"`
	popMaxGas                    int   `json:"number"`
	popCount                     int   `json:"number"`
	mloadNanoseconds             int64 `json:"number"`
	mloadMinNanoseconds          int64 `json:"number"`
	mloadMaxNanoseconds          int64 `json:"number"`
	mloadGas                     int   `json:"number"`
	mloadMinGas                  int   `json:"number"`
	mloadMaxGas                  int   `json:"number"`
	mloadCount                   int   `json:"number"`
	mstoreNanoseconds            int64 `json:"number"`
	mstoreMinNanoseconds         int64 `json:"number"`
	mstoreMaxNanoseconds         int64 `json:"number"`
	mstoreGas                    int   `json:"number"`
	mstoreMinGas                 int   `json:"number"`
	mstoreMaxGas                 int   `json:"number"`
	mstoreCount                  int   `json:"number"`
	mstore8Nanoseconds           int64 `json:"number"`
	mstore8MinNanoseconds        int64 `json:"number"`
	mstore8MaxNanoseconds        int64 `json:"number"`
	mstore8Gas                   int   `json:"number"`
	mstore8MinGas                int   `json:"number"`
	mstore8MaxGas                int   `json:"number"`
	mstore8Count                 int   `json:"number"`
	sloadNanoseconds             int64 `json:"number"`
	sloadMinNanoseconds          int64 `json:"number"`
	sloadMaxNanoseconds          int64 `json:"number"`
	sloadGas                     int   `json:"number"`
	sloadMinGas                  int   `json:"number"`
	sloadMaxGas                  int   `json:"number"`
	sloadCount                   int   `json:"number"`
	sstoreNanoseconds            int64 `json:"number"`
	sstoreMinNanoseconds         int64 `json:"number"`
	sstoreMaxNanoseconds         int64 `json:"number"`
	sstoreGas                    int   `json:"number"`
	sstoreMinGas                 int   `json:"number"`
	sstoreMaxGas                 int   `json:"number"`
	sstoreCount                  int   `json:"number"`
	jumpNanoseconds              int64 `json:"number"`
	jumpMinNanoseconds           int64 `json:"number"`
	jumpMaxNanoseconds           int64 `json:"number"`
	jumpGas                      int   `json:"number"`
	jumpMinGas                   int   `json:"number"`
	jumpMaxGas                   int   `json:"number"`
	jumpCount                    int   `json:"number"`
	jumpiNanoseconds             int64 `json:"number"`
	jumpiMinNanoseconds          int64 `json:"number"`
	jumpiMaxNanoseconds          int64 `json:"number"`
	jumpiGas                     int   `json:"number"`
	jumpiMinGas                  int   `json:"number"`
	jumpiMaxGas                  int   `json:"number"`
	jumpiCount                   int   `json:"number"`
	jumpdestNanoseconds          int64 `json:"number"`
	jumpdestMinNanoseconds       int64 `json:"number"`
	jumpdestMaxNanoseconds       int64 `json:"number"`
	jumpdestGas                  int   `json:"number"`
	jumpdestMinGas               int   `json:"number"`
	jumpdestMaxGas               int   `json:"number"`
	jumpdestCount                int   `json:"number"`
	pcNanoseconds                int64 `json:"number"`
	pcMinNanoseconds             int64 `json:"number"`
	pcMaxNanoseconds             int64 `json:"number"`
	pcGas                        int   `json:"number"`
	pcMinGas                     int   `json:"number"`
	pcMaxGas                     int   `json:"number"`
	pcCount                      int   `json:"number"`
	msizeNanoseconds             int64 `json:"number"`
	msizeMinNanoseconds          int64 `json:"number"`
	msizeMaxNanoseconds          int64 `json:"number"`
	msizeGas                     int   `json:"number"`
	msizeMinGas                  int   `json:"number"`
	msizeMaxGas                  int   `json:"number"`
	msizeCount                   int   `json:"number"`
	gasNanoseconds               int64 `json:"number"`
	gasMinNanoseconds            int64 `json:"number"`
	gasMaxNanoseconds            int64 `json:"number"`
	gasGas                       int   `json:"number"`
	gasMinGas                    int   `json:"number"`
	gasMaxGas                    int   `json:"number"`
	gasCount                     int   `json:"number"`
	createNanoseconds            int64 `json:"number"`
	createMinNanoseconds         int64 `json:"number"`
	createMaxNanoseconds         int64 `json:"number"`
	createGas                    int   `json:"number"`
	createMinGas                 int   `json:"number"`
	createMaxGas                 int   `json:"number"`
	createCount                  int   `json:"number"`
	create2Nanoseconds           int64 `json:"number"`
	create2MinNanoseconds        int64 `json:"number"`
	create2MaxNanoseconds        int64 `json:"number"`
	create2Gas                   int   `json:"number"`
	create2MinGas                int   `json:"number"`
	create2MaxGas                int   `json:"number"`
	create2Count                 int   `json:"number"`
	callNanoseconds              int64 `json:"number"`
	callMinNanoseconds           int64 `json:"number"`
	callMaxNanoseconds           int64 `json:"number"`
	callGas                      int   `json:"number"`
	callMinGas                   int   `json:"number"`
	callMaxGas                   int   `json:"number"`
	callCount                    int   `json:"number"`
	callCodeNanoseconds          int64 `json:"number"`
	callCodeMinNanoseconds       int64 `json:"number"`
	callCodeMaxNanoseconds       int64 `json:"number"`
	callCodeGas                  int   `json:"number"`
	callCodeMinGas               int   `json:"number"`
	callCodeMaxGas               int   `json:"number"`
	callCodeCount                int   `json:"number"`
	delegateCallNanoseconds      int64 `json:"number"`
	delegateCallMinNanoseconds   int64 `json:"number"`
	delegateCallMaxNanoseconds   int64 `json:"number"`
	delegateCallGas              int   `json:"number"`
	delegateCallMinGas           int   `json:"number"`
	delegateCallMaxGas           int   `json:"number"`
	delegateCallCount            int   `json:"number"`
	staticCallNanoseconds        int64 `json:"number"`
	staticCallMinNanoseconds     int64 `json:"number"`
	staticCallMaxNanoseconds     int64 `json:"number"`
	staticCallGas                int   `json:"number"`
	staticCallMinGas             int   `json:"number"`
	staticCallMaxGas             int   `json:"number"`
	staticCallCount              int   `json:"number"`
	returnNanoseconds            int64 `json:"number"`
	returnMinNanoseconds         int64 `json:"number"`
	returnMaxNanoseconds         int64 `json:"number"`
	returnGas                    int   `json:"number"`
	returnMinGas                 int   `json:"number"`
	returnMaxGas                 int   `json:"number"`
	returnCount                  int   `json:"number"`
	revertNanoseconds            int64 `json:"number"`
	revertMinNanoseconds         int64 `json:"number"`
	revertMaxNanoseconds         int64 `json:"number"`
	revertGas                    int   `json:"number"`
	revertMinGas                 int   `json:"number"`
	revertMaxGas                 int   `json:"number"`
	revertCount                  int   `json:"number"`
	undefinedNanoseconds         int64 `json:"number"`
	undefinedMinNanoseconds      int64 `json:"number"`
	undefinedMaxNanoseconds      int64 `json:"number"`
	undefinedGas                 int   `json:"number"`
	undefinedMinGas              int   `json:"number"`
	undefinedMaxGas              int   `json:"number"`
	undefinedCount               int   `json:"number"`
	stopNanoseconds              int64 `json:"number"`
	stopMinNanoseconds           int64 `json:"number"`
	stopMaxNanoseconds           int64 `json:"number"`
	stopGas                      int   `json:"number"`
	stopMinGas                   int   `json:"number"`
	stopMaxGas                   int   `json:"number"`
	stopCount                    int   `json:"number"`
	selfdestructNanoseconds      int64 `json:"number"`
	selfdestructMinNanoseconds   int64 `json:"number"`
	selfdestructMaxNanoseconds   int64 `json:"number"`
	selfdestructGas              int   `json:"number"`
	selfdestructMinGas           int   `json:"number"`
	selfdestructMaxGas           int   `json:"number"`
	selfdestructCount            int   `json:"number"`
	makeLogNanoseconds           int64 `json:"number"`
	makeLogMinNanoseconds        int64 `json:"number"`
	makeLogMaxNanoseconds        int64 `json:"number"`
	makeLogGas                   int   `json:"number"`
	makeLogMinGas                int   `json:"number"`
	makeLogMaxGas                int   `json:"number"`
	makeLogCount                 int   `json:"number"`
	push1Nanoseconds             int64 `json:"number"`
	push1MinNanoseconds          int64 `json:"number"`
	push1MaxNanoseconds          int64 `json:"number"`
	push1Gas                     int   `json:"number"`
	push1MinGas                  int   `json:"number"`
	push1MaxGas                  int   `json:"number"`
	push1Count                   int   `json:"number"`
	makePushNanoseconds          int64 `json:"number"`
	makePushMinNanoseconds       int64 `json:"number"`
	makePushMaxNanoseconds       int64 `json:"number"`
	makePushGas                  int   `json:"number"`
	makePushMinGas               int   `json:"number"`
	makePushMaxGas               int   `json:"number"`
	makePushCount                int   `json:"number"`
	makeDupNanoseconds           int64 `json:"number"`
	makeDupMinNanoseconds        int64 `json:"number"`
	makeDupMaxNanoseconds        int64 `json:"number"`
	makeDupGas                   int   `json:"number"`
	makeDupMinGas                int   `json:"number"`
	makeDupMaxGas                int   `json:"number"`
	makeDupCount                 int   `json:"number"`
	makeSwapNanoseconds          int64 `json:"number"`
	makeSwapMinNanoseconds       int64 `json:"number"`
	makeSwapMaxNanoseconds       int64 `json:"number"`
	makeSwapGas                  int   `json:"number"`
	makeSwapMinGas               int   `json:"number"`
	makeSwapMaxGas               int   `json:"number"`
	makeSwapCount                int   `json:"number"`
}

var bytecodeInfoLog bytecodeInfoType

// Cybersecurity Lab
type Export struct{}

func (export Export) GetBytecodeInfoStats() string {
	var output string = "{"
	output += `"add ns":"` + strconv.FormatInt(bytecodeInfoLog.addNanoseconds, 10) + `",`
	output += `"add min ns":"` + strconv.FormatInt(bytecodeInfoLog.addMinNanoseconds, 10) + `",`
	output += `"add max ns":"` + strconv.FormatInt(bytecodeInfoLog.addMaxNanoseconds, 10) + `",`
	output += `"add gas":"` + strconv.Itoa(bytecodeInfoLog.addGas) + `",`
	output += `"add min gas":"` + strconv.Itoa(bytecodeInfoLog.addMinGas) + `",`
	output += `"add max gas":"` + strconv.Itoa(bytecodeInfoLog.addMaxGas) + `",`
	output += `"addCount":"` + strconv.Itoa(bytecodeInfoLog.addCount) + `",`
	output += `"sub ns":"` + strconv.FormatInt(bytecodeInfoLog.subNanoseconds, 10) + `",`
	output += `"sub min ns":"` + strconv.FormatInt(bytecodeInfoLog.subMinNanoseconds, 10) + `",`
	output += `"sub max ns":"` + strconv.FormatInt(bytecodeInfoLog.subMaxNanoseconds, 10) + `",`
	output += `"sub gas":"` + strconv.Itoa(bytecodeInfoLog.subGas) + `",`
	output += `"sub min gas":"` + strconv.Itoa(bytecodeInfoLog.subMinGas) + `",`
	output += `"sub max gas":"` + strconv.Itoa(bytecodeInfoLog.subMaxGas) + `",`
	output += `"subCount":"` + strconv.Itoa(bytecodeInfoLog.subCount) + `",`
	output += `"mul ns":"` + strconv.FormatInt(bytecodeInfoLog.mulNanoseconds, 10) + `",`
	output += `"mul min ns":"` + strconv.FormatInt(bytecodeInfoLog.mulMinNanoseconds, 10) + `",`
	output += `"mul max ns":"` + strconv.FormatInt(bytecodeInfoLog.mulMaxNanoseconds, 10) + `",`
	output += `"mul gas":"` + strconv.Itoa(bytecodeInfoLog.mulGas) + `",`
	output += `"mul min gas":"` + strconv.Itoa(bytecodeInfoLog.mulMinGas) + `",`
	output += `"mul max gas":"` + strconv.Itoa(bytecodeInfoLog.mulMaxGas) + `",`
	output += `"mulCount":"` + strconv.Itoa(bytecodeInfoLog.mulCount) + `",`
	output += `"div ns":"` + strconv.FormatInt(bytecodeInfoLog.divNanoseconds, 10) + `",`
	output += `"div min ns":"` + strconv.FormatInt(bytecodeInfoLog.divMinNanoseconds, 10) + `",`
	output += `"div max ns":"` + strconv.FormatInt(bytecodeInfoLog.divMaxNanoseconds, 10) + `",`
	output += `"div gas":"` + strconv.Itoa(bytecodeInfoLog.divGas) + `",`
	output += `"div min gas":"` + strconv.Itoa(bytecodeInfoLog.divMinGas) + `",`
	output += `"div max gas":"` + strconv.Itoa(bytecodeInfoLog.divMaxGas) + `",`
	output += `"divCount":"` + strconv.Itoa(bytecodeInfoLog.divCount) + `",`
	output += `"sdiv ns":"` + strconv.FormatInt(bytecodeInfoLog.sdivNanoseconds, 10) + `",`
	output += `"sdiv min ns":"` + strconv.FormatInt(bytecodeInfoLog.sdivMinNanoseconds, 10) + `",`
	output += `"sdiv max ns":"` + strconv.FormatInt(bytecodeInfoLog.sdivMaxNanoseconds, 10) + `",`
	output += `"sdiv gas":"` + strconv.Itoa(bytecodeInfoLog.sdivGas) + `",`
	output += `"sdiv min gas":"` + strconv.Itoa(bytecodeInfoLog.sdivMinGas) + `",`
	output += `"sdiv max gas":"` + strconv.Itoa(bytecodeInfoLog.sdivMaxGas) + `",`
	output += `"sdivCount":"` + strconv.Itoa(bytecodeInfoLog.sdivCount) + `",`
	output += `"mod ns":"` + strconv.FormatInt(bytecodeInfoLog.modNanoseconds, 10) + `",`
	output += `"mod min ns":"` + strconv.FormatInt(bytecodeInfoLog.modMinNanoseconds, 10) + `",`
	output += `"mod max ns":"` + strconv.FormatInt(bytecodeInfoLog.modMaxNanoseconds, 10) + `",`
	output += `"mod gas":"` + strconv.Itoa(bytecodeInfoLog.modGas) + `",`
	output += `"mod min gas":"` + strconv.Itoa(bytecodeInfoLog.modMinGas) + `",`
	output += `"mod max gas":"` + strconv.Itoa(bytecodeInfoLog.modMaxGas) + `",`
	output += `"modCount":"` + strconv.Itoa(bytecodeInfoLog.modCount) + `",`
	output += `"smod ns":"` + strconv.FormatInt(bytecodeInfoLog.smodNanoseconds, 10) + `",`
	output += `"smod min ns":"` + strconv.FormatInt(bytecodeInfoLog.smodMinNanoseconds, 10) + `",`
	output += `"smod max ns":"` + strconv.FormatInt(bytecodeInfoLog.smodMaxNanoseconds, 10) + `",`
	output += `"smod gas":"` + strconv.Itoa(bytecodeInfoLog.smodGas) + `",`
	output += `"smod min gas":"` + strconv.Itoa(bytecodeInfoLog.smodMinGas) + `",`
	output += `"smod max gas":"` + strconv.Itoa(bytecodeInfoLog.smodMaxGas) + `",`
	output += `"smodCount":"` + strconv.Itoa(bytecodeInfoLog.smodCount) + `",`
	output += `"exp ns":"` + strconv.FormatInt(bytecodeInfoLog.expNanoseconds, 10) + `",`
	output += `"exp min ns":"` + strconv.FormatInt(bytecodeInfoLog.expMinNanoseconds, 10) + `",`
	output += `"exp max ns":"` + strconv.FormatInt(bytecodeInfoLog.expMaxNanoseconds, 10) + `",`
	output += `"exp gas":"` + strconv.Itoa(bytecodeInfoLog.expGas) + `",`
	output += `"exp min gas":"` + strconv.Itoa(bytecodeInfoLog.expMinGas) + `",`
	output += `"exp max gas":"` + strconv.Itoa(bytecodeInfoLog.expMaxGas) + `",`
	output += `"expCount":"` + strconv.Itoa(bytecodeInfoLog.expCount) + `",`
	output += `"signExtend ns":"` + strconv.FormatInt(bytecodeInfoLog.signExtendNanoseconds, 10) + `",`
	output += `"signExtend min ns":"` + strconv.FormatInt(bytecodeInfoLog.signExtendMinNanoseconds, 10) + `",`
	output += `"signExtend max ns":"` + strconv.FormatInt(bytecodeInfoLog.signExtendMaxNanoseconds, 10) + `",`
	output += `"signExtend gas":"` + strconv.Itoa(bytecodeInfoLog.signExtendGas) + `",`
	output += `"signExtend min gas":"` + strconv.Itoa(bytecodeInfoLog.signExtendMinGas) + `",`
	output += `"signExtend max gas":"` + strconv.Itoa(bytecodeInfoLog.signExtendMaxGas) + `",`
	output += `"signExtendCount":"` + strconv.Itoa(bytecodeInfoLog.signExtendCount) + `",`
	output += `"not ns":"` + strconv.FormatInt(bytecodeInfoLog.notNanoseconds, 10) + `",`
	output += `"not min ns":"` + strconv.FormatInt(bytecodeInfoLog.notMinNanoseconds, 10) + `",`
	output += `"not max ns":"` + strconv.FormatInt(bytecodeInfoLog.notMaxNanoseconds, 10) + `",`
	output += `"not gas":"` + strconv.Itoa(bytecodeInfoLog.notGas) + `",`
	output += `"not min gas":"` + strconv.Itoa(bytecodeInfoLog.notMinGas) + `",`
	output += `"not max gas":"` + strconv.Itoa(bytecodeInfoLog.notMaxGas) + `",`
	output += `"notCount":"` + strconv.Itoa(bytecodeInfoLog.notCount) + `",`
	output += `"lt ns":"` + strconv.FormatInt(bytecodeInfoLog.ltNanoseconds, 10) + `",`
	output += `"lt min ns":"` + strconv.FormatInt(bytecodeInfoLog.ltMinNanoseconds, 10) + `",`
	output += `"lt max ns":"` + strconv.FormatInt(bytecodeInfoLog.ltMaxNanoseconds, 10) + `",`
	output += `"lt gas":"` + strconv.Itoa(bytecodeInfoLog.ltGas) + `",`
	output += `"lt min gas":"` + strconv.Itoa(bytecodeInfoLog.ltMinGas) + `",`
	output += `"lt max gas":"` + strconv.Itoa(bytecodeInfoLog.ltMaxGas) + `",`
	output += `"ltCount":"` + strconv.Itoa(bytecodeInfoLog.ltCount) + `",`
	output += `"gt ns":"` + strconv.FormatInt(bytecodeInfoLog.gtNanoseconds, 10) + `",`
	output += `"gt min ns":"` + strconv.FormatInt(bytecodeInfoLog.gtMinNanoseconds, 10) + `",`
	output += `"gt max ns":"` + strconv.FormatInt(bytecodeInfoLog.gtMaxNanoseconds, 10) + `",`
	output += `"gt gas":"` + strconv.Itoa(bytecodeInfoLog.gtGas) + `",`
	output += `"gt min gas":"` + strconv.Itoa(bytecodeInfoLog.gtMinGas) + `",`
	output += `"gt max gas":"` + strconv.Itoa(bytecodeInfoLog.gtMaxGas) + `",`
	output += `"gtCount":"` + strconv.Itoa(bytecodeInfoLog.gtCount) + `",`
	output += `"slt ns":"` + strconv.FormatInt(bytecodeInfoLog.sltNanoseconds, 10) + `",`
	output += `"slt min ns":"` + strconv.FormatInt(bytecodeInfoLog.sltMinNanoseconds, 10) + `",`
	output += `"slt max ns":"` + strconv.FormatInt(bytecodeInfoLog.sltMaxNanoseconds, 10) + `",`
	output += `"slt gas":"` + strconv.Itoa(bytecodeInfoLog.sltGas) + `",`
	output += `"slt min gas":"` + strconv.Itoa(bytecodeInfoLog.sltMinGas) + `",`
	output += `"slt max gas":"` + strconv.Itoa(bytecodeInfoLog.sltMaxGas) + `",`
	output += `"sltCount":"` + strconv.Itoa(bytecodeInfoLog.sltCount) + `",`
	output += `"sgt ns":"` + strconv.FormatInt(bytecodeInfoLog.sgtNanoseconds, 10) + `",`
	output += `"sgt min ns":"` + strconv.FormatInt(bytecodeInfoLog.sgtMinNanoseconds, 10) + `",`
	output += `"sgt max ns":"` + strconv.FormatInt(bytecodeInfoLog.sgtMaxNanoseconds, 10) + `",`
	output += `"sgt gas":"` + strconv.Itoa(bytecodeInfoLog.sgtGas) + `",`
	output += `"sgt min gas":"` + strconv.Itoa(bytecodeInfoLog.sgtMinGas) + `",`
	output += `"sgt max gas":"` + strconv.Itoa(bytecodeInfoLog.sgtMaxGas) + `",`
	output += `"sgtCount":"` + strconv.Itoa(bytecodeInfoLog.sgtCount) + `",`
	output += `"eq ns":"` + strconv.FormatInt(bytecodeInfoLog.eqNanoseconds, 10) + `",`
	output += `"eq min ns":"` + strconv.FormatInt(bytecodeInfoLog.eqMinNanoseconds, 10) + `",`
	output += `"eq max ns":"` + strconv.FormatInt(bytecodeInfoLog.eqMaxNanoseconds, 10) + `",`
	output += `"eq gas":"` + strconv.Itoa(bytecodeInfoLog.eqGas) + `",`
	output += `"eq min gas":"` + strconv.Itoa(bytecodeInfoLog.eqMinGas) + `",`
	output += `"eq max gas":"` + strconv.Itoa(bytecodeInfoLog.eqMaxGas) + `",`
	output += `"eqCount":"` + strconv.Itoa(bytecodeInfoLog.eqCount) + `",`
	output += `"isZero ns":"` + strconv.FormatInt(bytecodeInfoLog.isZeroNanoseconds, 10) + `",`
	output += `"isZero min ns":"` + strconv.FormatInt(bytecodeInfoLog.isZeroMinNanoseconds, 10) + `",`
	output += `"isZero max ns":"` + strconv.FormatInt(bytecodeInfoLog.isZeroMaxNanoseconds, 10) + `",`
	output += `"isZero gas":"` + strconv.Itoa(bytecodeInfoLog.isZeroGas) + `",`
	output += `"isZero min gas":"` + strconv.Itoa(bytecodeInfoLog.isZeroMinGas) + `",`
	output += `"isZero max gas":"` + strconv.Itoa(bytecodeInfoLog.isZeroMaxGas) + `",`
	output += `"isZeroCount":"` + strconv.Itoa(bytecodeInfoLog.isZeroCount) + `",`
	output += `"and ns":"` + strconv.FormatInt(bytecodeInfoLog.andNanoseconds, 10) + `",`
	output += `"and min ns":"` + strconv.FormatInt(bytecodeInfoLog.andMinNanoseconds, 10) + `",`
	output += `"and max ns":"` + strconv.FormatInt(bytecodeInfoLog.andMaxNanoseconds, 10) + `",`
	output += `"and gas":"` + strconv.Itoa(bytecodeInfoLog.andGas) + `",`
	output += `"and min gas":"` + strconv.Itoa(bytecodeInfoLog.andMinGas) + `",`
	output += `"and max gas":"` + strconv.Itoa(bytecodeInfoLog.andMaxGas) + `",`
	output += `"andCount":"` + strconv.Itoa(bytecodeInfoLog.andCount) + `",`
	output += `"or ns":"` + strconv.FormatInt(bytecodeInfoLog.orNanoseconds, 10) + `",`
	output += `"or min ns":"` + strconv.FormatInt(bytecodeInfoLog.orMinNanoseconds, 10) + `",`
	output += `"or max ns":"` + strconv.FormatInt(bytecodeInfoLog.orMaxNanoseconds, 10) + `",`
	output += `"or gas":"` + strconv.Itoa(bytecodeInfoLog.orGas) + `",`
	output += `"or min gas":"` + strconv.Itoa(bytecodeInfoLog.orMinGas) + `",`
	output += `"or max gas":"` + strconv.Itoa(bytecodeInfoLog.orMaxGas) + `",`
	output += `"orCount":"` + strconv.Itoa(bytecodeInfoLog.orCount) + `",`
	output += `"xor ns":"` + strconv.FormatInt(bytecodeInfoLog.xorNanoseconds, 10) + `",`
	output += `"xor min ns":"` + strconv.FormatInt(bytecodeInfoLog.xorMinNanoseconds, 10) + `",`
	output += `"xor max ns":"` + strconv.FormatInt(bytecodeInfoLog.xorMaxNanoseconds, 10) + `",`
	output += `"xor gas":"` + strconv.Itoa(bytecodeInfoLog.xorGas) + `",`
	output += `"xor min gas":"` + strconv.Itoa(bytecodeInfoLog.xorMinGas) + `",`
	output += `"xor max gas":"` + strconv.Itoa(bytecodeInfoLog.xorMaxGas) + `",`
	output += `"xorCount":"` + strconv.Itoa(bytecodeInfoLog.xorCount) + `",`
	output += `"byte ns":"` + strconv.FormatInt(bytecodeInfoLog.byteNanoseconds, 10) + `",`
	output += `"byte min ns":"` + strconv.FormatInt(bytecodeInfoLog.byteMinNanoseconds, 10) + `",`
	output += `"byte max ns":"` + strconv.FormatInt(bytecodeInfoLog.byteMaxNanoseconds, 10) + `",`
	output += `"byte gas":"` + strconv.Itoa(bytecodeInfoLog.byteGas) + `",`
	output += `"byte min gas":"` + strconv.Itoa(bytecodeInfoLog.byteMinGas) + `",`
	output += `"byte max gas":"` + strconv.Itoa(bytecodeInfoLog.byteMaxGas) + `",`
	output += `"byteCount":"` + strconv.Itoa(bytecodeInfoLog.byteCount) + `",`
	output += `"addmod ns":"` + strconv.FormatInt(bytecodeInfoLog.addmodNanoseconds, 10) + `",`
	output += `"addmod min ns":"` + strconv.FormatInt(bytecodeInfoLog.addmodMinNanoseconds, 10) + `",`
	output += `"addmod max ns":"` + strconv.FormatInt(bytecodeInfoLog.addmodMaxNanoseconds, 10) + `",`
	output += `"addmod gas":"` + strconv.Itoa(bytecodeInfoLog.addmodGas) + `",`
	output += `"addmod min gas":"` + strconv.Itoa(bytecodeInfoLog.addmodMinGas) + `",`
	output += `"addmod max gas":"` + strconv.Itoa(bytecodeInfoLog.addmodMaxGas) + `",`
	output += `"addmodCount":"` + strconv.Itoa(bytecodeInfoLog.addmodCount) + `",`
	output += `"mulmod ns":"` + strconv.FormatInt(bytecodeInfoLog.mulmodNanoseconds, 10) + `",`
	output += `"mulmod min ns":"` + strconv.FormatInt(bytecodeInfoLog.mulmodMinNanoseconds, 10) + `",`
	output += `"mulmod max ns":"` + strconv.FormatInt(bytecodeInfoLog.mulmodMaxNanoseconds, 10) + `",`
	output += `"mulmod gas":"` + strconv.Itoa(bytecodeInfoLog.mulmodGas) + `",`
	output += `"mulmod min gas":"` + strconv.Itoa(bytecodeInfoLog.mulmodMinGas) + `",`
	output += `"mulmod max gas":"` + strconv.Itoa(bytecodeInfoLog.mulmodMaxGas) + `",`
	output += `"mulmodCount":"` + strconv.Itoa(bytecodeInfoLog.mulmodCount) + `",`
	output += `"shl ns":"` + strconv.FormatInt(bytecodeInfoLog.shlNanoseconds, 10) + `",`
	output += `"shl min ns":"` + strconv.FormatInt(bytecodeInfoLog.shlMinNanoseconds, 10) + `",`
	output += `"shl max ns":"` + strconv.FormatInt(bytecodeInfoLog.shlMaxNanoseconds, 10) + `",`
	output += `"shl gas":"` + strconv.Itoa(bytecodeInfoLog.shlGas) + `",`
	output += `"shl min gas":"` + strconv.Itoa(bytecodeInfoLog.shlMinGas) + `",`
	output += `"shl max gas":"` + strconv.Itoa(bytecodeInfoLog.shlMaxGas) + `",`
	output += `"shlCount":"` + strconv.Itoa(bytecodeInfoLog.shlCount) + `",`
	output += `"shr ns":"` + strconv.FormatInt(bytecodeInfoLog.shrNanoseconds, 10) + `",`
	output += `"shr min ns":"` + strconv.FormatInt(bytecodeInfoLog.shrMinNanoseconds, 10) + `",`
	output += `"shr max ns":"` + strconv.FormatInt(bytecodeInfoLog.shrMaxNanoseconds, 10) + `",`
	output += `"shr gas":"` + strconv.Itoa(bytecodeInfoLog.shrGas) + `",`
	output += `"shr min gas":"` + strconv.Itoa(bytecodeInfoLog.shrMinGas) + `",`
	output += `"shr max gas":"` + strconv.Itoa(bytecodeInfoLog.shrMaxGas) + `",`
	output += `"shrCount":"` + strconv.Itoa(bytecodeInfoLog.shrCount) + `",`
	output += `"sar ns":"` + strconv.FormatInt(bytecodeInfoLog.sarNanoseconds, 10) + `",`
	output += `"sar min ns":"` + strconv.FormatInt(bytecodeInfoLog.sarMinNanoseconds, 10) + `",`
	output += `"sar max ns":"` + strconv.FormatInt(bytecodeInfoLog.sarMaxNanoseconds, 10) + `",`
	output += `"sar gas":"` + strconv.Itoa(bytecodeInfoLog.sarGas) + `",`
	output += `"sar min gas":"` + strconv.Itoa(bytecodeInfoLog.sarMinGas) + `",`
	output += `"sar max gas":"` + strconv.Itoa(bytecodeInfoLog.sarMaxGas) + `",`
	output += `"sarCount":"` + strconv.Itoa(bytecodeInfoLog.sarCount) + `",`
	output += `"keccak256 ns":"` + strconv.FormatInt(bytecodeInfoLog.keccak256Nanoseconds, 10) + `",`
	output += `"keccak256 min ns":"` + strconv.FormatInt(bytecodeInfoLog.keccak256MinNanoseconds, 10) + `",`
	output += `"keccak256 max ns":"` + strconv.FormatInt(bytecodeInfoLog.keccak256MaxNanoseconds, 10) + `",`
	output += `"keccak256 gas":"` + strconv.Itoa(bytecodeInfoLog.keccak256Gas) + `",`
	output += `"keccak256 min gas":"` + strconv.Itoa(bytecodeInfoLog.keccak256MinGas) + `",`
	output += `"keccak256 max gas":"` + strconv.Itoa(bytecodeInfoLog.keccak256MaxGas) + `",`
	output += `"keccak256Count":"` + strconv.Itoa(bytecodeInfoLog.keccak256Count) + `",`
	output += `"address ns":"` + strconv.FormatInt(bytecodeInfoLog.addressNanoseconds, 10) + `",`
	output += `"address min ns":"` + strconv.FormatInt(bytecodeInfoLog.addressMinNanoseconds, 10) + `",`
	output += `"address max ns":"` + strconv.FormatInt(bytecodeInfoLog.addressMaxNanoseconds, 10) + `",`
	output += `"address gas":"` + strconv.Itoa(bytecodeInfoLog.addressGas) + `",`
	output += `"address min gas":"` + strconv.Itoa(bytecodeInfoLog.addressMinGas) + `",`
	output += `"address max gas":"` + strconv.Itoa(bytecodeInfoLog.addressMaxGas) + `",`
	output += `"addressCount":"` + strconv.Itoa(bytecodeInfoLog.addressCount) + `",`
	output += `"balance ns":"` + strconv.FormatInt(bytecodeInfoLog.balanceNanoseconds, 10) + `",`
	output += `"balance min ns":"` + strconv.FormatInt(bytecodeInfoLog.balanceMinNanoseconds, 10) + `",`
	output += `"balance max ns":"` + strconv.FormatInt(bytecodeInfoLog.balanceMaxNanoseconds, 10) + `",`
	output += `"balance gas":"` + strconv.Itoa(bytecodeInfoLog.balanceGas) + `",`
	output += `"balance min gas":"` + strconv.Itoa(bytecodeInfoLog.balanceMinGas) + `",`
	output += `"balance max gas":"` + strconv.Itoa(bytecodeInfoLog.balanceMaxGas) + `",`
	output += `"balanceCount":"` + strconv.Itoa(bytecodeInfoLog.balanceCount) + `",`
	output += `"origin ns":"` + strconv.FormatInt(bytecodeInfoLog.originNanoseconds, 10) + `",`
	output += `"origin min ns":"` + strconv.FormatInt(bytecodeInfoLog.originMinNanoseconds, 10) + `",`
	output += `"origin max ns":"` + strconv.FormatInt(bytecodeInfoLog.originMaxNanoseconds, 10) + `",`
	output += `"origin gas":"` + strconv.Itoa(bytecodeInfoLog.originGas) + `",`
	output += `"origin min gas":"` + strconv.Itoa(bytecodeInfoLog.originMinGas) + `",`
	output += `"origin max gas":"` + strconv.Itoa(bytecodeInfoLog.originMaxGas) + `",`
	output += `"originCount":"` + strconv.Itoa(bytecodeInfoLog.originCount) + `",`
	output += `"caller ns":"` + strconv.FormatInt(bytecodeInfoLog.callerNanoseconds, 10) + `",`
	output += `"caller min ns":"` + strconv.FormatInt(bytecodeInfoLog.callerMinNanoseconds, 10) + `",`
	output += `"caller max ns":"` + strconv.FormatInt(bytecodeInfoLog.callerMaxNanoseconds, 10) + `",`
	output += `"caller gas":"` + strconv.Itoa(bytecodeInfoLog.callerGas) + `",`
	output += `"caller min gas":"` + strconv.Itoa(bytecodeInfoLog.callerMinGas) + `",`
	output += `"caller max gas":"` + strconv.Itoa(bytecodeInfoLog.callerMaxGas) + `",`
	output += `"callerCount":"` + strconv.Itoa(bytecodeInfoLog.callerCount) + `",`
	output += `"callValue ns":"` + strconv.FormatInt(bytecodeInfoLog.callValueNanoseconds, 10) + `",`
	output += `"callValue min ns":"` + strconv.FormatInt(bytecodeInfoLog.callValueMinNanoseconds, 10) + `",`
	output += `"callValue max ns":"` + strconv.FormatInt(bytecodeInfoLog.callValueMaxNanoseconds, 10) + `",`
	output += `"callValue gas":"` + strconv.Itoa(bytecodeInfoLog.callValueGas) + `",`
	output += `"callValue min gas":"` + strconv.Itoa(bytecodeInfoLog.callValueMinGas) + `",`
	output += `"callValue max gas":"` + strconv.Itoa(bytecodeInfoLog.callValueMaxGas) + `",`
	output += `"callValueCount":"` + strconv.Itoa(bytecodeInfoLog.callValueCount) + `",`
	output += `"callDataLoad ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataLoadNanoseconds, 10) + `",`
	output += `"callDataLoad min ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataLoadMinNanoseconds, 10) + `",`
	output += `"callDataLoad max ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataLoadMaxNanoseconds, 10) + `",`
	output += `"callDataLoad gas":"` + strconv.Itoa(bytecodeInfoLog.callDataLoadGas) + `",`
	output += `"callDataLoad min gas":"` + strconv.Itoa(bytecodeInfoLog.callDataLoadMinGas) + `",`
	output += `"callDataLoad max gas":"` + strconv.Itoa(bytecodeInfoLog.callDataLoadMaxGas) + `",`
	output += `"callDataLoadCount":"` + strconv.Itoa(bytecodeInfoLog.callDataLoadCount) + `",`
	output += `"callDataSize ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataSizeNanoseconds, 10) + `",`
	output += `"callDataSize min ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataSizeMinNanoseconds, 10) + `",`
	output += `"callDataSize max ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataSizeMaxNanoseconds, 10) + `",`
	output += `"callDataSize gas":"` + strconv.Itoa(bytecodeInfoLog.callDataSizeGas) + `",`
	output += `"callDataSize min gas":"` + strconv.Itoa(bytecodeInfoLog.callDataSizeMinGas) + `",`
	output += `"callDataSize max gas":"` + strconv.Itoa(bytecodeInfoLog.callDataSizeMaxGas) + `",`
	output += `"callDataSizeCount":"` + strconv.Itoa(bytecodeInfoLog.callDataSizeCount) + `",`
	output += `"callDataCopy ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataCopyNanoseconds, 10) + `",`
	output += `"callDataCopy min ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataCopyMinNanoseconds, 10) + `",`
	output += `"callDataCopy max ns":"` + strconv.FormatInt(bytecodeInfoLog.callDataCopyMaxNanoseconds, 10) + `",`
	output += `"callDataCopy gas":"` + strconv.Itoa(bytecodeInfoLog.callDataCopyGas) + `",`
	output += `"callDataCopy min gas":"` + strconv.Itoa(bytecodeInfoLog.callDataCopyMinGas) + `",`
	output += `"callDataCopy max gas":"` + strconv.Itoa(bytecodeInfoLog.callDataCopyMaxGas) + `",`
	output += `"callDataCopyCount":"` + strconv.Itoa(bytecodeInfoLog.callDataCopyCount) + `",`
	output += `"returnDataSize ns":"` + strconv.FormatInt(bytecodeInfoLog.returnDataSizeNanoseconds, 10) + `",`
	output += `"returnDataSize min ns":"` + strconv.FormatInt(bytecodeInfoLog.returnDataSizeMinNanoseconds, 10) + `",`
	output += `"returnDataSize max ns":"` + strconv.FormatInt(bytecodeInfoLog.returnDataSizeMaxNanoseconds, 10) + `",`
	output += `"returnDataSize gas":"` + strconv.Itoa(bytecodeInfoLog.returnDataSizeGas) + `",`
	output += `"returnDataSize min gas":"` + strconv.Itoa(bytecodeInfoLog.returnDataSizeMinGas) + `",`
	output += `"returnDataSize max gas":"` + strconv.Itoa(bytecodeInfoLog.returnDataSizeMaxGas) + `",`
	output += `"returnDataSizeCount":"` + strconv.Itoa(bytecodeInfoLog.returnDataSizeCount) + `",`
	output += `"returnDataCopy ns":"` + strconv.FormatInt(bytecodeInfoLog.returnDataCopyNanoseconds, 10) + `",`
	output += `"returnDataCopy min ns":"` + strconv.FormatInt(bytecodeInfoLog.returnDataCopyMinNanoseconds, 10) + `",`
	output += `"returnDataCopy max ns":"` + strconv.FormatInt(bytecodeInfoLog.returnDataCopyMaxNanoseconds, 10) + `",`
	output += `"returnDataCopy gas":"` + strconv.Itoa(bytecodeInfoLog.returnDataCopyGas) + `",`
	output += `"returnDataCopy min gas":"` + strconv.Itoa(bytecodeInfoLog.returnDataCopyMinGas) + `",`
	output += `"returnDataCopy max gas":"` + strconv.Itoa(bytecodeInfoLog.returnDataCopyMaxGas) + `",`
	output += `"returnDataCopyCount":"` + strconv.Itoa(bytecodeInfoLog.returnDataCopyCount) + `",`
	output += `"extCodeSize ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeSizeNanoseconds, 10) + `",`
	output += `"extCodeSize min ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeSizeMinNanoseconds, 10) + `",`
	output += `"extCodeSize max ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeSizeMaxNanoseconds, 10) + `",`
	output += `"extCodeSize gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeSizeGas) + `",`
	output += `"extCodeSize min gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeSizeMinGas) + `",`
	output += `"extCodeSize max gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeSizeMaxGas) + `",`
	output += `"extCodeSizeCount":"` + strconv.Itoa(bytecodeInfoLog.extCodeSizeCount) + `",`
	output += `"codeSize ns":"` + strconv.FormatInt(bytecodeInfoLog.codeSizeNanoseconds, 10) + `",`
	output += `"codeSize min ns":"` + strconv.FormatInt(bytecodeInfoLog.codeSizeMinNanoseconds, 10) + `",`
	output += `"codeSize max ns":"` + strconv.FormatInt(bytecodeInfoLog.codeSizeMaxNanoseconds, 10) + `",`
	output += `"codeSize gas":"` + strconv.Itoa(bytecodeInfoLog.codeSizeGas) + `",`
	output += `"codeSize min gas":"` + strconv.Itoa(bytecodeInfoLog.codeSizeMinGas) + `",`
	output += `"codeSize max gas":"` + strconv.Itoa(bytecodeInfoLog.codeSizeMaxGas) + `",`
	output += `"codeSizeCount":"` + strconv.Itoa(bytecodeInfoLog.codeSizeCount) + `",`
	output += `"codeCopy ns":"` + strconv.FormatInt(bytecodeInfoLog.codeCopyNanoseconds, 10) + `",`
	output += `"codeCopy min ns":"` + strconv.FormatInt(bytecodeInfoLog.codeCopyMinNanoseconds, 10) + `",`
	output += `"codeCopy max ns":"` + strconv.FormatInt(bytecodeInfoLog.codeCopyMaxNanoseconds, 10) + `",`
	output += `"codeCopy gas":"` + strconv.Itoa(bytecodeInfoLog.codeCopyGas) + `",`
	output += `"codeCopy min gas":"` + strconv.Itoa(bytecodeInfoLog.codeCopyMinGas) + `",`
	output += `"codeCopy max gas":"` + strconv.Itoa(bytecodeInfoLog.codeCopyMaxGas) + `",`
	output += `"codeCopyCount":"` + strconv.Itoa(bytecodeInfoLog.codeCopyCount) + `",`
	output += `"extCodeCopy ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeCopyNanoseconds, 10) + `",`
	output += `"extCodeCopy min ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeCopyMinNanoseconds, 10) + `",`
	output += `"extCodeCopy max ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeCopyMaxNanoseconds, 10) + `",`
	output += `"extCodeCopy gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeCopyGas) + `",`
	output += `"extCodeCopy min gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeCopyMinGas) + `",`
	output += `"extCodeCopy max gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeCopyMaxGas) + `",`
	output += `"extCodeCopyCount":"` + strconv.Itoa(bytecodeInfoLog.extCodeCopyCount) + `",`
	output += `"extCodeHash ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeHashNanoseconds, 10) + `",`
	output += `"extCodeHash min ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeHashMinNanoseconds, 10) + `",`
	output += `"extCodeHash max ns":"` + strconv.FormatInt(bytecodeInfoLog.extCodeHashMaxNanoseconds, 10) + `",`
	output += `"extCodeHash gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeHashGas) + `",`
	output += `"extCodeHash min gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeHashMinGas) + `",`
	output += `"extCodeHash max gas":"` + strconv.Itoa(bytecodeInfoLog.extCodeHashMaxGas) + `",`
	output += `"extCodeHashCount":"` + strconv.Itoa(bytecodeInfoLog.extCodeHashCount) + `",`
	output += `"gasprice ns":"` + strconv.FormatInt(bytecodeInfoLog.gaspriceNanoseconds, 10) + `",`
	output += `"gasprice min ns":"` + strconv.FormatInt(bytecodeInfoLog.gaspriceMinNanoseconds, 10) + `",`
	output += `"gasprice max ns":"` + strconv.FormatInt(bytecodeInfoLog.gaspriceMaxNanoseconds, 10) + `",`
	output += `"gasprice gas":"` + strconv.Itoa(bytecodeInfoLog.gaspriceGas) + `",`
	output += `"gasprice min gas":"` + strconv.Itoa(bytecodeInfoLog.gaspriceMinGas) + `",`
	output += `"gasprice max gas":"` + strconv.Itoa(bytecodeInfoLog.gaspriceMaxGas) + `",`
	output += `"gaspriceCount":"` + strconv.Itoa(bytecodeInfoLog.gaspriceCount) + `",`
	output += `"blockhash ns":"` + strconv.FormatInt(bytecodeInfoLog.blockhashNanoseconds, 10) + `",`
	output += `"blockhash min ns":"` + strconv.FormatInt(bytecodeInfoLog.blockhashMinNanoseconds, 10) + `",`
	output += `"blockhash max ns":"` + strconv.FormatInt(bytecodeInfoLog.blockhashMaxNanoseconds, 10) + `",`
	output += `"blockhash gas":"` + strconv.Itoa(bytecodeInfoLog.blockhashGas) + `",`
	output += `"blockhash min gas":"` + strconv.Itoa(bytecodeInfoLog.blockhashMinGas) + `",`
	output += `"blockhash max gas":"` + strconv.Itoa(bytecodeInfoLog.blockhashMaxGas) + `",`
	output += `"blockhashCount":"` + strconv.Itoa(bytecodeInfoLog.blockhashCount) + `",`
	output += `"coinbase ns":"` + strconv.FormatInt(bytecodeInfoLog.coinbaseNanoseconds, 10) + `",`
	output += `"coinbase min ns":"` + strconv.FormatInt(bytecodeInfoLog.coinbaseMinNanoseconds, 10) + `",`
	output += `"coinbase max ns":"` + strconv.FormatInt(bytecodeInfoLog.coinbaseMaxNanoseconds, 10) + `",`
	output += `"coinbase gas":"` + strconv.Itoa(bytecodeInfoLog.coinbaseGas) + `",`
	output += `"coinbase min gas":"` + strconv.Itoa(bytecodeInfoLog.coinbaseMinGas) + `",`
	output += `"coinbase max gas":"` + strconv.Itoa(bytecodeInfoLog.coinbaseMaxGas) + `",`
	output += `"coinbaseCount":"` + strconv.Itoa(bytecodeInfoLog.coinbaseCount) + `",`
	output += `"timestamp ns":"` + strconv.FormatInt(bytecodeInfoLog.timestampNanoseconds, 10) + `",`
	output += `"timestamp min ns":"` + strconv.FormatInt(bytecodeInfoLog.timestampMinNanoseconds, 10) + `",`
	output += `"timestamp max ns":"` + strconv.FormatInt(bytecodeInfoLog.timestampMaxNanoseconds, 10) + `",`
	output += `"timestamp gas":"` + strconv.Itoa(bytecodeInfoLog.timestampGas) + `",`
	output += `"timestamp min gas":"` + strconv.Itoa(bytecodeInfoLog.timestampMinGas) + `",`
	output += `"timestamp max gas":"` + strconv.Itoa(bytecodeInfoLog.timestampMaxGas) + `",`
	output += `"timestampCount":"` + strconv.Itoa(bytecodeInfoLog.timestampCount) + `",`
	output += `"number ns":"` + strconv.FormatInt(bytecodeInfoLog.numberNanoseconds, 10) + `",`
	output += `"number min ns":"` + strconv.FormatInt(bytecodeInfoLog.numberMinNanoseconds, 10) + `",`
	output += `"number max ns":"` + strconv.FormatInt(bytecodeInfoLog.numberMaxNanoseconds, 10) + `",`
	output += `"number gas":"` + strconv.Itoa(bytecodeInfoLog.numberGas) + `",`
	output += `"number min gas":"` + strconv.Itoa(bytecodeInfoLog.numberMinGas) + `",`
	output += `"number max gas":"` + strconv.Itoa(bytecodeInfoLog.numberMaxGas) + `",`
	output += `"numberCount":"` + strconv.Itoa(bytecodeInfoLog.numberCount) + `",`
	output += `"difficulty ns":"` + strconv.FormatInt(bytecodeInfoLog.difficultyNanoseconds, 10) + `",`
	output += `"difficulty min ns":"` + strconv.FormatInt(bytecodeInfoLog.difficultyMinNanoseconds, 10) + `",`
	output += `"difficulty max ns":"` + strconv.FormatInt(bytecodeInfoLog.difficultyMaxNanoseconds, 10) + `",`
	output += `"difficulty gas":"` + strconv.Itoa(bytecodeInfoLog.difficultyGas) + `",`
	output += `"difficulty min gas":"` + strconv.Itoa(bytecodeInfoLog.difficultyMinGas) + `",`
	output += `"difficulty max gas":"` + strconv.Itoa(bytecodeInfoLog.difficultyMaxGas) + `",`
	output += `"difficultyCount":"` + strconv.Itoa(bytecodeInfoLog.difficultyCount) + `",`
	output += `"random ns":"` + strconv.FormatInt(bytecodeInfoLog.randomNanoseconds, 10) + `",`
	output += `"random min ns":"` + strconv.FormatInt(bytecodeInfoLog.randomMinNanoseconds, 10) + `",`
	output += `"random max ns":"` + strconv.FormatInt(bytecodeInfoLog.randomMaxNanoseconds, 10) + `",`
	output += `"random gas":"` + strconv.Itoa(bytecodeInfoLog.randomGas) + `",`
	output += `"random min gas":"` + strconv.Itoa(bytecodeInfoLog.randomMinGas) + `",`
	output += `"random max gas":"` + strconv.Itoa(bytecodeInfoLog.randomMaxGas) + `",`
	output += `"randomCount":"` + strconv.Itoa(bytecodeInfoLog.randomCount) + `",`
	output += `"gasLimit ns":"` + strconv.FormatInt(bytecodeInfoLog.gasLimitNanoseconds, 10) + `",`
	output += `"gasLimit min ns":"` + strconv.FormatInt(bytecodeInfoLog.gasLimitMinNanoseconds, 10) + `",`
	output += `"gasLimit max ns":"` + strconv.FormatInt(bytecodeInfoLog.gasLimitMaxNanoseconds, 10) + `",`
	output += `"gasLimit gas":"` + strconv.Itoa(bytecodeInfoLog.gasLimitGas) + `",`
	output += `"gasLimit min gas":"` + strconv.Itoa(bytecodeInfoLog.gasLimitMinGas) + `",`
	output += `"gasLimit max gas":"` + strconv.Itoa(bytecodeInfoLog.gasLimitMaxGas) + `",`
	output += `"gasLimitCount":"` + strconv.Itoa(bytecodeInfoLog.gasLimitCount) + `",`
	output += `"pop ns":"` + strconv.FormatInt(bytecodeInfoLog.popNanoseconds, 10) + `",`
	output += `"pop min ns":"` + strconv.FormatInt(bytecodeInfoLog.popMinNanoseconds, 10) + `",`
	output += `"pop max ns":"` + strconv.FormatInt(bytecodeInfoLog.popMaxNanoseconds, 10) + `",`
	output += `"pop gas":"` + strconv.Itoa(bytecodeInfoLog.popGas) + `",`
	output += `"pop min gas":"` + strconv.Itoa(bytecodeInfoLog.popMinGas) + `",`
	output += `"pop max gas":"` + strconv.Itoa(bytecodeInfoLog.popMaxGas) + `",`
	output += `"popCount":"` + strconv.Itoa(bytecodeInfoLog.popCount) + `",`
	output += `"mload ns":"` + strconv.FormatInt(bytecodeInfoLog.mloadNanoseconds, 10) + `",`
	output += `"mload min ns":"` + strconv.FormatInt(bytecodeInfoLog.mloadMinNanoseconds, 10) + `",`
	output += `"mload max ns":"` + strconv.FormatInt(bytecodeInfoLog.mloadMaxNanoseconds, 10) + `",`
	output += `"mload gas":"` + strconv.Itoa(bytecodeInfoLog.mloadGas) + `",`
	output += `"mload min gas":"` + strconv.Itoa(bytecodeInfoLog.mloadMinGas) + `",`
	output += `"mload max gas":"` + strconv.Itoa(bytecodeInfoLog.mloadMaxGas) + `",`
	output += `"mloadCount":"` + strconv.Itoa(bytecodeInfoLog.mloadCount) + `",`
	output += `"mstore ns":"` + strconv.FormatInt(bytecodeInfoLog.mstoreNanoseconds, 10) + `",`
	output += `"mstore min ns":"` + strconv.FormatInt(bytecodeInfoLog.mstoreMinNanoseconds, 10) + `",`
	output += `"mstore max ns":"` + strconv.FormatInt(bytecodeInfoLog.mstoreMaxNanoseconds, 10) + `",`
	output += `"mstore gas":"` + strconv.Itoa(bytecodeInfoLog.mstoreGas) + `",`
	output += `"mstore min gas":"` + strconv.Itoa(bytecodeInfoLog.mstoreMinGas) + `",`
	output += `"mstore max gas":"` + strconv.Itoa(bytecodeInfoLog.mstoreMaxGas) + `",`
	output += `"mstoreCount":"` + strconv.Itoa(bytecodeInfoLog.mstoreCount) + `",`
	output += `"mstore8 ns":"` + strconv.FormatInt(bytecodeInfoLog.mstore8Nanoseconds, 10) + `",`
	output += `"mstore8 min ns":"` + strconv.FormatInt(bytecodeInfoLog.mstore8MinNanoseconds, 10) + `",`
	output += `"mstore8 max ns":"` + strconv.FormatInt(bytecodeInfoLog.mstore8MaxNanoseconds, 10) + `",`
	output += `"mstore8 gas":"` + strconv.Itoa(bytecodeInfoLog.mstore8Gas) + `",`
	output += `"mstore8 min gas":"` + strconv.Itoa(bytecodeInfoLog.mstore8MinGas) + `",`
	output += `"mstore8 max gas":"` + strconv.Itoa(bytecodeInfoLog.mstore8MaxGas) + `",`
	output += `"mstore8Count":"` + strconv.Itoa(bytecodeInfoLog.mstore8Count) + `",`
	output += `"sload ns":"` + strconv.FormatInt(bytecodeInfoLog.sloadNanoseconds, 10) + `",`
	output += `"sload min ns":"` + strconv.FormatInt(bytecodeInfoLog.sloadMinNanoseconds, 10) + `",`
	output += `"sload max ns":"` + strconv.FormatInt(bytecodeInfoLog.sloadMaxNanoseconds, 10) + `",`
	output += `"sload gas":"` + strconv.Itoa(bytecodeInfoLog.sloadGas) + `",`
	output += `"sload min gas":"` + strconv.Itoa(bytecodeInfoLog.sloadMinGas) + `",`
	output += `"sload max gas":"` + strconv.Itoa(bytecodeInfoLog.sloadMaxGas) + `",`
	output += `"sloadCount":"` + strconv.Itoa(bytecodeInfoLog.sloadCount) + `",`
	output += `"sstore ns":"` + strconv.FormatInt(bytecodeInfoLog.sstoreNanoseconds, 10) + `",`
	output += `"sstore min ns":"` + strconv.FormatInt(bytecodeInfoLog.sstoreMinNanoseconds, 10) + `",`
	output += `"sstore max ns":"` + strconv.FormatInt(bytecodeInfoLog.sstoreMaxNanoseconds, 10) + `",`
	output += `"sstore gas":"` + strconv.Itoa(bytecodeInfoLog.sstoreGas) + `",`
	output += `"sstore min gas":"` + strconv.Itoa(bytecodeInfoLog.sstoreMinGas) + `",`
	output += `"sstore max gas":"` + strconv.Itoa(bytecodeInfoLog.sstoreMaxGas) + `",`
	output += `"sstoreCount":"` + strconv.Itoa(bytecodeInfoLog.sstoreCount) + `",`
	output += `"jump ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpNanoseconds, 10) + `",`
	output += `"jump min ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpMinNanoseconds, 10) + `",`
	output += `"jump max ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpMaxNanoseconds, 10) + `",`
	output += `"jump gas":"` + strconv.Itoa(bytecodeInfoLog.jumpGas) + `",`
	output += `"jump min gas":"` + strconv.Itoa(bytecodeInfoLog.jumpMinGas) + `",`
	output += `"jump max gas":"` + strconv.Itoa(bytecodeInfoLog.jumpMaxGas) + `",`
	output += `"jumpCount":"` + strconv.Itoa(bytecodeInfoLog.jumpCount) + `",`
	output += `"jumpi ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpiNanoseconds, 10) + `",`
	output += `"jumpi min ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpiMinNanoseconds, 10) + `",`
	output += `"jumpi max ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpiMaxNanoseconds, 10) + `",`
	output += `"jumpi gas":"` + strconv.Itoa(bytecodeInfoLog.jumpiGas) + `",`
	output += `"jumpi min gas":"` + strconv.Itoa(bytecodeInfoLog.jumpiMinGas) + `",`
	output += `"jumpi max gas":"` + strconv.Itoa(bytecodeInfoLog.jumpiMaxGas) + `",`
	output += `"jumpiCount":"` + strconv.Itoa(bytecodeInfoLog.jumpiCount) + `",`
	output += `"jumpdest ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpdestNanoseconds, 10) + `",`
	output += `"jumpdest min ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpdestMinNanoseconds, 10) + `",`
	output += `"jumpdest max ns":"` + strconv.FormatInt(bytecodeInfoLog.jumpdestMaxNanoseconds, 10) + `",`
	output += `"jumpdest gas":"` + strconv.Itoa(bytecodeInfoLog.jumpdestGas) + `",`
	output += `"jumpdest min gas":"` + strconv.Itoa(bytecodeInfoLog.jumpdestMinGas) + `",`
	output += `"jumpdest max gas":"` + strconv.Itoa(bytecodeInfoLog.jumpdestMaxGas) + `",`
	output += `"jumpdestCount":"` + strconv.Itoa(bytecodeInfoLog.jumpdestCount) + `",`
	output += `"pc ns":"` + strconv.FormatInt(bytecodeInfoLog.pcNanoseconds, 10) + `",`
	output += `"pc min ns":"` + strconv.FormatInt(bytecodeInfoLog.pcMinNanoseconds, 10) + `",`
	output += `"pc max ns":"` + strconv.FormatInt(bytecodeInfoLog.pcMaxNanoseconds, 10) + `",`
	output += `"pc gas":"` + strconv.Itoa(bytecodeInfoLog.pcGas) + `",`
	output += `"pc min gas":"` + strconv.Itoa(bytecodeInfoLog.pcMinGas) + `",`
	output += `"pc max gas":"` + strconv.Itoa(bytecodeInfoLog.pcMaxGas) + `",`
	output += `"pcCount":"` + strconv.Itoa(bytecodeInfoLog.pcCount) + `",`
	output += `"msize ns":"` + strconv.FormatInt(bytecodeInfoLog.msizeNanoseconds, 10) + `",`
	output += `"msize min ns":"` + strconv.FormatInt(bytecodeInfoLog.msizeMinNanoseconds, 10) + `",`
	output += `"msize max ns":"` + strconv.FormatInt(bytecodeInfoLog.msizeMaxNanoseconds, 10) + `",`
	output += `"msize gas":"` + strconv.Itoa(bytecodeInfoLog.msizeGas) + `",`
	output += `"msize min gas":"` + strconv.Itoa(bytecodeInfoLog.msizeMinGas) + `",`
	output += `"msize max gas":"` + strconv.Itoa(bytecodeInfoLog.msizeMaxGas) + `",`
	output += `"msizeCount":"` + strconv.Itoa(bytecodeInfoLog.msizeCount) + `",`
	output += `"gas ns":"` + strconv.FormatInt(bytecodeInfoLog.gasNanoseconds, 10) + `",`
	output += `"gas min ns":"` + strconv.FormatInt(bytecodeInfoLog.gasMinNanoseconds, 10) + `",`
	output += `"gas max ns":"` + strconv.FormatInt(bytecodeInfoLog.gasMaxNanoseconds, 10) + `",`
	output += `"gas gas":"` + strconv.Itoa(bytecodeInfoLog.gasGas) + `",`
	output += `"gas min gas":"` + strconv.Itoa(bytecodeInfoLog.gasMinGas) + `",`
	output += `"gas max gas":"` + strconv.Itoa(bytecodeInfoLog.gasMaxGas) + `",`
	output += `"gasCount":"` + strconv.Itoa(bytecodeInfoLog.gasCount) + `",`
	output += `"create ns":"` + strconv.FormatInt(bytecodeInfoLog.createNanoseconds, 10) + `",`
	output += `"create min ns":"` + strconv.FormatInt(bytecodeInfoLog.createMinNanoseconds, 10) + `",`
	output += `"create max ns":"` + strconv.FormatInt(bytecodeInfoLog.createMaxNanoseconds, 10) + `",`
	output += `"create gas":"` + strconv.Itoa(bytecodeInfoLog.createGas) + `",`
	output += `"create min gas":"` + strconv.Itoa(bytecodeInfoLog.createMinGas) + `",`
	output += `"create max gas":"` + strconv.Itoa(bytecodeInfoLog.createMaxGas) + `",`
	output += `"createCount":"` + strconv.Itoa(bytecodeInfoLog.createCount) + `",`
	output += `"create2 ns":"` + strconv.FormatInt(bytecodeInfoLog.create2Nanoseconds, 10) + `",`
	output += `"create2 min ns":"` + strconv.FormatInt(bytecodeInfoLog.create2MinNanoseconds, 10) + `",`
	output += `"create2 max ns":"` + strconv.FormatInt(bytecodeInfoLog.create2MaxNanoseconds, 10) + `",`
	output += `"create2 gas":"` + strconv.Itoa(bytecodeInfoLog.create2Gas) + `",`
	output += `"create2 min gas":"` + strconv.Itoa(bytecodeInfoLog.create2MinGas) + `",`
	output += `"create2 max gas":"` + strconv.Itoa(bytecodeInfoLog.create2MaxGas) + `",`
	output += `"create2Count":"` + strconv.Itoa(bytecodeInfoLog.create2Count) + `",`
	output += `"call ns":"` + strconv.FormatInt(bytecodeInfoLog.callNanoseconds, 10) + `",`
	output += `"call min ns":"` + strconv.FormatInt(bytecodeInfoLog.callMinNanoseconds, 10) + `",`
	output += `"call max ns":"` + strconv.FormatInt(bytecodeInfoLog.callMaxNanoseconds, 10) + `",`
	output += `"call gas":"` + strconv.Itoa(bytecodeInfoLog.callGas) + `",`
	output += `"call min gas":"` + strconv.Itoa(bytecodeInfoLog.callMinGas) + `",`
	output += `"call max gas":"` + strconv.Itoa(bytecodeInfoLog.callMaxGas) + `",`
	output += `"callCount":"` + strconv.Itoa(bytecodeInfoLog.callCount) + `",`
	output += `"callCode ns":"` + strconv.FormatInt(bytecodeInfoLog.callCodeNanoseconds, 10) + `",`
	output += `"callCode min ns":"` + strconv.FormatInt(bytecodeInfoLog.callCodeMinNanoseconds, 10) + `",`
	output += `"callCode max ns":"` + strconv.FormatInt(bytecodeInfoLog.callCodeMaxNanoseconds, 10) + `",`
	output += `"callCode gas":"` + strconv.Itoa(bytecodeInfoLog.callCodeGas) + `",`
	output += `"callCode min gas":"` + strconv.Itoa(bytecodeInfoLog.callCodeMinGas) + `",`
	output += `"callCode max gas":"` + strconv.Itoa(bytecodeInfoLog.callCodeMaxGas) + `",`
	output += `"callCodeCount":"` + strconv.Itoa(bytecodeInfoLog.callCodeCount) + `",`
	output += `"delegateCall ns":"` + strconv.FormatInt(bytecodeInfoLog.delegateCallNanoseconds, 10) + `",`
	output += `"delegateCall min ns":"` + strconv.FormatInt(bytecodeInfoLog.delegateCallMinNanoseconds, 10) + `",`
	output += `"delegateCall max ns":"` + strconv.FormatInt(bytecodeInfoLog.delegateCallMaxNanoseconds, 10) + `",`
	output += `"delegateCall gas":"` + strconv.Itoa(bytecodeInfoLog.delegateCallGas) + `",`
	output += `"delegateCall min gas":"` + strconv.Itoa(bytecodeInfoLog.delegateCallMinGas) + `",`
	output += `"delegateCall max gas":"` + strconv.Itoa(bytecodeInfoLog.delegateCallMaxGas) + `",`
	output += `"delegateCallCount":"` + strconv.Itoa(bytecodeInfoLog.delegateCallCount) + `",`
	output += `"staticCall ns":"` + strconv.FormatInt(bytecodeInfoLog.staticCallNanoseconds, 10) + `",`
	output += `"staticCall min ns":"` + strconv.FormatInt(bytecodeInfoLog.staticCallMinNanoseconds, 10) + `",`
	output += `"staticCall max ns":"` + strconv.FormatInt(bytecodeInfoLog.staticCallMaxNanoseconds, 10) + `",`
	output += `"staticCall gas":"` + strconv.Itoa(bytecodeInfoLog.staticCallGas) + `",`
	output += `"staticCall min gas":"` + strconv.Itoa(bytecodeInfoLog.staticCallMinGas) + `",`
	output += `"staticCall max gas":"` + strconv.Itoa(bytecodeInfoLog.staticCallMaxGas) + `",`
	output += `"staticCallCount":"` + strconv.Itoa(bytecodeInfoLog.staticCallCount) + `",`
	output += `"return ns":"` + strconv.FormatInt(bytecodeInfoLog.returnNanoseconds, 10) + `",`
	output += `"return min ns":"` + strconv.FormatInt(bytecodeInfoLog.returnMinNanoseconds, 10) + `",`
	output += `"return max ns":"` + strconv.FormatInt(bytecodeInfoLog.returnMaxNanoseconds, 10) + `",`
	output += `"return gas":"` + strconv.Itoa(bytecodeInfoLog.returnGas) + `",`
	output += `"return min gas":"` + strconv.Itoa(bytecodeInfoLog.returnMinGas) + `",`
	output += `"return max gas":"` + strconv.Itoa(bytecodeInfoLog.returnMaxGas) + `",`
	output += `"returnCount":"` + strconv.Itoa(bytecodeInfoLog.returnCount) + `",`
	output += `"revert ns":"` + strconv.FormatInt(bytecodeInfoLog.revertNanoseconds, 10) + `",`
	output += `"revert min ns":"` + strconv.FormatInt(bytecodeInfoLog.revertMinNanoseconds, 10) + `",`
	output += `"revert max ns":"` + strconv.FormatInt(bytecodeInfoLog.revertMaxNanoseconds, 10) + `",`
	output += `"revert gas":"` + strconv.Itoa(bytecodeInfoLog.revertGas) + `",`
	output += `"revert min gas":"` + strconv.Itoa(bytecodeInfoLog.revertMinGas) + `",`
	output += `"revert max gas":"` + strconv.Itoa(bytecodeInfoLog.revertMaxGas) + `",`
	output += `"revertCount":"` + strconv.Itoa(bytecodeInfoLog.revertCount) + `",`
	output += `"undefined ns":"` + strconv.FormatInt(bytecodeInfoLog.undefinedNanoseconds, 10) + `",`
	output += `"undefined min ns":"` + strconv.FormatInt(bytecodeInfoLog.undefinedMinNanoseconds, 10) + `",`
	output += `"undefined max ns":"` + strconv.FormatInt(bytecodeInfoLog.undefinedMaxNanoseconds, 10) + `",`
	output += `"undefined gas":"` + strconv.Itoa(bytecodeInfoLog.undefinedGas) + `",`
	output += `"undefined min gas":"` + strconv.Itoa(bytecodeInfoLog.undefinedMinGas) + `",`
	output += `"undefined max gas":"` + strconv.Itoa(bytecodeInfoLog.undefinedMaxGas) + `",`
	output += `"undefinedCount":"` + strconv.Itoa(bytecodeInfoLog.undefinedCount) + `",`
	output += `"stop ns":"` + strconv.FormatInt(bytecodeInfoLog.stopNanoseconds, 10) + `",`
	output += `"stop min ns":"` + strconv.FormatInt(bytecodeInfoLog.stopMinNanoseconds, 10) + `",`
	output += `"stop max ns":"` + strconv.FormatInt(bytecodeInfoLog.stopMaxNanoseconds, 10) + `",`
	output += `"stop gas":"` + strconv.Itoa(bytecodeInfoLog.stopGas) + `",`
	output += `"stop min gas":"` + strconv.Itoa(bytecodeInfoLog.stopMinGas) + `",`
	output += `"stop max gas":"` + strconv.Itoa(bytecodeInfoLog.stopMaxGas) + `",`
	output += `"stopCount":"` + strconv.Itoa(bytecodeInfoLog.stopCount) + `",`
	output += `"selfdestruct ns":"` + strconv.FormatInt(bytecodeInfoLog.selfdestructNanoseconds, 10) + `",`
	output += `"selfdestruct min ns":"` + strconv.FormatInt(bytecodeInfoLog.selfdestructMinNanoseconds, 10) + `",`
	output += `"selfdestruct max ns":"` + strconv.FormatInt(bytecodeInfoLog.selfdestructMaxNanoseconds, 10) + `",`
	output += `"selfdestruct gas":"` + strconv.Itoa(bytecodeInfoLog.selfdestructGas) + `",`
	output += `"selfdestruct min gas":"` + strconv.Itoa(bytecodeInfoLog.selfdestructMinGas) + `",`
	output += `"selfdestruct max gas":"` + strconv.Itoa(bytecodeInfoLog.selfdestructMaxGas) + `",`
	output += `"selfdestructCount":"` + strconv.Itoa(bytecodeInfoLog.selfdestructCount) + `",`
	output += `"makeLog ns":"` + strconv.FormatInt(bytecodeInfoLog.makeLogNanoseconds, 10) + `",`
	output += `"makeLog min ns":"` + strconv.FormatInt(bytecodeInfoLog.makeLogMinNanoseconds, 10) + `",`
	output += `"makeLog max ns":"` + strconv.FormatInt(bytecodeInfoLog.makeLogMaxNanoseconds, 10) + `",`
	output += `"makeLog gas":"` + strconv.Itoa(bytecodeInfoLog.makeLogGas) + `",`
	output += `"makeLog min gas":"` + strconv.Itoa(bytecodeInfoLog.makeLogMinGas) + `",`
	output += `"makeLog max gas":"` + strconv.Itoa(bytecodeInfoLog.makeLogMaxGas) + `",`
	output += `"makeLogCount":"` + strconv.Itoa(bytecodeInfoLog.makeLogCount) + `",`
	output += `"push1 ns":"` + strconv.FormatInt(bytecodeInfoLog.push1Nanoseconds, 10) + `",`
	output += `"push1 min ns":"` + strconv.FormatInt(bytecodeInfoLog.push1MinNanoseconds, 10) + `",`
	output += `"push1 max ns":"` + strconv.FormatInt(bytecodeInfoLog.push1MaxNanoseconds, 10) + `",`
	output += `"push1 gas":"` + strconv.Itoa(bytecodeInfoLog.push1Gas) + `",`
	output += `"push1 min gas":"` + strconv.Itoa(bytecodeInfoLog.push1MinGas) + `",`
	output += `"push1 max gas":"` + strconv.Itoa(bytecodeInfoLog.push1MaxGas) + `",`
	output += `"push1Count":"` + strconv.Itoa(bytecodeInfoLog.push1Count) + `",`
	output += `"makePush ns":"` + strconv.FormatInt(bytecodeInfoLog.makePushNanoseconds, 10) + `",`
	output += `"makePush min ns":"` + strconv.FormatInt(bytecodeInfoLog.makePushMinNanoseconds, 10) + `",`
	output += `"makePush max ns":"` + strconv.FormatInt(bytecodeInfoLog.makePushMaxNanoseconds, 10) + `",`
	output += `"makePush gas":"` + strconv.Itoa(bytecodeInfoLog.makePushGas) + `",`
	output += `"makePush min gas":"` + strconv.Itoa(bytecodeInfoLog.makePushMinGas) + `",`
	output += `"makePush max gas":"` + strconv.Itoa(bytecodeInfoLog.makePushMaxGas) + `",`
	output += `"makePushCount":"` + strconv.Itoa(bytecodeInfoLog.makePushCount) + `",`
	output += `"makeDup ns":"` + strconv.FormatInt(bytecodeInfoLog.makeDupNanoseconds, 10) + `",`
	output += `"makeDup min ns":"` + strconv.FormatInt(bytecodeInfoLog.makeDupMinNanoseconds, 10) + `",`
	output += `"makeDup max ns":"` + strconv.FormatInt(bytecodeInfoLog.makeDupMaxNanoseconds, 10) + `",`
	output += `"makeDup gas":"` + strconv.Itoa(bytecodeInfoLog.makeDupGas) + `",`
	output += `"makeDup min gas":"` + strconv.Itoa(bytecodeInfoLog.makeDupMinGas) + `",`
	output += `"makeDup max gas":"` + strconv.Itoa(bytecodeInfoLog.makeDupMaxGas) + `",`
	output += `"makeDupCount":"` + strconv.Itoa(bytecodeInfoLog.makeDupCount) + `",`
	output += `"makeSwap ns":"` + strconv.FormatInt(bytecodeInfoLog.makeSwapNanoseconds, 10) + `",`
	output += `"makeSwap min ns":"` + strconv.FormatInt(bytecodeInfoLog.makeSwapMinNanoseconds, 10) + `",`
	output += `"makeSwap max ns":"` + strconv.FormatInt(bytecodeInfoLog.makeSwapMaxNanoseconds, 10) + `",`
	output += `"makeSwap gas":"` + strconv.Itoa(bytecodeInfoLog.makeSwapGas) + `",`
	output += `"makeSwap min gas":"` + strconv.Itoa(bytecodeInfoLog.makeSwapMinGas) + `",`
	output += `"makeSwap max gas":"` + strconv.Itoa(bytecodeInfoLog.makeSwapMaxGas) + `",`
	output += `"makeSwapCount":"` + strconv.Itoa(bytecodeInfoLog.makeSwapCount) + `"`
	output += "}"

	return output
}

// Cybersecurity Lab: Measure duration
func opAdd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var a []byte
	var b error
	start := time.Now()
	a, b = _opAdd(pc, interpreter, scope)
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.addMaxNanoseconds {
		bytecodeInfoLog.addMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.addMinNanoseconds || bytecodeInfoLog.addMinNanoseconds == 0 {
		bytecodeInfoLog.addMinNanoseconds = duration
	}
	bytecodeInfoLog.addNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.subMaxNanoseconds {
		bytecodeInfoLog.subMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.subMinNanoseconds || bytecodeInfoLog.subMinNanoseconds == 0 {
		bytecodeInfoLog.subMinNanoseconds = duration
	}
	bytecodeInfoLog.subNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.mulMaxNanoseconds {
		bytecodeInfoLog.mulMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.mulMinNanoseconds || bytecodeInfoLog.mulMinNanoseconds == 0 {
		bytecodeInfoLog.mulMinNanoseconds = duration
	}
	bytecodeInfoLog.mulNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.divMaxNanoseconds {
		bytecodeInfoLog.divMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.divMinNanoseconds || bytecodeInfoLog.divMinNanoseconds == 0 {
		bytecodeInfoLog.divMinNanoseconds = duration
	}
	bytecodeInfoLog.divNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.sdivMaxNanoseconds {
		bytecodeInfoLog.sdivMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.sdivMinNanoseconds || bytecodeInfoLog.sdivMinNanoseconds == 0 {
		bytecodeInfoLog.sdivMinNanoseconds = duration
	}
	bytecodeInfoLog.sdivNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.modMaxNanoseconds {
		bytecodeInfoLog.modMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.modMinNanoseconds || bytecodeInfoLog.modMinNanoseconds == 0 {
		bytecodeInfoLog.modMinNanoseconds = duration
	}
	bytecodeInfoLog.modNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.smodMaxNanoseconds {
		bytecodeInfoLog.smodMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.smodMinNanoseconds || bytecodeInfoLog.smodMinNanoseconds == 0 {
		bytecodeInfoLog.smodMinNanoseconds = duration
	}
	bytecodeInfoLog.smodNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.expMaxNanoseconds {
		bytecodeInfoLog.expMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.expMinNanoseconds || bytecodeInfoLog.expMinNanoseconds == 0 {
		bytecodeInfoLog.expMinNanoseconds = duration
	}
	bytecodeInfoLog.expNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.signExtendMaxNanoseconds {
		bytecodeInfoLog.signExtendMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.signExtendMinNanoseconds || bytecodeInfoLog.signExtendMinNanoseconds == 0 {
		bytecodeInfoLog.signExtendMinNanoseconds = duration
	}
	bytecodeInfoLog.signExtendNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.notMaxNanoseconds {
		bytecodeInfoLog.notMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.notMinNanoseconds || bytecodeInfoLog.notMinNanoseconds == 0 {
		bytecodeInfoLog.notMinNanoseconds = duration
	}
	bytecodeInfoLog.notNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.ltMaxNanoseconds {
		bytecodeInfoLog.ltMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.ltMinNanoseconds || bytecodeInfoLog.ltMinNanoseconds == 0 {
		bytecodeInfoLog.ltMinNanoseconds = duration
	}
	bytecodeInfoLog.ltNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.gtMaxNanoseconds {
		bytecodeInfoLog.gtMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.gtMinNanoseconds || bytecodeInfoLog.gtMinNanoseconds == 0 {
		bytecodeInfoLog.gtMinNanoseconds = duration
	}
	bytecodeInfoLog.gtNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.sltMaxNanoseconds {
		bytecodeInfoLog.sltMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.sltMinNanoseconds || bytecodeInfoLog.sltMinNanoseconds == 0 {
		bytecodeInfoLog.sltMinNanoseconds = duration
	}
	bytecodeInfoLog.sltNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.sgtMaxNanoseconds {
		bytecodeInfoLog.sgtMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.sgtMinNanoseconds || bytecodeInfoLog.sgtMinNanoseconds == 0 {
		bytecodeInfoLog.sgtMinNanoseconds = duration
	}
	bytecodeInfoLog.sgtNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.eqMaxNanoseconds {
		bytecodeInfoLog.eqMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.eqMinNanoseconds || bytecodeInfoLog.eqMinNanoseconds == 0 {
		bytecodeInfoLog.eqMinNanoseconds = duration
	}
	bytecodeInfoLog.eqNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.isZeroMaxNanoseconds {
		bytecodeInfoLog.isZeroMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.isZeroMinNanoseconds || bytecodeInfoLog.isZeroMinNanoseconds == 0 {
		bytecodeInfoLog.isZeroMinNanoseconds = duration
	}
	bytecodeInfoLog.isZeroNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.andMaxNanoseconds {
		bytecodeInfoLog.andMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.andMinNanoseconds || bytecodeInfoLog.andMinNanoseconds == 0 {
		bytecodeInfoLog.andMinNanoseconds = duration
	}
	bytecodeInfoLog.andNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.orMaxNanoseconds {
		bytecodeInfoLog.orMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.orMinNanoseconds || bytecodeInfoLog.orMinNanoseconds == 0 {
		bytecodeInfoLog.orMinNanoseconds = duration
	}
	bytecodeInfoLog.orNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.xorMaxNanoseconds {
		bytecodeInfoLog.xorMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.xorMinNanoseconds || bytecodeInfoLog.xorMinNanoseconds == 0 {
		bytecodeInfoLog.xorMinNanoseconds = duration
	}
	bytecodeInfoLog.xorNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.byteMaxNanoseconds {
		bytecodeInfoLog.byteMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.byteMinNanoseconds || bytecodeInfoLog.byteMinNanoseconds == 0 {
		bytecodeInfoLog.byteMinNanoseconds = duration
	}
	bytecodeInfoLog.byteNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.addmodMaxNanoseconds {
		bytecodeInfoLog.addmodMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.addmodMinNanoseconds || bytecodeInfoLog.addmodMinNanoseconds == 0 {
		bytecodeInfoLog.addmodMinNanoseconds = duration
	}
	bytecodeInfoLog.addmodNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.mulmodMaxNanoseconds {
		bytecodeInfoLog.mulmodMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.mulmodMinNanoseconds || bytecodeInfoLog.mulmodMinNanoseconds == 0 {
		bytecodeInfoLog.mulmodMinNanoseconds = duration
	}
	bytecodeInfoLog.mulmodNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.shlMaxNanoseconds {
		bytecodeInfoLog.shlMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.shlMinNanoseconds || bytecodeInfoLog.shlMinNanoseconds == 0 {
		bytecodeInfoLog.shlMinNanoseconds = duration
	}
	bytecodeInfoLog.shlNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.shrMaxNanoseconds {
		bytecodeInfoLog.shrMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.shrMinNanoseconds || bytecodeInfoLog.shrMinNanoseconds == 0 {
		bytecodeInfoLog.shrMinNanoseconds = duration
	}
	bytecodeInfoLog.shrNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.sarMaxNanoseconds {
		bytecodeInfoLog.sarMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.sarMinNanoseconds || bytecodeInfoLog.sarMinNanoseconds == 0 {
		bytecodeInfoLog.sarMinNanoseconds = duration
	}
	bytecodeInfoLog.sarNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.keccak256MaxNanoseconds {
		bytecodeInfoLog.keccak256MaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.keccak256MinNanoseconds || bytecodeInfoLog.keccak256MinNanoseconds == 0 {
		bytecodeInfoLog.keccak256MinNanoseconds = duration
	}
	bytecodeInfoLog.keccak256Nanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.addressMaxNanoseconds {
		bytecodeInfoLog.addressMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.addressMinNanoseconds || bytecodeInfoLog.addressMinNanoseconds == 0 {
		bytecodeInfoLog.addressMinNanoseconds = duration
	}
	bytecodeInfoLog.addressNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.balanceMaxNanoseconds {
		bytecodeInfoLog.balanceMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.balanceMinNanoseconds || bytecodeInfoLog.balanceMinNanoseconds == 0 {
		bytecodeInfoLog.balanceMinNanoseconds = duration
	}
	bytecodeInfoLog.balanceNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.originMaxNanoseconds {
		bytecodeInfoLog.originMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.originMinNanoseconds || bytecodeInfoLog.originMinNanoseconds == 0 {
		bytecodeInfoLog.originMinNanoseconds = duration
	}
	bytecodeInfoLog.originNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callerMaxNanoseconds {
		bytecodeInfoLog.callerMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callerMinNanoseconds || bytecodeInfoLog.callerMinNanoseconds == 0 {
		bytecodeInfoLog.callerMinNanoseconds = duration
	}
	bytecodeInfoLog.callerNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callValueMaxNanoseconds {
		bytecodeInfoLog.callValueMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callValueMinNanoseconds || bytecodeInfoLog.callValueMinNanoseconds == 0 {
		bytecodeInfoLog.callValueMinNanoseconds = duration
	}
	bytecodeInfoLog.callValueNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callDataLoadMaxNanoseconds {
		bytecodeInfoLog.callDataLoadMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callDataLoadMinNanoseconds || bytecodeInfoLog.callDataLoadMinNanoseconds == 0 {
		bytecodeInfoLog.callDataLoadMinNanoseconds = duration
	}
	bytecodeInfoLog.callDataLoadNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callDataSizeMaxNanoseconds {
		bytecodeInfoLog.callDataSizeMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callDataSizeMinNanoseconds || bytecodeInfoLog.callDataSizeMinNanoseconds == 0 {
		bytecodeInfoLog.callDataSizeMinNanoseconds = duration
	}
	bytecodeInfoLog.callDataSizeNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callDataCopyMaxNanoseconds {
		bytecodeInfoLog.callDataCopyMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callDataCopyMinNanoseconds || bytecodeInfoLog.callDataCopyMinNanoseconds == 0 {
		bytecodeInfoLog.callDataCopyMinNanoseconds = duration
	}
	bytecodeInfoLog.callDataCopyNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.returnDataSizeMaxNanoseconds {
		bytecodeInfoLog.returnDataSizeMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.returnDataSizeMinNanoseconds || bytecodeInfoLog.returnDataSizeMinNanoseconds == 0 {
		bytecodeInfoLog.returnDataSizeMinNanoseconds = duration
	}
	bytecodeInfoLog.returnDataSizeNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.returnDataCopyMaxNanoseconds {
		bytecodeInfoLog.returnDataCopyMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.returnDataCopyMinNanoseconds || bytecodeInfoLog.returnDataCopyMinNanoseconds == 0 {
		bytecodeInfoLog.returnDataCopyMinNanoseconds = duration
	}
	bytecodeInfoLog.returnDataCopyNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.extCodeSizeMaxNanoseconds {
		bytecodeInfoLog.extCodeSizeMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.extCodeSizeMinNanoseconds || bytecodeInfoLog.extCodeSizeMinNanoseconds == 0 {
		bytecodeInfoLog.extCodeSizeMinNanoseconds = duration
	}
	bytecodeInfoLog.extCodeSizeNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.codeSizeMaxNanoseconds {
		bytecodeInfoLog.codeSizeMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.codeSizeMinNanoseconds || bytecodeInfoLog.codeSizeMinNanoseconds == 0 {
		bytecodeInfoLog.codeSizeMinNanoseconds = duration
	}
	bytecodeInfoLog.codeSizeNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.codeCopyMaxNanoseconds {
		bytecodeInfoLog.codeCopyMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.codeCopyMinNanoseconds || bytecodeInfoLog.codeCopyMinNanoseconds == 0 {
		bytecodeInfoLog.codeCopyMinNanoseconds = duration
	}
	bytecodeInfoLog.codeCopyNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.extCodeCopyMaxNanoseconds {
		bytecodeInfoLog.extCodeCopyMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.extCodeCopyMinNanoseconds || bytecodeInfoLog.extCodeCopyMinNanoseconds == 0 {
		bytecodeInfoLog.extCodeCopyMinNanoseconds = duration
	}
	bytecodeInfoLog.extCodeCopyNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.extCodeHashMaxNanoseconds {
		bytecodeInfoLog.extCodeHashMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.extCodeHashMinNanoseconds || bytecodeInfoLog.extCodeHashMinNanoseconds == 0 {
		bytecodeInfoLog.extCodeHashMinNanoseconds = duration
	}
	bytecodeInfoLog.extCodeHashNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.gaspriceMaxNanoseconds {
		bytecodeInfoLog.gaspriceMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.gaspriceMinNanoseconds || bytecodeInfoLog.gaspriceMinNanoseconds == 0 {
		bytecodeInfoLog.gaspriceMinNanoseconds = duration
	}
	bytecodeInfoLog.gaspriceNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.blockhashMaxNanoseconds {
		bytecodeInfoLog.blockhashMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.blockhashMinNanoseconds || bytecodeInfoLog.blockhashMinNanoseconds == 0 {
		bytecodeInfoLog.blockhashMinNanoseconds = duration
	}
	bytecodeInfoLog.blockhashNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.coinbaseMaxNanoseconds {
		bytecodeInfoLog.coinbaseMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.coinbaseMinNanoseconds || bytecodeInfoLog.coinbaseMinNanoseconds == 0 {
		bytecodeInfoLog.coinbaseMinNanoseconds = duration
	}
	bytecodeInfoLog.coinbaseNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.timestampMaxNanoseconds {
		bytecodeInfoLog.timestampMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.timestampMinNanoseconds || bytecodeInfoLog.timestampMinNanoseconds == 0 {
		bytecodeInfoLog.timestampMinNanoseconds = duration
	}
	bytecodeInfoLog.timestampNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.numberMaxNanoseconds {
		bytecodeInfoLog.numberMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.numberMinNanoseconds || bytecodeInfoLog.numberMinNanoseconds == 0 {
		bytecodeInfoLog.numberMinNanoseconds = duration
	}
	bytecodeInfoLog.numberNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.difficultyMaxNanoseconds {
		bytecodeInfoLog.difficultyMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.difficultyMinNanoseconds || bytecodeInfoLog.difficultyMinNanoseconds == 0 {
		bytecodeInfoLog.difficultyMinNanoseconds = duration
	}
	bytecodeInfoLog.difficultyNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.randomMaxNanoseconds {
		bytecodeInfoLog.randomMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.randomMinNanoseconds || bytecodeInfoLog.randomMinNanoseconds == 0 {
		bytecodeInfoLog.randomMinNanoseconds = duration
	}
	bytecodeInfoLog.randomNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.gasLimitMaxNanoseconds {
		bytecodeInfoLog.gasLimitMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.gasLimitMinNanoseconds || bytecodeInfoLog.gasLimitMinNanoseconds == 0 {
		bytecodeInfoLog.gasLimitMinNanoseconds = duration
	}
	bytecodeInfoLog.gasLimitNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.popMaxNanoseconds {
		bytecodeInfoLog.popMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.popMinNanoseconds || bytecodeInfoLog.popMinNanoseconds == 0 {
		bytecodeInfoLog.popMinNanoseconds = duration
	}
	bytecodeInfoLog.popNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.mloadMaxNanoseconds {
		bytecodeInfoLog.mloadMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.mloadMinNanoseconds || bytecodeInfoLog.mloadMinNanoseconds == 0 {
		bytecodeInfoLog.mloadMinNanoseconds = duration
	}
	bytecodeInfoLog.mloadNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.mstoreMaxNanoseconds {
		bytecodeInfoLog.mstoreMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.mstoreMinNanoseconds || bytecodeInfoLog.mstoreMinNanoseconds == 0 {
		bytecodeInfoLog.mstoreMinNanoseconds = duration
	}
	bytecodeInfoLog.mstoreNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.mstore8MaxNanoseconds {
		bytecodeInfoLog.mstore8MaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.mstore8MinNanoseconds || bytecodeInfoLog.mstore8MinNanoseconds == 0 {
		bytecodeInfoLog.mstore8MinNanoseconds = duration
	}
	bytecodeInfoLog.mstore8Nanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.sloadMaxNanoseconds {
		bytecodeInfoLog.sloadMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.sloadMinNanoseconds || bytecodeInfoLog.sloadMinNanoseconds == 0 {
		bytecodeInfoLog.sloadMinNanoseconds = duration
	}
	bytecodeInfoLog.sloadNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.sstoreMaxNanoseconds {
		bytecodeInfoLog.sstoreMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.sstoreMinNanoseconds || bytecodeInfoLog.sstoreMinNanoseconds == 0 {
		bytecodeInfoLog.sstoreMinNanoseconds = duration
	}
	bytecodeInfoLog.sstoreNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.jumpMaxNanoseconds {
		bytecodeInfoLog.jumpMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.jumpMinNanoseconds || bytecodeInfoLog.jumpMinNanoseconds == 0 {
		bytecodeInfoLog.jumpMinNanoseconds = duration
	}
	bytecodeInfoLog.jumpNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.jumpiMaxNanoseconds {
		bytecodeInfoLog.jumpiMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.jumpiMinNanoseconds || bytecodeInfoLog.jumpiMinNanoseconds == 0 {
		bytecodeInfoLog.jumpiMinNanoseconds = duration
	}
	bytecodeInfoLog.jumpiNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.jumpdestMaxNanoseconds {
		bytecodeInfoLog.jumpdestMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.jumpdestMinNanoseconds || bytecodeInfoLog.jumpdestMinNanoseconds == 0 {
		bytecodeInfoLog.jumpdestMinNanoseconds = duration
	}
	bytecodeInfoLog.jumpdestNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.pcMaxNanoseconds {
		bytecodeInfoLog.pcMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.pcMinNanoseconds || bytecodeInfoLog.pcMinNanoseconds == 0 {
		bytecodeInfoLog.pcMinNanoseconds = duration
	}
	bytecodeInfoLog.pcNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.msizeMaxNanoseconds {
		bytecodeInfoLog.msizeMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.msizeMinNanoseconds || bytecodeInfoLog.msizeMinNanoseconds == 0 {
		bytecodeInfoLog.msizeMinNanoseconds = duration
	}
	bytecodeInfoLog.msizeNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.gasMaxNanoseconds {
		bytecodeInfoLog.gasMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.gasMinNanoseconds || bytecodeInfoLog.gasMinNanoseconds == 0 {
		bytecodeInfoLog.gasMinNanoseconds = duration
	}
	bytecodeInfoLog.gasNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.createMaxNanoseconds {
		bytecodeInfoLog.createMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.createMinNanoseconds || bytecodeInfoLog.createMinNanoseconds == 0 {
		bytecodeInfoLog.createMinNanoseconds = duration
	}
	bytecodeInfoLog.createNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.create2MaxNanoseconds {
		bytecodeInfoLog.create2MaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.create2MinNanoseconds || bytecodeInfoLog.create2MinNanoseconds == 0 {
		bytecodeInfoLog.create2MinNanoseconds = duration
	}
	bytecodeInfoLog.create2Nanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callMaxNanoseconds {
		bytecodeInfoLog.callMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callMinNanoseconds || bytecodeInfoLog.callMinNanoseconds == 0 {
		bytecodeInfoLog.callMinNanoseconds = duration
	}
	bytecodeInfoLog.callNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.callCodeMaxNanoseconds {
		bytecodeInfoLog.callCodeMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.callCodeMinNanoseconds || bytecodeInfoLog.callCodeMinNanoseconds == 0 {
		bytecodeInfoLog.callCodeMinNanoseconds = duration
	}
	bytecodeInfoLog.callCodeNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.delegateCallMaxNanoseconds {
		bytecodeInfoLog.delegateCallMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.delegateCallMinNanoseconds || bytecodeInfoLog.delegateCallMinNanoseconds == 0 {
		bytecodeInfoLog.delegateCallMinNanoseconds = duration
	}
	bytecodeInfoLog.delegateCallNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.staticCallMaxNanoseconds {
		bytecodeInfoLog.staticCallMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.staticCallMinNanoseconds || bytecodeInfoLog.staticCallMinNanoseconds == 0 {
		bytecodeInfoLog.staticCallMinNanoseconds = duration
	}
	bytecodeInfoLog.staticCallNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.returnMaxNanoseconds {
		bytecodeInfoLog.returnMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.returnMinNanoseconds || bytecodeInfoLog.returnMinNanoseconds == 0 {
		bytecodeInfoLog.returnMinNanoseconds = duration
	}
	bytecodeInfoLog.returnNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.revertMaxNanoseconds {
		bytecodeInfoLog.revertMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.revertMinNanoseconds || bytecodeInfoLog.revertMinNanoseconds == 0 {
		bytecodeInfoLog.revertMinNanoseconds = duration
	}
	bytecodeInfoLog.revertNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.undefinedMaxNanoseconds {
		bytecodeInfoLog.undefinedMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.undefinedMinNanoseconds || bytecodeInfoLog.undefinedMinNanoseconds == 0 {
		bytecodeInfoLog.undefinedMinNanoseconds = duration
	}
	bytecodeInfoLog.undefinedNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.stopMaxNanoseconds {
		bytecodeInfoLog.stopMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.stopMinNanoseconds || bytecodeInfoLog.stopMinNanoseconds == 0 {
		bytecodeInfoLog.stopMinNanoseconds = duration
	}
	bytecodeInfoLog.stopNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.selfdestructMaxNanoseconds {
		bytecodeInfoLog.selfdestructMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.selfdestructMinNanoseconds || bytecodeInfoLog.selfdestructMinNanoseconds == 0 {
		bytecodeInfoLog.selfdestructMinNanoseconds = duration
	}
	bytecodeInfoLog.selfdestructNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.makeLogMaxNanoseconds {
		bytecodeInfoLog.makeLogMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.makeLogMinNanoseconds || bytecodeInfoLog.makeLogMinNanoseconds == 0 {
		bytecodeInfoLog.makeLogMinNanoseconds = duration
	}
	bytecodeInfoLog.makeLogNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.push1MaxNanoseconds {
		bytecodeInfoLog.push1MaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.push1MinNanoseconds || bytecodeInfoLog.push1MinNanoseconds == 0 {
		bytecodeInfoLog.push1MinNanoseconds = duration
	}
	bytecodeInfoLog.push1Nanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.makePushMaxNanoseconds {
		bytecodeInfoLog.makePushMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.makePushMinNanoseconds || bytecodeInfoLog.makePushMinNanoseconds == 0 {
		bytecodeInfoLog.makePushMinNanoseconds = duration
	}
	bytecodeInfoLog.makePushNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.makeDupMaxNanoseconds {
		bytecodeInfoLog.makeDupMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.makeDupMinNanoseconds || bytecodeInfoLog.makeDupMinNanoseconds == 0 {
		bytecodeInfoLog.makeDupMinNanoseconds = duration
	}
	bytecodeInfoLog.makeDupNanoseconds += duration
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
	duration := int64(time.Since(start) / time.Nanosecond)
	if duration > bytecodeInfoLog.makeSwapMaxNanoseconds {
		bytecodeInfoLog.makeSwapMaxNanoseconds = duration
	}
	if duration < bytecodeInfoLog.makeSwapMinNanoseconds || bytecodeInfoLog.makeSwapMinNanoseconds == 0 {
		bytecodeInfoLog.makeSwapMinNanoseconds = duration
	}
	bytecodeInfoLog.makeSwapNanoseconds += duration
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
