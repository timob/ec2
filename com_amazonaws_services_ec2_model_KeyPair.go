package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelKeyPairInterface interface {
	JavaLangObjectInterface

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.KeyPair withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelKeyPair

	// public void setKeyFingerprint(java.lang.String)
	SetKeyFingerprint(a string) 

	// public java.lang.String getKeyFingerprint()
	GetKeyFingerprint() string

	// public com.amazonaws.services.ec2.model.KeyPair withKeyFingerprint(java.lang.String)
	WithKeyFingerprint(a string) *ServicesEc2ModelKeyPair

	// public void setKeyMaterial(java.lang.String)
	SetKeyMaterial(a string) 

	// public java.lang.String getKeyMaterial()
	GetKeyMaterial() string

	// public com.amazonaws.services.ec2.model.KeyPair withKeyMaterial(java.lang.String)
	WithKeyMaterial(a string) *ServicesEc2ModelKeyPair

	// public com.amazonaws.services.ec2.model.KeyPair clone()
	Clone() *ServicesEc2ModelKeyPair
}

type ServicesEc2ModelKeyPair struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.KeyPair()
func NewServicesEc2ModelKeyPair() (*ServicesEc2ModelKeyPair) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/KeyPair")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelKeyPair{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPair) SetKeyName(a string)  {
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
func (jbobject *ServicesEc2ModelKeyPair) GetKeyName() string {
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

// public com.amazonaws.services.ec2.model.KeyPair withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPair) WithKeyName(a string) *ServicesEc2ModelKeyPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/KeyPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelKeyPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyFingerprint(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPair) SetKeyFingerprint(a string)  {
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
func (jbobject *ServicesEc2ModelKeyPair) GetKeyFingerprint() string {
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

// public com.amazonaws.services.ec2.model.KeyPair withKeyFingerprint(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPair) WithKeyFingerprint(a string) *ServicesEc2ModelKeyPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyFingerprint", "com/amazonaws/services/ec2/model/KeyPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelKeyPair{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyMaterial(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPair) SetKeyMaterial(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyMaterial", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKeyMaterial()
func (jbobject *ServicesEc2ModelKeyPair) GetKeyMaterial() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyMaterial", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.KeyPair withKeyMaterial(java.lang.String)
func (jbobject *ServicesEc2ModelKeyPair) WithKeyMaterial(a string) *ServicesEc2ModelKeyPair {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyMaterial", "com/amazonaws/services/ec2/model/KeyPair", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelKeyPair{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelKeyPair) ToString() string {
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
func (jbobject *ServicesEc2ModelKeyPair) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelKeyPair) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.KeyPair clone()
func (jbobject *ServicesEc2ModelKeyPair) Clone() *ServicesEc2ModelKeyPair {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/KeyPair")
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
	unique_x := &ServicesEc2ModelKeyPair{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelKeyPair) Clone2() (*JavaLangObject, error) {
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


