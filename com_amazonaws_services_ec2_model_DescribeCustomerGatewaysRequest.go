package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeCustomerGatewaysRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getCustomerGatewayIds()
	GetCustomerGatewayIds() []string

	// public void setCustomerGatewayIds(java.util.Collection<java.lang.String>)
	SetCustomerGatewayIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withCustomerGatewayIds(java.lang.String...)
	WithCustomerGatewayIds(a ...string) *ServicesEc2ModelDescribeCustomerGatewaysRequest

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withCustomerGatewayIds(java.util.Collection<java.lang.String>)
	WithCustomerGatewayIds2(a []string) *ServicesEc2ModelDescribeCustomerGatewaysRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeCustomerGatewaysRequest

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeCustomerGatewaysRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest clone()
	Clone3() *ServicesEc2ModelDescribeCustomerGatewaysRequest
}

type ServicesEc2ModelDescribeCustomerGatewaysRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest()
func NewServicesEc2ModelDescribeCustomerGatewaysRequest() (*ServicesEc2ModelDescribeCustomerGatewaysRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeCustomerGatewaysRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getCustomerGatewayIds()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) GetCustomerGatewayIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomerGatewayIds", "java/util/List")
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

// public void setCustomerGatewayIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) SetCustomerGatewayIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCustomerGatewayIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withCustomerGatewayIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) WithCustomerGatewayIds(a ...string) *ServicesEc2ModelDescribeCustomerGatewaysRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGatewayIds", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withCustomerGatewayIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) WithCustomerGatewayIds2(a []string) *ServicesEc2ModelDescribeCustomerGatewaysRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGatewayIds", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeCustomerGatewaysRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeCustomerGatewaysRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysRequest clone()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) Clone3() *ServicesEc2ModelDescribeCustomerGatewaysRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysRequest")
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysRequest) Clone2() (*JavaLangObject, error) {
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


