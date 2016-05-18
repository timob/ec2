package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSnapshotAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelDescribeSnapshotAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute2(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest withAttribute(java.lang.String)
	WithAttribute2(a string) *ServicesEc2ModelDescribeSnapshotAttributeRequest

	// public void setAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
	SetAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest withAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
	WithAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface) *ServicesEc2ModelDescribeSnapshotAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest clone()
	Clone3() *ServicesEc2ModelDescribeSnapshotAttributeRequest
}

type ServicesEc2ModelDescribeSnapshotAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest()
func NewServicesEc2ModelDescribeSnapshotAttributeRequest() (*ServicesEc2ModelDescribeSnapshotAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest(java.lang.String, java.lang.String)
func NewServicesEc2ModelDescribeSnapshotAttributeRequest3(a string, b string) (*ServicesEc2ModelDescribeSnapshotAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest(java.lang.String, com.amazonaws.services.ec2.model.SnapshotAttributeName)
func NewServicesEc2ModelDescribeSnapshotAttributeRequest2(a string, b ServicesEc2ModelSnapshotAttributeNameInterface) (*ServicesEc2ModelDescribeSnapshotAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/services/ec2/model/SnapshotAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) SetSnapshotId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSnapshotId()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) GetSnapshotId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) WithSnapshotId(a string) *ServicesEc2ModelDescribeSnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) SetAttribute2(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) GetAttribute() string {
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) WithAttribute2(a string) *ServicesEc2ModelDescribeSnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) SetAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest withAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) WithAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface) *ServicesEc2ModelDescribeSnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotAttributeName"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeRequest clone()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) Clone3() *ServicesEc2ModelDescribeSnapshotAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeRequest")
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeRequest) Clone2() (*JavaLangObject, error) {
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


