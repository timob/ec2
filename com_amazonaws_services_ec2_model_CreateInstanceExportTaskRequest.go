package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateInstanceExportTaskRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelCreateInstanceExportTaskRequest

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelCreateInstanceExportTaskRequest

	// public void setTargetEnvironment(java.lang.String)
	SetTargetEnvironment2(a string) 

	// public java.lang.String getTargetEnvironment()
	GetTargetEnvironment() string

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withTargetEnvironment(java.lang.String)
	WithTargetEnvironment2(a string) *ServicesEc2ModelCreateInstanceExportTaskRequest

	// public void setTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
	SetTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface) 

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
	WithTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface) *ServicesEc2ModelCreateInstanceExportTaskRequest

	// public void setExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3TaskSpecification)
	SetExportToS3Task(a ServicesEc2ModelExportToS3TaskSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification getExportToS3Task()
	GetExportToS3Task() *ServicesEc2ModelExportToS3TaskSpecification

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3TaskSpecification)
	WithExportToS3Task(a ServicesEc2ModelExportToS3TaskSpecificationInterface) *ServicesEc2ModelCreateInstanceExportTaskRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest clone()
	Clone3() *ServicesEc2ModelCreateInstanceExportTaskRequest
}

type ServicesEc2ModelCreateInstanceExportTaskRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest()
func NewServicesEc2ModelCreateInstanceExportTaskRequest() (*ServicesEc2ModelCreateInstanceExportTaskRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) WithDescription(a string) *ServicesEc2ModelCreateInstanceExportTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) WithInstanceId(a string) *ServicesEc2ModelCreateInstanceExportTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetEnvironment(java.lang.String)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) SetTargetEnvironment2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetEnvironment", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getTargetEnvironment()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) GetTargetEnvironment() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTargetEnvironment", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withTargetEnvironment(java.lang.String)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) WithTargetEnvironment2(a string) *ServicesEc2ModelCreateInstanceExportTaskRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetEnvironment", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) SetTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTargetEnvironment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportEnvironment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withTargetEnvironment(com.amazonaws.services.ec2.model.ExportEnvironment)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) WithTargetEnvironment(a ServicesEc2ModelExportEnvironmentInterface) *ServicesEc2ModelCreateInstanceExportTaskRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTargetEnvironment", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportEnvironment"))
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3TaskSpecification)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) SetExportToS3Task(a ServicesEc2ModelExportToS3TaskSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExportToS3Task", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportToS3TaskSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification getExportToS3Task()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) GetExportToS3Task() *ServicesEc2ModelExportToS3TaskSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExportToS3Task", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification")
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest withExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3TaskSpecification)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) WithExportToS3Task(a ServicesEc2ModelExportToS3TaskSpecificationInterface) *ServicesEc2ModelCreateInstanceExportTaskRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportToS3Task", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportToS3TaskSpecification"))
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskRequest clone()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) Clone3() *ServicesEc2ModelCreateInstanceExportTaskRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskRequest")
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskRequest) Clone2() (*JavaLangObject, error) {
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


