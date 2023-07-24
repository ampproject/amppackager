// Copyright 2022-2023 The sacloud/api-client-go Authors
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

package profile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// DirectoryNameEnv プロファイルの格納先を指定する環境変数
	DirectoryNameEnv = "SAKURACLOUD_PROFILE_DIR"
	// DirectoryNameEnvOld プロファイルの格納先を指定する環境変数(後方互換)
	DirectoryNameEnvOld = "USACLOUD_PROFILE_DIR"
	// DefaultProfileName デフォルトのプロファイル名
	DefaultProfileName = "default"

	// EnableAPITraceWord TraceModeに設定する、APIトレースを有効化するためのキーワード
	EnableAPITraceWord = "api"

	// EnableHTTPTraceWord TraceModeに設定する、HTTPトレースを有効化するためのキーワード
	EnableHTTPTraceWord = "http"
)

var (
	configDirName   = ".usacloud"
	configFileName  = "config.json"
	currentFileName = "current"
)

// ValidateName プロファイル名が有効か検証
func ValidateName(profileName string, invalidRunes ...rune) error {
	invalids := invalidRunes
	if len(invalids) == 0 {
		// validate profileName
		invalids = []rune{filepath.ListSeparator, filepath.Separator}
	}

	for _, r := range invalids {
		if strings.ContainsRune(profileName, r) {
			return fmt.Errorf("got invalid profile name: %s", profileName)
		}
	}
	return nil
}

func loadProfileDirFromEnvs() (string, error) {
	dir, err := loadProfileDirFromEnv(DirectoryNameEnv)
	if err != nil {
		return "", err
	}
	if dir == "" {
		// fallback
		dir, err = loadProfileDirFromEnv(DirectoryNameEnvOld)
		if err != nil {
			return "", err
		}
	}
	return dir, nil
}

func loadProfileDirFromEnv(key string) (string, error) {
	if path, ok := os.LookupEnv(key); ok {
		if err := ValidateName(path, filepath.ListSeparator); err != nil {
			return "", fmt.Errorf("loading ProfileDir from environment variables[%s] is failed: %s", key, err)
		}
		return filepath.Clean(path), nil
	}
	return "", nil
}

func baseDir() (string, error) {
	// from profileDirEnv var
	path, err := loadProfileDirFromEnvs()
	if path != "" || err != nil {
		return path, err
	}

	// default, use homedir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user's home dir is failed: %s", err)
	}
	return homeDir, nil
}

// ConfigDir プロファイルを格納するディレクトリのフルパス
func ConfigDir() (string, error) {
	baseDir, err := baseDir()
	if err != nil {
		return "", fmt.Errorf("getting profile base dir is failed: %s", err)
	}
	return filepath.Clean(filepath.Join(baseDir, configDirName)), nil
}

// ConfigFilePath 指定のプロファイル名のコンフィグファイルパスを取得
func ConfigFilePath(profileName string) (string, error) {
	if err := ValidateName(profileName); err != nil {
		return "", err
	}

	if profileName == "" {
		profileName = DefaultProfileName
	}
	baseDir, err := baseDir()
	if err != nil {
		return "", fmt.Errorf("getting profile base dir is failed: %s", err)
	}
	return filepath.Clean(filepath.Join(baseDir, configDirName, filepath.Clean(profileName), configFileName)), nil
}

