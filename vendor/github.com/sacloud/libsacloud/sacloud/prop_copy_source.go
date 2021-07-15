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

// propCopySource コピー元情報内包型
type propCopySource struct {
	SourceDisk    *Disk    `json:",omitempty"` // コピー元ディスク
	SourceArchive *Archive `json:",omitempty"` // コピー元アーカイブ

}

// SetSourceArchive ソースアーカイブ設定
func (p *propCopySource) SetSourceArchive(sourceID ID) {
	if sourceID.IsEmpty() {
		return
	}
	p.SourceArchive = &Archive{
		Resource: &Resource{ID: sourceID},
	}
	p.SourceDisk = nil
}

// SetSourceDisk ソースディスク設定
func (p *propCopySource) SetSourceDisk(sourceID ID) {
	if sourceID.IsEmpty() {
		return
	}
	p.SourceDisk = &Disk{
		Resource: &Resource{ID: sourceID},
	}
	p.SourceArchive = nil
}

// GetSourceArchive ソースアーカイブ取得
func (p *propCopySource) GetSourceArchive() *Archive {
	return p.SourceArchive
}

// GetSourceDisk ソースディスク取得
func (p *propCopySource) GetSourceDisk() *Disk {
	return p.SourceDisk
}

// GetSourceArchiveID ソースアーカイブID取得
func (p *propCopySource) GetSourceArchiveID() ID {
	if p.SourceArchive != nil {
		return p.SourceArchive.ID
	}
	return EmptyID
}

// GetSourceDiskID ソースディスクID取得
func (p *propCopySource) GetSourceDiskID() ID {
	if p.SourceDisk != nil {
		return p.SourceDisk.ID
	}
	return EmptyID
}
