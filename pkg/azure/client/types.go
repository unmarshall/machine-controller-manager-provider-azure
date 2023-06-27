package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// ConnectConfig is the configuration required for a client to connect to azure.
type ConnectConfig struct {
	// SubscriptionID
	SubscriptionID string
	TenantID       string
	ClientID       string
	ClientSecret   string
}

// CreateTokenCredential creates an azcore.TokenCredential using the ConnectConfig.
func (c ConnectConfig) CreateTokenCredential() (azcore.TokenCredential, error) {
	return azidentity.NewClientSecretCredential(c.TenantID, c.ClientID, c.ClientSecret, nil)
}

// ARMClientProvider is a factory providing methods to get clients for different resources.
type ARMClientProvider interface {
	CreateResourceGroupsClient(connectConfig ConnectConfig) (*armresources.ResourceGroupsClient, error)
	CreateVirtualMachinesClient(connectConfig ConnectConfig) (*armcompute.VirtualMachinesClient, error)
	CreateNetworkInterfacesClient(connectConfig ConnectConfig) (*armnetwork.InterfacesClient, error)
	CreateSubnetClient(connectConfig ConnectConfig) (*armnetwork.SubnetsClient, error)
	CreateDisksClient(connectConfig ConnectConfig) (*armcompute.DisksClient, error)
	CreateResourceGraphClient(connectConfig ConnectConfig) (*armresourcegraph.Client, error)
	CreateImagesClient(connectConfig ConnectConfig) (*armcompute.ImagesClient, error)
	CreateMarketPlaceAgreementsClient(connectConfig ConnectConfig) (*armmarketplaceordering.MarketplaceAgreementsClient, error)
}