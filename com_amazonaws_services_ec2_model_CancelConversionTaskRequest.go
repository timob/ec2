package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelConversionTaskRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setConversionTaskId(java.lang.String)
	SetConversionTaskId(a string) 

	// public java.lang.String getConversionTaskId()
	GetConversionTaskId() string

	// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest withConversionTaskId(java.lang.String)
	WithConversionTaskId(a string) *ServicesEc2ModelCancelConversionTaskRequest

	// public void setReasonMessage(java.lang.String)
	SetReasonMessage(a string) 

	// public java.lang.String getReasonMessage()
	GetReasonMessage() string

	// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest withReasonMessage(java.lang.String)
	WithReasonMessage(a string) *ServicesEc2ModelCancelConversionTaskRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelConversionTaskRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest clone()
	Clone3() *ServicesEc2ModelCancelConversionTaskRequest
}

type ServicesEc2ModelCancelConversionTaskRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest()
func NewServicesEc2ModelCancelConversionTaskRequest() (*ServicesEc2ModelCancelConversionTaskRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelConversionTaskRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelConversionTaskRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setConversionTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) SetConversionTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConversionTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getConversionTaskId()
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) GetConversionTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConversionTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest withConversionTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) WithConversionTaskId(a string) *ServicesEc2ModelCancelConversionTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConversionTaskId", "com/amazonaws/services/ec2/model/CancelConversionTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelConversionTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setReasonMessage(java.lang.String)
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) SetReasonMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReasonMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReasonMessage()
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) GetReasonMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReasonMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest withReasonMessage(java.lang.String)
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) WithReasonMessage(a string) *ServicesEc2ModelCancelConversionTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReasonMessage", "com/amazonaws/services/ec2/model/CancelConversionTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelConversionTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelConversionTaskRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelConversionTaskRequest clone()
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) Clone3() *ServicesEc2ModelCancelConversionTaskRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelConversionTaskRequest")
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
	unique_x := &ServicesEc2ModelCancelConversionTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCancelConversionTaskRequest) Clone2() (*JavaLangObject, error) {
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


