pragma solidity <=0.7;

import "../../node_modules/@openzeppelin/contracts/presets/ERC20PresetMinterPauser.sol";
import "../../node_modules/@openzeppelin/contracts/access/Ownable.sol";


contract Token is ERC20PresetMinterPauser, Ownable {
    constructor(string memory name, string memory symbol) ERC20PresetMinterPauser(name, symbol) public {
        
    }
}

