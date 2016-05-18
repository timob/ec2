package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelImportTaskResultInterface interface {
	JavaLangObjectInterface

	// public void setImportTaskId(java.lang.String)
	SetImportTaskId(a string) 

	// public java.lang.String getImportTaskId()
	GetImportTaskId() string

	// public com.amazonaws.services.ec2.model.CancelImportTaskResult withImportTaskId(java.lang.String)
	WithImportTaskId(a string) *ServicesEc2ModelCancelImportTaskResult

	// public void setState(java.lang.String)
	SetState(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.CancelImportTaskResult withState(java.lang.String)
	WithState(a string) *ServicesEc2ModelCancelImportTaskResult

	// public void setPreviousState(java.lang.String)
	SetPreviousState(a string) 

	// public java.lang.String getPreviousState()
	GetPreviousState() string

	// public com.amazonaws.services.ec2.model.CancelImportTaskResult withPreviousState(java.lang.String)
	WithPreviousState(a string) *ServicesEc2ModelCancelImportTaskResult

	// public com.amazonaws.services.ec2.model.CancelImportTaskResult clone()
	Clone() *ServicesEc2ModelCancelImportTaskResult
}

type ServicesEc2ModelCancelImportTaskResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelImportTaskResult()
func NewServicesEc2ModelCancelImportTaskResult() (*ServicesEc2ModelCancelImportTaskResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelImportTaskResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelImportTaskResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskResult) SetImportTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImportTaskId()
func (jbobject *ServicesEc2ModelCancelImportTaskResult) GetImportTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelImportTaskResult withImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskResult) WithImportTaskId(a string) *ServicesEc2ModelCancelImportTaskResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportTaskId", "com/amazonaws/services/ec2/model/CancelImportTaskResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelImportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskResult) SetState(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelCancelImportTaskResult) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelImportTaskResult withState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskResult) WithState(a string) *ServicesEc2ModelCancelImportTaskResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/CancelImportTaskResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelImportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPreviousState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskResult) SetPreviousState(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreviousState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPreviousState()
func (jbobject *ServicesEc2ModelCancelImportTaskResult) GetPreviousState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPreviousState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelImportTaskResult withPreviousState(java.lang.String)
func (jbobject *ServicesEc2ModelCancelImportTaskResult) WithPreviousState(a string) *ServicesEc2ModelCancelImportTaskResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPreviousState", "com/amazonaws/services/ec2/model/CancelImportTaskResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelImportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelImportTaskResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelImportTaskResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelImportTaskResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelImportTaskResult clone()
func (jbobject *ServicesEc2ModelCancelImportTaskResult) Clone() *ServicesEc2ModelCancelImportTaskResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelImportTaskResult")
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
	unique_x := &ServicesEc2ModelCancelImportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelImportTaskResult) Clone2() (*JavaLangObject, error) {
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


