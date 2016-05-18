package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateNetworkInterfaceResultInterface interface {
	JavaLangObjectInterface

	// public void setNetworkInterface(com.amazonaws.services.ec2.model.NetworkInterface)
	SetNetworkInterface(a ServicesEc2ModelNetworkInterfaceInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterface getNetworkInterface()
	GetNetworkInterface() *ServicesEc2ModelNetworkInterface

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult withNetworkInterface(com.amazonaws.services.ec2.model.NetworkInterface)
	WithNetworkInterface(a ServicesEc2ModelNetworkInterfaceInterface) *ServicesEc2ModelCreateNetworkInterfaceResult

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult clone()
	Clone() *ServicesEc2ModelCreateNetworkInterfaceResult
}

type ServicesEc2ModelCreateNetworkInterfaceResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult()
func NewServicesEc2ModelCreateNetworkInterfaceResult() (*ServicesEc2ModelCreateNetworkInterfaceResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateNetworkInterfaceResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateNetworkInterfaceResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterface(com.amazonaws.services.ec2.model.NetworkInterface)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) SetNetworkInterface(a ServicesEc2ModelNetworkInterfaceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterface", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterface"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkInterface getNetworkInterface()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) GetNetworkInterface() *ServicesEc2ModelNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterface", "com/amazonaws/services/ec2/model/NetworkInterface")
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
	unique_x := &ServicesEc2ModelNetworkInterface{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult withNetworkInterface(com.amazonaws.services.ec2.model.NetworkInterface)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) WithNetworkInterface(a ServicesEc2ModelNetworkInterfaceInterface) *ServicesEc2ModelCreateNetworkInterfaceResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterface", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterface"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceResult clone()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) Clone() *ServicesEc2ModelCreateNetworkInterfaceResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceResult")
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceResult) Clone2() (*JavaLangObject, error) {
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


