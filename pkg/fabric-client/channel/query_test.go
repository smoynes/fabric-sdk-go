/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package channel

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/mocks"
)

func TestQueryMethods(t *testing.T) {
	channel, _ := setupTestChannel()

	_, err := channel.QueryBlock(-1)
	if err == nil {
		t.Fatalf("Query block cannot be negative number")
	}

	_, err = channel.QueryBlockByHash(nil)
	if err == nil {
		t.Fatalf("Query hash cannot be nil")
	}
	_, err = channel.QueryByChaincode("", []string{"method"}, nil)
	if err == nil {
		t.Fatalf("QueryByChannelcode: name cannot be empty")
	}

	_, err = channel.QueryByChaincode("qscc", nil, nil)
	if err == nil {
		t.Fatalf("QueryByChannelcode: arguments cannot be empty")
	}

	_, err = channel.QueryByChaincode("qscc", []string{"method"}, nil)
	if err == nil {
		t.Fatalf("QueryByChannelcode: targets cannot be empty")
	}

}

func TestChannelQueryBlock(t *testing.T) {

	channel, _ := setupTestChannel()

	peer := mocks.MockPeer{MockName: "Peer1", MockURL: "http://peer1.com", MockRoles: []string{}, MockCert: nil}
	err := channel.AddPeer(&peer)

	_, err = channel.QueryBlock(1)

	if err != nil {
		t.Fatal("Test channel query block failed,")
	}

	_, err = channel.QueryBlockByHash([]byte(""))

	if err != nil {
		t.Fatal("Test channel query block by hash failed,")
	}

}

func TestQueryInstantiatedChaincodes(t *testing.T) {
	channel, _ := setupTestChannel()

	peer := mocks.MockPeer{MockName: "Peer1", MockURL: "http://peer1.com", MockRoles: []string{}, MockCert: nil}
	err := channel.AddPeer(&peer)

	res, err := channel.QueryInstantiatedChaincodes()

	if err != nil || res == nil {
		t.Fatal("Test QueryInstatiated chaincode failed")
	}

}

func TestQueryTransaction(t *testing.T) {
	channel, _ := setupTestChannel()

	peer := mocks.MockPeer{MockName: "Peer1", MockURL: "http://peer1.com", MockRoles: []string{}, MockCert: nil}
	err := channel.AddPeer(&peer)

	res, err := channel.QueryTransaction("txid")

	if err != nil || res == nil {
		t.Fatal("Test QueryTransaction failed")
	}
}

func TestQueryInfo(t *testing.T) {
	channel, _ := setupTestChannel()

	peer := mocks.MockPeer{MockName: "Peer1", MockURL: "http://peer1.com", MockRoles: []string{}, MockCert: nil}
	err := channel.AddPeer(&peer)

	res, err := channel.QueryInfo()

	if err != nil || res == nil {
		t.Fatal("Test QueryInfo failed")
	}
}
