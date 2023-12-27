package azure

import (
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/compute"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVirtualMachine(ctx *pulumi.Context, rsg *core.ResourceGroup, networkInterface *network.NetworkInterface) (*compute.VirtualMachine, error)  {
	
	vm, err := compute.NewVirtualMachine(ctx, "mainVirtualMachine", &compute.VirtualMachineArgs{
		Location:          rsg.Location,
		ResourceGroupName: rsg.Name,
		NetworkInterfaceIds: pulumi.StringArray{
			networkInterface.ID(),
		},
		Name: pulumi.String("pulumi_vm"),
		VmSize: pulumi.String("Standard_B1s"),
		StorageImageReference: &compute.VirtualMachineStorageImageReferenceArgs{
			// Publisher: pulumi.String("Canonical"),
			// Offer:     pulumi.String("0001-com-ubuntu-server-jammy"),
			// Sku:       pulumi.String("22_04-lts"),
			// Version:   pulumi.String("latest"),
			Id: pulumi.String("/subscriptions/1523240e-54cb-4a59-a661-f2f34da60861/resourceGroups/MyVMImageGroup/providers/Microsoft.Compute/images/Nginx_Package"),
		},
		StorageOsDisk: &compute.VirtualMachineStorageOsDiskArgs{
			Name:            pulumi.String("myosdisk1"),
			Caching:         pulumi.String("ReadWrite"),
			CreateOption:    pulumi.String("FromImage"),
			ManagedDiskType: pulumi.String("Standard_LRS"),
		},
		OsProfile: &compute.VirtualMachineOsProfileArgs{
			ComputerName:  pulumi.String("achyuth"),
			AdminUsername: pulumi.String("achyuth"),
			AdminPassword: pulumi.String("Password1234!"),
		},
		OsProfileLinuxConfig: &compute.VirtualMachineOsProfileLinuxConfigArgs{
			DisablePasswordAuthentication: pulumi.Bool(false),
		},
		Tags: pulumi.StringMap{
			"environment": pulumi.String("staging"),
		},
	})


	return vm, err
}