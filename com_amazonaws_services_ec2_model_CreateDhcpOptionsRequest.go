package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateDhcpOptionsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<com.amazonaws.services.ec2.model.DhcpConfiguration> getDhcpConfigurations()
	GetDhcpConfigurations() []*ServicesEc2ModelDhcpConfiguration

	// public void setDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
	SetDhcpConfigurations(a []*ServicesEc2ModelDhcpConfiguration) 

	// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest withDhcpConfigurations(com.amazonaws.services.ec2.model.DhcpConfiguration...)
	WithDhcpConfigurations(a ...*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelCreateDhcpOptionsRequest

	// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest withDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
	WithDhcpConfigurations2(a []*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelCreateDhcpOptionsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest clone()
	Clone3() *ServicesEc2ModelCreateDhcpOptionsRequest
}

type ServicesEc2ModelCreateDhcpOptionsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest()
func NewServicesEc2ModelCreateDhcpOptionsRequest() (*ServicesEc2ModelCreateDhcpOptionsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateDhcpOptionsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateDhcpOptionsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest(java.util.List<com.amazonaws.services.ec2.model.DhcpConfiguration>)
func NewServicesEc2ModelCreateDhcpOptionsRequest2(a []*ServicesEc2ModelDhcpConfiguration) (*ServicesEc2ModelCreateDhcpOptionsRequest) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateDhcpOptionsRequest", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelCreateDhcpOptionsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.DhcpConfiguration> getDhcpConfigurations()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) GetDhcpConfigurations() []*ServicesEc2ModelDhcpConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDhcpConfigurations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelDhcpConfiguration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) SetDhcpConfigurations(a []*ServicesEc2ModelDhcpConfiguration)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDhcpConfigurations", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest withDhcpConfigurations(com.amazonaws.services.ec2.model.DhcpConfiguration...)
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) WithDhcpConfigurations(a ...*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelCreateDhcpOptionsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/DhcpConfiguration")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpConfigurations", "com/amazonaws/services/ec2/model/CreateDhcpOptionsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DhcpConfiguration"))
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
	unique_x := &ServicesEc2ModelCreateDhcpOptionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest withDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) WithDhcpConfigurations2(a []*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelCreateDhcpOptionsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpConfigurations", "com/amazonaws/services/ec2/model/CreateDhcpOptionsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateDhcpOptionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsRequest clone()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) Clone3() *ServicesEc2ModelCreateDhcpOptionsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateDhcpOptionsRequest")
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
	unique_x := &ServicesEc2ModelCreateDhcpOptionsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateDhcpOptionsRequest) Clone2() (*JavaLangObject, error) {
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


