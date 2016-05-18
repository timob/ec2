package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDetachNetworkInterfaceRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setAttachmentId(java.lang.String)
	SetAttachmentId(a string) 

	// public java.lang.String getAttachmentId()
	GetAttachmentId() string

	// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest withAttachmentId(java.lang.String)
	WithAttachmentId(a string) *ServicesEc2ModelDetachNetworkInterfaceRequest

	// public void setForce(java.lang.Boolean)
	SetForce(a bool) 

	// public java.lang.Boolean getForce()
	GetForce() bool

	// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest withForce(java.lang.Boolean)
	WithForce(a bool) *ServicesEc2ModelDetachNetworkInterfaceRequest

	// public java.lang.Boolean isForce()
	IsForce() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest clone()
	Clone3() *ServicesEc2ModelDetachNetworkInterfaceRequest
}

type ServicesEc2ModelDetachNetworkInterfaceRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest()
func NewServicesEc2ModelDetachNetworkInterfaceRequest() (*ServicesEc2ModelDetachNetworkInterfaceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DetachNetworkInterfaceRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDetachNetworkInterfaceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) SetAttachmentId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachmentId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttachmentId()
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) GetAttachmentId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachmentId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest withAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) WithAttachmentId(a string) *ServicesEc2ModelDetachNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachmentId", "com/amazonaws/services/ec2/model/DetachNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDetachNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setForce(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) SetForce(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setForce", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getForce()
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) GetForce() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getForce", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest withForce(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) WithForce(a bool) *ServicesEc2ModelDetachNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withForce", "com/amazonaws/services/ec2/model/DetachNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelDetachNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isForce()
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) IsForce() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isForce", "java/lang/Boolean")
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

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DetachNetworkInterfaceRequest clone()
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) Clone3() *ServicesEc2ModelDetachNetworkInterfaceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DetachNetworkInterfaceRequest")
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
	unique_x := &ServicesEc2ModelDetachNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDetachNetworkInterfaceRequest) Clone2() (*JavaLangObject, error) {
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


