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

package naked

import (
	"encoding/json"
	"sort"
	"strings"
	"time"
)

// MonitorValues アクティビティモニタのレスポンス向け汎用エンベロープ
type MonitorValues struct {
	// CPU CPU-TIME
	CPU MonitorCPUTimeValues
	// Disk Read/Write
	Disk MonitorDiskValues
	// Interface Send/Receive
	Interface MonitorInterfaceValues
	// Router In/Out
	Router MonitorRouterValues
	// Database データベース
	Database MonitorDatabaseValues
	// FreeDiskSize 空きディスクサイズ(NFS)
	FreeDiskSize MonitorFreeDiskSizeValues
	// ResponseTimeSec 応答時間(シンプル監視)
	ResponseTimeSec MonitorResponseTimeSecValues
	// Link UplinkBPS/DownlinkBPS
	Link MonitorLinkValues
	// Connection 接続数
	Connection MonitorConnectionValues
	// LocalRouter Receive/Send bytes per sec
	LocalRouter MonitorLocalRouterValues
}

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorValues) UnmarshalJSON(data []byte) error {
	v := MonitorValues{}

	// CPU
	if err := json.Unmarshal(data, &v.CPU); err != nil {
		return nil
	}
	if len(v.CPU) > 0 {
		*m = v
		return nil
	}

	//	Disk
	if err := json.Unmarshal(data, &v.Disk); err != nil {
		return nil
	}
	if len(v.Disk) > 0 {
		*m = v
		return nil
	}

	//	Interface
	if err := json.Unmarshal(data, &v.Interface); err != nil {
		return nil
	}
	if len(v.Interface) > 0 {
		*m = v
		return nil
	}

	//	Router
	if err := json.Unmarshal(data, &v.Router); err != nil {
		return nil
	}
	if len(v.Router) > 0 {
		*m = v
		return nil
	}

	//	Database
	if err := json.Unmarshal(data, &v.Database); err != nil {
		return nil
	}
	if len(v.Database) > 0 {
		*m = v
		return nil
	}

	//	FreeDiskSize
	if err := json.Unmarshal(data, &v.FreeDiskSize); err != nil {
		return nil
	}
	if len(v.FreeDiskSize) > 0 {
		*m = v
		return nil
	}

	//	ResponseTimeSec
	if err := json.Unmarshal(data, &v.ResponseTimeSec); err != nil {
		return nil
	}
	if len(v.ResponseTimeSec) > 0 {
		*m = v
		return nil
	}

	//	Link
	if err := json.Unmarshal(data, &v.Link); err != nil {
		return nil
	}
	if len(v.Link) > 0 {
		*m = v
		return nil
	}

	//	Connection
	if err := json.Unmarshal(data, &v.Connection); err != nil {
		return nil
	}
	if len(v.Connection) > 0 {
		*m = v
		return nil
	}

	//	LocalRouter
	if err := json.Unmarshal(data, &v.LocalRouter); err != nil {
		return nil
	}
	if len(v.LocalRouter) > 0 {
		*m = v
		return nil
	}
	return nil
}

/************************************************
 * CPU-TIME
************************************************/

// MonitorCPUTimeValue CPU-TIMEアクティビティモニタ
type MonitorCPUTimeValue struct {
	Time    time.Time // 対象時刻
	CPUTime float64
}

// MonitorCPUTimeValues CPU-TIMEアクティビティモニタ
type MonitorCPUTimeValues []*MonitorCPUTimeValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorCPUTimeValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorCPUTimeValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * Disk(Read/Write)
************************************************/

// MonitorDiskValue アクティビティモニタ
type MonitorDiskValue struct {
	Time  time.Time
	Write float64
	Read  float64
}

// MonitorDiskValues アクティビティモニタ
type MonitorDiskValues []*MonitorDiskValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorDiskValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorDiskValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * Interface(Send/Receive)
************************************************/

// MonitorInterfaceValue アクティビティモニタ
type MonitorInterfaceValue struct {
	Time    time.Time
	Send    float64
	Receive float64
}

// MonitorInterfaceValues アクティビティモニタ
type MonitorInterfaceValues []*MonitorInterfaceValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorInterfaceValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorInterfaceValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * Router(In/Out)
************************************************/

