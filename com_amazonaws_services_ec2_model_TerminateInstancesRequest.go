package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelTerminateInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getInstanceIds()
	GetInstanceIds() []string

	// public void setInstanceIds(java.util.Collection<java.lang.String>)
	SetInstanceIds(a []string) 

	// public com.amazonaws.services.ec2.model.TerminateInstancesRequest withInstanceIds(java.lang.String...)
	WithInstanceIds(a ...string) *ServicesEc2ModelTerminateInstancesRequest

	// public com.amazonaws.services.ec2.model.TerminateInstancesRequest withInstanceIds(java.util.Collection<java.lang.String>)
	WithInstanceIds2(a []string) *ServicesEc2ModelTerminateInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.TerminateInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.TerminateInstancesRequest clone()
	Clone3() *ServicesEc2ModelTerminateInstancesRequest
}

type ServicesEc2ModelTerminateInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.TerminateInstancesRequest()
func NewServicesEc2ModelTerminateInstancesRequest() (*ServicesEc2ModelTerminateInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/TerminateInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelTerminateInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.TerminateInstancesRequest(java.util.List<java.lang.String>)
func NewServicesEc2ModelTerminateInstancesRequest2(a []string) (*ServicesEc2ModelTerminateInstancesRequest) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/TerminateInstancesRequest", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelTerminateInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getInstanceIds()
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) GetInstanceIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceIds", "java/util/List")
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

// public void setInstanceIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) SetInstanceIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.TerminateInstancesRequest withInstanceIds(java.lang.String...)
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) WithInstanceIds(a ...string) *ServicesEc2ModelTerminateInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceIds", "com/amazonaws/services/ec2/model/TerminateInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTerminateInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.TerminateInstancesRequest withInstanceIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) WithInstanceIds2(a []string) *ServicesEc2ModelTerminateInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceIds", "com/amazonaws/services/ec2/model/TerminateInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelTerminateInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.TerminateInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.TerminateInstancesRequest clone()
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) Clone3() *ServicesEc2ModelTerminateInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/TerminateInstancesRequest")
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
	unique_x := &ServicesEc2ModelTerminateInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelTerminateInstancesRequest) Clone2() (*JavaLangObject, error) {
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

