package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteNatGatewayResultInterface interface {
	JavaLangObjectInterface

	// public void setNatGatewayId(java.lang.String)
	SetNatGatewayId(a string) 

	// public java.lang.String getNatGatewayId()
	GetNatGatewayId() string

	// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult withNatGatewayId(java.lang.String)
	WithNatGatewayId(a string) *ServicesEc2ModelDeleteNatGatewayResult

	// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult clone()
	Clone() *ServicesEc2ModelDeleteNatGatewayResult
}

type ServicesEc2ModelDeleteNatGatewayResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult()
func NewServicesEc2ModelDeleteNatGatewayResult() (*ServicesEc2ModelDeleteNatGatewayResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteNatGatewayResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteNatGatewayResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNatGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) SetNatGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNatGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNatGatewayId()
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) GetNatGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNatGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult withNatGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) WithNatGatewayId(a string) *ServicesEc2ModelDeleteNatGatewayResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGatewayId", "com/amazonaws/services/ec2/model/DeleteNatGatewayResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteNatGatewayResult clone()
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) Clone() *ServicesEc2ModelDeleteNatGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteNatGatewayResult")
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
	unique_x := &ServicesEc2ModelDeleteNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDeleteNatGatewayResult) Clone2() (*JavaLangObject, error) {
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


