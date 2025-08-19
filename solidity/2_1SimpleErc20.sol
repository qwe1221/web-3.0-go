// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SimpleErc20 is IERC20, Ownable {
     // 存储账户余额
    mapping(address => uint256) private _balances;
    mapping(address => mapping (address =>uint256 )) private _allowances;
    uint256 private _totalSupply;//全网代币总供应量，只读不修改状态
    // 代币名称和符号
    string private _name;
    string private _symbol;
    // 代币小数位数
    uint8 private _decimals = 18;
    //事件定义
    event _Transfer(address indexed from, address indexed to, uint256 value);
    event _Approval(address indexed owner, address indexed spender, uint256 value);

    constructor(string memory name, string memory symbol, address initialOwner) Ownable(initialOwner) {
        _name = name;
        _symbol = symbol;
    }

    //接口实现
    function name() public view  returns (string memory) {
        return _name;
    }
    function symbol() public view  returns (string memory) {
        return _symbol;
    }
    function decimals() public view  returns (uint8) {
        return _decimals;
    }
 
    function totalSupply() public view override returns (uint256) {
        return _totalSupply;
    }
 
    function balanceOf(address account) public view override returns (uint256) {
        return _balances[account];
    }

    function transfer(address recipient, uint256 amount) public virtual override returns(bool) {
        require(_balances[recipient] >= amount,"Insufficient balance for transfer"); //检查余额是否足够;
        _transfer(msg.sender, recipient, amount);
        return true;
    }

     function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }
     function approve(address spender, uint256 amount) public virtual override returns (bool) {
        _approve(_msgSender(), spender, amount);
        return true;
    }
     function transferFrom(address sender, address recipient, uint256 amount) public virtual override returns (bool) {
        //执行代币转移
        _transfer(sender, recipient, amount);
        //更新授权额度
        _approve(sender, _msgSender(), _allowances[sender][_msgSender()] - amount);
        return true;
    }
    //增发代币函数（仅合约所有者可用）
    function mint(address to,uint256 amount) public onlyOwner{
        _mint(to,amount);
    }

    //内部函数，增发代币,零地址没有私钥，无法访问其中的代币，会导致代币永久丢失
    function _mint(address account,uint256 amount) internal virtual {
        require(account != address(0),"do not transfer to the zero address");
        //增加总供应量
        _totalSupply += amount;
        //更新目标地址得余额
        _balances[account] += amount;
        //触发事件记录代币增发事件便于浏览器和钱包跟踪
        emit _Transfer(address(0),account,amount);
    }

    function _transfer(address sender, address recipient, uint256 amount) internal virtual {
        require(sender != address(0),"do not transfer to the zero address");
        require(recipient != address(0),"do not transfer to the zero address");
        //检查转账金额是否足够
        require(_balances[sender] >= amount,"Insufficient balance for transfer");
        //更新转账双方的余额
        _balances[sender] -= amount;
        _balances[recipient] += amount;
        //发出转账事件
        emit _Transfer(sender,recipient,amount);
    }
    //定义一个内部函数 _approve用于设置代币授权
    function _approve(address owner,address spender,uint256 amount) internal virtual {
        require(owner != address(0),"do not transfer to the zero address");
        require(spender != address(0),"do not transfer to the zero address");
        _allowances[owner][spender] = amount;
        emit _Approval(owner,spender,amount);
    }
    //定义一个内部函数 _msgSender用于获取当前调用者的地址
    function _msgSender() internal view override  returns (address) {
        return msg.sender;
    }
}