.PHONY: ethereum waves

ethereum:
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "abigen" 2> /dev/null || echo 'Please install abigen'
	# Common
	abigen --pkg=luport --sol="./contracts/ethereum/LUPort.sol" --out="./abi/ethereum/luport/luport.go"
	abigen --pkg=ibport --sol="./contracts/ethereum/IBPort.sol" --out="./abi/ethereum/ibport/ibport.go"
	abigen --pkg=erc20 --sol="./contracts/ethereum/Token.sol" --out="./abi/ethereum/erc20/erc20.go"
	# USDN
	abigen --pkg=erc20 --sol="./contracts/stakable/USDN.sol" --out="./abi/ethereum/usdn-erc20/erc20.go"
	# SWOP
	abigen --pkg=erc20 --sol="./contracts/stakable/SWOP.sol" --out="./abi/ethereum/swop-erc20/erc20.go"
	# Autostaking NSBT
	abigen --pkg=erc20 --sol="./contracts/stakable/NSBT.sol" --out="./abi/ethereum/gwa-nsbt/erc20.go"
	# Autostaking IB Port
	abigen --pkg=ibport --sol="./contracts/stakable/IBPort.sol" --out="./abi/ethereum/autostaking-ibport/ibport.go"
	echo "Ethereum abi updated"

waves:
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "surfboard" 2> /dev/null || echo 'Please install sorfboard'
	surfboard compile ./contracts/waves/luport.ride > ./abi/waves/luport.abi
	surfboard compile ./contracts/waves/ibport.ride > ./abi/waves/ibport.abi
	echo "Waves abi updated"
