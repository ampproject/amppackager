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

// ENoteClass スタートアップスクリプトクラス
type ENoteClass string

var (
	// NoteClassShell shellクラス
	NoteClassShell = ENoteClass("shell")
	// NoteClassYAMLCloudConfig yaml_cloud_configクラス
	NoteClassYAMLCloudConfig = ENoteClass("yaml_cloud_config")
)

// ENoteClasses 設定可能なスタートアップスクリプトクラス
var ENoteClasses = []ENoteClass{NoteClassShell, NoteClassYAMLCloudConfig}

// propNoteClass スタートアップスクリプトクラス情報内包型
type propNoteClass struct {
	Class ENoteClass `json:",omitempty"` // クラス
}

// GetClass クラス 取得
func (p *propNoteClass) GetClass() ENoteClass {
	return p.Class
}

// SetClass クラス 設定
func (p *propNoteClass) SetClass(c ENoteClass) {
	p.Class = c
}

// GetClassStr クラス 取得(文字列)
func (p *propNoteClass) GetClassStr() string {
	return string(p.Class)
}

// SetClassByStr クラス 設定(文字列)
func (p *propNoteClass) SetClassByStr(c string) {
	p.Class = ENoteClass(c)
}
