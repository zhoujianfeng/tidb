// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package structure

import (
	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/terror"
)

// structure error codes.
const (
	codeInvalidHashKeyFlag   terror.ErrCode = 1
	codeInvalidHashKeyPrefix                = 2
	codeInvalidListIndex                    = 3
	codeInvalidListMetaData                 = 4
)

var (
	errInvalidHashKeyFlag   = terror.ClassStructure.New(codeInvalidHashKeyFlag, "invalid encoded hash key flag")
	errInvalidHashKeyPrefix = terror.ClassStructure.New(codeInvalidHashKeyPrefix, "invalid encoded hash key prefix")
	errInvalidListIndex     = terror.ClassStructure.New(codeInvalidListMetaData, "invalid list index")
	errInvalidListMetaData  = terror.ClassStructure.New(codeInvalidListMetaData, "invalid list meta data")
)

// NewStructure creates a TxStructure in transaction txn and with key prefix.
func NewStructure(txn kv.Transaction, prefix []byte) *TxStructure {
	return &TxStructure{
		txn:    txn,
		prefix: prefix,
	}
}

// TxStructure supports some simple data structures like string, hash, list, etc... and
// you can use these in a transaction.
type TxStructure struct {
	txn    kv.Transaction
	prefix []byte
}