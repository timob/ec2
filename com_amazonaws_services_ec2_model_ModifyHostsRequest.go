package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyHostsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getHostIds()
	GetHostIds() []string

	// public void setHostIds(java.util.Collection<java.lang.String>)
	SetHostIds(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyHostsRequest withHostIds(java.lang.String...)
	WithHostIds(a ...string) *ServicesEc2ModelModifyHostsRequest

	// public com.amazonaws.services.ec2.model.ModifyHostsRequest withHostIds(java.util.Collection<java.lang.String>)
	WithHostIds2(a []string) *ServicesEc2ModelModifyHostsRequest

	// public void setAutoPlacement(java.lang.String)
	SetAutoPlacement2(a string) 

	// public java.lang.String getAutoPlacement()
	GetAutoPlacement() string

	// public com.amazonaws.services.ec2.model.ModifyHostsRequest withAutoPlacement(java.lang.String)
	WithAutoPlacement2(a string) *ServicesEc2ModelModifyHostsRequest

	// public void setAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
	SetAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) 

	// public com.amazonaws.services.ec2.model.ModifyHostsRequest withAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
	WithAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) *ServicesEc2ModelModifyHostsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyHostsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyHostsRequest clone()
	Clone3() *ServicesEc2ModelModifyHostsRequest
}

type ServicesEc2ModelModifyHostsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyHostsRequest()
func NewServicesEc2ModelModifyHostsRequest() (*ServicesEc2ModelModifyHostsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyHostsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyHostsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getHostIds()
func (jbobject *ServicesEc2ModelModifyHostsRequest) GetHostIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostIds", "java/util/List")
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

// public void setHostIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyHostsRequest) SetHostIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyHostsRequest withHostIds(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyHostsRequest) WithHostIds(a ...string) *ServicesEc2ModelModifyHostsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostIds", "com/amazonaws/services/ec2/model/ModifyHostsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyHostsRequest withHostIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyHostsRequest) WithHostIds2(a []string) *ServicesEc2ModelModifyHostsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostIds", "com/amazonaws/services/ec2/model/ModifyHostsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAutoPlacement(java.lang.String)
func (jbobject *ServicesEc2ModelModifyHostsRequest) SetAutoPlacement2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoPlacement", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAutoPlacement()
func (jbobject *ServicesEc2ModelModifyHostsRequest) GetAutoPlacement() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAutoPlacement", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyHostsRequest withAutoPlacement(java.lang.String)
func (jbobject *ServicesEc2ModelModifyHostsRequest) WithAutoPlacement2(a string) *ServicesEc2ModelModifyHostsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoPlacement", "com/amazonaws/services/ec2/model/ModifyHostsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
func (jbobject *ServicesEc2ModelModifyHostsRequest) SetAutoPlacement(a ServicesEc2ModelAutoPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AutoPlacement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyHostsRequest withAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
func (jbobject *ServicesEc2ModelModifyHostsRequest) WithAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) *ServicesEc2ModelModifyHostsRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoPlacement", "com/amazonaws/services/ec2/model/ModifyHostsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AutoPlacement"))
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
	unique_x := &ServicesEc2ModelModifyHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyHostsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyHostsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyHostsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyHostsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyHostsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyHostsRequest clone()
func (jbobject *ServicesEc2ModelModifyHostsRequest) Clone3() *ServicesEc2ModelModifyHostsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyHostsRequest")
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
	unique_x := &ServicesEc2ModelModifyHostsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyHostsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyHostsRequest) Clone2() (*JavaLangObject, error) {
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


