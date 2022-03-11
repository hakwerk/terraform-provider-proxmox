/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	mkDataSourceVirtualEnvironmentNetworksAddresses  = "addresses"
	mkDataSourceVirtualEnvironmentNetworksNodeName   = "node_name"
	mkDataSourceVirtualEnvironmentNetworksPriorities = "priorities"
	mkDataSourceVirtualEnvironmentNetworksTypes      = "types"
)

func dataSourceVirtualEnvironmentNetworks() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkDataSourceVirtualEnvironmentNetworksAddresses: {
				Type:        schema.TypeList,
				Description: "Address of this network interface",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentNetworksNodeName: {
				Type:        schema.TypeString,
				Description: "The node name",
				Required:    true,
			},
			mkDataSourceVirtualEnvironmentNetworksPriorities: {
				Type:        schema.TypeList,
				Description: "The network interface priority",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			mkDataSourceVirtualEnvironmentNetworksTypes: {
				Type:        schema.TypeList,
				Description: "The network interface type",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
		Read: dataSourceVirtualEnvironmentNetworkRead,
	}
}

func dataSourceVirtualEnvironmentNetworkRead(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkDataSourceVirtualEnvironmentNetworksNodeName).(string)
	list, err := veClient.ListNetworks(nodeName, nil)

	if err != nil {
		return err
	}

	addresses := make([]interface{}, len(list))
	priorities := make([]interface{}, len(list))
	types := make([]interface{}, len(list))

	d.SetId(fmt.Sprintf("%s_networks", nodeName))

	d.Set(mkDataSourceVirtualEnvironmentNetworksAddresses, addresses)
	d.Set(mkDataSourceVirtualEnvironmentNetworksPriorities, priorities)
	d.Set(mkDataSourceVirtualEnvironmentNetworksTypes, types)

	return nil
}