// ConfigValue プロファイル コンフィグ
type ConfigValue struct {
	// AccessToken アクセストークン
	AccessToken string
	// AccessTokenSecret アクセスシークレット
	AccessTokenSecret string

	// Zone デフォルトゾーン
	Zone string
	// Zones 利用可能なゾーン
	Zones []string

	// UserAgent ユーザーエージェント
	UserAgent string `json:",omitempty"`
	// AcceptLanguage リクエスト時のAccept-Languageヘッダ
	AcceptLanguage string
	// Gzip Gzip圧縮の有効化
	Gzip bool

	// RetryMax 423/503時のリトライ回数
	RetryMax int
	// RetryMin 423/503時のリトライ間隔(最小) 単位:秒
	RetryWaitMin int
	// RetryMax 423/503時のリトライ間隔(最大) 単位:秒
	RetryWaitMax int

	// StatePollingTimeout StatePollWaiterでのタイムアウト 単位:秒
	StatePollingTimeout int
	// StatePollingInterval StatePollWaiterでのポーリング間隔 単位:秒
	StatePollingInterval int

	// HTTPRequestTimeout APIリクエスト時のHTTPタイムアウト 単位:秒
	HTTPRequestTimeout int
	// HTTPRequestRateLimit APIリクエスト時の1秒あたりのリクエスト上限数
	HTTPRequestRateLimit int

	// APIRootURL APIのルートURL
	APIRootURL string

	// DefaultZone グローバルリソースAPIを呼ぶ際に指定するゾーン
	DefaultZone string

	// TraceMode トレースモード
	TraceMode string
	// FakeMode フェイクモード有効化
	FakeMode bool
	// FakeStorePath フェイクモードでのファイルストアパス
	FakeStorePath string
}

func (o *ConfigValue) EnableHTTPTrace() bool {
	return EnableHTTPTrace(o.TraceMode)
}

func (o *ConfigValue) EnableAPITrace() bool {
	return EnableAPITrace(o.TraceMode)
}

func traceModeValue(strTraceMode string) string {
	return strings.ToLower(strings.TrimSpace(strTraceMode))
}

func EnableHTTPTrace(strTraceMode string) bool {
	traceMode := traceModeValue(strTraceMode)
	if traceMode == "" {
		return false
	}

	// TraceModeが"api"の場合はfalseにする(TraceMode=1などの場合はAPI/HTTP両方が有効になる)
	if traceMode == EnableAPITraceWord {
		return false
	}
	return true
}

func EnableAPITrace(strTraceMode string) bool {
	traceMode := traceModeValue(strTraceMode)
	if traceMode == "" {
		return false
	}

	// TraceModeが"http"の場合はfalseにする(TraceMode=1などの場合はAPI/HTTP両方が有効になる)
	if traceMode == EnableHTTPTraceWord || traceMode == "error" {
		return false
	}
	return true
}

// Save プロファイルコンフィグを保存
func Save(profileName string, val interface{}) error {
	if val == nil {
		return fmt.Errorf("config is required")
	}

	path, err := ConfigFilePath(profileName)
	if err != nil {
		return err
	}

	// create dir
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("creating profile directory[%q] is failed: %s", dir, err)
		}
	}

	rawBody, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling config to JSON is failed: %s", err)
	}

	// merge new value if current config exists
	if _, err := os.Stat(path); err == nil {
		currentData, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading current config %q failed: %s", path, err)
		}
		var currentDataMap map[string]interface{}
		if err := json.Unmarshal(currentData, &currentDataMap); err != nil {
			return fmt.Errorf("unmarshaling current config %q failed: %s", path, err)
		}

		var newDataMap map[string]interface{}
		if err := json.Unmarshal(rawBody, &newDataMap); err != nil {
			return fmt.Errorf("unmarshaling new config %q failed: %s", path, err)
		}

		// merge
		for k, v := range newDataMap {
			currentDataMap[k] = v
		}

		rawBody, err = json.MarshalIndent(currentDataMap, "", "  ")
		if err != nil {
			return fmt.Errorf("marshalling new config to JSON failed: %s", err)
		}
	}

	err = os.WriteFile(path, rawBody, 0600)
	if err != nil {
		return fmt.Errorf("writing config to %q is failed: %s", path, err)
	}

	return nil
}

// Load 指定のプロファイル名からロードする
//
// configValueには*profile.ConfigValue(派生)への参照を渡す
//
// 指定したプロファイル名に対応するコンフィグファイルが存在しない場合はエラーを返す
// ただしデフォルトのプロファイル名の場合はファイルが存在しなくてもエラーにしない
func Load(profileName string, configValue interface{}) error {
	filePath, err := ConfigFilePath(profileName)
	if err != nil {
		return err
	}

	// file exists?
	if _, err := os.Stat(filePath); err == nil {
		// read file
		buf, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("loading config from %q is failed: %s", filePath, err)
		}
		if err := json.Unmarshal(buf, configValue); err != nil {
			return fmt.Errorf("parsing config is failed: %s", err)
		}
	} else if profileName != DefaultProfileName {
		return fmt.Errorf("profile %q is not exists", profileName)
	}

	return nil
}

