/*
   Copyright The containerd Authors.

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

package encryption

import "github.com/gogo/protobuf/types"

// pbAny takes proto-generated Any type.
// https://developers.google.com/protocol-buffers/docs/proto3#any
type pbAny interface {
	GetTypeUrl() string
	GetValue() []byte
}

func fromAny(from pbAny) *types.Any {
	if from == nil {
		return nil
	}

	pbany, ok := from.(*types.Any)
	if ok {
		return pbany
	}

	return &types.Any{
		TypeUrl: from.GetTypeUrl(),
		Value:   from.GetValue(),
	}
}
