package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRestoreAddressToClassicResultInterface interface {
	JavaLangObjectInterface

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelRestoreAddressToClassicResult

	// public void setStatus(com.amazonaws.services.ec2.model.Status)
	SetStatus(a ServicesEc2ModelStatusInterface) 

	// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult withStatus(com.amazonaws.services.ec2.model.Status)
	WithStatus(a ServicesEc2ModelStatusInterface) *ServicesEc2ModelRestoreAddressToClassicResult

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelRestoreAddressToClassicResult

	// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult clone()
	Clone() *ServicesEc2ModelRestoreAddressToClassicResult
}

type ServicesEc2ModelRestoreAddressToClassicResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult()
func NewServicesEc2ModelRestoreAddressToClassicResult() (*ServicesEc2ModelRestoreAddressToClassicResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RestoreAddressToClassicResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRestoreAddressToClassicResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) SetStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) WithStatus2(a string) *ServicesEc2ModelRestoreAddressToClassicResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/RestoreAddressToClassicResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRestoreAddressToClassicResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.Status)
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) SetStatus(a ServicesEc2ModelStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Status"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult withStatus(com.amazonaws.services.ec2.model.Status)
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) WithStatus(a ServicesEc2ModelStatusInterface) *ServicesEc2ModelRestoreAddressToClassicResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/RestoreAddressToClassicResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Status"))
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
	unique_x := &ServicesEc2ModelRestoreAddressToClassicResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) SetPublicIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIp()
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) GetPublicIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) WithPublicIp(a string) *ServicesEc2ModelRestoreAddressToClassicResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/RestoreAddressToClassicResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRestoreAddressToClassicResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) ToString() string {
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
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RestoreAddressToClassicResult clone()
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) Clone() *ServicesEc2ModelRestoreAddressToClassicResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RestoreAddressToClassicResult")
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
	unique_x := &ServicesEc2ModelRestoreAddressToClassicResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRestoreAddressToClassicResult) Clone2() (*JavaLangObject, error) {
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


