package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceNetworkInterfaceInterface interface {
	JavaLangObjectInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
	SetStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface) 

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
	WithStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface) *ServicesEc2ModelInstanceNetworkInterface

	// public void setMacAddress(java.lang.String)
	SetMacAddress(a string) 

	// public java.lang.String getMacAddress()
	GetMacAddress() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withMacAddress(java.lang.String)
	WithMacAddress(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setPrivateDnsName(java.lang.String)
	SetPrivateDnsName(a string) 

	// public java.lang.String getPrivateDnsName()
	GetPrivateDnsName() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateDnsName(java.lang.String)
	WithPrivateDnsName(a string) *ServicesEc2ModelInstanceNetworkInterface

	// public void setSourceDestCheck(java.lang.Boolean)
	SetSourceDestCheck(a bool) 

	// public java.lang.Boolean getSourceDestCheck()
	GetSourceDestCheck() bool

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withSourceDestCheck(java.lang.Boolean)
	WithSourceDestCheck(a bool) *ServicesEc2ModelInstanceNetworkInterface

	// public java.lang.Boolean isSourceDestCheck()
	IsSourceDestCheck() bool

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getGroups()
	GetGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceNetworkInterface

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceNetworkInterface

	// public void setAttachment(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAttachment)
	SetAttachment(a ServicesEc2ModelInstanceNetworkInterfaceAttachmentInterface) 

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAttachment getAttachment()
	GetAttachment() *ServicesEc2ModelInstanceNetworkInterfaceAttachment

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withAttachment(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAttachment)
	WithAttachment(a ServicesEc2ModelInstanceNetworkInterfaceAttachmentInterface) *ServicesEc2ModelInstanceNetworkInterface

	// public void setAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
	SetAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface) 

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation getAssociation()
	GetAssociation() *ServicesEc2ModelInstanceNetworkInterfaceAssociation

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
	WithAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface) *ServicesEc2ModelInstanceNetworkInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstancePrivateIpAddress> getPrivateIpAddresses()
	GetPrivateIpAddresses() []*ServicesEc2ModelInstancePrivateIpAddress

	// public void setPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.InstancePrivateIpAddress>)
	SetPrivateIpAddresses(a []*ServicesEc2ModelInstancePrivateIpAddress) 

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateIpAddresses(com.amazonaws.services.ec2.model.InstancePrivateIpAddress...)
	WithPrivateIpAddresses(a ...*ServicesEc2ModelInstancePrivateIpAddress) *ServicesEc2ModelInstanceNetworkInterface

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.InstancePrivateIpAddress>)
	WithPrivateIpAddresses2(a []*ServicesEc2ModelInstancePrivateIpAddress) *ServicesEc2ModelInstanceNetworkInterface

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterface clone()
	Clone() *ServicesEc2ModelInstanceNetworkInterface
}

type ServicesEc2ModelInstanceNetworkInterface struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface()
func NewServicesEc2ModelInstanceNetworkInterface() (*ServicesEc2ModelInstanceNetworkInterface) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceNetworkInterface")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceNetworkInterface{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetNetworkInterfaceId(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetNetworkInterfaceId() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithNetworkInterfaceId(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithSubnetId(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithVpcId(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithDescription(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOwnerId()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithOwnerId(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithStatus2(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMacAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetMacAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMacAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getMacAddress()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetMacAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMacAddress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withMacAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithMacAddress(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMacAddress", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithPrivateIpAddress(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetPrivateDnsName(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetPrivateDnsName() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithPrivateDnsName(a string) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateDnsName", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetSourceDestCheck(a bool)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetSourceDestCheck() bool {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithSourceDestCheck(a bool) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceDestCheck", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isSourceDestCheck()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) IsSourceDestCheck() bool {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetGroups() []*ServicesEc2ModelGroupIdentifier {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetGroups(a []*ServicesEc2ModelGroupIdentifier)  {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttachment(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAttachment)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetAttachment(a ServicesEc2ModelInstanceNetworkInterfaceAttachmentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAttachment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAttachment getAttachment()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetAttachment() *ServicesEc2ModelInstanceNetworkInterfaceAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachment", "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAttachment")
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
	unique_x := &ServicesEc2ModelInstanceNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withAttachment(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAttachment)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithAttachment(a ServicesEc2ModelInstanceNetworkInterfaceAttachmentInterface) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachment", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAttachment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociation", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation getAssociation()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetAssociation() *ServicesEc2ModelInstanceNetworkInterfaceAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociation", "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation")
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
	unique_x := &ServicesEc2ModelInstanceNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociation", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstancePrivateIpAddress> getPrivateIpAddresses()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) GetPrivateIpAddresses() []*ServicesEc2ModelInstancePrivateIpAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddresses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstancePrivateIpAddress)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.InstancePrivateIpAddress>)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) SetPrivateIpAddresses(a []*ServicesEc2ModelInstancePrivateIpAddress)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddresses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateIpAddresses(com.amazonaws.services.ec2.model.InstancePrivateIpAddress...)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithPrivateIpAddresses(a ...*ServicesEc2ModelInstancePrivateIpAddress) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstancePrivateIpAddress")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstancePrivateIpAddress"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface withPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.InstancePrivateIpAddress>)
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) WithPrivateIpAddresses2(a []*ServicesEc2ModelInstancePrivateIpAddress) *ServicesEc2ModelInstanceNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/InstanceNetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterface clone()
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) Clone() *ServicesEc2ModelInstanceNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceNetworkInterface")
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
	unique_x := &ServicesEc2ModelInstanceNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceNetworkInterface) Clone2() (*JavaLangObject, error) {
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


