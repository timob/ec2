package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateInstanceExportTaskResultInterface interface {
	JavaLangObjectInterface

	// public void setExportTask(com.amazonaws.services.ec2.model.ExportTask)
	SetExportTask(a ServicesEc2ModelExportTaskInterface) 

	// public com.amazonaws.services.ec2.model.ExportTask getExportTask()
	GetExportTask() *ServicesEc2ModelExportTask

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult withExportTask(com.amazonaws.services.ec2.model.ExportTask)
	WithExportTask(a ServicesEc2ModelExportTaskInterface) *ServicesEc2ModelCreateInstanceExportTaskResult

	// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult clone()
	Clone() *ServicesEc2ModelCreateInstanceExportTaskResult
}

type ServicesEc2ModelCreateInstanceExportTaskResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult()
func NewServicesEc2ModelCreateInstanceExportTaskResult() (*ServicesEc2ModelCreateInstanceExportTaskResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateInstanceExportTaskResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateInstanceExportTaskResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setExportTask(com.amazonaws.services.ec2.model.ExportTask)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) SetExportTask(a ServicesEc2ModelExportTaskInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExportTask", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportTask"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ExportTask getExportTask()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) GetExportTask() *ServicesEc2ModelExportTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExportTask", "com/amazonaws/services/ec2/model/ExportTask")
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult withExportTask(com.amazonaws.services.ec2.model.ExportTask)
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) WithExportTask(a ServicesEc2ModelExportTaskInterface) *ServicesEc2ModelCreateInstanceExportTaskResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportTask", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportTask"))
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateInstanceExportTaskResult clone()
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) Clone() *ServicesEc2ModelCreateInstanceExportTaskResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateInstanceExportTaskResult")
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
	unique_x := &ServicesEc2ModelCreateInstanceExportTaskResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateInstanceExportTaskResult) Clone2() (*JavaLangObject, error) {
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


