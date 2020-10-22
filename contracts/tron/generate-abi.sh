#!/bin/bash

abigen --pkg luport --solc /usr/local/bin/solc --sol LUPort/LUPort.sol > ../api/luport/luport.go
abigen --pkg nebula --solc /usr/local/bin/solc --sol Nebula/Nebula.sol > ../api/nebula/nebula.go
abigen --pkg ibport --solc /usr/local/bin/solc --sol IBPort/IBPort.sol > ../api/ibport/ibport.go
