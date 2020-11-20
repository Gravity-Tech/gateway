pragma solidity <=0.7;

import "./interfaces/Ownable.sol";
import "../../node_modules/@openzeppelin/contracts/presets/ERC20PresetMinterPauser.sol";

contract Token is ERC20PresetMinterPauser, Ownable {
    constructor(string memory name, string memory symbol) ERC20PresetMinterPauser(name, symbol) public {
        
    }

    function transferOwnership(address owner) external {
        require(hasRole(MINTER_ROLE, msg.sender), "the caller is not a minter");
        require(hasRole(PAUSER_ROLE, msg.sender), "the caller is not a pauser");

        grantRole(MINTER_ROLE, owner);
        grantRole(PAUSER_ROLE, owner);
    }
}

