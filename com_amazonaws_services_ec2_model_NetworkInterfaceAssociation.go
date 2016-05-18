package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceAssociationInterface interface {
	JavaLangObjectInterface

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelNetworkInterfaceAssociation

	// public void setPublicDnsName(java.lang.String)
	SetPublicDnsName(a string) 

	// public java.lang.String getPublicDnsName()
	GetPublicDnsName() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withPublicDnsName(java.lang.String)
	WithPublicDnsName(a string) *ServicesEc2ModelNetworkInterfaceAssociation

	// public void setIpOwnerId(java.lang.String)
	SetIpOwnerId(a string) 

	// public java.lang.String getIpOwnerId()
	GetIpOwnerId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withIpOwnerId(java.lang.String)
	WithIpOwnerId(a string) *ServicesEc2ModelNetworkInterfaceAssociation

	// public void setAllocationId(java.lang.String)
	SetAllocationId(a string) 

	// public java.lang.String getAllocationId()
	GetAllocationId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withAllocationId(java.lang.String)
	WithAllocationId(a string) *ServicesEc2ModelNetworkInterfaceAssociation

	// public void setAssociationId(java.lang.String)
	SetAssociationId(a string) 

	// public java.lang.String getAssociationId()
	GetAssociationId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withAssociationId(java.lang.String)
	WithAssociationId(a string) *ServicesEc2ModelNetworkInterfaceAssociation

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation clone()
	Clone() *ServicesEc2ModelNetworkInterfaceAssociation
}

type ServicesEc2ModelNetworkInterfaceAssociation struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation()
func NewServicesEc2ModelNetworkInterfaceAssociation() (*ServicesEc2ModelNetworkInterfaceAssociation) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkInterfaceAssociation")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) SetPublicIp(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) GetPublicIp() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) WithPublicIp(a string) *ServicesEc2ModelNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublicDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) SetPublicDnsName(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) GetPublicDnsName() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withPublicDnsName(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) WithPublicDnsName(a string) *ServicesEc2ModelNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicDnsName", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIpOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) SetIpOwnerId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) GetIpOwnerId() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withIpOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) WithIpOwnerId(a string) *ServicesEc2ModelNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpOwnerId", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAllocationId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) SetAllocationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllocationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAllocationId()
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) GetAllocationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllocationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withAllocationId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) WithAllocationId(a string) *ServicesEc2ModelNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllocationId", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) SetAssociationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAssociationId()
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) GetAssociationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation withAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) WithAssociationId(a string) *ServicesEc2ModelNetworkInterfaceAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociationId", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAssociation clone()
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) Clone() *ServicesEc2ModelNetworkInterfaceAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkInterfaceAssociation")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkInterfaceAssociation) Clone2() (*JavaLangObject, error) {
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


