package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyInstancePlacementRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelModifyInstancePlacementRequest

	// public void setTenancy(java.lang.String)
	SetTenancy2(a string) 

	// public java.lang.String getTenancy()
	GetTenancy() string

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withTenancy(java.lang.String)
	WithTenancy2(a string) *ServicesEc2ModelModifyInstancePlacementRequest

	// public void setTenancy(com.amazonaws.services.ec2.model.HostTenancy)
	SetTenancy(a ServicesEc2ModelHostTenancyInterface) 

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withTenancy(com.amazonaws.services.ec2.model.HostTenancy)
	WithTenancy(a ServicesEc2ModelHostTenancyInterface) *ServicesEc2ModelModifyInstancePlacementRequest

	// public void setAffinity(java.lang.String)
	SetAffinity2(a string) 

	// public java.lang.String getAffinity()
	GetAffinity() string

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withAffinity(java.lang.String)
	WithAffinity2(a string) *ServicesEc2ModelModifyInstancePlacementRequest

	// public void setAffinity(com.amazonaws.services.ec2.model.Affinity)
	SetAffinity(a ServicesEc2ModelAffinityInterface) 

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withAffinity(com.amazonaws.services.ec2.model.Affinity)
	WithAffinity(a ServicesEc2ModelAffinityInterface) *ServicesEc2ModelModifyInstancePlacementRequest

	// public void setHostId(java.lang.String)
	SetHostId(a string) 

	// public java.lang.String getHostId()
	GetHostId() string

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withHostId(java.lang.String)
	WithHostId(a string) *ServicesEc2ModelModifyInstancePlacementRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest clone()
	Clone3() *ServicesEc2ModelModifyInstancePlacementRequest
}

type ServicesEc2ModelModifyInstancePlacementRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest()
func NewServicesEc2ModelModifyInstancePlacementRequest() (*ServicesEc2ModelModifyInstancePlacementRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) SetInstanceId(a string)  {
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
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) GetInstanceId() string {
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

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) WithInstanceId(a string) *ServicesEc2ModelModifyInstancePlacementRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) SetTenancy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTenancy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getTenancy()
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) GetTenancy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTenancy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) WithTenancy2(a string) *ServicesEc2ModelModifyInstancePlacementRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTenancy", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTenancy(com.amazonaws.services.ec2.model.HostTenancy)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) SetTenancy(a ServicesEc2ModelHostTenancyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTenancy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/HostTenancy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withTenancy(com.amazonaws.services.ec2.model.HostTenancy)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) WithTenancy(a ServicesEc2ModelHostTenancyInterface) *ServicesEc2ModelModifyInstancePlacementRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTenancy", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/HostTenancy"))
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAffinity(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) SetAffinity2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAffinity", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAffinity()
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) GetAffinity() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAffinity", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withAffinity(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) WithAffinity2(a string) *ServicesEc2ModelModifyInstancePlacementRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAffinity", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAffinity(com.amazonaws.services.ec2.model.Affinity)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) SetAffinity(a ServicesEc2ModelAffinityInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAffinity", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Affinity"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withAffinity(com.amazonaws.services.ec2.model.Affinity)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) WithAffinity(a ServicesEc2ModelAffinityInterface) *ServicesEc2ModelModifyInstancePlacementRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAffinity", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Affinity"))
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHostId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) SetHostId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHostId()
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) GetHostId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest withHostId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) WithHostId(a string) *ServicesEc2ModelModifyInstancePlacementRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostId", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyInstancePlacementRequest clone()
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) Clone3() *ServicesEc2ModelModifyInstancePlacementRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyInstancePlacementRequest")
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
	unique_x := &ServicesEc2ModelModifyInstancePlacementRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyInstancePlacementRequest) Clone2() (*JavaLangObject, error) {
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


