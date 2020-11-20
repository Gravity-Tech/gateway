pragma solidity <=0.7;


abstract contract Ownable {
  address private _owner;
  address private _admin;

  constructor () internal {
    _owner = msg.sender;
    _admin = msg.sender;
  }

  modifier onlyOwner() {
    require(_owner == msg.sender || _admin == msg.sender, "Ownable: caller is not the owner or admin");
    _;
  }

  function transferOwnership(address newOwner) external virtual onlyOwner {
    require(newOwner != address(0), "Ownable: new owner is the zero address");
    _owner = newOwner;
  }
}