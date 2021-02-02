// HTML form minimal example using golang web-server
// Visit: http://127.0.0.1:8080
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	os.Setenv("PATH", "/Users/jckim/Documents/Workspace/hyperledger/fabric-samples/bin")
	os.Setenv("FABRIC_CFG_PATH", "/Users/jckim/Documents/Workspace/hyperledger/fabric-samples/config/")
	os.Setenv("CORE_PEER_TLS_ENABLED", "true")
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "/Users/jckim/Documents/Workspace/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt")
	os.Setenv("CORE_PEER_MSPCONFIGPATH", "/Users/jckim/Documents/Workspace/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp")
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")

	command := exec.Command(
		"peer",
		"chaincode invoke",
		"--orderer localhost:7050",
		"--ordererTLSHostnameOverride orderer.example.com",
		"--tls",
		"--cafile /Users/jckim/Documents/Workspace/hyperledger/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem",
		"--tlsRootCertFiles /Users/jckim/Documents/Workspace/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt",
		"--tlsRootCertFiles /Users/jckim/Documents/Workspace/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt",
		"--peerAddresses localhost:7051",
		"--peerAddresses localhost:9051",
		"--channelID mychannel",
		"--name omni",
		"--ctor '{\"Args\":[\"getvcinfo\", \"9cdcd186-4396-491f-ac85-dc13b3898432\"]}'",
	)

	// set var to get the output
	var out bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(out.String())
}