// MonitorRouterValue アクティビティモニタ
type MonitorRouterValue struct {
	Time time.Time
	In   float64
	Out  float64
}

// MonitorRouterValues アクティビティモニタ
type MonitorRouterValues []*MonitorRouterValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorRouterValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorRouterValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * Database
************************************************/

// MonitorDatabaseValue アクティビティモニタ
type MonitorDatabaseValue struct {
	Time              time.Time // 対象時刻
	TotalMemorySize   float64
	UsedMemorySize    float64
	TotalDisk1Size    float64
	UsedDisk1Size     float64
	TotalDisk2Size    float64
	UsedDisk2Size     float64
	BinlogUsedSizeKiB float64
	DelayTimeSec      float64
}

// MonitorDatabaseValues アクティビティモニタ
type MonitorDatabaseValues []*MonitorDatabaseValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorDatabaseValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorDatabaseValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * FreeDiskSize
************************************************/

// MonitorFreeDiskSizeValue アクティビティモニタ
type MonitorFreeDiskSizeValue struct {
	Time         time.Time // 対象時刻
	FreeDiskSize float64
}

// MonitorFreeDiskSizeValues アクティビティモニタ
type MonitorFreeDiskSizeValues []*MonitorFreeDiskSizeValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorFreeDiskSizeValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorFreeDiskSizeValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * ResponseTimeSec
************************************************/

// MonitorResponseTimeSecValue アクティビティモニタ
type MonitorResponseTimeSecValue struct {
	Time            time.Time // 対象時刻
	ResponseTimeSec float64
}

// MonitorResponseTimeSecValues アクティビティモニタ
type MonitorResponseTimeSecValues []*MonitorResponseTimeSecValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorResponseTimeSecValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorResponseTimeSecValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * Link(up/down)
************************************************/

// MonitorLinkValue アクティビティモニタ
type MonitorLinkValue struct {
	Time        time.Time
	UplinkBPS   float64
	DownlinkBPS float64
}

// MonitorLinkValues アクティビティモニタ
type MonitorLinkValues []*MonitorLinkValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorLinkValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorLinkValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * Connection(ProxyLB)
************************************************/

// MonitorConnectionValue アクティビティモニタ
type MonitorConnectionValue struct {
	Time              time.Time
	ActiveConnections float64
	ConnectionsPerSec float64
}

// MonitorConnectionValues アクティビティモニタ
type MonitorConnectionValues []*MonitorConnectionValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorConnectionValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorConnectionValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

/************************************************
 * LocalRouter(Receive/Send BytesPerSec)
************************************************/

// MonitorLocalRouterValue アクティビティモニタ
type MonitorLocalRouterValue struct {
	Time               time.Time
	ReceiveBytesPerSec float64
	SendBytesPerSec    float64
}

// MonitorLocalRouterValues アクティビティモニタ
type MonitorLocalRouterValues []*MonitorLocalRouterValue

// UnmarshalJSON アクティビティモニタ向けUnmarshalJSON実装
func (m *MonitorLocalRouterValues) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	var rawMonitorValues rawMonitorValues
	if err := json.Unmarshal(data, &rawMonitorValues); err != nil {
		return err
	}
	values, err := rawMonitorValues.monitorLocalRouterValues()
	if err != nil {
		return err
	}

	*m = values
	return nil
}

