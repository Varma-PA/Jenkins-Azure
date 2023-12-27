package main

import (
	"infrastructure/azure"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		rsg, rsg_err := azure.CreateResourceGroup(ctx);

		if(rsg_err != nil){
			return rsg_err
		} 

		_, subnet, vNetError, subnetError := azure.CreateVirtualNetworkAndSubnets(ctx, rsg)
		

		networkInterface := azure.CreateNetworkingProtocols(ctx, rsg, subnet)

		_, vmError := azure.CreateVirtualMachine(ctx, rsg, networkInterface)	

		if(vNetError != nil) {
			return vNetError
		}
		
		if(subnetError != nil){
			return subnetError
		}

		if(vmError != nil){
			return vmError
		}
		
		return nil
	})
}
