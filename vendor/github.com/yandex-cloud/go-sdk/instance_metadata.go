// Copyright (c) 2019 Yandex LLC. All rights reserved.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package ycsdk

// InstanceMetadataOverrideEnvVar is a name of environment variable which contains override for a default value
const InstanceMetadataOverrideEnvVar = "YC_METADATA_ADDR"

// InstanceMetadataAddr is address at  the metadata server is accessible from inside the virtual machine.
// See doc for details: https://cloud.yandex.com/docs/compute/operations/vm-info/get-info#inside-instance
const InstanceMetadataAddr = "169.254.169.254"

// TODO(skipor): SDK methods for getting information from inside a VM.
// https://cloud.yandex.com/docs/compute/operations/vm-info/get-info#inside-instance
