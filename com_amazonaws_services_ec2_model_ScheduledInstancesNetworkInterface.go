package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstancesNetworkInterfaceInterface interface {
	JavaLangObjectInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setDeviceIndex(java.lang.Integer)
	SetDeviceIndex(a int) 

	// public java.lang.Integer getDeviceIndex()
	GetDeviceIndex() int

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withDeviceIndex(java.lang.Integer)
	WithDeviceIndex(a int) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig> getPrivateIpAddressConfigs()
	GetPrivateIpAddressConfigs() []*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig

	// public void setPrivateIpAddressConfigs(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig>)
	SetPrivateIpAddressConfigs(a []*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withPrivateIpAddressConfigs(com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig...)
	WithPrivateIpAddressConfigs(a ...*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withPrivateIpAddressConfigs(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig>)
	WithPrivateIpAddressConfigs2(a []*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setSecondaryPrivateIpAddressCount(java.lang.Integer)
	SetSecondaryPrivateIpAddressCount(a int) 

	// public java.lang.Integer getSecondaryPrivateIpAddressCount()
	GetSecondaryPrivateIpAddressCount() int

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withSecondaryPrivateIpAddressCount(java.lang.Integer)
	WithSecondaryPrivateIpAddressCount(a int) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setAssociatePublicIpAddress(java.lang.Boolean)
	SetAssociatePublicIpAddress(a bool) 

	// public java.lang.Boolean getAssociatePublicIpAddress()
	GetAssociatePublicIpAddress() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withAssociatePublicIpAddress(java.lang.Boolean)
	WithAssociatePublicIpAddress(a bool) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public java.lang.Boolean isAssociatePublicIpAddress()
	IsAssociatePublicIpAddress() bool

	// public java.util.List<java.lang.String> getGroups()
	GetGroups() []string

	// public void setGroups(java.util.Collection<java.lang.String>)
	SetGroups(a []string) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withGroups(java.lang.String...)
	WithGroups(a ...string) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withGroups(java.util.Collection<java.lang.String>)
	WithGroups2(a []string) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public void setDeleteOnTermination(java.lang.Boolean)
	SetDeleteOnTermination(a bool) 

	// public java.lang.Boolean getDeleteOnTermination()
	GetDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withDeleteOnTermination(java.lang.Boolean)
	WithDeleteOnTermination(a bool) *ServicesEc2ModelScheduledInstancesNetworkInterface

	// public java.lang.Boolean isDeleteOnTermination()
	IsDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface clone()
	Clone() *ServicesEc2ModelScheduledInstancesNetworkInterface
}

type ServicesEc2ModelScheduledInstancesNetworkInterface struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface()
func NewServicesEc2ModelScheduledInstancesNetworkInterface() (*ServicesEc2ModelScheduledInstancesNetworkInterface) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetNetworkInterfaceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterfaceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkInterfaceId()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetNetworkInterfaceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithNetworkInterfaceId(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeviceIndex(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetDeviceIndex(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeviceIndex", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getDeviceIndex()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetDeviceIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeviceIndex", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withDeviceIndex(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithDeviceIndex(a int) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceIndex", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithSubnetId(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithDescription(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithPrivateIpAddress(a string) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig> getPrivateIpAddressConfigs()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetPrivateIpAddressConfigs() []*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddressConfigs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPrivateIpAddressConfigs(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig>)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetPrivateIpAddressConfigs(a []*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddressConfigs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withPrivateIpAddressConfigs(com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig...)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithPrivateIpAddressConfigs(a ...*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ScheduledInstancesPrivateIpAddressConfig")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddressConfigs", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesPrivateIpAddressConfig"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withPrivateIpAddressConfigs(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig>)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithPrivateIpAddressConfigs2(a []*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddressConfigs", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSecondaryPrivateIpAddressCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetSecondaryPrivateIpAddressCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSecondaryPrivateIpAddressCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getSecondaryPrivateIpAddressCount()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetSecondaryPrivateIpAddressCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecondaryPrivateIpAddressCount", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withSecondaryPrivateIpAddressCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithSecondaryPrivateIpAddressCount(a int) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecondaryPrivateIpAddressCount", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAssociatePublicIpAddress(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetAssociatePublicIpAddress(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociatePublicIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getAssociatePublicIpAddress()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetAssociatePublicIpAddress() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociatePublicIpAddress", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withAssociatePublicIpAddress(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithAssociatePublicIpAddress(a bool) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociatePublicIpAddress", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isAssociatePublicIpAddress()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) IsAssociatePublicIpAddress() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAssociatePublicIpAddress", "java/lang/Boolean")
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

// public java.util.List<java.lang.String> getGroups()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetGroups() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroups", "java/util/List")
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

// public void setGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetGroups(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithGroups(a ...string) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithGroups2(a []string) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) SetDeleteOnTermination(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeleteOnTermination", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getDeleteOnTermination()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) GetDeleteOnTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeleteOnTermination", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface withDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) WithDeleteOnTermination(a bool) *ServicesEc2ModelScheduledInstancesNetworkInterface {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteOnTermination", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDeleteOnTermination()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) IsDeleteOnTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDeleteOnTermination", "java/lang/Boolean")
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
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesNetworkInterface clone()
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) Clone() *ServicesEc2ModelScheduledInstancesNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstancesNetworkInterface")
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
	unique_x := &ServicesEc2ModelScheduledInstancesNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstancesNetworkInterface) Clone2() (*JavaLangObject, error) {
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


