package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyReservedInstancesResultInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesModificationId(java.lang.String)
	SetReservedInstancesModificationId(a string) 

	// public java.lang.String getReservedInstancesModificationId()
	GetReservedInstancesModificationId() string

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult withReservedInstancesModificationId(java.lang.String)
	WithReservedInstancesModificationId(a string) *ServicesEc2ModelModifyReservedInstancesResult

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult clone()
	Clone() *ServicesEc2ModelModifyReservedInstancesResult
}

type ServicesEc2ModelModifyReservedInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult()
func NewServicesEc2ModelModifyReservedInstancesResult() (*ServicesEc2ModelModifyReservedInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyReservedInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyReservedInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesModificationId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) SetReservedInstancesModificationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesModificationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesModificationId()
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) GetReservedInstancesModificationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesModificationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult withReservedInstancesModificationId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) WithReservedInstancesModificationId(a string) *ServicesEc2ModelModifyReservedInstancesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesModificationId", "com/amazonaws/services/ec2/model/ModifyReservedInstancesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesResult clone()
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) Clone() *ServicesEc2ModelModifyReservedInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyReservedInstancesResult")
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelModifyReservedInstancesResult) Clone2() (*JavaLangObject, error) {
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


