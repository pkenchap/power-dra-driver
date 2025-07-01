/*
 * Copyright 2025 - IBM Corporation. All rights reserved
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/google/uuid"
	resourceapi "k8s.io/api/resource/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

const NXGZIPCAPS = "/host-sys/devices/vio/ibm,compression-v1/nx_gzip_caps"

func enumerateAllPossibleDevices(numNx int) (AllocatableDevices, error) {
	seed := os.Getenv("NODE_NAME")
	uuids := generateUUIDs(seed, numNx)
	alldevices := make(AllocatableDevices)
	if !existsNxGzip() {
		// If nx-gzip doesn't exist, then don't return any devices.
		return alldevices, nil
	}

	for i, uuid := range uuids {
		device := resourceapi.Device{
			Name: fmt.Sprintf("nx-%d", i),
			Basic: &resourceapi.BasicDevice{
				Attributes: map[resourceapi.QualifiedName]resourceapi.DeviceAttribute{
					"index": {
						IntValue: ptr.To(int64(i)),
					},
					"uuid": {
						StringValue: ptr.To(uuid),
					},
					"model": {
						StringValue: ptr.To("LATEST-NX-MODEL"),
					},
					"driverVersion": {
						VersionValue: ptr.To("0.1.0"),
					},
				},
				Capacity: map[resourceapi.QualifiedName]resourceapi.DeviceCapacity{
					"memory": {
						Value: resource.MustParse("80Gi"),
					},
				},
			},
		}
		alldevices[device.Name] = device
	}
	return alldevices, nil
}

func generateUUIDs(seed string, count int) []string {
	rand := rand.New(rand.NewSource(hash(seed)))

	uuids := make([]string, count)
	for i := 0; i < count; i++ {
		charset := make([]byte, 16)
		rand.Read(charset)
		uuid, _ := uuid.FromBytes(charset)
		uuids[i] = "nx-" + uuid.String()
	}

	return uuids
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
	_, err := os.Stat(NXGZIPCAPS)
	if err != nil {
		klog.V(5).ErrorS(err, "Failed to detect Nest Accelerator nx-gzip feature")
		return false
	}
	return true
}
