
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

// main uses the AWS SDK for Go V2 to create an Amazon Simple Storage Service
// (Amazon S3) client and list up to 10 hostzones in your account.
// This example uses the default settings specified in your shared credentials
// and config files.
func main() {
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	r53Client := route53.NewFromConfig(sdkConfig)
	count := 10
	fmt.Printf("Let's list up to %v hostzones for your account.\n", count)
	result, err := r53Client.ListHostedZones(ctx, &route53.ListHostedZonesInput{})
	if err != nil {
		fmt.Printf("Couldn't list hostzones for your account. Here's why: %v\n", err)
		return
	}
	if len(result.HostedZones) == 0 {
		fmt.Println("You don't have any hostzones!")
	} else {
		if count > len(result.HostedZones) {
			count = len(result.HostedZones)
		}
		for _, hz := range result.HostedZones[:count] {
			fmt.Printf("\t%v\n", *hz.Name)
		}
	}
}


