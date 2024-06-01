// SPDX-License-Identifier: MIT

pragma solidity =0.8.12;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/utils/Context.sol";
import "./IncredibleSquaringTaskManager.sol";

contract WBTC is Context, IERC20 {
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    address immutable private taskManagerAddress;
    uint256 private _totalSupply;

    // Struct to store mint requests
    struct MintRequestInfo {
        string signedMessage;
        string transactionHash;
        address recipient;
        uint256 amount;
        bool fulfilled;
    }

    // Mapping to store mint requests
    mapping(bytes32 => MintRequestInfo) public mintRequests;

    // Event for special mint request
    event MintRequest(
        bytes32 indexed requestId,
        address indexed requester,
        string signedMessage,
        string transactionHash,
        address recipient,
        uint256 amount
    );

    constructor(address _taskManagerAddress) {
        taskManagerAddress = _taskManagerAddress;
    }

    /**
     * @dev See {IERC20-totalSupply}.
     */
    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev See {IERC20-balanceOf}.
     */
    function balanceOf(
        address account
    ) public view virtual override returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev Regular mint function to create tokens.
     * Can be called by anyone.
     */
    function mint(address account, uint256 amount) public {
        _mint(account, amount);
    }

    /**
     * @dev Burns tokens from a specific account.
     * Includes a bitcoinAddress parameter for additional processing.
     */
    function burn(address account, uint256 amount, bytes memory bitcoinAddress) public {
        _burn(account, amount, bitcoinAddress);
    }

    /**
     * @dev See {IERC20-transfer}.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - the caller must have a balance of at least `amount`.
     */
    function transfer(
        address to,
        uint256 amount
    ) public virtual override returns (bool) {
        address owner = _msgSender();
        _transfer(owner, to, amount);
        return true;
    }

    /**
     * @dev See {IERC20-allowance}.
     */
    function allowance(
        address owner,
        address spender
    ) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev See {IERC20-approve}.
     *
     * NOTE: If `amount` is the maximum `uint256`, the allowance is not updated on
     * `transferFrom`. This is semantically equivalent to an infinite approval.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function approve(
        address /*spender*/,
        uint256 /*amount*/
    ) public virtual override returns (bool) {
        return true;
    }

    /**
     * @dev See {IERC20-transferFrom}.
     *
     * Emits an {Approval} event indicating the updated allowance. This is not
     * required by the EIP. See the note at the beginning of {ERC20}.
     *
     * NOTE: Does not update the allowance if the current allowance
     * is the maximum `uint256`.
     *
     * Requirements:
     *
     * - `from` and `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     * - the caller must have allowance for ``from``'s tokens of at least
     * `amount`.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public virtual override returns (bool) {
        _transfer(from, to, amount);
        return true;
    }

    /**
     * @dev Special mint request function.
     * Emits a MintRequest event with the provided parameters.
     * Saves the request details for later fulfillment.
     */
    function requestMint(string memory signedMessage, string memory transactionHash, address recipient, uint256 amount) public {
        bytes32 requestId = keccak256(abi.encodePacked(signedMessage, transactionHash, recipient, amount, block.timestamp));
        mintRequests[requestId] = MintRequestInfo({
            signedMessage: signedMessage,
            transactionHash: transactionHash,
            recipient: recipient,
            amount: amount,
            fulfilled: false
        });
        emit MintRequest(requestId, msg.sender, signedMessage, transactionHash, recipient, amount);
    }

    /**
     * @dev Function to fulfill a mint request.
     * Only allows fulfilling requests that have not been fulfilled yet.
     */
    function fulfillMintRequest(bytes32 requestId) public {
        MintRequestInfo storage request = mintRequests[requestId];
        require(!request.fulfilled, "Mint request already fulfilled");
        request.fulfilled = true;
        _mint(request.recipient, request.amount);
    }

    /**
     * @dev Internal transfer function.
     * Moves `amount` of tokens from `from` to `to`.
     * Emits a {Transfer} event.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     */
    function _transfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        require(
            _balances[from] >= amount,
            "ERC20: transfer amount exceeds balance"
        );
        unchecked {
            _balances[from] = _balances[from] - amount;
            _balances[to] += amount;
        }

        emit Transfer(from, to, amount);
    }

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     */
    function _mint(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: mint to the zero address");

        _totalSupply += amount;
        unchecked {
            _balances[account] += amount;
        }
        emit Transfer(address(0), account, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`, reducing the
     * total supply.
     * Includes a bitcoinAddress parameter for additional processing.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function _burn(address account, uint256 amount, bytes memory bitcoinAddress) internal virtual {
        require(account != address(0), "ERC20: burn from the zero address");

        _beforeTokenTransfer(bitcoinAddress, amount);

        uint256 accountBalance = _balances[account];
        require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
        unchecked {
            _balances[account] = accountBalance - amount;
            _totalSupply -= amount;
        }

        emit Transfer(account, address(0), amount);

        _afterTokenTransfer(account, address(0), amount);
    }

    /**
     * @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
     * This internal function is equivalent to `approve`, and can be used to
     * e.g. set automatic allowances for certain subsystems, etc.
     *
     * Emits an {Approval} event.
     *
     * Requirements:
     *
     * - `owner` cannot be the zero address.
     * - `spender` cannot be the zero address.
     */
    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) internal virtual {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Updates `owner` s allowance for `spender` based on spent `amount`.
     * Does not update the allowance amount in case of infinite allowance.
     * Revert if not enough allowance is available.
     *
     * Might emit an {Approval} event.
     */
    function _spendAllowance(
        address owner,
        address spender,
        uint256 amount
    ) internal virtual {
        uint256 currentAllowance = allowance(owner, spender);
        require(currentAllowance >= amount, "ERC20: insufficient allowance");
    }

    /**
     * @dev Hook that is called before any transfer of tokens. This includes
     * minting and burning.
     * Calls an external task manager to handle pre-transfer tasks.
     *
     * Calling conditions:
     *
     * - when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
     * will be transferred to `to`.
     * - when `from` is zero, `amount` tokens will be minted for `to`.
     * - when `to` is zero, `amount` of ``from``'s tokens will be burned.
     * - `from` and `to` are never both zero.
     */
    function _beforeTokenTransfer(
        bytes memory bitcoinAddress,
        uint256 amount
    ) internal virtual {
        IncredibleSquaringTaskManager(taskManagerAddress).createNewTask(msg.sender, amount, 100, abi.encode(0));
    }

    /**
     * @dev Hook that is called after any transfer of tokens. This includes
     * minting and burning.
     *
     * Calling conditions:
     *
     * - when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
     * has been transferred to `to`.
     * - when `from` is zero, `amount` tokens have been minted for `to`.
     * - when `to` is zero, `amount` of ``from``'s tokens have been burned.
     * - `from` and `to` are never both zero.
     */
    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {}
}
