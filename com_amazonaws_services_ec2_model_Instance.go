package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelInstanceInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.Instance withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelInstance

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.Instance withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelInstance

	// public void setState(com.amazonaws.services.ec2.model.InstanceState)
	SetState(a ServicesEc2ModelInstanceStateInterface) 

	// public com.amazonaws.services.ec2.model.InstanceState getState()
	GetState() *ServicesEc2ModelInstanceState

	// public com.amazonaws.services.ec2.model.Instance withState(com.amazonaws.services.ec2.model.InstanceState)
	WithState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstance

	// public void setPrivateDnsName(java.lang.String)
	SetPrivateDnsName(a string) 

	// public java.lang.String getPrivateDnsName()
	GetPrivateDnsName() string

	// public com.amazonaws.services.ec2.model.Instance withPrivateDnsName(java.lang.String)
	WithPrivateDnsName(a string) *ServicesEc2ModelInstance

	// public void setPublicDnsName(java.lang.String)
	SetPublicDnsName(a string) 

	// public java.lang.String getPublicDnsName()
	GetPublicDnsName() string

	// public com.amazonaws.services.ec2.model.Instance withPublicDnsName(java.lang.String)
	WithPublicDnsName(a string) *ServicesEc2ModelInstance

	// public void setStateTransitionReason(java.lang.String)
	SetStateTransitionReason(a string) 

	// public java.lang.String getStateTransitionReason()
	GetStateTransitionReason() string

	// public com.amazonaws.services.ec2.model.Instance withStateTransitionReason(java.lang.String)
	WithStateTransitionReason(a string) *ServicesEc2ModelInstance

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.Instance withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelInstance

	// public void setAmiLaunchIndex(java.lang.Integer)
	SetAmiLaunchIndex(a int) 

	// public java.lang.Integer getAmiLaunchIndex()
	GetAmiLaunchIndex() int

	// public com.amazonaws.services.ec2.model.Instance withAmiLaunchIndex(java.lang.Integer)
	WithAmiLaunchIndex(a int) *ServicesEc2ModelInstance

	// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
	GetProductCodes() []*ServicesEc2ModelProductCode

	// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	SetProductCodes(a []*ServicesEc2ModelProductCode) 

	// public com.amazonaws.services.ec2.model.Instance withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
	WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelInstance

	// public com.amazonaws.services.ec2.model.Instance withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelInstance

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.Instance withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelInstance

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.Instance withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelInstance

	// public void setLaunchTime(java.util.Date)
	SetLaunchTime(a time.Time) 

	// public java.util.Date getLaunchTime()
	GetLaunchTime() time.Time

	// public com.amazonaws.services.ec2.model.Instance withLaunchTime(java.util.Date)
	WithLaunchTime(a time.Time) *ServicesEc2ModelInstance

	// public void setPlacement(com.amazonaws.services.ec2.model.Placement)
	SetPlacement(a ServicesEc2ModelPlacementInterface) 

	// public com.amazonaws.services.ec2.model.Placement getPlacement()
	GetPlacement() *ServicesEc2ModelPlacement

	// public com.amazonaws.services.ec2.model.Instance withPlacement(com.amazonaws.services.ec2.model.Placement)
	WithPlacement(a ServicesEc2ModelPlacementInterface) *ServicesEc2ModelInstance

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.Instance withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelInstance

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.Instance withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelInstance

	// public void setPlatform(java.lang.String)
	SetPlatform2(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.Instance withPlatform(java.lang.String)
	WithPlatform2(a string) *ServicesEc2ModelInstance

	// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	SetPlatform(a ServicesEc2ModelPlatformValuesInterface) 

	// public com.amazonaws.services.ec2.model.Instance withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelInstance

	// public void setMonitoring(com.amazonaws.services.ec2.model.Monitoring)
	SetMonitoring(a ServicesEc2ModelMonitoringInterface) 

	// public com.amazonaws.services.ec2.model.Monitoring getMonitoring()
	GetMonitoring() *ServicesEc2ModelMonitoring

	// public com.amazonaws.services.ec2.model.Instance withMonitoring(com.amazonaws.services.ec2.model.Monitoring)
	WithMonitoring(a ServicesEc2ModelMonitoringInterface) *ServicesEc2ModelInstance

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.Instance withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelInstance

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.Instance withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelInstance

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.Instance withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelInstance

	// public void setPublicIpAddress(java.lang.String)
	SetPublicIpAddress(a string) 

	// public java.lang.String getPublicIpAddress()
	GetPublicIpAddress() string

	// public com.amazonaws.services.ec2.model.Instance withPublicIpAddress(java.lang.String)
	WithPublicIpAddress(a string) *ServicesEc2ModelInstance

	// public void setStateReason(com.amazonaws.services.ec2.model.StateReason)
	SetStateReason(a ServicesEc2ModelStateReasonInterface) 

	// public com.amazonaws.services.ec2.model.StateReason getStateReason()
	GetStateReason() *ServicesEc2ModelStateReason

	// public com.amazonaws.services.ec2.model.Instance withStateReason(com.amazonaws.services.ec2.model.StateReason)
	WithStateReason(a ServicesEc2ModelStateReasonInterface) *ServicesEc2ModelInstance

	// public void setArchitecture(java.lang.String)
	SetArchitecture2(a string) 

	// public java.lang.String getArchitecture()
	GetArchitecture() string

	// public com.amazonaws.services.ec2.model.Instance withArchitecture(java.lang.String)
	WithArchitecture2(a string) *ServicesEc2ModelInstance

	// public void setArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
	SetArchitecture(a ServicesEc2ModelArchitectureValuesInterface) 

	// public com.amazonaws.services.ec2.model.Instance withArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
	WithArchitecture(a ServicesEc2ModelArchitectureValuesInterface) *ServicesEc2ModelInstance

	// public void setRootDeviceType(java.lang.String)
	SetRootDeviceType2(a string) 

	// public java.lang.String getRootDeviceType()
	GetRootDeviceType() string

	// public com.amazonaws.services.ec2.model.Instance withRootDeviceType(java.lang.String)
	WithRootDeviceType2(a string) *ServicesEc2ModelInstance

	// public void setRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
	SetRootDeviceType(a ServicesEc2ModelDeviceTypeInterface) 

	// public com.amazonaws.services.ec2.model.Instance withRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
	WithRootDeviceType(a ServicesEc2ModelDeviceTypeInterface) *ServicesEc2ModelInstance

	// public void setRootDeviceName(java.lang.String)
	SetRootDeviceName(a string) 

	// public java.lang.String getRootDeviceName()
	GetRootDeviceName() string

	// public com.amazonaws.services.ec2.model.Instance withRootDeviceName(java.lang.String)
	WithRootDeviceName(a string) *ServicesEc2ModelInstance

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelInstanceBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelInstanceBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.Instance withBlockDeviceMappings(com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstance

	// public com.amazonaws.services.ec2.model.Instance withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstance

	// public void setVirtualizationType(java.lang.String)
	SetVirtualizationType2(a string) 

	// public java.lang.String getVirtualizationType()
	GetVirtualizationType() string

	// public com.amazonaws.services.ec2.model.Instance withVirtualizationType(java.lang.String)
	WithVirtualizationType2(a string) *ServicesEc2ModelInstance

	// public void setVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
	SetVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface) 

	// public com.amazonaws.services.ec2.model.Instance withVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
	WithVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface) *ServicesEc2ModelInstance

	// public void setInstanceLifecycle(java.lang.String)
	SetInstanceLifecycle2(a string) 

	// public java.lang.String getInstanceLifecycle()
	GetInstanceLifecycle() string

	// public com.amazonaws.services.ec2.model.Instance withInstanceLifecycle(java.lang.String)
	WithInstanceLifecycle2(a string) *ServicesEc2ModelInstance

	// public void setInstanceLifecycle(com.amazonaws.services.ec2.model.InstanceLifecycleType)
	SetInstanceLifecycle(a ServicesEc2ModelInstanceLifecycleTypeInterface) 

	// public com.amazonaws.services.ec2.model.Instance withInstanceLifecycle(com.amazonaws.services.ec2.model.InstanceLifecycleType)
	WithInstanceLifecycle(a ServicesEc2ModelInstanceLifecycleTypeInterface) *ServicesEc2ModelInstance

	// public void setSpotInstanceRequestId(java.lang.String)
	SetSpotInstanceRequestId(a string) 

	// public java.lang.String getSpotInstanceRequestId()
	GetSpotInstanceRequestId() string

	// public com.amazonaws.services.ec2.model.Instance withSpotInstanceRequestId(java.lang.String)
	WithSpotInstanceRequestId(a string) *ServicesEc2ModelInstance

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.Instance withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelInstance

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.Instance withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelInstance

	// public com.amazonaws.services.ec2.model.Instance withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelInstance

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getSecurityGroups()
	GetSecurityGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetSecurityGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.Instance withSecurityGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithSecurityGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstance

	// public com.amazonaws.services.ec2.model.Instance withSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithSecurityGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstance

	// public void setSourceDestCheck(java.lang.Boolean)
	SetSourceDestCheck(a bool) 

	// public java.lang.Boolean getSourceDestCheck()
	GetSourceDestCheck() bool

	// public com.amazonaws.services.ec2.model.Instance withSourceDestCheck(java.lang.Boolean)
	WithSourceDestCheck(a bool) *ServicesEc2ModelInstance

	// public java.lang.Boolean isSourceDestCheck()
	IsSourceDestCheck() bool

	// public void setHypervisor(java.lang.String)
	SetHypervisor2(a string) 

	// public java.lang.String getHypervisor()
	GetHypervisor() string

	// public com.amazonaws.services.ec2.model.Instance withHypervisor(java.lang.String)
	WithHypervisor2(a string) *ServicesEc2ModelInstance

	// public void setHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
	SetHypervisor(a ServicesEc2ModelHypervisorTypeInterface) 

	// public com.amazonaws.services.ec2.model.Instance withHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
	WithHypervisor(a ServicesEc2ModelHypervisorTypeInterface) *ServicesEc2ModelInstance

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterface> getNetworkInterfaces()
	GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterface

	// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterface>)
	SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterface) 

	// public com.amazonaws.services.ec2.model.Instance withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterface...)
	WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterface) *ServicesEc2ModelInstance

	// public com.amazonaws.services.ec2.model.Instance withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterface>)
	WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterface) *ServicesEc2ModelInstance

	// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfile)
	SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileInterface) 

	// public com.amazonaws.services.ec2.model.IamInstanceProfile getIamInstanceProfile()
	GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfile

	// public com.amazonaws.services.ec2.model.Instance withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfile)
	WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileInterface) *ServicesEc2ModelInstance

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.Instance withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelInstance

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public void setSriovNetSupport(java.lang.String)
	SetSriovNetSupport(a string) 

	// public java.lang.String getSriovNetSupport()
	GetSriovNetSupport() string

	// public com.amazonaws.services.ec2.model.Instance withSriovNetSupport(java.lang.String)
	WithSriovNetSupport(a string) *ServicesEc2ModelInstance

	// public com.amazonaws.services.ec2.model.Instance clone()
	Clone() *ServicesEc2ModelInstance
}

