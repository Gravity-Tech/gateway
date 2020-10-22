.PHONY: ethereum waves

ethereum:
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "abigen" 2> /dev/null || echo 'Please install abigen'
	abigen --pkg=luport --sol="./contracts/ethereum/LUPort.sol" --out="./abi/ethereum/luport/luport.go"
	abigen --pkg=ibport --sol="./contracts/ethereum/IBPort.sol" --out="./abi/ethereum/ibport/ibport.go"
	echo "Ethereum abi updated"

waves:
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "surfboard" 2> /dev/null || echo 'Please install sorfboard'
	npm update surfboard -g
	surfboard compile ./contracts/waves/luport.ride > ./abi/waves/luport.abi
	surfboard compile ./contracts/waves/ibport.ride > ./abi/waves/ibport.abi
	echo "Waves abi updated"