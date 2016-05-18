package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAttachInternetGatewayRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setInternetGatewayId(java.lang.String)
	SetInternetGatewayId(a string) 

	// public java.lang.String getInternetGatewayId()
	GetInternetGatewayId() string

	// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest withInternetGatewayId(java.lang.String)
	WithInternetGatewayId(a string) *ServicesEc2ModelAttachInternetGatewayRequest

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelAttachInternetGatewayRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AttachInternetGatewayRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest clone()
	Clone3() *ServicesEc2ModelAttachInternetGatewayRequest
}

type ServicesEc2ModelAttachInternetGatewayRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest()
func NewServicesEc2ModelAttachInternetGatewayRequest() (*ServicesEc2ModelAttachInternetGatewayRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AttachInternetGatewayRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAttachInternetGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInternetGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) SetInternetGatewayId(a string)  {
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
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) GetInternetGatewayId() string {
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

// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest withInternetGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) WithInternetGatewayId(a string) *ServicesEc2ModelAttachInternetGatewayRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInternetGatewayId", "com/amazonaws/services/ec2/model/AttachInternetGatewayRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachInternetGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) WithVpcId(a string) *ServicesEc2ModelAttachInternetGatewayRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/AttachInternetGatewayRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachInternetGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AttachInternetGatewayRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AttachInternetGatewayRequest clone()
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) Clone3() *ServicesEc2ModelAttachInternetGatewayRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AttachInternetGatewayRequest")
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
	unique_x := &ServicesEc2ModelAttachInternetGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelAttachInternetGatewayRequest) Clone2() (*JavaLangObject, error) {
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


