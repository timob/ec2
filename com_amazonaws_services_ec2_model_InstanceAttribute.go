package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceAttributeInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelInstanceAttribute

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelInstanceAttribute

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelInstanceAttribute

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelInstanceAttribute

	// public void setUserData(java.lang.String)
	SetUserData(a string) 

	// public java.lang.String getUserData()
	GetUserData() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withUserData(java.lang.String)
	WithUserData(a string) *ServicesEc2ModelInstanceAttribute

	// public void setDisableApiTermination(java.lang.Boolean)
	SetDisableApiTermination(a bool) 

	// public java.lang.Boolean getDisableApiTermination()
	GetDisableApiTermination() bool

	// public com.amazonaws.services.ec2.model.InstanceAttribute withDisableApiTermination(java.lang.Boolean)
	WithDisableApiTermination(a bool) *ServicesEc2ModelInstanceAttribute

	// public java.lang.Boolean isDisableApiTermination()
	IsDisableApiTermination() bool

	// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
	SetInstanceInitiatedShutdownBehavior(a string) 

	// public java.lang.String getInstanceInitiatedShutdownBehavior()
	GetInstanceInitiatedShutdownBehavior() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withInstanceInitiatedShutdownBehavior(java.lang.String)
	WithInstanceInitiatedShutdownBehavior(a string) *ServicesEc2ModelInstanceAttribute

	// public void setRootDeviceName(java.lang.String)
	SetRootDeviceName(a string) 

	// public java.lang.String getRootDeviceName()
	GetRootDeviceName() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withRootDeviceName(java.lang.String)
	WithRootDeviceName(a string) *ServicesEc2ModelInstanceAttribute

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelInstanceBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelInstanceBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.InstanceAttribute withBlockDeviceMappings(com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstanceAttribute

	// public com.amazonaws.services.ec2.model.InstanceAttribute withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstanceAttribute

	// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
	GetProductCodes() []*ServicesEc2ModelProductCode

	// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	SetProductCodes(a []*ServicesEc2ModelProductCode) 

	// public com.amazonaws.services.ec2.model.InstanceAttribute withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
	WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelInstanceAttribute

	// public com.amazonaws.services.ec2.model.InstanceAttribute withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelInstanceAttribute

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.InstanceAttribute withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelInstanceAttribute

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public void setSriovNetSupport(java.lang.String)
	SetSriovNetSupport(a string) 

	// public java.lang.String getSriovNetSupport()
	GetSriovNetSupport() string

	// public com.amazonaws.services.ec2.model.InstanceAttribute withSriovNetSupport(java.lang.String)
	WithSriovNetSupport(a string) *ServicesEc2ModelInstanceAttribute

	// public void setSourceDestCheck(java.lang.Boolean)
	SetSourceDestCheck(a bool) 

	// public java.lang.Boolean getSourceDestCheck()
	GetSourceDestCheck() bool

	// public com.amazonaws.services.ec2.model.InstanceAttribute withSourceDestCheck(java.lang.Boolean)
	WithSourceDestCheck(a bool) *ServicesEc2ModelInstanceAttribute

	// public java.lang.Boolean isSourceDestCheck()
	IsSourceDestCheck() bool

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getGroups()
	GetGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.InstanceAttribute withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceAttribute

	// public com.amazonaws.services.ec2.model.InstanceAttribute withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceAttribute

	// public com.amazonaws.services.ec2.model.InstanceAttribute clone()
	Clone() *ServicesEc2ModelInstanceAttribute
}

type ServicesEc2ModelInstanceAttribute struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceAttribute()
func NewServicesEc2ModelInstanceAttribute() (*ServicesEc2ModelInstanceAttribute) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceAttribute")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceAttribute{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetInstanceId(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetInstanceId() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithInstanceId(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetInstanceType(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithInstanceType(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetKernelId(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetKernelId() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithKernelId(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetRamdiskId(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetRamdiskId() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithRamdiskId(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetUserData(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetUserData() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withUserData(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithUserData(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDisableApiTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetDisableApiTermination(a bool)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetDisableApiTermination() bool {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withDisableApiTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithDisableApiTermination(a bool) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDisableApiTermination", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDisableApiTermination()
func (jbobject *ServicesEc2ModelInstanceAttribute) IsDisableApiTermination() bool {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) SetInstanceInitiatedShutdownBehavior(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetInstanceInitiatedShutdownBehavior() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithInstanceInitiatedShutdownBehavior(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetRootDeviceName(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetRootDeviceName() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withRootDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithRootDeviceName(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceName", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelInstanceAttribute) GetBlockDeviceMappings() []*ServicesEc2ModelInstanceBlockDeviceMapping {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) SetBlockDeviceMappings(a []*ServicesEc2ModelInstanceBlockDeviceMapping)  {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withBlockDeviceMappings(com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping...)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithBlockDeviceMappings(a ...*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceAttribute withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping>)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithBlockDeviceMappings2(a []*ServicesEc2ModelInstanceBlockDeviceMapping) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
func (jbobject *ServicesEc2ModelInstanceAttribute) GetProductCodes() []*ServicesEc2ModelProductCode {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) SetProductCodes(a []*ServicesEc2ModelProductCode)  {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ProductCode")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCode"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceAttribute withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetEbsOptimized(a bool)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetEbsOptimized() bool {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithEbsOptimized(a bool) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelInstanceAttribute) IsEbsOptimized() bool {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) SetSriovNetSupport(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetSriovNetSupport() string {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithSriovNetSupport(a string) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSriovNetSupport", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetSourceDestCheck(a bool)  {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) GetSourceDestCheck() bool {
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

// public com.amazonaws.services.ec2.model.InstanceAttribute withSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithSourceDestCheck(a bool) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceDestCheck", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isSourceDestCheck()
func (jbobject *ServicesEc2ModelInstanceAttribute) IsSourceDestCheck() bool {
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

// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getGroups()
func (jbobject *ServicesEc2ModelInstanceAttribute) GetGroups() []*ServicesEc2ModelGroupIdentifier {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroups", "java/util/List")
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

// public void setGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelInstanceAttribute) SetGroups(a []*ServicesEc2ModelGroupIdentifier)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceAttribute withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceAttribute withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelInstanceAttribute) WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceAttribute {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/InstanceAttribute", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceAttribute) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceAttribute) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceAttribute clone()
func (jbobject *ServicesEc2ModelInstanceAttribute) Clone() *ServicesEc2ModelInstanceAttribute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceAttribute")
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
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceAttribute) Clone2() (*JavaLangObject, error) {
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