// rawMonitorValue アクティビティモニター
type rawMonitorValue struct {
	CPUTime            *float64 `json:"CPU-TIME,omitempty" yaml:"cpu_timme,omitempty" structs:",omitempty"`                       // CPU時間
	Write              *float64 `json:",omitempty" yaml:"write,omitempty" structs:",omitempty"`                                   // ディスク書き込み
	Read               *float64 `json:",omitempty" yaml:"read,omitempty" structs:",omitempty"`                                    // ディスク読み取り
	Receive            *float64 `json:",omitempty" yaml:"receive,omitempty" structs:",omitempty"`                                 // パケット受信
	Send               *float64 `json:",omitempty" yaml:"send,omitempty" structs:",omitempty"`                                    // パケット送信
	In                 *float64 `json:",omitempty" yaml:"in,omitempty" structs:",omitempty"`                                      // パケット受信
	Out                *float64 `json:",omitempty" yaml:"out,omitempty" structs:",omitempty"`                                     // パケット送信
	TotalMemorySize    *float64 `json:"Total-Memory-Size,omitempty" yaml:"total_memory_size,omitempty" structs:",omitempty"`      // 総メモリサイズ
	UsedMemorySize     *float64 `json:"Used-Memory-Size,omitempty" yaml:"used_memory_size,omitempty" structs:",omitempty"`        // 使用済みメモリサイズ
	TotalDisk1Size     *float64 `json:"Total-Disk1-Size,omitempty" yaml:"total_disk1_size,omitempty" structs:",omitempty"`        // 総ディスクサイズ
	UsedDisk1Size      *float64 `json:"Used-Disk1-Size,omitempty" yaml:"used_disk1_size,omitempty" structs:",omitempty"`          // 使用済みディスクサイズ
	TotalDisk2Size     *float64 `json:"Total-Disk2-Size,omitempty" yaml:"total_disk2_size,omitempty" structs:",omitempty"`        // 総ディスクサイズ
	UsedDisk2Size      *float64 `json:"Used-Disk2-Size,omitempty" yaml:"used_disk2_size,omitempty" structs:",omitempty"`          // 使用済みディスクサイズ
	BinlogUsedSizeKiB  *float64 `json:"binlogUsedSizeKiB,omitempty" yaml:"binlog_used_size_kib,omitempty" structs:",omitempty"`   // バイナリログのサイズ(レプリケーション有効時のみ、master/slave両方で利用可能)
	DelayTimeSec       *float64 `json:"delayTimeSec,omitempty" yaml:"delay_time_sec,omitempty" structs:",omitempty"`              // レプリケーション遅延時間(レプリケーション有効時のみ、slave側のみ)
	FreeDiskSize       *float64 `json:"Free-Disk-Size,omitempty" yaml:"free_disk_size,omitempty" structs:",omitempty"`            // 空きディスクサイズ(NFS)
	ResponseTimeSec    *float64 `json:"responsetimesec,omitempty" yaml:"response_time_sec,omitempty" structs:",omitempty"`        // レスポンスタイム(シンプル監視)
	UplinkBPS          *float64 `json:"UplinkBps,omitempty" yaml:"uplink_bps,omitempty" structs:",omitempty"`                     // 上り方向トラフィック
	DownlinkBPS        *float64 `json:"DownlinkBps,omitempty" yaml:"downlink_bps,omitempty" structs:",omitempty"`                 // 下り方向トラフィック
	ActiveConnections  *float64 `json:"activeConnections,omitempty" yaml:"active_connections,omitempty" structs:",omitempty"`     // アクティブコネクション(プロキシLB)
	ConnectionsPerSec  *float64 `json:"connectionsPerSec,omitempty" yaml:"connections_per_sec,omitempty" structs:",omitempty"`    // 秒間コネクション数
	ReceiveBytesPerSec *float64 `json:"receiveBytesPerSec,omitempty" yaml:"receive_bytes_per_sec,omitempty" structs:",omitempty"` // 秒間受信バイト数
	SendBytesPerSec    *float64 `json:"sendBytesPerSec,omitempty" yaml:"send_bytes_per_sec,omitempty" structs:",omitempty"`       // 秒間送信バイト数
}

