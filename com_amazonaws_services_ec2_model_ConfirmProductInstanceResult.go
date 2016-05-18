package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelConfirmProductInstanceResultInterface interface {
	JavaLangObjectInterface

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelConfirmProductInstanceResult

	// public void setReturn(java.lang.Boolean)
	SetReturn(a bool) 

	// public java.lang.Boolean getReturn()
	GetReturn() bool

	// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult withReturn(java.lang.Boolean)
	WithReturn(a bool) *ServicesEc2ModelConfirmProductInstanceResult

	// public java.lang.Boolean isReturn()
	IsReturn() bool

	// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult clone()
	Clone() *ServicesEc2ModelConfirmProductInstanceResult
}

type ServicesEc2ModelConfirmProductInstanceResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult()
func NewServicesEc2ModelConfirmProductInstanceResult() (*ServicesEc2ModelConfirmProductInstanceResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ConfirmProductInstanceResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelConfirmProductInstanceResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) SetOwnerId(a string)  {
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
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) GetOwnerId() string {
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

// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) WithOwnerId(a string) *ServicesEc2ModelConfirmProductInstanceResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/ConfirmProductInstanceResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConfirmProductInstanceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setReturn(java.lang.Boolean)
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) SetReturn(a bool)  {
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
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) GetReturn() bool {
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

// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult withReturn(java.lang.Boolean)
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) WithReturn(a bool) *ServicesEc2ModelConfirmProductInstanceResult {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReturn", "com/amazonaws/services/ec2/model/ConfirmProductInstanceResult", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelConfirmProductInstanceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isReturn()
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) IsReturn() bool {
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
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) ToString() string {
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
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ConfirmProductInstanceResult clone()
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) Clone() *ServicesEc2ModelConfirmProductInstanceResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ConfirmProductInstanceResult")
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
	unique_x := &ServicesEc2ModelConfirmProductInstanceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelConfirmProductInstanceResult) Clone2() (*JavaLangObject, error) {
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