type ServicesEc2ModelInstance struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Instance()
func NewServicesEc2ModelInstance() (*ServicesEc2ModelInstance) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Instance")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstance{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelInstance) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithInstanceId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetImageId(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetImageId() string {
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

// public com.amazonaws.services.ec2.model.Instance withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithImageId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstance) SetState(a ServicesEc2ModelInstanceStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceState getState()
func (jbobject *ServicesEc2ModelInstance) GetState() *ServicesEc2ModelInstanceState {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "com/amazonaws/services/ec2/model/InstanceState")
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstance) WithState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetPrivateDnsName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateDnsName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrivateDnsName()
func (jbobject *ServicesEc2ModelInstance) GetPrivateDnsName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateDnsName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithPrivateDnsName(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateDnsName", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublicDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetPublicDnsName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicDnsName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicDnsName()
func (jbobject *ServicesEc2ModelInstance) GetPublicDnsName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicDnsName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withPublicDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithPublicDnsName(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicDnsName", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStateTransitionReason(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetStateTransitionReason(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStateTransitionReason", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStateTransitionReason()
func (jbobject *ServicesEc2ModelInstance) GetStateTransitionReason() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStateTransitionReason", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withStateTransitionReason(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithStateTransitionReason(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStateTransitionReason", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetKeyName(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetKeyName() string {
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

// public com.amazonaws.services.ec2.model.Instance withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithKeyName(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAmiLaunchIndex(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstance) SetAmiLaunchIndex(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAmiLaunchIndex", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getAmiLaunchIndex()
func (jbobject *ServicesEc2ModelInstance) GetAmiLaunchIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAmiLaunchIndex", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.Instance withAmiLaunchIndex(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstance) WithAmiLaunchIndex(a int) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAmiLaunchIndex", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
func (jbobject *ServicesEc2ModelInstance) GetProductCodes() []*ServicesEc2ModelProductCode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductCodes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelProductCode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelInstance) SetProductCodes(a []*ServicesEc2ModelProductCode)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
func (jbobject *ServicesEc2ModelInstance) WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ProductCode")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCode"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelInstance) WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetInstanceType2(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.Instance withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithInstanceType2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelInstance) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
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

// public com.amazonaws.services.ec2.model.Instance withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelInstance) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchTime(java.util.Date)
func (jbobject *ServicesEc2ModelInstance) SetLaunchTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getLaunchTime()
func (jbobject *ServicesEc2ModelInstance) GetLaunchTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchTime", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.Instance withLaunchTime(java.util.Date)
func (jbobject *ServicesEc2ModelInstance) WithLaunchTime(a time.Time) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchTime", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlacement(com.amazonaws.services.ec2.model.Placement)
func (jbobject *ServicesEc2ModelInstance) SetPlacement(a ServicesEc2ModelPlacementInterface)  {
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
func (jbobject *ServicesEc2ModelInstance) GetPlacement() *ServicesEc2ModelPlacement {
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

// public com.amazonaws.services.ec2.model.Instance withPlacement(com.amazonaws.services.ec2.model.Placement)
func (jbobject *ServicesEc2ModelInstance) WithPlacement(a ServicesEc2ModelPlacementInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacement", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Placement"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetKernelId(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetKernelId() string {
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

// public com.amazonaws.services.ec2.model.Instance withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithKernelId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetRamdiskId(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetRamdiskId() string {
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

// public com.amazonaws.services.ec2.model.Instance withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithRamdiskId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetPlatform2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPlatform()
func (jbobject *ServicesEc2ModelInstance) GetPlatform() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlatform", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithPlatform2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelInstance) SetPlatform(a ServicesEc2ModelPlatformValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelInstance) WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoring(com.amazonaws.services.ec2.model.Monitoring)
func (jbobject *ServicesEc2ModelInstance) SetMonitoring(a ServicesEc2ModelMonitoringInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoring", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Monitoring"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Monitoring getMonitoring()
func (jbobject *ServicesEc2ModelInstance) GetMonitoring() *ServicesEc2ModelMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoring", "com/amazonaws/services/ec2/model/Monitoring")
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
	unique_x := &ServicesEc2ModelMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withMonitoring(com.amazonaws.services.ec2.model.Monitoring)
func (jbobject *ServicesEc2ModelInstance) WithMonitoring(a ServicesEc2ModelMonitoringInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoring", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Monitoring"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.Instance withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithSubnetId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelInstance) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithVpcId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.Instance withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithPrivateIpAddress(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublicIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetPublicIpAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIpAddress()
func (jbobject *ServicesEc2ModelInstance) GetPublicIpAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIpAddress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withPublicIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithPublicIpAddress(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIpAddress", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStateReason(com.amazonaws.services.ec2.model.StateReason)
func (jbobject *ServicesEc2ModelInstance) SetStateReason(a ServicesEc2ModelStateReasonInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStateReason", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/StateReason"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.StateReason getStateReason()
func (jbobject *ServicesEc2ModelInstance) GetStateReason() *ServicesEc2ModelStateReason {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStateReason", "com/amazonaws/services/ec2/model/StateReason")
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
	unique_x := &ServicesEc2ModelStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withStateReason(com.amazonaws.services.ec2.model.StateReason)
func (jbobject *ServicesEc2ModelInstance) WithStateReason(a ServicesEc2ModelStateReasonInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStateReason", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/StateReason"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetArchitecture2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArchitecture", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getArchitecture()
func (jbobject *ServicesEc2ModelInstance) GetArchitecture() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArchitecture", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithArchitecture2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
func (jbobject *ServicesEc2ModelInstance) SetArchitecture(a ServicesEc2ModelArchitectureValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArchitecture", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ArchitectureValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
func (jbobject *ServicesEc2ModelInstance) WithArchitecture(a ServicesEc2ModelArchitectureValuesInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ArchitectureValues"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetRootDeviceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRootDeviceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRootDeviceType()
func (jbobject *ServicesEc2ModelInstance) GetRootDeviceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRootDeviceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withRootDeviceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithRootDeviceType2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceType", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
func (jbobject *ServicesEc2ModelInstance) SetRootDeviceType(a ServicesEc2ModelDeviceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRootDeviceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeviceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
func (jbobject *ServicesEc2ModelInstance) WithRootDeviceType(a ServicesEc2ModelDeviceTypeInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceType", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeviceType"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetRootDeviceName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRootDeviceName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRootDeviceName()
func (jbobject *ServicesEc2ModelInstance) GetRootDeviceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRootDeviceName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withRootDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithRootDeviceName(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceName", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelInstance) GetBlockDeviceMappings() []*ServicesEc2ModelInstanceBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockDeviceMappings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceBlockDeviceMapping)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
func (jbobject *ServicesEc2ModelInstance) SetBlockDeviceMappings(a []*ServicesEc2ModelInstanceBlockDeviceMapping)  {
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

// public com.amazonaws.services.ec2.model.Instance withBlockDeviceMappings(com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping...)
func (jbobject *ServicesEc2ModelInstance) WithBlockDeviceMappings(a ...*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
func (jbobject *ServicesEc2ModelInstance) WithBlockDeviceMappings2(a []*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVirtualizationType(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetVirtualizationType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVirtualizationType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVirtualizationType()
func (jbobject *ServicesEc2ModelInstance) GetVirtualizationType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVirtualizationType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withVirtualizationType(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithVirtualizationType2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualizationType", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
func (jbobject *ServicesEc2ModelInstance) SetVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVirtualizationType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VirtualizationType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
func (jbobject *ServicesEc2ModelInstance) WithVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualizationType", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VirtualizationType"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceLifecycle(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetInstanceLifecycle2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceLifecycle", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceLifecycle()
func (jbobject *ServicesEc2ModelInstance) GetInstanceLifecycle() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceLifecycle", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withInstanceLifecycle(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithInstanceLifecycle2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceLifecycle", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceLifecycle(com.amazonaws.services.ec2.model.InstanceLifecycleType)
func (jbobject *ServicesEc2ModelInstance) SetInstanceLifecycle(a ServicesEc2ModelInstanceLifecycleTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceLifecycle", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceLifecycleType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withInstanceLifecycle(com.amazonaws.services.ec2.model.InstanceLifecycleType)
func (jbobject *ServicesEc2ModelInstance) WithInstanceLifecycle(a ServicesEc2ModelInstanceLifecycleTypeInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceLifecycle", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceLifecycleType"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotInstanceRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetSpotInstanceRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotInstanceRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotInstanceRequestId()
func (jbobject *ServicesEc2ModelInstance) GetSpotInstanceRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotInstanceRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withSpotInstanceRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithSpotInstanceRequestId(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequestId", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelInstance) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.Instance withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithClientToken(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelInstance) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelInstance) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelInstance) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelInstance) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getSecurityGroups()
func (jbobject *ServicesEc2ModelInstance) GetSecurityGroups() []*ServicesEc2ModelGroupIdentifier {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecurityGroups", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelGroupIdentifier)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelInstance) SetSecurityGroups(a []*ServicesEc2ModelGroupIdentifier)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSecurityGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withSecurityGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelInstance) WithSecurityGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelInstance) WithSecurityGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstance) SetSourceDestCheck(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceDestCheck", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getSourceDestCheck()
func (jbobject *ServicesEc2ModelInstance) GetSourceDestCheck() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceDestCheck", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.Instance withSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstance) WithSourceDestCheck(a bool) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceDestCheck", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isSourceDestCheck()
func (jbobject *ServicesEc2ModelInstance) IsSourceDestCheck() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isSourceDestCheck", "java/lang/Boolean")
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

// public void setHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetHypervisor2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHypervisor", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHypervisor()
func (jbobject *ServicesEc2ModelInstance) GetHypervisor() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHypervisor", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithHypervisor2(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHypervisor", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
func (jbobject *ServicesEc2ModelInstance) SetHypervisor(a ServicesEc2ModelHypervisorTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHypervisor", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/HypervisorType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Instance withHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
func (jbobject *ServicesEc2ModelInstance) WithHypervisor(a ServicesEc2ModelHypervisorTypeInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHypervisor", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/HypervisorType"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterface> getNetworkInterfaces()
func (jbobject *ServicesEc2ModelInstance) GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaces", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceNetworkInterface)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterface>)
func (jbobject *ServicesEc2ModelInstance) SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterface)  {
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

// public com.amazonaws.services.ec2.model.Instance withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterface...)
func (jbobject *ServicesEc2ModelInstance) WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceNetworkInterface")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterface"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterface>)
func (jbobject *ServicesEc2ModelInstance) WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfile)
func (jbobject *ServicesEc2ModelInstance) SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIamInstanceProfile", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/IamInstanceProfile"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.IamInstanceProfile getIamInstanceProfile()
func (jbobject *ServicesEc2ModelInstance) GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIamInstanceProfile", "com/amazonaws/services/ec2/model/IamInstanceProfile")
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
	unique_x := &ServicesEc2ModelIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Instance withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfile)
func (jbobject *ServicesEc2ModelInstance) WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileInterface) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIamInstanceProfile", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IamInstanceProfile"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstance) SetEbsOptimized(a bool)  {
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
func (jbobject *ServicesEc2ModelInstance) GetEbsOptimized() bool {
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

// public com.amazonaws.services.ec2.model.Instance withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstance) WithEbsOptimized(a bool) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelInstance) IsEbsOptimized() bool {
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

// public void setSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) SetSriovNetSupport(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSriovNetSupport", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSriovNetSupport()
func (jbobject *ServicesEc2ModelInstance) GetSriovNetSupport() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSriovNetSupport", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Instance withSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelInstance) WithSriovNetSupport(a string) *ServicesEc2ModelInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSriovNetSupport", "com/amazonaws/services/ec2/model/Instance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstance) ToString() string {
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
func (jbobject *ServicesEc2ModelInstance) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstance) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Instance clone()
func (jbobject *ServicesEc2ModelInstance) Clone() *ServicesEc2ModelInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Instance")
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
	unique_x := &ServicesEc2ModelInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstance) Clone2() (*JavaLangObject, error) {
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


