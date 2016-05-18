package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpcEndpointResultInterface interface {
	JavaLangObjectInterface

	// public void setVpcEndpoint(com.amazonaws.services.ec2.model.VpcEndpoint)
	SetVpcEndpoint(a ServicesEc2ModelVpcEndpointInterface) 

	// public com.amazonaws.services.ec2.model.VpcEndpoint getVpcEndpoint()
	GetVpcEndpoint() *ServicesEc2ModelVpcEndpoint

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult withVpcEndpoint(com.amazonaws.services.ec2.model.VpcEndpoint)
	WithVpcEndpoint(a ServicesEc2ModelVpcEndpointInterface) *ServicesEc2ModelCreateVpcEndpointResult

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCreateVpcEndpointResult

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult clone()
	Clone() *ServicesEc2ModelCreateVpcEndpointResult
}

type ServicesEc2ModelCreateVpcEndpointResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult()
func NewServicesEc2ModelCreateVpcEndpointResult() (*ServicesEc2ModelCreateVpcEndpointResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcEndpointResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpcEndpointResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcEndpoint(com.amazonaws.services.ec2.model.VpcEndpoint)
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) SetVpcEndpoint(a ServicesEc2ModelVpcEndpointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcEndpoint", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcEndpoint"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcEndpoint getVpcEndpoint()
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) GetVpcEndpoint() *ServicesEc2ModelVpcEndpoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcEndpoint", "com/amazonaws/services/ec2/model/VpcEndpoint")
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
	unique_x := &ServicesEc2ModelVpcEndpoint{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult withVpcEndpoint(com.amazonaws.services.ec2.model.VpcEndpoint)
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) WithVpcEndpoint(a ServicesEc2ModelVpcEndpointInterface) *ServicesEc2ModelCreateVpcEndpointResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcEndpoint", "com/amazonaws/services/ec2/model/CreateVpcEndpointResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcEndpoint"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) WithClientToken(a string) *ServicesEc2ModelCreateVpcEndpointResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CreateVpcEndpointResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointResult clone()
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) Clone() *ServicesEc2ModelCreateVpcEndpointResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpcEndpointResult")
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVpcEndpointResult) Clone2() (*JavaLangObject, error) {
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


