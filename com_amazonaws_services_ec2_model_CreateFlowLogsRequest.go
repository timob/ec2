package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateFlowLogsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getResourceIds()
	GetResourceIds() []string

	// public void setResourceIds(java.util.Collection<java.lang.String>)
	SetResourceIds(a []string) 

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceIds(java.lang.String...)
	WithResourceIds(a ...string) *ServicesEc2ModelCreateFlowLogsRequest

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceIds(java.util.Collection<java.lang.String>)
	WithResourceIds2(a []string) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setResourceType(java.lang.String)
	SetResourceType2(a string) 

	// public java.lang.String getResourceType()
	GetResourceType() string

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceType(java.lang.String)
	WithResourceType2(a string) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setResourceType(com.amazonaws.services.ec2.model.FlowLogsResourceType)
	SetResourceType(a ServicesEc2ModelFlowLogsResourceTypeInterface) 

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceType(com.amazonaws.services.ec2.model.FlowLogsResourceType)
	WithResourceType(a ServicesEc2ModelFlowLogsResourceTypeInterface) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setTrafficType(java.lang.String)
	SetTrafficType2(a string) 

	// public java.lang.String getTrafficType()
	GetTrafficType() string

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withTrafficType(java.lang.String)
	WithTrafficType2(a string) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setTrafficType(com.amazonaws.services.ec2.model.TrafficType)
	SetTrafficType(a ServicesEc2ModelTrafficTypeInterface) 

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withTrafficType(com.amazonaws.services.ec2.model.TrafficType)
	WithTrafficType(a ServicesEc2ModelTrafficTypeInterface) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setLogGroupName(java.lang.String)
	SetLogGroupName(a string) 

	// public java.lang.String getLogGroupName()
	GetLogGroupName() string

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withLogGroupName(java.lang.String)
	WithLogGroupName(a string) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setDeliverLogsPermissionArn(java.lang.String)
	SetDeliverLogsPermissionArn(a string) 

	// public java.lang.String getDeliverLogsPermissionArn()
	GetDeliverLogsPermissionArn() string

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withDeliverLogsPermissionArn(java.lang.String)
	WithDeliverLogsPermissionArn(a string) *ServicesEc2ModelCreateFlowLogsRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCreateFlowLogsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateFlowLogsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest clone()
	Clone3() *ServicesEc2ModelCreateFlowLogsRequest
}

type ServicesEc2ModelCreateFlowLogsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest()
func NewServicesEc2ModelCreateFlowLogsRequest() (*ServicesEc2ModelCreateFlowLogsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateFlowLogsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateFlowLogsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getResourceIds()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetResourceIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResourceIds", "java/util/List")
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

// public void setResourceIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetResourceIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceIds(java.lang.String...)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithResourceIds(a ...string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceIds", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithResourceIds2(a []string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceIds", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setResourceType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetResourceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getResourceType()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetResourceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResourceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithResourceType2(a string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceType", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setResourceType(com.amazonaws.services.ec2.model.FlowLogsResourceType)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetResourceType(a ServicesEc2ModelFlowLogsResourceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/FlowLogsResourceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withResourceType(com.amazonaws.services.ec2.model.FlowLogsResourceType)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithResourceType(a ServicesEc2ModelFlowLogsResourceTypeInterface) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceType", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/FlowLogsResourceType"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTrafficType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetTrafficType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTrafficType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getTrafficType()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetTrafficType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTrafficType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withTrafficType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithTrafficType2(a string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTrafficType", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTrafficType(com.amazonaws.services.ec2.model.TrafficType)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetTrafficType(a ServicesEc2ModelTrafficTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTrafficType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/TrafficType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withTrafficType(com.amazonaws.services.ec2.model.TrafficType)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithTrafficType(a ServicesEc2ModelTrafficTypeInterface) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTrafficType", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/TrafficType"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLogGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetLogGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLogGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getLogGroupName()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetLogGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLogGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withLogGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithLogGroupName(a string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLogGroupName", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeliverLogsPermissionArn(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetDeliverLogsPermissionArn(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeliverLogsPermissionArn", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDeliverLogsPermissionArn()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetDeliverLogsPermissionArn() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeliverLogsPermissionArn", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withDeliverLogsPermissionArn(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithDeliverLogsPermissionArn(a string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeliverLogsPermissionArn", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) WithClientToken(a string) *ServicesEc2ModelCreateFlowLogsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateFlowLogsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateFlowLogsRequest clone()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) Clone3() *ServicesEc2ModelCreateFlowLogsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateFlowLogsRequest")
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
	unique_x := &ServicesEc2ModelCreateFlowLogsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateFlowLogsRequest) Clone2() (*JavaLangObject, error) {
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


