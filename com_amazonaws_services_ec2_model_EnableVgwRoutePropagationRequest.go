package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelEnableVgwRoutePropagationRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setRouteTableId(java.lang.String)
	SetRouteTableId(a string) 

	// public java.lang.String getRouteTableId()
	GetRouteTableId() string

	// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest withRouteTableId(java.lang.String)
	WithRouteTableId(a string) *ServicesEc2ModelEnableVgwRoutePropagationRequest

	// public void setGatewayId(java.lang.String)
	SetGatewayId(a string) 

	// public java.lang.String getGatewayId()
	GetGatewayId() string

	// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest withGatewayId(java.lang.String)
	WithGatewayId(a string) *ServicesEc2ModelEnableVgwRoutePropagationRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest clone()
	Clone3() *ServicesEc2ModelEnableVgwRoutePropagationRequest
}

type ServicesEc2ModelEnableVgwRoutePropagationRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest()
func NewServicesEc2ModelEnableVgwRoutePropagationRequest() (*ServicesEc2ModelEnableVgwRoutePropagationRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/EnableVgwRoutePropagationRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelEnableVgwRoutePropagationRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) SetRouteTableId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTableId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRouteTableId()
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) GetRouteTableId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTableId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest withRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) WithRouteTableId(a string) *ServicesEc2ModelEnableVgwRoutePropagationRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableId", "com/amazonaws/services/ec2/model/EnableVgwRoutePropagationRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEnableVgwRoutePropagationRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) SetGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGatewayId()
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) GetGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest withGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) WithGatewayId(a string) *ServicesEc2ModelEnableVgwRoutePropagationRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGatewayId", "com/amazonaws/services/ec2/model/EnableVgwRoutePropagationRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEnableVgwRoutePropagationRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.EnableVgwRoutePropagationRequest clone()
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) Clone3() *ServicesEc2ModelEnableVgwRoutePropagationRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/EnableVgwRoutePropagationRequest")
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
	unique_x := &ServicesEc2ModelEnableVgwRoutePropagationRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelEnableVgwRoutePropagationRequest) Clone2() (*JavaLangObject, error) {
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


