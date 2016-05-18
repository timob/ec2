package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCopySnapshotRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSourceRegion(java.lang.String)
	SetSourceRegion(a string) 

	// public java.lang.String getSourceRegion()
	GetSourceRegion() string

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withSourceRegion(java.lang.String)
	WithSourceRegion(a string) *ServicesEc2ModelCopySnapshotRequest

	// public void setSourceSnapshotId(java.lang.String)
	SetSourceSnapshotId(a string) 

	// public java.lang.String getSourceSnapshotId()
	GetSourceSnapshotId() string

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withSourceSnapshotId(java.lang.String)
	WithSourceSnapshotId(a string) *ServicesEc2ModelCopySnapshotRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelCopySnapshotRequest

	// public void setDestinationRegion(java.lang.String)
	SetDestinationRegion(a string) 

	// public java.lang.String getDestinationRegion()
	GetDestinationRegion() string

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withDestinationRegion(java.lang.String)
	WithDestinationRegion(a string) *ServicesEc2ModelCopySnapshotRequest

	// public void setPresignedUrl(java.lang.String)
	SetPresignedUrl(a string) 

	// public java.lang.String getPresignedUrl()
	GetPresignedUrl() string

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withPresignedUrl(java.lang.String)
	WithPresignedUrl(a string) *ServicesEc2ModelCopySnapshotRequest

	// public void setEncrypted(java.lang.Boolean)
	SetEncrypted(a bool) 

	// public java.lang.Boolean getEncrypted()
	GetEncrypted() bool

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withEncrypted(java.lang.Boolean)
	WithEncrypted(a bool) *ServicesEc2ModelCopySnapshotRequest

	// public java.lang.Boolean isEncrypted()
	IsEncrypted() bool

	// public void setKmsKeyId(java.lang.String)
	SetKmsKeyId(a string) 

	// public java.lang.String getKmsKeyId()
	GetKmsKeyId() string

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest withKmsKeyId(java.lang.String)
	WithKmsKeyId(a string) *ServicesEc2ModelCopySnapshotRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CopySnapshotRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CopySnapshotRequest clone()
	Clone3() *ServicesEc2ModelCopySnapshotRequest
}

type ServicesEc2ModelCopySnapshotRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CopySnapshotRequest()
func NewServicesEc2ModelCopySnapshotRequest() (*ServicesEc2ModelCopySnapshotRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CopySnapshotRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCopySnapshotRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSourceRegion(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetSourceRegion(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceRegion", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSourceRegion()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetSourceRegion() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceRegion", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withSourceRegion(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithSourceRegion(a string) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceRegion", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetSourceSnapshotId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceSnapshotId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSourceSnapshotId()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetSourceSnapshotId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceSnapshotId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withSourceSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithSourceSnapshotId(a string) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceSnapshotId", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithDescription(a string) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDestinationRegion(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetDestinationRegion(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDestinationRegion", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDestinationRegion()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetDestinationRegion() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDestinationRegion", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withDestinationRegion(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithDestinationRegion(a string) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDestinationRegion", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPresignedUrl(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetPresignedUrl(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPresignedUrl", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPresignedUrl()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetPresignedUrl() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPresignedUrl", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withPresignedUrl(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithPresignedUrl(a string) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPresignedUrl", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetEncrypted(a bool)  {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetEncrypted() bool {
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithEncrypted(a bool) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEncrypted", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEncrypted()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) IsEncrypted() bool {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) SetKmsKeyId(a string)  {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetKmsKeyId() string {
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

// public com.amazonaws.services.ec2.model.CopySnapshotRequest withKmsKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotRequest) WithKmsKeyId(a string) *ServicesEc2ModelCopySnapshotRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKmsKeyId", "com/amazonaws/services/ec2/model/CopySnapshotRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CopySnapshotRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CopySnapshotRequest clone()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) Clone3() *ServicesEc2ModelCopySnapshotRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CopySnapshotRequest")
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
	unique_x := &ServicesEc2ModelCopySnapshotRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCopySnapshotRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCopySnapshotRequest) Clone2() (*JavaLangObject, error) {
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


