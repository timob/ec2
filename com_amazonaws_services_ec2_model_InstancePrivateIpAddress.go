package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstancePrivateIpAddressInterface interface {
	JavaLangObjectInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelInstancePrivateIpAddress

	// public void setPrivateDnsName(java.lang.String)
	SetPrivateDnsName(a string) 

	// public java.lang.String getPrivateDnsName()
	GetPrivateDnsName() string

	// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withPrivateDnsName(java.lang.String)
	WithPrivateDnsName(a string) *ServicesEc2ModelInstancePrivateIpAddress

	// public void setPrimary(java.lang.Boolean)
	SetPrimary(a bool) 

	// public java.lang.Boolean getPrimary()
	GetPrimary() bool

	// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withPrimary(java.lang.Boolean)
	WithPrimary(a bool) *ServicesEc2ModelInstancePrivateIpAddress

	// public java.lang.Boolean isPrimary()
	IsPrimary() bool

	// public void setAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
	SetAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface) 

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation getAssociation()
	GetAssociation() *ServicesEc2ModelInstanceNetworkInterfaceAssociation

	// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
	WithAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface) *ServicesEc2ModelInstancePrivateIpAddress

	// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress clone()
	Clone() *ServicesEc2ModelInstancePrivateIpAddress
}

type ServicesEc2ModelInstancePrivateIpAddress struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress()
func NewServicesEc2ModelInstancePrivateIpAddress() (*ServicesEc2ModelInstancePrivateIpAddress) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstancePrivateIpAddress")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstancePrivateIpAddress{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) WithPrivateIpAddress(a string) *ServicesEc2ModelInstancePrivateIpAddress {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/InstancePrivateIpAddress", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstancePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) SetPrivateDnsName(a string)  {
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
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) GetPrivateDnsName() string {
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

// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withPrivateDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) WithPrivateDnsName(a string) *ServicesEc2ModelInstancePrivateIpAddress {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateDnsName", "com/amazonaws/services/ec2/model/InstancePrivateIpAddress", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstancePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) SetPrimary(a bool)  {
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
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) GetPrimary() bool {
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

// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) WithPrimary(a bool) *ServicesEc2ModelInstancePrivateIpAddress {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrimary", "com/amazonaws/services/ec2/model/InstancePrivateIpAddress", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelInstancePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isPrimary()
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) IsPrimary() bool {
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

// public void setAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) SetAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface)  {
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
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) GetAssociation() *ServicesEc2ModelInstanceNetworkInterfaceAssociation {
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

// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress withAssociation(com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation)
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) WithAssociation(a ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface) *ServicesEc2ModelInstancePrivateIpAddress {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociation", "com/amazonaws/services/ec2/model/InstancePrivateIpAddress", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation"))
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
	unique_x := &ServicesEc2ModelInstancePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) ToString() string {
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
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstancePrivateIpAddress clone()
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) Clone() *ServicesEc2ModelInstancePrivateIpAddress {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstancePrivateIpAddress")
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
	unique_x := &ServicesEc2ModelInstancePrivateIpAddress{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstancePrivateIpAddress) Clone2() (*JavaLangObject, error) {
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


