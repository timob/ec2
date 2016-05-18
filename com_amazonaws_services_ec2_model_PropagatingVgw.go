package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPropagatingVgwInterface interface {
	JavaLangObjectInterface

	// public void setGatewayId(java.lang.String)
	SetGatewayId(a string) 

	// public java.lang.String getGatewayId()
	GetGatewayId() string

	// public com.amazonaws.services.ec2.model.PropagatingVgw withGatewayId(java.lang.String)
	WithGatewayId(a string) *ServicesEc2ModelPropagatingVgw

	// public com.amazonaws.services.ec2.model.PropagatingVgw clone()
	Clone() *ServicesEc2ModelPropagatingVgw
}

type ServicesEc2ModelPropagatingVgw struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PropagatingVgw()
func NewServicesEc2ModelPropagatingVgw() (*ServicesEc2ModelPropagatingVgw) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PropagatingVgw")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPropagatingVgw{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelPropagatingVgw) SetGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGatewayId()
func (jbobject *ServicesEc2ModelPropagatingVgw) GetGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PropagatingVgw withGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelPropagatingVgw) WithGatewayId(a string) *ServicesEc2ModelPropagatingVgw {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGatewayId", "com/amazonaws/services/ec2/model/PropagatingVgw", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPropagatingVgw{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPropagatingVgw) ToString() string {
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
func (jbobject *ServicesEc2ModelPropagatingVgw) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPropagatingVgw) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PropagatingVgw clone()
func (jbobject *ServicesEc2ModelPropagatingVgw) Clone() *ServicesEc2ModelPropagatingVgw {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PropagatingVgw")
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
	unique_x := &ServicesEc2ModelPropagatingVgw{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPropagatingVgw) Clone2() (*JavaLangObject, error) {
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


