package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAttachClassicLinkVpcRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelAttachClassicLinkVpcRequest

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelAttachClassicLinkVpcRequest

	// public java.util.List<java.lang.String> getGroups()
	GetGroups() []string

	// public void setGroups(java.util.Collection<java.lang.String>)
	SetGroups(a []string) 

	// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withGroups(java.lang.String...)
	WithGroups(a ...string) *ServicesEc2ModelAttachClassicLinkVpcRequest

	// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withGroups(java.util.Collection<java.lang.String>)
	WithGroups2(a []string) *ServicesEc2ModelAttachClassicLinkVpcRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest clone()
	Clone3() *ServicesEc2ModelAttachClassicLinkVpcRequest
}

type ServicesEc2ModelAttachClassicLinkVpcRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest()
func NewServicesEc2ModelAttachClassicLinkVpcRequest() (*ServicesEc2ModelAttachClassicLinkVpcRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAttachClassicLinkVpcRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) WithInstanceId(a string) *ServicesEc2ModelAttachClassicLinkVpcRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachClassicLinkVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) WithVpcId(a string) *ServicesEc2ModelAttachClassicLinkVpcRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachClassicLinkVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroups()
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) GetGroups() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroups", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) SetGroups(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) WithGroups(a ...string) *ServicesEc2ModelAttachClassicLinkVpcRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachClassicLinkVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest withGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) WithGroups2(a []string) *ServicesEc2ModelAttachClassicLinkVpcRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelAttachClassicLinkVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AttachClassicLinkVpcRequest clone()
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) Clone3() *ServicesEc2ModelAttachClassicLinkVpcRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AttachClassicLinkVpcRequest")
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
	unique_x := &ServicesEc2ModelAttachClassicLinkVpcRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelAttachClassicLinkVpcRequest) Clone2() (*JavaLangObject, error) {
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


