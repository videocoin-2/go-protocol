// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package streams

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

// StreamMetaData contains all meta data concerning the Stream contract.
var StreamMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_client\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"profiles\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"_paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"ChunkProofScrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"ChunkProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"ChunkProofValidated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"wattage\",\"type\":\"uint256[]\"}],\"name\":\"addInputChunkId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"client\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endStream\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isChunk\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"}],\"name\":\"isProfileTranscoded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTranscodingDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outStreams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"required\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatedChunks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"scrapProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proof\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outChunkId\",\"type\":\"uint256\"}],\"name\":\"submitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profile\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"validateProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wattages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516122e73803806122e783398181016040528101906100319190610499565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361009f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161009690610573565b60405180910390fd5b836001819055508260025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816008908051906020019061017b92919061020a565b505f5f90505b600880549050811015610200575f600882815481106101a3576101a2610591565b5b905f5260205f2001549050600160075f8381526020019081526020015f205f015f6101000a81548160ff0219169083151502179055508160075f8381526020019081526020015f2060010181905550508080600101915050610181565b50505050506105be565b828054828255905f5260205f20908101928215610244579160200282015b82811115610243578251825591602001919060010190610228565b5b5090506102519190610255565b5090565b5b8082111561026c575f815f905550600101610256565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61029381610281565b811461029d575f5ffd5b50565b5f815190506102ae8161028a565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102dd826102b4565b9050919050565b6102ed816102d3565b81146102f7575f5ffd5b50565b5f81519050610308816102e4565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61035882610312565b810181811067ffffffffffffffff8211171561037757610376610322565b5b80604052505050565b5f610389610270565b9050610395828261034f565b919050565b5f67ffffffffffffffff8211156103b4576103b3610322565b5b602082029050602081019050919050565b5f5ffd5b5f6103db6103d68461039a565b610380565b905080838252602082019050602084028301858111156103fe576103fd6103c5565b5b835b81811015610427578061041388826102a0565b845260208401935050602081019050610400565b5050509392505050565b5f82601f8301126104455761044461030e565b5b81516104558482602086016103c9565b91505092915050565b5f610468826102d3565b9050919050565b6104788161045e565b8114610482575f5ffd5b50565b5f815190506104938161046f565b92915050565b5f5f5f5f608085870312156104b1576104b0610279565b5b5f6104be878288016102a0565b94505060206104cf878288016102fa565b935050604085015167ffffffffffffffff8111156104f0576104ef61027d565b5b6104fc87828801610431565b925050606061050d87828801610485565b91505092959194509250565b5f82825260208201905092915050565b7f496e76616c696420636c69656e742061646472657373000000000000000000005f82015250565b5f61055d601683610519565b915061056882610529565b602082019050919050565b5f6020820190508181035f83015261058a81610551565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b611d1c806105cb5f395ff3fe608060405234801561000f575f5ffd5b50600436106100f3575f3560e01c8063747f758911610095578063c5d0b14c11610064578063c5d0b14c1461024b578063d78e647f1461027d578063eda0ce1714610299578063fc1028bc146102c9576100f3565b8063747f7589146101eb578063af640d0f14610207578063bb57a5ed14610225578063bbe58b0c1461022f576100f3565b80631f54e5bd116100d15780631f54e5bd146101635780633013ce2914610193578063356939ab146101b1578063481c6a75146101cd576100f3565b8063109e94cf146100f757806312fa6feb146101155780631bb62fc414610133575b5f5ffd5b6100ff6102e7565b60405161010c9190611268565b60405180910390f35b61011d61030c565b60405161012a919061129b565b60405180910390f35b61014d600480360381019061014891906112f8565b61031f565b60405161015a919061129b565b60405180910390f35b61017d600480360381019061017891906112f8565b61033c565b60405161018a919061129b565b60405180910390f35b61019b6103b4565b6040516101a8919061137e565b60405180910390f35b6101cb60048036038101906101c691906114e7565b6103d9565b005b6101d5610557565b6040516101e29190611268565b60405180910390f35b61020560048036038101906102009190611541565b61057b565b005b61020f6107b3565b60405161021c91906115b4565b60405180910390f35b61022d6107b9565b005b610249600480360381019061024491906115cd565b6108b4565b005b610265600480360381019061026091906112f8565b610b1e565b6040516102749392919061160b565b60405180910390f35b610297600480360381019061029291906115cd565b610b50565b005b6102b360048036038101906102ae91906115cd565b61113d565b6040516102c091906115b4565b60405180910390f35b6102d1611168565b6040516102de919061129b565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260149054906101000a900460ff1681565b6004602052805f5260405f205f915054906101000a900460ff1681565b5f5f60075f8481526020019081526020015f209050805f015f9054906101000a900460ff166103a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103979061169a565b60405180910390fd5b600680549050816002015414915050919050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610467576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045e90611702565b60405180910390fd5b60045f8381526020019081526020015f205f9054906101000a900460ff1615801561049f5750600260149054906101000a900460ff16155b6104de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d59061176a565b60405180910390fd5b600160045f8481526020019081526020015f205f6101000a81548160ff0219169083151502179055508060055f8481526020019081526020015f20908051906020019061052c9291906111c3565b50600682908060018154018082558091505060019003905f5260205f20015f90919091909150555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60075f8681526020019081526020015f206003015f8581526020019081526020015f20905060045f8581526020019081526020015f205f9054906101000a900460ff16801561061957505f73ffffffffffffffffffffffffffffffffffffffff16816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610658576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064f906117f8565b60405180910390fd5b60075f8681526020019081526020015f205f015f9054906101000a900460ff166106b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ae9061169a565b60405180910390fd5b805f0160405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815260200185815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015550506001815f018054905061077e9190611843565b85857fb625ad3a7ea2d8a035a88ccb4001b583e72535a7093232080cc461ff6925917b60405160405180910390a45050505050565b60015481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610847576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083e90611702565b60405180910390fd5b600260149054906101000a900460ff1615610897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088e906118c0565b60405180910390fd5b6001600260146101000a81548160ff021916908315150217905550565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b815260040161090d9190611268565b602060405180830381865afa158015610928573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094c9190611908565b61098b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109829061197d565b60405180910390fd5b5f60075f8481526020019081526020015f206003015f8381526020019081526020015f20905060045f8381526020019081526020015f205f9054906101000a900460ff168015610a2957505f73ffffffffffffffffffffffffffffffffffffffff16816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610a68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5f906117f8565b60405180910390fd5b805f018054905060018260010154610a80919061199b565b1115610ac1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab890611a18565b60405180910390fd5b806001015f815480929190610ad590611a36565b919050555060018160010154610aeb9190611843565b82847fb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a60405160405180910390a4505050565b6007602052805f5260405f205f91509050805f015f9054906101000a900460ff16908060010154908060020154905083565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663facd743b336040518263ffffffff1660e01b8152600401610ba99190611268565b602060405180830381865afa158015610bc4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be89190611908565b610c27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1e9061197d565b60405180910390fd5b5f60075f8481526020019081526020015f209050805f015f9054906101000a900460ff16610c8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c819061169a565b60405180910390fd5b5f816003015f8481526020019081526020015f20905060045f8481526020019081526020015f205f9054906101000a900460ff168015610d1857505f73ffffffffffffffffffffffffffffffffffffffff16816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610d57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4e906117f8565b60405180910390fd5b805f0180549050816001015410610da3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9a90611ac7565b60405180910390fd5b5f815f01826001015481548110610dbd57610dbc611ae5565b5b905f5260205f20906003020190505f60055f8681526020019081526020015f20846001015481548110610df357610df2611ae5565b5b905f5260205f20015490505f60645f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323be9d276040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e6a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e8e9190611b26565b83610e999190611b51565b610ea39190611bbf565b90508082610eb19190611843565b915060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb845f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401610f31929190611bef565b6020604051808303815f875af1158015610f4d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f719190611908565b610fb0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa790611c60565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161102c929190611bef565b6020604051808303815f875af1158015611048573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106c9190611908565b6110ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a290611cc8565b60405180910390fd5b33846002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846002015f81548092919061110190611a36565b919050555085877f6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e3760405160405180910390a350505050505050565b6005602052815f5260405f208181548110611156575f80fd5b905f5260205f20015f91509150505481565b5f5f5f90505b6008805490508110156111ba576111a06008828154811061119257611191611ae5565b5b905f5260205f20015461033c565b6111ad575f9150506111c0565b808060010191505061116e565b50600190505b90565b828054828255905f5260205f209081019282156111fd579160200282015b828111156111fc5782518255916020019190600101906111e1565b5b50905061120a919061120e565b5090565b5b80821115611225575f815f90555060010161120f565b5090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61125282611229565b9050919050565b61126281611248565b82525050565b5f60208201905061127b5f830184611259565b92915050565b5f8115159050919050565b61129581611281565b82525050565b5f6020820190506112ae5f83018461128c565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6112d7816112c5565b81146112e1575f5ffd5b50565b5f813590506112f2816112ce565b92915050565b5f6020828403121561130d5761130c6112bd565b5b5f61131a848285016112e4565b91505092915050565b5f819050919050565b5f61134661134161133c84611229565b611323565b611229565b9050919050565b5f6113578261132c565b9050919050565b5f6113688261134d565b9050919050565b6113788161135e565b82525050565b5f6020820190506113915f83018461136f565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6113e18261139b565b810181811067ffffffffffffffff82111715611400576113ff6113ab565b5b80604052505050565b5f6114126112b4565b905061141e82826113d8565b919050565b5f67ffffffffffffffff82111561143d5761143c6113ab565b5b602082029050602081019050919050565b5f5ffd5b5f61146461145f84611423565b611409565b905080838252602082019050602084028301858111156114875761148661144e565b5b835b818110156114b0578061149c88826112e4565b845260208401935050602081019050611489565b5050509392505050565b5f82601f8301126114ce576114cd611397565b5b81356114de848260208601611452565b91505092915050565b5f5f604083850312156114fd576114fc6112bd565b5b5f61150a858286016112e4565b925050602083013567ffffffffffffffff81111561152b5761152a6112c1565b5b611537858286016114ba565b9150509250929050565b5f5f5f5f60808587031215611559576115586112bd565b5b5f611566878288016112e4565b9450506020611577878288016112e4565b9350506040611588878288016112e4565b9250506060611599878288016112e4565b91505092959194509250565b6115ae816112c5565b82525050565b5f6020820190506115c75f8301846115a5565b92915050565b5f5f604083850312156115e3576115e26112bd565b5b5f6115f0858286016112e4565b9250506020611601858286016112e4565b9150509250929050565b5f60608201905061161e5f83018661128c565b61162b60208301856115a5565b61163860408301846115a5565b949350505050565b5f82825260208201905092915050565b7f50726f66696c65206e6f742072657175697265640000000000000000000000005f82015250565b5f611684601483611640565b915061168f82611650565b602082019050919050565b5f6020820190508181035f8301526116b181611678565b9050919050565b7f4e6f742061206d616e61676572000000000000000000000000000000000000005f82015250565b5f6116ec600d83611640565b91506116f7826116b8565b602082019050919050565b5f6020820190508181035f830152611719816116e0565b9050919050565b7f496e76616c6964206368756e6b206f722073747265616d20656e6465640000005f82015250565b5f611754601d83611640565b915061175f82611720565b602082019050919050565b5f6020820190508181035f83015261178181611748565b9050919050565b7f496e76616c6964206368756e6b206f7220616c72656164792076616c696461745f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f6117e2602283611640565b91506117ed82611788565b604082019050919050565b5f6020820190508181035f83015261180f816117d6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61184d826112c5565b9150611858836112c5565b92508282039050818111156118705761186f611816565b5b92915050565b7f53747265616d20616c726561647920656e6465640000000000000000000000005f82015250565b5f6118aa601483611640565b91506118b582611876565b602082019050919050565b5f6020820190508181035f8301526118d78161189e565b9050919050565b6118e781611281565b81146118f1575f5ffd5b50565b5f81519050611902816118de565b92915050565b5f6020828403121561191d5761191c6112bd565b5b5f61192a848285016118f4565b91505092915050565b7f4e6f7420612076616c696461746f7200000000000000000000000000000000005f82015250565b5f611967600f83611640565b915061197282611933565b602082019050919050565b5f6020820190508181035f8301526119948161195b565b9050919050565b5f6119a5826112c5565b91506119b0836112c5565b92508282019050808211156119c8576119c7611816565b5b92915050565b7f4e6f2070726f6f667320746f20736372617000000000000000000000000000005f82015250565b5f611a02601283611640565b9150611a0d826119ce565b602082019050919050565b5f6020820190508181035f830152611a2f816119f6565b9050919050565b5f611a40826112c5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a7257611a71611816565b5b600182019050919050565b7f4e6f2070726f6f667320617661696c61626c65000000000000000000000000005f82015250565b5f611ab1601383611640565b9150611abc82611a7d565b602082019050919050565b5f6020820190508181035f830152611ade81611aa5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050611b20816112ce565b92915050565b5f60208284031215611b3b57611b3a6112bd565b5b5f611b4884828501611b12565b91505092915050565b5f611b5b826112c5565b9150611b66836112c5565b9250828202611b74816112c5565b91508282048414831517611b8b57611b8a611816565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611bc9826112c5565b9150611bd4836112c5565b925082611be457611be3611b92565b5b828204905092915050565b5f604082019050611c025f830185611259565b611c0f60208301846115a5565b9392505050565b7f4d696e6572207061796d656e74206661696c65640000000000000000000000005f82015250565b5f611c4a601483611640565b9150611c5582611c16565b602082019050919050565b5f6020820190508181035f830152611c7781611c3e565b9050919050565b7f5365727669636520666565207061796d656e74206661696c65640000000000005f82015250565b5f611cb2601a83611640565b9150611cbd82611c7e565b602082019050919050565b5f6020820190508181035f830152611cdf81611ca6565b905091905056fea2646970667358221220dda308ed18d851b2cca7106c293394b8a99e427a0b04854f4c7a53cd5238516a64736f6c634300081c0033",
}

