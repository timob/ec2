package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest clone()
	Clone3() *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest
}

type ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest()
func NewServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest() (*ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotDatafeedSubscriptionRequest clone()
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) Clone3() *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotDatafeedSubscriptionRequest")
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
	unique_x := &ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeSpotDatafeedSubscriptionRequest) Clone2() (*JavaLangObject, error) {
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


