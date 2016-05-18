package main

import (
	"github.com/timob/ec2"
	"github.com/timob/javabind"
	"fmt"
	"log"
	"os"
	"flag"
)

func listRegions(client *ec2.ServicesEc2AmazonEC2Client) (r []string, err error) {
	res := client.DescribeRegions()
	for _, reg := range res.GetRegions() {
		r = append(r, reg.GetEndpoint())
	}
	return
}

// prints out instance information
// make sure credentials are in ~/.aws/credentials
func main() {
	endPoint := flag.String("r", "", "amazon region (can use \"all\")")
	listRegionsOpt := flag.Bool("listregions", false, "list amazon regions")
	flag.Parse()
		
	err := javabind.SetupJVM(os.Getenv("CLASSPATH"))
	if err != nil {
		log.Fatal(err)
	}																								
	
	
	credentials := ec2.NewAuthProfileProfileCredentialsProvider().GetCredentials()
	client := ec2.NewServicesEc2AmazonEC2Client3(credentials)

	var endPointList []string
	if *endPoint == "all" {
		endPointList, err = listRegions(client)
		if err != nil {
			log.Fatal(err)
		}
	} else if *endPoint != "" {
		endPointList = append(endPointList, *endPoint)
		//client.SetEndpoint(*endPoint)
	} else {
		var regions *ec2.RegionsRegions
		endPointList = append(endPointList, "ec2." + regions.DEFAULT_REGION().GetName() + ".amazonaws.com")
	}

	if *listRegionsOpt {
		l, err := listRegions(client)
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range l {
			fmt.Printf("%s\n", r)
		}
		return
	}

	for _, s := range endPointList {
		client.SetEndpoint(s)
		instances := client.DescribeInstances()

		for _, r := range instances.GetReservations() {
			for _, i := range r.GetInstances() {
				fmt.Printf("%s %s %s %s", i.GetInstanceId(), i.GetInstanceType(), i.GetPlacement().GetAvailabilityZone(), i.GetState().GetName())
				ipAddr := i.GetPublicIpAddress()
				if ipAddr != "" {
					fmt.Printf(" %s", ipAddr)
				}
				fmt.Printf("\n")
			}
		}
	}
}
