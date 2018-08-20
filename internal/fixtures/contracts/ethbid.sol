pragma solidity ^0.4.24;

import "solidity/contracts/Chainlinked.sol";

contract RopstenConsumer is Chainlinked, Ownable {
  uint256 public currentPrice;

  bytes32 constant SPEC_ID = bytes32("3e8f43664bf1449baae3903b44565e7f");

  event RequestEthereumPriceFulfilled(
    bytes32 indexed requestId,
    uint256 indexed price
  );

  constructor(address _linkAddress, address _oracleAddress) Ownable() public {
    setLinkToken(_linkAddress);
    setOracle(_oracleAddress);
  }

  function requestEthereumPrice()
    public
    onlyOwner
  {
    ChainlinkLib.Run memory run = newRun(SPEC_ID, this, "fulfillEthereumPrice(bytes32,uint256)");
    run.add("url", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,EUR,JPY");
    string[] memory path = new string[](1);
    path[0] = "USD";
    run.addStringArray("path", path);
    run.add("someExtra", "my_value");
    run.addInt("times", 100);
    chainlinkRequest(run, LINK(1));
  }

  function fulfillEthereumPrice(bytes32 _requestId, uint256 _price)
    public
    checkChainlinkFulfillment(_requestId)
  {
    emit RequestEthereumPriceFulfilled(_requestId, _price);
    currentPrice = _price;
  }

  function withdrawLink() public onlyOwner {
    require(link.transfer(owner, link.balanceOf(address(this))), "Unable to transfer");
  }

}
