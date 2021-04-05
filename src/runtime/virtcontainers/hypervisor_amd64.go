// Copyright (c) 2021 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

import "os"

// Implementation of this function is architecture specific
func availableGuestProtection() (guestProtection, error) {
	flags, err := CPUFlags(procCPUInfo)
	if err != nil {
		return noneProtection, err
	}

	// TDX is supported and properly loaded when the firmware directory exists or `tdx` is part of the CPU flags
	if d, err := os.Stat(tdxSysFirmwareDir); (err == nil && d.IsDir()) || flags[tdxCPUFlag] {
		return tdxProtection, nil
	}
	// SEV is supported and properly loaded when `/sys/module/kvm_amd/parameters/sev` is set to `1`
	if _, err := os.Stat(sevKvmParameter); err == nil && !os.IsNotExist(err) {
		return sevProtection, nil
	}

	return noneProtection, nil
}
