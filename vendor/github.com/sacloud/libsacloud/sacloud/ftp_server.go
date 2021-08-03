// Copyright 2016-2020 The Libsacloud Authors
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

package sacloud

// FTPServer FTPサーバー接続情報
type FTPServer struct {
	HostName  string `json:",omitempty"` // FTPサーバーホスト名
	IPAddress string `json:",omitempty"` // FTPサーバー IPアドレス
	User      string `json:",omitempty"` // 接続ユーザー名
	Password  string `json:",omitempty"` // パスワード

}

// FTPOpenRequest FTP接続オープンリクエスト
type FTPOpenRequest struct {
	ChangePassword bool // パスワード変更フラグ
}
