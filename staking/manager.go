// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingManagerUnbondingRequest is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerUnbondingRequest struct {
	Transcoder common.Address
	Timestamp  *big.Int
	Amount     *big.Int
}

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transcoderApprovalPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_slashPoolAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Jailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"TranscoderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"readiness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnbondingRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Unjailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"_slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegateManaged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"managed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"getSlashableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTrancoderSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTranscoderState\",\"outputs\":[{\"internalType\":\"enumStakingManager.TranscoderState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"}],\"name\":\"getUnbondingRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.UnbondingRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"isJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"isManaged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingWithdrawalsExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"registerTranscoder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbondingManaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setApprovalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setSelfMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSlashFundAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setSlashRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"}],\"name\":\"setZone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transcoderApprovalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transcoders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"effectiveMinSelfStake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transcodersArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transcodersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"unjail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051614e6a380380614e6a83398181016040528101906100319190610457565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100999190610503565b60405180910390fd5b6100b1816101a560201b60201c565b508660025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600381905550846004819055508360058190555082600681905550816007819055508060085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101675f5f1b3361026660201b60201c565b506101987f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b083361026660201b60201c565b505050505050505061051c565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f610277838361035b60201b60201c565b610351576001805f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506102ee6103bf60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610355565b5f90505b92915050565b5f60015f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103f3826103ca565b9050919050565b610403816103e9565b811461040d575f5ffd5b50565b5f8151905061041e816103fa565b92915050565b5f819050919050565b61043681610424565b8114610440575f5ffd5b50565b5f815190506104518161042d565b92915050565b5f5f5f5f5f5f5f60e0888a031215610472576104716103c6565b5b5f61047f8a828b01610410565b97505060206104908a828b01610443565b96505060406104a18a828b01610443565b95505060606104b28a828b01610443565b94505060806104c38a828b01610443565b93505060a06104d48a828b01610443565b92505060c06104e58a828b01610410565b91505092959891949750929550565b6104fd816103e9565b82525050565b5f6020820190506105165f8301846104f4565b92915050565b614941806105295f395ff3fe608060405234801561000f575f5ffd5b50600436106102d4575f3560e01c806372f702f311610180578063c5f530af116100e7578063ec69a1bb116100a0578063f2fde38b1161007a578063f2fde38b14610909578063f3ae241514610925578063fa76015414610955578063fece707d14610973576102d4565b8063ec69a1bb1461088b578063ec810a0a146108bb578063ec87621c146108eb576102d4565b8063c5f530af146107ae578063d547741f146107cc578063e2dc17f6146107e8578063e341181d1461081f578063e4c6990b1461083d578063e71824bc1461085b576102d4565b806396fb4a9c1161013957806396fb4a9c146106ee5780639e79b1221461070a578063a217fddf1461073a578063a79e726314610758578063ac18de4314610774578063c57cc20014610790576102d4565b806372f702f3146106165780637edbceb1146106345780638d23fc611461063e5780638da5cb5b146106705780638efc97a11461068e57806391d14854146106be576102d4565b806336568abe1161023f578063503074ef116101f85780636c2227de116101d25780636c2227de146105a25780636cf6d675146105be5780636db28909146105dc578063715018a61461060c576102d4565b8063503074ef14610560578063523f87f11461057c5780635afd2faa14610598576102d4565b806336568abe146104a25780633939e608146104be578063399f57c0146104dc5780634151766b146104f8578063449ecfe6146105145780635028e2e114610530576102d4565b8063248a9ca311610291578063248a9ca3146103d2578063254124fa1461040257806326e348ba146104325780632c9f0f2e1461044e5780632d06177a1461046a5780632f2ff15d14610486576102d4565b806301ffc9a7146102d8578063026e402b14610308578063029859921461032457806314bfb527146103425780631e7ff8f614610372578063220bb14e146103a2575b5f5ffd5b6102f260048036038101906102ed91906138e1565b6109a3565b6040516102ff9190613926565b60405180910390f35b610322600480360381019061031d91906139cc565b610a1c565b005b61032c610abd565b6040516103399190613a19565b60405180910390f35b61035c60048036038101906103579190613a32565b610ac3565b6040516103699190613926565b60405180910390f35b61038c60048036038101906103879190613a32565b610bd1565b6040516103999190613a19565b60405180910390f35b6103bc60048036038101906103b79190613a32565b610c87565b6040516103c99190613926565b60405180910390f35b6103ec60048036038101906103e79190613a90565b610ce1565b6040516103f99190613aca565b60405180910390f35b61041c60048036038101906104179190613ae3565b610cfe565b6040516104299190613a19565b60405180910390f35b61044c60048036038101906104479190613b33565b610dec565b005b61046860048036038101906104639190613b33565b610e09565b005b610484600480360381019061047f9190613a32565b610e1b565b005b6104a0600480360381019061049b9190613b5e565b610e94565b005b6104bc60048036038101906104b79190613b5e565b610eb6565b005b6104c6610f31565b6040516104d39190613a19565b60405180910390f35b6104f660048036038101906104f19190613b33565b610f37565b005b610512600480360381019061050d9190613b33565b6110c9565b005b61052e60048036038101906105299190613a32565b6110db565b005b61054a60048036038101906105459190613b9c565b611239565b6040516105579190613a19565b60405180910390f35b61057a600480360381019061057591906139cc565b6113b2565b005b61059660048036038101906105919190613ae3565b61144e565b005b6105a06114e6565b005b6105bc60048036038101906105b79190613a32565b6115cb565b005b6105c6611952565b6040516105d39190613a19565b60405180910390f35b6105f660048036038101906105f191906139cc565b611958565b6040516106039190613a19565b60405180910390f35b6106146119fe565b005b61061e611a11565b60405161062b9190613c35565b60405180910390f35b61063c611a36565b005b61065860048036038101906106539190613a32565b611b29565b60405161066793929190613c4e565b60405180910390f35b610678611b5c565b6040516106859190613c92565b60405180910390f35b6106a860048036038101906106a39190613b9c565b611b83565b6040516106b59190613a19565b60405180910390f35b6106d860048036038101906106d39190613b5e565b611d0c565b6040516106e59190613926565b60405180910390f35b610708600480360381019061070391906139cc565b611d70565b005b610724600480360381019061071f91906139cc565b611e0c565b6040516107319190613d09565b60405180910390f35b610742611ee3565b60405161074f9190613aca565b60405180910390f35b610772600480360381019061076d9190613d5d565b611ee9565b005b61078e60048036038101906107899190613a32565b611fc3565b005b61079861203c565b6040516107a59190613926565b60405180910390f35b6107b66120e3565b6040516107c39190613a19565b60405180910390f35b6107e660048036038101906107e19190613b5e565b6120e9565b005b61080260048036038101906107fd9190613a32565b61210b565b604051610816989796959493929190613d88565b60405180910390f35b61082761215b565b6040516108349190613a19565b60405180910390f35b610845612161565b6040516108529190613c92565b60405180910390f35b61087560048036038101906108709190613a32565b612186565b6040516108829190613e77565b60405180910390f35b6108a560048036038101906108a09190613a32565b612449565b6040516108b29190613a19565b60405180910390f35b6108d560048036038101906108d09190613b33565b61254e565b6040516108e29190613c92565b60405180910390f35b6108f3612589565b6040516109009190613aca565b60405180910390f35b610923600480360381019061091e9190613a32565b6125ad565b005b61093f600480360381019061093a9190613a32565b612631565b60405161094c9190613926565b60405180910390f35b61095d612663565b60405161096a9190613a19565b60405180910390f35b61098d60048036038101906109889190613a32565b61266f565b60405161099a9190613a19565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a155750610a1482612681565b5b9050919050565b5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806005015f9054906101000a900460ff1615610aad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa490613eea565b60405180910390fd5b610ab88333846126ea565b505050565b60035481565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2990613f52565b60405180910390fd5b5f60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816001015411610bb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610baf90613fba565b60405180910390fd5b806008015f9054906101000a900460ff16915050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3790613f52565b60405180910390fd5b60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f01549050919050565b5f5f600a5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806005015f9054906101000a900460ff16915050919050565b5f60015f8381526020019081526020015f20600101549050919050565b5f610d0833612631565b610d47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3e90614022565b60405180910390fd5b5f600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806005015f9054906101000a900460ff16610dd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dce906140d6565b60405180910390fd5b610de2858585612ad0565b9150509392505050565b610df4613030565b5f8111610dff575f5ffd5b8060048190555050565b610e11613030565b8060058190555050565b610e23613030565b610e4d7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08826130b7565b508073ffffffffffffffffffffffffffffffffffffffff167f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a60405160405180910390a250565b610e9d82610ce1565b610ea6816131a0565b610eb083836130b7565b50505050565b610ebe6131b4565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610f22576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f2c82826131bb565b505050565b60055481565b60648110610f7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7190614164565b60405180910390fd5b5f3390505f60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816001015414611004576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ffb906141cc565b60405180910390fd5b4281600101819055508281600201819055506004548160090181905550600b82908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167f6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b60405160405180910390a2505050565b6110d1613030565b8060078190555050565b6110e3613030565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611151576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114890613f52565b60405180910390fd5b5f60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f8160010154116111d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ce90614234565b60405180910390fd5b5f816008015f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167ffa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff60405160405180910390a25050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036112a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161129f90613f52565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130d90613f52565b60405180910390fd5b5f6113218484611b83565b905080600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546113a9919061427f565b91505092915050565b6113ba613030565b5f60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816001015411611440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143790613fba565b60405180910390fd5b818160050181905550505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611478816131a0565b5f600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090506001816005015f6101000a81548160ff0219169083151502179055506114df8585856126ea565b5050505050565b5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508060040154816003015410611570576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611567906142fc565b60405180910390fd5b5f816003015490505b81600401548110156115c6575f61159082336132a5565b90508061159f575050506115c9565b600183600301546115b0919061431a565b8360030181905550508080600101915050611579565b50505b565b6115d3613030565b5f60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167890614397565b60405180910390fd5b5f8160010154116116c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116be90613fba565b60405180910390fd5b5f6116d183612186565b9050600160048111156116e7576116e6613e04565b5b8160048111156116fa576116f9613e04565b5b148061172a57506003600481111561171557611714613e04565b5b81600481111561172857611727613e04565b5b145b611769576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611760906143ff565b60405180910390fd5b5f6064600754845f015461177d919061441d565b611787919061448b565b905080835f015f82825461179b919061427f565b92505081905550826006016040518060400160405280428152602001600754815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015560208201518160010155505060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016118749291906144bb565b6020604051808303815f875af1158015611890573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118b4919061450c565b6118f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ea90614581565b60405180910390fd5b6118fc846134b6565b8373ffffffffffffffffffffffffffffffffffffffff167f4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd6007546040516119449190613a19565b60405180910390a250505050565b60065481565b5f5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806005015f9054906101000a900460ff16156119ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119e190614635565b60405180910390fd5b6119f5843385612ad0565b91505092915050565b611a06613030565b611a0f5f61359f565b565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508060040154816003015410611ac0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ab7906142fc565b60405180910390fd5b611ace8160030154336132a5565b611b0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b049061469d565b60405180910390fd5b60018160030154611b1e919061431a565b816003018190555050565b600a602052805f5260405f205f91509050806003015490806004015490806005015f9054906101000a900460ff16905083565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f5f60095f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f5f90505f826001015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f8190505b8460060180549050811015611cfe575f856006018281548110611c7657611c756146bb565b5b905f5260205f2090600202016001015490505f606482875f015f8d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054611cd5919061441d565b611cdf919061448b565b90508085611ced919061431a565b945050508080600101915050611c50565b508194505050505092915050565b5f60015f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b611d78613030565b5f60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816001015411611dfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611df590613fba565b60405180910390fd5b818160040181905550505050565b611e14613854565b5f600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806002015f8481526020019081526020015f206040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152505091505092915050565b5f5f1b81565b611ef1613030565b8073ffffffffffffffffffffffffffffffffffffffff1660085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611f80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f7790614732565b60405180910390fd5b8060085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611fcb613030565b611ff57f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08826131bb565b508073ffffffffffffffffffffffffffffffffffffffff167fef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd3160405160405180910390a250565b5f5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816003015490505b81600401548110156120da575f826002015f8381526020019081526020015f2090506006548160010154426120bb919061427f565b106120cc57600193505050506120e0565b508080600101915050612086565b505f9150505b90565b60045481565b6120f282610ce1565b6120fb816131a0565b61210583836131bb565b50505050565b6009602052805f5260405f205f91509050805f015490806001015490806002015490806003015490806004015490806005015490806008015f9054906101000a900460ff16908060090154905088565b60075481565b60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036121f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121ec90613f52565b60405180910390fd5b5f60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f82600101540361228b57600492505050612444565b816008015f9054906101000a900460ff16156122ac57600292505050612444565b6005548260010154426122bf919061427f565b1061243e578160090154815f015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541061231857600192505050612444565b5f5f90505f826003015490505b82600401548110156123dd575f836002015f8381526020019081526020015f2090508673ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146123a257506123d0565b6006548160010154426123b5919061427f565b10156123ce578060020154836123cb919061431a565b92505b505b8080600101915050612325565b50600454825f015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548261242b919061431a565b1061243c5760039350505050612444565b505b5f925050505b919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036124b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124af90613f52565b60405180910390fd5b5f60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f81600101541161253e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161253590613fba565b60405180910390fd5b8060060180549050915050919050565b600b818154811061255d575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b6125b5613030565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612625575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161261c9190613c92565b60405180910390fd5b61262e8161359f565b50565b5f61265c7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0883611d0c565b9050919050565b5f600b80549050905090565b5f61267a8283611239565b9050919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60095f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f600a5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16036127d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127cf90614397565b60405180910390fd5b60035483101561281d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128149061479a565b60405180910390fd5b5f815f015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403612910578160070184908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160060180549050816001015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b61291a8585613660565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8530866040518463ffffffff1660e01b8152600401612978939291906147b8565b6020604051808303815f875af1158015612994573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129b8919061450c565b6129f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ee90614581565b60405180910390fd5b82825f015f828254612a09919061431a565b9250508190555082815f015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612a5d919061431a565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b85604051612ac19190613a19565b60405180910390a35050505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603612b3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b3690614837565b60405180910390fd5b5f60095f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f600a5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050612bc98686613660565b805f015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054841115612c4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c419061489f565b60405180910390fd5b5f612c5487612186565b905084825f015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054612ca0919061427f565b825f015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555084835f0154612cf1919061427f565b835f01819055505f8260040154905060018360040154612d11919061431a565b83600401819055505f6004811115612d2c57612d2b613e04565b5b826004811115612d3f57612d3e613e04565b5b1480612d6f575060026004811115612d5a57612d59613e04565b5b826004811115612d6d57612d6c613e04565b5b145b80612da657508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614155b15612f0d578773ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16827f6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5428a604051612e0b9291906148bd565b60405180910390a460405180606001604052808973ffffffffffffffffffffffffffffffffffffffff16815260200160065442612e48919061427f565b815260200187815250836002015f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050612ec981886132a5565b612f08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612eff9061469d565b60405180910390fd5b613022565b8773ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16827f6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da560065442612f6b919061431a565b8a604051612f7a9291906148bd565b60405180910390a460405180606001604052808973ffffffffffffffffffffffffffffffffffffffff16815260200142815260200187815250836002015f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050505b809450505050509392505050565b6130386131b4565b73ffffffffffffffffffffffffffffffffffffffff16613056611b5c565b73ffffffffffffffffffffffffffffffffffffffff16146130b5576130796131b4565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016130ac9190613c92565b60405180910390fd5b565b5f6130c28383611d0c565b613196576001805f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506131336131b4565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001905061319a565b5f90505b92915050565b6131b1816131ac6131b4565b613803565b50565b5f33905090565b5f6131c68383611d0c565b1561329b575f60015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506132386131b4565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001905061329f565b5f90505b92915050565b5f5f600a5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816002015f8681526020019081526020015f2090505f816002015403613311575f925050506134b0565b600654816001015442613324919061427f565b1015613334575f925050506134b0565b5f816002015490505f826002018190555060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86836040518363ffffffff1660e01b81526004016133a19291906144bb565b6020604051808303815f875af11580156133bd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133e1919061450c565b613420576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161341790614581565b60405180910390fd5b815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16877f544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84846040516134a09190613a19565b60405180910390a4600193505050505b92915050565b5f60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f81600101541161353c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161353390613fba565b60405180910390fd5b6001816008015f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e760405160405180910390a25050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f600a5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f6136ab8484611b83565b905080825f015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546136fa919061427f565b9250508190555060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161377e9291906144bb565b6020604051808303815f875af115801561379a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137be919061450c565b6137fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137f490614581565b60405180910390fd5b50505050565b61380d8282611d0c565b6138505780826040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526004016138479291906148e4565b60405180910390fd5b5050565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6138c08161388c565b81146138ca575f5ffd5b50565b5f813590506138db816138b7565b92915050565b5f602082840312156138f6576138f5613888565b5b5f613903848285016138cd565b91505092915050565b5f8115159050919050565b6139208161390c565b82525050565b5f6020820190506139395f830184613917565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6139688261393f565b9050919050565b6139788161395e565b8114613982575f5ffd5b50565b5f813590506139938161396f565b92915050565b5f819050919050565b6139ab81613999565b81146139b5575f5ffd5b50565b5f813590506139c6816139a2565b92915050565b5f5f604083850312156139e2576139e1613888565b5b5f6139ef85828601613985565b9250506020613a00858286016139b8565b9150509250929050565b613a1381613999565b82525050565b5f602082019050613a2c5f830184613a0a565b92915050565b5f60208284031215613a4757613a46613888565b5b5f613a5484828501613985565b91505092915050565b5f819050919050565b613a6f81613a5d565b8114613a79575f5ffd5b50565b5f81359050613a8a81613a66565b92915050565b5f60208284031215613aa557613aa4613888565b5b5f613ab284828501613a7c565b91505092915050565b613ac481613a5d565b82525050565b5f602082019050613add5f830184613abb565b92915050565b5f5f5f60608486031215613afa57613af9613888565b5b5f613b0786828701613985565b9350506020613b1886828701613985565b9250506040613b29868287016139b8565b9150509250925092565b5f60208284031215613b4857613b47613888565b5b5f613b55848285016139b8565b91505092915050565b5f5f60408385031215613b7457613b73613888565b5b5f613b8185828601613a7c565b9250506020613b9285828601613985565b9150509250929050565b5f5f60408385031215613bb257613bb1613888565b5b5f613bbf85828601613985565b9250506020613bd085828601613985565b9150509250929050565b5f819050919050565b5f613bfd613bf8613bf38461393f565b613bda565b61393f565b9050919050565b5f613c0e82613be3565b9050919050565b5f613c1f82613c04565b9050919050565b613c2f81613c15565b82525050565b5f602082019050613c485f830184613c26565b92915050565b5f606082019050613c615f830186613a0a565b613c6e6020830185613a0a565b613c7b6040830184613917565b949350505050565b613c8c8161395e565b82525050565b5f602082019050613ca55f830184613c83565b92915050565b613cb48161395e565b82525050565b613cc381613999565b82525050565b606082015f820151613cdd5f850182613cab565b506020820151613cf06020850182613cba565b506040820151613d036040850182613cba565b50505050565b5f606082019050613d1c5f830184613cc9565b92915050565b5f613d2c8261393f565b9050919050565b613d3c81613d22565b8114613d46575f5ffd5b50565b5f81359050613d5781613d33565b92915050565b5f60208284031215613d7257613d71613888565b5b5f613d7f84828501613d49565b91505092915050565b5f61010082019050613d9c5f83018b613a0a565b613da9602083018a613a0a565b613db66040830189613a0a565b613dc36060830188613a0a565b613dd06080830187613a0a565b613ddd60a0830186613a0a565b613dea60c0830185613917565b613df760e0830184613a0a565b9998505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60058110613e4257613e41613e04565b5b50565b5f819050613e5282613e31565b919050565b5f613e6182613e45565b9050919050565b613e7181613e57565b82525050565b5f602082019050613e8a5f830184613e68565b92915050565b5f82825260208201905092915050565b7f44656c656761746f72206973206d616e616765640000000000000000000000005f82015250565b5f613ed4601483613e90565b9150613edf82613ea0565b602082019050919050565b5f6020820190508181035f830152613f0181613ec8565b9050919050565b7f63616e277420757365207a65726f2061646472657373000000000000000000005f82015250565b5f613f3c601683613e90565b9150613f4782613f08565b602082019050919050565b5f6020820190508181035f830152613f6981613f30565b9050919050565b7f5472616e73636f646572206e6f742072656769737465726564000000000000005f82015250565b5f613fa4601983613e90565b9150613faf82613f70565b602082019050919050565b5f6020820190508181035f830152613fd181613f98565b9050919050565b7f6e6f742061206d616e61676572000000000000000000000000000000000000005f82015250565b5f61400c600d83613e90565b915061401782613fd8565b602082019050919050565b5f6020820190508181035f83015261403981614000565b9050919050565b7f74686973206d6574686f642063616e206f6e6c792062652075736564206f6e6c5f8201527f7920666f722064656c656761746f722074686174206465706f7369746564204560208201527f5243323020746f6b656e73000000000000000000000000000000000000000000604082015250565b5f6140c0604b83613e90565b91506140cb82614040565b606082019050919050565b5f6020820190508181035f8301526140ed816140b4565b9050919050565b7f52617465206d75737420626520612070657263656e74616765206265747765655f8201527f6e203020616e6420313030000000000000000000000000000000000000000000602082015250565b5f61414e602b83613e90565b9150614159826140f4565b604082019050919050565b5f6020820190508181035f83015261417b81614142565b9050919050565b7f5472616e73636f64657220616c726561647920726567697374657265640000005f82015250565b5f6141b6601d83613e90565b91506141c182614182565b602082019050919050565b5f6020820190508181035f8301526141e3816141aa565b9050919050565b7f52656769737465726564207472616e73636f646572206f6e6c790000000000005f82015250565b5f61421e601a83613e90565b9150614229826141ea565b602082019050919050565b5f6020820190508181035f83015261424b81614212565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61428982613999565b915061429483613999565b92508282039050818111156142ac576142ab614252565b5b92915050565b7f6e6f2070656e64696e67207265717565737473000000000000000000000000005f82015250565b5f6142e6601383613e90565b91506142f1826142b2565b602082019050919050565b5f6020820190508181035f830152614313816142da565b9050919050565b5f61432482613999565b915061432f83613999565b925082820190508082111561434757614346614252565b5b92915050565b7f496e76616c6964207472616e73636f64657220616464726573730000000000005f82015250565b5f614381601a83613e90565b915061438c8261434d565b602082019050919050565b5f6020820190508181035f8301526143ae81614375565b9050919050565b7f43616e6e6f7420736c6173682074686973207472616e73636f646572000000005f82015250565b5f6143e9601c83613e90565b91506143f4826143b5565b602082019050919050565b5f6020820190508181035f830152614416816143dd565b9050919050565b5f61442782613999565b915061443283613999565b925082820261444081613999565b9150828204841483151761445757614456614252565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61449582613999565b91506144a083613999565b9250826144b0576144af61445e565b5b828204905092915050565b5f6040820190506144ce5f830185613c83565b6144db6020830184613a0a565b9392505050565b6144eb8161390c565b81146144f5575f5ffd5b50565b5f81519050614506816144e2565b92915050565b5f6020828403121561452157614520613888565b5b5f61452e848285016144f8565b91505092915050565b7f546f6b656e207472616e73666572206661696c656400000000000000000000005f82015250565b5f61456b601583613e90565b915061457682614537565b602082019050919050565b5f6020820190508181035f8301526145988161455f565b9050919050565b7f74686973206d6574686f642063616e277420626520757365642062792064656c5f8201527f656761746f722074686174206465706f736974656420455243323020746f6b6560208201527f6e73000000000000000000000000000000000000000000000000000000000000604082015250565b5f61461f604283613e90565b915061462a8261459f565b606082019050919050565b5f6020820190508181035f83015261464c81614613565b9050919050565b7f6661696c656420746f207769746864726177207374616b6500000000000000005f82015250565b5f614687601883613e90565b915061469282614653565b602082019050919050565b5f6020820190508181035f8301526146b48161467b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f416c72656164792073657420746f2074686973206164647265737300000000005f82015250565b5f61471c601b83613e90565b9150614727826146e8565b602082019050919050565b5f6020820190508181035f83015261474981614710565b9050919050565b7f416d6f756e742062656c6f77206d696e696d756d2064656c65676174696f6e005f82015250565b5f614784601f83613e90565b915061478f82614750565b602082019050919050565b5f6020820190508181035f8301526147b181614778565b9050919050565b5f6060820190506147cb5f830186613c83565b6147d86020830185613c83565b6147e56040830184613a0a565b949350505050565b7f43616e60742075736520616464726573732030783000000000000000000000005f82015250565b5f614821601583613e90565b915061482c826147ed565b602082019050919050565b5f6020820190508181035f83015261484e81614815565b9050919050565b7f4e6f7420656e6f7567682066756e6473000000000000000000000000000000005f82015250565b5f614889601083613e90565b915061489482614855565b602082019050919050565b5f6020820190508181035f8301526148b68161487d565b9050919050565b5f6040820190506148d05f830185613a0a565b6148dd6020830184613a0a565b9392505050565b5f6040820190506148f75f830185613c83565b6149046020830184613abb565b939250505056fea264697066735822122090f1fbef45f2623ae719257b2c458367c9dcaa3d79e685261d49d29114f42a7864736f6c634300081c0033",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingManagerMetaData.Bin instead.
var StakingManagerBin = StakingManagerMetaData.Bin

