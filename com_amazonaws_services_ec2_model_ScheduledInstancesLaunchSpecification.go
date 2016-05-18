package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstancesLaunchSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public java.util.List<java.lang.String> getSecurityGroupIds()
	GetSecurityGroupIds() []string

	// public void setSecurityGroupIds(java.util.Collection<java.lang.String>)
	SetSecurityGroupIds(a []string) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withSecurityGroupIds(java.lang.String...)
	WithSecurityGroupIds(a ...string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withSecurityGroupIds(java.util.Collection<java.lang.String>)
	WithSecurityGroupIds2(a []string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setUserData(java.lang.String)
	SetUserData(a string) 

	// public java.lang.String getUserData()
	GetUserData() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withUserData(java.lang.String)
	WithUserData(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setPlacement(com.amazonaws.services.ec2.model.ScheduledInstancesPlacement)
	SetPlacement(a ServicesEc2ModelScheduledInstancesPlacementInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesPlacement getPlacement()
	GetPlacement() *ServicesEc2ModelScheduledInstancesPlacement

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withPlacement(com.amazonaws.services.ec2.model.ScheduledInstancesPlacement)
	WithPlacement(a ServicesEc2ModelScheduledInstancesPlacementInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelScheduledInstancesBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelScheduledInstancesBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withBlockDeviceMappings(com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelScheduledInstancesBlockDeviceMapping) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelScheduledInstancesBlockDeviceMapping) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setMonitoring(com.amazonaws.services.ec2.model.ScheduledInstancesMonitoring)
	SetMonitoring(a ServicesEc2ModelScheduledInstancesMonitoringInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesMonitoring getMonitoring()
	GetMonitoring() *ServicesEc2ModelScheduledInstancesMonitoring

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withMonitoring(com.amazonaws.services.ec2.model.ScheduledInstancesMonitoring)
	WithMonitoring(a ServicesEc2ModelScheduledInstancesMonitoringInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface> getNetworkInterfaces()
	GetNetworkInterfaces() []*ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface>)
	SetNetworkInterfaces(a []*ServicesEc2ModelScheduledInstancesNetworkInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withNetworkInterfaces(com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface...)
	WithNetworkInterfaces(a ...*ServicesEc2ModelScheduledInstancesNetworkInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface>)
	WithNetworkInterfaces2(a []*ServicesEc2ModelScheduledInstancesNetworkInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile)
	SetIamInstanceProfile(a ServicesEc2ModelScheduledInstancesIamInstanceProfileInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile getIamInstanceProfile()
	GetIamInstanceProfile() *ServicesEc2ModelScheduledInstancesIamInstanceProfile

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withIamInstanceProfile(com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile)
	WithIamInstanceProfile(a ServicesEc2ModelScheduledInstancesIamInstanceProfileInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelScheduledInstancesLaunchSpecification

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification clone()
	Clone() *ServicesEc2ModelScheduledInstancesLaunchSpecification
}

type ServicesEc2ModelScheduledInstancesLaunchSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification()
func NewServicesEc2ModelScheduledInstancesLaunchSpecification() (*ServicesEc2ModelScheduledInstancesLaunchSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetImageId(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetImageId() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithImageId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetKeyName(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetKeyName() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithKeyName(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getSecurityGroupIds()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetSecurityGroupIds() []string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetSecurityGroupIds(a []string)  {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withSecurityGroupIds(java.lang.String...)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithSecurityGroupIds(a ...string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroupIds", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withSecurityGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithSecurityGroupIds2(a []string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroupIds", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetUserData(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetUserData() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withUserData(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithUserData(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlacement(com.amazonaws.services.ec2.model.ScheduledInstancesPlacement)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetPlacement(a ServicesEc2ModelScheduledInstancesPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesPlacement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesPlacement getPlacement()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetPlacement() *ServicesEc2ModelScheduledInstancesPlacement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlacement", "com/amazonaws/services/ec2/model/ScheduledInstancesPlacement")
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
	unique_x := &ServicesEc2ModelScheduledInstancesPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withPlacement(com.amazonaws.services.ec2.model.ScheduledInstancesPlacement)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithPlacement(a ServicesEc2ModelScheduledInstancesPlacementInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacement", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesPlacement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetKernelId(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetKernelId() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithKernelId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetInstanceType(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithInstanceType(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetRamdiskId(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetRamdiskId() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithRamdiskId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetBlockDeviceMappings() []*ServicesEc2ModelScheduledInstancesBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockDeviceMappings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelScheduledInstancesBlockDeviceMapping)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping>)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetBlockDeviceMappings(a []*ServicesEc2ModelScheduledInstancesBlockDeviceMapping)  {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withBlockDeviceMappings(com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping...)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithBlockDeviceMappings(a ...*ServicesEc2ModelScheduledInstancesBlockDeviceMapping) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping>)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithBlockDeviceMappings2(a []*ServicesEc2ModelScheduledInstancesBlockDeviceMapping) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoring(com.amazonaws.services.ec2.model.ScheduledInstancesMonitoring)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetMonitoring(a ServicesEc2ModelScheduledInstancesMonitoringInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoring", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesMonitoring"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesMonitoring getMonitoring()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetMonitoring() *ServicesEc2ModelScheduledInstancesMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoring", "com/amazonaws/services/ec2/model/ScheduledInstancesMonitoring")
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
	unique_x := &ServicesEc2ModelScheduledInstancesMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withMonitoring(com.amazonaws.services.ec2.model.ScheduledInstancesMonitoring)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithMonitoring(a ServicesEc2ModelScheduledInstancesMonitoringInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoring", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesMonitoring"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithSubnetId(a string) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface> getNetworkInterfaces()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetNetworkInterfaces() []*ServicesEc2ModelScheduledInstancesNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaces", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelScheduledInstancesNetworkInterface)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface>)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetNetworkInterfaces(a []*ServicesEc2ModelScheduledInstancesNetworkInterface)  {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withNetworkInterfaces(com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface...)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithNetworkInterfaces(a ...*ServicesEc2ModelScheduledInstancesNetworkInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface>)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithNetworkInterfaces2(a []*ServicesEc2ModelScheduledInstancesNetworkInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIamInstanceProfile(com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetIamInstanceProfile(a ServicesEc2ModelScheduledInstancesIamInstanceProfileInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIamInstanceProfile", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile getIamInstanceProfile()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetIamInstanceProfile() *ServicesEc2ModelScheduledInstancesIamInstanceProfile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIamInstanceProfile", "com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile")
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
	unique_x := &ServicesEc2ModelScheduledInstancesIamInstanceProfile{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withIamInstanceProfile(com.amazonaws.services.ec2.model.ScheduledInstancesIamInstanceProfile)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithIamInstanceProfile(a ServicesEc2ModelScheduledInstancesIamInstanceProfileInterface) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIamInstanceProfile", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesIamInstanceProfile"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) SetEbsOptimized(a bool)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) GetEbsOptimized() bool {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) WithEbsOptimized(a bool) *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) IsEbsOptimized() bool {
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

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesLaunchSpecification clone()
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) Clone() *ServicesEc2ModelScheduledInstancesLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstancesLaunchSpecification")
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
	unique_x := &ServicesEc2ModelScheduledInstancesLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstancesLaunchSpecification) Clone2() (*JavaLangObject, error) {
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


