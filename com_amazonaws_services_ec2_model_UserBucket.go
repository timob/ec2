package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelUserBucketInterface interface {
	JavaLangObjectInterface

	// public void setS3Bucket(java.lang.String)
	SetS3Bucket(a string) 

	// public java.lang.String getS3Bucket()
	GetS3Bucket() string

	// public com.amazonaws.services.ec2.model.UserBucket withS3Bucket(java.lang.String)
	WithS3Bucket(a string) *ServicesEc2ModelUserBucket

	// public void setS3Key(java.lang.String)
	SetS3Key(a string) 

	// public java.lang.String getS3Key()
	GetS3Key() string

	// public com.amazonaws.services.ec2.model.UserBucket withS3Key(java.lang.String)
	WithS3Key(a string) *ServicesEc2ModelUserBucket

	// public com.amazonaws.services.ec2.model.UserBucket clone()
	Clone() *ServicesEc2ModelUserBucket
}

type ServicesEc2ModelUserBucket struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.UserBucket()
func NewServicesEc2ModelUserBucket() (*ServicesEc2ModelUserBucket) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/UserBucket")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelUserBucket{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setS3Bucket(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucket) SetS3Bucket(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setS3Bucket", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getS3Bucket()
func (jbobject *ServicesEc2ModelUserBucket) GetS3Bucket() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getS3Bucket", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserBucket withS3Bucket(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucket) WithS3Bucket(a string) *ServicesEc2ModelUserBucket {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withS3Bucket", "com/amazonaws/services/ec2/model/UserBucket", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserBucket{}
	unique_x.Callable = dst
	return unique_x
}

// public void setS3Key(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucket) SetS3Key(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setS3Key", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getS3Key()
func (jbobject *ServicesEc2ModelUserBucket) GetS3Key() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getS3Key", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UserBucket withS3Key(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucket) WithS3Key(a string) *ServicesEc2ModelUserBucket {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withS3Key", "com/amazonaws/services/ec2/model/UserBucket", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserBucket{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelUserBucket) ToString() string {
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
func (jbobject *ServicesEc2ModelUserBucket) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelUserBucket) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.UserBucket clone()
func (jbobject *ServicesEc2ModelUserBucket) Clone() *ServicesEc2ModelUserBucket {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/UserBucket")
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
	unique_x := &ServicesEc2ModelUserBucket{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelUserBucket) Clone2() (*JavaLangObject, error) {
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


