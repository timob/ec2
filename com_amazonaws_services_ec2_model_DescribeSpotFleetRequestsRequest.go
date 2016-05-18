package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotFleetRequestsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getSpotFleetRequestIds()
	GetSpotFleetRequestIds() []string

	// public void setSpotFleetRequestIds(java.util.Collection<java.lang.String>)
	SetSpotFleetRequestIds(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withSpotFleetRequestIds(java.lang.String...)
	WithSpotFleetRequestIds(a ...string) *ServicesEc2ModelDescribeSpotFleetRequestsRequest

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withSpotFleetRequestIds(java.util.Collection<java.lang.String>)
	WithSpotFleetRequestIds2(a []string) *ServicesEc2ModelDescribeSpotFleetRequestsRequest

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetRequestsRequest

	// public void setMaxResults(java.lang.Integer)
	SetMaxResults(a int) 

	// public java.lang.Integer getMaxResults()
	GetMaxResults() int

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withMaxResults(java.lang.Integer)
	WithMaxResults(a int) *ServicesEc2ModelDescribeSpotFleetRequestsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest clone()
	Clone3() *ServicesEc2ModelDescribeSpotFleetRequestsRequest
}

type ServicesEc2ModelDescribeSpotFleetRequestsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest()
func NewServicesEc2ModelDescribeSpotFleetRequestsRequest() (*ServicesEc2ModelDescribeSpotFleetRequestsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotFleetRequestsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getSpotFleetRequestIds()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) GetSpotFleetRequestIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestIds", "java/util/List")
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

// public void setSpotFleetRequestIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) SetSpotFleetRequestIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withSpotFleetRequestIds(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) WithSpotFleetRequestIds(a ...string) *ServicesEc2ModelDescribeSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestIds", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withSpotFleetRequestIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) WithSpotFleetRequestIds2(a []string) *ServicesEc2ModelDescribeSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestIds", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) SetMaxResults(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxResults", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMaxResults()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) GetMaxResults() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxResults", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest withMaxResults(java.lang.Integer)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) WithMaxResults(a int) *ServicesEc2ModelDescribeSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxResults", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsRequest clone()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) Clone3() *ServicesEc2ModelDescribeSpotFleetRequestsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsRequest")
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsRequest) Clone2() (*JavaLangObject, error) {
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


