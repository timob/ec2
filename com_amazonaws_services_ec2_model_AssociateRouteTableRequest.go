package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAssociateRouteTableRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelAssociateRouteTableRequest

	// public void setRouteTableId(java.lang.String)
	SetRouteTableId(a string) 

	// public java.lang.String getRouteTableId()
	GetRouteTableId() string

	// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest withRouteTableId(java.lang.String)
	WithRouteTableId(a string) *ServicesEc2ModelAssociateRouteTableRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AssociateRouteTableRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest clone()
	Clone3() *ServicesEc2ModelAssociateRouteTableRequest
}

type ServicesEc2ModelAssociateRouteTableRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest()
func NewServicesEc2ModelAssociateRouteTableRequest() (*ServicesEc2ModelAssociateRouteTableRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AssociateRouteTableRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAssociateRouteTableRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) SetSubnetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnetId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSubnetId()
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) GetSubnetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnetId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) WithSubnetId(a string) *ServicesEc2ModelAssociateRouteTableRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/AssociateRouteTableRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAssociateRouteTableRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) SetRouteTableId(a string)  {
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
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) GetRouteTableId() string {
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

// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest withRouteTableId(java.lang.String)
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) WithRouteTableId(a string) *ServicesEc2ModelAssociateRouteTableRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableId", "com/amazonaws/services/ec2/model/AssociateRouteTableRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAssociateRouteTableRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AssociateRouteTableRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AssociateRouteTableRequest clone()
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) Clone3() *ServicesEc2ModelAssociateRouteTableRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AssociateRouteTableRequest")
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
	unique_x := &ServicesEc2ModelAssociateRouteTableRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelAssociateRouteTableRequest) Clone2() (*JavaLangObject, error) {
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


