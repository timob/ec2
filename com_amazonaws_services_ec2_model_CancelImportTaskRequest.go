package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelImportTaskRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setImportTaskId(java.lang.String)
	SetImportTaskId(a string) 

	// public java.lang.String getImportTaskId()
	GetImportTaskId() string

	// public com.amazonaws.services.ec2.model.CancelImportTaskRequest withImportTaskId(java.lang.String)
	WithImportTaskId(a string) *ServicesEc2ModelCancelImportTaskRequest

	// public void setCancelReason(java.lang.String)
	SetCancelReason(a string) 

	// public java.lang.String getCancelReason()
	GetCancelReason() string

	// public com.amazonaws.services.ec2.model.CancelImportTaskRequest withCancelReason(java.lang.String)
	WithCancelReason(a string) *ServicesEc2ModelCancelImportTaskRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelImportTaskRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CancelImportTaskRequest clone()
	Clone3() *ServicesEc2ModelCancelImportTaskRequest
}

type ServicesEc2ModelCancelImportTaskRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CancelImportTaskRequest()
func NewServicesEc2ModelCancelImportTaskRequest() (*ServicesEc2ModelCancelImportTaskRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelImportTaskRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelImportTaskRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) SetImportTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImportTaskId()
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) GetImportTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelImportTaskRequest withImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) WithImportTaskId(a string) *ServicesEc2ModelCancelImportTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportTaskId", "com/amazonaws/services/ec2/model/CancelImportTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelImportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCancelReason(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) SetCancelReason(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCancelReason", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCancelReason()
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) GetCancelReason() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCancelReason", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelImportTaskRequest withCancelReason(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) WithCancelReason(a string) *ServicesEc2ModelCancelImportTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCancelReason", "com/amazonaws/services/ec2/model/CancelImportTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelImportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelImportTaskRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelImportTaskRequest clone()
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) Clone3() *ServicesEc2ModelCancelImportTaskRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelImportTaskRequest")
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
	unique_x := &ServicesEc2ModelCancelImportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCancelImportTaskRequest) Clone2() (*JavaLangObject, error) {
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


