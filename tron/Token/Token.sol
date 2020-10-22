pragma solidity ^0.5;

import "../@openzeppelin/contracts/token/ERC20/ERC20Burnable.sol";
import "../@openzeppelin/contracts/token/ERC20/ERC20Mintable.sol";

contract Token is ERC20Burnable, ERC20Mintable {
    string public name = "TST";
    uint public decimals = 18;
}

