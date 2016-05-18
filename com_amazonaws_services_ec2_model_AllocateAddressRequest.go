package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAllocateAddressRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setDomain(java.lang.String)
	SetDomain2(a string) 

	// public java.lang.String getDomain()
	GetDomain() string

	// public com.amazonaws.services.ec2.model.AllocateAddressRequest withDomain(java.lang.String)
	WithDomain2(a string) *ServicesEc2ModelAllocateAddressRequest

	// public void setDomain(com.amazonaws.services.ec2.model.DomainType)
	SetDomain(a ServicesEc2ModelDomainTypeInterface) 

	// public com.amazonaws.services.ec2.model.AllocateAddressRequest withDomain(com.amazonaws.services.ec2.model.DomainType)
	WithDomain(a ServicesEc2ModelDomainTypeInterface) *ServicesEc2ModelAllocateAddressRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AllocateAddressRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.AllocateAddressRequest clone()
	Clone3() *ServicesEc2ModelAllocateAddressRequest
}

type ServicesEc2ModelAllocateAddressRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.AllocateAddressRequest()
func NewServicesEc2ModelAllocateAddressRequest() (*ServicesEc2ModelAllocateAddressRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AllocateAddressRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAllocateAddressRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDomain(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressRequest) SetDomain2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDomain", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDomain()
func (jbobject *ServicesEc2ModelAllocateAddressRequest) GetDomain() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDomain", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AllocateAddressRequest withDomain(java.lang.String)
func (jbobject *ServicesEc2ModelAllocateAddressRequest) WithDomain2(a string) *ServicesEc2ModelAllocateAddressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDomain", "com/amazonaws/services/ec2/model/AllocateAddressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateAddressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDomain(com.amazonaws.services.ec2.model.DomainType)
func (jbobject *ServicesEc2ModelAllocateAddressRequest) SetDomain(a ServicesEc2ModelDomainTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDomain", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DomainType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AllocateAddressRequest withDomain(com.amazonaws.services.ec2.model.DomainType)
func (jbobject *ServicesEc2ModelAllocateAddressRequest) WithDomain(a ServicesEc2ModelDomainTypeInterface) *ServicesEc2ModelAllocateAddressRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDomain", "com/amazonaws/services/ec2/model/AllocateAddressRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DomainType"))
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
	unique_x := &ServicesEc2ModelAllocateAddressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AllocateAddressRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelAllocateAddressRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelAllocateAddressRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelAllocateAddressRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAllocateAddressRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AllocateAddressRequest clone()
func (jbobject *ServicesEc2ModelAllocateAddressRequest) Clone3() *ServicesEc2ModelAllocateAddressRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AllocateAddressRequest")
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
	unique_x := &ServicesEc2ModelAllocateAddressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelAllocateAddressRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelAllocateAddressRequest) Clone2() (*JavaLangObject, error) {
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


