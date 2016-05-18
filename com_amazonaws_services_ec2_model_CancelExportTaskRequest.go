package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelExportTaskRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setExportTaskId(java.lang.String)
	SetExportTaskId(a string) 

	// public java.lang.String getExportTaskId()
	GetExportTaskId() string

	// public com.amazonaws.services.ec2.model.CancelExportTaskRequest withExportTaskId(java.lang.String)
	WithExportTaskId(a string) *ServicesEc2ModelCancelExportTaskRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelExportTaskRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CancelExportTaskRequest clone()
	Clone3() *ServicesEc2ModelCancelExportTaskRequest
}

type ServicesEc2ModelCancelExportTaskRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CancelExportTaskRequest()
func NewServicesEc2ModelCancelExportTaskRequest() (*ServicesEc2ModelCancelExportTaskRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelExportTaskRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelExportTaskRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setExportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) SetExportTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExportTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getExportTaskId()
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) GetExportTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExportTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelExportTaskRequest withExportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) WithExportTaskId(a string) *ServicesEc2ModelCancelExportTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportTaskId", "com/amazonaws/services/ec2/model/CancelExportTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelExportTaskRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelExportTaskRequest clone()
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) Clone3() *ServicesEc2ModelCancelExportTaskRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelExportTaskRequest")
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
	unique_x := &ServicesEc2ModelCancelExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCancelExportTaskRequest) Clone2() (*JavaLangObject, error) {
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


