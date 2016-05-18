package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcPeeringConnectionsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getVpcPeeringConnectionIds()
	GetVpcPeeringConnectionIds() []string

	// public void setVpcPeeringConnectionIds(java.util.Collection<java.lang.String>)
	SetVpcPeeringConnectionIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withVpcPeeringConnectionIds(java.lang.String...)
	WithVpcPeeringConnectionIds(a ...string) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withVpcPeeringConnectionIds(java.util.Collection<java.lang.String>)
	WithVpcPeeringConnectionIds2(a []string) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest clone()
	Clone3() *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest
}

type ServicesEc2ModelDescribeVpcPeeringConnectionsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest()
func NewServicesEc2ModelDescribeVpcPeeringConnectionsRequest() (*ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcPeeringConnectionsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getVpcPeeringConnectionIds()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) GetVpcPeeringConnectionIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcPeeringConnectionIds", "java/util/List")
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

// public void setVpcPeeringConnectionIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) SetVpcPeeringConnectionIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcPeeringConnectionIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withVpcPeeringConnectionIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) WithVpcPeeringConnectionIds(a ...string) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionIds", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withVpcPeeringConnectionIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) WithVpcPeeringConnectionIds2(a []string) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionIds", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsRequest clone()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) Clone3() *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsRequest")
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsRequest) Clone2() (*JavaLangObject, error) {
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


