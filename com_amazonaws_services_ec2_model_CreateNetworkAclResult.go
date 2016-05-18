package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateNetworkAclResultInterface interface {
	JavaLangObjectInterface

	// public void setNetworkAcl(com.amazonaws.services.ec2.model.NetworkAcl)
	SetNetworkAcl(a ServicesEc2ModelNetworkAclInterface) 

	// public com.amazonaws.services.ec2.model.NetworkAcl getNetworkAcl()
	GetNetworkAcl() *ServicesEc2ModelNetworkAcl

	// public com.amazonaws.services.ec2.model.CreateNetworkAclResult withNetworkAcl(com.amazonaws.services.ec2.model.NetworkAcl)
	WithNetworkAcl(a ServicesEc2ModelNetworkAclInterface) *ServicesEc2ModelCreateNetworkAclResult

	// public com.amazonaws.services.ec2.model.CreateNetworkAclResult clone()
	Clone() *ServicesEc2ModelCreateNetworkAclResult
}

type ServicesEc2ModelCreateNetworkAclResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclResult()
func NewServicesEc2ModelCreateNetworkAclResult() (*ServicesEc2ModelCreateNetworkAclResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateNetworkAclResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateNetworkAclResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkAcl(com.amazonaws.services.ec2.model.NetworkAcl)
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) SetNetworkAcl(a ServicesEc2ModelNetworkAclInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkAcl", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkAcl"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkAcl getNetworkAcl()
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) GetNetworkAcl() *ServicesEc2ModelNetworkAcl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkAcl", "com/amazonaws/services/ec2/model/NetworkAcl")
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclResult withNetworkAcl(com.amazonaws.services.ec2.model.NetworkAcl)
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) WithNetworkAcl(a ServicesEc2ModelNetworkAclInterface) *ServicesEc2ModelCreateNetworkAclResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAcl", "com/amazonaws/services/ec2/model/CreateNetworkAclResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkAcl"))
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
	unique_x := &ServicesEc2ModelCreateNetworkAclResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateNetworkAclResult clone()
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) Clone() *ServicesEc2ModelCreateNetworkAclResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateNetworkAclResult")
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
	unique_x := &ServicesEc2ModelCreateNetworkAclResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateNetworkAclResult) Clone2() (*JavaLangObject, error) {
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


