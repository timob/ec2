package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpcPeeringConnectionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelCreateVpcPeeringConnectionRequest

	// public void setPeerVpcId(java.lang.String)
	SetPeerVpcId(a string) 

	// public java.lang.String getPeerVpcId()
	GetPeerVpcId() string

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest withPeerVpcId(java.lang.String)
	WithPeerVpcId(a string) *ServicesEc2ModelCreateVpcPeeringConnectionRequest

	// public void setPeerOwnerId(java.lang.String)
	SetPeerOwnerId(a string) 

	// public java.lang.String getPeerOwnerId()
	GetPeerOwnerId() string

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest withPeerOwnerId(java.lang.String)
	WithPeerOwnerId(a string) *ServicesEc2ModelCreateVpcPeeringConnectionRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest clone()
	Clone3() *ServicesEc2ModelCreateVpcPeeringConnectionRequest
}

type ServicesEc2ModelCreateVpcPeeringConnectionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest()
func NewServicesEc2ModelCreateVpcPeeringConnectionRequest() (*ServicesEc2ModelCreateVpcPeeringConnectionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpcPeeringConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) WithVpcId(a string) *ServicesEc2ModelCreateVpcPeeringConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPeerVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) SetPeerVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPeerVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPeerVpcId()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) GetPeerVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPeerVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest withPeerVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) WithPeerVpcId(a string) *ServicesEc2ModelCreateVpcPeeringConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPeerVpcId", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPeerOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) SetPeerOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPeerOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPeerOwnerId()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) GetPeerOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPeerOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest withPeerOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) WithPeerOwnerId(a string) *ServicesEc2ModelCreateVpcPeeringConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPeerOwnerId", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionRequest clone()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) Clone3() *ServicesEc2ModelCreateVpcPeeringConnectionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionRequest")
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
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionRequest) Clone2() (*JavaLangObject, error) {
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


