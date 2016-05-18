package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportVolumeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ImportVolumeRequest withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelImportVolumeRequest

	// public void setImage(com.amazonaws.services.ec2.model.DiskImageDetail)
	SetImage(a ServicesEc2ModelDiskImageDetailInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageDetail getImage()
	GetImage() *ServicesEc2ModelDiskImageDetail

	// public com.amazonaws.services.ec2.model.ImportVolumeRequest withImage(com.amazonaws.services.ec2.model.DiskImageDetail)
	WithImage(a ServicesEc2ModelDiskImageDetailInterface) *ServicesEc2ModelImportVolumeRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportVolumeRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportVolumeRequest

	// public void setVolume(com.amazonaws.services.ec2.model.VolumeDetail)
	SetVolume(a ServicesEc2ModelVolumeDetailInterface) 

	// public com.amazonaws.services.ec2.model.VolumeDetail getVolume()
	GetVolume() *ServicesEc2ModelVolumeDetail

	// public com.amazonaws.services.ec2.model.ImportVolumeRequest withVolume(com.amazonaws.services.ec2.model.VolumeDetail)
	WithVolume(a ServicesEc2ModelVolumeDetailInterface) *ServicesEc2ModelImportVolumeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportVolumeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ImportVolumeRequest clone()
	Clone3() *ServicesEc2ModelImportVolumeRequest
}

type ServicesEc2ModelImportVolumeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ImportVolumeRequest()
func NewServicesEc2ModelImportVolumeRequest() (*ServicesEc2ModelImportVolumeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportVolumeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportVolumeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeRequest) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelImportVolumeRequest) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportVolumeRequest withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeRequest) WithAvailabilityZone(a string) *ServicesEc2ModelImportVolumeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ImportVolumeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImage(com.amazonaws.services.ec2.model.DiskImageDetail)
func (jbobject *ServicesEc2ModelImportVolumeRequest) SetImage(a ServicesEc2ModelDiskImageDetailInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDetail"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DiskImageDetail getImage()
func (jbobject *ServicesEc2ModelImportVolumeRequest) GetImage() *ServicesEc2ModelDiskImageDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImage", "com/amazonaws/services/ec2/model/DiskImageDetail")
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportVolumeRequest withImage(com.amazonaws.services.ec2.model.DiskImageDetail)
func (jbobject *ServicesEc2ModelImportVolumeRequest) WithImage(a ServicesEc2ModelDiskImageDetailInterface) *ServicesEc2ModelImportVolumeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImage", "com/amazonaws/services/ec2/model/ImportVolumeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDetail"))
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
	unique_x := &ServicesEc2ModelImportVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportVolumeRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportVolumeRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeRequest) WithDescription(a string) *ServicesEc2ModelImportVolumeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportVolumeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolume(com.amazonaws.services.ec2.model.VolumeDetail)
func (jbobject *ServicesEc2ModelImportVolumeRequest) SetVolume(a ServicesEc2ModelVolumeDetailInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolume", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeDetail"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeDetail getVolume()
func (jbobject *ServicesEc2ModelImportVolumeRequest) GetVolume() *ServicesEc2ModelVolumeDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolume", "com/amazonaws/services/ec2/model/VolumeDetail")
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
	unique_x := &ServicesEc2ModelVolumeDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportVolumeRequest withVolume(com.amazonaws.services.ec2.model.VolumeDetail)
func (jbobject *ServicesEc2ModelImportVolumeRequest) WithVolume(a ServicesEc2ModelVolumeDetailInterface) *ServicesEc2ModelImportVolumeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolume", "com/amazonaws/services/ec2/model/ImportVolumeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeDetail"))
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
	unique_x := &ServicesEc2ModelImportVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ImportVolumeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelImportVolumeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelImportVolumeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelImportVolumeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportVolumeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportVolumeRequest clone()
func (jbobject *ServicesEc2ModelImportVolumeRequest) Clone3() *ServicesEc2ModelImportVolumeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportVolumeRequest")
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
	unique_x := &ServicesEc2ModelImportVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelImportVolumeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelImportVolumeRequest) Clone2() (*JavaLangObject, error) {
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


