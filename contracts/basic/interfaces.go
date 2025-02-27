/*
Copyright 2023 The Bestchains Authors.

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

package basic

import (
	"github.com/bestchains/bestchains-contracts/contracts/access"
	"github.com/bestchains/bestchains-contracts/contracts/nonce"
	"github.com/bestchains/bestchains-contracts/library/context"
)

type EventPutValue struct {
	Index uint64
	KID   string
}

// IBasic provides common data Put/Get
type IBasic interface {
	nonce.INonce
	access.IAccessControl
	// Initialize the contract
	Initialize(ctx context.ContextInterface) error
	// Total k/v paris stored
	Total(ctx context.ContextInterface) (uint64, error)
	// PutValue stores kval with pre-defined key calculation
	PutValue(ctx context.ContextInterface, msg context.Message, val string) (string, error)
	// GetValueByIndex get kval with index
	GetValueByIndex(ctx context.ContextInterface, index string) (string, error)
	// GetValueByKID get kval with key id
	GetValueByKID(ctx context.ContextInterface, kid string) (string, error)
}
