package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelS3StorageInterface interface {
	JavaLangObjectInterface

	// public void setBucket(java.lang.String)
	SetBucket(a string) 

	// public java.lang.String getBucket()
	GetBucket() string

	// public com.amazonaws.services.ec2.model.S3Storage withBucket(java.lang.String)
	WithBucket(a string) *ServicesEc2ModelS3Storage

	// public void setPrefix(java.lang.String)
	SetPrefix(a string) 

	// public java.lang.String getPrefix()
	GetPrefix() string

	// public com.amazonaws.services.ec2.model.S3Storage withPrefix(java.lang.String)
	WithPrefix(a string) *ServicesEc2ModelS3Storage

	// public void setAWSAccessKeyId(java.lang.String)
	SetAWSAccessKeyId(a string) 

	// public java.lang.String getAWSAccessKeyId()
	GetAWSAccessKeyId() string

	// public com.amazonaws.services.ec2.model.S3Storage withAWSAccessKeyId(java.lang.String)
	WithAWSAccessKeyId(a string) *ServicesEc2ModelS3Storage

	// public void setUploadPolicy(java.lang.String)
	SetUploadPolicy(a string) 

	// public java.lang.String getUploadPolicy()
	GetUploadPolicy() string

	// public com.amazonaws.services.ec2.model.S3Storage withUploadPolicy(java.lang.String)
	WithUploadPolicy(a string) *ServicesEc2ModelS3Storage

	// public void setUploadPolicySignature(java.lang.String)
	SetUploadPolicySignature(a string) 

	// public java.lang.String getUploadPolicySignature()
	GetUploadPolicySignature() string

	// public com.amazonaws.services.ec2.model.S3Storage withUploadPolicySignature(java.lang.String)
	WithUploadPolicySignature(a string) *ServicesEc2ModelS3Storage

	// public com.amazonaws.services.ec2.model.S3Storage clone()
	Clone() *ServicesEc2ModelS3Storage
}

type ServicesEc2ModelS3Storage struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.S3Storage()
func NewServicesEc2ModelS3Storage() (*ServicesEc2ModelS3Storage) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/S3Storage")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelS3Storage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setBucket(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) SetBucket(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBucket", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getBucket()
func (jbobject *ServicesEc2ModelS3Storage) GetBucket() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBucket", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.S3Storage withBucket(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) WithBucket(a string) *ServicesEc2ModelS3Storage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBucket", "com/amazonaws/services/ec2/model/S3Storage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelS3Storage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrefix(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) SetPrefix(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefix", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrefix()
func (jbobject *ServicesEc2ModelS3Storage) GetPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefix", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.S3Storage withPrefix(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) WithPrefix(a string) *ServicesEc2ModelS3Storage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefix", "com/amazonaws/services/ec2/model/S3Storage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelS3Storage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAWSAccessKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) SetAWSAccessKeyId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAWSAccessKeyId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAWSAccessKeyId()
func (jbobject *ServicesEc2ModelS3Storage) GetAWSAccessKeyId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAWSAccessKeyId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.S3Storage withAWSAccessKeyId(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) WithAWSAccessKeyId(a string) *ServicesEc2ModelS3Storage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAWSAccessKeyId", "com/amazonaws/services/ec2/model/S3Storage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelS3Storage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUploadPolicy(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) SetUploadPolicy(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUploadPolicy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUploadPolicy()
func (jbobject *ServicesEc2ModelS3Storage) GetUploadPolicy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUploadPolicy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.S3Storage withUploadPolicy(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) WithUploadPolicy(a string) *ServicesEc2ModelS3Storage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUploadPolicy", "com/amazonaws/services/ec2/model/S3Storage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelS3Storage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUploadPolicySignature(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) SetUploadPolicySignature(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUploadPolicySignature", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUploadPolicySignature()
func (jbobject *ServicesEc2ModelS3Storage) GetUploadPolicySignature() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUploadPolicySignature", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.S3Storage withUploadPolicySignature(java.lang.String)
func (jbobject *ServicesEc2ModelS3Storage) WithUploadPolicySignature(a string) *ServicesEc2ModelS3Storage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUploadPolicySignature", "com/amazonaws/services/ec2/model/S3Storage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelS3Storage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelS3Storage) ToString() string {
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
func (jbobject *ServicesEc2ModelS3Storage) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelS3Storage) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.S3Storage clone()
func (jbobject *ServicesEc2ModelS3Storage) Clone() *ServicesEc2ModelS3Storage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/S3Storage")
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
	unique_x := &ServicesEc2ModelS3Storage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelS3Storage) Clone2() (*JavaLangObject, error) {
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


