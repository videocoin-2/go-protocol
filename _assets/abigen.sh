#!/bin/bash

proto_path=$(readlink -f $1)
output=$2

solc --abi -o $output --overwrite --allow-paths $proto_path openzeppelin-solidity=$proto_path/node_modules/openzeppelin-solidity $proto_path/contracts/StakingManager.sol $proto_path/contracts/Stream.sol $proto_path/contracts/StreamManager.sol