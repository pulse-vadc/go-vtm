// Copyright (C) 2018-2019, Pulse Secure, LLC.
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

package vtm

import "os"

var testTm *VirtualTrafficManager

func init() {
	var contactable bool
	var err *vtmErrorResponse
	testTm, contactable, err = NewVirtualTrafficManager(
		os.Getenv("VTM_BASE_URL"),
		os.Getenv("VTM_USERNAME"),
		os.Getenv("VTM_PASSWORD"),
		false,
		false,
	)
	if contactable == false {
		panic(err)
	}
}
