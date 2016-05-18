package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteTagsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getResources()
	GetResources() []string

	// public void setResources(java.util.Collection<java.lang.String>)
	SetResources(a []string) 

	// public com.amazonaws.services.ec2.model.DeleteTagsRequest withResources(java.lang.String...)
	WithResources(a ...string) *ServicesEc2ModelDeleteTagsRequest

	// public com.amazonaws.services.ec2.model.DeleteTagsRequest withResources(java.util.Collection<java.lang.String>)
	WithResources2(a []string) *ServicesEc2ModelDeleteTagsRequest

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.DeleteTagsRequest withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelDeleteTagsRequest

	// public com.amazonaws.services.ec2.model.DeleteTagsRequest withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelDeleteTagsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteTagsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteTagsRequest clone()
	Clone3() *ServicesEc2ModelDeleteTagsRequest
}

type ServicesEc2ModelDeleteTagsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest()
func NewServicesEc2ModelDeleteTagsRequest() (*ServicesEc2ModelDeleteTagsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteTagsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteTagsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest(java.util.List<java.lang.String>)
func NewServicesEc2ModelDeleteTagsRequest2(a []string) (*ServicesEc2ModelDeleteTagsRequest) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteTagsRequest", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelDeleteTagsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getResources()
func (jbobject *ServicesEc2ModelDeleteTagsRequest) GetResources() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResources", "java/util/List")
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

// public void setResources(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDeleteTagsRequest) SetResources(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResources", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest withResources(java.lang.String...)
func (jbobject *ServicesEc2ModelDeleteTagsRequest) WithResources(a ...string) *ServicesEc2ModelDeleteTagsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResources", "com/amazonaws/services/ec2/model/DeleteTagsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteTagsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest withResources(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDeleteTagsRequest) WithResources2(a []string) *ServicesEc2ModelDeleteTagsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResources", "com/amazonaws/services/ec2/model/DeleteTagsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDeleteTagsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelDeleteTagsRequest) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelDeleteTagsRequest) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelDeleteTagsRequest) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelDeleteTagsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/DeleteTagsRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelDeleteTagsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelDeleteTagsRequest) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelDeleteTagsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/DeleteTagsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDeleteTagsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteTagsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteTagsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDeleteTagsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteTagsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteTagsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteTagsRequest clone()
func (jbobject *ServicesEc2ModelDeleteTagsRequest) Clone3() *ServicesEc2ModelDeleteTagsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteTagsRequest")
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
	unique_x := &ServicesEc2ModelDeleteTagsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteTagsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDeleteTagsRequest) Clone2() (*JavaLangObject, error) {
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


