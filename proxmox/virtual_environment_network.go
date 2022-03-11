/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
)

// ListNetworks retrieves a list of networks.
func (c *VirtualEnvironmentClient) ListNetworks(nodeName string, d *VirtualEnvironmentNetworkListRequestBody) ([]*VirtualEnvironmentNetworkListResponseData, error) {
	resBody := &VirtualEnvironmentNetworkListResponseBody{}
	err := c.DoRequest(hmGET, fmt.Sprintf("nodes/%s/network", url.PathEscape(nodeName)), d, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	sort.Slice(resBody.Data, func(i, j int) bool {
		return resBody.Data[i].Priority < resBody.Data[j].Priority
	})

	return resBody.Data, nil
}
