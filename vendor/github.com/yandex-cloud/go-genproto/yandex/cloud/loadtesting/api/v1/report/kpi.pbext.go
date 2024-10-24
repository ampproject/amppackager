// Code generated by protoc-gen-goext. DO NOT EDIT.

package report

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/common"
)

func (m *Kpi) SetSelector(v *KpiSelector) {
	m.Selector = v
}

func (m *Kpi) SetThreshold(v *KpiThreshold) {
	m.Threshold = v
}

func (m *KpiThreshold) SetValue(v float64) {
	m.Value = v
}

func (m *KpiThreshold) SetComparison(v Comparison) {
	m.Comparison = v
}

type KpiSelector_Kind = isKpiSelector_Kind

func (m *KpiSelector) SetKind(v KpiSelector_Kind) {
	m.Kind = v
}

func (m *KpiSelector) SetResponseTime(v *KpiSelector_ResponseTime) {
	m.Kind = &KpiSelector_ResponseTime_{
		ResponseTime: v,
	}
}

func (m *KpiSelector) SetInstances(v *KpiSelector_Instances) {
	m.Kind = &KpiSelector_Instances_{
		Instances: v,
	}
}

func (m *KpiSelector) SetImbalanceRps(v *KpiSelector_ImbalanceRps) {
	m.Kind = &KpiSelector_ImbalanceRps_{
		ImbalanceRps: v,
	}
}

func (m *KpiSelector) SetProtocolCodesAbsolute(v *KpiSelector_ProtocolCodesAbsolute) {
	m.Kind = &KpiSelector_ProtocolCodesAbsolute_{
		ProtocolCodesAbsolute: v,
	}
}

func (m *KpiSelector) SetProtocolCodesRelative(v *KpiSelector_ProtocolCodesRelative) {
	m.Kind = &KpiSelector_ProtocolCodesRelative_{
		ProtocolCodesRelative: v,
	}
}

func (m *KpiSelector) SetNetworkCodesAbsolute(v *KpiSelector_NetworkCodesAbsolute) {
	m.Kind = &KpiSelector_NetworkCodesAbsolute_{
		NetworkCodesAbsolute: v,
	}
}

func (m *KpiSelector) SetNetworkCodesRelative(v *KpiSelector_NetworkCodesRelative) {
	m.Kind = &KpiSelector_NetworkCodesRelative_{
		NetworkCodesRelative: v,
	}
}

func (m *KpiSelector_ResponseTime) SetQuantile(v common.QuantileType) {
	m.Quantile = v
}

func (m *KpiSelector_Instances) SetAgg(v Aggregation) {
	m.Agg = v
}

func (m *KpiSelector_ProtocolCodesAbsolute) SetCodesPatterns(v []string) {
	m.CodesPatterns = v
}

func (m *KpiSelector_ProtocolCodesRelative) SetCodesPatterns(v []string) {
	m.CodesPatterns = v
}

func (m *KpiSelector_NetworkCodesAbsolute) SetCodesPatterns(v []string) {
	m.CodesPatterns = v
}

func (m *KpiSelector_NetworkCodesRelative) SetCodesPatterns(v []string) {
	m.CodesPatterns = v
}