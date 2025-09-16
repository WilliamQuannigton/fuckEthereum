// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Counter
 * @dev A simple counter contract that allows incrementing, decrementing, and reading the count
 */
contract Counter {
    uint256 private count;
    
    // Events for tracking state changes
    event CountIncremented(uint256 newCount);
    event CountDecremented(uint256 newCount);
    event CountReset(uint256 newCount);
    
    // Constructor to initialize the counter
    constructor() {
        count = 0;
    }
    
    /**
     * @dev Increment the counter by 1
     * @return The new count value
     */
    function increment() public returns (uint256) {
        count += 1;
        emit CountIncremented(count);
        return count;
    }
    
    /**
     * @dev Decrement the counter by 1
     * @return The new count value
     */
    function decrement() public returns (uint256) {
        require(count > 0, "Counter cannot be negative");
        count -= 1;
        emit CountDecremented(count);
        return count;
    }
    
    /**
     * @dev Reset the counter to 0
     * @return The new count value
     */
    function reset() public returns (uint256) {
        count = 0;
        emit CountReset(count);
        return count;
    }
    
    /**
     * @dev Get the current count value
     * @return The current count
     */
    function getCount() public view returns (uint256) {
        return count;
    }
    
    /**
     * @dev Get the current count value (alternative function name)
     * @return The current count
     */
    function getCurrentCount() public view returns (uint256) {
        return count;
    }
}
