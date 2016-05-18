package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeReservedInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getReservedInstancesIds()
	GetReservedInstancesIds() []string

	// public void setReservedInstancesIds(java.util.Collection<java.lang.String>)
	SetReservedInstancesIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withReservedInstancesIds(java.lang.String...)
	WithReservedInstancesIds(a ...string) *ServicesEc2ModelDescribeReservedInstancesRequest

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withReservedInstancesIds(java.util.Collection<java.lang.String>)
	WithReservedInstancesIds2(a []string) *ServicesEc2ModelDescribeReservedInstancesRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
	GetFilters() []*ServicesEc2ModelFilter

	// public void setFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	SetFilters(a []*ServicesEc2ModelFilter) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
	WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesRequest

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
	WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesRequest

	// public void setOfferingType(java.lang.String)
	SetOfferingType2(a string) 

	// public java.lang.String getOfferingType()
	GetOfferingType() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withOfferingType(java.lang.String)
	WithOfferingType2(a string) *ServicesEc2ModelDescribeReservedInstancesRequest

	// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelDescribeReservedInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest clone()
	Clone3() *ServicesEc2ModelDescribeReservedInstancesRequest
}

type ServicesEc2ModelDescribeReservedInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest()
func NewServicesEc2ModelDescribeReservedInstancesRequest() (*ServicesEc2ModelDescribeReservedInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getReservedInstancesIds()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) GetReservedInstancesIds() []string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) SetReservedInstancesIds(a []string)  {
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withReservedInstancesIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) WithReservedInstancesIds(a ...string) *ServicesEc2ModelDescribeReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesIds", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withReservedInstancesIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) WithReservedInstancesIds2(a []string) *ServicesEc2ModelDescribeReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesIds", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Filter> getFilters()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) GetFilters() []*ServicesEc2ModelFilter {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) SetFilters(a []*ServicesEc2ModelFilter)  {
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withFilters(com.amazonaws.services.ec2.model.Filter...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) WithFilters(a ...*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Filter")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Filter"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withFilters(java.util.Collection<com.amazonaws.services.ec2.model.Filter>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) WithFilters2(a []*ServicesEc2ModelFilter) *ServicesEc2ModelDescribeReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFilters", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) SetOfferingType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOfferingType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOfferingType()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) GetOfferingType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOfferingType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) WithOfferingType2(a string) *ServicesEc2ModelDescribeReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOfferingType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelDescribeReservedInstancesRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesRequest clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) Clone3() *ServicesEc2ModelDescribeReservedInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesRequest")
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesRequest) Clone2() (*JavaLangObject, error) {
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