// DeployStakingManager deploys a new Ethereum contract, binding an instance of StakingManager to it.
func DeployStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, _stakingToken common.Address, _minDelegation *big.Int, _minSelfStake *big.Int, _transcoderApprovalPeriod *big.Int, _unbondingPeriod *big.Int, _slashRate *big.Int, _slashPoolAddress common.Address) (common.Address, *types.Transaction, *StakingManager, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingManagerBin), backend, _stakingToken, _minDelegation, _minSelfStake, _transcoderApprovalPeriod, _unbondingPeriod, _slashRate, _slashPoolAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) MANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.MANAGERROLE(&_StakingManager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) MANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.MANAGERROLE(&_StakingManager.CallOpts)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerCaller) Delegators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "delegators", arg0)

	outstruct := new(struct {
		Pending *big.Int
		Next    *big.Int
		Managed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pending = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Next = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Managed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerSession) Delegators(arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerCallerSession) Delegators(arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetDelegatorStake(opts *bind.CallOpts, transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getDelegatorStake", transcoderAddr, delegAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSelfStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getSelfStake", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSlashableAmount(opts *bind.CallOpts, transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getSlashableAmount", transcoderAddr, delegatorAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTotalStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getTotalStake", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTrancoderSlashes(opts *bind.CallOpts, transcoderAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getTrancoderSlashes", transcoderAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerCaller) GetTranscoderState(opts *bind.CallOpts, transcoderAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getTranscoderState", transcoderAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerCallerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns((address,uint256,uint256))
func (_StakingManager *StakingManagerCaller) GetUnbondingRequest(opts *bind.CallOpts, delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getUnbondingRequest", delegatorAddr, unbondingID)

	if err != nil {
		return *new(StakingManagerUnbondingRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingManagerUnbondingRequest)).(*StakingManagerUnbondingRequest)

	return out0, err

}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns((address,uint256,uint256))
func (_StakingManager *StakingManagerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns((address,uint256,uint256))
func (_StakingManager *StakingManagerCallerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsJailed(opts *bind.CallOpts, transcoderAddr common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "isJailed", transcoderAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsManaged(opts *bind.CallOpts, delegatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "isManaged", delegatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerSession) IsManaged(delegatorAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsManaged(&_StakingManager.CallOpts, delegatorAddr)
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsManaged(delegatorAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsManaged(&_StakingManager.CallOpts, delegatorAddr)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsManager(opts *bind.CallOpts, v common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "isManager", v)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerSession) IsManager(v common.Address) (bool, error) {
	return _StakingManager.Contract.IsManager(&_StakingManager.CallOpts, v)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsManager(v common.Address) (bool, error) {
	return _StakingManager.Contract.IsManager(&_StakingManager.CallOpts, v)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minSelfStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerCaller) PendingWithdrawalsExist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "pendingWithdrawalsExist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerSession) PendingWithdrawalsExist() (bool, error) {
	return _StakingManager.Contract.PendingWithdrawalsExist(&_StakingManager.CallOpts)
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) PendingWithdrawalsExist() (bool, error) {
	return _StakingManager.Contract.PendingWithdrawalsExist(&_StakingManager.CallOpts)
}

// SlashPoolAddress is a free data retrieval call binding the contract method 0xe4c6990b.
//
// Solidity: function slashPoolAddress() view returns(address)
func (_StakingManager *StakingManagerCaller) SlashPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "slashPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashPoolAddress is a free data retrieval call binding the contract method 0xe4c6990b.
//
// Solidity: function slashPoolAddress() view returns(address)
func (_StakingManager *StakingManagerSession) SlashPoolAddress() (common.Address, error) {
	return _StakingManager.Contract.SlashPoolAddress(&_StakingManager.CallOpts)
}

// SlashPoolAddress is a free data retrieval call binding the contract method 0xe4c6990b.
//
// Solidity: function slashPoolAddress() view returns(address)
func (_StakingManager *StakingManagerCallerSession) SlashPoolAddress() (common.Address, error) {
	return _StakingManager.Contract.SlashPoolAddress(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerCaller) SlashRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "slashRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakingManager *StakingManagerCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakingManager *StakingManagerSession) StakingToken() (common.Address, error) {
	return _StakingManager.Contract.StakingToken(&_StakingManager.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakingManager *StakingManagerCallerSession) StakingToken() (common.Address, error) {
	return _StakingManager.Contract.StakingToken(&_StakingManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscoderApprovalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "transcoderApprovalPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerCaller) Transcoders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "transcoders", arg0)

	outstruct := new(struct {
		Total                 *big.Int
		Timestamp             *big.Int
		RewardRate            *big.Int
		Rewards               *big.Int
		Zone                  *big.Int
		Capacity              *big.Int
		Jailed                bool
		EffectiveMinSelfStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Zone = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Capacity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Jailed = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.EffectiveMinSelfStake = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerSession) Transcoders(arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerCallerSession) Transcoders(arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerCaller) TranscodersArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "transcodersArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerSession) TranscodersArray(arg0 *big.Int) (common.Address, error) {
	return _StakingManager.Contract.TranscodersArray(&_StakingManager.CallOpts, arg0)
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerCallerSession) TranscodersArray(arg0 *big.Int) (common.Address, error) {
	return _StakingManager.Contract.TranscodersArray(&_StakingManager.CallOpts, arg0)
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscodersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "transcodersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerSession) TranscodersCount() (*big.Int, error) {
	return _StakingManager.Contract.TranscodersCount(&_StakingManager.CallOpts)
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscodersCount() (*big.Int, error) {
	return _StakingManager.Contract.TranscodersCount(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unbondingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x6c2227de.
//
// Solidity: function _slash(address addr) returns()
func (_StakingManager *StakingManagerTransactor) Slash(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "_slash", addr)
}

// Slash is a paid mutator transaction binding the contract method 0x6c2227de.
//
// Solidity: function _slash(address addr) returns()
func (_StakingManager *StakingManagerSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// Slash is a paid mutator transaction binding the contract method 0x6c2227de.
//
// Solidity: function _slash(address addr) returns()
func (_StakingManager *StakingManagerTransactorSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerTransactor) AddManager(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "addManager", v)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerSession) AddManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddManager(&_StakingManager.TransactOpts, v)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerTransactorSession) AddManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddManager(&_StakingManager.TransactOpts, v)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address transcoderAddr, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) Delegate(opts *bind.TransactOpts, transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegate", transcoderAddr, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address transcoderAddr, uint256 amount) returns()
func (_StakingManager *StakingManagerSession) Delegate(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address transcoderAddr, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) Delegate(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x523f87f1.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) DelegateManaged(opts *bind.TransactOpts, transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegateManaged", transcoderAddr, delegatorAddr, amount)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x523f87f1.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns()
func (_StakingManager *StakingManagerSession) DelegateManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DelegateManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x523f87f1.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) DelegateManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DelegateManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactor) RegisterTranscoder(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerTranscoder", rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerTransactor) RemoveManager(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "removeManager", v)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerSession) RemoveManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RemoveManager(&_StakingManager.TransactOpts, v)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerTransactorSession) RemoveManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RemoveManager(&_StakingManager.TransactOpts, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbonding(opts *bind.TransactOpts, transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbonding", transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbondingManaged(opts *bind.TransactOpts, transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbondingManaged", transcoderAddr, delegatorAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbondingManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbondingManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbondingManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbondingManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactor) SetApprovalPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setApprovalPeriod", period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactorSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactor) SetCapacity(opts *bind.TransactOpts, addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setCapacity", addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactorSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) SetSelfMinStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSelfMinStake", amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactor) SetSlashFundAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlashFundAddress", addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetSlashRate is a paid mutator transaction binding the contract method 0x4151766b.
//
// Solidity: function setSlashRate(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactor) SetSlashRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlashRate", rate)
}

// SetSlashRate is a paid mutator transaction binding the contract method 0x4151766b.
//
// Solidity: function setSlashRate(uint256 rate) returns()
func (_StakingManager *StakingManagerSession) SetSlashRate(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashRate(&_StakingManager.TransactOpts, rate)
}

// SetSlashRate is a paid mutator transaction binding the contract method 0x4151766b.
//
// Solidity: function setSlashRate(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlashRate(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashRate(&_StakingManager.TransactOpts, rate)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactor) SetZone(opts *bind.TransactOpts, addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setZone", addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactorSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactor) Unjail(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unjail", transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactorSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerTransactor) WithdrawAllPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawAllPending")
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerSession) WithdrawAllPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawAllPending(&_StakingManager.TransactOpts)
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawAllPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawAllPending(&_StakingManager.TransactOpts)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerTransactor) WithdrawPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawPending")
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerSession) WithdrawPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawPending(&_StakingManager.TransactOpts)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawPending(&_StakingManager.TransactOpts)
}

// StakingManagerDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the StakingManager contract.
type StakingManagerDelegatedIterator struct {
	Event *StakingManagerDelegated // Event containing the contract specifics and raw log

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
func (it *StakingManagerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerDelegated)
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
		it.Event = new(StakingManagerDelegated)
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
func (it *StakingManagerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerDelegated represents a Delegated event raised by the StakingManager contract.
type StakingManagerDelegated struct {
	Transcoder common.Address
	Delegator  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterDelegated(opts *bind.FilterOpts, transcoder []common.Address, delegator []common.Address) (*StakingManagerDelegatedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Delegated", transcoderRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerDelegatedIterator{contract: _StakingManager.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakingManagerDelegated, transcoder []common.Address, delegator []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Delegated", transcoderRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerDelegated)
				if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseDelegated(log types.Log) (*StakingManagerDelegated, error) {
	event := new(StakingManagerDelegated)
	if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerJailedIterator is returned from FilterJailed and is used to iterate over the raw logs and unpacked data for Jailed events raised by the StakingManager contract.
type StakingManagerJailedIterator struct {
	Event *StakingManagerJailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerJailed)
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
		it.Event = new(StakingManagerJailed)
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
func (it *StakingManagerJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerJailed represents a Jailed event raised by the StakingManager contract.
type StakingManagerJailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJailed is a free log retrieval operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterJailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerJailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerJailedIterator{contract: _StakingManager.contract, event: "Jailed", logs: logs, sub: sub}, nil
}

// WatchJailed is a free log subscription operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchJailed(opts *bind.WatchOpts, sink chan<- *StakingManagerJailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerJailed)
				if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
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

// ParseJailed is a log parse operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseJailed(log types.Log) (*StakingManagerJailed, error) {
	event := new(StakingManagerJailed)
	if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the StakingManager contract.
type StakingManagerManagerAddedIterator struct {
	Event *StakingManagerManagerAdded // Event containing the contract specifics and raw log

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
func (it *StakingManagerManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerManagerAdded)
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
		it.Event = new(StakingManagerManagerAdded)
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
func (it *StakingManagerManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerManagerAdded represents a ManagerAdded event raised by the StakingManager contract.
type StakingManagerManagerAdded struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) FilterManagerAdded(opts *bind.FilterOpts, manager []common.Address) (*StakingManagerManagerAddedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerManagerAddedIterator{contract: _StakingManager.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *StakingManagerManagerAdded, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerManagerAdded)
				if err := _StakingManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) ParseManagerAdded(log types.Log) (*StakingManagerManagerAdded, error) {
	event := new(StakingManagerManagerAdded)
	if err := _StakingManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the StakingManager contract.
type StakingManagerManagerRemovedIterator struct {
	Event *StakingManagerManagerRemoved // Event containing the contract specifics and raw log

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
func (it *StakingManagerManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerManagerRemoved)
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
		it.Event = new(StakingManagerManagerRemoved)
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
func (it *StakingManagerManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerManagerRemoved represents a ManagerRemoved event raised by the StakingManager contract.
type StakingManagerManagerRemoved struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) FilterManagerRemoved(opts *bind.FilterOpts, manager []common.Address) (*StakingManagerManagerRemovedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerManagerRemovedIterator{contract: _StakingManager.contract, event: "ManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *StakingManagerManagerRemoved, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerManagerRemoved)
				if err := _StakingManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
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

// ParseManagerRemoved is a log parse operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) ParseManagerRemoved(log types.Log) (*StakingManagerManagerRemoved, error) {
	event := new(StakingManagerManagerRemoved)
	if err := _StakingManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingManager contract.
type StakingManagerRoleAdminChangedIterator struct {
	Event *StakingManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleAdminChanged)
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
		it.Event = new(StakingManagerRoleAdminChanged)
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
func (it *StakingManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleAdminChanged represents a RoleAdminChanged event raised by the StakingManager contract.
type StakingManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleAdminChangedIterator{contract: _StakingManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleAdminChanged)
				if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakingManagerRoleAdminChanged, error) {
	event := new(StakingManagerRoleAdminChanged)
	if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingManager contract.
type StakingManagerRoleGrantedIterator struct {
	Event *StakingManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleGranted)
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
		it.Event = new(StakingManagerRoleGranted)
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
func (it *StakingManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleGranted represents a RoleGranted event raised by the StakingManager contract.
type StakingManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleGrantedIterator{contract: _StakingManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleGranted)
				if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) ParseRoleGranted(log types.Log) (*StakingManagerRoleGranted, error) {
	event := new(StakingManagerRoleGranted)
	if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingManager contract.
type StakingManagerRoleRevokedIterator struct {
	Event *StakingManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleRevoked)
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
		it.Event = new(StakingManagerRoleRevoked)
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
func (it *StakingManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleRevoked represents a RoleRevoked event raised by the StakingManager contract.
type StakingManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleRevokedIterator{contract: _StakingManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleRevoked)
				if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) ParseRoleRevoked(log types.Log) (*StakingManagerRoleRevoked, error) {
	event := new(StakingManagerRoleRevoked)
	if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the StakingManager contract.
type StakingManagerSlashedIterator struct {
	Event *StakingManagerSlashed // Event containing the contract specifics and raw log

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
func (it *StakingManagerSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerSlashed)
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
		it.Event = new(StakingManagerSlashed)
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
func (it *StakingManagerSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerSlashed represents a Slashed event raised by the StakingManager contract.
type StakingManagerSlashed struct {
	Transcoder common.Address
	Rate       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 rate)
func (_StakingManager *StakingManagerFilterer) FilterSlashed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerSlashedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Slashed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerSlashedIterator{contract: _StakingManager.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 rate)
func (_StakingManager *StakingManagerFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingManagerSlashed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Slashed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerSlashed)
				if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 rate)
func (_StakingManager *StakingManagerFilterer) ParseSlashed(log types.Log) (*StakingManagerSlashed, error) {
	event := new(StakingManagerSlashed)
	if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeWithdrawalIterator is returned from FilterStakeWithdrawal and is used to iterate over the raw logs and unpacked data for StakeWithdrawal events raised by the StakingManager contract.
type StakingManagerStakeWithdrawalIterator struct {
	Event *StakingManagerStakeWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeWithdrawal)
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
		it.Event = new(StakingManagerStakeWithdrawal)
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
func (it *StakingManagerStakeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeWithdrawal represents a StakeWithdrawal event raised by the StakingManager contract.
type StakingManagerStakeWithdrawal struct {
	UnbondingID *big.Int
	Delegator   common.Address
	Transcoder  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawal is a free log retrieval operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterStakeWithdrawal(opts *bind.FilterOpts, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (*StakingManagerStakeWithdrawalIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeWithdrawal", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeWithdrawalIterator{contract: _StakingManager.contract, event: "StakeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawal is a free log subscription operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchStakeWithdrawal(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeWithdrawal, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeWithdrawal", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeWithdrawal)
				if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
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

// ParseStakeWithdrawal is a log parse operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseStakeWithdrawal(log types.Log) (*StakingManagerStakeWithdrawal, error) {
	event := new(StakingManagerStakeWithdrawal)
	if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerTranscoderRegisteredIterator is returned from FilterTranscoderRegistered and is used to iterate over the raw logs and unpacked data for TranscoderRegistered events raised by the StakingManager contract.
type StakingManagerTranscoderRegisteredIterator struct {
	Event *StakingManagerTranscoderRegistered // Event containing the contract specifics and raw log

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
func (it *StakingManagerTranscoderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerTranscoderRegistered)
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
		it.Event = new(StakingManagerTranscoderRegistered)
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
func (it *StakingManagerTranscoderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerTranscoderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerTranscoderRegistered represents a TranscoderRegistered event raised by the StakingManager contract.
type StakingManagerTranscoderRegistered struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTranscoderRegistered is a free log retrieval operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterTranscoderRegistered(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerTranscoderRegisteredIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTranscoderRegisteredIterator{contract: _StakingManager.contract, event: "TranscoderRegistered", logs: logs, sub: sub}, nil
}

// WatchTranscoderRegistered is a free log subscription operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchTranscoderRegistered(opts *bind.WatchOpts, sink chan<- *StakingManagerTranscoderRegistered, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerTranscoderRegistered)
				if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
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

// ParseTranscoderRegistered is a log parse operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseTranscoderRegistered(log types.Log) (*StakingManagerTranscoderRegistered, error) {
	event := new(StakingManagerTranscoderRegistered)
	if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnbondingRequestedIterator is returned from FilterUnbondingRequested and is used to iterate over the raw logs and unpacked data for UnbondingRequested events raised by the StakingManager contract.
type StakingManagerUnbondingRequestedIterator struct {
	Event *StakingManagerUnbondingRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnbondingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnbondingRequested)
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
		it.Event = new(StakingManagerUnbondingRequested)
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
func (it *StakingManagerUnbondingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnbondingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnbondingRequested represents a UnbondingRequested event raised by the StakingManager contract.
type StakingManagerUnbondingRequested struct {
	UnbondingID *big.Int
	Delegator   common.Address
	Transcoder  common.Address
	Readiness   *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbondingRequested is a free log retrieval operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterUnbondingRequested(opts *bind.FilterOpts, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (*StakingManagerUnbondingRequestedIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnbondingRequested", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnbondingRequestedIterator{contract: _StakingManager.contract, event: "UnbondingRequested", logs: logs, sub: sub}, nil
}

// WatchUnbondingRequested is a free log subscription operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchUnbondingRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnbondingRequested, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnbondingRequested", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnbondingRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
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

// ParseUnbondingRequested is a log parse operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseUnbondingRequested(log types.Log) (*StakingManagerUnbondingRequested, error) {
	event := new(StakingManagerUnbondingRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnjailedIterator is returned from FilterUnjailed and is used to iterate over the raw logs and unpacked data for Unjailed events raised by the StakingManager contract.
type StakingManagerUnjailedIterator struct {
	Event *StakingManagerUnjailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnjailed)
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
		it.Event = new(StakingManagerUnjailed)
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
func (it *StakingManagerUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnjailed represents a Unjailed event raised by the StakingManager contract.
type StakingManagerUnjailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnjailed is a free log retrieval operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterUnjailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerUnjailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnjailedIterator{contract: _StakingManager.contract, event: "Unjailed", logs: logs, sub: sub}, nil
}

// WatchUnjailed is a free log subscription operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchUnjailed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnjailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnjailed)
				if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
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

// ParseUnjailed is a log parse operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseUnjailed(log types.Log) (*StakingManagerUnjailed, error) {
	event := new(StakingManagerUnjailed)
	if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
