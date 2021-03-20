pragma solidity <=0.7;

import "./StandartToken.sol";


contract USDN is StandartToken {
  function name() external pure returns (string memory) {
    return "Wrapped Neutrino USD";
  }

  function symbol() external pure returns (string memory) {
    return "gwUSDN";
  }

  function decimals() external pure returns (uint8) {
    return 18;
  }
}
