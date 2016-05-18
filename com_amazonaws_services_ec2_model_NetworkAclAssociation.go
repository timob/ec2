package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkAclAssociationInterface interface {
	JavaLangObjectInterface

	// public void setNetworkAclAssociationId(java.lang.String)
	SetNetworkAclAssociationId(a string) 

	// public java.lang.String getNetworkAclAssociationId()
	GetNetworkAclAssociationId() string

	// public com.amazonaws.services.ec2.model.NetworkAclAssociation withNetworkAclAssociationId(java.lang.String)
	WithNetworkAclAssociationId(a string) *ServicesEc2ModelNetworkAclAssociation

	// public void setNetworkAclId(java.lang.String)
	SetNetworkAclId(a string) 

	// public java.lang.String getNetworkAclId()
	GetNetworkAclId() string

	// public com.amazonaws.services.ec2.model.NetworkAclAssociation withNetworkAclId(java.lang.String)
	WithNetworkAclId(a string) *ServicesEc2ModelNetworkAclAssociation

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.NetworkAclAssociation withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelNetworkAclAssociation

	// public com.amazonaws.services.ec2.model.NetworkAclAssociation clone()
	Clone() *ServicesEc2ModelNetworkAclAssociation
}

type ServicesEc2ModelNetworkAclAssociation struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkAclAssociation()
func NewServicesEc2ModelNetworkAclAssociation() (*ServicesEc2ModelNetworkAclAssociation) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkAclAssociation")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkAclAssociation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkAclAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAclAssociation) SetNetworkAclAssociationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkAclAssociationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkAclAssociationId()
func (jbobject *ServicesEc2ModelNetworkAclAssociation) GetNetworkAclAssociationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkAclAssociationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkAclAssociation withNetworkAclAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAclAssociation) WithNetworkAclAssociationId(a string) *ServicesEc2ModelNetworkAclAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAclAssociationId", "com/amazonaws/services/ec2/model/NetworkAclAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkAclAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAclAssociation) SetNetworkAclId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkAclId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkAclId()
func (jbobject *ServicesEc2ModelNetworkAclAssociation) GetNetworkAclId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkAclId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkAclAssociation withNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAclAssociation) WithNetworkAclId(a string) *ServicesEc2ModelNetworkAclAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAclId", "com/amazonaws/services/ec2/model/NetworkAclAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkAclAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAclAssociation) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkAclAssociation) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.NetworkAclAssociation withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAclAssociation) WithSubnetId(a string) *ServicesEc2ModelNetworkAclAssociation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/NetworkAclAssociation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkAclAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkAclAssociation) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkAclAssociation) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkAclAssociation) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkAclAssociation clone()
func (jbobject *ServicesEc2ModelNetworkAclAssociation) Clone() *ServicesEc2ModelNetworkAclAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkAclAssociation")
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
	unique_x := &ServicesEc2ModelNetworkAclAssociation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkAclAssociation) Clone2() (*JavaLangObject, error) {
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


