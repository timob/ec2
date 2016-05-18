package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReleaseAddressRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.ReleaseAddressRequest withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelReleaseAddressRequest

	// public void setAllocationId(java.lang.String)
	SetAllocationId(a string) 

	// public java.lang.String getAllocationId()
	GetAllocationId() string

	// public com.amazonaws.services.ec2.model.ReleaseAddressRequest withAllocationId(java.lang.String)
	WithAllocationId(a string) *ServicesEc2ModelReleaseAddressRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ReleaseAddressRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ReleaseAddressRequest clone()
	Clone3() *ServicesEc2ModelReleaseAddressRequest
}

type ServicesEc2ModelReleaseAddressRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ReleaseAddressRequest()
func NewServicesEc2ModelReleaseAddressRequest() (*ServicesEc2ModelReleaseAddressRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReleaseAddressRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReleaseAddressRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ReleaseAddressRequest(java.lang.String)
func NewServicesEc2ModelReleaseAddressRequest2(a string) (*ServicesEc2ModelReleaseAddressRequest) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReleaseAddressRequest", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelReleaseAddressRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelReleaseAddressRequest) SetPublicIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIp()
func (jbobject *ServicesEc2ModelReleaseAddressRequest) GetPublicIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReleaseAddressRequest withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelReleaseAddressRequest) WithPublicIp(a string) *ServicesEc2ModelReleaseAddressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/ReleaseAddressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReleaseAddressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAllocationId(java.lang.String)
func (jbobject *ServicesEc2ModelReleaseAddressRequest) SetAllocationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllocationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAllocationId()
func (jbobject *ServicesEc2ModelReleaseAddressRequest) GetAllocationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllocationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReleaseAddressRequest withAllocationId(java.lang.String)
func (jbobject *ServicesEc2ModelReleaseAddressRequest) WithAllocationId(a string) *ServicesEc2ModelReleaseAddressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllocationId", "com/amazonaws/services/ec2/model/ReleaseAddressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReleaseAddressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ReleaseAddressRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelReleaseAddressRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelReleaseAddressRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelReleaseAddressRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReleaseAddressRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReleaseAddressRequest clone()
func (jbobject *ServicesEc2ModelReleaseAddressRequest) Clone3() *ServicesEc2ModelReleaseAddressRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReleaseAddressRequest")
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
	unique_x := &ServicesEc2ModelReleaseAddressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelReleaseAddressRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelReleaseAddressRequest) Clone2() (*JavaLangObject, error) {
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


