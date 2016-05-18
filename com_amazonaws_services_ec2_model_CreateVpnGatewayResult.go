package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpnGatewayResultInterface interface {
	JavaLangObjectInterface

	// public void setVpnGateway(com.amazonaws.services.ec2.model.VpnGateway)
	SetVpnGateway(a ServicesEc2ModelVpnGatewayInterface) 

	// public com.amazonaws.services.ec2.model.VpnGateway getVpnGateway()
	GetVpnGateway() *ServicesEc2ModelVpnGateway

	// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult withVpnGateway(com.amazonaws.services.ec2.model.VpnGateway)
	WithVpnGateway(a ServicesEc2ModelVpnGatewayInterface) *ServicesEc2ModelCreateVpnGatewayResult

	// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult clone()
	Clone() *ServicesEc2ModelCreateVpnGatewayResult
}

type ServicesEc2ModelCreateVpnGatewayResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult()
func NewServicesEc2ModelCreateVpnGatewayResult() (*ServicesEc2ModelCreateVpnGatewayResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpnGatewayResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpnGatewayResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnGateway(com.amazonaws.services.ec2.model.VpnGateway)
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) SetVpnGateway(a ServicesEc2ModelVpnGatewayInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnGateway"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnGateway getVpnGateway()
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) GetVpnGateway() *ServicesEc2ModelVpnGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnGateway", "com/amazonaws/services/ec2/model/VpnGateway")
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
	unique_x := &ServicesEc2ModelVpnGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult withVpnGateway(com.amazonaws.services.ec2.model.VpnGateway)
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) WithVpnGateway(a ServicesEc2ModelVpnGatewayInterface) *ServicesEc2ModelCreateVpnGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnGateway", "com/amazonaws/services/ec2/model/CreateVpnGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnGateway"))
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
	unique_x := &ServicesEc2ModelCreateVpnGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpnGatewayResult clone()
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) Clone() *ServicesEc2ModelCreateVpnGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpnGatewayResult")
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
	unique_x := &ServicesEc2ModelCreateVpnGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVpnGatewayResult) Clone2() (*JavaLangObject, error) {
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


