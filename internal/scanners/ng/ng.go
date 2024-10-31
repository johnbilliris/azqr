// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package ng

import (
	"github.com/Azure/azqr/internal/azqr"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// NatGatewayScanner - Scanner for NAT Gateway
type NatGatewayScanner struct {
	config *azqr.ScannerConfig
	client *armnetwork.NatGatewaysClient
}

// Init - Initializes the NAT Gateway Scanner
func (a *NatGatewayScanner) Init(config *azqr.ScannerConfig) error {
	a.config = config
	var err error
	a.client, err = armnetwork.NewNatGatewaysClient(config.SubscriptionID, config.Cred, config.ClientOptions)
	return err
}

// Scan - Scans all NAT Gateway in a Resource Group
func (c *NatGatewayScanner) Scan(scanContext *azqr.ScanContext) ([]azqr.AzqrServiceResult, error) {
	azqr.LogSubscriptionScan(c.config.SubscriptionID, c.ResourceTypes()[0])

	svcs, err := c.list()
	if err != nil {
		return nil, err
	}
	engine := azqr.RecommendationEngine{}
	rules := c.GetRecommendations()
	results := []azqr.AzqrServiceResult{}

	for _, w := range svcs {
		rr := engine.EvaluateRecommendations(rules, w, scanContext)

		results = append(results, azqr.AzqrServiceResult{
			SubscriptionID:   c.config.SubscriptionID,
			SubscriptionName: c.config.SubscriptionName,
			ResourceGroup:    azqr.GetResourceGroupFromResourceID(*w.ID),
			ServiceName:      *w.Name,
			Type:             *w.Type,
			Location:         *w.Location,
			Recommendations:  rr,
		})
	}
	return results, nil
}

func (c *NatGatewayScanner) list() ([]*armnetwork.NatGateway, error) {
	pager := c.client.NewListAllPager(nil)

	svcs := make([]*armnetwork.NatGateway, 0)
	for pager.More() {
		resp, err := pager.NextPage(c.config.Ctx)
		if err != nil {
			return nil, err
		}
		svcs = append(svcs, resp.Value...)
	}
	return svcs, nil
}

func (a *NatGatewayScanner) ResourceTypes() []string {
	return []string{"Microsoft.Network/natGateways"}
}