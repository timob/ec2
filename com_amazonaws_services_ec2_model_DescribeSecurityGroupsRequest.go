package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSecurityGroupsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getGroupNames()
	GetGroupNames() []string

	// public void setGroupNames(java.util.Collection<java.lang.String>)
	SetGroupNames(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupNames(java.lang.String...)
	WithGroupNames(a ...string) *ServicesEc2ModelDescribeSecurityGroupsRequest

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupNames(java.util.Collection<java.lang.String>)
	WithGroupNames2(a []string) *ServicesEc2ModelDescribeSecurityGroupsRequest

	// public java.util.List<java.lang.String> getGroupIds()
	GetGroupIds() []string

	// public void setGroupIds(java.util.Collection<java.lang.String>)
	SetGroupIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupIds(java.lang.String...)
	WithGroupIds(a ...string) *ServicesEc2ModelDescribeSecurityGroupsRequest

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupIds(java.util.Collection<java.lang.String>)
	WithGroupIds2(a []string) *ServicesEc2ModelDescribeSecurityGroupsRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSecurityGroupsRequest

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSecurityGroupsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest clone()
	Clone3() *ServicesEc2ModelDescribeSecurityGroupsRequest
}

type ServicesEc2ModelDescribeSecurityGroupsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest()
func NewServicesEc2ModelDescribeSecurityGroupsRequest() (*ServicesEc2ModelDescribeSecurityGroupsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getGroupNames()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) GetGroupNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupNames", "java/util/List")
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

// public void setGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) SetGroupNames(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupNames", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupNames(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) WithGroupNames(a ...string) *ServicesEc2ModelDescribeSecurityGroupsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) WithGroupNames2(a []string) *ServicesEc2ModelDescribeSecurityGroupsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroupIds()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) GetGroupIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupIds", "java/util/List")
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

// public void setGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) SetGroupIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) WithGroupIds(a ...string) *ServicesEc2ModelDescribeSecurityGroupsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupIds", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) WithGroupIds2(a []string) *ServicesEc2ModelDescribeSecurityGroupsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupIds", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSecurityGroupsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeSecurityGroupsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsRequest clone()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) Clone3() *ServicesEc2ModelDescribeSecurityGroupsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsRequest")
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsRequest) Clone2() (*JavaLangObject, error) {
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


