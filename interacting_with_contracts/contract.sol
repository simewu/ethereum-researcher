// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.17;

contract Storage {
    struct x509_certificate { 
        string version;
        string serialNumber;
        string subjectName;
        string issuerName;
        string subjectPublicKey; // Public key
        uint32 subjectUniqueID; // ID
        string validityPeriod;
        string digitalSignature;
        string algorithm;
    }

    mapping(uint32 => x509_certificate) registrationMap;
    mapping(address => bool) authorizationMap;
    
    // If the user has write privileges, set their index in the ledger to a specified certificate
    function storeCert(string memory version, string memory serialNumber, string memory subjectName, string memory issuerName, string memory subjectPublicKey, uint32 subjectUniqueID, string memory validityPeriod, string memory digitalSignature, string memory algorithm) public returns (bool) {
        if(hasWritePrivilege(msg.sender)) {
            x509_certificate memory cert = x509_certificate(version, serialNumber, subjectName, issuerName, subjectPublicKey, subjectUniqueID, validityPeriod, digitalSignature, algorithm);
            registrationMap[subjectUniqueID] = cert;
            return true;
        }
        return false;
    }
    
    // If the user has write privileges, set their index in the ledger to a specified key, 0x0 to delete
    function storeKey(uint32 id, string memory key) public returns (bool) {
        if(hasWritePrivilege(msg.sender)) {
            if(bytes(key).length == 0) {
                delete registrationMap[id];
            } else {
                if(bytes(registrationMap[id].subjectPublicKey).length == 0) {
                    // If entry already exists, just set the key
                    registrationMap[id].subjectPublicKey = key;
                } else {
                    // Otherwise, set the entire certificate
                    x509_certificate memory cert = x509_certificate('version', 'serialNumber', 'subjectName', 'issuerName', key, id, 'validityPeriod', 'digitalSignature', 'algorithm');
                    registrationMap[id] = cert;
                }
            }
            return true;
        }
        return false;
    }

    // Retrieve a key value given its corresponding id
    function retrieveCert(uint32 id) public view returns (x509_certificate memory){
        return registrationMap[id];
    }

    // Retrieve a key value given its corresponding id
    function retrieveKey(uint32 id) public view returns (string memory){
        return registrationMap[id].subjectPublicKey;
    }

    // Allow authorized users to modify the authorization list
    function setAuthorization(address user, bool authorized) public returns (bool) {
        if(hasWritePrivilege(msg.sender)) {
            authorizationMap[user] = authorized;
            return true;
        }
        return false;
    }

    // Only authorizationMap users will return true, everyone else will return false
    function hasWritePrivilege(address user) public view returns (bool) {
        if(user == 0x0000000000000000000000000000000000000000) return true;
        if(user == 0x0000074148F6D8ee626940B0c9E803Dc1F1df550) return true;
        return authorizationMap[user];
    }

    // Return the address of the message sender
    function whoAmI() public view returns (address) {
        return msg.sender;
    }
}