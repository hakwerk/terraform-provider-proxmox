/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

//import (
//)

// VirtualEnvironmentNetworkListRequestBody contains the body for a network list request.
type VirtualEnvironmentNetworkListRequestBody struct {
	Type *string `json:"type,omitempty" url:"type,omitempty"`
}

// VirtualEnvironmentNetworkListResponseBody contains the body from a network list response.
type VirtualEnvironmentNetworkListResponseBody struct {
	Data []*VirtualEnvironmentNetworkListResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentNetworkListResponseData contains the data from a network list response.
type VirtualEnvironmentNetworkListResponseData struct {
	Address  string `json:"address,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Type     string `json:"type,omitempty"`
}
