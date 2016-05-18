package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelLaunchSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelLaunchSpecification

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelLaunchSpecification

	// public void setUserData(java.lang.String)
	SetUserData(a string) 

	// public java.lang.String getUserData()
	GetUserData() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withUserData(java.lang.String)
	WithUserData(a string) *ServicesEc2ModelLaunchSpecification

	// public void setAddressingType(java.lang.String)
	SetAddressingType(a string) 

	// public java.lang.String getAddressingType()
	GetAddressingType() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withAddressingType(java.lang.String)
	WithAddressingType(a string) *ServicesEc2ModelLaunchSpecification

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelLaunchSpecification

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelLaunchSpecification

	// public void setPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
	SetPlacement(a ServicesEc2ModelSpotPlacementInterface) 

	// public com.amazonaws.services.ec2.model.SpotPlacement getPlacement()
	GetPlacement() *ServicesEc2ModelSpotPlacement

	// public com.amazonaws.services.ec2.model.LaunchSpecification withPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
	WithPlacement(a ServicesEc2ModelSpotPlacementInterface) *ServicesEc2ModelLaunchSpecification

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelLaunchSpecification

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.LaunchSpecification withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelLaunchSpecification

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.LaunchSpecification withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification> getNetworkInterfaces()
	GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterfaceSpecification

	// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
	SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification...)
	WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.LaunchSpecification withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
	WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelLaunchSpecification

	// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
	SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification getIamInstanceProfile()
	GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfileSpecification

	// public com.amazonaws.services.ec2.model.LaunchSpecification withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
	WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) *ServicesEc2ModelLaunchSpecification

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.LaunchSpecification withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelLaunchSpecification

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getAllSecurityGroups()
	GetAllSecurityGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setAllSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetAllSecurityGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification withAllSecurityGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithAllSecurityGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.LaunchSpecification withAllSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithAllSecurityGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelLaunchSpecification

	// public void setMonitoringEnabled(java.lang.Boolean)
	SetMonitoringEnabled(a bool) 

	// public java.lang.Boolean getMonitoringEnabled()
	GetMonitoringEnabled() bool

	// public com.amazonaws.services.ec2.model.LaunchSpecification withMonitoringEnabled(java.lang.Boolean)
	WithMonitoringEnabled(a bool) *ServicesEc2ModelLaunchSpecification

	// public java.lang.Boolean isMonitoringEnabled()
	IsMonitoringEnabled() bool

	// public java.util.List<java.lang.String> getSecurityGroups()
	GetSecurityGroups() []string

	// public void setSecurityGroups(java.util.Collection<java.lang.String>)
	SetSecurityGroups(a []string) 

	// public com.amazonaws.services.ec2.model.LaunchSpecification withSecurityGroups(java.lang.String...)
	WithSecurityGroups(a ...string) *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.LaunchSpecification withSecurityGroups(java.util.Collection<java.lang.String>)
	WithSecurityGroups2(a []string) *ServicesEc2ModelLaunchSpecification

	// public com.amazonaws.services.ec2.model.LaunchSpecification clone()
	Clone() *ServicesEc2ModelLaunchSpecification
}

type ServicesEc2ModelLaunchSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.LaunchSpecification()
func NewServicesEc2ModelLaunchSpecification() (*ServicesEc2ModelLaunchSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/LaunchSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelLaunchSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetImageId(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetImageId() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithImageId(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetKeyName(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetKeyName() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithKeyName(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetUserData(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetUserData() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withUserData(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithUserData(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAddressingType(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetAddressingType(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetAddressingType() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withAddressingType(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithAddressingType(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAddressingType", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetInstanceType2(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithInstanceType2(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetPlacement(a ServicesEc2ModelSpotPlacementInterface)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetPlacement() *ServicesEc2ModelSpotPlacement {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withPlacement(com.amazonaws.services.ec2.model.SpotPlacement)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithPlacement(a ServicesEc2ModelSpotPlacementInterface) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacement", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotPlacement"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetKernelId(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetKernelId() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithKernelId(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetRamdiskId(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetRamdiskId() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithRamdiskId(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelLaunchSpecification) GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping)  {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/BlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BlockDeviceMapping"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.LaunchSpecification withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithSubnetId(a string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification> getNetworkInterfaces()
func (jbobject *ServicesEc2ModelLaunchSpecification) GetNetworkInterfaces() []*ServicesEc2ModelInstanceNetworkInterfaceSpecification {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) SetNetworkInterfaces(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification)  {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withNetworkInterfaces(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification...)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithNetworkInterfaces(a ...*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceSpecification"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.LaunchSpecification withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.InstanceNetworkInterfaceSpecification>)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithNetworkInterfaces2(a []*ServicesEc2ModelInstanceNetworkInterfaceSpecification) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetIamInstanceProfile() *ServicesEc2ModelIamInstanceProfileSpecification {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withIamInstanceProfile(com.amazonaws.services.ec2.model.IamInstanceProfileSpecification)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithIamInstanceProfile(a ServicesEc2ModelIamInstanceProfileSpecificationInterface) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIamInstanceProfile", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IamInstanceProfileSpecification"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetEbsOptimized(a bool)  {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) GetEbsOptimized() bool {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithEbsOptimized(a bool) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelLaunchSpecification) IsEbsOptimized() bool {
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

// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getAllSecurityGroups()
func (jbobject *ServicesEc2ModelLaunchSpecification) GetAllSecurityGroups() []*ServicesEc2ModelGroupIdentifier {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllSecurityGroups", "java/util/List")
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

// public void setAllSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetAllSecurityGroups(a []*ServicesEc2ModelGroupIdentifier)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllSecurityGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.LaunchSpecification withAllSecurityGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithAllSecurityGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllSecurityGroups", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.LaunchSpecification withAllSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithAllSecurityGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllSecurityGroups", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoringEnabled(java.lang.Boolean)
func (jbobject *ServicesEc2ModelLaunchSpecification) SetMonitoringEnabled(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoringEnabled", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getMonitoringEnabled()
func (jbobject *ServicesEc2ModelLaunchSpecification) GetMonitoringEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoringEnabled", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withMonitoringEnabled(java.lang.Boolean)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithMonitoringEnabled(a bool) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoringEnabled", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isMonitoringEnabled()
func (jbobject *ServicesEc2ModelLaunchSpecification) IsMonitoringEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMonitoringEnabled", "java/lang/Boolean")
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

// public java.util.List<java.lang.String> getSecurityGroups()
func (jbobject *ServicesEc2ModelLaunchSpecification) GetSecurityGroups() []string {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) SetSecurityGroups(a []string)  {
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

// public com.amazonaws.services.ec2.model.LaunchSpecification withSecurityGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithSecurityGroups(a ...string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.LaunchSpecification withSecurityGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelLaunchSpecification) WithSecurityGroups2(a []string) *ServicesEc2ModelLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/LaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelLaunchSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelLaunchSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.LaunchSpecification clone()
func (jbobject *ServicesEc2ModelLaunchSpecification) Clone() *ServicesEc2ModelLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/LaunchSpecification")
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
	unique_x := &ServicesEc2ModelLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelLaunchSpecification) Clone2() (*JavaLangObject, error) {
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


