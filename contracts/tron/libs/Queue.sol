pragma solidity ^0.5;

contract QueueLib {
    struct Queue {
        bytes32 first;
        bytes32 last;
        mapping(bytes32=>bytes32) nextElement;
        mapping(bytes32=>bytes32) prevElement;
    }

    function Queue_drop(Queue storage queue, bytes32 rqHash) internal {
        bytes32 prevElement = queue.prevElement[rqHash];
        bytes32 nextElement = queue.nextElement[rqHash];

        if (prevElement != bytes32(0)) {
            queue.nextElement[prevElement] = nextElement;
        } else {
            queue.first = nextElement;
        }

        if (nextElement != bytes32(0)) {
            queue.prevElement[nextElement] = prevElement;
        } else {
            queue.last = prevElement;
        }
    }

    function Queue_push(Queue storage queue, bytes32 elementHash) internal {
        if (queue.first == 0x000) {
            queue.first = elementHash;
            queue.last = elementHash;
        } else {
            queue.nextElement[queue.last] = elementHash;
            queue.prevElement[elementHash] = queue.last;
            queue.nextElement[elementHash] = bytes32(0);
            queue.last = elementHash;
        }
    }

}