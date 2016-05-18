package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelDescribeVpcAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute2(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest withAttribute(java.lang.String)
	WithAttribute2(a string) *ServicesEc2ModelDescribeVpcAttributeRequest

	// public void setAttribute(com.amazonaws.services.ec2.model.VpcAttributeName)
	SetAttribute(a ServicesEc2ModelVpcAttributeNameInterface) 

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest withAttribute(com.amazonaws.services.ec2.model.VpcAttributeName)
	WithAttribute(a ServicesEc2ModelVpcAttributeNameInterface) *ServicesEc2ModelDescribeVpcAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest clone()
	Clone3() *ServicesEc2ModelDescribeVpcAttributeRequest
}

type ServicesEc2ModelDescribeVpcAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest()
func NewServicesEc2ModelDescribeVpcAttributeRequest() (*ServicesEc2ModelDescribeVpcAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) WithVpcId(a string) *ServicesEc2ModelDescribeVpcAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/DescribeVpcAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) SetAttribute2(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) GetAttribute() string {
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

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) WithAttribute2(a string) *ServicesEc2ModelDescribeVpcAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/DescribeVpcAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(com.amazonaws.services.ec2.model.VpcAttributeName)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) SetAttribute(a ServicesEc2ModelVpcAttributeNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest withAttribute(com.amazonaws.services.ec2.model.VpcAttributeName)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) WithAttribute(a ServicesEc2ModelVpcAttributeNameInterface) *ServicesEc2ModelDescribeVpcAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/DescribeVpcAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcAttributeName"))
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeRequest clone()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) Clone3() *ServicesEc2ModelDescribeVpcAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcAttributeRequest")
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeRequest) Clone2() (*JavaLangObject, error) {
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


