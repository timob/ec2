package ec2

import "github.com/timob/javabind"

type ServicesEc2AmazonEC2ClientInterface interface {
	AmazonWebServiceClientInterface

	// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult acceptVpcPeeringConnection(com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionRequest)
	AcceptVpcPeeringConnection2(a ServicesEc2ModelAcceptVpcPeeringConnectionRequestInterface) *ServicesEc2ModelAcceptVpcPeeringConnectionResult

	// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult acceptVpcPeeringConnection()
	AcceptVpcPeeringConnection() *ServicesEc2ModelAcceptVpcPeeringConnectionResult

	// public com.amazonaws.services.ec2.model.AllocateAddressResult allocateAddress(com.amazonaws.services.ec2.model.AllocateAddressRequest)
	AllocateAddress2(a ServicesEc2ModelAllocateAddressRequestInterface) *ServicesEc2ModelAllocateAddressResult

	// public com.amazonaws.services.ec2.model.AllocateAddressResult allocateAddress()
	AllocateAddress() *ServicesEc2ModelAllocateAddressResult

	// public com.amazonaws.services.ec2.model.AllocateHostsResult allocateHosts(com.amazonaws.services.ec2.model.AllocateHostsRequest)
	AllocateHosts(a ServicesEc2ModelAllocateHostsRequestInterface) *ServicesEc2ModelAllocateHostsResult

	// public void assignPrivateIpAddresses(com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest)
	AssignPrivateIpAddresses(a ServicesEc2ModelAssignPrivateIpAddressesRequestInterface) 

	// public com.amazonaws.services.ec2.model.AssociateAddressResult associateAddress(com.amazonaws.services.ec2.model.AssociateAddressRequest)
	AssociateAddress(a ServicesEc2ModelAssociateAddressRequestInterface) *ServicesEc2ModelAssociateAddressResult

	// public void associateDhcpOptions(com.amazonaws.services.ec2.model.AssociateDhcpOptionsRequest)
	AssociateDhcpOptions(a ServicesEc2ModelAssociateDhcpOptionsRequestInterface) 

	// public com.amazonaws.services.ec2.model.AssociateRouteTableResult associateRouteTable(com.amazonaws.services.ec2.model.AssociateRouteTableRequest)
	AssociateRouteTable(a ServicesEc2ModelAssociateRouteTableRequestInterface) *ServicesEc2ModelAssociateRouteTableResult

	// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcResult attachClassicLinkVpc(com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest)
	AttachClassicLinkVpc(a ServicesEc2ModelAttachClassicLinkVpcRequestInterface) *ServicesEc2ModelAttachClassicLinkVpcResult

	// public void attachInternetGateway(com.amazonaws.services.ec2.model.AttachInternetGatewayRequest)
	AttachInternetGateway(a ServicesEc2ModelAttachInternetGatewayRequestInterface) 

	// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult attachNetworkInterface(com.amazonaws.services.ec2.model.AttachNetworkInterfaceRequest)
	AttachNetworkInterface(a ServicesEc2ModelAttachNetworkInterfaceRequestInterface) *ServicesEc2ModelAttachNetworkInterfaceResult

	// public com.amazonaws.services.ec2.model.AttachVolumeResult attachVolume(com.amazonaws.services.ec2.model.AttachVolumeRequest)
	AttachVolume(a ServicesEc2ModelAttachVolumeRequestInterface) *ServicesEc2ModelAttachVolumeResult

	// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult attachVpnGateway(com.amazonaws.services.ec2.model.AttachVpnGatewayRequest)
	AttachVpnGateway(a ServicesEc2ModelAttachVpnGatewayRequestInterface) *ServicesEc2ModelAttachVpnGatewayResult

	// public void authorizeSecurityGroupEgress(com.amazonaws.services.ec2.model.AuthorizeSecurityGroupEgressRequest)
	AuthorizeSecurityGroupEgress(a ServicesEc2ModelAuthorizeSecurityGroupEgressRequestInterface) 

	// public void authorizeSecurityGroupIngress(com.amazonaws.services.ec2.model.AuthorizeSecurityGroupIngressRequest)
	AuthorizeSecurityGroupIngress(a ServicesEc2ModelAuthorizeSecurityGroupIngressRequestInterface) 

	// public com.amazonaws.services.ec2.model.BundleInstanceResult bundleInstance(com.amazonaws.services.ec2.model.BundleInstanceRequest)
	BundleInstance(a ServicesEc2ModelBundleInstanceRequestInterface) *ServicesEc2ModelBundleInstanceResult

	// public com.amazonaws.services.ec2.model.CancelBundleTaskResult cancelBundleTask(com.amazonaws.services.ec2.model.CancelBundleTaskRequest)
	CancelBundleTask(a ServicesEc2ModelCancelBundleTaskRequestInterface) *ServicesEc2ModelCancelBundleTaskResult

	// public void cancelConversionTask(com.amazonaws.services.ec2.model.CancelConversionTaskRequest)
	CancelConversionTask(a ServicesEc2ModelCancelConversionTaskRequestInterface) 

	// public void cancelExportTask(com.amazonaws.services.ec2.model.CancelExportTaskRequest)
	CancelExportTask(a ServicesEc2ModelCancelExportTaskRequestInterface) 

	// public com.amazonaws.services.ec2.model.CancelImportTaskResult cancelImportTask(com.amazonaws.services.ec2.model.CancelImportTaskRequest)
	CancelImportTask2(a ServicesEc2ModelCancelImportTaskRequestInterface) *ServicesEc2ModelCancelImportTaskResult

	// public com.amazonaws.services.ec2.model.CancelImportTaskResult cancelImportTask()
	CancelImportTask() *ServicesEc2ModelCancelImportTaskResult

	// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult cancelReservedInstancesListing(com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest)
	CancelReservedInstancesListing(a ServicesEc2ModelCancelReservedInstancesListingRequestInterface) *ServicesEc2ModelCancelReservedInstancesListingResult

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult cancelSpotFleetRequests(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest)
	CancelSpotFleetRequests(a ServicesEc2ModelCancelSpotFleetRequestsRequestInterface) *ServicesEc2ModelCancelSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult cancelSpotInstanceRequests(com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest)
	CancelSpotInstanceRequests(a ServicesEc2ModelCancelSpotInstanceRequestsRequestInterface) *ServicesEc2ModelCancelSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult confirmProductInstance(com.amazonaws.services.ec2.model.ConfirmProductInstanceRequest)
	ConfirmProductInstance(a ServicesEc2ModelConfirmProductInstanceRequestInterface) *ServicesEc2ModelConfirmProductInstanceResult

	// public com.amazonaws.services.ec2.model.CopyImageResult copyImage(com.amazonaws.services.ec2.model.CopyImageRequest)
	CopyImage(a ServicesEc2ModelCopyImageRequestInterface) *ServicesEc2ModelCopyImageResult

	// public com.amazonaws.services.ec2.model.CopySnapshotResult copySnapshot(com.amazonaws.services.ec2.model.CopySnapshotRequest)
	CopySnapshot(a ServicesEc2ModelCopySnapshotRequestInterface) *ServicesEc2ModelCopySnapshotResult

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult createCustomerGateway(com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest)
	CreateCustomerGateway(a ServicesEc2ModelCreateCustomerGatewayRequestInterface) *ServicesEc2ModelCreateCustomerGatewayResult

	// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult createDhcpOptions(com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest)
	CreateDhcpOptions(a ServicesEc2ModelCreateDhcpOptionsRequestInterface) *ServicesEc2ModelCreateDhcpOptionsResult

	// public com.amazonaws.services.ec2.model.CreateFlowLogsResult createFlowLogs(com.amazonaws.services.ec2.model.CreateFlowLogsRequest)
	CreateFlowLogs(a ServicesEc2ModelCreateFlowLogsRequestInterface) *ServicesEc2ModelCreateFlowLogsResult

	// public com.amazonaws.services.ec2.model.CreateImageResult createImage(com.amazonaws.services.ec2.model.CreateImageRequest)
	CreateImage(a ServicesEc2ModelCreateImageRequestInterface) *ServicesEc2ModelCreateImageResult

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult createInstanceExportTask(com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest)
	CreateInstanceExportTask(a ServicesEc2ModelCreateInstanceExportTaskRequestInterface) *ServicesEc2ModelCreateInstanceExportTaskResult

	// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult createInternetGateway(com.amazonaws.services.ec2.model.CreateInternetGatewayRequest)
	CreateInternetGateway2(a ServicesEc2ModelCreateInternetGatewayRequestInterface) *ServicesEc2ModelCreateInternetGatewayResult

	// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult createInternetGateway()
	CreateInternetGateway() *ServicesEc2ModelCreateInternetGatewayResult

	// public com.amazonaws.services.ec2.model.CreateKeyPairResult createKeyPair(com.amazonaws.services.ec2.model.CreateKeyPairRequest)
	CreateKeyPair(a ServicesEc2ModelCreateKeyPairRequestInterface) *ServicesEc2ModelCreateKeyPairResult

	// public com.amazonaws.services.ec2.model.CreateNatGatewayResult createNatGateway(com.amazonaws.services.ec2.model.CreateNatGatewayRequest)
	CreateNatGateway(a ServicesEc2ModelCreateNatGatewayRequestInterface) *ServicesEc2ModelCreateNatGatewayResult

	// public com.amazonaws.services.ec2.model.CreateNetworkAclResult createNetworkAcl(com.amazonaws.services.ec2.model.CreateNetworkAclRequest)
	CreateNetworkAcl(a ServicesEc2ModelCreateNetworkAclRequestInterface) *ServicesEc2ModelCreateNetworkAclResult

	// public void createNetworkAclEntry(com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest)
	CreateNetworkAclEntry(a ServicesEc2ModelCreateNetworkAclEntryRequestInterface) 

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult createNetworkInterface(com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest)
	CreateNetworkInterface(a ServicesEc2ModelCreateNetworkInterfaceRequestInterface) *ServicesEc2ModelCreateNetworkInterfaceResult

	// public void createPlacementGroup(com.amazonaws.services.ec2.model.CreatePlacementGroupRequest)
	CreatePlacementGroup(a ServicesEc2ModelCreatePlacementGroupRequestInterface) 

	// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult createReservedInstancesListing(com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest)
	CreateReservedInstancesListing(a ServicesEc2ModelCreateReservedInstancesListingRequestInterface) *ServicesEc2ModelCreateReservedInstancesListingResult

	// public com.amazonaws.services.ec2.model.CreateRouteResult createRoute(com.amazonaws.services.ec2.model.CreateRouteRequest)
	CreateRoute(a ServicesEc2ModelCreateRouteRequestInterface) *ServicesEc2ModelCreateRouteResult

	// public com.amazonaws.services.ec2.model.CreateRouteTableResult createRouteTable(com.amazonaws.services.ec2.model.CreateRouteTableRequest)
	CreateRouteTable(a ServicesEc2ModelCreateRouteTableRequestInterface) *ServicesEc2ModelCreateRouteTableResult

	// public com.amazonaws.services.ec2.model.CreateSecurityGroupResult createSecurityGroup(com.amazonaws.services.ec2.model.CreateSecurityGroupRequest)
	CreateSecurityGroup(a ServicesEc2ModelCreateSecurityGroupRequestInterface) *ServicesEc2ModelCreateSecurityGroupResult

	// public com.amazonaws.services.ec2.model.CreateSnapshotResult createSnapshot(com.amazonaws.services.ec2.model.CreateSnapshotRequest)
	CreateSnapshot(a ServicesEc2ModelCreateSnapshotRequestInterface) *ServicesEc2ModelCreateSnapshotResult

	// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult createSpotDatafeedSubscription(com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest)
	CreateSpotDatafeedSubscription(a ServicesEc2ModelCreateSpotDatafeedSubscriptionRequestInterface) *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult

	// public com.amazonaws.services.ec2.model.CreateSubnetResult createSubnet(com.amazonaws.services.ec2.model.CreateSubnetRequest)
	CreateSubnet(a ServicesEc2ModelCreateSubnetRequestInterface) *ServicesEc2ModelCreateSubnetResult

	// public void createTags(com.amazonaws.services.ec2.model.CreateTagsRequest)
	CreateTags(a ServicesEc2ModelCreateTagsRequestInterface) 

	// public com.amazonaws.services.ec2.model.CreateVolumeResult createVolume(com.amazonaws.services.ec2.model.CreateVolumeRequest)
	CreateVolume(a ServicesEc2ModelCreateVolumeRequestInterface) *ServicesEc2ModelCreateVolumeResult

	// public com.amazonaws.services.ec2.model.CreateVpcResult createVpc(com.amazonaws.services.ec2.model.CreateVpcRequest)
	CreateVpc(a ServicesEc2ModelCreateVpcRequestInterface) *ServicesEc2ModelCreateVpcResult

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult createVpcEndpoint(com.amazonaws.services.ec2.model.CreateVpcEndpointRequest)
	CreateVpcEndpoint(a ServicesEc2ModelCreateVpcEndpointRequestInterface) *ServicesEc2ModelCreateVpcEndpointResult

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult createVpcPeeringConnection(com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest)
	CreateVpcPeeringConnection2(a ServicesEc2ModelCreateVpcPeeringConnectionRequestInterface) *ServicesEc2ModelCreateVpcPeeringConnectionResult

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult createVpcPeeringConnection()
	CreateVpcPeeringConnection() *ServicesEc2ModelCreateVpcPeeringConnectionResult

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult createVpnConnection(com.amazonaws.services.ec2.model.CreateVpnConnectionRequest)
	CreateVpnConnection(a ServicesEc2ModelCreateVpnConnectionRequestInterface) *ServicesEc2ModelCreateVpnConnectionResult

	// public void createVpnConnectionRoute(com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest)
	CreateVpnConnectionRoute(a ServicesEc2ModelCreateVpnConnectionRouteRequestInterface) 

	// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult createVpnGateway(com.amazonaws.services.ec2.model.CreateVpnGatewayRequest)
	CreateVpnGateway(a ServicesEc2ModelCreateVpnGatewayRequestInterface) *ServicesEc2ModelCreateVpnGatewayResult

	// public void deleteCustomerGateway(com.amazonaws.services.ec2.model.DeleteCustomerGatewayRequest)
	DeleteCustomerGateway(a ServicesEc2ModelDeleteCustomerGatewayRequestInterface) 

	// public void deleteDhcpOptions(com.amazonaws.services.ec2.model.DeleteDhcpOptionsRequest)
	DeleteDhcpOptions(a ServicesEc2ModelDeleteDhcpOptionsRequestInterface) 

	// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult deleteFlowLogs(com.amazonaws.services.ec2.model.DeleteFlowLogsRequest)
	DeleteFlowLogs(a ServicesEc2ModelDeleteFlowLogsRequestInterface) *ServicesEc2ModelDeleteFlowLogsResult

	// public void deleteInternetGateway(com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest)
	DeleteInternetGateway(a ServicesEc2ModelDeleteInternetGatewayRequestInterface) 

	// public void deleteKeyPair(com.amazonaws.services.ec2.model.DeleteKeyPairRequest)
	DeleteKeyPair(a ServicesEc2ModelDeleteKeyPairRequestInterface) 

	// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult deleteNatGateway(com.amazonaws.services.ec2.model.DeleteNatGatewayRequest)
	DeleteNatGateway(a ServicesEc2ModelDeleteNatGatewayRequestInterface) *ServicesEc2ModelDeleteNatGatewayResult

	// public void deleteNetworkAcl(com.amazonaws.services.ec2.model.DeleteNetworkAclRequest)
	DeleteNetworkAcl(a ServicesEc2ModelDeleteNetworkAclRequestInterface) 

	// public void deleteNetworkAclEntry(com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest)
	DeleteNetworkAclEntry(a ServicesEc2ModelDeleteNetworkAclEntryRequestInterface) 

	// public void deleteNetworkInterface(com.amazonaws.services.ec2.model.DeleteNetworkInterfaceRequest)
	DeleteNetworkInterface(a ServicesEc2ModelDeleteNetworkInterfaceRequestInterface) 

	// public void deletePlacementGroup(com.amazonaws.services.ec2.model.DeletePlacementGroupRequest)
	DeletePlacementGroup(a ServicesEc2ModelDeletePlacementGroupRequestInterface) 

	// public void deleteRoute(com.amazonaws.services.ec2.model.DeleteRouteRequest)
	DeleteRoute(a ServicesEc2ModelDeleteRouteRequestInterface) 

	// public void deleteRouteTable(com.amazonaws.services.ec2.model.DeleteRouteTableRequest)
	DeleteRouteTable(a ServicesEc2ModelDeleteRouteTableRequestInterface) 

	// public void deleteSecurityGroup(com.amazonaws.services.ec2.model.DeleteSecurityGroupRequest)
	DeleteSecurityGroup(a ServicesEc2ModelDeleteSecurityGroupRequestInterface) 

	// public void deleteSnapshot(com.amazonaws.services.ec2.model.DeleteSnapshotRequest)
	DeleteSnapshot(a ServicesEc2ModelDeleteSnapshotRequestInterface) 

	// public void deleteSpotDatafeedSubscription(com.amazonaws.services.ec2.model.DeleteSpotDatafeedSubscriptionRequest)
	DeleteSpotDatafeedSubscription2(a ServicesEc2ModelDeleteSpotDatafeedSubscriptionRequestInterface) 

	// public void deleteSpotDatafeedSubscription()
	DeleteSpotDatafeedSubscription() 

	// public void deleteSubnet(com.amazonaws.services.ec2.model.DeleteSubnetRequest)
	DeleteSubnet(a ServicesEc2ModelDeleteSubnetRequestInterface) 

	// public void deleteTags(com.amazonaws.services.ec2.model.DeleteTagsRequest)
	DeleteTags(a ServicesEc2ModelDeleteTagsRequestInterface) 

	// public void deleteVolume(com.amazonaws.services.ec2.model.DeleteVolumeRequest)
	DeleteVolume(a ServicesEc2ModelDeleteVolumeRequestInterface) 

	// public void deleteVpc(com.amazonaws.services.ec2.model.DeleteVpcRequest)
	DeleteVpc(a ServicesEc2ModelDeleteVpcRequestInterface) 

	// public com.amazonaws.services.ec2.model.DeleteVpcEndpointsResult deleteVpcEndpoints(com.amazonaws.services.ec2.model.DeleteVpcEndpointsRequest)
	DeleteVpcEndpoints(a ServicesEc2ModelDeleteVpcEndpointsRequestInterface) *ServicesEc2ModelDeleteVpcEndpointsResult

	// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionResult deleteVpcPeeringConnection(com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest)
	DeleteVpcPeeringConnection(a ServicesEc2ModelDeleteVpcPeeringConnectionRequestInterface) *ServicesEc2ModelDeleteVpcPeeringConnectionResult

	// public void deleteVpnConnection(com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest)
	DeleteVpnConnection(a ServicesEc2ModelDeleteVpnConnectionRequestInterface) 

	// public void deleteVpnConnectionRoute(com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest)
	DeleteVpnConnectionRoute(a ServicesEc2ModelDeleteVpnConnectionRouteRequestInterface) 

	// public void deleteVpnGateway(com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest)
	DeleteVpnGateway(a ServicesEc2ModelDeleteVpnGatewayRequestInterface) 

	// public void deregisterImage(com.amazonaws.services.ec2.model.DeregisterImageRequest)
	DeregisterImage(a ServicesEc2ModelDeregisterImageRequestInterface) 

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult describeAccountAttributes(com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest)
	DescribeAccountAttributes2(a ServicesEc2ModelDescribeAccountAttributesRequestInterface) *ServicesEc2ModelDescribeAccountAttributesResult

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult describeAccountAttributes()
	DescribeAccountAttributes() *ServicesEc2ModelDescribeAccountAttributesResult

	// public com.amazonaws.services.ec2.model.DescribeAddressesResult describeAddresses(com.amazonaws.services.ec2.model.DescribeAddressesRequest)
	DescribeAddresses2(a ServicesEc2ModelDescribeAddressesRequestInterface) *ServicesEc2ModelDescribeAddressesResult

	// public com.amazonaws.services.ec2.model.DescribeAddressesResult describeAddresses()
	DescribeAddresses() *ServicesEc2ModelDescribeAddressesResult

	// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult describeAvailabilityZones(com.amazonaws.services.ec2.model.DescribeAvailabilityZonesRequest)
	DescribeAvailabilityZones2(a ServicesEc2ModelDescribeAvailabilityZonesRequestInterface) *ServicesEc2ModelDescribeAvailabilityZonesResult

	// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult describeAvailabilityZones()
	DescribeAvailabilityZones() *ServicesEc2ModelDescribeAvailabilityZonesResult

	// public com.amazonaws.services.ec2.model.DescribeBundleTasksResult describeBundleTasks(com.amazonaws.services.ec2.model.DescribeBundleTasksRequest)
	DescribeBundleTasks2(a ServicesEc2ModelDescribeBundleTasksRequestInterface) *ServicesEc2ModelDescribeBundleTasksResult

	// public com.amazonaws.services.ec2.model.DescribeBundleTasksResult describeBundleTasks()
	DescribeBundleTasks() *ServicesEc2ModelDescribeBundleTasksResult

	// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult describeClassicLinkInstances(com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesRequest)
	DescribeClassicLinkInstances2(a ServicesEc2ModelDescribeClassicLinkInstancesRequestInterface) *ServicesEc2ModelDescribeClassicLinkInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult describeClassicLinkInstances()
	DescribeClassicLinkInstances() *ServicesEc2ModelDescribeClassicLinkInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult describeConversionTasks(com.amazonaws.services.ec2.model.DescribeConversionTasksRequest)
	DescribeConversionTasks2(a ServicesEc2ModelDescribeConversionTasksRequestInterface) *ServicesEc2ModelDescribeConversionTasksResult

	// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult describeConversionTasks()
	DescribeConversionTasks() *ServicesEc2ModelDescribeConversionTasksResult

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult describeCustomerGateways(com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest)
	DescribeCustomerGateways2(a ServicesEc2ModelDescribeCustomerGatewaysRequestInterface) *ServicesEc2ModelDescribeCustomerGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult describeCustomerGateways()
	DescribeCustomerGateways() *ServicesEc2ModelDescribeCustomerGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeDhcpOptionsResult describeDhcpOptions(com.amazonaws.services.ec2.model.DescribeDhcpOptionsRequest)
	DescribeDhcpOptions2(a ServicesEc2ModelDescribeDhcpOptionsRequestInterface) *ServicesEc2ModelDescribeDhcpOptionsResult

	// public com.amazonaws.services.ec2.model.DescribeDhcpOptionsResult describeDhcpOptions()
	DescribeDhcpOptions() *ServicesEc2ModelDescribeDhcpOptionsResult

	// public com.amazonaws.services.ec2.model.DescribeExportTasksResult describeExportTasks(com.amazonaws.services.ec2.model.DescribeExportTasksRequest)
	DescribeExportTasks2(a ServicesEc2ModelDescribeExportTasksRequestInterface) *ServicesEc2ModelDescribeExportTasksResult

	// public com.amazonaws.services.ec2.model.DescribeExportTasksResult describeExportTasks()
	DescribeExportTasks() *ServicesEc2ModelDescribeExportTasksResult

	// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult describeFlowLogs(com.amazonaws.services.ec2.model.DescribeFlowLogsRequest)
	DescribeFlowLogs2(a ServicesEc2ModelDescribeFlowLogsRequestInterface) *ServicesEc2ModelDescribeFlowLogsResult

	// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult describeFlowLogs()
	DescribeFlowLogs() *ServicesEc2ModelDescribeFlowLogsResult

	// public com.amazonaws.services.ec2.model.DescribeHostsResult describeHosts(com.amazonaws.services.ec2.model.DescribeHostsRequest)
	DescribeHosts2(a ServicesEc2ModelDescribeHostsRequestInterface) *ServicesEc2ModelDescribeHostsResult

	// public com.amazonaws.services.ec2.model.DescribeHostsResult describeHosts()
	DescribeHosts() *ServicesEc2ModelDescribeHostsResult

	// public com.amazonaws.services.ec2.model.DescribeIdFormatResult describeIdFormat(com.amazonaws.services.ec2.model.DescribeIdFormatRequest)
	DescribeIdFormat2(a ServicesEc2ModelDescribeIdFormatRequestInterface) *ServicesEc2ModelDescribeIdFormatResult

	// public com.amazonaws.services.ec2.model.DescribeIdFormatResult describeIdFormat()
	DescribeIdFormat() *ServicesEc2ModelDescribeIdFormatResult

	// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult describeImageAttribute(com.amazonaws.services.ec2.model.DescribeImageAttributeRequest)
	DescribeImageAttribute(a ServicesEc2ModelDescribeImageAttributeRequestInterface) *ServicesEc2ModelDescribeImageAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeImagesResult describeImages(com.amazonaws.services.ec2.model.DescribeImagesRequest)
	DescribeImages2(a ServicesEc2ModelDescribeImagesRequestInterface) *ServicesEc2ModelDescribeImagesResult

	// public com.amazonaws.services.ec2.model.DescribeImagesResult describeImages()
	DescribeImages() *ServicesEc2ModelDescribeImagesResult

	// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult describeImportImageTasks(com.amazonaws.services.ec2.model.DescribeImportImageTasksRequest)
	DescribeImportImageTasks2(a ServicesEc2ModelDescribeImportImageTasksRequestInterface) *ServicesEc2ModelDescribeImportImageTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult describeImportImageTasks()
	DescribeImportImageTasks() *ServicesEc2ModelDescribeImportImageTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult describeImportSnapshotTasks(com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksRequest)
	DescribeImportSnapshotTasks2(a ServicesEc2ModelDescribeImportSnapshotTasksRequestInterface) *ServicesEc2ModelDescribeImportSnapshotTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult describeImportSnapshotTasks()
	DescribeImportSnapshotTasks() *ServicesEc2ModelDescribeImportSnapshotTasksResult

	// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult describeInstanceAttribute(com.amazonaws.services.ec2.model.DescribeInstanceAttributeRequest)
	DescribeInstanceAttribute(a ServicesEc2ModelDescribeInstanceAttributeRequestInterface) *ServicesEc2ModelDescribeInstanceAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult describeInstanceStatus(com.amazonaws.services.ec2.model.DescribeInstanceStatusRequest)
	DescribeInstanceStatus2(a ServicesEc2ModelDescribeInstanceStatusRequestInterface) *ServicesEc2ModelDescribeInstanceStatusResult

	// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult describeInstanceStatus()
	DescribeInstanceStatus() *ServicesEc2ModelDescribeInstanceStatusResult

	// public com.amazonaws.services.ec2.model.DescribeInstancesResult describeInstances(com.amazonaws.services.ec2.model.DescribeInstancesRequest)
	DescribeInstances2(a ServicesEc2ModelDescribeInstancesRequestInterface) *ServicesEc2ModelDescribeInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeInstancesResult describeInstances()
	DescribeInstances() *ServicesEc2ModelDescribeInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult describeInternetGateways(com.amazonaws.services.ec2.model.DescribeInternetGatewaysRequest)
	DescribeInternetGateways2(a ServicesEc2ModelDescribeInternetGatewaysRequestInterface) *ServicesEc2ModelDescribeInternetGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult describeInternetGateways()
	DescribeInternetGateways() *ServicesEc2ModelDescribeInternetGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult describeKeyPairs(com.amazonaws.services.ec2.model.DescribeKeyPairsRequest)
	DescribeKeyPairs2(a ServicesEc2ModelDescribeKeyPairsRequestInterface) *ServicesEc2ModelDescribeKeyPairsResult

	// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult describeKeyPairs()
	DescribeKeyPairs() *ServicesEc2ModelDescribeKeyPairsResult

	// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult describeMovingAddresses(com.amazonaws.services.ec2.model.DescribeMovingAddressesRequest)
	DescribeMovingAddresses2(a ServicesEc2ModelDescribeMovingAddressesRequestInterface) *ServicesEc2ModelDescribeMovingAddressesResult

	// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult describeMovingAddresses()
	DescribeMovingAddresses() *ServicesEc2ModelDescribeMovingAddressesResult

	// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult describeNatGateways(com.amazonaws.services.ec2.model.DescribeNatGatewaysRequest)
	DescribeNatGateways(a ServicesEc2ModelDescribeNatGatewaysRequestInterface) *ServicesEc2ModelDescribeNatGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult describeNetworkAcls(com.amazonaws.services.ec2.model.DescribeNetworkAclsRequest)
	DescribeNetworkAcls2(a ServicesEc2ModelDescribeNetworkAclsRequestInterface) *ServicesEc2ModelDescribeNetworkAclsResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult describeNetworkAcls()
	DescribeNetworkAcls() *ServicesEc2ModelDescribeNetworkAclsResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeResult describeNetworkInterfaceAttribute(com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest)
	DescribeNetworkInterfaceAttribute(a ServicesEc2ModelDescribeNetworkInterfaceAttributeRequestInterface) *ServicesEc2ModelDescribeNetworkInterfaceAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult describeNetworkInterfaces(com.amazonaws.services.ec2.model.DescribeNetworkInterfacesRequest)
	DescribeNetworkInterfaces2(a ServicesEc2ModelDescribeNetworkInterfacesRequestInterface) *ServicesEc2ModelDescribeNetworkInterfacesResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult describeNetworkInterfaces()
	DescribeNetworkInterfaces() *ServicesEc2ModelDescribeNetworkInterfacesResult

	// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult describePlacementGroups(com.amazonaws.services.ec2.model.DescribePlacementGroupsRequest)
	DescribePlacementGroups2(a ServicesEc2ModelDescribePlacementGroupsRequestInterface) *ServicesEc2ModelDescribePlacementGroupsResult

	// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult describePlacementGroups()
	DescribePlacementGroups() *ServicesEc2ModelDescribePlacementGroupsResult

	// public com.amazonaws.services.ec2.model.DescribePrefixListsResult describePrefixLists(com.amazonaws.services.ec2.model.DescribePrefixListsRequest)
	DescribePrefixLists2(a ServicesEc2ModelDescribePrefixListsRequestInterface) *ServicesEc2ModelDescribePrefixListsResult

	// public com.amazonaws.services.ec2.model.DescribePrefixListsResult describePrefixLists()
	DescribePrefixLists() *ServicesEc2ModelDescribePrefixListsResult

	// public com.amazonaws.services.ec2.model.DescribeRegionsResult describeRegions(com.amazonaws.services.ec2.model.DescribeRegionsRequest)
	DescribeRegions2(a ServicesEc2ModelDescribeRegionsRequestInterface) *ServicesEc2ModelDescribeRegionsResult

	// public com.amazonaws.services.ec2.model.DescribeRegionsResult describeRegions()
	DescribeRegions() *ServicesEc2ModelDescribeRegionsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult describeReservedInstances(com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest)
	DescribeReservedInstances2(a ServicesEc2ModelDescribeReservedInstancesRequestInterface) *ServicesEc2ModelDescribeReservedInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult describeReservedInstances()
	DescribeReservedInstances() *ServicesEc2ModelDescribeReservedInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesListingsResult describeReservedInstancesListings(com.amazonaws.services.ec2.model.DescribeReservedInstancesListingsRequest)
	DescribeReservedInstancesListings2(a ServicesEc2ModelDescribeReservedInstancesListingsRequestInterface) *ServicesEc2ModelDescribeReservedInstancesListingsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesListingsResult describeReservedInstancesListings()
	DescribeReservedInstancesListings() *ServicesEc2ModelDescribeReservedInstancesListingsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult describeReservedInstancesModifications(com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest)
	DescribeReservedInstancesModifications2(a ServicesEc2ModelDescribeReservedInstancesModificationsRequestInterface) *ServicesEc2ModelDescribeReservedInstancesModificationsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult describeReservedInstancesModifications()
	DescribeReservedInstancesModifications() *ServicesEc2ModelDescribeReservedInstancesModificationsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult describeReservedInstancesOfferings(com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest)
	DescribeReservedInstancesOfferings2(a ServicesEc2ModelDescribeReservedInstancesOfferingsRequestInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult describeReservedInstancesOfferings()
	DescribeReservedInstancesOfferings() *ServicesEc2ModelDescribeReservedInstancesOfferingsResult

	// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult describeRouteTables(com.amazonaws.services.ec2.model.DescribeRouteTablesRequest)
	DescribeRouteTables2(a ServicesEc2ModelDescribeRouteTablesRequestInterface) *ServicesEc2ModelDescribeRouteTablesResult

	// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult describeRouteTables()
	DescribeRouteTables() *ServicesEc2ModelDescribeRouteTablesResult

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult describeScheduledInstanceAvailability(com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest)
	DescribeScheduledInstanceAvailability(a ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequestInterface) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult describeScheduledInstances(com.amazonaws.services.ec2.model.DescribeScheduledInstancesRequest)
	DescribeScheduledInstances(a ServicesEc2ModelDescribeScheduledInstancesRequestInterface) *ServicesEc2ModelDescribeScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult describeSecurityGroups(com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest)
	DescribeSecurityGroups2(a ServicesEc2ModelDescribeSecurityGroupsRequestInterface) *ServicesEc2ModelDescribeSecurityGroupsResult

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult describeSecurityGroups()
	DescribeSecurityGroups() *ServicesEc2ModelDescribeSecurityGroupsResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult describeSnapshotAttribute(com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest)
	DescribeSnapshotAttribute(a ServicesEc2ModelDescribeSnapshotAttributeRequestInterface) *ServicesEc2ModelDescribeSnapshotAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult describeSnapshots(com.amazonaws.services.ec2.model.DescribeSnapshotsRequest)
	DescribeSnapshots2(a ServicesEc2ModelDescribeSnapshotsRequestInterface) *ServicesEc2ModelDescribeSnapshotsResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult describeSnapshots()
	DescribeSnapshots() *ServicesEc2ModelDescribeSnapshotsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult describeSpotDatafeedSubscription(com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest)
	DescribeSpotDatafeedSubscription2(a ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequestInterface) *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult

	// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult describeSpotDatafeedSubscription()
	DescribeSpotDatafeedSubscription() *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult describeSpotFleetInstances(com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesRequest)
	DescribeSpotFleetInstances(a ServicesEc2ModelDescribeSpotFleetInstancesRequestInterface) *ServicesEc2ModelDescribeSpotFleetInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult describeSpotFleetRequestHistory(com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryRequest)
	DescribeSpotFleetRequestHistory(a ServicesEc2ModelDescribeSpotFleetRequestHistoryRequestInterface) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult describeSpotFleetRequests(com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest)
	DescribeSpotFleetRequests2(a ServicesEc2ModelDescribeSpotFleetRequestsRequestInterface) *ServicesEc2ModelDescribeSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult describeSpotFleetRequests()
	DescribeSpotFleetRequests() *ServicesEc2ModelDescribeSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult describeSpotInstanceRequests(com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsRequest)
	DescribeSpotInstanceRequests2(a ServicesEc2ModelDescribeSpotInstanceRequestsRequestInterface) *ServicesEc2ModelDescribeSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult describeSpotInstanceRequests()
	DescribeSpotInstanceRequests() *ServicesEc2ModelDescribeSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult describeSpotPriceHistory(com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest)
	DescribeSpotPriceHistory2(a ServicesEc2ModelDescribeSpotPriceHistoryRequestInterface) *ServicesEc2ModelDescribeSpotPriceHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult describeSpotPriceHistory()
	DescribeSpotPriceHistory() *ServicesEc2ModelDescribeSpotPriceHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSubnetsResult describeSubnets(com.amazonaws.services.ec2.model.DescribeSubnetsRequest)
	DescribeSubnets2(a ServicesEc2ModelDescribeSubnetsRequestInterface) *ServicesEc2ModelDescribeSubnetsResult

	// public com.amazonaws.services.ec2.model.DescribeSubnetsResult describeSubnets()
	DescribeSubnets() *ServicesEc2ModelDescribeSubnetsResult

	// public com.amazonaws.services.ec2.model.DescribeTagsResult describeTags(com.amazonaws.services.ec2.model.DescribeTagsRequest)
	DescribeTags2(a ServicesEc2ModelDescribeTagsRequestInterface) *ServicesEc2ModelDescribeTagsResult

	// public com.amazonaws.services.ec2.model.DescribeTagsResult describeTags()
	DescribeTags() *ServicesEc2ModelDescribeTagsResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult describeVolumeAttribute(com.amazonaws.services.ec2.model.DescribeVolumeAttributeRequest)
	DescribeVolumeAttribute(a ServicesEc2ModelDescribeVolumeAttributeRequestInterface) *ServicesEc2ModelDescribeVolumeAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult describeVolumeStatus(com.amazonaws.services.ec2.model.DescribeVolumeStatusRequest)
	DescribeVolumeStatus2(a ServicesEc2ModelDescribeVolumeStatusRequestInterface) *ServicesEc2ModelDescribeVolumeStatusResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult describeVolumeStatus()
	DescribeVolumeStatus() *ServicesEc2ModelDescribeVolumeStatusResult

	// public com.amazonaws.services.ec2.model.DescribeVolumesResult describeVolumes(com.amazonaws.services.ec2.model.DescribeVolumesRequest)
	DescribeVolumes2(a ServicesEc2ModelDescribeVolumesRequestInterface) *ServicesEc2ModelDescribeVolumesResult

	// public com.amazonaws.services.ec2.model.DescribeVolumesResult describeVolumes()
	DescribeVolumes() *ServicesEc2ModelDescribeVolumesResult

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult describeVpcAttribute(com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest)
	DescribeVpcAttribute(a ServicesEc2ModelDescribeVpcAttributeRequestInterface) *ServicesEc2ModelDescribeVpcAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult describeVpcClassicLink(com.amazonaws.services.ec2.model.DescribeVpcClassicLinkRequest)
	DescribeVpcClassicLink2(a ServicesEc2ModelDescribeVpcClassicLinkRequestInterface) *ServicesEc2ModelDescribeVpcClassicLinkResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult describeVpcClassicLink()
	DescribeVpcClassicLink() *ServicesEc2ModelDescribeVpcClassicLinkResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult describeVpcClassicLinkDnsSupport(com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportRequest)
	DescribeVpcClassicLinkDnsSupport(a ServicesEc2ModelDescribeVpcClassicLinkDnsSupportRequestInterface) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult describeVpcEndpointServices(com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesRequest)
	DescribeVpcEndpointServices2(a ServicesEc2ModelDescribeVpcEndpointServicesRequestInterface) *ServicesEc2ModelDescribeVpcEndpointServicesResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult describeVpcEndpointServices()
	DescribeVpcEndpointServices() *ServicesEc2ModelDescribeVpcEndpointServicesResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult describeVpcEndpoints(com.amazonaws.services.ec2.model.DescribeVpcEndpointsRequest)
	DescribeVpcEndpoints2(a ServicesEc2ModelDescribeVpcEndpointsRequestInterface) *ServicesEc2ModelDescribeVpcEndpointsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult describeVpcEndpoints()
	DescribeVpcEndpoints() *ServicesEc2ModelDescribeVpcEndpointsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult describeVpcPeeringConnections(com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest)
	DescribeVpcPeeringConnections2(a ServicesEc2ModelDescribeVpcPeeringConnectionsRequestInterface) *ServicesEc2ModelDescribeVpcPeeringConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult describeVpcPeeringConnections()
	DescribeVpcPeeringConnections() *ServicesEc2ModelDescribeVpcPeeringConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcsResult describeVpcs(com.amazonaws.services.ec2.model.DescribeVpcsRequest)
	DescribeVpcs2(a ServicesEc2ModelDescribeVpcsRequestInterface) *ServicesEc2ModelDescribeVpcsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcsResult describeVpcs()
	DescribeVpcs() *ServicesEc2ModelDescribeVpcsResult

	// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult describeVpnConnections(com.amazonaws.services.ec2.model.DescribeVpnConnectionsRequest)
	DescribeVpnConnections2(a ServicesEc2ModelDescribeVpnConnectionsRequestInterface) *ServicesEc2ModelDescribeVpnConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult describeVpnConnections()
	DescribeVpnConnections() *ServicesEc2ModelDescribeVpnConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult describeVpnGateways(com.amazonaws.services.ec2.model.DescribeVpnGatewaysRequest)
	DescribeVpnGateways2(a ServicesEc2ModelDescribeVpnGatewaysRequestInterface) *ServicesEc2ModelDescribeVpnGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult describeVpnGateways()
	DescribeVpnGateways() *ServicesEc2ModelDescribeVpnGatewaysResult

	// public com.amazonaws.services.ec2.model.DetachClassicLinkVpcResult detachClassicLinkVpc(com.amazonaws.services.ec2.model.DetachClassicLinkVpcRequest)
	DetachClassicLinkVpc(a ServicesEc2ModelDetachClassicLinkVpcRequestInterface) *ServicesEc2ModelDetachClassicLinkVpcResult

	// public void detachInternetGateway(com.amazonaws.services.ec2.model.DetachInternetGatewayRequest)
	DetachInternetGateway(a ServicesEc2ModelDetachInternetGatewayRequestInterface) 

	// public void detachNetworkInterface(com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest)
	DetachNetworkInterface(a ServicesEc2ModelDetachNetworkInterfaceRequestInterface) 

	// public com.amazonaws.services.ec2.model.DetachVolumeResult detachVolume(com.amazonaws.services.ec2.model.DetachVolumeRequest)
	DetachVolume(a ServicesEc2ModelDetachVolumeRequestInterface) *ServicesEc2ModelDetachVolumeResult

	// public void detachVpnGateway(com.amazonaws.services.ec2.model.DetachVpnGatewayRequest)
	DetachVpnGateway(a ServicesEc2ModelDetachVpnGatewayRequestInterface) 

	// public void disableVgwRoutePropagation(com.amazonaws.services.ec2.model.DisableVgwRoutePropagationRequest)
	DisableVgwRoutePropagation(a ServicesEc2ModelDisableVgwRoutePropagationRequestInterface) 

	// public com.amazonaws.services.ec2.model.DisableVpcClassicLinkResult disableVpcClassicLink(com.amazonaws.services.ec2.model.DisableVpcClassicLinkRequest)
	DisableVpcClassicLink(a ServicesEc2ModelDisableVpcClassicLinkRequestInterface) *ServicesEc2ModelDisableVpcClassicLinkResult

	// public com.amazonaws.services.ec2.model.DisableVpcClassicLinkDnsSupportResult disableVpcClassicLinkDnsSupport(com.amazonaws.services.ec2.model.DisableVpcClassicLinkDnsSupportRequest)
	DisableVpcClassicLinkDnsSupport(a ServicesEc2ModelDisableVpcClassicLinkDnsSupportRequestInterface) *ServicesEc2ModelDisableVpcClassicLinkDnsSupportResult

	// public void disassociateAddress(com.amazonaws.services.ec2.model.DisassociateAddressRequest)
	DisassociateAddress(a ServicesEc2ModelDisassociateAddressRequestInterface) 

	// public void disassociateRouteTable(com.amazonaws.services.ec2.model.DisassociateRouteTableRequest)
	DisassociateRouteTable(a ServicesEc2ModelDisassociateRouteTableRequestInterface) 

	// public void enableVgwRoutePropagation(com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest)
	EnableVgwRoutePropagation(a ServicesEc2ModelEnableVgwRoutePropagationRequestInterface) 

	// public void enableVolumeIO(com.amazonaws.services.ec2.model.EnableVolumeIORequest)
	EnableVolumeIO(a ServicesEc2ModelEnableVolumeIORequestInterface) 

	// public com.amazonaws.services.ec2.model.EnableVpcClassicLinkResult enableVpcClassicLink(com.amazonaws.services.ec2.model.EnableVpcClassicLinkRequest)
	EnableVpcClassicLink(a ServicesEc2ModelEnableVpcClassicLinkRequestInterface) *ServicesEc2ModelEnableVpcClassicLinkResult

	// public com.amazonaws.services.ec2.model.EnableVpcClassicLinkDnsSupportResult enableVpcClassicLinkDnsSupport(com.amazonaws.services.ec2.model.EnableVpcClassicLinkDnsSupportRequest)
	EnableVpcClassicLinkDnsSupport(a ServicesEc2ModelEnableVpcClassicLinkDnsSupportRequestInterface) *ServicesEc2ModelEnableVpcClassicLinkDnsSupportResult

	// public com.amazonaws.services.ec2.model.GetConsoleOutputResult getConsoleOutput(com.amazonaws.services.ec2.model.GetConsoleOutputRequest)
	GetConsoleOutput(a ServicesEc2ModelGetConsoleOutputRequestInterface) *ServicesEc2ModelGetConsoleOutputResult

	// public com.amazonaws.services.ec2.model.GetPasswordDataResult getPasswordData(com.amazonaws.services.ec2.model.GetPasswordDataRequest)
	GetPasswordData(a ServicesEc2ModelGetPasswordDataRequestInterface) *ServicesEc2ModelGetPasswordDataResult

	// public com.amazonaws.services.ec2.model.ImportImageResult importImage(com.amazonaws.services.ec2.model.ImportImageRequest)
	ImportImage2(a ServicesEc2ModelImportImageRequestInterface) *ServicesEc2ModelImportImageResult

	// public com.amazonaws.services.ec2.model.ImportImageResult importImage()
	ImportImage() *ServicesEc2ModelImportImageResult

	// public com.amazonaws.services.ec2.model.ImportInstanceResult importInstance(com.amazonaws.services.ec2.model.ImportInstanceRequest)
	ImportInstance(a ServicesEc2ModelImportInstanceRequestInterface) *ServicesEc2ModelImportInstanceResult

	// public com.amazonaws.services.ec2.model.ImportKeyPairResult importKeyPair(com.amazonaws.services.ec2.model.ImportKeyPairRequest)
	ImportKeyPair(a ServicesEc2ModelImportKeyPairRequestInterface) *ServicesEc2ModelImportKeyPairResult

	// public com.amazonaws.services.ec2.model.ImportSnapshotResult importSnapshot(com.amazonaws.services.ec2.model.ImportSnapshotRequest)
	ImportSnapshot2(a ServicesEc2ModelImportSnapshotRequestInterface) *ServicesEc2ModelImportSnapshotResult

	// public com.amazonaws.services.ec2.model.ImportSnapshotResult importSnapshot()
	ImportSnapshot() *ServicesEc2ModelImportSnapshotResult

	// public com.amazonaws.services.ec2.model.ImportVolumeResult importVolume(com.amazonaws.services.ec2.model.ImportVolumeRequest)
	ImportVolume(a ServicesEc2ModelImportVolumeRequestInterface) *ServicesEc2ModelImportVolumeResult

	// public com.amazonaws.services.ec2.model.ModifyHostsResult modifyHosts(com.amazonaws.services.ec2.model.ModifyHostsRequest)
	ModifyHosts(a ServicesEc2ModelModifyHostsRequestInterface) *ServicesEc2ModelModifyHostsResult

	// public void modifyIdFormat(com.amazonaws.services.ec2.model.ModifyIdFormatRequest)
	ModifyIdFormat(a ServicesEc2ModelModifyIdFormatRequestInterface) 

	// public void modifyImageAttribute(com.amazonaws.services.ec2.model.ModifyImageAttributeRequest)
	ModifyImageAttribute(a ServicesEc2ModelModifyImageAttributeRequestInterface) 

	// public void modifyInstanceAttribute(com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest)
	ModifyInstanceAttribute(a ServicesEc2ModelModifyInstanceAttributeRequestInterface) 

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementResult modifyInstancePlacement(com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest)
	ModifyInstancePlacement(a ServicesEc2ModelModifyInstancePlacementRequestInterface) *ServicesEc2ModelModifyInstancePlacementResult

	// public void modifyNetworkInterfaceAttribute(com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest)
	ModifyNetworkInterfaceAttribute(a ServicesEc2ModelModifyNetworkInterfaceAttributeRequestInterface) 

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult modifyReservedInstances(com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest)
	ModifyReservedInstances(a ServicesEc2ModelModifyReservedInstancesRequestInterface) *ServicesEc2ModelModifyReservedInstancesResult

	// public void modifySnapshotAttribute(com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest)
	ModifySnapshotAttribute(a ServicesEc2ModelModifySnapshotAttributeRequestInterface) 

	// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestResult modifySpotFleetRequest(com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest)
	ModifySpotFleetRequest(a ServicesEc2ModelModifySpotFleetRequestRequestInterface) *ServicesEc2ModelModifySpotFleetRequestResult

	// public void modifySubnetAttribute(com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest)
	ModifySubnetAttribute(a ServicesEc2ModelModifySubnetAttributeRequestInterface) 

	// public void modifyVolumeAttribute(com.amazonaws.services.ec2.model.ModifyVolumeAttributeRequest)
	ModifyVolumeAttribute(a ServicesEc2ModelModifyVolumeAttributeRequestInterface) 

	// public void modifyVpcAttribute(com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest)
	ModifyVpcAttribute(a ServicesEc2ModelModifyVpcAttributeRequestInterface) 

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult modifyVpcEndpoint(com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest)
	ModifyVpcEndpoint(a ServicesEc2ModelModifyVpcEndpointRequestInterface) *ServicesEc2ModelModifyVpcEndpointResult

	// public com.amazonaws.services.ec2.model.MonitorInstancesResult monitorInstances(com.amazonaws.services.ec2.model.MonitorInstancesRequest)
	MonitorInstances(a ServicesEc2ModelMonitorInstancesRequestInterface) *ServicesEc2ModelMonitorInstancesResult

	// public com.amazonaws.services.ec2.model.MoveAddressToVpcResult moveAddressToVpc(com.amazonaws.services.ec2.model.MoveAddressToVpcRequest)
	MoveAddressToVpc(a ServicesEc2ModelMoveAddressToVpcRequestInterface) *ServicesEc2ModelMoveAddressToVpcResult

	// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult purchaseReservedInstancesOffering(com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest)
	PurchaseReservedInstancesOffering(a ServicesEc2ModelPurchaseReservedInstancesOfferingRequestInterface) *ServicesEc2ModelPurchaseReservedInstancesOfferingResult

	// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult purchaseScheduledInstances(com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest)
	PurchaseScheduledInstances(a ServicesEc2ModelPurchaseScheduledInstancesRequestInterface) *ServicesEc2ModelPurchaseScheduledInstancesResult

	// public void rebootInstances(com.amazonaws.services.ec2.model.RebootInstancesRequest)
	RebootInstances(a ServicesEc2ModelRebootInstancesRequestInterface) 

	// public com.amazonaws.services.ec2.model.RegisterImageResult registerImage(com.amazonaws.services.ec2.model.RegisterImageRequest)
	RegisterImage(a ServicesEc2ModelRegisterImageRequestInterface) *ServicesEc2ModelRegisterImageResult

	// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionResult rejectVpcPeeringConnection(com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest)
	RejectVpcPeeringConnection(a ServicesEc2ModelRejectVpcPeeringConnectionRequestInterface) *ServicesEc2ModelRejectVpcPeeringConnectionResult

	// public void releaseAddress(com.amazonaws.services.ec2.model.ReleaseAddressRequest)
	ReleaseAddress(a ServicesEc2ModelReleaseAddressRequestInterface) 

	// public com.amazonaws.services.ec2.model.ReleaseHostsResult releaseHosts(com.amazonaws.services.ec2.model.ReleaseHostsRequest)
	ReleaseHosts(a ServicesEc2ModelReleaseHostsRequestInterface) *ServicesEc2ModelReleaseHostsResult

	// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult replaceNetworkAclAssociation(com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationRequest)
	ReplaceNetworkAclAssociation(a ServicesEc2ModelReplaceNetworkAclAssociationRequestInterface) *ServicesEc2ModelReplaceNetworkAclAssociationResult

	// public void replaceNetworkAclEntry(com.amazonaws.services.ec2.model.ReplaceNetworkAclEntryRequest)
	ReplaceNetworkAclEntry(a ServicesEc2ModelReplaceNetworkAclEntryRequestInterface) 

	// public void replaceRoute(com.amazonaws.services.ec2.model.ReplaceRouteRequest)
	ReplaceRoute(a ServicesEc2ModelReplaceRouteRequestInterface) 

	// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult replaceRouteTableAssociation(com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationRequest)
	ReplaceRouteTableAssociation(a ServicesEc2ModelReplaceRouteTableAssociationRequestInterface) *ServicesEc2ModelReplaceRouteTableAssociationResult

	// public void reportInstanceStatus(com.amazonaws.services.ec2.model.ReportInstanceStatusRequest)
	ReportInstanceStatus(a ServicesEc2ModelReportInstanceStatusRequestInterface) 

	// public com.amazonaws.services.ec2.model.RequestSpotFleetResult requestSpotFleet(com.amazonaws.services.ec2.model.RequestSpotFleetRequest)
	RequestSpotFleet(a ServicesEc2ModelRequestSpotFleetRequestInterface) *ServicesEc2ModelRequestSpotFleetResult

	// public com.amazonaws.services.ec2.model.RequestSpotInstancesResult requestSpotInstances(com.amazonaws.services.ec2.model.RequestSpotInstancesRequest)
	RequestSpotInstances(a ServicesEc2ModelRequestSpotInstancesRequestInterface) *ServicesEc2ModelRequestSpotInstancesResult

	// public void resetImageAttribute(com.amazonaws.services.ec2.model.ResetImageAttributeRequest)
	ResetImageAttribute(a ServicesEc2ModelResetImageAttributeRequestInterface) 

	// public void resetInstanceAttribute(com.amazonaws.services.ec2.model.ResetInstanceAttributeRequest)
	ResetInstanceAttribute(a ServicesEc2ModelResetInstanceAttributeRequestInterface) 

	// public void resetNetworkInterfaceAttribute(com.amazonaws.services.ec2.model.ResetNetworkInterfaceAttributeRequest)
	ResetNetworkInterfaceAttribute(a ServicesEc2ModelResetNetworkInterfaceAttributeRequestInterface) 

	// public void resetSnapshotAttribute(com.amazonaws.services.ec2.model.ResetSnapshotAttributeRequest)
	ResetSnapshotAttribute(a ServicesEc2ModelResetSnapshotAttributeRequestInterface) 

	// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult restoreAddressToClassic(com.amazonaws.services.ec2.model.RestoreAddressToClassicRequest)
	RestoreAddressToClassic(a ServicesEc2ModelRestoreAddressToClassicRequestInterface) *ServicesEc2ModelRestoreAddressToClassicResult

	// public void revokeSecurityGroupEgress(com.amazonaws.services.ec2.model.RevokeSecurityGroupEgressRequest)
	RevokeSecurityGroupEgress(a ServicesEc2ModelRevokeSecurityGroupEgressRequestInterface) 

	// public void revokeSecurityGroupIngress(com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest)
	RevokeSecurityGroupIngress2(a ServicesEc2ModelRevokeSecurityGroupIngressRequestInterface) 

	// public void revokeSecurityGroupIngress()
	RevokeSecurityGroupIngress() 

	// public com.amazonaws.services.ec2.model.RunInstancesResult runInstances(com.amazonaws.services.ec2.model.RunInstancesRequest)
	RunInstances(a ServicesEc2ModelRunInstancesRequestInterface) *ServicesEc2ModelRunInstancesResult

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult runScheduledInstances(com.amazonaws.services.ec2.model.RunScheduledInstancesRequest)
	RunScheduledInstances(a ServicesEc2ModelRunScheduledInstancesRequestInterface) *ServicesEc2ModelRunScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.StartInstancesResult startInstances(com.amazonaws.services.ec2.model.StartInstancesRequest)
	StartInstances(a ServicesEc2ModelStartInstancesRequestInterface) *ServicesEc2ModelStartInstancesResult

	// public com.amazonaws.services.ec2.model.StopInstancesResult stopInstances(com.amazonaws.services.ec2.model.StopInstancesRequest)
	StopInstances(a ServicesEc2ModelStopInstancesRequestInterface) *ServicesEc2ModelStopInstancesResult

	// public com.amazonaws.services.ec2.model.TerminateInstancesResult terminateInstances(com.amazonaws.services.ec2.model.TerminateInstancesRequest)
	TerminateInstances(a ServicesEc2ModelTerminateInstancesRequestInterface) *ServicesEc2ModelTerminateInstancesResult

	// public void unassignPrivateIpAddresses(com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest)
	UnassignPrivateIpAddresses(a ServicesEc2ModelUnassignPrivateIpAddressesRequestInterface) 

	// public com.amazonaws.services.ec2.model.UnmonitorInstancesResult unmonitorInstances(com.amazonaws.services.ec2.model.UnmonitorInstancesRequest)
	UnmonitorInstances(a ServicesEc2ModelUnmonitorInstancesRequestInterface) *ServicesEc2ModelUnmonitorInstancesResult

	// public com.amazonaws.ResponseMetadata getCachedResponseMetadata(com.amazonaws.AmazonWebServiceRequest)
	GetCachedResponseMetadata(a AmazonWebServiceRequestInterface) *ResponseMetadata
}

type ServicesEc2AmazonEC2Client struct {
	AmazonWebServiceClient
}

// public com.amazonaws.services.ec2.AmazonEC2Client()
func NewServicesEc2AmazonEC2Client() (*ServicesEc2AmazonEC2Client) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.AmazonEC2Client(com.amazonaws.ClientConfiguration)
func NewServicesEc2AmazonEC2Client2(a ClientConfigurationInterface) (*ServicesEc2AmazonEC2Client) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client", conv_a.Value().Cast("com/amazonaws/ClientConfiguration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.AmazonEC2Client(com.amazonaws.auth.AWSCredentials)
func NewServicesEc2AmazonEC2Client3(a AuthAWSCredentialsInterface) (*ServicesEc2AmazonEC2Client) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client", conv_a.Value().Cast("com/amazonaws/auth/AWSCredentials"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.AmazonEC2Client(com.amazonaws.auth.AWSCredentials, com.amazonaws.ClientConfiguration)
func NewServicesEc2AmazonEC2Client5(a AuthAWSCredentialsInterface, b ClientConfigurationInterface) (*ServicesEc2AmazonEC2Client) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client", conv_a.Value().Cast("com/amazonaws/auth/AWSCredentials"), conv_b.Value().Cast("com/amazonaws/ClientConfiguration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.AmazonEC2Client(com.amazonaws.auth.AWSCredentialsProvider)
func NewServicesEc2AmazonEC2Client4(a AuthAWSCredentialsProviderInterface) (*ServicesEc2AmazonEC2Client) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client", conv_a.Value().Cast("com/amazonaws/auth/AWSCredentialsProvider"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.AmazonEC2Client(com.amazonaws.auth.AWSCredentialsProvider, com.amazonaws.ClientConfiguration)
func NewServicesEc2AmazonEC2Client6(a AuthAWSCredentialsProviderInterface, b ClientConfigurationInterface) (*ServicesEc2AmazonEC2Client) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client", conv_a.Value().Cast("com/amazonaws/auth/AWSCredentialsProvider"), conv_b.Value().Cast("com/amazonaws/ClientConfiguration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.AmazonEC2Client(com.amazonaws.auth.AWSCredentialsProvider, com.amazonaws.ClientConfiguration, com.amazonaws.metrics.RequestMetricCollector)
func NewServicesEc2AmazonEC2Client7(a AuthAWSCredentialsProviderInterface, b ClientConfigurationInterface, c MetricsRequestMetricCollectorInterface) (*ServicesEc2AmazonEC2Client) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/AmazonEC2Client", conv_a.Value().Cast("com/amazonaws/auth/AWSCredentialsProvider"), conv_b.Value().Cast("com/amazonaws/ClientConfiguration"), conv_c.Value().Cast("com/amazonaws/metrics/RequestMetricCollector"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2AmazonEC2Client{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult acceptVpcPeeringConnection(com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AcceptVpcPeeringConnection2(a ServicesEc2ModelAcceptVpcPeeringConnectionRequestInterface) *ServicesEc2ModelAcceptVpcPeeringConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "acceptVpcPeeringConnection", "com/amazonaws/services/ec2/model/AcceptVpcPeeringConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AcceptVpcPeeringConnectionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAcceptVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult acceptVpcPeeringConnection()
func (jbobject *ServicesEc2AmazonEC2Client) AcceptVpcPeeringConnection() *ServicesEc2ModelAcceptVpcPeeringConnectionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "acceptVpcPeeringConnection", "com/amazonaws/services/ec2/model/AcceptVpcPeeringConnectionResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAcceptVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AllocateAddressResult allocateAddress(com.amazonaws.services.ec2.model.AllocateAddressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AllocateAddress2(a ServicesEc2ModelAllocateAddressRequestInterface) *ServicesEc2ModelAllocateAddressResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "allocateAddress", "com/amazonaws/services/ec2/model/AllocateAddressResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AllocateAddressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AllocateAddressResult allocateAddress()
func (jbobject *ServicesEc2AmazonEC2Client) AllocateAddress() *ServicesEc2ModelAllocateAddressResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "allocateAddress", "com/amazonaws/services/ec2/model/AllocateAddressResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAllocateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AllocateHostsResult allocateHosts(com.amazonaws.services.ec2.model.AllocateHostsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AllocateHosts(a ServicesEc2ModelAllocateHostsRequestInterface) *ServicesEc2ModelAllocateHostsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "allocateHosts", "com/amazonaws/services/ec2/model/AllocateHostsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AllocateHostsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAllocateHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void assignPrivateIpAddresses(com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AssignPrivateIpAddresses(a ServicesEc2ModelAssignPrivateIpAddressesRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "assignPrivateIpAddresses", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AssociateAddressResult associateAddress(com.amazonaws.services.ec2.model.AssociateAddressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AssociateAddress(a ServicesEc2ModelAssociateAddressRequestInterface) *ServicesEc2ModelAssociateAddressResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "associateAddress", "com/amazonaws/services/ec2/model/AssociateAddressResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AssociateAddressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAssociateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void associateDhcpOptions(com.amazonaws.services.ec2.model.AssociateDhcpOptionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AssociateDhcpOptions(a ServicesEc2ModelAssociateDhcpOptionsRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "associateDhcpOptions", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AssociateDhcpOptionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AssociateRouteTableResult associateRouteTable(com.amazonaws.services.ec2.model.AssociateRouteTableRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AssociateRouteTable(a ServicesEc2ModelAssociateRouteTableRequestInterface) *ServicesEc2ModelAssociateRouteTableResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "associateRouteTable", "com/amazonaws/services/ec2/model/AssociateRouteTableResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AssociateRouteTableRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAssociateRouteTableResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcResult attachClassicLinkVpc(com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AttachClassicLinkVpc(a ServicesEc2ModelAttachClassicLinkVpcRequestInterface) *ServicesEc2ModelAttachClassicLinkVpcResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attachClassicLinkVpc", "com/amazonaws/services/ec2/model/AttachClassicLinkVpcResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAttachClassicLinkVpcResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void attachInternetGateway(com.amazonaws.services.ec2.model.AttachInternetGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AttachInternetGateway(a ServicesEc2ModelAttachInternetGatewayRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "attachInternetGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachInternetGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult attachNetworkInterface(com.amazonaws.services.ec2.model.AttachNetworkInterfaceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AttachNetworkInterface(a ServicesEc2ModelAttachNetworkInterfaceRequestInterface) *ServicesEc2ModelAttachNetworkInterfaceResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attachNetworkInterface", "com/amazonaws/services/ec2/model/AttachNetworkInterfaceResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachNetworkInterfaceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAttachNetworkInterfaceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AttachVolumeResult attachVolume(com.amazonaws.services.ec2.model.AttachVolumeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AttachVolume(a ServicesEc2ModelAttachVolumeRequestInterface) *ServicesEc2ModelAttachVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attachVolume", "com/amazonaws/services/ec2/model/AttachVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachVolumeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAttachVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AttachVpnGatewayResult attachVpnGateway(com.amazonaws.services.ec2.model.AttachVpnGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AttachVpnGateway(a ServicesEc2ModelAttachVpnGatewayRequestInterface) *ServicesEc2ModelAttachVpnGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attachVpnGateway", "com/amazonaws/services/ec2/model/AttachVpnGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachVpnGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelAttachVpnGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void authorizeSecurityGroupEgress(com.amazonaws.services.ec2.model.AuthorizeSecurityGroupEgressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AuthorizeSecurityGroupEgress(a ServicesEc2ModelAuthorizeSecurityGroupEgressRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "authorizeSecurityGroupEgress", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AuthorizeSecurityGroupEgressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void authorizeSecurityGroupIngress(com.amazonaws.services.ec2.model.AuthorizeSecurityGroupIngressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) AuthorizeSecurityGroupIngress(a ServicesEc2ModelAuthorizeSecurityGroupIngressRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "authorizeSecurityGroupIngress", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AuthorizeSecurityGroupIngressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.BundleInstanceResult bundleInstance(com.amazonaws.services.ec2.model.BundleInstanceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) BundleInstance(a ServicesEc2ModelBundleInstanceRequestInterface) *ServicesEc2ModelBundleInstanceResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "bundleInstance", "com/amazonaws/services/ec2/model/BundleInstanceResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleInstanceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelBundleInstanceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelBundleTaskResult cancelBundleTask(com.amazonaws.services.ec2.model.CancelBundleTaskRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelBundleTask(a ServicesEc2ModelCancelBundleTaskRequestInterface) *ServicesEc2ModelCancelBundleTaskResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cancelBundleTask", "com/amazonaws/services/ec2/model/CancelBundleTaskResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelBundleTaskRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCancelBundleTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void cancelConversionTask(com.amazonaws.services.ec2.model.CancelConversionTaskRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelConversionTask(a ServicesEc2ModelCancelConversionTaskRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelConversionTask", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelConversionTaskRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void cancelExportTask(com.amazonaws.services.ec2.model.CancelExportTaskRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelExportTask(a ServicesEc2ModelCancelExportTaskRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelExportTask", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelExportTaskRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelImportTaskResult cancelImportTask(com.amazonaws.services.ec2.model.CancelImportTaskRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelImportTask2(a ServicesEc2ModelCancelImportTaskRequestInterface) *ServicesEc2ModelCancelImportTaskResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cancelImportTask", "com/amazonaws/services/ec2/model/CancelImportTaskResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelImportTaskRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCancelImportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelImportTaskResult cancelImportTask()
func (jbobject *ServicesEc2AmazonEC2Client) CancelImportTask() *ServicesEc2ModelCancelImportTaskResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cancelImportTask", "com/amazonaws/services/ec2/model/CancelImportTaskResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCancelImportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingResult cancelReservedInstancesListing(com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelReservedInstancesListing(a ServicesEc2ModelCancelReservedInstancesListingRequestInterface) *ServicesEc2ModelCancelReservedInstancesListingResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cancelReservedInstancesListing", "com/amazonaws/services/ec2/model/CancelReservedInstancesListingResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelReservedInstancesListingRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCancelReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsResult cancelSpotFleetRequests(com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelSpotFleetRequests(a ServicesEc2ModelCancelSpotFleetRequestsRequestInterface) *ServicesEc2ModelCancelSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cancelSpotFleetRequests", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult cancelSpotInstanceRequests(com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CancelSpotInstanceRequests(a ServicesEc2ModelCancelSpotInstanceRequestsRequestInterface) *ServicesEc2ModelCancelSpotInstanceRequestsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cancelSpotInstanceRequests", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult confirmProductInstance(com.amazonaws.services.ec2.model.ConfirmProductInstanceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ConfirmProductInstance(a ServicesEc2ModelConfirmProductInstanceRequestInterface) *ServicesEc2ModelConfirmProductInstanceResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "confirmProductInstance", "com/amazonaws/services/ec2/model/ConfirmProductInstanceResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ConfirmProductInstanceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelConfirmProductInstanceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CopyImageResult copyImage(com.amazonaws.services.ec2.model.CopyImageRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CopyImage(a ServicesEc2ModelCopyImageRequestInterface) *ServicesEc2ModelCopyImageResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copyImage", "com/amazonaws/services/ec2/model/CopyImageResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CopyImageRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCopyImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CopySnapshotResult copySnapshot(com.amazonaws.services.ec2.model.CopySnapshotRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CopySnapshot(a ServicesEc2ModelCopySnapshotRequestInterface) *ServicesEc2ModelCopySnapshotResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copySnapshot", "com/amazonaws/services/ec2/model/CopySnapshotResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CopySnapshotRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCopySnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult createCustomerGateway(com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateCustomerGateway(a ServicesEc2ModelCreateCustomerGatewayRequestInterface) *ServicesEc2ModelCreateCustomerGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createCustomerGateway", "com/amazonaws/services/ec2/model/CreateCustomerGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateCustomerGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult createDhcpOptions(com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateDhcpOptions(a ServicesEc2ModelCreateDhcpOptionsRequestInterface) *ServicesEc2ModelCreateDhcpOptionsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createDhcpOptions", "com/amazonaws/services/ec2/model/CreateDhcpOptionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateDhcpOptionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateDhcpOptionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsResult createFlowLogs(com.amazonaws.services.ec2.model.CreateFlowLogsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateFlowLogs(a ServicesEc2ModelCreateFlowLogsRequestInterface) *ServicesEc2ModelCreateFlowLogsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createFlowLogs", "com/amazonaws/services/ec2/model/CreateFlowLogsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateFlowLogsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateImageResult createImage(com.amazonaws.services.ec2.model.CreateImageRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateImage(a ServicesEc2ModelCreateImageRequestInterface) *ServicesEc2ModelCreateImageResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createImage", "com/amazonaws/services/ec2/model/CreateImageResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateImageRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult createInstanceExportTask(com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateInstanceExportTask(a ServicesEc2ModelCreateInstanceExportTaskRequestInterface) *ServicesEc2ModelCreateInstanceExportTaskResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createInstanceExportTask", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult createInternetGateway(com.amazonaws.services.ec2.model.CreateInternetGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateInternetGateway2(a ServicesEc2ModelCreateInternetGatewayRequestInterface) *ServicesEc2ModelCreateInternetGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createInternetGateway", "com/amazonaws/services/ec2/model/CreateInternetGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateInternetGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateInternetGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult createInternetGateway()
func (jbobject *ServicesEc2AmazonEC2Client) CreateInternetGateway() *ServicesEc2ModelCreateInternetGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createInternetGateway", "com/amazonaws/services/ec2/model/CreateInternetGatewayResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateInternetGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateKeyPairResult createKeyPair(com.amazonaws.services.ec2.model.CreateKeyPairRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateKeyPair(a ServicesEc2ModelCreateKeyPairRequestInterface) *ServicesEc2ModelCreateKeyPairResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createKeyPair", "com/amazonaws/services/ec2/model/CreateKeyPairResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateKeyPairRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNatGatewayResult createNatGateway(com.amazonaws.services.ec2.model.CreateNatGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateNatGateway(a ServicesEc2ModelCreateNatGatewayRequestInterface) *ServicesEc2ModelCreateNatGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createNatGateway", "com/amazonaws/services/ec2/model/CreateNatGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateNatGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclResult createNetworkAcl(com.amazonaws.services.ec2.model.CreateNetworkAclRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateNetworkAcl(a ServicesEc2ModelCreateNetworkAclRequestInterface) *ServicesEc2ModelCreateNetworkAclResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createNetworkAcl", "com/amazonaws/services/ec2/model/CreateNetworkAclResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateNetworkAclRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateNetworkAclResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void createNetworkAclEntry(com.amazonaws.services.ec2.model.CreateNetworkAclEntryRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateNetworkAclEntry(a ServicesEc2ModelCreateNetworkAclEntryRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "createNetworkAclEntry", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateNetworkAclEntryRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult createNetworkInterface(com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateNetworkInterface(a ServicesEc2ModelCreateNetworkInterfaceRequestInterface) *ServicesEc2ModelCreateNetworkInterfaceResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createNetworkInterface", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void createPlacementGroup(com.amazonaws.services.ec2.model.CreatePlacementGroupRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreatePlacementGroup(a ServicesEc2ModelCreatePlacementGroupRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "createPlacementGroup", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreatePlacementGroupRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateReservedInstancesListingResult createReservedInstancesListing(com.amazonaws.services.ec2.model.CreateReservedInstancesListingRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateReservedInstancesListing(a ServicesEc2ModelCreateReservedInstancesListingRequestInterface) *ServicesEc2ModelCreateReservedInstancesListingResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createReservedInstancesListing", "com/amazonaws/services/ec2/model/CreateReservedInstancesListingResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateReservedInstancesListingRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateReservedInstancesListingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateRouteResult createRoute(com.amazonaws.services.ec2.model.CreateRouteRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateRoute(a ServicesEc2ModelCreateRouteRequestInterface) *ServicesEc2ModelCreateRouteResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createRoute", "com/amazonaws/services/ec2/model/CreateRouteResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateRouteRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateRouteResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateRouteTableResult createRouteTable(com.amazonaws.services.ec2.model.CreateRouteTableRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateRouteTable(a ServicesEc2ModelCreateRouteTableRequestInterface) *ServicesEc2ModelCreateRouteTableResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createRouteTable", "com/amazonaws/services/ec2/model/CreateRouteTableResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateRouteTableRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateRouteTableResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSecurityGroupResult createSecurityGroup(com.amazonaws.services.ec2.model.CreateSecurityGroupRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateSecurityGroup(a ServicesEc2ModelCreateSecurityGroupRequestInterface) *ServicesEc2ModelCreateSecurityGroupResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createSecurityGroup", "com/amazonaws/services/ec2/model/CreateSecurityGroupResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateSecurityGroupRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateSecurityGroupResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSnapshotResult createSnapshot(com.amazonaws.services.ec2.model.CreateSnapshotRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateSnapshot(a ServicesEc2ModelCreateSnapshotRequestInterface) *ServicesEc2ModelCreateSnapshotResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createSnapshot", "com/amazonaws/services/ec2/model/CreateSnapshotResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateSnapshotRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateSnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionResult createSpotDatafeedSubscription(com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateSpotDatafeedSubscription(a ServicesEc2ModelCreateSpotDatafeedSubscriptionRequestInterface) *ServicesEc2ModelCreateSpotDatafeedSubscriptionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createSpotDatafeedSubscription", "com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSubnetResult createSubnet(com.amazonaws.services.ec2.model.CreateSubnetRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateSubnet(a ServicesEc2ModelCreateSubnetRequestInterface) *ServicesEc2ModelCreateSubnetResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createSubnet", "com/amazonaws/services/ec2/model/CreateSubnetResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateSubnetRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateSubnetResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void createTags(com.amazonaws.services.ec2.model.CreateTagsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateTags(a ServicesEc2ModelCreateTagsRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "createTags", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateTagsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVolumeResult createVolume(com.amazonaws.services.ec2.model.CreateVolumeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVolume(a ServicesEc2ModelCreateVolumeRequestInterface) *ServicesEc2ModelCreateVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVolume", "com/amazonaws/services/ec2/model/CreateVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVolumeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcResult createVpc(com.amazonaws.services.ec2.model.CreateVpcRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpc(a ServicesEc2ModelCreateVpcRequestInterface) *ServicesEc2ModelCreateVpcResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVpc", "com/amazonaws/services/ec2/model/CreateVpcResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVpcRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVpcResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult createVpcEndpoint(com.amazonaws.services.ec2.model.CreateVpcEndpointRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpcEndpoint(a ServicesEc2ModelCreateVpcEndpointRequestInterface) *ServicesEc2ModelCreateVpcEndpointResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVpcEndpoint", "com/amazonaws/services/ec2/model/CreateVpcEndpointResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVpcEndpointRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult createVpcPeeringConnection(com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpcPeeringConnection2(a ServicesEc2ModelCreateVpcPeeringConnectionRequestInterface) *ServicesEc2ModelCreateVpcPeeringConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVpcPeeringConnection", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult createVpcPeeringConnection()
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpcPeeringConnection() *ServicesEc2ModelCreateVpcPeeringConnectionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVpcPeeringConnection", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult createVpnConnection(com.amazonaws.services.ec2.model.CreateVpnConnectionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpnConnection(a ServicesEc2ModelCreateVpnConnectionRequestInterface) *ServicesEc2ModelCreateVpnConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVpnConnection", "com/amazonaws/services/ec2/model/CreateVpnConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVpnConnectionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVpnConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void createVpnConnectionRoute(com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpnConnectionRoute(a ServicesEc2ModelCreateVpnConnectionRouteRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "createVpnConnectionRoute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVpnConnectionRouteRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult createVpnGateway(com.amazonaws.services.ec2.model.CreateVpnGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) CreateVpnGateway(a ServicesEc2ModelCreateVpnGatewayRequestInterface) *ServicesEc2ModelCreateVpnGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createVpnGateway", "com/amazonaws/services/ec2/model/CreateVpnGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVpnGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelCreateVpnGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void deleteCustomerGateway(com.amazonaws.services.ec2.model.DeleteCustomerGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteCustomerGateway(a ServicesEc2ModelDeleteCustomerGatewayRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteCustomerGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteCustomerGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteDhcpOptions(com.amazonaws.services.ec2.model.DeleteDhcpOptionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteDhcpOptions(a ServicesEc2ModelDeleteDhcpOptionsRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteDhcpOptions", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteDhcpOptionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DeleteFlowLogsResult deleteFlowLogs(com.amazonaws.services.ec2.model.DeleteFlowLogsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteFlowLogs(a ServicesEc2ModelDeleteFlowLogsRequestInterface) *ServicesEc2ModelDeleteFlowLogsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "deleteFlowLogs", "com/amazonaws/services/ec2/model/DeleteFlowLogsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteFlowLogsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDeleteFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void deleteInternetGateway(com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteInternetGateway(a ServicesEc2ModelDeleteInternetGatewayRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteInternetGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteInternetGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteKeyPair(com.amazonaws.services.ec2.model.DeleteKeyPairRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteKeyPair(a ServicesEc2ModelDeleteKeyPairRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteKeyPair", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteKeyPairRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult deleteNatGateway(com.amazonaws.services.ec2.model.DeleteNatGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteNatGateway(a ServicesEc2ModelDeleteNatGatewayRequestInterface) *ServicesEc2ModelDeleteNatGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "deleteNatGateway", "com/amazonaws/services/ec2/model/DeleteNatGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteNatGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDeleteNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void deleteNetworkAcl(com.amazonaws.services.ec2.model.DeleteNetworkAclRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteNetworkAcl(a ServicesEc2ModelDeleteNetworkAclRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteNetworkAcl", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteNetworkAclRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteNetworkAclEntry(com.amazonaws.services.ec2.model.DeleteNetworkAclEntryRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteNetworkAclEntry(a ServicesEc2ModelDeleteNetworkAclEntryRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteNetworkAclEntry", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteNetworkAclEntryRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteNetworkInterface(com.amazonaws.services.ec2.model.DeleteNetworkInterfaceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteNetworkInterface(a ServicesEc2ModelDeleteNetworkInterfaceRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteNetworkInterface", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteNetworkInterfaceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deletePlacementGroup(com.amazonaws.services.ec2.model.DeletePlacementGroupRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeletePlacementGroup(a ServicesEc2ModelDeletePlacementGroupRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deletePlacementGroup", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeletePlacementGroupRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteRoute(com.amazonaws.services.ec2.model.DeleteRouteRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteRoute(a ServicesEc2ModelDeleteRouteRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteRoute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteRouteRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteRouteTable(com.amazonaws.services.ec2.model.DeleteRouteTableRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteRouteTable(a ServicesEc2ModelDeleteRouteTableRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteRouteTable", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteRouteTableRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteSecurityGroup(com.amazonaws.services.ec2.model.DeleteSecurityGroupRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteSecurityGroup(a ServicesEc2ModelDeleteSecurityGroupRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteSecurityGroup", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteSecurityGroupRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteSnapshot(com.amazonaws.services.ec2.model.DeleteSnapshotRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteSnapshot(a ServicesEc2ModelDeleteSnapshotRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteSnapshot", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteSnapshotRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteSpotDatafeedSubscription(com.amazonaws.services.ec2.model.DeleteSpotDatafeedSubscriptionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteSpotDatafeedSubscription2(a ServicesEc2ModelDeleteSpotDatafeedSubscriptionRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteSpotDatafeedSubscription", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteSpotDatafeedSubscriptionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteSpotDatafeedSubscription()
func (jbobject *ServicesEc2AmazonEC2Client) DeleteSpotDatafeedSubscription()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteSpotDatafeedSubscription", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void deleteSubnet(com.amazonaws.services.ec2.model.DeleteSubnetRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteSubnet(a ServicesEc2ModelDeleteSubnetRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteSubnet", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteSubnetRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteTags(com.amazonaws.services.ec2.model.DeleteTagsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteTags(a ServicesEc2ModelDeleteTagsRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteTags", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteTagsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteVolume(com.amazonaws.services.ec2.model.DeleteVolumeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVolume(a ServicesEc2ModelDeleteVolumeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVolume", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVolumeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteVpc(com.amazonaws.services.ec2.model.DeleteVpcRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVpc(a ServicesEc2ModelDeleteVpcRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVpc", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVpcRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DeleteVpcEndpointsResult deleteVpcEndpoints(com.amazonaws.services.ec2.model.DeleteVpcEndpointsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVpcEndpoints(a ServicesEc2ModelDeleteVpcEndpointsRequestInterface) *ServicesEc2ModelDeleteVpcEndpointsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVpcEndpoints", "com/amazonaws/services/ec2/model/DeleteVpcEndpointsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVpcEndpointsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDeleteVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionResult deleteVpcPeeringConnection(com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVpcPeeringConnection(a ServicesEc2ModelDeleteVpcPeeringConnectionRequestInterface) *ServicesEc2ModelDeleteVpcPeeringConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVpcPeeringConnection", "com/amazonaws/services/ec2/model/DeleteVpcPeeringConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVpcPeeringConnectionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDeleteVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void deleteVpnConnection(com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVpnConnection(a ServicesEc2ModelDeleteVpnConnectionRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVpnConnection", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVpnConnectionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteVpnConnectionRoute(com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVpnConnectionRoute(a ServicesEc2ModelDeleteVpnConnectionRouteRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVpnConnectionRoute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVpnConnectionRouteRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deleteVpnGateway(com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeleteVpnGateway(a ServicesEc2ModelDeleteVpnGatewayRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteVpnGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeleteVpnGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void deregisterImage(com.amazonaws.services.ec2.model.DeregisterImageRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DeregisterImage(a ServicesEc2ModelDeregisterImageRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deregisterImage", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeregisterImageRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult describeAccountAttributes(com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeAccountAttributes2(a ServicesEc2ModelDescribeAccountAttributesRequestInterface) *ServicesEc2ModelDescribeAccountAttributesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeAccountAttributes", "com/amazonaws/services/ec2/model/DescribeAccountAttributesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeAccountAttributesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeAccountAttributesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult describeAccountAttributes()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeAccountAttributes() *ServicesEc2ModelDescribeAccountAttributesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeAccountAttributes", "com/amazonaws/services/ec2/model/DescribeAccountAttributesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeAccountAttributesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAddressesResult describeAddresses(com.amazonaws.services.ec2.model.DescribeAddressesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeAddresses2(a ServicesEc2ModelDescribeAddressesRequestInterface) *ServicesEc2ModelDescribeAddressesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeAddresses", "com/amazonaws/services/ec2/model/DescribeAddressesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeAddressesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAddressesResult describeAddresses()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeAddresses() *ServicesEc2ModelDescribeAddressesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeAddresses", "com/amazonaws/services/ec2/model/DescribeAddressesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult describeAvailabilityZones(com.amazonaws.services.ec2.model.DescribeAvailabilityZonesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeAvailabilityZones2(a ServicesEc2ModelDescribeAvailabilityZonesRequestInterface) *ServicesEc2ModelDescribeAvailabilityZonesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeAvailabilityZones", "com/amazonaws/services/ec2/model/DescribeAvailabilityZonesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeAvailabilityZonesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeAvailabilityZonesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAvailabilityZonesResult describeAvailabilityZones()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeAvailabilityZones() *ServicesEc2ModelDescribeAvailabilityZonesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeAvailabilityZones", "com/amazonaws/services/ec2/model/DescribeAvailabilityZonesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeAvailabilityZonesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeBundleTasksResult describeBundleTasks(com.amazonaws.services.ec2.model.DescribeBundleTasksRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeBundleTasks2(a ServicesEc2ModelDescribeBundleTasksRequestInterface) *ServicesEc2ModelDescribeBundleTasksResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeBundleTasks", "com/amazonaws/services/ec2/model/DescribeBundleTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeBundleTasksRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeBundleTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeBundleTasksResult describeBundleTasks()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeBundleTasks() *ServicesEc2ModelDescribeBundleTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeBundleTasks", "com/amazonaws/services/ec2/model/DescribeBundleTasksResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeBundleTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult describeClassicLinkInstances(com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeClassicLinkInstances2(a ServicesEc2ModelDescribeClassicLinkInstancesRequestInterface) *ServicesEc2ModelDescribeClassicLinkInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeClassicLinkInstances", "com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult describeClassicLinkInstances()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeClassicLinkInstances() *ServicesEc2ModelDescribeClassicLinkInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeClassicLinkInstances", "com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult describeConversionTasks(com.amazonaws.services.ec2.model.DescribeConversionTasksRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeConversionTasks2(a ServicesEc2ModelDescribeConversionTasksRequestInterface) *ServicesEc2ModelDescribeConversionTasksResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeConversionTasks", "com/amazonaws/services/ec2/model/DescribeConversionTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeConversionTasksRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeConversionTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult describeConversionTasks()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeConversionTasks() *ServicesEc2ModelDescribeConversionTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeConversionTasks", "com/amazonaws/services/ec2/model/DescribeConversionTasksResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeConversionTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult describeCustomerGateways(com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeCustomerGateways2(a ServicesEc2ModelDescribeCustomerGatewaysRequestInterface) *ServicesEc2ModelDescribeCustomerGatewaysResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeCustomerGateways", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult describeCustomerGateways()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeCustomerGateways() *ServicesEc2ModelDescribeCustomerGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeCustomerGateways", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeDhcpOptionsResult describeDhcpOptions(com.amazonaws.services.ec2.model.DescribeDhcpOptionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeDhcpOptions2(a ServicesEc2ModelDescribeDhcpOptionsRequestInterface) *ServicesEc2ModelDescribeDhcpOptionsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeDhcpOptions", "com/amazonaws/services/ec2/model/DescribeDhcpOptionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeDhcpOptionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeDhcpOptionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeDhcpOptionsResult describeDhcpOptions()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeDhcpOptions() *ServicesEc2ModelDescribeDhcpOptionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeDhcpOptions", "com/amazonaws/services/ec2/model/DescribeDhcpOptionsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeDhcpOptionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeExportTasksResult describeExportTasks(com.amazonaws.services.ec2.model.DescribeExportTasksRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeExportTasks2(a ServicesEc2ModelDescribeExportTasksRequestInterface) *ServicesEc2ModelDescribeExportTasksResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeExportTasks", "com/amazonaws/services/ec2/model/DescribeExportTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeExportTasksRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeExportTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeExportTasksResult describeExportTasks()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeExportTasks() *ServicesEc2ModelDescribeExportTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeExportTasks", "com/amazonaws/services/ec2/model/DescribeExportTasksResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeExportTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult describeFlowLogs(com.amazonaws.services.ec2.model.DescribeFlowLogsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeFlowLogs2(a ServicesEc2ModelDescribeFlowLogsRequestInterface) *ServicesEc2ModelDescribeFlowLogsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeFlowLogs", "com/amazonaws/services/ec2/model/DescribeFlowLogsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeFlowLogsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult describeFlowLogs()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeFlowLogs() *ServicesEc2ModelDescribeFlowLogsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeFlowLogs", "com/amazonaws/services/ec2/model/DescribeFlowLogsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeHostsResult describeHosts(com.amazonaws.services.ec2.model.DescribeHostsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeHosts2(a ServicesEc2ModelDescribeHostsRequestInterface) *ServicesEc2ModelDescribeHostsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeHosts", "com/amazonaws/services/ec2/model/DescribeHostsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeHostsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeHostsResult describeHosts()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeHosts() *ServicesEc2ModelDescribeHostsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeHosts", "com/amazonaws/services/ec2/model/DescribeHostsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeIdFormatResult describeIdFormat(com.amazonaws.services.ec2.model.DescribeIdFormatRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeIdFormat2(a ServicesEc2ModelDescribeIdFormatRequestInterface) *ServicesEc2ModelDescribeIdFormatResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeIdFormat", "com/amazonaws/services/ec2/model/DescribeIdFormatResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeIdFormatRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeIdFormatResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeIdFormatResult describeIdFormat()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeIdFormat() *ServicesEc2ModelDescribeIdFormatResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeIdFormat", "com/amazonaws/services/ec2/model/DescribeIdFormatResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeIdFormatResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult describeImageAttribute(com.amazonaws.services.ec2.model.DescribeImageAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImageAttribute(a ServicesEc2ModelDescribeImageAttributeRequestInterface) *ServicesEc2ModelDescribeImageAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImageAttribute", "com/amazonaws/services/ec2/model/DescribeImageAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeImageAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImageAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImagesResult describeImages(com.amazonaws.services.ec2.model.DescribeImagesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImages2(a ServicesEc2ModelDescribeImagesRequestInterface) *ServicesEc2ModelDescribeImagesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImages", "com/amazonaws/services/ec2/model/DescribeImagesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeImagesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImagesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImagesResult describeImages()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImages() *ServicesEc2ModelDescribeImagesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImages", "com/amazonaws/services/ec2/model/DescribeImagesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImagesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult describeImportImageTasks(com.amazonaws.services.ec2.model.DescribeImportImageTasksRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImportImageTasks2(a ServicesEc2ModelDescribeImportImageTasksRequestInterface) *ServicesEc2ModelDescribeImportImageTasksResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImportImageTasks", "com/amazonaws/services/ec2/model/DescribeImportImageTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeImportImageTasksRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult describeImportImageTasks()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImportImageTasks() *ServicesEc2ModelDescribeImportImageTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImportImageTasks", "com/amazonaws/services/ec2/model/DescribeImportImageTasksResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult describeImportSnapshotTasks(com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImportSnapshotTasks2(a ServicesEc2ModelDescribeImportSnapshotTasksRequestInterface) *ServicesEc2ModelDescribeImportSnapshotTasksResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImportSnapshotTasks", "com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult describeImportSnapshotTasks()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeImportSnapshotTasks() *ServicesEc2ModelDescribeImportSnapshotTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeImportSnapshotTasks", "com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult describeInstanceAttribute(com.amazonaws.services.ec2.model.DescribeInstanceAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInstanceAttribute(a ServicesEc2ModelDescribeInstanceAttributeRequestInterface) *ServicesEc2ModelDescribeInstanceAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInstanceAttribute", "com/amazonaws/services/ec2/model/DescribeInstanceAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeInstanceAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInstanceAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult describeInstanceStatus(com.amazonaws.services.ec2.model.DescribeInstanceStatusRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInstanceStatus2(a ServicesEc2ModelDescribeInstanceStatusRequestInterface) *ServicesEc2ModelDescribeInstanceStatusResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInstanceStatus", "com/amazonaws/services/ec2/model/DescribeInstanceStatusResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeInstanceStatusRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult describeInstanceStatus()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInstanceStatus() *ServicesEc2ModelDescribeInstanceStatusResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInstanceStatus", "com/amazonaws/services/ec2/model/DescribeInstanceStatusResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstancesResult describeInstances(com.amazonaws.services.ec2.model.DescribeInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInstances2(a ServicesEc2ModelDescribeInstancesRequestInterface) *ServicesEc2ModelDescribeInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInstances", "com/amazonaws/services/ec2/model/DescribeInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstancesResult describeInstances()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInstances() *ServicesEc2ModelDescribeInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInstances", "com/amazonaws/services/ec2/model/DescribeInstancesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult describeInternetGateways(com.amazonaws.services.ec2.model.DescribeInternetGatewaysRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInternetGateways2(a ServicesEc2ModelDescribeInternetGatewaysRequestInterface) *ServicesEc2ModelDescribeInternetGatewaysResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInternetGateways", "com/amazonaws/services/ec2/model/DescribeInternetGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeInternetGatewaysRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInternetGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult describeInternetGateways()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeInternetGateways() *ServicesEc2ModelDescribeInternetGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeInternetGateways", "com/amazonaws/services/ec2/model/DescribeInternetGatewaysResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeInternetGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult describeKeyPairs(com.amazonaws.services.ec2.model.DescribeKeyPairsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeKeyPairs2(a ServicesEc2ModelDescribeKeyPairsRequestInterface) *ServicesEc2ModelDescribeKeyPairsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeKeyPairs", "com/amazonaws/services/ec2/model/DescribeKeyPairsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeKeyPairsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeKeyPairsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult describeKeyPairs()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeKeyPairs() *ServicesEc2ModelDescribeKeyPairsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeKeyPairs", "com/amazonaws/services/ec2/model/DescribeKeyPairsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeKeyPairsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult describeMovingAddresses(com.amazonaws.services.ec2.model.DescribeMovingAddressesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeMovingAddresses2(a ServicesEc2ModelDescribeMovingAddressesRequestInterface) *ServicesEc2ModelDescribeMovingAddressesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeMovingAddresses", "com/amazonaws/services/ec2/model/DescribeMovingAddressesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeMovingAddressesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult describeMovingAddresses()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeMovingAddresses() *ServicesEc2ModelDescribeMovingAddressesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeMovingAddresses", "com/amazonaws/services/ec2/model/DescribeMovingAddressesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult describeNatGateways(com.amazonaws.services.ec2.model.DescribeNatGatewaysRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeNatGateways(a ServicesEc2ModelDescribeNatGatewaysRequestInterface) *ServicesEc2ModelDescribeNatGatewaysResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeNatGateways", "com/amazonaws/services/ec2/model/DescribeNatGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeNatGatewaysRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeNatGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult describeNetworkAcls(com.amazonaws.services.ec2.model.DescribeNetworkAclsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeNetworkAcls2(a ServicesEc2ModelDescribeNetworkAclsRequestInterface) *ServicesEc2ModelDescribeNetworkAclsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeNetworkAcls", "com/amazonaws/services/ec2/model/DescribeNetworkAclsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeNetworkAclsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeNetworkAclsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult describeNetworkAcls()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeNetworkAcls() *ServicesEc2ModelDescribeNetworkAclsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeNetworkAcls", "com/amazonaws/services/ec2/model/DescribeNetworkAclsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeNetworkAclsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeResult describeNetworkInterfaceAttribute(com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeNetworkInterfaceAttribute(a ServicesEc2ModelDescribeNetworkInterfaceAttributeRequestInterface) *ServicesEc2ModelDescribeNetworkInterfaceAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeNetworkInterfaceAttribute", "com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeNetworkInterfaceAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult describeNetworkInterfaces(com.amazonaws.services.ec2.model.DescribeNetworkInterfacesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeNetworkInterfaces2(a ServicesEc2ModelDescribeNetworkInterfacesRequestInterface) *ServicesEc2ModelDescribeNetworkInterfacesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeNetworkInterfaces", "com/amazonaws/services/ec2/model/DescribeNetworkInterfacesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeNetworkInterfacesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeNetworkInterfacesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult describeNetworkInterfaces()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeNetworkInterfaces() *ServicesEc2ModelDescribeNetworkInterfacesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeNetworkInterfaces", "com/amazonaws/services/ec2/model/DescribeNetworkInterfacesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeNetworkInterfacesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult describePlacementGroups(com.amazonaws.services.ec2.model.DescribePlacementGroupsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribePlacementGroups2(a ServicesEc2ModelDescribePlacementGroupsRequestInterface) *ServicesEc2ModelDescribePlacementGroupsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describePlacementGroups", "com/amazonaws/services/ec2/model/DescribePlacementGroupsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribePlacementGroupsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribePlacementGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribePlacementGroupsResult describePlacementGroups()
func (jbobject *ServicesEc2AmazonEC2Client) DescribePlacementGroups() *ServicesEc2ModelDescribePlacementGroupsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describePlacementGroups", "com/amazonaws/services/ec2/model/DescribePlacementGroupsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribePlacementGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult describePrefixLists(com.amazonaws.services.ec2.model.DescribePrefixListsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribePrefixLists2(a ServicesEc2ModelDescribePrefixListsRequestInterface) *ServicesEc2ModelDescribePrefixListsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describePrefixLists", "com/amazonaws/services/ec2/model/DescribePrefixListsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribePrefixListsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribePrefixListsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribePrefixListsResult describePrefixLists()
func (jbobject *ServicesEc2AmazonEC2Client) DescribePrefixLists() *ServicesEc2ModelDescribePrefixListsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describePrefixLists", "com/amazonaws/services/ec2/model/DescribePrefixListsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribePrefixListsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeRegionsResult describeRegions(com.amazonaws.services.ec2.model.DescribeRegionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeRegions2(a ServicesEc2ModelDescribeRegionsRequestInterface) *ServicesEc2ModelDescribeRegionsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeRegions", "com/amazonaws/services/ec2/model/DescribeRegionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeRegionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeRegionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeRegionsResult describeRegions()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeRegions() *ServicesEc2ModelDescribeRegionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeRegions", "com/amazonaws/services/ec2/model/DescribeRegionsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeRegionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult describeReservedInstances(com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstances2(a ServicesEc2ModelDescribeReservedInstancesRequestInterface) *ServicesEc2ModelDescribeReservedInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstances", "com/amazonaws/services/ec2/model/DescribeReservedInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult describeReservedInstances()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstances() *ServicesEc2ModelDescribeReservedInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstances", "com/amazonaws/services/ec2/model/DescribeReservedInstancesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesListingsResult describeReservedInstancesListings(com.amazonaws.services.ec2.model.DescribeReservedInstancesListingsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstancesListings2(a ServicesEc2ModelDescribeReservedInstancesListingsRequestInterface) *ServicesEc2ModelDescribeReservedInstancesListingsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstancesListings", "com/amazonaws/services/ec2/model/DescribeReservedInstancesListingsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeReservedInstancesListingsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesListingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesListingsResult describeReservedInstancesListings()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstancesListings() *ServicesEc2ModelDescribeReservedInstancesListingsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstancesListings", "com/amazonaws/services/ec2/model/DescribeReservedInstancesListingsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesListingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult describeReservedInstancesModifications(com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstancesModifications2(a ServicesEc2ModelDescribeReservedInstancesModificationsRequestInterface) *ServicesEc2ModelDescribeReservedInstancesModificationsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstancesModifications", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult describeReservedInstancesModifications()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstancesModifications() *ServicesEc2ModelDescribeReservedInstancesModificationsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstancesModifications", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult describeReservedInstancesOfferings(com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstancesOfferings2(a ServicesEc2ModelDescribeReservedInstancesOfferingsRequestInterface) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstancesOfferings", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult describeReservedInstancesOfferings()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeReservedInstancesOfferings() *ServicesEc2ModelDescribeReservedInstancesOfferingsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeReservedInstancesOfferings", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult describeRouteTables(com.amazonaws.services.ec2.model.DescribeRouteTablesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeRouteTables2(a ServicesEc2ModelDescribeRouteTablesRequestInterface) *ServicesEc2ModelDescribeRouteTablesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeRouteTables", "com/amazonaws/services/ec2/model/DescribeRouteTablesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeRouteTablesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeRouteTablesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult describeRouteTables()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeRouteTables() *ServicesEc2ModelDescribeRouteTablesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeRouteTables", "com/amazonaws/services/ec2/model/DescribeRouteTablesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeRouteTablesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult describeScheduledInstanceAvailability(com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeScheduledInstanceAvailability(a ServicesEc2ModelDescribeScheduledInstanceAvailabilityRequestInterface) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeScheduledInstanceAvailability", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult describeScheduledInstances(com.amazonaws.services.ec2.model.DescribeScheduledInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeScheduledInstances(a ServicesEc2ModelDescribeScheduledInstancesRequestInterface) *ServicesEc2ModelDescribeScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeScheduledInstances", "com/amazonaws/services/ec2/model/DescribeScheduledInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeScheduledInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult describeSecurityGroups(com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSecurityGroups2(a ServicesEc2ModelDescribeSecurityGroupsRequestInterface) *ServicesEc2ModelDescribeSecurityGroupsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSecurityGroups", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult describeSecurityGroups()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSecurityGroups() *ServicesEc2ModelDescribeSecurityGroupsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSecurityGroups", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult describeSnapshotAttribute(com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSnapshotAttribute(a ServicesEc2ModelDescribeSnapshotAttributeRequestInterface) *ServicesEc2ModelDescribeSnapshotAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSnapshotAttribute", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult describeSnapshots(com.amazonaws.services.ec2.model.DescribeSnapshotsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSnapshots2(a ServicesEc2ModelDescribeSnapshotsRequestInterface) *ServicesEc2ModelDescribeSnapshotsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSnapshots", "com/amazonaws/services/ec2/model/DescribeSnapshotsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSnapshotsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSnapshotsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult describeSnapshots()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSnapshots() *ServicesEc2ModelDescribeSnapshotsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSnapshots", "com/amazonaws/services/ec2/model/DescribeSnapshotsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSnapshotsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult describeSpotDatafeedSubscription(com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotDatafeedSubscription2(a ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequestInterface) *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotDatafeedSubscription", "com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionResult describeSpotDatafeedSubscription()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotDatafeedSubscription() *ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotDatafeedSubscription", "com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult describeSpotFleetInstances(com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotFleetInstances(a ServicesEc2ModelDescribeSpotFleetInstancesRequestInterface) *ServicesEc2ModelDescribeSpotFleetInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotFleetInstances", "com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryResult describeSpotFleetRequestHistory(com.amazonaws.services.ec2.model.DescribeSpotFleetRequestHistoryRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotFleetRequestHistory(a ServicesEc2ModelDescribeSpotFleetRequestHistoryRequestInterface) *ServicesEc2ModelDescribeSpotFleetRequestHistoryResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotFleetRequestHistory", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSpotFleetRequestHistoryRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult describeSpotFleetRequests(com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotFleetRequests2(a ServicesEc2ModelDescribeSpotFleetRequestsRequestInterface) *ServicesEc2ModelDescribeSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotFleetRequests", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult describeSpotFleetRequests()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotFleetRequests() *ServicesEc2ModelDescribeSpotFleetRequestsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotFleetRequests", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult describeSpotInstanceRequests(com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotInstanceRequests2(a ServicesEc2ModelDescribeSpotInstanceRequestsRequestInterface) *ServicesEc2ModelDescribeSpotInstanceRequestsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotInstanceRequests", "com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult describeSpotInstanceRequests()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotInstanceRequests() *ServicesEc2ModelDescribeSpotInstanceRequestsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotInstanceRequests", "com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult describeSpotPriceHistory(com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotPriceHistory2(a ServicesEc2ModelDescribeSpotPriceHistoryRequestInterface) *ServicesEc2ModelDescribeSpotPriceHistoryResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotPriceHistory", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult describeSpotPriceHistory()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSpotPriceHistory() *ServicesEc2ModelDescribeSpotPriceHistoryResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSpotPriceHistory", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSubnetsResult describeSubnets(com.amazonaws.services.ec2.model.DescribeSubnetsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSubnets2(a ServicesEc2ModelDescribeSubnetsRequestInterface) *ServicesEc2ModelDescribeSubnetsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSubnets", "com/amazonaws/services/ec2/model/DescribeSubnetsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeSubnetsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSubnetsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSubnetsResult describeSubnets()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeSubnets() *ServicesEc2ModelDescribeSubnetsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeSubnets", "com/amazonaws/services/ec2/model/DescribeSubnetsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeSubnetsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeTagsResult describeTags(com.amazonaws.services.ec2.model.DescribeTagsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeTags2(a ServicesEc2ModelDescribeTagsRequestInterface) *ServicesEc2ModelDescribeTagsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeTags", "com/amazonaws/services/ec2/model/DescribeTagsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeTagsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeTagsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeTagsResult describeTags()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeTags() *ServicesEc2ModelDescribeTagsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeTags", "com/amazonaws/services/ec2/model/DescribeTagsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeTagsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult describeVolumeAttribute(com.amazonaws.services.ec2.model.DescribeVolumeAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVolumeAttribute(a ServicesEc2ModelDescribeVolumeAttributeRequestInterface) *ServicesEc2ModelDescribeVolumeAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVolumeAttribute", "com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVolumeAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult describeVolumeStatus(com.amazonaws.services.ec2.model.DescribeVolumeStatusRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVolumeStatus2(a ServicesEc2ModelDescribeVolumeStatusRequestInterface) *ServicesEc2ModelDescribeVolumeStatusResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVolumeStatus", "com/amazonaws/services/ec2/model/DescribeVolumeStatusResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVolumeStatusRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult describeVolumeStatus()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVolumeStatus() *ServicesEc2ModelDescribeVolumeStatusResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVolumeStatus", "com/amazonaws/services/ec2/model/DescribeVolumeStatusResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumesResult describeVolumes(com.amazonaws.services.ec2.model.DescribeVolumesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVolumes2(a ServicesEc2ModelDescribeVolumesRequestInterface) *ServicesEc2ModelDescribeVolumesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVolumes", "com/amazonaws/services/ec2/model/DescribeVolumesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVolumesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVolumesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumesResult describeVolumes()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVolumes() *ServicesEc2ModelDescribeVolumesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVolumes", "com/amazonaws/services/ec2/model/DescribeVolumesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVolumesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult describeVpcAttribute(com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcAttribute(a ServicesEc2ModelDescribeVpcAttributeRequestInterface) *ServicesEc2ModelDescribeVpcAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcAttribute", "com/amazonaws/services/ec2/model/DescribeVpcAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult describeVpcClassicLink(com.amazonaws.services.ec2.model.DescribeVpcClassicLinkRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcClassicLink2(a ServicesEc2ModelDescribeVpcClassicLinkRequestInterface) *ServicesEc2ModelDescribeVpcClassicLinkResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcClassicLink", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcClassicLinkRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult describeVpcClassicLink()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcClassicLink() *ServicesEc2ModelDescribeVpcClassicLinkResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcClassicLink", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult describeVpcClassicLinkDnsSupport(com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcClassicLinkDnsSupport(a ServicesEc2ModelDescribeVpcClassicLinkDnsSupportRequestInterface) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcClassicLinkDnsSupport", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult describeVpcEndpointServices(com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcEndpointServices2(a ServicesEc2ModelDescribeVpcEndpointServicesRequestInterface) *ServicesEc2ModelDescribeVpcEndpointServicesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcEndpointServices", "com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult describeVpcEndpointServices()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcEndpointServices() *ServicesEc2ModelDescribeVpcEndpointServicesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcEndpointServices", "com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult describeVpcEndpoints(com.amazonaws.services.ec2.model.DescribeVpcEndpointsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcEndpoints2(a ServicesEc2ModelDescribeVpcEndpointsRequestInterface) *ServicesEc2ModelDescribeVpcEndpointsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcEndpoints", "com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcEndpointsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult describeVpcEndpoints()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcEndpoints() *ServicesEc2ModelDescribeVpcEndpointsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcEndpoints", "com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult describeVpcPeeringConnections(com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcPeeringConnections2(a ServicesEc2ModelDescribeVpcPeeringConnectionsRequestInterface) *ServicesEc2ModelDescribeVpcPeeringConnectionsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcPeeringConnections", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult describeVpcPeeringConnections()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcPeeringConnections() *ServicesEc2ModelDescribeVpcPeeringConnectionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcPeeringConnections", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcsResult describeVpcs(com.amazonaws.services.ec2.model.DescribeVpcsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcs2(a ServicesEc2ModelDescribeVpcsRequestInterface) *ServicesEc2ModelDescribeVpcsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcs", "com/amazonaws/services/ec2/model/DescribeVpcsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpcsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcsResult describeVpcs()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpcs() *ServicesEc2ModelDescribeVpcsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpcs", "com/amazonaws/services/ec2/model/DescribeVpcsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpcsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult describeVpnConnections(com.amazonaws.services.ec2.model.DescribeVpnConnectionsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpnConnections2(a ServicesEc2ModelDescribeVpnConnectionsRequestInterface) *ServicesEc2ModelDescribeVpnConnectionsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpnConnections", "com/amazonaws/services/ec2/model/DescribeVpnConnectionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpnConnectionsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult describeVpnConnections()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpnConnections() *ServicesEc2ModelDescribeVpnConnectionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpnConnections", "com/amazonaws/services/ec2/model/DescribeVpnConnectionsResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult describeVpnGateways(com.amazonaws.services.ec2.model.DescribeVpnGatewaysRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpnGateways2(a ServicesEc2ModelDescribeVpnGatewaysRequestInterface) *ServicesEc2ModelDescribeVpnGatewaysResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpnGateways", "com/amazonaws/services/ec2/model/DescribeVpnGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DescribeVpnGatewaysRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult describeVpnGateways()
func (jbobject *ServicesEc2AmazonEC2Client) DescribeVpnGateways() *ServicesEc2ModelDescribeVpnGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "describeVpnGateways", "com/amazonaws/services/ec2/model/DescribeVpnGatewaysResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DetachClassicLinkVpcResult detachClassicLinkVpc(com.amazonaws.services.ec2.model.DetachClassicLinkVpcRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DetachClassicLinkVpc(a ServicesEc2ModelDetachClassicLinkVpcRequestInterface) *ServicesEc2ModelDetachClassicLinkVpcResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "detachClassicLinkVpc", "com/amazonaws/services/ec2/model/DetachClassicLinkVpcResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DetachClassicLinkVpcRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDetachClassicLinkVpcResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void detachInternetGateway(com.amazonaws.services.ec2.model.DetachInternetGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DetachInternetGateway(a ServicesEc2ModelDetachInternetGatewayRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "detachInternetGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DetachInternetGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void detachNetworkInterface(com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DetachNetworkInterface(a ServicesEc2ModelDetachNetworkInterfaceRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "detachNetworkInterface", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DetachNetworkInterfaceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DetachVolumeResult detachVolume(com.amazonaws.services.ec2.model.DetachVolumeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DetachVolume(a ServicesEc2ModelDetachVolumeRequestInterface) *ServicesEc2ModelDetachVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "detachVolume", "com/amazonaws/services/ec2/model/DetachVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DetachVolumeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDetachVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void detachVpnGateway(com.amazonaws.services.ec2.model.DetachVpnGatewayRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DetachVpnGateway(a ServicesEc2ModelDetachVpnGatewayRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "detachVpnGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DetachVpnGatewayRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void disableVgwRoutePropagation(com.amazonaws.services.ec2.model.DisableVgwRoutePropagationRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DisableVgwRoutePropagation(a ServicesEc2ModelDisableVgwRoutePropagationRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "disableVgwRoutePropagation", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DisableVgwRoutePropagationRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DisableVpcClassicLinkResult disableVpcClassicLink(com.amazonaws.services.ec2.model.DisableVpcClassicLinkRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DisableVpcClassicLink(a ServicesEc2ModelDisableVpcClassicLinkRequestInterface) *ServicesEc2ModelDisableVpcClassicLinkResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "disableVpcClassicLink", "com/amazonaws/services/ec2/model/DisableVpcClassicLinkResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DisableVpcClassicLinkRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDisableVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DisableVpcClassicLinkDnsSupportResult disableVpcClassicLinkDnsSupport(com.amazonaws.services.ec2.model.DisableVpcClassicLinkDnsSupportRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DisableVpcClassicLinkDnsSupport(a ServicesEc2ModelDisableVpcClassicLinkDnsSupportRequestInterface) *ServicesEc2ModelDisableVpcClassicLinkDnsSupportResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "disableVpcClassicLinkDnsSupport", "com/amazonaws/services/ec2/model/DisableVpcClassicLinkDnsSupportResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DisableVpcClassicLinkDnsSupportRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDisableVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void disassociateAddress(com.amazonaws.services.ec2.model.DisassociateAddressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DisassociateAddress(a ServicesEc2ModelDisassociateAddressRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "disassociateAddress", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DisassociateAddressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void disassociateRouteTable(com.amazonaws.services.ec2.model.DisassociateRouteTableRequest)
func (jbobject *ServicesEc2AmazonEC2Client) DisassociateRouteTable(a ServicesEc2ModelDisassociateRouteTableRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "disassociateRouteTable", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DisassociateRouteTableRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void enableVgwRoutePropagation(com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest)
func (jbobject *ServicesEc2AmazonEC2Client) EnableVgwRoutePropagation(a ServicesEc2ModelEnableVgwRoutePropagationRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "enableVgwRoutePropagation", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EnableVgwRoutePropagationRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void enableVolumeIO(com.amazonaws.services.ec2.model.EnableVolumeIORequest)
func (jbobject *ServicesEc2AmazonEC2Client) EnableVolumeIO(a ServicesEc2ModelEnableVolumeIORequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "enableVolumeIO", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EnableVolumeIORequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EnableVpcClassicLinkResult enableVpcClassicLink(com.amazonaws.services.ec2.model.EnableVpcClassicLinkRequest)
func (jbobject *ServicesEc2AmazonEC2Client) EnableVpcClassicLink(a ServicesEc2ModelEnableVpcClassicLinkRequestInterface) *ServicesEc2ModelEnableVpcClassicLinkResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "enableVpcClassicLink", "com/amazonaws/services/ec2/model/EnableVpcClassicLinkResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EnableVpcClassicLinkRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelEnableVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.EnableVpcClassicLinkDnsSupportResult enableVpcClassicLinkDnsSupport(com.amazonaws.services.ec2.model.EnableVpcClassicLinkDnsSupportRequest)
func (jbobject *ServicesEc2AmazonEC2Client) EnableVpcClassicLinkDnsSupport(a ServicesEc2ModelEnableVpcClassicLinkDnsSupportRequestInterface) *ServicesEc2ModelEnableVpcClassicLinkDnsSupportResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "enableVpcClassicLinkDnsSupport", "com/amazonaws/services/ec2/model/EnableVpcClassicLinkDnsSupportResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EnableVpcClassicLinkDnsSupportRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelEnableVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.GetConsoleOutputResult getConsoleOutput(com.amazonaws.services.ec2.model.GetConsoleOutputRequest)
func (jbobject *ServicesEc2AmazonEC2Client) GetConsoleOutput(a ServicesEc2ModelGetConsoleOutputRequestInterface) *ServicesEc2ModelGetConsoleOutputResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConsoleOutput", "com/amazonaws/services/ec2/model/GetConsoleOutputResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GetConsoleOutputRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelGetConsoleOutputResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.GetPasswordDataResult getPasswordData(com.amazonaws.services.ec2.model.GetPasswordDataRequest)
func (jbobject *ServicesEc2AmazonEC2Client) GetPasswordData(a ServicesEc2ModelGetPasswordDataRequestInterface) *ServicesEc2ModelGetPasswordDataResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPasswordData", "com/amazonaws/services/ec2/model/GetPasswordDataResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GetPasswordDataRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelGetPasswordDataResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportImageResult importImage(com.amazonaws.services.ec2.model.ImportImageRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ImportImage2(a ServicesEc2ModelImportImageRequestInterface) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importImage", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportImageRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportImageResult importImage()
func (jbobject *ServicesEc2AmazonEC2Client) ImportImage() *ServicesEc2ModelImportImageResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importImage", "com/amazonaws/services/ec2/model/ImportImageResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceResult importInstance(com.amazonaws.services.ec2.model.ImportInstanceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ImportInstance(a ServicesEc2ModelImportInstanceRequestInterface) *ServicesEc2ModelImportInstanceResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importInstance", "com/amazonaws/services/ec2/model/ImportInstanceResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportInstanceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportInstanceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportKeyPairResult importKeyPair(com.amazonaws.services.ec2.model.ImportKeyPairRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ImportKeyPair(a ServicesEc2ModelImportKeyPairRequestInterface) *ServicesEc2ModelImportKeyPairResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importKeyPair", "com/amazonaws/services/ec2/model/ImportKeyPairResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportKeyPairRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportSnapshotResult importSnapshot(com.amazonaws.services.ec2.model.ImportSnapshotRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ImportSnapshot2(a ServicesEc2ModelImportSnapshotRequestInterface) *ServicesEc2ModelImportSnapshotResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importSnapshot", "com/amazonaws/services/ec2/model/ImportSnapshotResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportSnapshotRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportSnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportSnapshotResult importSnapshot()
func (jbobject *ServicesEc2AmazonEC2Client) ImportSnapshot() *ServicesEc2ModelImportSnapshotResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importSnapshot", "com/amazonaws/services/ec2/model/ImportSnapshotResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportSnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportVolumeResult importVolume(com.amazonaws.services.ec2.model.ImportVolumeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ImportVolume(a ServicesEc2ModelImportVolumeRequestInterface) *ServicesEc2ModelImportVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "importVolume", "com/amazonaws/services/ec2/model/ImportVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportVolumeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelImportVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyHostsResult modifyHosts(com.amazonaws.services.ec2.model.ModifyHostsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyHosts(a ServicesEc2ModelModifyHostsRequestInterface) *ServicesEc2ModelModifyHostsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "modifyHosts", "com/amazonaws/services/ec2/model/ModifyHostsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyHostsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelModifyHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void modifyIdFormat(com.amazonaws.services.ec2.model.ModifyIdFormatRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyIdFormat(a ServicesEc2ModelModifyIdFormatRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyIdFormat", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyIdFormatRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void modifyImageAttribute(com.amazonaws.services.ec2.model.ModifyImageAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyImageAttribute(a ServicesEc2ModelModifyImageAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyImageAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyImageAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void modifyInstanceAttribute(com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyInstanceAttribute(a ServicesEc2ModelModifyInstanceAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyInstanceAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementResult modifyInstancePlacement(com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyInstancePlacement(a ServicesEc2ModelModifyInstancePlacementRequestInterface) *ServicesEc2ModelModifyInstancePlacementResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "modifyInstancePlacement", "com/amazonaws/services/ec2/model/ModifyInstancePlacementResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelModifyInstancePlacementResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void modifyNetworkInterfaceAttribute(com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyNetworkInterfaceAttribute(a ServicesEc2ModelModifyNetworkInterfaceAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyNetworkInterfaceAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult modifyReservedInstances(com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyReservedInstances(a ServicesEc2ModelModifyReservedInstancesRequestInterface) *ServicesEc2ModelModifyReservedInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "modifyReservedInstances", "com/amazonaws/services/ec2/model/ModifyReservedInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelModifyReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void modifySnapshotAttribute(com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifySnapshotAttribute(a ServicesEc2ModelModifySnapshotAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifySnapshotAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifySpotFleetRequestResult modifySpotFleetRequest(com.amazonaws.services.ec2.model.ModifySpotFleetRequestRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifySpotFleetRequest(a ServicesEc2ModelModifySpotFleetRequestRequestInterface) *ServicesEc2ModelModifySpotFleetRequestResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "modifySpotFleetRequest", "com/amazonaws/services/ec2/model/ModifySpotFleetRequestResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifySpotFleetRequestRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelModifySpotFleetRequestResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void modifySubnetAttribute(com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifySubnetAttribute(a ServicesEc2ModelModifySubnetAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifySubnetAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifySubnetAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void modifyVolumeAttribute(com.amazonaws.services.ec2.model.ModifyVolumeAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyVolumeAttribute(a ServicesEc2ModelModifyVolumeAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyVolumeAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyVolumeAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void modifyVpcAttribute(com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyVpcAttribute(a ServicesEc2ModelModifyVpcAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "modifyVpcAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyVpcAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult modifyVpcEndpoint(com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ModifyVpcEndpoint(a ServicesEc2ModelModifyVpcEndpointRequestInterface) *ServicesEc2ModelModifyVpcEndpointResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "modifyVpcEndpoint", "com/amazonaws/services/ec2/model/ModifyVpcEndpointResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelModifyVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.MonitorInstancesResult monitorInstances(com.amazonaws.services.ec2.model.MonitorInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) MonitorInstances(a ServicesEc2ModelMonitorInstancesRequestInterface) *ServicesEc2ModelMonitorInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "monitorInstances", "com/amazonaws/services/ec2/model/MonitorInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/MonitorInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelMonitorInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.MoveAddressToVpcResult moveAddressToVpc(com.amazonaws.services.ec2.model.MoveAddressToVpcRequest)
func (jbobject *ServicesEc2AmazonEC2Client) MoveAddressToVpc(a ServicesEc2ModelMoveAddressToVpcRequestInterface) *ServicesEc2ModelMoveAddressToVpcResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "moveAddressToVpc", "com/amazonaws/services/ec2/model/MoveAddressToVpcResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/MoveAddressToVpcRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelMoveAddressToVpcResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingResult purchaseReservedInstancesOffering(com.amazonaws.services.ec2.model.PurchaseReservedInstancesOfferingRequest)
func (jbobject *ServicesEc2AmazonEC2Client) PurchaseReservedInstancesOffering(a ServicesEc2ModelPurchaseReservedInstancesOfferingRequestInterface) *ServicesEc2ModelPurchaseReservedInstancesOfferingResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "purchaseReservedInstancesOffering", "com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PurchaseReservedInstancesOfferingRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelPurchaseReservedInstancesOfferingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.PurchaseScheduledInstancesResult purchaseScheduledInstances(com.amazonaws.services.ec2.model.PurchaseScheduledInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) PurchaseScheduledInstances(a ServicesEc2ModelPurchaseScheduledInstancesRequestInterface) *ServicesEc2ModelPurchaseScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "purchaseScheduledInstances", "com/amazonaws/services/ec2/model/PurchaseScheduledInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PurchaseScheduledInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelPurchaseScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void rebootInstances(com.amazonaws.services.ec2.model.RebootInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RebootInstances(a ServicesEc2ModelRebootInstancesRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "rebootInstances", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RebootInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RegisterImageResult registerImage(com.amazonaws.services.ec2.model.RegisterImageRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RegisterImage(a ServicesEc2ModelRegisterImageRequestInterface) *ServicesEc2ModelRegisterImageResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "registerImage", "com/amazonaws/services/ec2/model/RegisterImageResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RegisterImageRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRegisterImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionResult rejectVpcPeeringConnection(com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RejectVpcPeeringConnection(a ServicesEc2ModelRejectVpcPeeringConnectionRequestInterface) *ServicesEc2ModelRejectVpcPeeringConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rejectVpcPeeringConnection", "com/amazonaws/services/ec2/model/RejectVpcPeeringConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RejectVpcPeeringConnectionRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRejectVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void releaseAddress(com.amazonaws.services.ec2.model.ReleaseAddressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReleaseAddress(a ServicesEc2ModelReleaseAddressRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseAddress", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReleaseAddressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReleaseHostsResult releaseHosts(com.amazonaws.services.ec2.model.ReleaseHostsRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReleaseHosts(a ServicesEc2ModelReleaseHostsRequestInterface) *ServicesEc2ModelReleaseHostsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "releaseHosts", "com/amazonaws/services/ec2/model/ReleaseHostsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReleaseHostsRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelReleaseHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationResult replaceNetworkAclAssociation(com.amazonaws.services.ec2.model.ReplaceNetworkAclAssociationRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReplaceNetworkAclAssociation(a ServicesEc2ModelReplaceNetworkAclAssociationRequestInterface) *ServicesEc2ModelReplaceNetworkAclAssociationResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "replaceNetworkAclAssociation", "com/amazonaws/services/ec2/model/ReplaceNetworkAclAssociationResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReplaceNetworkAclAssociationRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelReplaceNetworkAclAssociationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void replaceNetworkAclEntry(com.amazonaws.services.ec2.model.ReplaceNetworkAclEntryRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReplaceNetworkAclEntry(a ServicesEc2ModelReplaceNetworkAclEntryRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "replaceNetworkAclEntry", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReplaceNetworkAclEntryRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void replaceRoute(com.amazonaws.services.ec2.model.ReplaceRouteRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReplaceRoute(a ServicesEc2ModelReplaceRouteRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "replaceRoute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReplaceRouteRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationResult replaceRouteTableAssociation(com.amazonaws.services.ec2.model.ReplaceRouteTableAssociationRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReplaceRouteTableAssociation(a ServicesEc2ModelReplaceRouteTableAssociationRequestInterface) *ServicesEc2ModelReplaceRouteTableAssociationResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "replaceRouteTableAssociation", "com/amazonaws/services/ec2/model/ReplaceRouteTableAssociationResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReplaceRouteTableAssociationRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelReplaceRouteTableAssociationResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void reportInstanceStatus(com.amazonaws.services.ec2.model.ReportInstanceStatusRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ReportInstanceStatus(a ServicesEc2ModelReportInstanceStatusRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reportInstanceStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReportInstanceStatusRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RequestSpotFleetResult requestSpotFleet(com.amazonaws.services.ec2.model.RequestSpotFleetRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RequestSpotFleet(a ServicesEc2ModelRequestSpotFleetRequestInterface) *ServicesEc2ModelRequestSpotFleetResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "requestSpotFleet", "com/amazonaws/services/ec2/model/RequestSpotFleetResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RequestSpotFleetRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRequestSpotFleetResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RequestSpotInstancesResult requestSpotInstances(com.amazonaws.services.ec2.model.RequestSpotInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RequestSpotInstances(a ServicesEc2ModelRequestSpotInstancesRequestInterface) *ServicesEc2ModelRequestSpotInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "requestSpotInstances", "com/amazonaws/services/ec2/model/RequestSpotInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RequestSpotInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRequestSpotInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void resetImageAttribute(com.amazonaws.services.ec2.model.ResetImageAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ResetImageAttribute(a ServicesEc2ModelResetImageAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "resetImageAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResetImageAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void resetInstanceAttribute(com.amazonaws.services.ec2.model.ResetInstanceAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ResetInstanceAttribute(a ServicesEc2ModelResetInstanceAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "resetInstanceAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResetInstanceAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void resetNetworkInterfaceAttribute(com.amazonaws.services.ec2.model.ResetNetworkInterfaceAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ResetNetworkInterfaceAttribute(a ServicesEc2ModelResetNetworkInterfaceAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "resetNetworkInterfaceAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResetNetworkInterfaceAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void resetSnapshotAttribute(com.amazonaws.services.ec2.model.ResetSnapshotAttributeRequest)
func (jbobject *ServicesEc2AmazonEC2Client) ResetSnapshotAttribute(a ServicesEc2ModelResetSnapshotAttributeRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "resetSnapshotAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResetSnapshotAttributeRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult restoreAddressToClassic(com.amazonaws.services.ec2.model.RestoreAddressToClassicRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RestoreAddressToClassic(a ServicesEc2ModelRestoreAddressToClassicRequestInterface) *ServicesEc2ModelRestoreAddressToClassicResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "restoreAddressToClassic", "com/amazonaws/services/ec2/model/RestoreAddressToClassicResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RestoreAddressToClassicRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRestoreAddressToClassicResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void revokeSecurityGroupEgress(com.amazonaws.services.ec2.model.RevokeSecurityGroupEgressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RevokeSecurityGroupEgress(a ServicesEc2ModelRevokeSecurityGroupEgressRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "revokeSecurityGroupEgress", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RevokeSecurityGroupEgressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void revokeSecurityGroupIngress(com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RevokeSecurityGroupIngress2(a ServicesEc2ModelRevokeSecurityGroupIngressRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "revokeSecurityGroupIngress", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void revokeSecurityGroupIngress()
func (jbobject *ServicesEc2AmazonEC2Client) RevokeSecurityGroupIngress()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "revokeSecurityGroupIngress", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.services.ec2.model.RunInstancesResult runInstances(com.amazonaws.services.ec2.model.RunInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RunInstances(a ServicesEc2ModelRunInstancesRequestInterface) *ServicesEc2ModelRunInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runInstances", "com/amazonaws/services/ec2/model/RunInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RunInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRunInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult runScheduledInstances(com.amazonaws.services.ec2.model.RunScheduledInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) RunScheduledInstances(a ServicesEc2ModelRunScheduledInstancesRequestInterface) *ServicesEc2ModelRunScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runScheduledInstances", "com/amazonaws/services/ec2/model/RunScheduledInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RunScheduledInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelRunScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.StartInstancesResult startInstances(com.amazonaws.services.ec2.model.StartInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) StartInstances(a ServicesEc2ModelStartInstancesRequestInterface) *ServicesEc2ModelStartInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "startInstances", "com/amazonaws/services/ec2/model/StartInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/StartInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelStartInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.StopInstancesResult stopInstances(com.amazonaws.services.ec2.model.StopInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) StopInstances(a ServicesEc2ModelStopInstancesRequestInterface) *ServicesEc2ModelStopInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stopInstances", "com/amazonaws/services/ec2/model/StopInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/StopInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelStopInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.TerminateInstancesResult terminateInstances(com.amazonaws.services.ec2.model.TerminateInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) TerminateInstances(a ServicesEc2ModelTerminateInstancesRequestInterface) *ServicesEc2ModelTerminateInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "terminateInstances", "com/amazonaws/services/ec2/model/TerminateInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/TerminateInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelTerminateInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void unassignPrivateIpAddresses(com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) UnassignPrivateIpAddresses(a ServicesEc2ModelUnassignPrivateIpAddressesRequestInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unassignPrivateIpAddresses", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/UnassignPrivateIpAddressesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.UnmonitorInstancesResult unmonitorInstances(com.amazonaws.services.ec2.model.UnmonitorInstancesRequest)
func (jbobject *ServicesEc2AmazonEC2Client) UnmonitorInstances(a ServicesEc2ModelUnmonitorInstancesRequestInterface) *ServicesEc2ModelUnmonitorInstancesResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "unmonitorInstances", "com/amazonaws/services/ec2/model/UnmonitorInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UnmonitorInstancesRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelUnmonitorInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.ResponseMetadata getCachedResponseMetadata(com.amazonaws.AmazonWebServiceRequest)
func (jbobject *ServicesEc2AmazonEC2Client) GetCachedResponseMetadata(a AmazonWebServiceRequestInterface) *ResponseMetadata {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCachedResponseMetadata", "com/amazonaws/ResponseMetadata", conv_a.Value().Cast("com/amazonaws/AmazonWebServiceRequest"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ResponseMetadata{}
	unique_x.Callable = dst
	return unique_x
}


