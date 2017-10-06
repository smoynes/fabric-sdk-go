/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/

package util

import (
	"reflect"

	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/common/util"
)

type key struct {
	reflect.Value
	str string
}

type keys []*key

// ComputeStringHash computes the hash of the given string
func ComputeStringHash(input string) []byte {
	return ComputeHash([]byte(input))
}

// ComputeHash computes the hash of the given bytes
func ComputeHash(input []byte) []byte {
	return util.ComputeSHA256(input)
}
