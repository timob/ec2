package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyVpcAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelModifyVpcAttributeRequest

	// public void setEnableDnsSupport(java.lang.Boolean)
	SetEnableDnsSupport(a bool) 

	// public java.lang.Boolean getEnableDnsSupport()
	GetEnableDnsSupport() bool

	// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest withEnableDnsSupport(java.lang.Boolean)
	WithEnableDnsSupport(a bool) *ServicesEc2ModelModifyVpcAttributeRequest

	// public java.lang.Boolean isEnableDnsSupport()
	IsEnableDnsSupport() bool

	// public void setEnableDnsHostnames(java.lang.Boolean)
	SetEnableDnsHostnames(a bool) 

	// public java.lang.Boolean getEnableDnsHostnames()
	GetEnableDnsHostnames() bool

	// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest withEnableDnsHostnames(java.lang.Boolean)
	WithEnableDnsHostnames(a bool) *ServicesEc2ModelModifyVpcAttributeRequest

	// public java.lang.Boolean isEnableDnsHostnames()
	IsEnableDnsHostnames() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest clone()
	Clone3() *ServicesEc2ModelModifyVpcAttributeRequest
}

type ServicesEc2ModelModifyVpcAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest()
func NewServicesEc2ModelModifyVpcAttributeRequest() (*ServicesEc2ModelModifyVpcAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyVpcAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyVpcAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) WithVpcId(a string) *ServicesEc2ModelModifyVpcAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/ModifyVpcAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEnableDnsSupport(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) SetEnableDnsSupport(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnableDnsSupport", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEnableDnsSupport()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) GetEnableDnsSupport() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnableDnsSupport", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest withEnableDnsSupport(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) WithEnableDnsSupport(a bool) *ServicesEc2ModelModifyVpcAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEnableDnsSupport", "com/amazonaws/services/ec2/model/ModifyVpcAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEnableDnsSupport()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) IsEnableDnsSupport() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnableDnsSupport", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setEnableDnsHostnames(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) SetEnableDnsHostnames(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnableDnsHostnames", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEnableDnsHostnames()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) GetEnableDnsHostnames() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnableDnsHostnames", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest withEnableDnsHostnames(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) WithEnableDnsHostnames(a bool) *ServicesEc2ModelModifyVpcAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEnableDnsHostnames", "com/amazonaws/services/ec2/model/ModifyVpcAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEnableDnsHostnames()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) IsEnableDnsHostnames() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnableDnsHostnames", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyVpcAttributeRequest clone()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) Clone3() *ServicesEc2ModelModifyVpcAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyVpcAttributeRequest")
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
	unique_x := &ServicesEc2ModelModifyVpcAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyVpcAttributeRequest) Clone2() (*JavaLangObject, error) {
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


