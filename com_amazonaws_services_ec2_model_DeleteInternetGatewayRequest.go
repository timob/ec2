package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteInternetGatewayRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setInternetGatewayId(java.lang.String)
	SetInternetGatewayId(a string) 

	// public java.lang.String getInternetGatewayId()
	GetInternetGatewayId() string

	// public com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest withInternetGatewayId(java.lang.String)
	WithInternetGatewayId(a string) *ServicesEc2ModelDeleteInternetGatewayRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest clone()
	Clone3() *ServicesEc2ModelDeleteInternetGatewayRequest
}

type ServicesEc2ModelDeleteInternetGatewayRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest()
func NewServicesEc2ModelDeleteInternetGatewayRequest() (*ServicesEc2ModelDeleteInternetGatewayRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteInternetGatewayRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteInternetGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInternetGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) SetInternetGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInternetGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInternetGatewayId()
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) GetInternetGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInternetGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest withInternetGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) WithInternetGatewayId(a string) *ServicesEc2ModelDeleteInternetGatewayRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInternetGatewayId", "com/amazonaws/services/ec2/model/DeleteInternetGatewayRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteInternetGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteInternetGatewayRequest clone()
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) Clone3() *ServicesEc2ModelDeleteInternetGatewayRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteInternetGatewayRequest")
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
	unique_x := &ServicesEc2ModelDeleteInternetGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDeleteInternetGatewayRequest) Clone2() (*JavaLangObject, error) {
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


