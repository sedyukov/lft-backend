// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.7.6;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `from` to `to` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);
}


/**
 * @dev Interface for the optional metadata functions from the ERC20 standard.
 *
 * _Available since v4.1._
 */
interface IERC20Metadata is IERC20 {
    /**
     * @dev Returns the name of the token.
     */
    function name() external view returns (string memory);

    /**
     * @dev Returns the symbol of the token.
     */
    function symbol() external view returns (string memory);

    /**
     * @dev Returns the decimals places of the token.
     */
    function decimals() external view returns (uint8);
}


/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}


/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}


contract LevelFiveToken is IERC20, IERC20Metadata, Context, Ownable {

    string private constant TOKEN_NAME = "Level Five Token";
    string private constant TOKEN_SYMBOL = "LFT";
    uint256 private constant INITIAL_TOKEN_AMOUNT = 1e27; // 1,000,000,000
    uint256 private constant REFERRAL_MIN_AMOUNT = 1e21; // 1,000

    uint8 private constant FEE_DEVELOPER = 1;
    uint8 private constant FEE_STAKERS = 2;
    uint8 private constant FEE_LEVEL_1 = 1;
    uint8 private constant FEE_LEVEL_2 = 1;
    uint8 private constant FEE_LEVEL_3 = 1;
    uint8 private constant FEE_LEVEL_4 = 1;
    uint8 private constant FEE_LEVEL_5 = 3;

    ////////////////////////////////////////////////////////////////

    event Register(address indexed referral, address trader);
    event Stake(address indexed staker, uint256 amount);
    event Unstake(address indexed staker, uint256 amount);
    event RewardReferral(address indexed trader, address indexed referral, uint8 indexed level, uint256 amount);
    event RewardStakers(address indexed trader, uint256 amount);

    ////////////////////////////////////////////////////////////////

    string private _name;
    string private _symbol;

    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;

    uint256 private _totalSupply;

    ////////////////////////////////////////////////////////////////

    address private _developer;
    address private _lpToken;

    mapping(address => address[5]) private _referrals;

    mapping(address => uint256) private _shares;
    uint256 private _totalShare;

    ////////////////////////////////////////////////////////////////

    constructor(address developer, address[5] memory refs) {
        require(developer != address(0), "Developer cannot be zero address");
        require(refs.length == 5, "Five initial referrals required");

        _name = TOKEN_NAME;
        _symbol = TOKEN_SYMBOL;

        // set up the Developer and the LP token address
        _developer = developer;
        _lpToken = 0x997af3d9295df06B511ff7121Cf9D3eF6f65E749; // TODO: calculate it instead

        // mint initial amount of the token to the owner
        address deployer = _msgSender();
        _mint(deployer, INITIAL_TOKEN_AMOUNT);

        // transfer minimal required amount of tokens to initial referrals
        _transfer(deployer, refs[0], REFERRAL_MIN_AMOUNT);
        _transfer(deployer, refs[1], REFERRAL_MIN_AMOUNT);
        _transfer(deployer, refs[2], REFERRAL_MIN_AMOUNT);
        _transfer(deployer, refs[3], REFERRAL_MIN_AMOUNT);
        _transfer(deployer, refs[4], REFERRAL_MIN_AMOUNT);

        // register deployer and initial referrals
        _register(deployer, address(0));
        _register(refs[0], deployer);
        _register(refs[1], refs[0]);
        _register(refs[2], refs[1]);
        _register(refs[3], refs[2]);
        _register(refs[4], refs[3]);
    }

    function lpToken() public view returns (address) {
        return _lpToken;
    }

    // TODO: Remove this! This is for tests only!
    function setLPToken(address lpt) public onlyOwner {
        _lpToken = lpt;
    }
    /**
     * @dev Returns amount of currently staked tokens (which are also available to unstake right now).
     */
    function staked(address staker) public view returns (uint256) {
        uint256 totalStaked = _balances[address(this)];
        if (totalStaked == 0 || _totalShare == 0) {
            return 0;
        }
        return _shares[staker] * totalStaked / _totalShare;
    }

    function referrals(address trader) public view returns (address[5] memory) {
        return _referrals[trader];
    }

    function share(address trader) public view returns (uint256) {
        return _shares[trader];
    }

    function totalShare() public view returns (uint256) {
        return _totalShare;
    }

    function acceptInvite(address referral) public {
        address trader = _msgSender();
        require(_referrals[trader][4] == address(0), "Registered already");
        require(_referrals[referral][4] != address(0), "Provided referral is not registered");

        // register new trader
        _register(trader, referral);

        emit Register(referral, trader);
    }

    function stake(uint256 amount) public {
        address staker = _msgSender();
        uint256 totalStaked = _balances[address(this)];
        uint256 shares = amount;
        if (totalStaked > 0 && _totalShare > 0) {
            shares = amount * _totalShare / totalStaked;
        }

        // update storage
        _shares[staker] += shares;
        _totalShare += shares;

        // transfer tokens to the contract
        _transfer(staker, address(this), amount);

        emit Stake(staker, amount);
    }

    function unstake(uint256 amount) public {
        address staker = _msgSender();
        uint256 totalStaked = _balances[address(this)];
        uint256 shares = amount * _totalShare / totalStaked;

        // update storage
        _shares[staker] -= shares;
        _totalShare -= shares;

        // transfer tokens back to the staker
        _transfer(address(this), staker, amount);

        emit Unstake(staker, amount);
    }

    function _register(address trader, address referral) private {
        address[5] memory refs = _referrals[referral];
        _referrals[trader] = [referral, refs[0], refs[1], refs[2], refs[3]];
    }

    function _transferTaxed(address from, address to, uint256 amount, address trader) internal {
        address[5] memory refs = _referrals[trader];
        require(refs[4] != address(0), "Trading is not allowed (not registered)");

         {
            // calculate amounts to transfer as fees
            uint256 feeDeveloper = amount * FEE_DEVELOPER / 100;
            uint256 feeStakers = amount * FEE_STAKERS / 100;
            uint256[5] memory feeRefs = [
                amount * FEE_LEVEL_1 / 100,
                amount * FEE_LEVEL_2 / 100,
                amount * FEE_LEVEL_3 / 100,
                amount * FEE_LEVEL_4 / 100,
                amount * FEE_LEVEL_5 / 100
            ];

            // decrease balance of sender
            uint256 fromBalance = _balances[from];
            require(fromBalance >= amount, "ERC20: transfer amount exceeds balance");
            _balances[from] = fromBalance - amount;

            // transfer common part of the amount to the recipient
            uint256 feeTotal = feeDeveloper + feeStakers + feeRefs[0] + feeRefs[1] + feeRefs[2] + feeRefs[3] + feeRefs[4];
            uint256 remain = amount - feeTotal;
            _balances[to] += remain;
            emit Transfer(from, to, remain);

            // transfer fees to active referrals only
            for (uint8 i = 0; i < 5; i++) {
                uint256 feeRef = feeRefs[i];
                uint256 balance = _balances[refs[i]];
                if (balance >= REFERRAL_MIN_AMOUNT) {
                    // referral is active so transfer
                    _balances[refs[i]] = balance + feeRef;
                    emit Transfer(from, refs[i], feeRef);
                    emit RewardReferral(trader, refs[i], i+1, feeRef);
                } else {
                    // referral is not active at current time so transfer to stakers pool instead
                    feeStakers += feeRef;
                }
            }
            
            // transfer fees to the Developer
            _balances[_developer] += feeDeveloper;
            emit Transfer(from, _developer, feeDeveloper);
            
            // transfer fees to stakers pool
            _balances[address(this)] += feeStakers;
            emit Transfer(from, address(this), feeStakers);
            emit RewardStakers(trader, feeStakers);
        }
    }

    ////////////////////////////////////////////////////////////////
    // ERC20 related functions
    ////////////////////////////////////////////////////////////////

    /**
     * @dev Returns the name of the token.
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the symbol of the token, usually a shorter version of the
     * name.
     */
    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns the number of decimals used to get its user representation.
     * For example, if `decimals` equals `2`, a balance of `505` tokens should
     * be displayed to a user as `5.05` (`505 / 10 ** 2`).
     *
     * Tokens usually opt for a value of 18, imitating the relationship between
     * Ether and Wei. This is the value {ERC20} uses, unless this function is
     * overridden;
     *
     * NOTE: This information is only used for _display_ purposes: it in
     * no way affects any of the arithmetic of the contract, including
     * {IERC20-balanceOf} and {IERC20-transfer}.
     */
    function decimals() public view virtual override returns (uint8) {
        return 18;
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
    function balanceOf(address account) public view virtual override returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev See {IERC20-allowance}.
     */
    function allowance(address owner, address spender) public view virtual override returns (uint256) {
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
    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, amount);
        return true;
    }

    /**
     * @dev See {IERC20-transfer}.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - the caller must have a balance of at least `amount`.
     */
    function transfer(address to, uint256 amount) public virtual override returns (bool) {
        address owner = _msgSender();
        if (owner == _lpToken) {
            // TODO: Is it possible case?
            _transferTaxed(owner, to, amount, to);
        } else if (to == _lpToken) {
            // TODO: Is it possible case?
            _transferTaxed(owner, to, amount, owner);
        } else {
            _transfer(owner, to, amount);
        }
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
    function transferFrom(address from, address to, uint256 amount) public virtual override returns (bool) {
        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        if (from == _lpToken) {
            // TODO: Is it possible case?
            _transferTaxed(from, to, amount, to);
        } else if (to == _lpToken) {
            // TODO: Is it possible case?
            _transferTaxed(from, to, amount, from);
        } else {
            _transfer(from, to, amount);
        }
        return true;
    }

    /**
     * @dev Atomically increases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, allowance(owner, spender) + addedValue);
        return true;
    }

    /**
     * @dev Atomically decreases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     * - `spender` must have allowance for the caller of at least
     * `subtractedValue`.
     */
    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
        address owner = _msgSender();
        uint256 currentAllowance = allowance(owner, spender);
        require(currentAllowance >= subtractedValue, "ERC20: decreased allowance below zero");
         {
            _approve(owner, spender, currentAllowance - subtractedValue);
        }
        return true;
    }

    /**
     * @dev Moves `amount` of tokens from `from` to `to`.
     *
     * This internal function is equivalent to {transfer}, and can be used to
     * e.g. implement automatic token fees, slashing mechanisms, etc.
     *
     * Emits a {Transfer} event.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     */
    function _transfer(address from, address to, uint256 amount) internal virtual {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");

        uint256 fromBalance = _balances[from];
        require(fromBalance >= amount, "ERC20: transfer amount exceeds balance");
         {
            _balances[from] = fromBalance - amount;
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
         {
            // Overflow not possible: balance + amount is at most totalSupply + amount, which is checked above.
            _balances[account] += amount;
        }
        emit Transfer(address(0), account, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`, reducing the
     * total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function _burn(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: burn from the zero address");

        uint256 accountBalance = _balances[account];
        require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
         {
            _balances[account] = accountBalance - amount;
            // Overflow not possible: amount <= accountBalance <= totalSupply.
            _totalSupply -= amount;
        }

        emit Transfer(account, address(0), amount);
    }

    /**
     * @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
     *
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
    function _approve(address owner, address spender, uint256 amount) internal virtual {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Updates `owner` s allowance for `spender` based on spent `amount`.
     *
     * Does not update the allowance amount in case of infinite allowance.
     * Revert if not enough allowance is available.
     *
     * Might emit an {Approval} event.
     */
    function _spendAllowance(address owner, address spender, uint256 amount) internal virtual {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= amount, "ERC20: insufficient allowance");
             {
                _approve(owner, spender, currentAllowance - amount);
            }
        }
    }
}