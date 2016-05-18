package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelResetImageAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelResetImageAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute2(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest withAttribute(java.lang.String)
	WithAttribute2(a string) *ServicesEc2ModelResetImageAttributeRequest

	// public void setAttribute(com.amazonaws.services.ec2.model.ResetImageAttributeName)
	SetAttribute(a ServicesEc2ModelResetImageAttributeNameInterface) 

	// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest withAttribute(com.amazonaws.services.ec2.model.ResetImageAttributeName)
	WithAttribute(a ServicesEc2ModelResetImageAttributeNameInterface) *ServicesEc2ModelResetImageAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ResetImageAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest clone()
	Clone3() *ServicesEc2ModelResetImageAttributeRequest
}

type ServicesEc2ModelResetImageAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest()
func NewServicesEc2ModelResetImageAttributeRequest() (*ServicesEc2ModelResetImageAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ResetImageAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelResetImageAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest(java.lang.String, java.lang.String)
func NewServicesEc2ModelResetImageAttributeRequest3(a string, b string) (*ServicesEc2ModelResetImageAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ResetImageAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelResetImageAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest(java.lang.String, com.amazonaws.services.ec2.model.ResetImageAttributeName)
func NewServicesEc2ModelResetImageAttributeRequest2(a string, b ServicesEc2ModelResetImageAttributeNameInterface) (*ServicesEc2ModelResetImageAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ResetImageAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/services/ec2/model/ResetImageAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelResetImageAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) SetImageId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImageId()
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) GetImageId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) WithImageId(a string) *ServicesEc2ModelResetImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/ResetImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelResetImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) SetAttribute2(a string)  {
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
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) GetAttribute() string {
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

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) WithAttribute2(a string) *ServicesEc2ModelResetImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ResetImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelResetImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(com.amazonaws.services.ec2.model.ResetImageAttributeName)
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) SetAttribute(a ServicesEc2ModelResetImageAttributeNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResetImageAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest withAttribute(com.amazonaws.services.ec2.model.ResetImageAttributeName)
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) WithAttribute(a ServicesEc2ModelResetImageAttributeNameInterface) *ServicesEc2ModelResetImageAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ResetImageAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResetImageAttributeName"))
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
	unique_x := &ServicesEc2ModelResetImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ResetImageAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ResetImageAttributeRequest clone()
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) Clone3() *ServicesEc2ModelResetImageAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ResetImageAttributeRequest")
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
	unique_x := &ServicesEc2ModelResetImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelResetImageAttributeRequest) Clone2() (*JavaLangObject, error) {
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


