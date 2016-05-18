package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVolumeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSize(java.lang.Integer)
	SetSize(a int) 

	// public java.lang.Integer getSize()
	GetSize() int

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withSize(java.lang.Integer)
	WithSize(a int) *ServicesEc2ModelCreateVolumeRequest

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelCreateVolumeRequest

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelCreateVolumeRequest

	// public void setVolumeType(java.lang.String)
	SetVolumeType2(a string) 

	// public java.lang.String getVolumeType()
	GetVolumeType() string

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withVolumeType(java.lang.String)
	WithVolumeType2(a string) *ServicesEc2ModelCreateVolumeRequest

	// public void setVolumeType(com.amazonaws.services.ec2.model.VolumeType)
	SetVolumeType(a ServicesEc2ModelVolumeTypeInterface) 

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withVolumeType(com.amazonaws.services.ec2.model.VolumeType)
	WithVolumeType(a ServicesEc2ModelVolumeTypeInterface) *ServicesEc2ModelCreateVolumeRequest

	// public void setIops(java.lang.Integer)
	SetIops(a int) 

	// public java.lang.Integer getIops()
	GetIops() int

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withIops(java.lang.Integer)
	WithIops(a int) *ServicesEc2ModelCreateVolumeRequest

	// public void setEncrypted(java.lang.Boolean)
	SetEncrypted(a bool) 

	// public java.lang.Boolean getEncrypted()
	GetEncrypted() bool

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withEncrypted(java.lang.Boolean)
	WithEncrypted(a bool) *ServicesEc2ModelCreateVolumeRequest

	// public java.lang.Boolean isEncrypted()
	IsEncrypted() bool

	// public void setKmsKeyId(java.lang.String)
	SetKmsKeyId(a string) 

	// public java.lang.String getKmsKeyId()
	GetKmsKeyId() string

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest withKmsKeyId(java.lang.String)
	WithKmsKeyId(a string) *ServicesEc2ModelCreateVolumeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVolumeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateVolumeRequest clone()
	Clone3() *ServicesEc2ModelCreateVolumeRequest
}

type ServicesEc2ModelCreateVolumeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest()
func NewServicesEc2ModelCreateVolumeRequest() (*ServicesEc2ModelCreateVolumeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVolumeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVolumeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest(java.lang.Integer, java.lang.String)
func NewServicesEc2ModelCreateVolumeRequest2(a int, b string) (*ServicesEc2ModelCreateVolumeRequest) {
	conv_a := javabind.NewGoToJavaInteger()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/Integer"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelCreateVolumeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest(java.lang.String, java.lang.String)
func NewServicesEc2ModelCreateVolumeRequest3(a string, b string) (*ServicesEc2ModelCreateVolumeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelCreateVolumeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetSize(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getSize()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetSize() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithSize(a int) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSize", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetSnapshotId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSnapshotId()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetSnapshotId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithSnapshotId(a string) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithAvailabilityZone(a string) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetVolumeType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeType()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetVolumeType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithVolumeType2(a string) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeType(com.amazonaws.services.ec2.model.VolumeType)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetVolumeType(a ServicesEc2ModelVolumeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withVolumeType(com.amazonaws.services.ec2.model.VolumeType)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithVolumeType(a ServicesEc2ModelVolumeTypeInterface) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeType"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetIops(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIops", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getIops()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetIops() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIops", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithIops(a int) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIops", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetEncrypted(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEncrypted", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEncrypted()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetEncrypted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEncrypted", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithEncrypted(a bool) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEncrypted", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEncrypted()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) IsEncrypted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEncrypted", "java/lang/Boolean")
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

// public void setKmsKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) SetKmsKeyId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKmsKeyId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKmsKeyId()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetKmsKeyId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKmsKeyId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVolumeRequest withKmsKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumeRequest) WithKmsKeyId(a string) *ServicesEc2ModelCreateVolumeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKmsKeyId", "com/amazonaws/services/ec2/model/CreateVolumeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVolumeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateVolumeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVolumeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVolumeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVolumeRequest clone()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) Clone3() *ServicesEc2ModelCreateVolumeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVolumeRequest")
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
	unique_x := &ServicesEc2ModelCreateVolumeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateVolumeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateVolumeRequest) Clone2() (*JavaLangObject, error) {
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


