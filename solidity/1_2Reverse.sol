// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

//反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
contract Reverse {
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