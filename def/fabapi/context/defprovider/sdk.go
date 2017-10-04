/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package defprovider

import (
	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"

	"github.com/hyperledger/fabric-sdk-go/def/fabapi/opt"
	configImpl "github.com/hyperledger/fabric-sdk-go/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/errors"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"
	signingMgr "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/signingmgr"
	discovery "github.com/hyperledger/fabric-sdk-go/pkg/fabric-txn/discovery"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/bccsp"
	bccspFactory "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/bccsp/factory"
)

// DefaultProviderFactory represents the default SDK provider factory.
type DefaultProviderFactory struct{}

// NewDefaultProviderFactory returns the default SDK provider factory.
func NewDefaultProviderFactory() *DefaultProviderFactory {
	f := DefaultProviderFactory{}
	return &f
}

// NewConfigProvider creates a Config using the SDK's default implementation
func (f *DefaultProviderFactory) NewConfigProvider(o opt.ConfigOpts, a opt.SDKOpts) (apiconfig.Config, error) {
	return configImpl.InitConfig(a.ConfigFile)
}

// NewStateStoreProvider creates a KeyValueStore using the SDK's default implementation
func (f *DefaultProviderFactory) NewStateStoreProvider(o opt.StateStoreOpts, config apiconfig.Config) (fab.KeyValueStore, error) {

	var stateStorePath = o.Path
	if stateStorePath == "" {
		clientCofig, err := config.Client()
		if err != nil {
			return nil, err
		}
		stateStorePath = clientCofig.CredentialStore.Path
	}

	stateStore, err := kvs.CreateNewFileKeyValueStore(stateStorePath)
	if err != nil {
		return nil, errors.WithMessage(err, "CreateNewFileKeyValueStore failed")
	}
	return stateStore, nil
}

// NewCryptoSuiteProvider returns a new default implementation of BCCSP
func (f *DefaultProviderFactory) NewCryptoSuiteProvider(config *bccspFactory.FactoryOpts) (bccsp.BCCSP, error) {
	return bccspFactory.GetBCCSPFromOpts(config)
}

// NewSigningManager returns a new default implementation of signing manager
func (f *DefaultProviderFactory) NewSigningManager(cryptoProvider bccsp.BCCSP, config apiconfig.Config) (fab.SigningManager, error) {
	return signingMgr.NewSigningManager(cryptoProvider, config)
}

// NewDiscoveryProvider returns a new default implementation of discovery provider
func (f *DefaultProviderFactory) NewDiscoveryProvider(config apiconfig.Config) (fab.DiscoveryProvider, error) {
	return discovery.NewDiscoveryProvider(config)

}
