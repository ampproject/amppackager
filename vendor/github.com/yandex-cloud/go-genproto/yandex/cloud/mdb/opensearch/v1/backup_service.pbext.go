// Code generated by protoc-gen-goext. DO NOT EDIT.

package opensearch

func (m *GetBackupRequest) SetBackupId(v string) {
	m.BackupId = v
}

func (m *ListBackupsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListBackupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBackupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBackupsResponse) SetBackups(v []*Backup) {
	m.Backups = v
}

func (m *ListBackupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
