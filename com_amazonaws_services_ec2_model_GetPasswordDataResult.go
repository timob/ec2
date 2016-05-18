package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelGetPasswordDataResultInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.GetPasswordDataResult withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelGetPasswordDataResult

	// public void setTimestamp(java.util.Date)
	SetTimestamp(a time.Time) 

	// public java.util.Date getTimestamp()
	GetTimestamp() time.Time

	// public com.amazonaws.services.ec2.model.GetPasswordDataResult withTimestamp(java.util.Date)
	WithTimestamp(a time.Time) *ServicesEc2ModelGetPasswordDataResult

	// public void setPasswordData(java.lang.String)
	SetPasswordData(a string) 

	// public java.lang.String getPasswordData()
	GetPasswordData() string

	// public com.amazonaws.services.ec2.model.GetPasswordDataResult withPasswordData(java.lang.String)
	WithPasswordData(a string) *ServicesEc2ModelGetPasswordDataResult

	// public com.amazonaws.services.ec2.model.GetPasswordDataResult clone()
	Clone() *ServicesEc2ModelGetPasswordDataResult
}

type ServicesEc2ModelGetPasswordDataResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.GetPasswordDataResult()
func NewServicesEc2ModelGetPasswordDataResult() (*ServicesEc2ModelGetPasswordDataResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/GetPasswordDataResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelGetPasswordDataResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelGetPasswordDataResult) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelGetPasswordDataResult) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.GetPasswordDataResult withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelGetPasswordDataResult) WithInstanceId(a string) *ServicesEc2ModelGetPasswordDataResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/GetPasswordDataResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelGetPasswordDataResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelGetPasswordDataResult) SetTimestamp(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTimestamp", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getTimestamp()
func (jbobject *ServicesEc2ModelGetPasswordDataResult) GetTimestamp() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimestamp", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.GetPasswordDataResult withTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelGetPasswordDataResult) WithTimestamp(a time.Time) *ServicesEc2ModelGetPasswordDataResult {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTimestamp", "com/amazonaws/services/ec2/model/GetPasswordDataResult", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelGetPasswordDataResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPasswordData(java.lang.String)
func (jbobject *ServicesEc2ModelGetPasswordDataResult) SetPasswordData(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPasswordData", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPasswordData()
func (jbobject *ServicesEc2ModelGetPasswordDataResult) GetPasswordData() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPasswordData", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.GetPasswordDataResult withPasswordData(java.lang.String)
func (jbobject *ServicesEc2ModelGetPasswordDataResult) WithPasswordData(a string) *ServicesEc2ModelGetPasswordDataResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPasswordData", "com/amazonaws/services/ec2/model/GetPasswordDataResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelGetPasswordDataResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelGetPasswordDataResult) ToString() string {
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
func (jbobject *ServicesEc2ModelGetPasswordDataResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelGetPasswordDataResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.GetPasswordDataResult clone()
func (jbobject *ServicesEc2ModelGetPasswordDataResult) Clone() *ServicesEc2ModelGetPasswordDataResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/GetPasswordDataResult")
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
	unique_x := &ServicesEc2ModelGetPasswordDataResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelGetPasswordDataResult) Clone2() (*JavaLangObject, error) {
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


