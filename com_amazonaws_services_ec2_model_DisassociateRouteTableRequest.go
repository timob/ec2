package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDisassociateRouteTableRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setAssociationId(java.lang.String)
	SetAssociationId(a string) 

	// public java.lang.String getAssociationId()
	GetAssociationId() string

	// public com.amazonaws.services.ec2.model.DisassociateRouteTableRequest withAssociationId(java.lang.String)
	WithAssociationId(a string) *ServicesEc2ModelDisassociateRouteTableRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DisassociateRouteTableRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DisassociateRouteTableRequest clone()
	Clone3() *ServicesEc2ModelDisassociateRouteTableRequest
}

type ServicesEc2ModelDisassociateRouteTableRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DisassociateRouteTableRequest()
func NewServicesEc2ModelDisassociateRouteTableRequest() (*ServicesEc2ModelDisassociateRouteTableRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DisassociateRouteTableRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDisassociateRouteTableRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) SetAssociationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAssociationId()
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) GetAssociationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DisassociateRouteTableRequest withAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) WithAssociationId(a string) *ServicesEc2ModelDisassociateRouteTableRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociationId", "com/amazonaws/services/ec2/model/DisassociateRouteTableRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDisassociateRouteTableRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DisassociateRouteTableRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DisassociateRouteTableRequest clone()
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) Clone3() *ServicesEc2ModelDisassociateRouteTableRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DisassociateRouteTableRequest")
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
	unique_x := &ServicesEc2ModelDisassociateRouteTableRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDisassociateRouteTableRequest) Clone2() (*JavaLangObject, error) {
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


