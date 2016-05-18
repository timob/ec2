package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateKeyPairResultInterface interface {
	JavaLangObjectInterface

	// public void setKeyPair(com.amazonaws.services.ec2.model.KeyPair)
	SetKeyPair(a ServicesEc2ModelKeyPairInterface) 

	// public com.amazonaws.services.ec2.model.KeyPair getKeyPair()
	GetKeyPair() *ServicesEc2ModelKeyPair

	// public com.amazonaws.services.ec2.model.CreateKeyPairResult withKeyPair(com.amazonaws.services.ec2.model.KeyPair)
	WithKeyPair(a ServicesEc2ModelKeyPairInterface) *ServicesEc2ModelCreateKeyPairResult

	// public com.amazonaws.services.ec2.model.CreateKeyPairResult clone()
	Clone() *ServicesEc2ModelCreateKeyPairResult
}

type ServicesEc2ModelCreateKeyPairResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateKeyPairResult()
func NewServicesEc2ModelCreateKeyPairResult() (*ServicesEc2ModelCreateKeyPairResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateKeyPairResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateKeyPairResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setKeyPair(com.amazonaws.services.ec2.model.KeyPair)
func (jbobject *ServicesEc2ModelCreateKeyPairResult) SetKeyPair(a ServicesEc2ModelKeyPairInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyPair", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/KeyPair"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.KeyPair getKeyPair()
func (jbobject *ServicesEc2ModelCreateKeyPairResult) GetKeyPair() *ServicesEc2ModelKeyPair {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyPair", "com/amazonaws/services/ec2/model/KeyPair")
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

// public com.amazonaws.services.ec2.model.CreateKeyPairResult withKeyPair(com.amazonaws.services.ec2.model.KeyPair)
func (jbobject *ServicesEc2ModelCreateKeyPairResult) WithKeyPair(a ServicesEc2ModelKeyPairInterface) *ServicesEc2ModelCreateKeyPairResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyPair", "com/amazonaws/services/ec2/model/CreateKeyPairResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/KeyPair"))
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
	unique_x := &ServicesEc2ModelCreateKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateKeyPairResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateKeyPairResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateKeyPairResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateKeyPairResult clone()
func (jbobject *ServicesEc2ModelCreateKeyPairResult) Clone() *ServicesEc2ModelCreateKeyPairResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateKeyPairResult")
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
	unique_x := &ServicesEc2ModelCreateKeyPairResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateKeyPairResult) Clone2() (*JavaLangObject, error) {
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


