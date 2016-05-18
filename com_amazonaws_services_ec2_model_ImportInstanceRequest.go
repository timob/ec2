package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportInstanceRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportInstanceRequest

	// public void setLaunchSpecification(com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification)
	SetLaunchSpecification(a ServicesEc2ModelImportInstanceLaunchSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification getLaunchSpecification()
	GetLaunchSpecification() *ServicesEc2ModelImportInstanceLaunchSpecification

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest withLaunchSpecification(com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification)
	WithLaunchSpecification(a ServicesEc2ModelImportInstanceLaunchSpecificationInterface) *ServicesEc2ModelImportInstanceRequest

	// public java.util.List<com.amazonaws.services.ec2.model.DiskImage> getDiskImages()
	GetDiskImages() []*ServicesEc2ModelDiskImage

	// public void setDiskImages(java.util.Collection<com.amazonaws.services.ec2.model.DiskImage>)
	SetDiskImages(a []*ServicesEc2ModelDiskImage) 

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest withDiskImages(com.amazonaws.services.ec2.model.DiskImage...)
	WithDiskImages(a ...*ServicesEc2ModelDiskImage) *ServicesEc2ModelImportInstanceRequest

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest withDiskImages(java.util.Collection<com.amazonaws.services.ec2.model.DiskImage>)
	WithDiskImages2(a []*ServicesEc2ModelDiskImage) *ServicesEc2ModelImportInstanceRequest

	// public void setPlatform(java.lang.String)
	SetPlatform2(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest withPlatform(java.lang.String)
	WithPlatform2(a string) *ServicesEc2ModelImportInstanceRequest

	// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	SetPlatform(a ServicesEc2ModelPlatformValuesInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelImportInstanceRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportInstanceRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ImportInstanceRequest clone()
	Clone3() *ServicesEc2ModelImportInstanceRequest
}

type ServicesEc2ModelImportInstanceRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ImportInstanceRequest()
func NewServicesEc2ModelImportInstanceRequest() (*ServicesEc2ModelImportInstanceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportInstanceRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportInstanceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceRequest) WithDescription(a string) *ServicesEc2ModelImportInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchSpecification(com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification)
func (jbobject *ServicesEc2ModelImportInstanceRequest) SetLaunchSpecification(a ServicesEc2ModelImportInstanceLaunchSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchSpecification", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification getLaunchSpecification()
func (jbobject *ServicesEc2ModelImportInstanceRequest) GetLaunchSpecification() *ServicesEc2ModelImportInstanceLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchSpecification", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification")
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceRequest withLaunchSpecification(com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification)
func (jbobject *ServicesEc2ModelImportInstanceRequest) WithLaunchSpecification(a ServicesEc2ModelImportInstanceLaunchSpecificationInterface) *ServicesEc2ModelImportInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchSpecification", "com/amazonaws/services/ec2/model/ImportInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification"))
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.DiskImage> getDiskImages()
func (jbobject *ServicesEc2ModelImportInstanceRequest) GetDiskImages() []*ServicesEc2ModelDiskImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDiskImages", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelDiskImage)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDiskImages(java.util.Collection<com.amazonaws.services.ec2.model.DiskImage>)
func (jbobject *ServicesEc2ModelImportInstanceRequest) SetDiskImages(a []*ServicesEc2ModelDiskImage)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDiskImages", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceRequest withDiskImages(com.amazonaws.services.ec2.model.DiskImage...)
func (jbobject *ServicesEc2ModelImportInstanceRequest) WithDiskImages(a ...*ServicesEc2ModelDiskImage) *ServicesEc2ModelImportInstanceRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/DiskImage")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskImages", "com/amazonaws/services/ec2/model/ImportInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImage"))
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceRequest withDiskImages(java.util.Collection<com.amazonaws.services.ec2.model.DiskImage>)
func (jbobject *ServicesEc2ModelImportInstanceRequest) WithDiskImages2(a []*ServicesEc2ModelDiskImage) *ServicesEc2ModelImportInstanceRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskImages", "com/amazonaws/services/ec2/model/ImportInstanceRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceRequest) SetPlatform2(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceRequest) GetPlatform() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceRequest withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceRequest) WithPlatform2(a string) *ServicesEc2ModelImportInstanceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ImportInstanceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelImportInstanceRequest) SetPlatform(a ServicesEc2ModelPlatformValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceRequest withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelImportInstanceRequest) WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelImportInstanceRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ImportInstanceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportInstanceRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelImportInstanceRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelImportInstanceRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelImportInstanceRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportInstanceRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportInstanceRequest clone()
func (jbobject *ServicesEc2ModelImportInstanceRequest) Clone3() *ServicesEc2ModelImportInstanceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportInstanceRequest")
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
	unique_x := &ServicesEc2ModelImportInstanceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelImportInstanceRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelImportInstanceRequest) Clone2() (*JavaLangObject, error) {
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


