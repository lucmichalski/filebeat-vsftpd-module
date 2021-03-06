// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package vsftpd

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "vsftpd", asset.ModuleFieldsPri, AssetVsftpd); err != nil {
		panic(err)
	}
}

// AssetVsftpd returns asset data.
// This is the base64 encoded gzipped contents of module/vsftpd.
func AssetVsftpd() string {
	return "eJzEkEFugzAQRfc+xRd7LuBFb1CpV7CYAVkY28JDC7evaAlxghfJIsTL5/Gf912j50XjO7USSQFixbFG9Q8qBRCnZrRRbPAaHwrANo3PQJNjBbSWHSX9d1fDm4GzxPXIElmjG8MUN1JIvc3Js8gI7/Du7ZdjkxiGKOfZNM9miGunbHJjh0UudC/dU/qJUvGrUbSU0WedHrW6eLngu4JEEiNTOtGj5+UnjFRQMc0h+l0qNp6gsS/zbTixtfAs6jcAAP//VkkqUA=="
}
