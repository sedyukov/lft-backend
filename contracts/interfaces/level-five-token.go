// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"developer\",\"type\":\"address\"},{\"internalType\":\"address[5]\",\"name\":\"refs\",\"type\":\"address[5]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardReferral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardStakers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"acceptInvite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"referrals\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lpt\",\"type\":\"address\"}],\"name\":\"setLPToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"share\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"staked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d8a38038062001d8a833981810160405260c08110156200003757600080fd5b50805190602001620000526200004c62000244565b62000248565b6001600160a01b038216620000ae576040805162461bcd60e51b815260206004820181905260248201527f446576656c6f7065722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6040805180820190915260108082526f2632bb32b6102334bb32902a37b5b2b760811b6020909201918252620000e7916001916200055f565b506040805180820190915260038082526213119560ea1b602090920191825262000114916002916200055f565b50600680546001600160a01b0384166001600160a01b0319918216179091556007805490911673997af3d9295df06b511ff7121cf9d3ef6f65e74917905560006200015e62000244565b905062000178816b033b2e3c9fd0803ce800000062000298565b62000196818360005b6020020151683635c9adc5dea0000062000342565b620001a48183600162000181565b620001b28183600262000181565b620001c08183600362000181565b620001ce8183600462000181565b620001db81600062000484565b8151620001e9908262000484565b602082015162000202908360005b602002015162000484565b60408201516200021590836001620001f7565b60608201516200022890836002620001f7565b60808201516200023b90836003620001f7565b50505062000656565b3390565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038216620002f4576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b60058054820190556001600160a01b03821660008181526003602090815260408083208054860190558051858152905160008051602062001d45833981519152929181900390910190a35050565b6001600160a01b038316620003895760405162461bcd60e51b815260040180806020018281038252602581526020018062001d656025913960400191505060405180910390fd5b6001600160a01b038216620003d05760405162461bcd60e51b815260040180806020018281038252602381526020018062001cfc6023913960400191505060405180910390fd5b6001600160a01b038316600090815260036020526040902054818110156200042a5760405162461bcd60e51b815260040180806020018281038252602681526020018062001d1f6026913960400191505060405180910390fd5b6001600160a01b03808516600081815260036020908152604080832087870390559387168083529184902080548701905583518681529351919360008051602062001d45833981519152929081900390910190a350505050565b6001600160a01b038116600090815260086020526040808220815160a08101928390529160059082845b81546001600160a01b03168152600190910190602001808311620004ae57505050505090506040518060a00160405280836001600160a01b03166001600160a01b03168152602001826000600581106200050457fe5b602090810291909101516001600160a01b0390811683528482015181168383015260408086015182168185015260608087015183169401939093528616600090815260089091522062000559916005620005f4565b50505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620005975760008555620005e2565b82601f10620005b257805160ff1916838001178555620005e2565b82800160010185558215620005e2579182015b82811115620005e2578251825591602001919060010190620005c5565b50620005f09291506200063f565b5090565b8260058101928215620005e2579160200282015b82811115620005e257825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000608565b5b80821115620005f0576000815560010162000640565b61169680620006666000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806370a08231116100c3578063a457c2d71161007c578063a457c2d714610402578063a694fc3a1461042e578063a9059cbb1461044b578063dbdb556314610477578063dd62ed3e1461049d578063f2fde38b146104cb5761014d565b806370a0823114610340578063715018a6146103665780638da5cb5b1461036e57806395d89b411461037657806398807d841461037e5780639ca423b3146103a45761014d565b806323b872dd1161011557806323b872dd146102575780632e17de781461028d578063313ce567146102ac57806339509351146102ca5780635fcbd285146102f65780636a678a9c1461031a5761014d565b8063026c42071461015257806306fdde031461016c578063095ea7b3146101e957806318160ddd146102295780631877bb5c14610231575b600080fd5b61015a6104f1565b60408051918252519081900360200190f35b6101746104f7565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ae578181015183820152602001610196565b50505050905090810190601f1680156101db5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610215600480360360408110156101ff57600080fd5b506001600160a01b03813516906020013561058c565b604080519115158252519081900360200190f35b61015a6105ae565b61015a6004803603602081101561024757600080fd5b50356001600160a01b03166105b4565b6102156004803603606081101561026d57600080fd5b506001600160a01b038135811691602081013590911690604001356105d3565b6102aa600480360360208110156102a357600080fd5b503561064a565b005b6102b46106f5565b6040805160ff9092168252519081900360200190f35b610215600480360360408110156102e057600080fd5b506001600160a01b0381351690602001356106fa565b6102fe61071d565b604080516001600160a01b039092168252519081900360200190f35b6102aa6004803603602081101561033057600080fd5b50356001600160a01b031661072c565b61015a6004803603602081101561035657600080fd5b50356001600160a01b0316610756565b6102aa610771565b6102fe610785565b610174610794565b61015a6004803603602081101561039457600080fd5b50356001600160a01b03166107f2565b6103ca600480360360208110156103ba57600080fd5b50356001600160a01b031661084b565b604051808260a080838360005b838110156103ef5781810151838201526020016103d7565b5050505090500191505060405180910390f35b6102156004803603604081101561041857600080fd5b506001600160a01b0381351690602001356108a7565b6102aa6004803603602081101561044457600080fd5b503561090e565b6102156004803603604081101561046157600080fd5b506001600160a01b0381351690602001356109c9565b6102aa6004803603602081101561048d57600080fd5b50356001600160a01b0316610a2b565b61015a600480360360408110156104b357600080fd5b506001600160a01b0381358116916020013516610b4e565b6102aa600480360360208110156104e157600080fd5b50356001600160a01b0316610b79565b600a5490565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156105825780601f1061055757610100808354040283529160200191610582565b820191906000526020600020905b81548152906001019060200180831161056557829003601f168201915b5050505050905090565b600080610597610bd2565b90506105a4818585610bd6565b5060019392505050565b60055490565b6001600160a01b0381166000908152600960205260409020545b919050565b6000806105de610bd2565b90506105eb858285610cc2565b6007546001600160a01b03868116911614156106125761060d85858587610d41565b61063f565b6007546001600160a01b03858116911614156106345761060d85858588610d41565b61063f85858561119c565b506001949350505050565b6000610654610bd2565b30600090815260036020526040812054600a5492935091829085028161067657fe5b6001600160a01b0385166000908152600960205260409020805492909104918290039055600a8054829003905590506106b030848661119c565b6040805185815290516001600160a01b038516917f85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd919081900360200190a250505050565b601290565b600080610705610bd2565b90506105a48185856107178589610b4e565b01610bd6565b6007546001600160a01b031690565b6107346112d7565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b031660009081526003602052604090205490565b6107796112d7565b610783600061134b565b565b6000546001600160a01b031690565b60028054604080516020601f60001961010060018716150201909416859004938401819004810282018101909252828152606093909290918301828280156105825780601f1061055757610100808354040283529160200191610582565b3060009081526003602052604081205480158061080f5750600a54155b1561081e5760009150506105ce565b600a546001600160a01b03841660009081526009602052604090205482028161084357fe5b049392505050565b61085361146c565b6001600160a01b03821660009081526008602052604090819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161087e5750505050509050919050565b6000806108b2610bd2565b905060006108c08286610b4e565b9050838110156109015760405162461bcd60e51b815260040180806020018281038252602581526020018061163c6025913960400191505060405180910390fd5b61063f8286868403610bd6565b6000610918610bd2565b3060009081526003602052604090205490915082811580159061093d57506000600a54115b156109535781600a5485028161094f57fe5b0490505b6001600160a01b0383166000908152600960205260409020805482019055600a80548201905561098483308661119c565b6040805185815290516001600160a01b038516917febedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a919081900360200190a250505050565b6000806109d4610bd2565b6007549091506001600160a01b03808316911614156109fe576109f981858587610d41565b6105a4565b6007546001600160a01b0385811691161415610a20576109f981858584610d41565b6105a481858561119c565b6000610a35610bd2565b6001600160a01b038181166000908152600860205260409020600401549192501615610a9d576040805162461bcd60e51b81526020600482015260126024820152715265676973746572656420616c726561647960701b604482015290519081900360640190fd5b6001600160a01b0382166000908152600860205260408120600401546001600160a01b03161415610aff5760405162461bcd60e51b81526004018080602001828103825260238152602001806116196023913960400191505060405180910390fd5b610b09818361139b565b604080516001600160a01b0383811682529151918416917f98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a3515199181900360200190a25050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205490565b610b816112d7565b6001600160a01b038116610bc65760405162461bcd60e51b81526004018080602001828103825260268152602001806115426026913960400191505060405180910390fd5b610bcf8161134b565b50565b3390565b6001600160a01b038316610c1b5760405162461bcd60e51b81526004018080602001828103825260248152602001806115f56024913960400191505060405180910390fd5b6001600160a01b038216610c605760405162461bcd60e51b81526004018080602001828103825260228152602001806115686022913960400191505060405180910390fd5b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6000610cce8484610b4e565b90506000198114610d3b5781811015610d2e576040805162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015290519081900360640190fd5b610d3b8484848403610bd6565b50505050565b6001600160a01b038116600090815260086020526040808220815160a08101928390529160059082845b81546001600160a01b03168152600190910190602001808311610d6b575050505050905060006001600160a01b031681600460058110610da757fe5b60200201516001600160a01b03161415610df25760405162461bcd60e51b81526004018080602001828103825260278152602001806114f86027913960400191505060405180910390fd5b6040805160a0810182526064808604808352602080840182905283850182905260608401829052600380890284900460808601526001600160a01b038b166000908152915293909320546002870291909104919086811015610e855760405162461bcd60e51b815260040180806020018281038252602681526020018061158a6026913960400191505060405180910390fd5b6001600160a01b0389166000908152600360205260408120888303905582600460200201518360036020020151846002602002015185600160200201518660006020020151888a01010101010190506000818903905080600360008c6001600160a01b03166001600160a01b0316815260200190815260200160002060008282540192505081905550896001600160a01b03168b6001600160a01b03166000805160206115b0833981519152836040518082815260200191505060405180910390a360005b60058160ff1610156110bf576000858260ff1660058110610f6757fe5b602002015190506000600360008b8560ff1660058110610f8357fe5b60200201516001600160a01b03166001600160a01b03168152602001908152602001600020549050683635c9adc5dea0000081106110b057818101600360008c8660ff1660058110610fd157fe5b60200201516001600160a01b03166001600160a01b0316815260200190815260200160002081905550898360ff166005811061100957fe5b60200201516001600160a01b03168e6001600160a01b03166000805160206115b0833981519152846040518082815260200191505060405180910390a38260010160ff168a8460ff166005811061105c57fe5b60200201516001600160a01b03168c6001600160a01b03167f465e6d39f39851c67c8c607ade83d797e24fb74952d23e916b51472d756a5fc7856040518082815260200191505060405180910390a46110b5565b968101965b5050600101610f4a565b50600680546001600160a01b0390811660009081526003602090815260409182902080548b019055925481518a8152915190831693928f16926000805160206115b0833981519152928290030190a330600081815260036020908152604091829020805489019055815188815291516001600160a01b038f16926000805160206115b083398151915292908290030190a36040805186815290516001600160a01b038a16917fa1710d49c174fe37d3ab453bf039e0b530c5d2b4be719b29c9d9250dc3ad73b3919081900360200190a25050505050505050505050565b6001600160a01b0383166111e15760405162461bcd60e51b81526004018080602001828103825260258152602001806115d06025913960400191505060405180910390fd5b6001600160a01b0382166112265760405162461bcd60e51b815260040180806020018281038252602381526020018061151f6023913960400191505060405180910390fd5b6001600160a01b0383166000908152600360205260409020548181101561127e5760405162461bcd60e51b815260040180806020018281038252602681526020018061158a6026913960400191505060405180910390fd5b6001600160a01b0380851660008181526003602090815260408083208787039055938716808352918490208054870190558351868152935191936000805160206115b0833981519152929081900390910190a350505050565b6112df610bd2565b6001600160a01b03166112f0610785565b6001600160a01b031614610783576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116600090815260086020526040808220815160a08101928390529160059082845b81546001600160a01b031681526001909101906020018083116113c557505050505090506040518060a00160405280836001600160a01b03166001600160a01b031681526020018260006005811061141957fe5b602090810291909101516001600160a01b03908116835284820151811683830152604080860151821681850152606080870151831694019390935286166000908152600890915220610d3b91600561148a565b6040518060a001604052806005906020820280368337509192915050565b82600581019282156114d2579160200282015b828111156114d257825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061149d565b506114de9291506114e2565b5090565b5b808211156114de57600081556001016114e356fe54726164696e67206973206e6f7420616c6c6f77656420286e6f7420726567697374657265642945524332303a207472616e7366657220746f20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737350726f766964656420726566657272616c206973206e6f74207265676973746572656445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122088baaa95f7237b6a946147c26695738a9fc921a708b7b4152409942ecfd0fc8664736f6c6343000706003345524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a207472616e736665722066726f6d20746865207a65726f2061646472657373",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, developer common.Address, refs [5]common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, developer, refs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contract *ContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCallerSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Contract *ContractCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Contract *ContractSession) LpToken() (common.Address, error) {
	return _Contract.Contract.LpToken(&_Contract.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Contract *ContractCallerSession) LpToken() (common.Address, error) {
	return _Contract.Contract.LpToken(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Referrals is a free data retrieval call binding the contract method 0x9ca423b3.
//
// Solidity: function referrals(address trader) view returns(address[5])
func (_Contract *ContractCaller) Referrals(opts *bind.CallOpts, trader common.Address) ([5]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "referrals", trader)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// Referrals is a free data retrieval call binding the contract method 0x9ca423b3.
//
// Solidity: function referrals(address trader) view returns(address[5])
func (_Contract *ContractSession) Referrals(trader common.Address) ([5]common.Address, error) {
	return _Contract.Contract.Referrals(&_Contract.CallOpts, trader)
}

// Referrals is a free data retrieval call binding the contract method 0x9ca423b3.
//
// Solidity: function referrals(address trader) view returns(address[5])
func (_Contract *ContractCallerSession) Referrals(trader common.Address) ([5]common.Address, error) {
	return _Contract.Contract.Referrals(&_Contract.CallOpts, trader)
}

// Share is a free data retrieval call binding the contract method 0x1877bb5c.
//
// Solidity: function share(address trader) view returns(uint256)
func (_Contract *ContractCaller) Share(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "share", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Share is a free data retrieval call binding the contract method 0x1877bb5c.
//
// Solidity: function share(address trader) view returns(uint256)
func (_Contract *ContractSession) Share(trader common.Address) (*big.Int, error) {
	return _Contract.Contract.Share(&_Contract.CallOpts, trader)
}

// Share is a free data retrieval call binding the contract method 0x1877bb5c.
//
// Solidity: function share(address trader) view returns(uint256)
func (_Contract *ContractCallerSession) Share(trader common.Address) (*big.Int, error) {
	return _Contract.Contract.Share(&_Contract.CallOpts, trader)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address staker) view returns(uint256)
func (_Contract *ContractCaller) Staked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "staked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address staker) view returns(uint256)
func (_Contract *ContractSession) Staked(staker common.Address) (*big.Int, error) {
	return _Contract.Contract.Staked(&_Contract.CallOpts, staker)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address staker) view returns(uint256)
func (_Contract *ContractCallerSession) Staked(staker common.Address) (*big.Int, error) {
	return _Contract.Contract.Staked(&_Contract.CallOpts, staker)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_Contract *ContractCaller) TotalShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_Contract *ContractSession) TotalShare() (*big.Int, error) {
	return _Contract.Contract.TotalShare(&_Contract.CallOpts)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_Contract *ContractCallerSession) TotalShare() (*big.Int, error) {
	return _Contract.Contract.TotalShare(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// AcceptInvite is a paid mutator transaction binding the contract method 0xdbdb5563.
//
// Solidity: function acceptInvite(address referral) returns()
func (_Contract *ContractTransactor) AcceptInvite(opts *bind.TransactOpts, referral common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "acceptInvite", referral)
}

// AcceptInvite is a paid mutator transaction binding the contract method 0xdbdb5563.
//
// Solidity: function acceptInvite(address referral) returns()
func (_Contract *ContractSession) AcceptInvite(referral common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AcceptInvite(&_Contract.TransactOpts, referral)
}

// AcceptInvite is a paid mutator transaction binding the contract method 0xdbdb5563.
//
// Solidity: function acceptInvite(address referral) returns()
func (_Contract *ContractTransactorSession) AcceptInvite(referral common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AcceptInvite(&_Contract.TransactOpts, referral)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contract *ContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contract *ContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contract *ContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DecreaseAllowance(&_Contract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contract *ContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DecreaseAllowance(&_Contract.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contract *ContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contract *ContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contract *ContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetLPToken is a paid mutator transaction binding the contract method 0x6a678a9c.
//
// Solidity: function setLPToken(address lpt) returns()
func (_Contract *ContractTransactor) SetLPToken(opts *bind.TransactOpts, lpt common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLPToken", lpt)
}

// SetLPToken is a paid mutator transaction binding the contract method 0x6a678a9c.
//
// Solidity: function setLPToken(address lpt) returns()
func (_Contract *ContractSession) SetLPToken(lpt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetLPToken(&_Contract.TransactOpts, lpt)
}

// SetLPToken is a paid mutator transaction binding the contract method 0x6a678a9c.
//
// Solidity: function setLPToken(address lpt) returns()
func (_Contract *ContractTransactorSession) SetLPToken(lpt common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetLPToken(&_Contract.TransactOpts, lpt)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Contract *ContractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Contract *ContractSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Stake(&_Contract.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Stake(&_Contract.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contract *ContractSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_Contract *ContractTransactor) Unstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unstake", amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_Contract *ContractSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Unstake(&_Contract.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Unstake(&_Contract.TransactOpts, amount)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the Contract contract.
type ContractRegisterIterator struct {
	Event *ContractRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegister represents a Register event raised by the Contract contract.
type ContractRegister struct {
	Referral common.Address
	Trader   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0x98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a351519.
//
// Solidity: event Register(address indexed referral, address trader)
func (_Contract *ContractFilterer) FilterRegister(opts *bind.FilterOpts, referral []common.Address) (*ContractRegisterIterator, error) {

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Register", referralRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegisterIterator{contract: _Contract.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0x98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a351519.
//
// Solidity: event Register(address indexed referral, address trader)
func (_Contract *ContractFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *ContractRegister, referral []common.Address) (event.Subscription, error) {

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Register", referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegister)
				if err := _Contract.contract.UnpackLog(event, "Register", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegister is a log parse operation binding the contract event 0x98ada70a1cb506dc4591465e1ee9be3fd7a2b6c73ecf3b949009718c9a351519.
//
// Solidity: event Register(address indexed referral, address trader)
func (_Contract *ContractFilterer) ParseRegister(log types.Log) (*ContractRegister, error) {
	event := new(ContractRegister)
	if err := _Contract.contract.UnpackLog(event, "Register", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardReferralIterator is returned from FilterRewardReferral and is used to iterate over the raw logs and unpacked data for RewardReferral events raised by the Contract contract.
type ContractRewardReferralIterator struct {
	Event *ContractRewardReferral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRewardReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardReferral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRewardReferral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRewardReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardReferral represents a RewardReferral event raised by the Contract contract.
type ContractRewardReferral struct {
	Trader   common.Address
	Referral common.Address
	Level    uint8
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardReferral is a free log retrieval operation binding the contract event 0x465e6d39f39851c67c8c607ade83d797e24fb74952d23e916b51472d756a5fc7.
//
// Solidity: event RewardReferral(address indexed trader, address indexed referral, uint8 indexed level, uint256 amount)
func (_Contract *ContractFilterer) FilterRewardReferral(opts *bind.FilterOpts, trader []common.Address, referral []common.Address, level []uint8) (*ContractRewardReferralIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}
	var levelRule []interface{}
	for _, levelItem := range level {
		levelRule = append(levelRule, levelItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardReferral", traderRule, referralRule, levelRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardReferralIterator{contract: _Contract.contract, event: "RewardReferral", logs: logs, sub: sub}, nil
}

// WatchRewardReferral is a free log subscription operation binding the contract event 0x465e6d39f39851c67c8c607ade83d797e24fb74952d23e916b51472d756a5fc7.
//
// Solidity: event RewardReferral(address indexed trader, address indexed referral, uint8 indexed level, uint256 amount)
func (_Contract *ContractFilterer) WatchRewardReferral(opts *bind.WatchOpts, sink chan<- *ContractRewardReferral, trader []common.Address, referral []common.Address, level []uint8) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}
	var levelRule []interface{}
	for _, levelItem := range level {
		levelRule = append(levelRule, levelItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardReferral", traderRule, referralRule, levelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardReferral)
				if err := _Contract.contract.UnpackLog(event, "RewardReferral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardReferral is a log parse operation binding the contract event 0x465e6d39f39851c67c8c607ade83d797e24fb74952d23e916b51472d756a5fc7.
//
// Solidity: event RewardReferral(address indexed trader, address indexed referral, uint8 indexed level, uint256 amount)
func (_Contract *ContractFilterer) ParseRewardReferral(log types.Log) (*ContractRewardReferral, error) {
	event := new(ContractRewardReferral)
	if err := _Contract.contract.UnpackLog(event, "RewardReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardStakersIterator is returned from FilterRewardStakers and is used to iterate over the raw logs and unpacked data for RewardStakers events raised by the Contract contract.
type ContractRewardStakersIterator struct {
	Event *ContractRewardStakers // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRewardStakersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardStakers)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRewardStakers)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRewardStakersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardStakersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardStakers represents a RewardStakers event raised by the Contract contract.
type ContractRewardStakers struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardStakers is a free log retrieval operation binding the contract event 0xa1710d49c174fe37d3ab453bf039e0b530c5d2b4be719b29c9d9250dc3ad73b3.
//
// Solidity: event RewardStakers(address indexed trader, uint256 amount)
func (_Contract *ContractFilterer) FilterRewardStakers(opts *bind.FilterOpts, trader []common.Address) (*ContractRewardStakersIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardStakers", traderRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardStakersIterator{contract: _Contract.contract, event: "RewardStakers", logs: logs, sub: sub}, nil
}

// WatchRewardStakers is a free log subscription operation binding the contract event 0xa1710d49c174fe37d3ab453bf039e0b530c5d2b4be719b29c9d9250dc3ad73b3.
//
// Solidity: event RewardStakers(address indexed trader, uint256 amount)
func (_Contract *ContractFilterer) WatchRewardStakers(opts *bind.WatchOpts, sink chan<- *ContractRewardStakers, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardStakers", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardStakers)
				if err := _Contract.contract.UnpackLog(event, "RewardStakers", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardStakers is a log parse operation binding the contract event 0xa1710d49c174fe37d3ab453bf039e0b530c5d2b4be719b29c9d9250dc3ad73b3.
//
// Solidity: event RewardStakers(address indexed trader, uint256 amount)
func (_Contract *ContractFilterer) ParseRewardStakers(log types.Log) (*ContractRewardStakers, error) {
	event := new(ContractRewardStakers)
	if err := _Contract.contract.UnpackLog(event, "RewardStakers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Contract contract.
type ContractStakeIterator struct {
	Event *ContractStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStake represents a Stake event raised by the Contract contract.
type ContractStake struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed staker, uint256 amount)
func (_Contract *ContractFilterer) FilterStake(opts *bind.FilterOpts, staker []common.Address) (*ContractStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Stake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ContractStakeIterator{contract: _Contract.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed staker, uint256 amount)
func (_Contract *ContractFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *ContractStake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Stake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStake)
				if err := _Contract.contract.UnpackLog(event, "Stake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStake is a log parse operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed staker, uint256 amount)
func (_Contract *ContractFilterer) ParseStake(log types.Log) (*ContractStake, error) {
	event := new(ContractStake)
	if err := _Contract.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the Contract contract.
type ContractUnstakeIterator struct {
	Event *ContractUnstake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnstake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnstake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnstake represents a Unstake event raised by the Contract contract.
type ContractUnstake struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed staker, uint256 amount)
func (_Contract *ContractFilterer) FilterUnstake(opts *bind.FilterOpts, staker []common.Address) (*ContractUnstakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unstake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnstakeIterator{contract: _Contract.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed staker, uint256 amount)
func (_Contract *ContractFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *ContractUnstake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unstake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnstake)
				if err := _Contract.contract.UnpackLog(event, "Unstake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstake is a log parse operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed staker, uint256 amount)
func (_Contract *ContractFilterer) ParseUnstake(log types.Log) (*ContractUnstake, error) {
	event := new(ContractUnstake)
	if err := _Contract.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
