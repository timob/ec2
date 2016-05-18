package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRunInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setMinCount(java.lang.Integer)
	SetMinCount(a int) 

	// public java.lang.Integer getMinCount()
	GetMinCount() int

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withMinCount(java.lang.Integer)
	WithMinCount(a int) *ServicesEc2ModelRunInstancesRequest

	// public void setMaxCount(java.lang.Integer)
	SetMaxCount(a int) 

	// public java.lang.Integer getMaxCount()
	GetMaxCount() int

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withMaxCount(java.lang.Integer)
	WithMaxCount(a int) *ServicesEc2ModelRunInstancesRequest

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelRunInstancesRequest

	// public java.util.List<java.lang.String> getSecurityGroups()
	GetSecurityGroups() []string

	// public void setSecurityGroups(java.util.Collection<java.lang.String>)
	SetSecurityGroups(a []string) 

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroups(java.lang.String...)
	WithSecurityGroups(a ...string) *ServicesEc2ModelRunInstancesRequest

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroups(java.util.Collection<java.lang.String>)
	WithSecurityGroups2(a []string) *ServicesEc2ModelRunInstancesRequest

	// public java.util.List<java.lang.String> getSecurityGroupIds()
	GetSecurityGroupIds() []string

	// public void setSecurityGroupIds(java.util.Collection<java.lang.String>)
	SetSecurityGroupIds(a []string) 

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroupIds(java.lang.String...)
	WithSecurityGroupIds(a ...string) *ServicesEc2ModelRunInstancesRequest

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroupIds(java.util.Collection<java.lang.String>)
	WithSecurityGroupIds2(a []string) *ServicesEc2ModelRunInstancesRequest

	// public void setUserData(java.lang.String)
	SetUserData(a string) 

	// public java.lang.String getUserData()
	GetUserData() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withUserData(java.lang.String)
	WithUserData(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelRunInstancesRequest

	// public void setPlacement(com.amazonaws.services.ec2.model.Placement)
	SetPlacement(a ServicesEc2ModelPlacementInterface) 

	// public com.amazonaws.services.ec2.model.Placement getPlacement()
	GetPlacement() *ServicesEc2ModelPlacement

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withPlacement(com.amazonaws.services.ec2.model.Placement)
	WithPlacement(a ServicesEc2ModelPlacementInterface) *ServicesEc2ModelRunInstancesRequest

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelRunInstancesRequest

	// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelRunInstancesRequest

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelRunInstancesRequest

	// public void setMonitoring(java.lang.Boolean)
	SetMonitoring(a bool) 

	// public java.lang.Boolean getMonitoring()
	GetMonitoring() bool

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withMonitoring(java.lang.Boolean)
	WithMonitoring(a bool) *ServicesEc2ModelRunInstancesRequest

	// public java.lang.Boolean isMonitoring()
	IsMonitoring() bool

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setDisableApiTermination(java.lang.Boolean)
	SetDisableApiTermination(a bool) 

	// public java.lang.Boolean getDisableApiTermination()
	GetDisableApiTermination() bool

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withDisableApiTermination(java.lang.Boolean)
	WithDisableApiTermination(a bool) *ServicesEc2ModelRunInstancesRequest

	// public java.lang.Boolean isDisableApiTermination()
	IsDisableApiTermination() bool

	// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
	SetInstanceInitiatedShutdownBehavior2(a string) 

	// public java.lang.String getInstanceInitiatedShutdownBehavior()
	GetInstanceInitiatedShutdownBehavior() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceInitiatedShutdownBehavior(java.lang.String)
	WithInstanceInitiatedShutdownBehavior2(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
	SetInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface) 

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
	WithInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface) *ServicesEc2ModelRunInstancesRequest

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelRunInstancesRequest

	// public void setAdditionalInfo(java.lang.String)
	SetAdditionalInfo(a string) 

	// public java.lang.String getAdditionalInfo()
	GetAdditionalInfo() string

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withAdditionalInfo(java.lang.String)
	WithAdditionalInfo(a string) *ServicesEc2ModelRunInstancesRequest

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification> getNetworkInterfaces()
	GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterfaceSpecification

	// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
	SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) 

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification...)
	WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelRunInstancesRequest

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
	WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelRunInstancesRequest

	// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
	SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification getIamInstanceProfile()
	GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfileSpecification

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
	WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) *ServicesEc2ModelRunInstancesRequest

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.RunInstancesRequest withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelRunInstancesRequest

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RunInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.RunInstancesRequest clone()
	Clone3() *ServicesEc2ModelRunInstancesRequest
}