// StreamABI is the input ABI used to generate the binding from.
// Deprecated: Use StreamMetaData.ABI instead.
var StreamABI = StreamMetaData.ABI

// StreamBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StreamMetaData.Bin instead.
var StreamBin = StreamMetaData.Bin

// DeployStream deploys a new Ethereum contract, binding an instance of Stream to it.
func DeployStream(auth *bind.TransactOpts, backend bind.ContractBackend, _id *big.Int, _client common.Address, profiles []*big.Int, _paymentToken common.Address) (common.Address, *types.Transaction, *Stream, error) {
	parsed, err := StreamMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StreamBin), backend, _id, _client, profiles, _paymentToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stream{StreamCaller: StreamCaller{contract: contract}, StreamTransactor: StreamTransactor{contract: contract}, StreamFilterer: StreamFilterer{contract: contract}}, nil
}

// Stream is an auto generated Go binding around an Ethereum contract.
type Stream struct {
	StreamCaller     // Read-only binding to the contract
	StreamTransactor // Write-only binding to the contract
	StreamFilterer   // Log filterer for contract events
}

// StreamCaller is an auto generated read-only Go binding around an Ethereum contract.
type StreamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StreamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StreamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StreamSession struct {
	Contract     *Stream           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StreamCallerSession struct {
	Contract *StreamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StreamTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StreamTransactorSession struct {
	Contract     *StreamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamRaw is an auto generated low-level Go binding around an Ethereum contract.
type StreamRaw struct {
	Contract *Stream // Generic contract binding to access the raw methods on
}

// StreamCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StreamCallerRaw struct {
	Contract *StreamCaller // Generic read-only contract binding to access the raw methods on
}

// StreamTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StreamTransactorRaw struct {
	Contract *StreamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStream creates a new instance of Stream, bound to a specific deployed contract.
func NewStream(address common.Address, backend bind.ContractBackend) (*Stream, error) {
	contract, err := bindStream(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stream{StreamCaller: StreamCaller{contract: contract}, StreamTransactor: StreamTransactor{contract: contract}, StreamFilterer: StreamFilterer{contract: contract}}, nil
}

// NewStreamCaller creates a new read-only instance of Stream, bound to a specific deployed contract.
func NewStreamCaller(address common.Address, caller bind.ContractCaller) (*StreamCaller, error) {
	contract, err := bindStream(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StreamCaller{contract: contract}, nil
}

// NewStreamTransactor creates a new write-only instance of Stream, bound to a specific deployed contract.
func NewStreamTransactor(address common.Address, transactor bind.ContractTransactor) (*StreamTransactor, error) {
	contract, err := bindStream(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StreamTransactor{contract: contract}, nil
}

// NewStreamFilterer creates a new log filterer instance of Stream, bound to a specific deployed contract.
func NewStreamFilterer(address common.Address, filterer bind.ContractFilterer) (*StreamFilterer, error) {
	contract, err := bindStream(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StreamFilterer{contract: contract}, nil
}

// bindStream binds a generic wrapper to an already deployed contract.
func bindStream(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StreamMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stream *StreamRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stream.Contract.StreamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stream *StreamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.Contract.StreamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stream *StreamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stream.Contract.StreamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stream *StreamCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stream.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stream *StreamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stream *StreamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stream.Contract.contract.Transact(opts, method, params...)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Stream *StreamCaller) Client(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "client")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Stream *StreamSession) Client() (common.Address, error) {
	return _Stream.Contract.Client(&_Stream.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Stream *StreamCallerSession) Client() (common.Address, error) {
	return _Stream.Contract.Client(&_Stream.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Stream *StreamCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Stream *StreamSession) Ended() (bool, error) {
	return _Stream.Contract.Ended(&_Stream.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Stream *StreamCallerSession) Ended() (bool, error) {
	return _Stream.Contract.Ended(&_Stream.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Stream *StreamCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Stream *StreamSession) Id() (*big.Int, error) {
	return _Stream.Contract.Id(&_Stream.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Stream *StreamCallerSession) Id() (*big.Int, error) {
	return _Stream.Contract.Id(&_Stream.CallOpts)
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) view returns(bool)
func (_Stream *StreamCaller) IsChunk(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "isChunk", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) view returns(bool)
func (_Stream *StreamSession) IsChunk(arg0 *big.Int) (bool, error) {
	return _Stream.Contract.IsChunk(&_Stream.CallOpts, arg0)
}

// IsChunk is a free data retrieval call binding the contract method 0x1bb62fc4.
//
// Solidity: function isChunk(uint256 ) view returns(bool)
func (_Stream *StreamCallerSession) IsChunk(arg0 *big.Int) (bool, error) {
	return _Stream.Contract.IsChunk(&_Stream.CallOpts, arg0)
}

// IsProfileTranscoded is a free data retrieval call binding the contract method 0x1f54e5bd.
//
// Solidity: function isProfileTranscoded(uint256 profile) view returns(bool)
func (_Stream *StreamCaller) IsProfileTranscoded(opts *bind.CallOpts, profile *big.Int) (bool, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "isProfileTranscoded", profile)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProfileTranscoded is a free data retrieval call binding the contract method 0x1f54e5bd.
//
// Solidity: function isProfileTranscoded(uint256 profile) view returns(bool)
func (_Stream *StreamSession) IsProfileTranscoded(profile *big.Int) (bool, error) {
	return _Stream.Contract.IsProfileTranscoded(&_Stream.CallOpts, profile)
}

// IsProfileTranscoded is a free data retrieval call binding the contract method 0x1f54e5bd.
//
// Solidity: function isProfileTranscoded(uint256 profile) view returns(bool)
func (_Stream *StreamCallerSession) IsProfileTranscoded(profile *big.Int) (bool, error) {
	return _Stream.Contract.IsProfileTranscoded(&_Stream.CallOpts, profile)
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() view returns(bool)
func (_Stream *StreamCaller) IsTranscodingDone(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "isTranscodingDone")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() view returns(bool)
func (_Stream *StreamSession) IsTranscodingDone() (bool, error) {
	return _Stream.Contract.IsTranscodingDone(&_Stream.CallOpts)
}

// IsTranscodingDone is a free data retrieval call binding the contract method 0xfc1028bc.
//
// Solidity: function isTranscodingDone() view returns(bool)
func (_Stream *StreamCallerSession) IsTranscodingDone() (bool, error) {
	return _Stream.Contract.IsTranscodingDone(&_Stream.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Stream *StreamCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Stream *StreamSession) Manager() (common.Address, error) {
	return _Stream.Contract.Manager(&_Stream.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Stream *StreamCallerSession) Manager() (common.Address, error) {
	return _Stream.Contract.Manager(&_Stream.CallOpts)
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) view returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamCaller) OutStreams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "outStreams", arg0)

	outstruct := new(struct {
		Required        bool
		Index           *big.Int
		ValidatedChunks *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Required = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValidatedChunks = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) view returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamSession) OutStreams(arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	return _Stream.Contract.OutStreams(&_Stream.CallOpts, arg0)
}

// OutStreams is a free data retrieval call binding the contract method 0xc5d0b14c.
//
// Solidity: function outStreams(uint256 ) view returns(bool required, uint256 index, uint256 validatedChunks)
func (_Stream *StreamCallerSession) OutStreams(arg0 *big.Int) (struct {
	Required        bool
	Index           *big.Int
	ValidatedChunks *big.Int
}, error) {
	return _Stream.Contract.OutStreams(&_Stream.CallOpts, arg0)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Stream *StreamCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Stream *StreamSession) PaymentToken() (common.Address, error) {
	return _Stream.Contract.PaymentToken(&_Stream.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Stream *StreamCallerSession) PaymentToken() (common.Address, error) {
	return _Stream.Contract.PaymentToken(&_Stream.CallOpts)
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) view returns(uint256)
func (_Stream *StreamCaller) Wattages(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stream.contract.Call(opts, &out, "wattages", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) view returns(uint256)
func (_Stream *StreamSession) Wattages(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Stream.Contract.Wattages(&_Stream.CallOpts, arg0, arg1)
}

// Wattages is a free data retrieval call binding the contract method 0xeda0ce17.
//
// Solidity: function wattages(uint256 , uint256 ) view returns(uint256)
func (_Stream *StreamCallerSession) Wattages(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Stream.Contract.Wattages(&_Stream.CallOpts, arg0, arg1)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamTransactor) AddInputChunkId(opts *bind.TransactOpts, chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "addInputChunkId", chunkId, wattage)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamSession) AddInputChunkId(chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.Contract.AddInputChunkId(&_Stream.TransactOpts, chunkId, wattage)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x356939ab.
//
// Solidity: function addInputChunkId(uint256 chunkId, uint256[] wattage) returns()
func (_Stream *StreamTransactorSession) AddInputChunkId(chunkId *big.Int, wattage []*big.Int) (*types.Transaction, error) {
	return _Stream.Contract.AddInputChunkId(&_Stream.TransactOpts, chunkId, wattage)
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamTransactor) EndStream(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "endStream")
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamSession) EndStream() (*types.Transaction, error) {
	return _Stream.Contract.EndStream(&_Stream.TransactOpts)
}

// EndStream is a paid mutator transaction binding the contract method 0xbb57a5ed.
//
// Solidity: function endStream() returns()
func (_Stream *StreamTransactorSession) EndStream() (*types.Transaction, error) {
	return _Stream.Contract.EndStream(&_Stream.TransactOpts)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactor) ScrapProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "scrapProof", profile, chunkId)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamSession) ScrapProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ScrapProof(&_Stream.TransactOpts, profile, chunkId)
}

// ScrapProof is a paid mutator transaction binding the contract method 0xbbe58b0c.
//
// Solidity: function scrapProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactorSession) ScrapProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ScrapProof(&_Stream.TransactOpts, profile, chunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamTransactor) SubmitProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "submitProof", profile, chunkId, proof, outChunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamSession) SubmitProof(profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.SubmitProof(&_Stream.TransactOpts, profile, chunkId, proof, outChunkId)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x747f7589.
//
// Solidity: function submitProof(uint256 profile, uint256 chunkId, uint256 proof, uint256 outChunkId) returns()
func (_Stream *StreamTransactorSession) SubmitProof(profile *big.Int, chunkId *big.Int, proof *big.Int, outChunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.SubmitProof(&_Stream.TransactOpts, profile, chunkId, proof, outChunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactor) ValidateProof(opts *bind.TransactOpts, profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.contract.Transact(opts, "validateProof", profile, chunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamSession) ValidateProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ValidateProof(&_Stream.TransactOpts, profile, chunkId)
}

// ValidateProof is a paid mutator transaction binding the contract method 0xd78e647f.
//
// Solidity: function validateProof(uint256 profile, uint256 chunkId) returns()
func (_Stream *StreamTransactorSession) ValidateProof(profile *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _Stream.Contract.ValidateProof(&_Stream.TransactOpts, profile, chunkId)
}

// StreamChunkProofScrappedIterator is returned from FilterChunkProofScrapped and is used to iterate over the raw logs and unpacked data for ChunkProofScrapped events raised by the Stream contract.
type StreamChunkProofScrappedIterator struct {
	Event *StreamChunkProofScrapped // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofScrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofScrapped)
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
		it.Event = new(StreamChunkProofScrapped)
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
func (it *StreamChunkProofScrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofScrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofScrapped represents a ChunkProofScrapped event raised by the Stream contract.
type StreamChunkProofScrapped struct {
	Profile *big.Int
	ChunkId *big.Int
	Idx     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofScrapped is a free log retrieval operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) FilterChunkProofScrapped(opts *bind.FilterOpts, profile []*big.Int, chunkId []*big.Int, idx []*big.Int) (*StreamChunkProofScrappedIterator, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofScrapped", profileRule, chunkIdRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofScrappedIterator{contract: _Stream.contract, event: "ChunkProofScrapped", logs: logs, sub: sub}, nil
}

// WatchChunkProofScrapped is a free log subscription operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) WatchChunkProofScrapped(opts *bind.WatchOpts, sink chan<- *StreamChunkProofScrapped, profile []*big.Int, chunkId []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofScrapped", profileRule, chunkIdRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofScrapped)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofScrapped", log); err != nil {
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

// ParseChunkProofScrapped is a log parse operation binding the contract event 0xb25faf3bb93b5105c8bc184c6af20f251543b1fdd93cd5c6f9f31b13c138b91a.
//
// Solidity: event ChunkProofScrapped(uint256 indexed profile, uint256 indexed chunkId, uint256 indexed idx)
func (_Stream *StreamFilterer) ParseChunkProofScrapped(log types.Log) (*StreamChunkProofScrapped, error) {
	event := new(StreamChunkProofScrapped)
	if err := _Stream.contract.UnpackLog(event, "ChunkProofScrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StreamChunkProofSubmittedIterator is returned from FilterChunkProofSubmitted and is used to iterate over the raw logs and unpacked data for ChunkProofSubmitted events raised by the Stream contract.
type StreamChunkProofSubmittedIterator struct {
	Event *StreamChunkProofSubmitted // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofSubmitted)
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
		it.Event = new(StreamChunkProofSubmitted)
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
func (it *StreamChunkProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofSubmitted represents a ChunkProofSubmitted event raised by the Stream contract.
type StreamChunkProofSubmitted struct {
	ChunkId *big.Int
	Profile *big.Int
	Idx     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofSubmitted is a free log retrieval operation binding the contract event 0xb625ad3a7ea2d8a035a88ccb4001b583e72535a7093232080cc461ff6925917b.
//
// Solidity: event ChunkProofSubmitted(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) FilterChunkProofSubmitted(opts *bind.FilterOpts, chunkId []*big.Int, profile []*big.Int, idx []*big.Int) (*StreamChunkProofSubmittedIterator, error) {

	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofSubmitted", chunkIdRule, profileRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofSubmittedIterator{contract: _Stream.contract, event: "ChunkProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchChunkProofSubmitted is a free log subscription operation binding the contract event 0xb625ad3a7ea2d8a035a88ccb4001b583e72535a7093232080cc461ff6925917b.
//
// Solidity: event ChunkProofSubmitted(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) WatchChunkProofSubmitted(opts *bind.WatchOpts, sink chan<- *StreamChunkProofSubmitted, chunkId []*big.Int, profile []*big.Int, idx []*big.Int) (event.Subscription, error) {

	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}
	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofSubmitted", chunkIdRule, profileRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofSubmitted)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofSubmitted", log); err != nil {
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

// ParseChunkProofSubmitted is a log parse operation binding the contract event 0xb625ad3a7ea2d8a035a88ccb4001b583e72535a7093232080cc461ff6925917b.
//
// Solidity: event ChunkProofSubmitted(uint256 indexed chunkId, uint256 indexed profile, uint256 indexed idx)
func (_Stream *StreamFilterer) ParseChunkProofSubmitted(log types.Log) (*StreamChunkProofSubmitted, error) {
	event := new(StreamChunkProofSubmitted)
	if err := _Stream.contract.UnpackLog(event, "ChunkProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StreamChunkProofValidatedIterator is returned from FilterChunkProofValidated and is used to iterate over the raw logs and unpacked data for ChunkProofValidated events raised by the Stream contract.
type StreamChunkProofValidatedIterator struct {
	Event *StreamChunkProofValidated // Event containing the contract specifics and raw log

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
func (it *StreamChunkProofValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamChunkProofValidated)
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
		it.Event = new(StreamChunkProofValidated)
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
func (it *StreamChunkProofValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamChunkProofValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamChunkProofValidated represents a ChunkProofValidated event raised by the Stream contract.
type StreamChunkProofValidated struct {
	Profile *big.Int
	ChunkId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChunkProofValidated is a free log retrieval operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) FilterChunkProofValidated(opts *bind.FilterOpts, profile []*big.Int, chunkId []*big.Int) (*StreamChunkProofValidatedIterator, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Stream.contract.FilterLogs(opts, "ChunkProofValidated", profileRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamChunkProofValidatedIterator{contract: _Stream.contract, event: "ChunkProofValidated", logs: logs, sub: sub}, nil
}

// WatchChunkProofValidated is a free log subscription operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) WatchChunkProofValidated(opts *bind.WatchOpts, sink chan<- *StreamChunkProofValidated, profile []*big.Int, chunkId []*big.Int) (event.Subscription, error) {

	var profileRule []interface{}
	for _, profileItem := range profile {
		profileRule = append(profileRule, profileItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _Stream.contract.WatchLogs(opts, "ChunkProofValidated", profileRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamChunkProofValidated)
				if err := _Stream.contract.UnpackLog(event, "ChunkProofValidated", log); err != nil {
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

// ParseChunkProofValidated is a log parse operation binding the contract event 0x6e0dd2be4d72dbf12213042a47491e605ec38b66d076e6f74e9f458373080e37.
//
// Solidity: event ChunkProofValidated(uint256 indexed profile, uint256 indexed chunkId)
func (_Stream *StreamFilterer) ParseChunkProofValidated(log types.Log) (*StreamChunkProofValidated, error) {
	event := new(StreamChunkProofValidated)
	if err := _Stream.contract.UnpackLog(event, "ChunkProofValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
