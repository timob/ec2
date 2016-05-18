package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelUserBucketDetailsInterface interface {
	JavaLangObjectInterface

	// public void setS3Bucket(java.lang.String)
	SetS3Bucket(a string) 

	// public java.lang.String getS3Bucket()
	GetS3Bucket() string

	// public com.amazonaws.services.ec2.model.UserBucketDetails withS3Bucket(java.lang.String)
	WithS3Bucket(a string) *ServicesEc2ModelUserBucketDetails

	// public void setS3Key(java.lang.String)
	SetS3Key(a string) 

	// public java.lang.String getS3Key()
	GetS3Key() string

	// public com.amazonaws.services.ec2.model.UserBucketDetails withS3Key(java.lang.String)
	WithS3Key(a string) *ServicesEc2ModelUserBucketDetails

	// public com.amazonaws.services.ec2.model.UserBucketDetails clone()
	Clone() *ServicesEc2ModelUserBucketDetails
}

type ServicesEc2ModelUserBucketDetails struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.UserBucketDetails()
func NewServicesEc2ModelUserBucketDetails() (*ServicesEc2ModelUserBucketDetails) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/UserBucketDetails")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelUserBucketDetails{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setS3Bucket(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucketDetails) SetS3Bucket(a string)  {
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
func (jbobject *ServicesEc2ModelUserBucketDetails) GetS3Bucket() string {
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

// public com.amazonaws.services.ec2.model.UserBucketDetails withS3Bucket(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucketDetails) WithS3Bucket(a string) *ServicesEc2ModelUserBucketDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withS3Bucket", "com/amazonaws/services/ec2/model/UserBucketDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserBucketDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setS3Key(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucketDetails) SetS3Key(a string)  {
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
func (jbobject *ServicesEc2ModelUserBucketDetails) GetS3Key() string {
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

// public com.amazonaws.services.ec2.model.UserBucketDetails withS3Key(java.lang.String)
func (jbobject *ServicesEc2ModelUserBucketDetails) WithS3Key(a string) *ServicesEc2ModelUserBucketDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withS3Key", "com/amazonaws/services/ec2/model/UserBucketDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUserBucketDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelUserBucketDetails) ToString() string {
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
func (jbobject *ServicesEc2ModelUserBucketDetails) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelUserBucketDetails) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.UserBucketDetails clone()
func (jbobject *ServicesEc2ModelUserBucketDetails) Clone() *ServicesEc2ModelUserBucketDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/UserBucketDetails")
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
	unique_x := &ServicesEc2ModelUserBucketDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelUserBucketDetails) Clone2() (*JavaLangObject, error) {
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


