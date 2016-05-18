package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStartInstancesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getInstanceIds()
	GetInstanceIds() []string

	// public void setInstanceIds(java.util.Collection<java.lang.String>)
	SetInstanceIds(a []string) 

	// public com.amazonaws.services.ec2.model.StartInstancesRequest withInstanceIds(java.lang.String...)
	WithInstanceIds(a ...string) *ServicesEc2ModelStartInstancesRequest

	// public com.amazonaws.services.ec2.model.StartInstancesRequest withInstanceIds(java.util.Collection<java.lang.String>)
	WithInstanceIds2(a []string) *ServicesEc2ModelStartInstancesRequest

	// public void setAdditionalInfo(java.lang.String)
	SetAdditionalInfo(a string) 

	// public java.lang.String getAdditionalInfo()
	GetAdditionalInfo() string

	// public com.amazonaws.services.ec2.model.StartInstancesRequest withAdditionalInfo(java.lang.String)
	WithAdditionalInfo(a string) *ServicesEc2ModelStartInstancesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.StartInstancesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.StartInstancesRequest clone()
	Clone3() *ServicesEc2ModelStartInstancesRequest
}

type ServicesEc2ModelStartInstancesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.StartInstancesRequest()
func NewServicesEc2ModelStartInstancesRequest() (*ServicesEc2ModelStartInstancesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/StartInstancesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelStartInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.StartInstancesRequest(java.util.List<java.lang.String>)
func NewServicesEc2ModelStartInstancesRequest2(a []string) (*ServicesEc2ModelStartInstancesRequest) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/StartInstancesRequest", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelStartInstancesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getInstanceIds()
func (jbobject *ServicesEc2ModelStartInstancesRequest) GetInstanceIds() []string {
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
func (jbobject *ServicesEc2ModelStartInstancesRequest) SetInstanceIds(a []string)  {
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

// public com.amazonaws.services.ec2.model.StartInstancesRequest withInstanceIds(java.lang.String...)
func (jbobject *ServicesEc2ModelStartInstancesRequest) WithInstanceIds(a ...string) *ServicesEc2ModelStartInstancesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceIds", "com/amazonaws/services/ec2/model/StartInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStartInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.StartInstancesRequest withInstanceIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelStartInstancesRequest) WithInstanceIds2(a []string) *ServicesEc2ModelStartInstancesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceIds", "com/amazonaws/services/ec2/model/StartInstancesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelStartInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAdditionalInfo(java.lang.String)
func (jbobject *ServicesEc2ModelStartInstancesRequest) SetAdditionalInfo(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAdditionalInfo", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAdditionalInfo()
func (jbobject *ServicesEc2ModelStartInstancesRequest) GetAdditionalInfo() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdditionalInfo", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.StartInstancesRequest withAdditionalInfo(java.lang.String)
func (jbobject *ServicesEc2ModelStartInstancesRequest) WithAdditionalInfo(a string) *ServicesEc2ModelStartInstancesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdditionalInfo", "com/amazonaws/services/ec2/model/StartInstancesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStartInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.StartInstancesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelStartInstancesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelStartInstancesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelStartInstancesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelStartInstancesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.StartInstancesRequest clone()
func (jbobject *ServicesEc2ModelStartInstancesRequest) Clone3() *ServicesEc2ModelStartInstancesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/StartInstancesRequest")
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
	unique_x := &ServicesEc2ModelStartInstancesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelStartInstancesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelStartInstancesRequest) Clone2() (*JavaLangObject, error) {
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


