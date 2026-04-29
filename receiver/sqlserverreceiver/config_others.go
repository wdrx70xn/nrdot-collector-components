// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package sqlserverreceiver // import "github.com/newrelic/nrdot-collector-components/receiver/sqlserverreceiver"

func (*Config) validateInstanceAndComputerName() error {
	return nil
}
