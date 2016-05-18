package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifySubnetAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelModifySubnetAttributeRequest

	// public void setMapPublicIpOnLaunch(java.lang.Boolean)
	SetMapPublicIpOnLaunch(a bool) 

	// public java.lang.Boolean getMapPublicIpOnLaunch()
	GetMapPublicIpOnLaunch() bool

	// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest withMapPublicIpOnLaunch(java.lang.Boolean)
	WithMapPublicIpOnLaunch(a bool) *ServicesEc2ModelModifySubnetAttributeRequest

	// public java.lang.Boolean isMapPublicIpOnLaunch()
	IsMapPublicIpOnLaunch() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest clone()
	Clone3() *ServicesEc2ModelModifySubnetAttributeRequest
}

type ServicesEc2ModelModifySubnetAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest()
func NewServicesEc2ModelModifySubnetAttributeRequest() (*ServicesEc2ModelModifySubnetAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifySubnetAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifySubnetAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) WithSubnetId(a string) *ServicesEc2ModelModifySubnetAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/ModifySubnetAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySubnetAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMapPublicIpOnLaunch(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) SetMapPublicIpOnLaunch(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMapPublicIpOnLaunch", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getMapPublicIpOnLaunch()
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) GetMapPublicIpOnLaunch() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMapPublicIpOnLaunch", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest withMapPublicIpOnLaunch(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) WithMapPublicIpOnLaunch(a bool) *ServicesEc2ModelModifySubnetAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMapPublicIpOnLaunch", "com/amazonaws/services/ec2/model/ModifySubnetAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifySubnetAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isMapPublicIpOnLaunch()
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) IsMapPublicIpOnLaunch() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMapPublicIpOnLaunch", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifySubnetAttributeRequest clone()
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) Clone3() *ServicesEc2ModelModifySubnetAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifySubnetAttributeRequest")
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
	unique_x := &ServicesEc2ModelModifySubnetAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifySubnetAttributeRequest) Clone2() (*JavaLangObject, error) {
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


