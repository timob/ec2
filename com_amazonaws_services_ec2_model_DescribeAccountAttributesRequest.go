package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeAccountAttributesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getAttributeNames()
	GetAttributeNames() []string

	// public void setAttributeNames(java.util.Collection<java.lang.String>)
	SetAttributeNames(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest withAttributeNames(java.lang.String...)
	WithAttributeNames2(a ...string) *ServicesEc2ModelDescribeAccountAttributesRequest

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest withAttributeNames(java.util.Collection<java.lang.String>)
	WithAttributeNames3(a []string) *ServicesEc2ModelDescribeAccountAttributesRequest

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest withAttributeNames(com.amazonaws.services.ec2.model.AccountAttributeName...)
	WithAttributeNames(a ...*ServicesEc2ModelAccountAttributeName) *ServicesEc2ModelDescribeAccountAttributesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest clone()
	Clone3() *ServicesEc2ModelDescribeAccountAttributesRequest
}

type ServicesEc2ModelDescribeAccountAttributesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest()
func NewServicesEc2ModelDescribeAccountAttributesRequest() (*ServicesEc2ModelDescribeAccountAttributesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeAccountAttributesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeAccountAttributesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getAttributeNames()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) GetAttributeNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttributeNames", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAttributeNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) SetAttributeNames(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttributeNames", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest withAttributeNames(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) WithAttributeNames2(a ...string) *ServicesEc2ModelDescribeAccountAttributesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeNames", "com/amazonaws/services/ec2/model/DescribeAccountAttributesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest withAttributeNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) WithAttributeNames3(a []string) *ServicesEc2ModelDescribeAccountAttributesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeNames", "com/amazonaws/services/ec2/model/DescribeAccountAttributesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest withAttributeNames(com.amazonaws.services.ec2.model.AccountAttributeName...)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) WithAttributeNames(a ...*ServicesEc2ModelAccountAttributeName) *ServicesEc2ModelDescribeAccountAttributesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/AccountAttributeName")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeNames", "com/amazonaws/services/ec2/model/DescribeAccountAttributesRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AccountAttributeName"))
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesRequest clone()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) Clone3() *ServicesEc2ModelDescribeAccountAttributesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeAccountAttributesRequest")
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDescribeAccountAttributesRequest) Clone2() (*JavaLangObject, error) {
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


