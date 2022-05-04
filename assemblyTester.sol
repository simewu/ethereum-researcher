// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.13;

contract assemblyTester {
    /*
        1:		STOP
        2:		ADD
        3:		MUL
        4:		SUB
        5:		DIV
        6:		SDIV
        7:		MOD
        8:		SMOD
        9:		EXP
        10:		NOT
        11:		LT
        12:		GT
        13:		SLT
        14:		SGT
        15:		EQ
        16:		ISZERO
        17:		SIGNEXTEND
        18:		AND
        19:		OR
        20:		XOR
        21:		BYTE
        22:		SHL
        23:		SHR
        24:		SAR
        25:		ADDMOD
        26:		MULMOD
        27:		KECCAK256
        28:		ADDRESS
        29:		BALANCE
        30:		ORIGIN
        31:		CALLER
        32:		CALLVALUE
        33:		CALLDATALOAD
        34:		CALLDATASIZE
        35:		CALLDATACOPY
        36:		CHAINID
        37:		BASEFEE
        38:		DELEGATECALL
        39:		STATICCALL
        40:		CODESIZE
        41:		CODECOPY
        42:		GASPRICE
        43:		EXTCODESIZE
        44:		EXTCODECOPY
        45:		RETURNDATASIZE
        46:		RETURNDATACOPY
        47:		EXTCODEHASH
        48:		BLOCKHASH
        49:		COINBASE
        50:		TIMESTAMP
        51:		NUMBER
        52:		DIFFICULTY
        53:		GASLIMIT
        54:		SELFBALANCE
        55:		POP
        56:		MLOAD
        57:		MSTORE
        58:		MSTORE8
        59:		SLOAD
        60:		SSTORE
        61:		MSIZE
        62:		GAS
        63:		JUMPDEST
        64:		LOG0
        65:		LOG1
        66:		LOG2
        67:		LOG3
        68:		LOG4
        69:		CREATE
        70:		CREATE2
        71:		CALL
        72:		RETURN
        73:		CALLCODE
        74:		REVERT
        75:		INVALID
        76:		SELFDESTRUCT
    */
    
    function run(uint32 cmd, uint32 samples) public {
        if(cmd == 1) {
            // Opcode: STOP
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    stop()
                }
            }
        } else if(cmd == 2) {
            // Opcode: ADD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := add(5, 0x42)
                }
            }
        } else if(cmd == 3) {
            // Opcode: MUL
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := mul(5, 0x42)
                }
            }
        } else if(cmd == 4) {
            // Opcode: SUB
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := sub(5, 0x42)
                }
            }
        } else if(cmd == 5) {
            // Opcode: DIV
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := div(5, 0x42)
                }
            }
        } else if(cmd == 6) {
            // Opcode: SDIV
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := sdiv(5, 0x42)
                }
            }
        } else if(cmd == 7) {
            // Opcode: MOD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := mod(0x42, 5)
                }
            }
        } else if(cmd == 8) {
            // Opcode: SMOD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := smod(0x42, 5)
                }
            }
        } else if(cmd == 9) {
            // Opcode: EXP
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := exp(5, 0x42)
                }
            }
        } else if(cmd == 10) {
            // Opcode: NOT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := not(0x42)
                }
            }
        } else if(cmd == 11) {
            // Opcode: LT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := lt(5, 0x42)
                }
            }
        } else if(cmd == 12) {
            // Opcode: GT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := gt(5, 0x42)
                }
            }
        } else if(cmd == 13) {
            // Opcode: SLT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := slt(5, 0x42)
                }
            }
        } else if(cmd == 14) {
            // Opcode: SGT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := sgt(5, 0x42)
                }
            }
        } else if(cmd == 15) {
            // Opcode: EQ
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := eq(0x42, 0x42)
                }
            }
        } else if(cmd == 16) {
            // Opcode: ISZERO
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := iszero(0x42)
                }
            }
        } else if(cmd == 17) {
            // Opcode: SIGNEXTEND
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := signextend(5, 0x42)
                }
            }
        } else if(cmd == 18) {
            // Opcode: AND
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := and(5, 0x42)
                }
            }
        } else if(cmd == 19) {
            // Opcode: OR
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := or(5, 0x42)
                }
            }
        } else if(cmd == 20) {
            // Opcode: XOR
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := xor(5, 0x42)
                }
            }
        } else if(cmd == 21) {
            // Opcode: BYTE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := byte(5, 0x42)
                }
            }
        } else if(cmd == 22) {
            // Opcode: SHL
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := shl(5, 0x42)
                }
            }
        } else if(cmd == 23) {
            // Opcode: SHR
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := shr(5, 0x42)
                }
            }
        } else if(cmd == 24) {
            // Opcode: SAR
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := sar(5, 0x42)
                }
            }
        } else if(cmd == 25) {
            // Opcode: ADDMOD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := addmod(5, 0x42, 7)
                }
            }
        } else if(cmd == 26) {
            // Opcode: MULMOD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := mulmod(5, 0x42, 7)
                }
            }
        } else if(cmd == 27) {
            // Opcode: KECCAK256
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := keccak256(0, 0)
                }
            }
        } else if(cmd == 28) {
            // Opcode: ADDRESS
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := address()
                }
            }
        } else if(cmd == 29) {
            // Opcode: BALANCE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := balance(0x42)
                }
            }
        } else if(cmd == 30) {
            // Opcode: ORIGIN
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := origin()
                }
            }
        } else if(cmd == 31) {
            // Opcode: CALLER
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := caller()
                }
            }
        } else if(cmd == 32) {
            // Opcode: CALLVALUE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := callvalue()
                }
            }
        } else if(cmd == 33) {
            // Opcode: CALLDATALOAD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := calldataload(0x42)
                }
            }
        } else if(cmd == 34) {
            // Opcode: CALLDATASIZE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := calldatasize()
                }
            }
        } else if(cmd == 35) {
            // Opcode: CALLDATACOPY
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    calldatacopy(5, 0x42, 7)
                }
            }
        } else if(cmd == 36) {
            // Opcode: CHAINID
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := chainid()
                }
            }
        } else if(cmd == 37) {
            // Opcode: BASEFEE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := basefee()
                }
            }
        } else if(cmd == 38) {
            // Opcode: DELEGATECALL
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := delegatecall(5000, 0x42, 0, 0x44, 0, 0x20)
                }
            }
        } else if(cmd == 39) {
            // Opcode: STATICCALL
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := staticcall(5000, 0x42, 0, 0x44, 0, 0x20)
                }
            }
        } else if(cmd == 40) {
            // Opcode: CODESIZE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := codesize()
                }
            }
        } else if(cmd == 41) {
            // Opcode: CODECOPY
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    codecopy(5, 0x42, 7)
                }
            }
        } else if(cmd == 42) {
            // Opcode: GASPRICE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := gasprice()
                }
            }
        } else if(cmd == 43) {
            // Opcode: EXTCODESIZE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := extcodesize(0x42)
                }
            }
        } else if(cmd == 44) {
            // Opcode: EXTCODECOPY
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    extcodecopy(0x42, 4, 0, 0x20)
                }
            }
        } else if(cmd == 45) {
            // Opcode: RETURNDATASIZE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := returndatasize()
                }
            }
        } else if(cmd == 46) {
            // Opcode: RETURNDATACOPY
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    returndatacopy(5, 0x42, 7)
                }
            }
        } else if(cmd == 47) {
            // Opcode: EXTCODEHASH
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := extcodehash(0x42)
                }
            }
        } else if(cmd == 48) {
            // Opcode: BLOCKHASH
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := blockhash(0x42)
                }
            }
        } else if(cmd == 49) {
            // Opcode: COINBASE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := coinbase()
                }
            }
        } else if(cmd == 50) {
            // Opcode: TIMESTAMP
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := timestamp()
                }
            }
        } else if(cmd == 51) {
            // Opcode: NUMBER
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := number()
                }
            }
        } else if(cmd == 52) {
            // Opcode: DIFFICULTY
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := difficulty()
                }
            }
        } else if(cmd == 53) {
            // Opcode: GASLIMIT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := gaslimit()
                }
            }
        } else if(cmd == 54) {
            // Opcode: SELFBALANCE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := selfbalance()
                }
            }
        } else if(cmd == 55) {
            // Opcode: POP
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    pop(0x42)
                }
            }
        } else if(cmd == 56) {
            // Opcode: MLOAD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := mload(0x42)
                }
            }
        } else if(cmd == 57) {
            // Opcode: MSTORE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    mstore(5, 0x42)
                }
            }
        } else if(cmd == 58) {
            // Opcode: MSTORE8
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    mstore8(5, 0x42)
                }
            }
        } else if(cmd == 59) {
            // Opcode: SLOAD
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := sload(0x42)
                }
            }
        } else if(cmd == 60) {
            // Opcode: SSTORE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    sstore(5, 0x42)
                }
            }
        } else if(cmd == 61) {
            // Opcode: MSIZE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := msize()
                }
            }
        } else if(cmd == 62) {
            // Opcode: GAS
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := gas()
                }
            }
        } else if(cmd == 63) {
            // Opcode: JUMPDEST
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    //let x := jumpdest()
                }
            }
        } else if(cmd == 64) {
            // Opcode: LOG0
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    log0(0x42, 0x42)
                }
            }
        } else if(cmd == 65) {
            // Opcode: LOG1
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    log1(0x42, 0x42, 0x42)
                }
            }
        } else if(cmd == 66) {
            // Opcode: LOG2
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    log2(0x42, 0x42, 0x42, 0x42)
                }
            }
        } else if(cmd == 67) {
            // Opcode: LOG3
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    log3(0x42, 0x42, 0x42, 0x42, 0x42)
                }
            }
        } else if(cmd == 68) {
            // Opcode: LOG4
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    log4(0x42, 0x42, 0x42, 0x42, 0x42, 0x42)
                }
            }
        } else if(cmd == 69) {
            // Opcode: CREATE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := create(5, 0x42, 7)
                }
            }
        } else if(cmd == 70) {
            // Opcode: CREATE2
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := create2(5, 0x42, 7, 64)
                }
            }
        } else if(cmd == 71) {
            // Opcode: CALL
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := call(5, 0x42, 0, 0, 0x44, 0, 0x20)
                }
            }
        } else if(cmd == 72) {
            // Opcode: RETURN
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    return(5, 0x42)
                }
            }
        } else if(cmd == 73) {
            // Opcode: CALLCODE
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    let x := callcode(5, 0x42, 0, 0, 0x44, 0, 0x20)
                }
            }
        } else if(cmd == 74) {
            // Opcode: REVERT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    revert(5, 0x42)
                }
            }
        } else if(cmd == 75) {
            // Opcode: INVALID
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    invalid()
                }
            }
        } else if(cmd == 76) {
            // Opcode: SELFDESTRUCT
            for (uint32 i = 0; i < samples; i++) {
                assembly {
                    selfdestruct(0x42)
                }
            }
        } else {
            revert('Please enter a valid command.');
        }
    }
}
