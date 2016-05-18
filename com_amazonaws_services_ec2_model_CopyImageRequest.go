package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCopyImageRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSourceRegion(java.lang.String)
	SetSourceRegion(a string) 

	// public java.lang.String getSourceRegion()
	GetSourceRegion() string

	// public com.amazonaws.services.ec2.model.CopyImageRequest withSourceRegion(java.lang.String)
	WithSourceRegion(a string) *ServicesEc2ModelCopyImageRequest

	// public void setSourceImageId(java.lang.String)
	SetSourceImageId(a string) 

	// public java.lang.String getSourceImageId()
	GetSourceImageId() string

	// public com.amazonaws.services.ec2.model.CopyImageRequest withSourceImageId(java.lang.String)
	WithSourceImageId(a string) *ServicesEc2ModelCopyImageRequest

	// public void setName(java.lang.String)
	SetName(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.CopyImageRequest withName(java.lang.String)
	WithName(a string) *ServicesEc2ModelCopyImageRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.CopyImageRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelCopyImageRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CopyImageRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCopyImageRequest

	// public void setEncrypted(java.lang.Boolean)
	SetEncrypted(a bool) 

	// public java.lang.Boolean getEncrypted()
	GetEncrypted() bool

	// public com.amazonaws.services.ec2.model.CopyImageRequest withEncrypted(java.lang.Boolean)
	WithEncrypted(a bool) *ServicesEc2ModelCopyImageRequest

	// public java.lang.Boolean isEncrypted()
	IsEncrypted() bool

	// public void setKmsKeyId(java.lang.String)
	SetKmsKeyId(a string) 

	// public java.lang.String getKmsKeyId()
	GetKmsKeyId() string

	// public com.amazonaws.services.ec2.model.CopyImageRequest withKmsKeyId(java.lang.String)
	WithKmsKeyId(a string) *ServicesEc2ModelCopyImageRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CopyImageRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CopyImageRequest clone()
	Clone3() *ServicesEc2ModelCopyImageRequest
}

type ServicesEc2ModelCopyImageRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CopyImageRequest()
func NewServicesEc2ModelCopyImageRequest() (*ServicesEc2ModelCopyImageRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CopyImageRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCopyImageRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSourceRegion(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) SetSourceRegion(a string)  {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) GetSourceRegion() string {
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withSourceRegion(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithSourceRegion(a string) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceRegion", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceImageId(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) SetSourceImageId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceImageId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSourceImageId()
func (jbobject *ServicesEc2ModelCopyImageRequest) GetSourceImageId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceImageId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withSourceImageId(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithSourceImageId(a string) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceImageId", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getName()
func (jbobject *ServicesEc2ModelCopyImageRequest) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withName(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithName(a string) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithDescription(a string) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) SetClientToken(a string)  {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) GetClientToken() string {
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithClientToken(a string) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCopyImageRequest) SetEncrypted(a bool)  {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) GetEncrypted() bool {
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithEncrypted(a bool) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEncrypted", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEncrypted()
func (jbobject *ServicesEc2ModelCopyImageRequest) IsEncrypted() bool {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) SetKmsKeyId(a string)  {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) GetKmsKeyId() string {
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

// public com.amazonaws.services.ec2.model.CopyImageRequest withKmsKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelCopyImageRequest) WithKmsKeyId(a string) *ServicesEc2ModelCopyImageRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKmsKeyId", "com/amazonaws/services/ec2/model/CopyImageRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CopyImageRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCopyImageRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CopyImageRequest clone()
func (jbobject *ServicesEc2ModelCopyImageRequest) Clone3() *ServicesEc2ModelCopyImageRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CopyImageRequest")
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
	unique_x := &ServicesEc2ModelCopyImageRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCopyImageRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCopyImageRequest) Clone2() (*JavaLangObject, error) {
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


