package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceNetworkInterfaceAssociationInterface interface {
	JavaLangObjectInterface

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelInstanceNetworkInterfaceAssociation

	// public void setPublicDnsName(java.lang.String)
	SetPublicDnsName(a string) 

	// public java.lang.String getPublicDnsName()
	GetPublicDnsName() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation withPublicDnsName(java.lang.String)
	WithPublicDnsName(a string) *ServicesEc2ModelInstanceNetworkInterfaceAssociation

	// public void setIpOwnerId(java.lang.String)
	SetIpOwnerId(a string) 

	// public java.lang.String getIpOwnerId()
	GetIpOwnerId() string

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation withIpOwnerId(java.lang.String)
	WithIpOwnerId(a string) *ServicesEc2ModelInstanceNetworkInterfaceAssociation

	// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation clone()
	Clone() *ServicesEc2ModelInstanceNetworkInterfaceAssociation
}

type ServicesEc2ModelInstanceNetworkInterfaceAssociation struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation()
func NewServicesEc2ModelInstanceNetworkInterfaceAssociation() (*ServicesEc2ModelInstanceNetworkInterfaceAssociation) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceNetworkInterfaceAssociation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) SetPublicIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIp()
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) GetPublicIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) WithPublicIp(a string) *ServicesEc2ModelInstanceNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublicDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) SetPublicDnsName(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) GetPublicDnsName() string {
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation withPublicDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) WithPublicDnsName(a string) *ServicesEc2ModelInstanceNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicDnsName", "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIpOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) SetIpOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getIpOwnerId()
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) GetIpOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation withIpOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) WithIpOwnerId(a string) *ServicesEc2ModelInstanceNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpOwnerId", "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceNetworkInterfaceAssociation clone()
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) Clone() *ServicesEc2ModelInstanceNetworkInterfaceAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceNetworkInterfaceAssociation")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceNetworkInterfaceAssociation) Clone2() (*JavaLangObject, error) {
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


