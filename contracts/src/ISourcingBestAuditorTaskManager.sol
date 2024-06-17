// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/libraries/BN254.sol";

interface ISourcingBestAuditorTaskManager {
    // EVENTS
    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskIndex);

    event TaskChallengedSuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    event TaskChallengedUnsuccessfully(
        uint32 indexed taskIndex,
        address indexed challenger
    );

    event TaskCreated(
        uint256 taskId,
        uint256 budget,
        bytes AuditJobSpecificationURI,
        bytes quorumNumbers,
        uint32 quorumThresholdPercentage
    );

    // STRUCTS
    struct Task {
        uint256 taskId;
        bytes AuditJobSpecificationURI;
        uint256 budget;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
        Bid[] bids;
    }

    struct Bid {
        uint256 bidId;
        string zkp;
        string bidPitchDocURI;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        uint32 referenceTaskIndex; // Can be obtained by the operator from the event NewTaskCreated.
        string evaluatedBidURI; // URI pointing to the operator's evaluation of the bids.
        uint256 selectedBidId; // The ID of the selected best bid.
        // Other fields related to the operator's response
    }

    // Extra information related to taskResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than TaskResponse
    // This metadata is needed by the challenger, so we emit it in the TaskResponded event
    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    mapping(uint256 => Task) public tasks;
    uint256 public taskCounter;

    // FUNCTIONS
    // NOTE: this function creates new task.
    function createNewTask(
        uint256 numberToBeSquared,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    // // NOTE: this function raises challenge to existing tasks.
    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external;

    /// @notice Returns the TASK_RESPONSE_WINDOW_BLOCK
    function getTaskResponseWindowBlock() external view returns (uint32);
}
