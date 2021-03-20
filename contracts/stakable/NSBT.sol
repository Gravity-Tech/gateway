pragma solidity <=0.7;

import "./StandartToken.sol";


contract NSBT is StandartToken {
  function name() external pure returns (string memory) {
    return "Wrapped Autostaking NSBT";
  }

  function symbol() external pure returns (string memory) {
    return "gwaNSBT";
  }

  function decimals() external pure returns (uint8) {
    return 18;
  }
}
