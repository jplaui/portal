// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "./IdentityInterface.sol";

contract IdentityManager {
    struct PrivateClaim {
        string value; // commitment depending on public or private
        string statement; // phi
        string ipfsCircuitMetadata; // ipfs circuit
        string eventHash; // verification event hash
        uint256 timestamp;
        int id;
        int version;
    }

    struct PublicClaim {
        string value; // value
        string statement; // phi
        uint256 timestamp;
        int id;
        int version;
    }

    struct Attestation {
        bytes signature; // signs(t, claim={commiment/public_value}, statement)
        address attestor; // issuer
        uint256 expires; // unix timestamp
        uint256 timestamp;
        int claimId;
        int id;
    }

    struct Revocation {
        address attestedTo;
        int attestationId;
        string status;
        uint256 timestamp;
    }

    IdentityInterface public registry;
    address public owner;

    //events
    event NewClaim(string key, string value, bool isPrivate, int version);
    event NewAttestation(address indexed identity, string key, uint256 timestamp);
    event AttestationRevoked(address indexed identity, int id, string reason);

    mapping(string => PrivateClaim) public privateClaims;
    mapping(string => PublicClaim) public publicClaims;
    mapping(string => string) public ipfsClaims;     // key: claim hash  value: ipfs URI

    mapping(string => Revocation) public revocations;
    mapping(string => Attestation) public attestations;

    modifier authorised {
        require(msg.sender == owner);
        _;
    }

    constructor(IdentityInterface _registry) {
        owner = msg.sender;
        registry = _registry;
    }

    function setPublicClaim(string calldata key, string calldata value,
        string calldata statement, uint256 timestamp, int id) external authorised {
        int version = publicClaims[key].version;
        publicClaims[key] = PublicClaim(value, statement, timestamp, id, version + 1);
        emit NewClaim(key, value, false, version + 1);
    }

    function setPrivateClaim(
        string memory key,
        string calldata value,
        string calldata statement,
        string calldata ipfsURI,
        string calldata eventHash,
        uint256 timestamp,
        int identifier
    ) external authorised {
        int version = privateClaims[key].version;
        privateClaims[key] = PrivateClaim(value, statement, ipfsURI, eventHash, timestamp, identifier, version + 1);
        emit NewClaim(key, value, true, version + 1);
    }

    function setAttestation(
        string calldata key,
        address attestor,
        uint256 expires,
        bytes calldata signature,
        int claimId,
        int id
    ) external authorised {
        // check if record not exists
        require(attestations[key].timestamp == 0);
        attestations[key] = Attestation(signature, attestor, expires, block.timestamp, claimId, id);
        emit NewAttestation(attestor, key, expires);
    }

    function revokeAttestation(
        string calldata key,
        string calldata reason,
        address attestedTo,
        int attestationId
    ) external authorised {
        revocations[key] = Revocation(attestedTo, attestationId, reason, block.timestamp);
        emit AttestationRevoked(attestations[key].attestor, attestations[key].id, reason);
    }

    function setClaimURI(
        string calldata key,
        string calldata value
    ) external authorised {
        ipfsClaims[key] = value;
        emit NewClaim(key, value, false, 0);
    }

    function deleteClaim(
        string calldata key,
        bool isPrivate
    ) external authorised {
        if (isPrivate) {
            delete privateClaims[key];
        } else {
            delete publicClaims[key];
        }
    }

    function deleteClaimURI(
        string calldata key
    ) external authorised {
        delete ipfsClaims[key];
    }

}
