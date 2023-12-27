package azure

import (
	"fmt"

	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateNetworkingProtocols(ctx *pulumi.Context, rsg *core.ResourceGroup, subnet *network.Subnet) *network.NetworkInterface {
	
	publicIp, errorPublicIP := network.NewPublicIp(ctx, "MyPublicIP", &network.PublicIpArgs{
		ResourceGroupName: rsg.Name,
		Location:          rsg.Location,
		AllocationMethod:  pulumi.String("Static"),
		Tags: pulumi.StringMap{
			"environment": pulumi.String("Production"),
		},
	})
	if errorPublicIP != nil {
		fmt.Print(errorPublicIP)
		return nil
	}

	networkInterface, errorNetwork := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
		Location:          rsg.Location,
		ResourceGroupName: rsg.Name,
		EnableIpForwarding: pulumi.Bool(true),
		IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
			&network.NetworkInterfaceIpConfigurationArgs{
				Name:                       pulumi.String("internal"),
				SubnetId:                   subnet.ID(),
				PrivateIpAddressAllocation: pulumi.String("Dynamic"),
				PublicIpAddressId: publicIp.ID(),
			},
		},
	})
	if errorNetwork != nil {
		fmt.Println(errorNetwork)
		return nil
	}

	networkSecurityGroup, errorNetworkSecurityGroup := network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
		Location:          rsg.Location,
		ResourceGroupName: rsg.Name,

		SecurityRules: network.NetworkSecurityGroupSecurityRuleArray{
			&network.NetworkSecurityGroupSecurityRuleArgs{
				Name:                     pulumi.String("test123"),
				Priority:                 pulumi.Int(100),
				Direction:                pulumi.String("Inbound"),
				Access:                   pulumi.String("Allow"),
				Protocol:                 pulumi.String("Tcp"),
				SourcePortRange:          pulumi.String("*"),
				DestinationPortRange:     pulumi.String("*"),
				SourceAddressPrefix:      pulumi.String("*"),
				DestinationAddressPrefix: pulumi.String("*"),
			},
		},
		Tags: pulumi.StringMap{
			"environment": pulumi.String("Production"),
		},
	})
	if errorNetworkSecurityGroup != nil{
		fmt.Println(errorNetworkSecurityGroup)
		return nil
	}

	_, error_association := network.NewSubnetNetworkSecurityGroupAssociation(ctx, "exampleSubnetNetworkSecurityGroupAssociation", &network.SubnetNetworkSecurityGroupAssociationArgs{
		SubnetId:               subnet.ID(),
		NetworkSecurityGroupId: networkSecurityGroup.ID(),
	})
	if error_association != nil {
		fmt.Println(error_association)
		return nil
	}

	_, errorNetworkAssociationGroup := network.NewNetworkInterfaceSecurityGroupAssociation(ctx, "exampleNetworkInterfaceSecurityGroupAssociation", &network.NetworkInterfaceSecurityGroupAssociationArgs{
		NetworkInterfaceId:     networkInterface.ID(),
		NetworkSecurityGroupId: networkSecurityGroup.ID(),
	})
	if errorNetworkAssociationGroup != nil {
		fmt.Println(errorNetworkAssociationGroup)
		return nil
	}

	return networkInterface

}