type ServicesEc2ModelRunInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest()
func NewServicesEc2ModelRunInstancesRequest() (*ServicesEc2ModelRunInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RunInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRunInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest(java.lang.String, java.lang.Integer, java.lang.Integer)
func NewServicesEc2ModelRunInstancesRequest2(a string, b int, c int) (*ServicesEc2ModelRunInstancesRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaInteger()
	conv_c := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Integer"), conv_c.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2ModelRunInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetImageId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImageId()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetImageId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageId", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithImageId(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMinCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetMinCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMinCount()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetMinCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinCount", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withMinCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithMinCount(a int) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMinCount", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetMaxCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMaxCount()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetMaxCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxCount", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withMaxCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithMaxCount(a int) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxCount", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetKeyName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKeyName()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetKeyName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyName", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithKeyName(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getSecurityGroups()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetSecurityGroups() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecurityGroups", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSecurityGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetSecurityGroups(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSecurityGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithSecurityGroups(a ...string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithSecurityGroups2(a []string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getSecurityGroupIds()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetSecurityGroupIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecurityGroupIds", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSecurityGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetSecurityGroupIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSecurityGroupIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroupIds(java.lang.String...)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithSecurityGroupIds(a ...string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroupIds", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withSecurityGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithSecurityGroupIds2(a []string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroupIds", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetUserData(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserData", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUserData()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetUserData() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserData", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withUserData(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithUserData(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetInstanceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithInstanceType2(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlacement(com.amazonaws.services.ec2.model.Placement)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetPlacement(a ServicesEc2ModelPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Placement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Placement getPlacement()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetPlacement() *ServicesEc2ModelPlacement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlacement", "com/amazonaws/services/ec2/model/Placement")
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withPlacement(com.amazonaws.services.ec2.model.Placement)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithPlacement(a ServicesEc2ModelPlacementInterface) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacement", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Placement"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetKernelId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKernelId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKernelId()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetKernelId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKernelId", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithKernelId(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetRamdiskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRamdiskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRamdiskId()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetRamdiskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRamdiskId", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithRamdiskId(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockDeviceMappings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelBlockDeviceMapping)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlockDeviceMappings", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/BlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BlockDeviceMapping"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoring(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetMonitoring(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoring", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getMonitoring()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetMonitoring() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoring", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withMonitoring(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithMonitoring(a bool) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoring", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isMonitoring()
func (jbobject *ServicesEc2ModelRunInstancesRequest) IsMonitoring() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMonitoring", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetSubnetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnetId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSubnetId()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetSubnetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnetId", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithSubnetId(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDisableApiTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetDisableApiTermination(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDisableApiTermination", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getDisableApiTermination()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetDisableApiTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisableApiTermination", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withDisableApiTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithDisableApiTermination(a bool) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDisableApiTermination", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDisableApiTermination()
func (jbobject *ServicesEc2ModelRunInstancesRequest) IsDisableApiTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisableApiTermination", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetInstanceInitiatedShutdownBehavior2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceInitiatedShutdownBehavior", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceInitiatedShutdownBehavior()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetInstanceInitiatedShutdownBehavior() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceInitiatedShutdownBehavior", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithInstanceInitiatedShutdownBehavior2(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceInitiatedShutdownBehavior", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ShutdownBehavior"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ShutdownBehavior"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetPrivateIpAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrivateIpAddress()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetPrivateIpAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddress", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithPrivateIpAddress(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithClientToken(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAdditionalInfo(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetAdditionalInfo(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAdditionalInfo", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAdditionalInfo()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetAdditionalInfo() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdditionalInfo", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withAdditionalInfo(java.lang.String)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithAdditionalInfo(a string) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdditionalInfo", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification> getNetworkInterfaces()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterfaceSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaces", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceNetworkInterfaceSpecification)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterfaces", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification...)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceSpecification"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIamInstanceProfile", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/IamInstanceProfileSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification getIamInstanceProfile()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfileSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIamInstanceProfile", "com/amazonaws/services/ec2/model/IamInstanceProfileSpecification")
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
	unique_x := &ServicesEc2ModelIamInstanceProfileSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIamInstanceProfile", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IamInstanceProfileSpecification"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRunInstancesRequest) SetEbsOptimized(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEbsOptimized", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEbsOptimized()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetEbsOptimized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEbsOptimized", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelRunInstancesRequest) WithEbsOptimized(a bool) *ServicesEc2ModelRunInstancesRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/RunInstancesRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelRunInstancesRequest) IsEbsOptimized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEbsOptimized", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RunInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelRunInstancesRequest) GetDryRunRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDryRunRequest", "com/amazonaws/Request")
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
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRunInstancesRequest) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelRunInstancesRequest) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelRunInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RunInstancesRequest clone()
func (jbobject *ServicesEc2ModelRunInstancesRequest) Clone3() *ServicesEc2ModelRunInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RunInstancesRequest")
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
	unique_x := &ServicesEc2ModelRunInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelRunInstancesRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRunInstancesRequest) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


