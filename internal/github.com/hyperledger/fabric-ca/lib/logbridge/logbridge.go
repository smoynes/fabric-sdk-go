/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package logbridge

import (
	clog "github.com/cloudflare/cfssl/log"
	"github.com/hyperledger/fabric-sdk-go/pkg/logging"
)

var logger *logging.Logger
var cfLogBridge *cLogger

func init() {
	logger = logging.NewLogger("fabric_sdk_go")
	cfLogBridge = &cLogger{}
	clog.SetLogger(cfLogBridge)
}

// Debug bridges calls to the Go SDK logger's Debug.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf bridges calls to the Go SDK logger's Debugf.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

// Info bridges calls to the Go SDK logger's Info.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof bridges calls to the Go SDK logger's Debugf.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Fatalf bridges calls to the Go SDK logger's Debugf.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
