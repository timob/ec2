package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSpotFleetLaunchSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getSecurityGroups()
	GetSecurityGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetSecurityGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSecurityGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithSecurityGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithSecurityGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setUserData(java.lang.String)
	SetUserData(a string) 

	// public java.lang.String getUserData()
	GetUserData() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withUserData(java.lang.String)
	WithUserData(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setAddressingType(java.lang.String)
	SetAddressingType(a string) 

	// public java.lang.String getAddressingType()
	GetAddressingType() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withAddressingType(java.lang.String)
	WithAddressingType(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
	SetPlacement(a ServicesEc2ModelSpotPlacementInterface) 

	// public com.amazonaws.services.ec2.model.SpotPlacement getPlacement()
	GetPlacement() *ServicesEc2ModelSpotPlacement

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
	WithPlacement(a ServicesEc2ModelSpotPlacementInterface) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setMonitoring(com.amazonaws.services.ec2.model.SpotFleetMonitoring)
	SetMonitoring(a ServicesEc2ModelSpotFleetMonitoringInterface) 

	// public com.amazonaws.services.ec2.model.SpotFleetMonitoring getMonitoring()
	GetMonitoring() *ServicesEc2ModelSpotFleetMonitoring

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withMonitoring(com.amazonaws.services.ec2.model.SpotFleetMonitoring)
	WithMonitoring(a ServicesEc2ModelSpotFleetMonitoringInterface) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification> getNetworkInterfaces()
	GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterfaceSpecification

	// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
	SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) 

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification...)
	WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
	WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
	SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification getIamInstanceProfile()
	GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfileSpecification

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
	WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public void setWeightedCapacity(java.lang.Double)
	SetWeightedCapacity(a float64) 

	// public java.lang.Double getWeightedCapacity()
	GetWeightedCapacity() float64

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withWeightedCapacity(java.lang.Double)
	WithWeightedCapacity(a float64) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public void setSpotPrice(java.lang.String)
	SetSpotPrice(a string) 

	// public java.lang.String getSpotPrice()
	GetSpotPrice() string

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSpotPrice(java.lang.String)
	WithSpotPrice(a string) *ServicesEc2ModelSpotFleetLaunchSpecification

	// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification clone()
	Clone() *ServicesEc2ModelSpotFleetLaunchSpecification
}

type ServicesEc2ModelSpotFleetLaunchSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification()
func NewServicesEc2ModelSpotFleetLaunchSpecification() (*ServicesEc2ModelSpotFleetLaunchSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetImageId(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetImageId() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithImageId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetKeyName(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetKeyName() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithKeyName(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getSecurityGroups()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetSecurityGroups() []*ServicesEc2ModelGroupIdentifier {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetSecurityGroups(a []*ServicesEc2ModelGroupIdentifier)  {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSecurityGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithSecurityGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithSecurityGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetUserData(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetUserData() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withUserData(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithUserData(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAddressingType(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetAddressingType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAddressingType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAddressingType()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetAddressingType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAddressingType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withAddressingType(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithAddressingType(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAddressingType", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetInstanceType2(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithInstanceType2(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetPlacement(a ServicesEc2ModelSpotPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotPlacement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotPlacement getPlacement()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetPlacement() *ServicesEc2ModelSpotPlacement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlacement", "com/amazonaws/services/ec2/model/SpotPlacement")
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
	unique_x := &ServicesEc2ModelSpotPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithPlacement(a ServicesEc2ModelSpotPlacementInterface) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacement", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotPlacement"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetKernelId(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetKernelId() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithKernelId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetRamdiskId(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetRamdiskId() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithRamdiskId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping)  {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/BlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BlockDeviceMapping"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoring(com.amazonaws.services.ec2.model.SpotFleetMonitoring)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetMonitoring(a ServicesEc2ModelSpotFleetMonitoringInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoring", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetMonitoring"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SpotFleetMonitoring getMonitoring()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetMonitoring() *ServicesEc2ModelSpotFleetMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoring", "com/amazonaws/services/ec2/model/SpotFleetMonitoring")
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
	unique_x := &ServicesEc2ModelSpotFleetMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withMonitoring(com.amazonaws.services.ec2.model.SpotFleetMonitoring)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithMonitoring(a ServicesEc2ModelSpotFleetMonitoringInterface) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoring", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetMonitoring"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithSubnetId(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification> getNetworkInterfaces()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterfaceSpecification {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification)  {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification...)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceSpecification"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfileSpecification {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIamInstanceProfile", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IamInstanceProfileSpecification"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetEbsOptimized(a bool)  {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetEbsOptimized() bool {
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithEbsOptimized(a bool) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) IsEbsOptimized() bool {
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

// public void setWeightedCapacity(java.lang.Double)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetWeightedCapacity(a float64)  {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setWeightedCapacity", javabind.Void, conv_a.Value().Cast("java/lang/Double"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Double getWeightedCapacity()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetWeightedCapacity() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getWeightedCapacity", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withWeightedCapacity(java.lang.Double)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithWeightedCapacity(a float64) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withWeightedCapacity", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) SetSpotPrice(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotPrice", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotPrice()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) GetSpotPrice() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotPrice", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification withSpotPrice(java.lang.String)
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) WithSpotPrice(a string) *ServicesEc2ModelSpotFleetLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotPrice", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SpotFleetLaunchSpecification clone()
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) Clone() *ServicesEc2ModelSpotFleetLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SpotFleetLaunchSpecification")
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
	unique_x := &ServicesEc2ModelSpotFleetLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSpotFleetLaunchSpecification) Clone2() (*JavaLangObject, error) {
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


