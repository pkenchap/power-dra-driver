/*
 * Copyright 2025 - IBM Corporation. All rights reserved
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"os"

	resourceapi "k8s.io/api/resource/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

const CapsLocation = "/host-sys/devices/vio/ibm,compression-v1/nx_gzip_caps"
const UpperLimit = "200"
const Uuid = "b2ccae49-efdd-4d90-bc8e-6fec4a2b19f7"

func enumerateAllPossibleDevices(numNx int) (AllocatableDevices, error) {
	alldevices := make(AllocatableDevices)

	// If nx-gzip doesn't exist, then don't return any devices.
	if !existsNxGzip() {
		return alldevices, nil
	}

	// prepopulate with a single nx-gzip
	device := resourceapi.Device{
		Name: "cryptonxgzip",
		Basic: &resourceapi.BasicDevice{
			Attributes: map[resourceapi.QualifiedName]resourceapi.DeviceAttribute{
				"index": {
					IntValue: ptr.To(int64(0)),
				},
				"uuid": {
					StringValue: ptr.To(Uuid),
				},
				"model": {
					StringValue: ptr.To("LATEST-NX-MODEL"),
				},
				"driverVersion": {
					VersionValue: ptr.To("0.1.0"),
				},
			},
			Capacity: map[resourceapi.QualifiedName]resourceapi.DeviceCapacity{
				"nx-gzip": {
					Value: resource.MustParse(UpperLimit),
				},
			},
		},
	}
	alldevices[device.Name] = device

	return alldevices, nil
}

func hash(s string) int64 {
	h := int64(0)
	for _, c := range s {
		h = 31*h + int64(c)
	}
	return h
}

// Detect NXGZIPCAPS exists
func existsNxGzip() bool {
	_, err := os.Stat(CapsLocation)
	if err != nil {
		klog.V(5).ErrorS(err, "Failed to detect Nest Accelerator nx-gzip feature")
		return false
	}
	return true
}
