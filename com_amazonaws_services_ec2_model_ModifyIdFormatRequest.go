package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyIdFormatRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setResource(java.lang.String)
	SetResource(a string) 

	// public java.lang.String getResource()
	GetResource() string

	// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest withResource(java.lang.String)
	WithResource(a string) *ServicesEc2ModelModifyIdFormatRequest

	// public void setUseLongIds(java.lang.Boolean)
	SetUseLongIds(a bool) 

	// public java.lang.Boolean getUseLongIds()
	GetUseLongIds() bool

	// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest withUseLongIds(java.lang.Boolean)
	WithUseLongIds(a bool) *ServicesEc2ModelModifyIdFormatRequest

	// public java.lang.Boolean isUseLongIds()
	IsUseLongIds() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyIdFormatRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest clone()
	Clone3() *ServicesEc2ModelModifyIdFormatRequest
}

type ServicesEc2ModelModifyIdFormatRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest()
func NewServicesEc2ModelModifyIdFormatRequest() (*ServicesEc2ModelModifyIdFormatRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyIdFormatRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyIdFormatRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setResource(java.lang.String)
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) SetResource(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResource", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getResource()
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) GetResource() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResource", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest withResource(java.lang.String)
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) WithResource(a string) *ServicesEc2ModelModifyIdFormatRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResource", "com/amazonaws/services/ec2/model/ModifyIdFormatRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyIdFormatRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUseLongIds(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) SetUseLongIds(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseLongIds", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getUseLongIds()
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) GetUseLongIds() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUseLongIds", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest withUseLongIds(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) WithUseLongIds(a bool) *ServicesEc2ModelModifyIdFormatRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUseLongIds", "com/amazonaws/services/ec2/model/ModifyIdFormatRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyIdFormatRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isUseLongIds()
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) IsUseLongIds() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isUseLongIds", "java/lang/Boolean")
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

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyIdFormatRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyIdFormatRequest clone()
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) Clone3() *ServicesEc2ModelModifyIdFormatRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyIdFormatRequest")
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
	unique_x := &ServicesEc2ModelModifyIdFormatRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyIdFormatRequest) Clone2() (*JavaLangObject, error) {
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


