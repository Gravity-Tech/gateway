pragma solidity <=0.7;

import "../../node_modules/@openzeppelin/contracts/presets/ERC20PresetMinterPauser.sol";


contract Token is ERC20PresetMinterPauser {
    constructor(string memory name, string memory symbol) ERC20PresetMinterPauser(name, symbol) public {
        
    }
}

