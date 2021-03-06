package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpcResultInterface interface {
	JavaLangObjectInterface

	// public void setVpc(com.amazonaws.services.ec2.model.Vpc)
	SetVpc(a ServicesEc2ModelVpcInterface) 

	// public com.amazonaws.services.ec2.model.Vpc getVpc()
	GetVpc() *ServicesEc2ModelVpc

	// public com.amazonaws.services.ec2.model.CreateVpcResult withVpc(com.amazonaws.services.ec2.model.Vpc)
	WithVpc(a ServicesEc2ModelVpcInterface) *ServicesEc2ModelCreateVpcResult

	// public com.amazonaws.services.ec2.model.CreateVpcResult clone()
	Clone() *ServicesEc2ModelCreateVpcResult
}

type ServicesEc2ModelCreateVpcResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVpcResult()
func NewServicesEc2ModelCreateVpcResult() (*ServicesEc2ModelCreateVpcResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpcResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpc(com.amazonaws.services.ec2.model.Vpc)
func (jbobject *ServicesEc2ModelCreateVpcResult) SetVpc(a ServicesEc2ModelVpcInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpc", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Vpc"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Vpc getVpc()
func (jbobject *ServicesEc2ModelCreateVpcResult) GetVpc() *ServicesEc2ModelVpc {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpc", "com/amazonaws/services/ec2/model/Vpc")
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcResult withVpc(com.amazonaws.services.ec2.model.Vpc)
func (jbobject *ServicesEc2ModelCreateVpcResult) WithVpc(a ServicesEc2ModelVpcInterface) *ServicesEc2ModelCreateVpcResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpc", "com/amazonaws/services/ec2/model/CreateVpcResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Vpc"))
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
	unique_x := &ServicesEc2ModelCreateVpcResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVpcResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpcResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpcResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpcResult clone()
func (jbobject *ServicesEc2ModelCreateVpcResult) Clone() *ServicesEc2ModelCreateVpcResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpcResult")
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
	unique_x := &ServicesEc2ModelCreateVpcResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVpcResult) Clone2() (*JavaLangObject, error) {
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