// UnmarshalJSON JSONアンマーシャル(配列、オブジェクトが混在するためここで対応)
func (m *rawMonitorValue) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	tmp := &struct {
		CPUTime            *float64 `json:"CPU-TIME,omitempty" yaml:"cpu_timme,omitempty" structs:",omitempty"`
		Write              *float64 `json:",omitempty" yaml:"write,omitempty" structs:",omitempty"`
		Read               *float64 `json:",omitempty" yaml:"read,omitempty" structs:",omitempty"`
		Receive            *float64 `json:",omitempty" yaml:"receive,omitempty" structs:",omitempty"`
		Send               *float64 `json:",omitempty" yaml:"send,omitempty" structs:",omitempty"`
		In                 *float64 `json:",omitempty" yaml:"in,omitempty" structs:",omitempty"`
		Out                *float64 `json:",omitempty" yaml:"out,omitempty" structs:",omitempty"`
		TotalMemorySize    *float64 `json:"Total-Memory-Size,omitempty" yaml:"total_memory_size,omitempty" structs:",omitempty"`
		UsedMemorySize     *float64 `json:"Used-Memory-Size,omitempty" yaml:"used_memory_size,omitempty" structs:",omitempty"`
		TotalDisk1Size     *float64 `json:"Total-Disk1-Size,omitempty" yaml:"total_disk1_size,omitempty" structs:",omitempty"`
		UsedDisk1Size      *float64 `json:"Used-Disk1-Size,omitempty" yaml:"used_disk1_size,omitempty" structs:",omitempty"`
		TotalDisk2Size     *float64 `json:"Total-Disk2-Size,omitempty" yaml:"total_disk2_size,omitempty" structs:",omitempty"`
		UsedDisk2Size      *float64 `json:"Used-Disk2-Size,omitempty" yaml:"used_disk2_size,omitempty" structs:",omitempty"`
		BinlogUsedSizeKiB  *float64 `json:"binlogUsedSizeKiB,omitempty" yaml:"binlog_used_size_kib,omitempty" structs:",omitempty"`
		DelayTimeSec       *float64 `json:"delayTimeSec,omitempty" yaml:"delay_time_sec,omitempty" structs:",omitempty"`
		FreeDiskSize       *float64 `json:"Free-Disk-Size,omitempty" yaml:"free_disk_size,omitempty" structs:",omitempty"`
		ResponseTimeSec    *float64 `json:"responsetimesec,omitempty" yaml:"response_time_sec,omitempty" structs:",omitempty"`
		UplinkBPS          *float64 `json:"UplinkBps,omitempty" yaml:"uplink_bps,omitempty" structs:",omitempty"`
		DownlinkBPS        *float64 `json:"DownlinkBps,omitempty" yaml:"downlink_bps,omitempty" structs:",omitempty"`
		ActiveConnections  *float64 `json:"activeConnections,omitempty" yaml:"active_connections,omitempty" structs:",omitempty"`
		ConnectionsPerSec  *float64 `json:"connectionsPerSec,omitempty" yaml:"connections_per_sec,omitempty" structs:",omitempty"`
		ReceiveBytesPerSec *float64 `json:"receiveBytesPerSec,omitempty" yaml:"receive_bytes_per_sec,omitempty" structs:",omitempty"`
		SendBytesPerSec    *float64 `json:"sendBytesPerSec,omitempty" yaml:"send_bytes_per_sec,omitempty" structs:",omitempty"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	m.CPUTime = tmp.CPUTime
	m.Write = tmp.Write
	m.Read = tmp.Read
	m.Receive = tmp.Receive
	m.Send = tmp.Send
	m.In = tmp.In
	m.Out = tmp.Out
	m.TotalMemorySize = tmp.TotalMemorySize
	m.UsedMemorySize = tmp.UsedMemorySize
	m.TotalDisk1Size = tmp.TotalDisk1Size
	m.UsedDisk1Size = tmp.UsedDisk1Size
	m.TotalDisk2Size = tmp.TotalDisk2Size
	m.UsedDisk2Size = tmp.UsedDisk2Size
	m.BinlogUsedSizeKiB = tmp.BinlogUsedSizeKiB
	m.DelayTimeSec = tmp.DelayTimeSec
	m.FreeDiskSize = tmp.FreeDiskSize
	m.ResponseTimeSec = tmp.ResponseTimeSec
	m.UplinkBPS = tmp.UplinkBPS
	m.DownlinkBPS = tmp.DownlinkBPS
	m.ActiveConnections = tmp.ActiveConnections
	m.ConnectionsPerSec = tmp.ConnectionsPerSec
	m.ReceiveBytesPerSec = tmp.ReceiveBytesPerSec
	m.SendBytesPerSec = tmp.SendBytesPerSec

	return nil
}

type rawMonitorValues map[string]*rawMonitorValue

func (m *rawMonitorValues) monitorCPUTimeValues() (MonitorCPUTimeValues, error) {
	var values MonitorCPUTimeValues

	for k, v := range *m {
		if v.CPUTime == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorCPUTimeValue{
			Time:    time,
			CPUTime: *v.CPUTime,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorDiskValues() (MonitorDiskValues, error) {
	var values MonitorDiskValues

	for k, v := range *m {
		if v.Read == nil || v.Write == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorDiskValue{
			Time:  time,
			Read:  *v.Read,
			Write: *v.Write,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorInterfaceValues() (MonitorInterfaceValues, error) {
	var values MonitorInterfaceValues

	for k, v := range *m {
		if v.Send == nil || v.Receive == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorInterfaceValue{
			Time:    time,
			Send:    *v.Send,
			Receive: *v.Receive,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorRouterValues() (MonitorRouterValues, error) {
	var values MonitorRouterValues

	for k, v := range *m {
		if v.In == nil || v.Out == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorRouterValue{
			Time: time,
			In:   *v.In,
			Out:  *v.Out,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorDatabaseValues() (MonitorDatabaseValues, error) {
	var values MonitorDatabaseValues

	for k, v := range *m {
		if v.TotalMemorySize == nil || v.UsedMemorySize == nil ||
			v.TotalDisk1Size == nil || v.UsedDisk1Size == nil ||
			v.TotalDisk2Size == nil || v.UsedDisk2Size == nil ||
			v.BinlogUsedSizeKiB == nil || v.DelayTimeSec == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorDatabaseValue{
			Time:              time,
			TotalMemorySize:   *v.TotalMemorySize,
			UsedMemorySize:    *v.UsedMemorySize,
			TotalDisk1Size:    *v.TotalDisk1Size,
			UsedDisk1Size:     *v.UsedDisk1Size,
			TotalDisk2Size:    *v.TotalDisk2Size,
			UsedDisk2Size:     *v.UsedDisk2Size,
			BinlogUsedSizeKiB: *v.BinlogUsedSizeKiB,
			DelayTimeSec:      *v.DelayTimeSec,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorFreeDiskSizeValues() (MonitorFreeDiskSizeValues, error) {
	var values MonitorFreeDiskSizeValues

	for k, v := range *m {
		if v.FreeDiskSize == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorFreeDiskSizeValue{
			Time:         time,
			FreeDiskSize: *v.FreeDiskSize,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorResponseTimeSecValues() (MonitorResponseTimeSecValues, error) {
	var values MonitorResponseTimeSecValues

	for k, v := range *m {
		if v.ResponseTimeSec == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorResponseTimeSecValue{
			Time:            time,
			ResponseTimeSec: *v.ResponseTimeSec,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorLinkValues() (MonitorLinkValues, error) {
	var values MonitorLinkValues

	for k, v := range *m {
		if v.UplinkBPS == nil || v.DownlinkBPS == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorLinkValue{
			Time:        time,
			UplinkBPS:   *v.UplinkBPS,
			DownlinkBPS: *v.DownlinkBPS,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorConnectionValues() (MonitorConnectionValues, error) {
	var values MonitorConnectionValues

	for k, v := range *m {
		if v.ActiveConnections == nil || v.ConnectionsPerSec == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorConnectionValue{
			Time:              time,
			ActiveConnections: *v.ActiveConnections,
			ConnectionsPerSec: *v.ConnectionsPerSec,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}

func (m *rawMonitorValues) monitorLocalRouterValues() (MonitorLocalRouterValues, error) {
	var values MonitorLocalRouterValues

	for k, v := range *m {
		if v.ReceiveBytesPerSec == nil || v.SendBytesPerSec == nil {
			continue
		}
		time, err := time.Parse(time.RFC3339, k) // RFC3339 ≒ ISO8601
		if err != nil {
			return nil, err
		}
		values = append(values, &MonitorLocalRouterValue{
			Time:               time,
			ReceiveBytesPerSec: *v.ReceiveBytesPerSec,
			SendBytesPerSec:    *v.SendBytesPerSec,
		})
	}

	sort.Slice(values, func(i, j int) bool { return values[i].Time.Before(values[j].Time) })
	return values, nil
}
