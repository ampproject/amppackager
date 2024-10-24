package mailinabox

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SystemStatus Represents a system status.
type SystemStatus struct {
	Type  string        `json:"type,omitempty"`
	Text  string        `json:"text,omitempty"`
	Extra []ExtraStatus `json:"extra,omitempty"`
}

// ExtraStatus Represents extra status.
type ExtraStatus struct {
	Monospace bool   `json:"monospace,omitempty"`
	Text      string `json:"text,omitempty"`
}

// BackupStatus Represents a backup status.
type BackupStatus struct {
	Backups           []Backup `json:"backups,omitempty"`
	UnmatchedFileSize int      `json:"unmatched_file_size,omitempty"`
	Error             string   `json:"error,omitempty"`
}

// Backup Represents a backup.
type Backup struct {
	Date      string `json:"date,omitempty"`
	DateDelta string `json:"date_delta,omitempty"`
	DateStr   string `json:"date_str,omitempty"`
	DeletedIn string `json:"deleted_in,omitempty"`
	Full      bool   `json:"full,omitempty"`
	Size      int    `json:"size,omitempty"`
	Volumes   int    `json:"volumes,omitempty"`
}

// BackupConfig Represents a backup configuration.
type BackupConfig struct {
	EncPwFile           string `json:"enc_pw_file,omitempty"`
	FileTargetDirectory string `json:"file_target_directory,omitempty"`
	MinAgeInDays        int    `json:"min_age_in_days,omitempty"`
	SSHPubKey           string `json:"ssh_pub_key,omitempty"`
	Target              string `json:"target,omitempty"`
	TargetUser          string `json:"target_user,omitempty"`
	TargetPass          string `json:"target_pass,omitempty"`
}

// SystemService System API.
// https://mailinabox.email/api-docs.html#tag/System
type SystemService service

// GetStatus Returns an array of statuses which can include headings.
// https://mailinabox.email/api-docs.html#operation/getSystemStatus
func (s *SystemService) GetStatus(ctx context.Context) ([]SystemStatus, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "status")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var results []SystemStatus

	err = s.client.doJSON(req, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetVersion Returns installed Mail-in-a-Box version.
// https://mailinabox.email/api-docs.html#operation/getSystemVersion
func (s *SystemService) GetVersion(ctx context.Context) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "version")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetUpstreamVersion Returns Mail-in-a-Box upstream version.
// https://mailinabox.email/api-docs.html#operation/getSystemUpstreamVersion
func (s *SystemService) GetUpstreamVersion(ctx context.Context) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "latest-upstream-version")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetUpdates Returns system (apt) updates.
// https://mailinabox.email/api-docs.html#operation/getSystemUpdates
func (s *SystemService) GetUpdates(ctx context.Context) ([]string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "updates")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return nil, err
	}

	updates := strings.Split(strings.TrimSpace(string(resp)), "\n")

	return updates, nil
}

// UpdatePackages Updates system (apt) packages.
// https://mailinabox.email/api-docs.html#operation/getSystemUpdates
func (s *SystemService) UpdatePackages(ctx context.Context) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "update-packages")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetPrivacyStatus Returns system privacy (new-version check) status.
// https://mailinabox.email/api-docs.html#operation/getSystemPrivacyStatus
func (s *SystemService) GetPrivacyStatus(ctx context.Context) (bool, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "privacy")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return false, fmt.Errorf("unable to create request: %w", err)
	}

	var result bool

	err = s.client.doJSON(req, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

// UpdatePrivacyStatus Updates system privacy (new-version checks).
// - value: `private`: Disable new version checks
// - value: `off`: Enable new version checks
// https://mailinabox.email/api-docs.html#operation/updateSystemPrivacy
func (s *SystemService) UpdatePrivacyStatus(ctx context.Context, value string) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "privacy")

	data := url.Values{}
	data.Set("value", value)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetRebootStatus Returns the system reboot status.
// https://mailinabox.email/api-docs.html#operation/getSystemRebootStatus
func (s *SystemService) GetRebootStatus(ctx context.Context) (bool, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "reboot")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return false, fmt.Errorf("unable to create request: %w", err)
	}

	var result bool

	err = s.client.doJSON(req, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

// Reboot Reboots the system.
// https://mailinabox.email/api-docs.html#operation/rebootSystem
func (s *SystemService) Reboot(ctx context.Context) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "reboot")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), http.NoBody)
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}

// GetBackupStatus Returns the system backup status.
// https://mailinabox.email/api-docs.html#operation/getSystemBackupStatus
func (s *SystemService) GetBackupStatus(ctx context.Context) (*BackupStatus, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "backup", "status")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var result BackupStatus

	err = s.client.doJSON(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBackupConfig Returns the system backup config.
// https://mailinabox.email/api-docs.html#operation/getSystemBackupConfig
func (s *SystemService) GetBackupConfig(ctx context.Context) (*BackupConfig, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "backup", "config")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	var result BackupConfig

	err = s.client.doJSON(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateBackupConfig Updates the system backup config.
// https://mailinabox.email/api-docs.html#operation/updateSystemBackupConfig
func (s *SystemService) UpdateBackupConfig(ctx context.Context, target, targetUser, targetPass string, minAge int) (string, error) {
	endpoint := s.client.baseURL.JoinPath("admin", "system", "backup", "config")

	data := url.Values{}
	data.Set("target", target)
	data.Set("targetUser", targetUser)
	data.Set("targetPass", targetPass)
	data.Set("minAge", strconv.Itoa(minAge))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.doPlain(req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resp)), nil
}
