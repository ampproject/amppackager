// Copyright 2022-2023 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// EDiskEncryptionAlgorithm ディスク暗号化アルゴリズム
type EDiskEncryptionAlgorithm string

// String EDiskEncryptionAlgorithmの文字列表現
func (c EDiskEncryptionAlgorithm) String() string {
	return string(c)
}

var (
	// DiskEncryptionAlgorithms ディスク接続方法
	DiskEncryptionAlgorithms = struct {
		// None unencrypted
		None EDiskEncryptionAlgorithm
		// AES256XTS aes256_xts
		AES256XTS EDiskEncryptionAlgorithm
	}{
		None:      EDiskEncryptionAlgorithm("none"),
		AES256XTS: EDiskEncryptionAlgorithm("aes256_xts"),
	}

	// DiskEncryptionAlgorithmMap 文字列とDiskEncryptionAlgorithmのマップ
	DiskEncryptionAlgorithmMap = map[string]EDiskEncryptionAlgorithm{
		DiskEncryptionAlgorithms.None.String():      DiskEncryptionAlgorithms.None,
		DiskEncryptionAlgorithms.AES256XTS.String(): DiskEncryptionAlgorithms.AES256XTS,
	}

	// DiskEncryptionAlgorithmStrings DiskEncryptionAlgorithmに指定できる有効な文字列
	DiskEncryptionAlgorithmStrings = []string{
		DiskEncryptionAlgorithms.None.String(),
		DiskEncryptionAlgorithms.AES256XTS.String(),
	}
)
