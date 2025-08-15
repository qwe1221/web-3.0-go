//创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数

// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

contract Voting {
    mapping(string => uint256) public mappingVotes;
    string[] public candidates = ["zhangsan", "lisi", "wangwu"];

   
    function vote(string memory candidate) public {
        mappingVotes[candidate]++;
    }

    function getVotes(string memory candidate) public view returns (uint256) {
        return mappingVotes[candidate];
    }

    function resetVotes() public {
        for (uint256 i = 0; i < candidates.length ; i++) {
            delete mappingVotes[candidates[i]];
        }
    }

    function everseString(string memory str) public pure returns (string memory) {
        bytes memory strBytes = bytes(str); // 转为 bytes 以便按索引访问
        uint256 len = strBytes.length;
        bytes memory result = new bytes(len); // 创建相同长度的 bytes
        for (uint256 i = 0; i < len ;i++ ){
            result[i] = strBytes[len - 1 - i];
        }
        return string(result);
    }
}