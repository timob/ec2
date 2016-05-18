package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelKeyPairInfoInterface interface {
	JavaLangObjectInterface

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.KeyPairInfo withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelKeyPairInfo

	// public void setKeyFingerprint(java.lang.String)
	SetKeyFingerprint(a string) 

	// public java.lang.String getKeyFingerprint()
	GetKeyFingerprint() string

	// public com.amazonaws.services.ec2.model.KeyPairInfo withKeyFingerprint(java.lang.String)
	WithKeyFingerprint(a string) *ServicesEc2ModelKeyPairInfo

	// public com.amazonaws.services.ec2.model.KeyPairInfo clone()
	Clone() *ServicesEc2ModelKeyPairInfo
}

type ServicesEc2ModelKeyPairInfo struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.KeyPairInfo()
func NewServicesEc2ModelKeyPairInfo() (*ServicesEc2ModelKeyPairInfo) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/KeyPairInfo")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelKeyPairInfo{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPairInfo) SetKeyName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKeyName()
func (jbobject *ServicesEc2ModelKeyPairInfo) GetKeyName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.KeyPairInfo withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPairInfo) WithKeyName(a string) *ServicesEc2ModelKeyPairInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/KeyPairInfo", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelKeyPairInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyFingerprint(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPairInfo) SetKeyFingerprint(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyFingerprint", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKeyFingerprint()
func (jbobject *ServicesEc2ModelKeyPairInfo) GetKeyFingerprint() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyFingerprint", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.KeyPairInfo withKeyFingerprint(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPairInfo) WithKeyFingerprint(a string) *ServicesEc2ModelKeyPairInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyFingerprint", "com/amazonaws/services/ec2/model/KeyPairInfo", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelKeyPairInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelKeyPairInfo) ToString() string {
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
func (jbobject *ServicesEc2ModelKeyPairInfo) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelKeyPairInfo) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.KeyPairInfo clone()
func (jbobject *ServicesEc2ModelKeyPairInfo) Clone() *ServicesEc2ModelKeyPairInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/KeyPairInfo")
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
	unique_x := &ServicesEc2ModelKeyPairInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelKeyPairInfo) Clone2() (*JavaLangObject, error) {
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


