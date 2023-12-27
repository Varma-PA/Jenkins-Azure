package azure

import (
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVirtualNetworkAndSubnets(ctx *pulumi.Context, rsg *core.ResourceGroup) (*network.VirtualNetwork, *network.Subnet, error, error)  {
	
		//  Creating the Virtual Network
		vnet, errorNetwork := network.NewVirtualNetwork(ctx, "virtualNetwork", &network.VirtualNetworkArgs{
			AddressSpaces: pulumi.StringArray{
				pulumi.String("10.0.0.0/16"),
			},
			FlowTimeoutInMinutes: pulumi.Int(10),
			Location:             pulumi.String("eastus"),
			ResourceGroupName:    rsg.Name,
			Name:   			  pulumi.String("test-vnet"),
		})

		subnet, errorSubnet := network.NewSubnet(ctx, "subnet", &network.SubnetArgs{
			AddressPrefixes: pulumi.StringArray{
				pulumi.String("10.0.0.0/24"),
			},
			ResourceGroupName:  rsg.Name,
			Name:         pulumi.String("Public_Subnet"),
			VirtualNetworkName: vnet.Name,
			
		})


		return vnet, subnet, errorNetwork, errorSubnet

}