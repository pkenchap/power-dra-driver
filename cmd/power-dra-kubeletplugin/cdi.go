/*
 * Copyright 2025 - IBM Corporation. All rights reserved
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	"k8s.io/klog/v2"
	cdiapi "tags.cncf.io/container-device-interface/pkg/cdi"
	cdiparser "tags.cncf.io/container-device-interface/pkg/parser"
	cdispec "tags.cncf.io/container-device-interface/specs-go"
)

const (
	cdiVendor = "k8s." + DriverName
	cdiClass  = "nx"
	cdiKind   = cdiVendor + "/" + cdiClass

	cdiCommonDeviceName = "common"
)

type CDIHandler struct {
	cache *cdiapi.Cache
}

func NewCDIHandler(config *Config) (*CDIHandler, error) {
	cache, err := cdiapi.NewCache(
		cdiapi.WithSpecDirs(config.flags.cdiRoot),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create a new CDI cache: %w", err)
	}
	handler := &CDIHandler{
		cache: cache,
	}

	return handler, nil
}

func (cdi *CDIHandler) CreateCommonSpecFile() error {
	spec := &cdispec.Spec{
		Kind: cdiKind,
		Devices: []cdispec.Device{
			{
				Name: cdiCommonDeviceName,
				ContainerEdits: cdispec.ContainerEdits{
					Env: []string{
						"NX_DEVICE_CLAIM=added",
					},
					DeviceNodes: []*cdispec.DeviceNode{
						{
							Path:     "/dev/crypto/nx-gzip",
							HostPath: "/dev/crypto/nx-gzip",
						},
					},
				},
			},
		},
	}

	minVersion, err := cdiapi.MinimumRequiredVersion(spec)
	if err != nil {
		return fmt.Errorf("failed to get minimum required CDI spec version: %v", err)
	}
	spec.Version = minVersion

	specName, err := cdiapi.GenerateNameForTransientSpec(spec, cdiCommonDeviceName)
	if err != nil {
		return fmt.Errorf("failed to generate Spec name: %w", err)
	}

	return cdi.cache.WriteSpec(spec, specName)
}

func (cdi *CDIHandler) CreateClaimSpecFile(claimUID string, devices PreparedDevices) error {
	specName := cdiapi.GenerateTransientSpecName(cdiVendor, cdiClass, claimUID)

	// Only one device is used for nx-gzip
	deviceAdds := []cdispec.Device{
		{
			Name: "crypto-nx-gzip",
			ContainerEdits: cdispec.ContainerEdits{
				Env: []string{
					"NX_DEVICE_CLAIM=added",
				},
				DeviceNodes: []*cdispec.DeviceNode{
					{
						Path:     "/dev/crypto/nx-gzip",
						HostPath: "/dev/crypto/nx-gzip",
					},
				},
			},
		},
	}

	spec := &cdispec.Spec{
		Kind:    cdiKind,
		Devices: deviceAdds,
	}

	minVersion, err := cdiapi.MinimumRequiredVersion(spec)
	if err != nil {
		return fmt.Errorf("failed to get minimum required CDI spec version: %v", err)
	}
	spec.Version = minVersion

	return cdi.cache.WriteSpec(spec, specName)
}

func (cdi *CDIHandler) DeleteClaimSpecFile(claimUID string) error {
	specName := cdiapi.GenerateTransientSpecName(cdiVendor, cdiClass, claimUID)
	return cdi.cache.RemoveSpec(specName)
}

func (cdi *CDIHandler) GetClaimDevices(claimUID string, devices []string) []string {
	cdiDevices := []string{
		cdiparser.QualifiedName(cdiVendor, cdiClass, cdiCommonDeviceName),
	}

	cdiDevice := cdiparser.QualifiedName(cdiVendor, cdiClass, "crypto-nx-gzip")
	klog.V(10).Infof("cdi device: %v", cdiDevice)
	cdiDevices = append(cdiDevices, cdiDevice)

	return cdiDevices
}
