package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelGetConsoleOutputResultInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.GetConsoleOutputResult withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelGetConsoleOutputResult

	// public void setTimestamp(java.util.Date)
	SetTimestamp(a time.Time) 

	// public java.util.Date getTimestamp()
	GetTimestamp() time.Time

	// public com.amazonaws.services.ec2.model.GetConsoleOutputResult withTimestamp(java.util.Date)
	WithTimestamp(a time.Time) *ServicesEc2ModelGetConsoleOutputResult

	// public void setOutput(java.lang.String)
	SetOutput(a string) 

	// public java.lang.String getOutput()
	GetOutput() string

	// public com.amazonaws.services.ec2.model.GetConsoleOutputResult withOutput(java.lang.String)
	WithOutput(a string) *ServicesEc2ModelGetConsoleOutputResult

	// public java.lang.String getDecodedOutput()
	GetDecodedOutput() string

	// public com.amazonaws.services.ec2.model.GetConsoleOutputResult clone()
	Clone() *ServicesEc2ModelGetConsoleOutputResult
}

type ServicesEc2ModelGetConsoleOutputResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.GetConsoleOutputResult()
func NewServicesEc2ModelGetConsoleOutputResult() (*ServicesEc2ModelGetConsoleOutputResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/GetConsoleOutputResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelGetConsoleOutputResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) SetInstanceId(a string)  {
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
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) GetInstanceId() string {
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

// public com.amazonaws.services.ec2.model.GetConsoleOutputResult withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) WithInstanceId(a string) *ServicesEc2ModelGetConsoleOutputResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/GetConsoleOutputResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelGetConsoleOutputResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) SetTimestamp(a time.Time)  {
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
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) GetTimestamp() time.Time {
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

// public com.amazonaws.services.ec2.model.GetConsoleOutputResult withTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) WithTimestamp(a time.Time) *ServicesEc2ModelGetConsoleOutputResult {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTimestamp", "com/amazonaws/services/ec2/model/GetConsoleOutputResult", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelGetConsoleOutputResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOutput(java.lang.String)
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) SetOutput(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOutput", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOutput()
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) GetOutput() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOutput", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.GetConsoleOutputResult withOutput(java.lang.String)
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) WithOutput(a string) *ServicesEc2ModelGetConsoleOutputResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOutput", "com/amazonaws/services/ec2/model/GetConsoleOutputResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelGetConsoleOutputResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getDecodedOutput()
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) GetDecodedOutput() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDecodedOutput", "java/lang/String")
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

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) ToString() string {
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
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.GetConsoleOutputResult clone()
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) Clone() *ServicesEc2ModelGetConsoleOutputResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/GetConsoleOutputResult")
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
	unique_x := &ServicesEc2ModelGetConsoleOutputResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelGetConsoleOutputResult) Clone2() (*JavaLangObject, error) {
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


