package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyReservedInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelModifyReservedInstancesRequest

	// public java.util.List<java.lang.String> getReservedInstancesIds()
	GetReservedInstancesIds() []string

	// public void setReservedInstancesIds(java.util.Collection<java.lang.String>)
	SetReservedInstancesIds(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withReservedInstancesIds(java.lang.String...)
	WithReservedInstancesIds(a ...string) *ServicesEc2ModelModifyReservedInstancesRequest

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withReservedInstancesIds(java.util.Collection<java.lang.String>)
	WithReservedInstancesIds2(a []string) *ServicesEc2ModelModifyReservedInstancesRequest

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesConfiguration> getTargetConfigurations()
	GetTargetConfigurations() []*ServicesEc2ModelReservedInstancesConfiguration

	// public void setTargetConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesConfiguration>)
	SetTargetConfigurations(a []*ServicesEc2ModelReservedInstancesConfiguration) 

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withTargetConfigurations(com.amazonaws.services.ec2.model.ReservedInstancesConfiguration...)
	WithTargetConfigurations(a ...*ServicesEc2ModelReservedInstancesConfiguration) *ServicesEc2ModelModifyReservedInstancesRequest

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withTargetConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesConfiguration>)
	WithTargetConfigurations2(a []*ServicesEc2ModelReservedInstancesConfiguration) *ServicesEc2ModelModifyReservedInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest clone()
	Clone3() *ServicesEc2ModelModifyReservedInstancesRequest
}

type ServicesEc2ModelModifyReservedInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest()
func NewServicesEc2ModelModifyReservedInstancesRequest() (*ServicesEc2ModelModifyReservedInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) WithClientToken(a string) *ServicesEc2ModelModifyReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getReservedInstancesIds()
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) GetReservedInstancesIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesIds", "java/util/List")
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

// public void setReservedInstancesIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) SetReservedInstancesIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withReservedInstancesIds(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) WithReservedInstancesIds(a ...string) *ServicesEc2ModelModifyReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesIds", "com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withReservedInstancesIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) WithReservedInstancesIds2(a []string) *ServicesEc2ModelModifyReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesIds", "com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesConfiguration> getTargetConfigurations()
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) GetTargetConfigurations() []*ServicesEc2ModelReservedInstancesConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTargetConfigurations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstancesConfiguration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTargetConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesConfiguration>)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) SetTargetConfigurations(a []*ServicesEc2ModelReservedInstancesConfiguration)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetConfigurations", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withTargetConfigurations(com.amazonaws.services.ec2.model.ReservedInstancesConfiguration...)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) WithTargetConfigurations(a ...*ServicesEc2ModelReservedInstancesConfiguration) *ServicesEc2ModelModifyReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesConfiguration")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetConfigurations", "com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesConfiguration"))
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest withTargetConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesConfiguration>)
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) WithTargetConfigurations2(a []*ServicesEc2ModelReservedInstancesConfiguration) *ServicesEc2ModelModifyReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetConfigurations", "com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyReservedInstancesRequest clone()
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) Clone3() *ServicesEc2ModelModifyReservedInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyReservedInstancesRequest")
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
	unique_x := &ServicesEc2ModelModifyReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyReservedInstancesRequest) Clone2() (*JavaLangObject, error) {
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


