package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpcRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setCidrBlock(java.lang.String)
	SetCidrBlock(a string) 

	// public java.lang.String getCidrBlock()
	GetCidrBlock() string

	// public com.amazonaws.services.ec2.model.CreateVpcRequest withCidrBlock(java.lang.String)
	WithCidrBlock(a string) *ServicesEc2ModelCreateVpcRequest

	// public void setInstanceTenancy(java.lang.String)
	SetInstanceTenancy2(a string) 

	// public java.lang.String getInstanceTenancy()
	GetInstanceTenancy() string

	// public com.amazonaws.services.ec2.model.CreateVpcRequest withInstanceTenancy(java.lang.String)
	WithInstanceTenancy2(a string) *ServicesEc2ModelCreateVpcRequest

	// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	SetInstanceTenancy(a ServicesEc2ModelTenancyInterface) 

	// public com.amazonaws.services.ec2.model.CreateVpcRequest withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelCreateVpcRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpcRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateVpcRequest clone()
	Clone3() *ServicesEc2ModelCreateVpcRequest
}

type ServicesEc2ModelCreateVpcRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateVpcRequest()
func NewServicesEc2ModelCreateVpcRequest() (*ServicesEc2ModelCreateVpcRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpcRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateVpcRequest(java.lang.String)
func NewServicesEc2ModelCreateVpcRequest2(a string) (*ServicesEc2ModelCreateVpcRequest) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcRequest", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelCreateVpcRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcRequest) SetCidrBlock(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCidrBlock", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCidrBlock()
func (jbobject *ServicesEc2ModelCreateVpcRequest) GetCidrBlock() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCidrBlock", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcRequest withCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcRequest) WithCidrBlock(a string) *ServicesEc2ModelCreateVpcRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCidrBlock", "com/amazonaws/services/ec2/model/CreateVpcRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcRequest) SetInstanceTenancy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceTenancy()
func (jbobject *ServicesEc2ModelCreateVpcRequest) GetInstanceTenancy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceTenancy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcRequest withInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcRequest) WithInstanceTenancy2(a string) *ServicesEc2ModelCreateVpcRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/CreateVpcRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelCreateVpcRequest) SetInstanceTenancy(a ServicesEc2ModelTenancyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVpcRequest withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelCreateVpcRequest) WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelCreateVpcRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/CreateVpcRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
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
	unique_x := &ServicesEc2ModelCreateVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpcRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateVpcRequest) GetDryRunRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDryRunRequest", "com/amazonaws/Request")
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
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVpcRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpcRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpcRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpcRequest clone()
func (jbobject *ServicesEc2ModelCreateVpcRequest) Clone3() *ServicesEc2ModelCreateVpcRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpcRequest")
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
	unique_x := &ServicesEc2ModelCreateVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateVpcRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVpcRequest) Clone2() (*JavaLangObject, error) {
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


