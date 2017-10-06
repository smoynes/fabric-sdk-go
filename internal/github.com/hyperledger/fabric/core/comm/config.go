/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/

package comm

import (
	"github.com/spf13/viper"
)

var (
	// Is the configuration cached?
	configurationCached = false
	// Is TLS enabled
	tlsEnabled bool
	// Max send and receive bytes for grpc clients and servers
	maxRecvMsgSize = 100 * 1024 * 1024
	maxSendMsgSize = 100 * 1024 * 1024
	// Default keepalive options
	keepaliveOptions = KeepaliveOptions{
		ClientKeepaliveTime:    60,   // 1 min
		ClientKeepaliveTimeout: 20,   // 20 sec - gRPC default
		ServerKeepaliveTime:    7200, // 2 hours - gRPC default
		ServerKeepaliveTimeout: 20,   // 20 sec - gRPC default
	}
)

// KeepAliveOptions is used to set the gRPC keepalive settings for both
// clients and servers
type KeepaliveOptions struct {
	// ClientKeepaliveTime is the duration in seconds after which if the client
	// does not see any activity from the server it pings the server to see
	// if it is alive
	ClientKeepaliveTime int
	// ClientKeepaliveTimeout is the duration the client waits for a response
	// from the server after sending a ping before closing the connection
	ClientKeepaliveTimeout int
	// ServerKeepaliveTime is the duration in seconds after which if the server
	// does not see any activity from the client it pings the client to see
	// if it is alive
	ServerKeepaliveTime int
	// ServerKeepaliveTimeout is the duration the server waits for a response
	// from the client after sending a ping before closing the connection
	ServerKeepaliveTimeout int
}

// cacheConfiguration caches common package scoped variables
func cacheConfiguration() {
	if !configurationCached {
		tlsEnabled = viper.GetBool("peer.tls.enabled")
		configurationCached = true
	}
}

// TLSEnabled return cached value for "peer.tls.enabled" configuration value
func TLSEnabled() bool {
	if !configurationCached {
		cacheConfiguration()
	}
	return tlsEnabled
}

// MaxRecvMsgSize returns the maximum message size in bytes that gRPC clients
// and servers can receive
func MaxRecvMsgSize() int {
	return maxRecvMsgSize
}

// MaxSendMsgSize returns the maximum message size in bytes that gRPC clients
// and servers can send
func MaxSendMsgSize() int {
	return maxSendMsgSize
}
