package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpnConnectionResultInterface interface {
	JavaLangObjectInterface

	// public void setVpnConnection(com.amazonaws.services.ec2.model.VpnConnection)
	SetVpnConnection(a ServicesEc2ModelVpnConnectionInterface) 

	// public com.amazonaws.services.ec2.model.VpnConnection getVpnConnection()
	GetVpnConnection() *ServicesEc2ModelVpnConnection

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult withVpnConnection(com.amazonaws.services.ec2.model.VpnConnection)
	WithVpnConnection(a ServicesEc2ModelVpnConnectionInterface) *ServicesEc2ModelCreateVpnConnectionResult

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult clone()
	Clone() *ServicesEc2ModelCreateVpnConnectionResult
}

type ServicesEc2ModelCreateVpnConnectionResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult()
func NewServicesEc2ModelCreateVpnConnectionResult() (*ServicesEc2ModelCreateVpnConnectionResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpnConnectionResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpnConnectionResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnConnection(com.amazonaws.services.ec2.model.VpnConnection)
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) SetVpnConnection(a ServicesEc2ModelVpnConnectionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnConnection", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnection getVpnConnection()
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) GetVpnConnection() *ServicesEc2ModelVpnConnection {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnConnection", "com/amazonaws/services/ec2/model/VpnConnection")
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult withVpnConnection(com.amazonaws.services.ec2.model.VpnConnection)
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) WithVpnConnection(a ServicesEc2ModelVpnConnectionInterface) *ServicesEc2ModelCreateVpnConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnection", "com/amazonaws/services/ec2/model/CreateVpnConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnection"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionResult clone()
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) Clone() *ServicesEc2ModelCreateVpnConnectionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpnConnectionResult")
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVpnConnectionResult) Clone2() (*JavaLangObject, error) {
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


