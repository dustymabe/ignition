// Copyright 2018 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package partitions

import (
	"github.com/coreos/ignition/v2/tests/register"
	"github.com/coreos/ignition/v2/tests/types"
)

func init() {
	// Tests that have do not involve zeros
	register.Register(register.NegativeTest, ShouldNotExistNoWipeEntry())
	register.Register(register.NegativeTest, DoesNotMatchNoWipeEntry())
	register.Register(register.NegativeTest, ValidAndDoesNotMatchNoWipeEntry())
	register.Register(register.NegativeTest, NotThereAndDoesNotMatchNoWipeEntry())
}

func ShouldNotExistNoWipeEntry() types.Test {
	name := "partition.delete.nowipe"
	in := types.GetBaseDisk()
	out := in
	config := `{
		"ignition": {"version": "$version"},
		"storage": {
			"disks": [
			{
				"device": "$disk0",
				"partitions": [
				{
					"number": 9,
					"shouldExist": false
				}
				]
			}
			]
		}
	}`
	configMinVersion := "3.0.0"

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}

func DoesNotMatchNoWipeEntry() types.Test {
	name := "partition.match.fail"
	in := types.GetBaseDisk()
	out := in
	config := `{
		"ignition": {"version": "$version"},
		"storage": {
			"disks": [
			{
				"device": "$disk0",
				"partitions": [
				{
					"number": 9,
					"sizeMiB": 2
				}
				]
			}
			]
		}
	}`
	configMinVersion := "3.0.0"

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}

func ValidAndDoesNotMatchNoWipeEntry() types.Test {
	name := "partition.partialmatch.fail"
	in := types.GetBaseDisk()
	out := in
	config := `{
		"ignition": {"version": "$version"},
		"storage": {
			"disks": [
			{
				"device": "$disk0",
				"partitions": [
				{
					"number": 1
				},
				{
					"number": 9,
					"sizeMiB": 2
				}
				]
			}
			]
		}
	}`
	configMinVersion := "3.0.0"

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}

func NotThereAndDoesNotMatchNoWipeEntry() types.Test {
	name := "partition.partialmatch.fail.nonexistent"
	in := types.GetBaseDisk()
	out := in
	config := `{
		"ignition": {"version": "$version"},
		"storage": {
			"disks": [
			{
				"device": "$disk0",
				"partitions": [
				{
					"number": 10,
					"shouldExist": false
				},
				{
					"number": 9,
					"sizeMiB": 2
				}
				]
			}
			]
		}
	}`
	configMinVersion := "3.0.0"

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}
