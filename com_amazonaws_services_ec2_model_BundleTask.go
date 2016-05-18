package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelBundleTaskInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.BundleTask withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelBundleTask

	// public void setBundleId(java.lang.String)
	SetBundleId(a string) 

	// public java.lang.String getBundleId()
	GetBundleId() string

	// public com.amazonaws.services.ec2.model.BundleTask withBundleId(java.lang.String)
	WithBundleId(a string) *ServicesEc2ModelBundleTask

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.BundleTask withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelBundleTask

	// public void setState(com.amazonaws.services.ec2.model.BundleTaskState)
	SetState(a ServicesEc2ModelBundleTaskStateInterface) 

	// public com.amazonaws.services.ec2.model.BundleTask withState(com.amazonaws.services.ec2.model.BundleTaskState)
	WithState(a ServicesEc2ModelBundleTaskStateInterface) *ServicesEc2ModelBundleTask

	// public void setStartTime(java.util.Date)
	SetStartTime(a time.Time) 

	// public java.util.Date getStartTime()
	GetStartTime() time.Time

	// public com.amazonaws.services.ec2.model.BundleTask withStartTime(java.util.Date)
	WithStartTime(a time.Time) *ServicesEc2ModelBundleTask

	// public void setUpdateTime(java.util.Date)
	SetUpdateTime(a time.Time) 

	// public java.util.Date getUpdateTime()
	GetUpdateTime() time.Time

	// public com.amazonaws.services.ec2.model.BundleTask withUpdateTime(java.util.Date)
	WithUpdateTime(a time.Time) *ServicesEc2ModelBundleTask

	// public void setStorage(com.amazonaws.services.ec2.model.Storage)
	SetStorage(a ServicesEc2ModelStorageInterface) 

	// public com.amazonaws.services.ec2.model.Storage getStorage()
	GetStorage() *ServicesEc2ModelStorage

	// public com.amazonaws.services.ec2.model.BundleTask withStorage(com.amazonaws.services.ec2.model.Storage)
	WithStorage(a ServicesEc2ModelStorageInterface) *ServicesEc2ModelBundleTask

	// public void setProgress(java.lang.String)
	SetProgress(a string) 

	// public java.lang.String getProgress()
	GetProgress() string

	// public com.amazonaws.services.ec2.model.BundleTask withProgress(java.lang.String)
	WithProgress(a string) *ServicesEc2ModelBundleTask

	// public void setBundleTaskError(com.amazonaws.services.ec2.model.BundleTaskError)
	SetBundleTaskError(a ServicesEc2ModelBundleTaskErrorInterface) 

	// public com.amazonaws.services.ec2.model.BundleTaskError getBundleTaskError()
	GetBundleTaskError() *ServicesEc2ModelBundleTaskError

	// public com.amazonaws.services.ec2.model.BundleTask withBundleTaskError(com.amazonaws.services.ec2.model.BundleTaskError)
	WithBundleTaskError(a ServicesEc2ModelBundleTaskErrorInterface) *ServicesEc2ModelBundleTask

	// public com.amazonaws.services.ec2.model.BundleTask clone()
	Clone() *ServicesEc2ModelBundleTask
}

type ServicesEc2ModelBundleTask struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.BundleTask()
func NewServicesEc2ModelBundleTask() (*ServicesEc2ModelBundleTask) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/BundleTask")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelBundleTask{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) SetInstanceId(a string)  {
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
func (jbobject *ServicesEc2ModelBundleTask) GetInstanceId() string {
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

// public com.amazonaws.services.ec2.model.BundleTask withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) WithInstanceId(a string) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBundleId(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) SetBundleId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBundleId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getBundleId()
func (jbobject *ServicesEc2ModelBundleTask) GetBundleId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBundleId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.BundleTask withBundleId(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) WithBundleId(a string) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBundleId", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelBundleTask) GetState() string {
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

// public com.amazonaws.services.ec2.model.BundleTask withState(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) WithState2(a string) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.BundleTaskState)
func (jbobject *ServicesEc2ModelBundleTask) SetState(a ServicesEc2ModelBundleTaskStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleTaskState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.BundleTask withState(com.amazonaws.services.ec2.model.BundleTaskState)
func (jbobject *ServicesEc2ModelBundleTask) WithState(a ServicesEc2ModelBundleTaskStateInterface) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleTaskState"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelBundleTask) SetStartTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStartTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getStartTime()
func (jbobject *ServicesEc2ModelBundleTask) GetStartTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.BundleTask withStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelBundleTask) WithStartTime(a time.Time) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStartTime", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUpdateTime(java.util.Date)
func (jbobject *ServicesEc2ModelBundleTask) SetUpdateTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUpdateTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getUpdateTime()
func (jbobject *ServicesEc2ModelBundleTask) GetUpdateTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUpdateTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.BundleTask withUpdateTime(java.util.Date)
func (jbobject *ServicesEc2ModelBundleTask) WithUpdateTime(a time.Time) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUpdateTime", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStorage(com.amazonaws.services.ec2.model.Storage)
func (jbobject *ServicesEc2ModelBundleTask) SetStorage(a ServicesEc2ModelStorageInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStorage", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Storage"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Storage getStorage()
func (jbobject *ServicesEc2ModelBundleTask) GetStorage() *ServicesEc2ModelStorage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStorage", "com/amazonaws/services/ec2/model/Storage")
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
	unique_x := &ServicesEc2ModelStorage{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.BundleTask withStorage(com.amazonaws.services.ec2.model.Storage)
func (jbobject *ServicesEc2ModelBundleTask) WithStorage(a ServicesEc2ModelStorageInterface) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStorage", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Storage"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProgress(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) SetProgress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProgress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProgress()
func (jbobject *ServicesEc2ModelBundleTask) GetProgress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProgress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.BundleTask withProgress(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTask) WithProgress(a string) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProgress", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBundleTaskError(com.amazonaws.services.ec2.model.BundleTaskError)
func (jbobject *ServicesEc2ModelBundleTask) SetBundleTaskError(a ServicesEc2ModelBundleTaskErrorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBundleTaskError", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleTaskError"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.BundleTaskError getBundleTaskError()
func (jbobject *ServicesEc2ModelBundleTask) GetBundleTaskError() *ServicesEc2ModelBundleTaskError {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBundleTaskError", "com/amazonaws/services/ec2/model/BundleTaskError")
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
	unique_x := &ServicesEc2ModelBundleTaskError{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.BundleTask withBundleTaskError(com.amazonaws.services.ec2.model.BundleTaskError)
func (jbobject *ServicesEc2ModelBundleTask) WithBundleTaskError(a ServicesEc2ModelBundleTaskErrorInterface) *ServicesEc2ModelBundleTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBundleTaskError", "com/amazonaws/services/ec2/model/BundleTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BundleTaskError"))
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelBundleTask) ToString() string {
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
func (jbobject *ServicesEc2ModelBundleTask) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelBundleTask) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.BundleTask clone()
func (jbobject *ServicesEc2ModelBundleTask) Clone() *ServicesEc2ModelBundleTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/BundleTask")
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
	unique_x := &ServicesEc2ModelBundleTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelBundleTask) Clone2() (*JavaLangObject, error) {
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


