package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeNetworkInterfaceAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute2(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest withAttribute(java.lang.String)
	WithAttribute2(a string) *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest

	// public void setAttribute(com.amazonaws.services.ec2.model.NetworkInterfaceAttribute)
	SetAttribute(a ServicesEc2ModelNetworkInterfaceAttributeInterface) 

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest withAttribute(com.amazonaws.services.ec2.model.NetworkInterfaceAttribute)
	WithAttribute(a ServicesEc2ModelNetworkInterfaceAttributeInterface) *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest clone()
	Clone3() *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest
}

type ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest()
func NewServicesEc2ModelDescribeNetworkInterfaceAttributeRequest() (*ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) SetNetworkInterfaceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterfaceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkInterfaceId()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) GetNetworkInterfaceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) WithNetworkInterfaceId(a string) *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) SetAttribute2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttribute()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) GetAttribute() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttribute", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) WithAttribute2(a string) *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(com.amazonaws.services.ec2.model.NetworkInterfaceAttribute)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) SetAttribute(a ServicesEc2ModelNetworkInterfaceAttributeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest withAttribute(com.amazonaws.services.ec2.model.NetworkInterfaceAttribute)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) WithAttribute(a ServicesEc2ModelNetworkInterfaceAttributeInterface) *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfaceAttributeRequest clone()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) Clone3() *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeNetworkInterfaceAttributeRequest")
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeNetworkInterfaceAttributeRequest) Clone2() (*JavaLangObject, error) {
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