// Remove 指定のプロファイルのコンフィグを削除する
//
// プロファイルディレクトリが空になる場合はディレクトリも合わせて削除する
// Currentプロファイルが削除された場合はCurrentをデフォルトに設定する
func Remove(profileName string) error {
	path, err := ConfigFilePath(profileName)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		return fmt.Errorf("removing directory is failed: %q is not exists", dir)
	}

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("removing config is failed: %q is not exists", path)
	}

	// remove file
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("removing config %q is failed: %s", path, err)
	}

	// remove dir if dir is empty
	info, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("removing config file is failed: reading %q is failed: %s", dir, err)
	}
	if len(info) == 0 {
		// remove dir
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("removing config dir %q is failed: %s", dir, err)
		}
	}

	current, err := CurrentName()
	if err != nil {
		return fmt.Errorf("removing config is failed: CurrentName() returns error: %s", err)
	}

	if current == profileName {
		if err := SetCurrentName(DefaultProfileName); err != nil {
			return fmt.Errorf("removing config is failed: SetCurrentName() returns error: %s", err)
		}
	}
	return nil
}

// CurrentName カレントプロファイル名
func CurrentName() (string, error) {
	baseDir, err := baseDir()
	if err != nil {
		return "", err
	}

	profNameFile := filepath.Join(baseDir, configDirName, currentFileName)
	if _, err := os.Stat(profNameFile); err == nil {
		data, err := os.ReadFile(profNameFile)
		if err != nil {
			return "", fmt.Errorf("reading current profile is failed: %s", err)
		}
		profileName := string(data)
		if err := ValidateName(profileName); err != nil {
			return "", err
		}

		profileName = cleanupProfileName(profileName)
		if profileName == "" {
			profileName = DefaultProfileName
		}
		return profileName, nil
	}

	return DefaultProfileName, nil
}

func cleanupProfileName(profileName string) string {
	targets := []string{"　", "\t", "\n"}
	res := profileName
	for _, s := range targets {
		res = strings.ReplaceAll(res, s, "")
	}
	return strings.Trim(res, " ")
}

// SetCurrentName カレントプロファイル名を設定する
func SetCurrentName(profileName string) error {
	if err := ValidateName(profileName); err != nil {
		return err
	}

	profileName = cleanupProfileName(profileName)

	baseDir, err := baseDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(baseDir, configDirName)
	if _, err := os.Stat(configDir); err != nil {
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			return fmt.Errorf("creating config dir %q is failed: %s", configDir, err)
		}
	}

	if profileName != DefaultProfileName {
		profileConfigPath := filepath.Join(configDir, profileName, configFileName)
		if _, err := os.Stat(profileConfigPath); err != nil {
			return fmt.Errorf("profile %q is not exists", profileName)
		}
	}

	profNameFile := filepath.Join(baseDir, configDirName, currentFileName)
	if err := os.WriteFile(profNameFile, []byte(profileName), 0600); err != nil {
		return fmt.Errorf("writing profile to %q is failed: %s", profNameFile, err)
	}

	return nil
}

// List プロファイル名の一覧を返す
func List() ([]string, error) {
	res := []string{"default"}

	// get profile dirs under base dir
	baseDir, err := baseDir()
	if err != nil {
		return []string{}, fmt.Errorf("listing profiles is failed: %s", err)
	}
	configDirPath := filepath.Join(baseDir, configDirName)
	if _, err := os.Stat(configDirPath); err != nil {
		return res, nil
	}
	entries, err := os.ReadDir(filepath.Join(baseDir, configDirName))
	if err != nil {
		return []string{}, fmt.Errorf("listing profiles is failed: %s", err)
	}

	// validate each profile dir
	for _, fi := range entries {
		if fi.IsDir() {
			profile := filepath.Base(fi.Name())
			if profile != DefaultProfileName {
				if profile != DefaultProfileName {
					c := &ConfigValue{}
					if err := Load(profile, c); err == nil {
						res = append(res, profile)
					}
				}
			}
		}
	}

	return res, nil
}
