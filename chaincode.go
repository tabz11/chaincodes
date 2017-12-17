package main

import( "fmt"
		"github.com/hyperledger/fabric/core/chaincode/shim"
		sc "github.com/hyperledger/fabric/protos/peer"
	)

	type SimpleChainCode struct {}
	func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
		return shim.Success([]byte ("Init Called"))
	}

	func (t *SimpleChainCode) Invoke (stub shim.ChaincodeStubInterface) pb.Response {
		return shim.Success([]byte ("Invoke is called"))
	}

	func main ()
	{ err:= shim.Start(new(SimpleChainCode))
	if err!=nil {fmt.Printf("Error : ", err)}}