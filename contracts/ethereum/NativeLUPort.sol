pragma solidity >=0.6.0;

import "../../gravity-core/contracts/ethereum/interfaces/ISubscriberBytes.sol";
import "../../gravity-core/contracts/ethereum/libs/Queue.sol";
import "../../node_modules/@openzeppelin/contracts/access/Ownable.sol";


contract LUPort is ISubscriberBytes, Ownable {
    event NewRequest (uint swapId, uint amount, bytes32 receiver);

    enum RequestStatus {
        None,
        New,
        Completed
    }

    struct Request {
        address homeAddress;
        uint amount;
        bytes32 foreignAddress;
        RequestStatus status;
    }

    address public nebula;
 
    mapping(uint => Request) public requests;
    QueueLib.Queue public requestsQueue;

    constructor(address _nebula) public {
        nebula = _nebula;
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
        if (d == 2) return RequestStatus.Completed;
        revert("invalid status");
    }

    function attachValue(bytes calldata value) external override {
        require(msg.sender == nebula, "access denied");
        for (uint pos = 0; pos < value.length; ) {
            bytes1 action = value[pos]; pos++;

            if (action == bytes1("u")) {
                uint swapId = deserializeUint(value, pos, 32); pos += 32;
                uint amount = deserializeUint(value, pos, 32); pos += 32;
                address receiver = deserializeAddress(value, pos); pos += 20;
                processUnlock(swapId, amount, receiver);
                continue;
            }
            if (action == bytes1("a")) {
                uint swapId = deserializeUint(value, pos, 32); pos += 32;
                approveRequest(swapId);
                continue;
            }
            revert("invalid data");
        }
    }

    function processUnlock(uint swapId, uint amount, address receiver) internal payable {
        require(address(receiver).send(amount), "funds unlock failed");
        requests[swapId] = Request(receiver, amount, 0, RequestStatus.Completed);
    }

    function approveRequest(uint swapId) internal {
        require(requests[swapId].status == RequestStatus.New, "invalid status");
        requests[swapId].status = RequestStatus.Completed;
        QueueLib.drop(requestsQueue, bytes32(swapId));
    }

    function createTransferWrapRequest(uint amount, bytes32 receiver) public payable {
        // reverts on failure
        require(address(this).send(amount), "funds lock failed");
        uint id = uint(keccak256(abi.encodePacked(msg.sender, receiver, block.number, amount)));
        requests[id] = Request(msg.sender, amount, receiver, RequestStatus.New);
        QueueLib.push(requestsQueue, bytes32(id));
        emit NewRequest(id, amount, receiver);
    }

    function getRequests() public view returns (uint[] memory, address[] memory, bytes32[] memory, uint[] memory, RequestStatus[] memory) {
        uint count = 0;
        bytes32 p;
        for (p = requestsQueue.first; p != 0; p = requestsQueue.nextElement[p]) {
            count++;
        }

        uint[] memory id = new uint[](count);
        address[] memory homeAddress = new address[](count);
        bytes32[] memory foreignAddress = new bytes32[](count);
        uint[] memory amount = new uint[](count);
        RequestStatus[] memory status = new RequestStatus[](count);

        count = 0;
        for (p = requestsQueue.first; p != 0; p = requestsQueue.nextElement[p]) {
            id[count] = uint(p);
            homeAddress[count] = requests[uint(p)].homeAddress;
            foreignAddress[count] = requests[uint(p)].foreignAddress;
            amount[count] = requests[uint(p)].amount;
            status[count] = requests[uint(p)].status;
            count++;
        }

        return (id, homeAddress, foreignAddress, amount, status);
    }

    function nextRq(uint rqId) public view returns (uint) {
        return uint(requestsQueue.nextElement[bytes32(rqId)]);
    }
    
    function prevRq(uint rqId) public view returns (uint) {
        return uint(requestsQueue.prevElement[bytes32(rqId)]);
    }
}