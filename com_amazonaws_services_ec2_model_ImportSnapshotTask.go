package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportSnapshotTaskInterface interface {
	JavaLangObjectInterface

	// public void setImportTaskId(java.lang.String)
	SetImportTaskId(a string) 

	// public java.lang.String getImportTaskId()
	GetImportTaskId() string

	// public com.amazonaws.services.ec2.model.ImportSnapshotTask withImportTaskId(java.lang.String)
	WithImportTaskId(a string) *ServicesEc2ModelImportSnapshotTask

	// public void setSnapshotTaskDetail(com.amazonaws.services.ec2.model.SnapshotTaskDetail)
	SetSnapshotTaskDetail(a ServicesEc2ModelSnapshotTaskDetailInterface) 

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail getSnapshotTaskDetail()
	GetSnapshotTaskDetail() *ServicesEc2ModelSnapshotTaskDetail

	// public com.amazonaws.services.ec2.model.ImportSnapshotTask withSnapshotTaskDetail(com.amazonaws.services.ec2.model.SnapshotTaskDetail)
	WithSnapshotTaskDetail(a ServicesEc2ModelSnapshotTaskDetailInterface) *ServicesEc2ModelImportSnapshotTask

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportSnapshotTask withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportSnapshotTask

	// public com.amazonaws.services.ec2.model.ImportSnapshotTask clone()
	Clone() *ServicesEc2ModelImportSnapshotTask
}

type ServicesEc2ModelImportSnapshotTask struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportSnapshotTask()
func NewServicesEc2ModelImportSnapshotTask() (*ServicesEc2ModelImportSnapshotTask) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportSnapshotTask")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportSnapshotTask{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotTask) SetImportTaskId(a string)  {
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
func (jbobject *ServicesEc2ModelImportSnapshotTask) GetImportTaskId() string {
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

// public com.amazonaws.services.ec2.model.ImportSnapshotTask withImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotTask) WithImportTaskId(a string) *ServicesEc2ModelImportSnapshotTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportTaskId", "com/amazonaws/services/ec2/model/ImportSnapshotTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportSnapshotTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSnapshotTaskDetail(com.amazonaws.services.ec2.model.SnapshotTaskDetail)
func (jbobject *ServicesEc2ModelImportSnapshotTask) SetSnapshotTaskDetail(a ServicesEc2ModelSnapshotTaskDetailInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotTaskDetail", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotTaskDetail"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail getSnapshotTaskDetail()
func (jbobject *ServicesEc2ModelImportSnapshotTask) GetSnapshotTaskDetail() *ServicesEc2ModelSnapshotTaskDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotTaskDetail", "com/amazonaws/services/ec2/model/SnapshotTaskDetail")
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportSnapshotTask withSnapshotTaskDetail(com.amazonaws.services.ec2.model.SnapshotTaskDetail)
func (jbobject *ServicesEc2ModelImportSnapshotTask) WithSnapshotTaskDetail(a ServicesEc2ModelSnapshotTaskDetailInterface) *ServicesEc2ModelImportSnapshotTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotTaskDetail", "com/amazonaws/services/ec2/model/ImportSnapshotTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotTaskDetail"))
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
	unique_x := &ServicesEc2ModelImportSnapshotTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotTask) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelImportSnapshotTask) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportSnapshotTask withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportSnapshotTask) WithDescription(a string) *ServicesEc2ModelImportSnapshotTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportSnapshotTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportSnapshotTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportSnapshotTask) ToString() string {
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
func (jbobject *ServicesEc2ModelImportSnapshotTask) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportSnapshotTask) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportSnapshotTask clone()
func (jbobject *ServicesEc2ModelImportSnapshotTask) Clone() *ServicesEc2ModelImportSnapshotTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportSnapshotTask")
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
	unique_x := &ServicesEc2ModelImportSnapshotTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportSnapshotTask) Clone2() (*JavaLangObject, error) {
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


