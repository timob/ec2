package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceInterface interface {
	JavaLangObjectInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelNetworkInterface

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelNetworkInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelNetworkInterface

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelNetworkInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelNetworkInterface

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelNetworkInterface

	// public void setRequesterId(java.lang.String)
	SetRequesterId(a string) 

	// public java.lang.String getRequesterId()
	GetRequesterId() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withRequesterId(java.lang.String)
	WithRequesterId(a string) *ServicesEc2ModelNetworkInterface

	// public void setRequesterManaged(java.lang.Boolean)
	SetRequesterManaged(a bool) 

	// public java.lang.Boolean getRequesterManaged()
	GetRequesterManaged() bool

	// public com.amazonaws.services.ec2.model.NetworkInterface withRequesterManaged(java.lang.Boolean)
	WithRequesterManaged(a bool) *ServicesEc2ModelNetworkInterface

	// public java.lang.Boolean isRequesterManaged()
	IsRequesterManaged() bool

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelNetworkInterface

	// public void setStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
	SetStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterface withStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
	WithStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface) *ServicesEc2ModelNetworkInterface

	// public void setMacAddress(java.lang.String)
	SetMacAddress(a string) 

	// public java.lang.String getMacAddress()
	GetMacAddress() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withMacAddress(java.lang.String)
	WithMacAddress(a string) *ServicesEc2ModelNetworkInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelNetworkInterface

	// public void setPrivateDnsName(java.lang.String)
	SetPrivateDnsName(a string) 

	// public java.lang.String getPrivateDnsName()
	GetPrivateDnsName() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateDnsName(java.lang.String)
	WithPrivateDnsName(a string) *ServicesEc2ModelNetworkInterface

	// public void setSourceDestCheck(java.lang.Boolean)
	SetSourceDestCheck(a bool) 

	// public java.lang.Boolean getSourceDestCheck()
	GetSourceDestCheck() bool

	// public com.amazonaws.services.ec2.model.NetworkInterface withSourceDestCheck(java.lang.Boolean)
	WithSourceDestCheck(a bool) *ServicesEc2ModelNetworkInterface

	// public java.lang.Boolean isSourceDestCheck()
	IsSourceDestCheck() bool

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getGroups()
	GetGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.NetworkInterface withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelNetworkInterface

	// public com.amazonaws.services.ec2.model.NetworkInterface withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelNetworkInterface

	// public void setAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachment)
	SetAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment getAttachment()
	GetAttachment() *ServicesEc2ModelNetworkInterfaceAttachment

	// public com.amazonaws.services.ec2.model.NetworkInterface withAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachment)
	WithAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentInterface) *ServicesEc2ModelNetworkInterface

	// public void setAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
	SetAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation getAssociation()
	GetAssociation() *ServicesEc2ModelNetworkInterfaceAssociation

	// public com.amazonaws.services.ec2.model.NetworkInterface withAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
	WithAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface) *ServicesEc2ModelNetworkInterface

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTagSet()
	GetTagSet() []*ServicesEc2ModelTag

	// public void setTagSet(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTagSet(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.NetworkInterface withTagSet(com.amazonaws.services.ec2.model.Tag...)
	WithTagSet(a ...*ServicesEc2ModelTag) *ServicesEc2ModelNetworkInterface

	// public com.amazonaws.services.ec2.model.NetworkInterface withTagSet(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTagSet2(a []*ServicesEc2ModelTag) *ServicesEc2ModelNetworkInterface

	// public java.util.List<com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress> getPrivateIpAddresses()
	GetPrivateIpAddresses() []*ServicesEc2ModelNetworkInterfacePrivateIpAddress

	// public void setPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress>)
	SetPrivateIpAddresses(a []*ServicesEc2ModelNetworkInterfacePrivateIpAddress) 

	// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateIpAddresses(com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress...)
	WithPrivateIpAddresses(a ...*ServicesEc2ModelNetworkInterfacePrivateIpAddress) *ServicesEc2ModelNetworkInterface

	// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress>)
	WithPrivateIpAddresses2(a []*ServicesEc2ModelNetworkInterfacePrivateIpAddress) *ServicesEc2ModelNetworkInterface

	// public void setInterfaceType(java.lang.String)
	SetInterfaceType2(a string) 

	// public java.lang.String getInterfaceType()
	GetInterfaceType() string

	// public com.amazonaws.services.ec2.model.NetworkInterface withInterfaceType(java.lang.String)
	WithInterfaceType2(a string) *ServicesEc2ModelNetworkInterface

	// public void setInterfaceType(com.amazonaws.services.ec2.model.NetworkInterfaceType)
	SetInterfaceType(a ServicesEc2ModelNetworkInterfaceTypeInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterface withInterfaceType(com.amazonaws.services.ec2.model.NetworkInterfaceType)
	WithInterfaceType(a ServicesEc2ModelNetworkInterfaceTypeInterface) *ServicesEc2ModelNetworkInterface

	// public com.amazonaws.services.ec2.model.NetworkInterface clone()
	Clone() *ServicesEc2ModelNetworkInterface
}

