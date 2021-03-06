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

package histmgr

import "github.com/hyperledger/fabric/protos/common"
import "github.com/hyperledger/fabric/core/ledger"

// HistMgr - an interface that a history manager should implement
type HistMgr interface {
	NewHistoryQueryExecutor() (ledger.HistoryQueryExecutor, error)
	Commit(block *common.Block) error
	GetBlockNumFromSavepoint() (uint64, error)
}
