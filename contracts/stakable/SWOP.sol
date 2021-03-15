pragma solidity <=0.7;

import "../ethereum/interfaces/Ownable.sol";
import "./StandartToken.sol";


contract SWOP is StandartToken {
  function name() external pure returns (string memory) {
    return "Wrapped SWOP";
  }

  function symbol() external pure returns (string memory) {
    return "gwSWOP";
  }

  function decimals() external pure returns (uint8) {
    return 18;
  }
}
