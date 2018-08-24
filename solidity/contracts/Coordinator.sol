pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2; //solium-disable-line

// Coordinator handles oracle service aggreements between one or more oracles.
contract Coordinator {
  struct ServiceAgreement {
    uint256 payment;
    uint256 expiration;
    address[] oracles;
    bytes32 requestDigest;
  }

  mapping(bytes32 => ServiceAgreement) public serviceAgreements;

  event EmitString(string msg);
  event EmitAddress(address msg);
  event EmitV(uint8 v);
  event EmitA(bytes32 v);

  function getId(
    uint256 _payment,
    uint256 _expiration,
    address[] _oracles,
    bytes32 _requestDigest
  )
    public pure returns (bytes32)
  {
    return keccak256(abi.encodePacked(_payment, _expiration, _oracles, _requestDigest));
  }

  // XXX: No nested structs in web3
  //struct Signature {
    //uint8 v;
    //bytes32 r;
    //bytes32 s;
  //}

  function initiateServiceAgreement(
    uint256 _payment,
    uint256 _expiration,
    address[] _oracles,
    uint8[] _vs,
    bytes32[] _rs,
    bytes32[] _ss,
    // XXX: no nested structs in web3
    // bytes[][] _signatures,
    // Signature[] _signatures,
    bytes32 _requestDigest
  ) public
  {
    //require(_oracles.length == _signatures.length);

    for (uint i = 0; i < _oracles.length; i++) {
      emit EmitString("!!! SHOULD verify each participant");

      //bytes[] signature = _signatures[i];

      uint8 v = _vs[i];
      bytes32 r = _rs[i];
      bytes32 s = _ss[i];

      address signer = getOracleAddressFromSASignature(_requestDigest, v, r, s);// signature);
      emit EmitAddress(signer);

      address oracle = _oracles[i];
      emit EmitAddress(oracle);
      //require(
        //oracle == signer,
        //"!!! oracle is not the signer: TODO: can it do string interpolation of the addresses???"
      //);
    }

    bytes32 id = getId(_payment, _expiration, _oracles, _requestDigest);

    serviceAgreements[id] = ServiceAgreement(
      _payment,
      _expiration,
      _oracles,
      _requestDigest
    );
  }

  //function getOracleAddressFromSASignature(bytes32 _hash, bytes32 _sig) returns (address) {
  //function getOracleAddressFromSASignature(bytes32 _hash, bytes[] _sig) returns (address) {
  function getOracleAddressFromSASignature(bytes32 _hash, uint8 _v, bytes32 _r, bytes32 _s) returns (address) {
    //bytes32 r;
    //bytes32 s;
    //uint8 v;

    //if (sig.length != 65) {
      //return 0;
    //}

    //assembly {
      //v := byte(0, sig)
      //r := and(sig, 0x00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff)
      ////s := mload(add(sig, 33))
    //}

    // https://github.com/ethereum/go-ethereum/issues/2053
    //if (v < 27) {
      //v += 27;
    //}

    //if (v != 27 && v != 28) {
      //return 0;
    //}

    ///* prefix might be needed for geth only
     //* https://github.com/ethereum/go-ethereum/issues/3731
     //*/
    //// bytes memory prefix = "\x19Ethereum Signed Message:\n32";
    //// hash = sha3(prefix, hash);

    return ecrecover(_hash, _v, _r, _s);
  }
}
