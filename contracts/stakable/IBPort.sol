pragma solidity >=0.6.0;


import "./Usdn.sol";
import "../../gravity-core/contracts/ethereum/interfaces/ISubscriberBytes.sol";
import "../../gravity-core/contracts/ethereum/libs/Queue.sol";

contract IBPort is ISubscriberBytes, Ownable {
    enum RequestStatus {
        None,
        New,
        Rejected,
        Success,
        Returned
    }

    struct UnwrapRequest {
        address homeAddress;
        bytes32 foreignAddress;
        uint amount;
    }

    event RequestCreated(uint, address, bytes32, uint);

    address public nebula;
    USDN public token;

    mapping(uint => RequestStatus) public swapStatus;
    mapping(uint => UnwrapRequest) public unwrapRequests;
    QueueLib.Queue public requestsQueue;

    constructor(address _nebula, address _tokenAdress) public {
        nebula = _nebula;
        token = USDN(_tokenAdress);
    }

    function transferTokenOwnership(address newOwner) external virtual onlyOwner {
        token.transferOwnership(newOwner);
    }

    function deserializeUint(bytes memory b, uint startPos, uint len) internal pure returns (uint) {
        uint v = 0;
        for (uint p = startPos; p < startPos + len; p++) {
            v = v * 256 + uint(uint8(b[p]));
        }
        return v;
    }

    function deserializeAddress(bytes memory b, uint startPos) internal pure returns (address) {
        return address(uint160(deserializeUint(b, startPos, 20)));
    }

    function deserializeStatus(bytes memory b, uint pos) internal pure returns (RequestStatus) {
        uint d = uint(uint8(b[pos]));
        if (d == 0) return RequestStatus.None;
        if (d == 1) return RequestStatus.New;
        if (d == 2) return RequestStatus.Rejected;
        if (d == 3) return RequestStatus.Success;
        if (d == 4) return RequestStatus.Returned;
        revert("invalid status");
    }

    function attachValue(bytes calldata value) override external {
        require(msg.sender == nebula, "access denied");
        for (uint pos = 0; pos < value.length; ) {
            bytes1 action = value[pos]; pos++;

            if (action == bytes1("m")) {
                uint swapId = deserializeUint(value, pos, 32); pos += 32;
                uint amount = deserializeUint(value, pos, 32); pos += 32;
                address receiver = deserializeAddress(value, pos); pos += 20;
                mint(swapId, amount, receiver);
                continue;
            }

            if (action == bytes1("c")) {
                uint swapId = deserializeUint(value, pos, 32); pos += 32;
                RequestStatus newStatus = deserializeStatus(value, pos); pos += 1;
                changeStatus(swapId, newStatus);
                continue;
            }
            revert("invalid data");
        }
    }

    function mint(uint swapId, uint amount, address receiver) internal {
        require(swapStatus[swapId] == RequestStatus.None, "invalid request status");
        if (receiver == address(this)) {
            token.stake(amount);
        }
        else {
            token.deposit(receiver, amount); // function deposit(address account, uint256 amount) external onlyOwner
        }
        swapStatus[swapId] = RequestStatus.Success;
    }

    function changeStatus(uint swapId, RequestStatus newStatus) internal {
        require(swapStatus[swapId] == RequestStatus.New, "invalid request status");
        swapStatus[swapId] = newStatus;
        QueueLib.drop(requestsQueue, bytes32(swapId));
    }

    function createTransferUnwrapRequest(uint amount, bytes32 receiver) public {
        uint id = uint(keccak256(abi.encodePacked(msg.sender, receiver, block.number, amount)));
        unwrapRequests[id] = UnwrapRequest(msg.sender, receiver, amount);
        swapStatus[id] = RequestStatus.New;
        uint remainBalance = token.balanceOf(msg.sender) - amount;
        token.withdraw(msg.sender);
        token.deposit(msg.sender, remainBalance);
        QueueLib.push(requestsQueue, bytes32(id));
        emit RequestCreated(id, msg.sender, receiver, amount);
    }

    function nextRq(uint rqId) public view returns (uint) {
        return uint(requestsQueue.nextElement[bytes32(rqId)]);
    }
     function prevRq(uint rqId) public view returns (uint) {
        return uint(requestsQueue.prevElement[bytes32(rqId)]);
    }
}

