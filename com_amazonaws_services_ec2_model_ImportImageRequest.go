package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportImageRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportImageRequest

	// public java.util.List<com.amazonaws.services.ec2.model.ImageDiskContainer> getDiskContainers()
	GetDiskContainers() []*ServicesEc2ModelImageDiskContainer

	// public void setDiskContainers(java.util.Collection<com.amazonaws.services.ec2.model.ImageDiskContainer>)
	SetDiskContainers(a []*ServicesEc2ModelImageDiskContainer) 

	// public com.amazonaws.services.ec2.model.ImportImageRequest withDiskContainers(com.amazonaws.services.ec2.model.ImageDiskContainer...)
	WithDiskContainers(a ...*ServicesEc2ModelImageDiskContainer) *ServicesEc2ModelImportImageRequest

	// public com.amazonaws.services.ec2.model.ImportImageRequest withDiskContainers(java.util.Collection<com.amazonaws.services.ec2.model.ImageDiskContainer>)
	WithDiskContainers2(a []*ServicesEc2ModelImageDiskContainer) *ServicesEc2ModelImportImageRequest

	// public void setLicenseType(java.lang.String)
	SetLicenseType(a string) 

	// public java.lang.String getLicenseType()
	GetLicenseType() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withLicenseType(java.lang.String)
	WithLicenseType(a string) *ServicesEc2ModelImportImageRequest

	// public void setHypervisor(java.lang.String)
	SetHypervisor(a string) 

	// public java.lang.String getHypervisor()
	GetHypervisor() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withHypervisor(java.lang.String)
	WithHypervisor(a string) *ServicesEc2ModelImportImageRequest

	// public void setArchitecture(java.lang.String)
	SetArchitecture(a string) 

	// public java.lang.String getArchitecture()
	GetArchitecture() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withArchitecture(java.lang.String)
	WithArchitecture(a string) *ServicesEc2ModelImportImageRequest

	// public void setPlatform(java.lang.String)
	SetPlatform(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withPlatform(java.lang.String)
	WithPlatform(a string) *ServicesEc2ModelImportImageRequest

	// public void setClientData(com.amazonaws.services.ec2.model.ClientData)
	SetClientData(a ServicesEc2ModelClientDataInterface) 

	// public com.amazonaws.services.ec2.model.ClientData getClientData()
	GetClientData() *ServicesEc2ModelClientData

	// public com.amazonaws.services.ec2.model.ImportImageRequest withClientData(com.amazonaws.services.ec2.model.ClientData)
	WithClientData(a ServicesEc2ModelClientDataInterface) *ServicesEc2ModelImportImageRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelImportImageRequest

	// public void setRoleName(java.lang.String)
	SetRoleName(a string) 

	// public java.lang.String getRoleName()
	GetRoleName() string

	// public com.amazonaws.services.ec2.model.ImportImageRequest withRoleName(java.lang.String)
	WithRoleName(a string) *ServicesEc2ModelImportImageRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportImageRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ImportImageRequest clone()
	Clone3() *ServicesEc2ModelImportImageRequest
}

type ServicesEc2ModelImportImageRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ImportImageRequest()
func NewServicesEc2ModelImportImageRequest() (*ServicesEc2ModelImportImageRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportImageRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportImageRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithDescription(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ImageDiskContainer> getDiskContainers()
func (jbobject *ServicesEc2ModelImportImageRequest) GetDiskContainers() []*ServicesEc2ModelImageDiskContainer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDiskContainers", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelImageDiskContainer)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDiskContainers(java.util.Collection<com.amazonaws.services.ec2.model.ImageDiskContainer>)
func (jbobject *ServicesEc2ModelImportImageRequest) SetDiskContainers(a []*ServicesEc2ModelImageDiskContainer)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDiskContainers", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportImageRequest withDiskContainers(com.amazonaws.services.ec2.model.ImageDiskContainer...)
func (jbobject *ServicesEc2ModelImportImageRequest) WithDiskContainers(a ...*ServicesEc2ModelImageDiskContainer) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ImageDiskContainer")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskContainers", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageDiskContainer"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportImageRequest withDiskContainers(java.util.Collection<com.amazonaws.services.ec2.model.ImageDiskContainer>)
func (jbobject *ServicesEc2ModelImportImageRequest) WithDiskContainers2(a []*ServicesEc2ModelImageDiskContainer) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskContainers", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLicenseType(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetLicenseType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLicenseType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getLicenseType()
func (jbobject *ServicesEc2ModelImportImageRequest) GetLicenseType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLicenseType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withLicenseType(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithLicenseType(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLicenseType", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetHypervisor(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHypervisor", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHypervisor()
func (jbobject *ServicesEc2ModelImportImageRequest) GetHypervisor() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHypervisor", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithHypervisor(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHypervisor", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetArchitecture(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArchitecture", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getArchitecture()
func (jbobject *ServicesEc2ModelImportImageRequest) GetArchitecture() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArchitecture", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithArchitecture(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetPlatform(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPlatform()
func (jbobject *ServicesEc2ModelImportImageRequest) GetPlatform() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlatform", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithPlatform(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientData(com.amazonaws.services.ec2.model.ClientData)
func (jbobject *ServicesEc2ModelImportImageRequest) SetClientData(a ServicesEc2ModelClientDataInterface)  {
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
func (jbobject *ServicesEc2ModelImportImageRequest) GetClientData() *ServicesEc2ModelClientData {
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withClientData(com.amazonaws.services.ec2.model.ClientData)
func (jbobject *ServicesEc2ModelImportImageRequest) WithClientData(a ServicesEc2ModelClientDataInterface) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientData", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ClientData"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithClientToken(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRoleName(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) SetRoleName(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageRequest) GetRoleName() string {
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

// public com.amazonaws.services.ec2.model.ImportImageRequest withRoleName(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageRequest) WithRoleName(a string) *ServicesEc2ModelImportImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRoleName", "com/amazonaws/services/ec2/model/ImportImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportImageRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelImportImageRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelImportImageRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelImportImageRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportImageRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportImageRequest clone()
func (jbobject *ServicesEc2ModelImportImageRequest) Clone3() *ServicesEc2ModelImportImageRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportImageRequest")
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
	unique_x := &ServicesEc2ModelImportImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelImportImageRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelImportImageRequest) Clone2() (*JavaLangObject, error) {
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


