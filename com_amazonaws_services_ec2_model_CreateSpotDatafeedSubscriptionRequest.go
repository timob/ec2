package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateSpotDatafeedSubscriptionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setBucket(java.lang.String)
	SetBucket(a string) 

	// public java.lang.String getBucket()
	GetBucket() string

	// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest withBucket(java.lang.String)
	WithBucket(a string) *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest

	// public void setPrefix(java.lang.String)
	SetPrefix(a string) 

	// public java.lang.String getPrefix()
	GetPrefix() string

	// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest withPrefix(java.lang.String)
	WithPrefix(a string) *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest clone()
	Clone3() *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest
}

type ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest()
func NewServicesEc2ModelCreateSpotDatafeedSubscriptionRequest() (*ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest(java.lang.String)
func NewServicesEc2ModelCreateSpotDatafeedSubscriptionRequest2(a string) (*ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionRequest", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setBucket(java.lang.String)
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) SetBucket(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBucket", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getBucket()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) GetBucket() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBucket", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest withBucket(java.lang.String)
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) WithBucket(a string) *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBucket", "com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrefix(java.lang.String)
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) SetPrefix(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefix", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrefix()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) GetPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefix", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest withPrefix(java.lang.String)
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) WithPrefix(a string) *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefix", "com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateSpotDatafeedSubscriptionRequest clone()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) Clone3() *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateSpotDatafeedSubscriptionRequest")
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
	unique_x := &ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateSpotDatafeedSubscriptionRequest) Clone2() (*JavaLangObject, error) {
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


