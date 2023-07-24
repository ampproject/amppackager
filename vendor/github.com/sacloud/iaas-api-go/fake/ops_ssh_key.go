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

package fake

import (
	"context"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"golang.org/x/crypto/ssh"
)

var (
	// GeneratedPublicKey ダミー公開鍵
	GeneratedPublicKey = `ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAs7YFtxjGrI49MCBnSFbUPxqz0e5HSGQPnLlPJ0u/9w4WLpoOZYmoQDTMfuFA61qv+0dp5mpMZPj3f5YEGlwUFKPy3Cmrp0ub1nYDb7n62s+Xf68TNvbVgQMLF0xdOaWxdRsQwmH8lOWan1Ubc8iwfOa3TNGwOzGLMjdW3PiJ7hcE7nFqnmbQUabHWow8G6JYDHKyjAdpz+edK8u+LY0iEP8M8VAjRJKJVg4p1/oDjHFKI0qjfjitKzoLm5FGaFv8afH2WQSpu/2To7d/RaLhfoMZsUReLSxeDnQkKGERXrAywTHnFu60cOaT3EvaAhP1H3BPj2LESm8M4ja9FaARnQ== `
	// GeneratedPrivateKey ダミー秘密鍵
	GeneratedPrivateKey = `-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: DES-EDE3-CBC,44767E1C8C7F1420\n\nNjG0a+HcHUnly4nCF4EjIzWk0A7QleHBXK2NFlaeEJJO5i+cK6fkhW5lkWyA/MoP\nwuipGVM209LUcax/EABmxUcQUpuu28uMcIaJJBdd8PF7WsLgMJJRF82YhvKoyvPV\nmGCn+LXlJEsIVKYR+KEGLak2G3KVpNsKozgvoE5ytd/J+AP09w56pSwXmIPnRBiV\nV+adP8if1eWJZvj2sBfaHlYZiw374xWKvCGS0h/Ezaj1CQICk3CyNRxI3jYPCwte\n706tvPLVh0fk9PMaiZuvqGAi49z+AP7OkOzfnoZVHh8BPYNXVln0IZvUogPdhmTP\nYxFEcwvJF9WQGkA2Uz6zdX5LE60+Nw6qsL6f95hvAe4f9RC7MHnnhJKan68V7lKg\nm/659evXjUlMs5XMPDaQOWd4FKva4TRlqdjU4j+KcMlj+JICnqUfRb0mge8lXg91\nYFNhlM1JTEj7LvQQM54Af/lnytZdR8RvHluC4XMUurPVrJVouBO7ssW518vkbuzk\nqgAOvVDBJOA5O446+aL2lBJEMnP3/KCgBOHZp0hJzRCmFXmLZ1SWCSN6J9mB9O40\nM9EoVmjZVWGhlbM2RM79QUtwavYbyLVxRuMv5/TTjA4SnYnTqW3SLPfhnClkk1u5\nqM9ILoZgFVV4ENAI48gdIonr209nSv5wCD5knkMEC+U2r3DOypL9IjAnmR5VrxOy\ngKnr9cqYDT0gvP6jbyxLmNxGrzFfxogiyxQKGj61wAwkgTtOc9rK4j2w0UJ1sN6f\npkOoqtchEGR+sb+msiCv54OqLK87/XnEQKZFqo1JUjw4HXvy3Mz/cq+7zbI2ZhKh\nOnkZPSNIUq4dFilCIoRqIoUb/hNAsSAWZjvuSo24+OaQK6I5D+Q42ZMg+dJeeITB\nlTNAwbOYLzALrgbOfN10Xfid2NL3hP28b/8DCuU1NpQRBPTl8S1DAbKQskeTqVIj\n0x98ojlrDmcBchHP9B6MCBtt/1BiwCJzq0sAUtuvDXGtKSHBh7K857n1nQm+9ru9\nvfxi+zL1MPsObNChLc9YAqL63RCGDKR8u09I3MokpifI567r7Y68O1Y7rNTwtOrM\nxGEq+/7Xqa5Eu7Vj/yHp5abPqdn83ws8UueR7ufTmz2a2W7tA6mo/ObYYNktyP7A\nF4Gdl/6hDQZEmMr6u5zeiFKqAi5ZlWqfx7kYxT+Cu/qBfBf9Vs7OcsU83U06Wnye\nAYomPxideCkWYB+qjxU9+DViPg0fbbH6WDcqW2WUrMvshGaPG1pnVqhchzoDSo36\nvMvUXq9M21GLWL+69vNXpCA+pIuke4+yiskbTKWR36+AxRTo5Em6KiXwSK8icdQH\nExELbH3WYaw6C6M9+BEmzqUsLerqzvT5IKDlAxCSwOMxSUr2uJpr6prHkL8jvCWg\nX4Ch4HISWm43iGucH2Y6AWzHU1dZYn3zXisQ12D+kq+Lp55a+RHWHfuzujJAk9li\nsn3TaC1cHLXgv9aQpRiAmGO2Lb81gpcaNvZ1p3Tg3RfTMml1moogk2dYTpE2WfG8\nU1eVQG1McqeIoduvBL2YAWdOW4GS+bh3Fv6dfT7wq3bSDmdX4Gy8Og==\n-----END RSA PRIVATE KEY-----\n`
	// GeneratedFingerprint ダミーフィンガープリント
	GeneratedFingerprint = "79:d7:ac:b8:cf:cf:01:44:b2:19:ba:d4:82:fd:c4:2d"
)

// Find is fake implementation
func (o *SSHKeyOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.SSHKeyFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.SSHKey
	for _, res := range results {
		dest := &iaas.SSHKey{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.SSHKeyFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		SSHKeys: values,
	}, nil
}

// Create is fake implementation
func (o *SSHKeyOp) Create(ctx context.Context, param *iaas.SSHKeyCreateRequest) (*iaas.SSHKey, error) {
	result := &iaas.SSHKey{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	pk, _, _, _, err := ssh.ParseAuthorizedKey([]byte(result.PublicKey))
	if err != nil {
		return nil, err
	}
	result.Fingerprint = ssh.FingerprintLegacyMD5(pk)

	putSSHKey(iaas.APIDefaultZone, result)
	return result, nil
}

// Generate is fake implementation
func (o *SSHKeyOp) Generate(ctx context.Context, param *iaas.SSHKeyGenerateRequest) (*iaas.SSHKeyGenerated, error) {
	key := &iaas.SSHKey{}
	copySameNameField(param, key)
	fill(key, fillID, fillCreatedAt)

	result := &iaas.SSHKeyGenerated{}
	copySameNameField(key, result)

	result.PublicKey = GeneratedPublicKey
	result.PrivateKey = GeneratedPrivateKey
	result.Fingerprint = GeneratedFingerprint

	putSSHKey(iaas.APIDefaultZone, key)
	return result, nil
}

// Read is fake implementation
func (o *SSHKeyOp) Read(ctx context.Context, id types.ID) (*iaas.SSHKey, error) {
	value := getSSHKeyByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.SSHKey{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SSHKeyOp) Update(ctx context.Context, id types.ID, param *iaas.SSHKeyUpdateRequest) (*iaas.SSHKey, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)

	putSSHKey(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *SSHKeyOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}
