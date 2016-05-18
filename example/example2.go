package main

import (
	"github.com/timob/ec2"
	"github.com/timob/javabind"
	"fmt"
	"log"
	"os"
	"flag"
	"errors"
	"time"
	"strings"
)

var client *ec2.ServicesEc2AmazonEC2Client

func describeInstance(instanceId string) (*ec2.ServicesEc2ModelInstance, error) {
	req := ec2.NewServicesEc2ModelDescribeInstancesRequest()
	req.SetInstanceIds([]string{instanceId})
	if r := client.DescribeInstances2(req).GetReservations(); len(r) == 1 {
		if i := r[0].GetInstances(); len(i) == 1 {
			return i[0], nil
		}
	}
	return nil, errors.New("instance id not found")
}

func getInstanceState(instanceId string) (state string, err error) {
	instance, err := describeInstance(instanceId)
	if err != nil {
		return
	}
	return instance.GetState().GetName(), nil
}

func getInstanceIpAddress(instanceId string) (ipaddr string, err error) {
	instance, err := describeInstance(instanceId)
	if err != nil {
		return
	}
	return instance.GetPublicIpAddress(), nil
}

func startInstance(instanceId string) {
	req := ec2.NewServicesEc2ModelStartInstancesRequest()
	req.SetInstanceIds([]string{instanceId})
	client.StartInstances(req)
}

func stopInstance(instanceId string) {
	req := ec2.NewServicesEc2ModelStopInstancesRequest()
	req.SetInstanceIds([]string{instanceId})
	client.StopInstances(req)
}


// prints out instance information
// make sure credentials are in ~/.aws/credentials
func main() {
	instanceId := flag.String("id", "", "instance id to start")
	start := flag.Bool("start", false, "start instance")
	stop := flag.Bool("stop", false, "stop instance")
	verbose := flag.Bool("v", false, "verbose")
	endPoint := flag.String("r", "", "amazon region")
	flag.Parse()
	
	if *instanceId == "" {
		log.Fatal("instance id must be supplied")
	}
	
	if (*start && *stop) || !(*start || *stop) {
		log.Fatal("one of -start -stop must be selected")	
	}
	
	err := javabind.SetupJVM(os.Getenv("CLASSPATH"))
	if err != nil {
		log.Fatal(err)
	}																								
	defer func() {
		if x := recover(); x != nil {
			if v, ok := x.(error); ok && strings.Index(v.Error(), "com.amazonaws") != -1 {
				log.Fatal(v)
			} else {
				panic(x)
			}		
		}		
	}()
	
	credentials := ec2.NewAuthProfileProfileCredentialsProvider().GetCredentials()
	client = ec2.NewServicesEc2AmazonEC2Client3(credentials)

	if *endPoint != "" {
		client.SetEndpoint(*endPoint)
	}

	state, err := getInstanceState(*instanceId)
	if err != nil {
		log.Fatal(err)
	}
	
	if *start && state != "running" {
		startInstance(*instanceId)
	} else if *stop && state != "stopped" {
		stopInstance(*instanceId)	
	} else {
		log.Fatalf("instance already %s", state)
	}
	
	for {
		state, err = getInstanceState(*instanceId)
		if err != nil {
			log.Fatal(err)
		}		
		if *verbose {
			log.Printf("state = %s", state)
		}
		if (*start && state == "running") || (*stop && state == "stopped") {
			break
		}
		time.Sleep(time.Second * 10)
	}
	
	if state == "running" {
		addr, err := getInstanceIpAddress(*instanceId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", addr)
	}	
}		
