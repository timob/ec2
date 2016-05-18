package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportSnapshotRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportSnapshotRequest

	// public void setDiskContainer(com.amazonaws.services.ec2.model.SnapshotDiskContainer)
	SetDiskContainer(a ServicesEc2ModelSnapshotDiskContainerInterface) 

	// public com.amazonaws.services.ec2.model.SnapshotDiskContainer getDiskContainer()
	GetDiskContainer() *ServicesEc2ModelSnapshotDiskContainer

	// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withDiskContainer(com.amazonaws.services.ec2.model.SnapshotDiskContainer)
	WithDiskContainer(a ServicesEc2ModelSnapshotDiskContainerInterface) *ServicesEc2ModelImportSnapshotRequest

	// public void setClientData(com.amazonaws.services.ec2.model.ClientData)
	SetClientData(a ServicesEc2ModelClientDataInterface) 

	// public com.amazonaws.services.ec2.model.ClientData getClientData()
	GetClientData() *ServicesEc2ModelClientData

	// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withClientData(com.amazonaws.services.ec2.model.ClientData)
	WithClientData(a ServicesEc2ModelClientDataInterface) *ServicesEc2ModelImportSnapshotRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelImportSnapshotRequest

	// public void setRoleName(java.lang.String)
	SetRoleName(a string) 

	// public java.lang.String getRoleName()
	GetRoleName() string

	// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withRoleName(java.lang.String)
	WithRoleName(a string) *ServicesEc2ModelImportSnapshotRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportSnapshotRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ImportSnapshotRequest clone()
	Clone3() *ServicesEc2ModelImportSnapshotRequest
}

type ServicesEc2ModelImportSnapshotRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest()
func NewServicesEc2ModelImportSnapshotRequest() (*ServicesEc2ModelImportSnapshotRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportSnapshotRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportSnapshotRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportSnapshotRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) WithDescription(a string) *ServicesEc2ModelImportSnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportSnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportSnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDiskContainer(com.amazonaws.services.ec2.model.SnapshotDiskContainer)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) SetDiskContainer(a ServicesEc2ModelSnapshotDiskContainerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDiskContainer", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotDiskContainer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer getDiskContainer()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) GetDiskContainer() *ServicesEc2ModelSnapshotDiskContainer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDiskContainer", "com/amazonaws/services/ec2/model/SnapshotDiskContainer")
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
	unique_x := &ServicesEc2ModelSnapshotDiskContainer{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withDiskContainer(com.amazonaws.services.ec2.model.SnapshotDiskContainer)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) WithDiskContainer(a ServicesEc2ModelSnapshotDiskContainerInterface) *ServicesEc2ModelImportSnapshotRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskContainer", "com/amazonaws/services/ec2/model/ImportSnapshotRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotDiskContainer"))
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
	unique_x := &ServicesEc2ModelImportSnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientData(com.amazonaws.services.ec2.model.ClientData)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) SetClientData(a ServicesEc2ModelClientDataInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientData", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ClientData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ClientData getClientData()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) GetClientData() *ServicesEc2ModelClientData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientData", "com/amazonaws/services/ec2/model/ClientData")
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
	unique_x := &ServicesEc2ModelClientData{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withClientData(com.amazonaws.services.ec2.model.ClientData)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) WithClientData(a ServicesEc2ModelClientDataInterface) *ServicesEc2ModelImportSnapshotRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientData", "com/amazonaws/services/ec2/model/ImportSnapshotRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ClientData"))
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
	unique_x := &ServicesEc2ModelImportSnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) WithClientToken(a string) *ServicesEc2ModelImportSnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/ImportSnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportSnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRoleName(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) SetRoleName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRoleName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRoleName()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) GetRoleName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRoleName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest withRoleName(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotRequest) WithRoleName(a string) *ServicesEc2ModelImportSnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRoleName", "com/amazonaws/services/ec2/model/ImportSnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportSnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportSnapshotRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelImportSnapshotRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelImportSnapshotRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportSnapshotRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportSnapshotRequest clone()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) Clone3() *ServicesEc2ModelImportSnapshotRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportSnapshotRequest")
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
	unique_x := &ServicesEc2ModelImportSnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelImportSnapshotRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelImportSnapshotRequest) Clone2() (*JavaLangObject, error) {
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


