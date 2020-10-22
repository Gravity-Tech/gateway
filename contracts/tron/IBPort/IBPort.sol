pragma solidity ^0.5;

import "../Token/Token.sol";
import "../interfaces/ISubscriberBytes.sol";
import "../libs/Queue.sol";

contract IBPort is ISubscriberBytes {
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
    Token public tokenAddress;

    uint public requestPosition = 1;

    mapping(uint => RequestStatus) public swapStatus;
    mapping(uint => UnwrapRequest) public unwrapRequests;
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
        if (d == 2) return RequestStatus.Rejected;
        if (d == 3) return RequestStatus.Success;
        if (d == 4) return RequestStatus.Returned;
        revert("invalid status");
    }

    function attachValue(bytes calldata value) external {
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
        Token(tokenAddress).mint(receiver, amount);
        swapStatus[swapId] = RequestStatus.Success;
    }

    function changeStatus(uint swapId, RequestStatus newStatus) internal {
        require(swapStatus[swapId] == RequestStatus.New, "invalid request status");
        swapStatus[swapId] = newStatus;
    }


    function createTransferUnwrapRequest(uint amount, bytes32 receiver) public {
        unwrapRequests[requestPosition] = UnwrapRequest(msg.sender, receiver, amount);
        swapStatus[requestPosition] = RequestStatus.New;
        tokenAddress.burnFrom(msg.sender, amount);
        emit RequestCreated(requestPosition, msg.sender, receiver, amount);
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
            homeAddress[count] = unwrapRequests[uint(p)].homeAddress;
            foreignAddress[count] = unwrapRequests[uint(p)].foreignAddress;
            amount[count] = unwrapRequests[uint(p)].amount;
            status[count] = swapStatus[uint(p)];
            count++;
        }

        return (id, homeAddress, foreignAddress, amount, status);
    }
}
