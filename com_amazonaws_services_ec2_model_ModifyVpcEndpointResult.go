package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyVpcEndpointResultInterface interface {
	JavaLangObjectInterface

	// public void setReturn(java.lang.Boolean)
	SetReturn(a bool) 

	// public java.lang.Boolean getReturn()
	GetReturn() bool

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult withReturn(java.lang.Boolean)
	WithReturn(a bool) *ServicesEc2ModelModifyVpcEndpointResult

	// public java.lang.Boolean isReturn()
	IsReturn() bool

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult clone()
	Clone() *ServicesEc2ModelModifyVpcEndpointResult
}

type ServicesEc2ModelModifyVpcEndpointResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult()
func NewServicesEc2ModelModifyVpcEndpointResult() (*ServicesEc2ModelModifyVpcEndpointResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyVpcEndpointResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyVpcEndpointResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReturn(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) SetReturn(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReturn", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getReturn()
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) GetReturn() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReturn", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult withReturn(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) WithReturn(a bool) *ServicesEc2ModelModifyVpcEndpointResult {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReturn", "com/amazonaws/services/ec2/model/ModifyVpcEndpointResult", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isReturn()
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) IsReturn() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isReturn", "java/lang/Boolean")
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointResult clone()
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) Clone() *ServicesEc2ModelModifyVpcEndpointResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyVpcEndpointResult")
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelModifyVpcEndpointResult) Clone2() (*JavaLangObject, error) {
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


