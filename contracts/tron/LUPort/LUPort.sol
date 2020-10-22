pragma solidity ^0.5;

import "../Token/Token.sol";
import "../interfaces/ISubscriberBytes.sol";
import "../libs/Queue.sol";

contract LUPort is ISubscriberBytes, QueueLib {
    address public nebula;
    Token public tokenAddress;

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

    uint public lastReqId = 1;

    mapping(uint => Request) public requests;
    QueueLib.Queue public requestsQueue;

    constructor(address _nebula, address _tokenAddress) public {
        nebula = _nebula;
        tokenAddress = Token(_tokenAddress);
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

    function attachValue(bytes calldata value) external {
        // require(msg.sender == nebula, "access denied");
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

    function processUnlock(uint swapId, uint amount, address receiver) internal {
        require(tokenAddress.transfer(receiver, amount), "can't transfer from contract");
        requests[swapId] = Request(receiver, amount, 0, RequestStatus.Completed);
    }

    function approveRequest(uint swapId) internal {
        require(requests[swapId].status == RequestStatus.New, "invalid status");
        requests[swapId].status = RequestStatus.Completed;
        Queue_drop(requestsQueue, bytes32(swapId));
    }

    function createTransferUnwrapRequest(uint amount, bytes32 receiver) public {
        require(tokenAddress.transferFrom(msg.sender, address(this), amount), "can't transfer from");
        requests[lastReqId] = Request(msg.sender, amount, receiver, RequestStatus.New);
        Queue_push(requestsQueue, bytes32(lastReqId));
        emit NewRequest(lastReqId, amount, receiver);
        lastReqId++;
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
}
