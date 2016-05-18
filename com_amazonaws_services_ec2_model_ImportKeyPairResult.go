package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportKeyPairResultInterface interface {
	JavaLangObjectInterface

	// public void setKeyName(java.lang.String)
	SetKeyName(a string) 

	// public java.lang.String getKeyName()
	GetKeyName() string

	// public com.amazonaws.services.ec2.model.ImportKeyPairResult withKeyName(java.lang.String)
	WithKeyName(a string) *ServicesEc2ModelImportKeyPairResult

	// public void setKeyFingerprint(java.lang.String)
	SetKeyFingerprint(a string) 

	// public java.lang.String getKeyFingerprint()
	GetKeyFingerprint() string

	// public com.amazonaws.services.ec2.model.ImportKeyPairResult withKeyFingerprint(java.lang.String)
	WithKeyFingerprint(a string) *ServicesEc2ModelImportKeyPairResult

	// public com.amazonaws.services.ec2.model.ImportKeyPairResult clone()
	Clone() *ServicesEc2ModelImportKeyPairResult
}

type ServicesEc2ModelImportKeyPairResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportKeyPairResult()
func NewServicesEc2ModelImportKeyPairResult() (*ServicesEc2ModelImportKeyPairResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportKeyPairResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportKeyPairResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelImportKeyPairResult) SetKeyName(a string)  {
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
func (jbobject *ServicesEc2ModelImportKeyPairResult) GetKeyName() string {
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

// public com.amazonaws.services.ec2.model.ImportKeyPairResult withKeyName(java.lang.String)
func (jbobject *ServicesEc2ModelImportKeyPairResult) WithKeyName(a string) *ServicesEc2ModelImportKeyPairResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyName", "com/amazonaws/services/ec2/model/ImportKeyPairResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKeyFingerprint(java.lang.String)
func (jbobject *ServicesEc2ModelImportKeyPairResult) SetKeyFingerprint(a string)  {
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
func (jbobject *ServicesEc2ModelImportKeyPairResult) GetKeyFingerprint() string {
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

// public com.amazonaws.services.ec2.model.ImportKeyPairResult withKeyFingerprint(java.lang.String)
func (jbobject *ServicesEc2ModelImportKeyPairResult) WithKeyFingerprint(a string) *ServicesEc2ModelImportKeyPairResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyFingerprint", "com/amazonaws/services/ec2/model/ImportKeyPairResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportKeyPairResult) ToString() string {
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
func (jbobject *ServicesEc2ModelImportKeyPairResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportKeyPairResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportKeyPairResult clone()
func (jbobject *ServicesEc2ModelImportKeyPairResult) Clone() *ServicesEc2ModelImportKeyPairResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportKeyPairResult")
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
	unique_x := &ServicesEc2ModelImportKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportKeyPairResult) Clone2() (*JavaLangObject, error) {
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