type ServicesEc2ModelNetworkInterface struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkInterface()
func NewServicesEc2ModelNetworkInterface() (*ServicesEc2ModelNetworkInterface) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkInterface")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkInterface{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetNetworkInterfaceId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetNetworkInterfaceId() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithNetworkInterfaceId(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithSubnetId(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithVpcId(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelNetworkInterface) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterface withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithAvailabilityZone(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithDescription(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetOwnerId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetOwnerId() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithOwnerId(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRequesterId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetRequesterId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequesterId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRequesterId()
func (jbobject *ServicesEc2ModelNetworkInterface) GetRequesterId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequesterId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterface withRequesterId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithRequesterId(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRequesterId", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRequesterManaged(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterface) SetRequesterManaged(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequesterManaged", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getRequesterManaged()
func (jbobject *ServicesEc2ModelNetworkInterface) GetRequesterManaged() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequesterManaged", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.NetworkInterface withRequesterManaged(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterface) WithRequesterManaged(a bool) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRequesterManaged", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isRequesterManaged()
func (jbobject *ServicesEc2ModelNetworkInterface) IsRequesterManaged() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isRequesterManaged", "java/lang/Boolean")
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

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetStatus2(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithStatus2(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
func (jbobject *ServicesEc2ModelNetworkInterface) SetStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface)  {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withStatus(com.amazonaws.services.ec2.model.NetworkInterfaceStatus)
func (jbobject *ServicesEc2ModelNetworkInterface) WithStatus(a ServicesEc2ModelNetworkInterfaceStatusInterface) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMacAddress(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetMacAddress(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetMacAddress() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withMacAddress(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithMacAddress(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMacAddress", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithPrivateIpAddress(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetPrivateDnsName(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetPrivateDnsName() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithPrivateDnsName(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateDnsName", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterface) SetSourceDestCheck(a bool)  {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetSourceDestCheck() bool {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterface) WithSourceDestCheck(a bool) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceDestCheck", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isSourceDestCheck()
func (jbobject *ServicesEc2ModelNetworkInterface) IsSourceDestCheck() bool {
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
func (jbobject *ServicesEc2ModelNetworkInterface) GetGroups() []*ServicesEc2ModelGroupIdentifier {
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
func (jbobject *ServicesEc2ModelNetworkInterface) SetGroups(a []*ServicesEc2ModelGroupIdentifier)  {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelNetworkInterface) WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkInterface withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelNetworkInterface) WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachment)
func (jbobject *ServicesEc2ModelNetworkInterface) SetAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAttachment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment getAttachment()
func (jbobject *ServicesEc2ModelNetworkInterface) GetAttachment() *ServicesEc2ModelNetworkInterfaceAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachment", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkInterface withAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachment)
func (jbobject *ServicesEc2ModelNetworkInterface) WithAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentInterface) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachment", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAttachment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelNetworkInterface) SetAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociation", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAssociation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation getAssociation()
func (jbobject *ServicesEc2ModelNetworkInterface) GetAssociation() *ServicesEc2ModelNetworkInterfaceAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociation", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkInterface withAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelNetworkInterface) WithAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociation", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAssociation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTagSet()
func (jbobject *ServicesEc2ModelNetworkInterface) GetTagSet() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTagSet", "java/util/List")
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

// public void setTagSet(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelNetworkInterface) SetTagSet(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTagSet", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkInterface withTagSet(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelNetworkInterface) WithTagSet(a ...*ServicesEc2ModelTag) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTagSet", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkInterface withTagSet(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelNetworkInterface) WithTagSet2(a []*ServicesEc2ModelTag) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTagSet", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress> getPrivateIpAddresses()
func (jbobject *ServicesEc2ModelNetworkInterface) GetPrivateIpAddresses() []*ServicesEc2ModelNetworkInterfacePrivateIpAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddresses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNetworkInterfacePrivateIpAddress)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress>)
func (jbobject *ServicesEc2ModelNetworkInterface) SetPrivateIpAddresses(a []*ServicesEc2ModelNetworkInterfacePrivateIpAddress)  {
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

// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateIpAddresses(com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress...)
func (jbobject *ServicesEc2ModelNetworkInterface) WithPrivateIpAddresses(a ...*ServicesEc2ModelNetworkInterfacePrivateIpAddress) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkInterface withPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress>)
func (jbobject *ServicesEc2ModelNetworkInterface) WithPrivateIpAddresses2(a []*ServicesEc2ModelNetworkInterfacePrivateIpAddress) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInterfaceType(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) SetInterfaceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInterfaceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInterfaceType()
func (jbobject *ServicesEc2ModelNetworkInterface) GetInterfaceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInterfaceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterface withInterfaceType(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterface) WithInterfaceType2(a string) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInterfaceType", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInterfaceType(com.amazonaws.services.ec2.model.NetworkInterfaceType)
func (jbobject *ServicesEc2ModelNetworkInterface) SetInterfaceType(a ServicesEc2ModelNetworkInterfaceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInterfaceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkInterface withInterfaceType(com.amazonaws.services.ec2.model.NetworkInterfaceType)
func (jbobject *ServicesEc2ModelNetworkInterface) WithInterfaceType(a ServicesEc2ModelNetworkInterfaceTypeInterface) *ServicesEc2ModelNetworkInterface {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInterfaceType", "com/amazonaws/services/ec2/model/NetworkInterface", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkInterface) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkInterface) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkInterface) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkInterface clone()
func (jbobject *ServicesEc2ModelNetworkInterface) Clone() *ServicesEc2ModelNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkInterface")
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
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkInterface) Clone2() (*JavaLangObject, error) {
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


