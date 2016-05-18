package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfacePrivateIpAddressInterface interface {
	JavaLangObjectInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelNetworkInterfacePrivateIpAddress

	// public void setPrivateDnsName(java.lang.String)
	SetPrivateDnsName(a string) 

	// public java.lang.String getPrivateDnsName()
	GetPrivateDnsName() string

	// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withPrivateDnsName(java.lang.String)
	WithPrivateDnsName(a string) *ServicesEc2ModelNetworkInterfacePrivateIpAddress

	// public void setPrimary(java.lang.Boolean)
	SetPrimary(a bool) 

	// public java.lang.Boolean getPrimary()
	GetPrimary() bool

	// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withPrimary(java.lang.Boolean)
	WithPrimary(a bool) *ServicesEc2ModelNetworkInterfacePrivateIpAddress

	// public java.lang.Boolean isPrimary()
	IsPrimary() bool

	// public void setAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
	SetAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation getAssociation()
	GetAssociation() *ServicesEc2ModelNetworkInterfaceAssociation

	// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
	WithAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface) *ServicesEc2ModelNetworkInterfacePrivateIpAddress

	// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress clone()
	Clone() *ServicesEc2ModelNetworkInterfacePrivateIpAddress
}

type ServicesEc2ModelNetworkInterfacePrivateIpAddress struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress()
func NewServicesEc2ModelNetworkInterfacePrivateIpAddress() (*ServicesEc2ModelNetworkInterfacePrivateIpAddress) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkInterfacePrivateIpAddress{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) WithPrivateIpAddress(a string) *ServicesEc2ModelNetworkInterfacePrivateIpAddress {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterfacePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) SetPrivateDnsName(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) GetPrivateDnsName() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) WithPrivateDnsName(a string) *ServicesEc2ModelNetworkInterfacePrivateIpAddress {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateDnsName", "com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterfacePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) SetPrimary(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrimary", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getPrimary()
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) GetPrimary() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrimary", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) WithPrimary(a bool) *ServicesEc2ModelNetworkInterfacePrivateIpAddress {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrimary", "com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress", conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterfacePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isPrimary()
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) IsPrimary() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPrimary", "java/lang/Boolean")
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

// public void setAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) SetAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) GetAssociation() *ServicesEc2ModelNetworkInterfaceAssociation {
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

// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress withAssociation(com.amazonaws.services.ec2.model.NetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) WithAssociation(a ServicesEc2ModelNetworkInterfaceAssociationInterface) *ServicesEc2ModelNetworkInterfacePrivateIpAddress {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociation", "com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAssociation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelNetworkInterfacePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkInterfacePrivateIpAddress clone()
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) Clone() *ServicesEc2ModelNetworkInterfacePrivateIpAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkInterfacePrivateIpAddress")
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
	unique_x := &ServicesEc2ModelNetworkInterfacePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkInterfacePrivateIpAddress) Clone2() (*JavaLangObject, error) {
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


