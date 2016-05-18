package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeReservedInstancesModificationsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getReservedInstancesModificationIds()
	GetReservedInstancesModificationIds() []string

	// public void setReservedInstancesModificationIds(java.util.Collection<java.lang.String>)
	SetReservedInstancesModificationIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withReservedInstancesModificationIds(java.lang.String...)
	WithReservedInstancesModificationIds(a ...string) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withReservedInstancesModificationIds(java.util.Collection<java.lang.String>)
	WithReservedInstancesModificationIds2(a []string) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest clone()
	Clone3() *ServicesEc2ModelDescribeReservedInstancesModificationsRequest
}

type ServicesEc2ModelDescribeReservedInstancesModificationsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest()
func NewServicesEc2ModelDescribeReservedInstancesModificationsRequest() (*ServicesEc2ModelDescribeReservedInstancesModificationsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getReservedInstancesModificationIds()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) GetReservedInstancesModificationIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesModificationIds", "java/util/List")
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

// public void setReservedInstancesModificationIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) SetReservedInstancesModificationIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesModificationIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withReservedInstancesModificationIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) WithReservedInstancesModificationIds(a ...string) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesModificationIds", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withReservedInstancesModificationIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) WithReservedInstancesModificationIds2(a []string) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesModificationIds", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) SetNextToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNextToken()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) GetNextToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) GetFilters() []*ServicesEc2ModelFilter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFilters", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelFilter)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFilters", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesModificationsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsRequest clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) Clone3() *ServicesEc2ModelDescribeReservedInstancesModificationsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsRequest")
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsRequest) Clone2() (*JavaLangObject, error) {
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


