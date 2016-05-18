package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelExportTaskInterface interface {
	JavaLangObjectInterface

	// public void setExportTaskId(java.lang.String)
	SetExportTaskId(a string) 

	// public java.lang.String getExportTaskId()
	GetExportTaskId() string

	// public com.amazonaws.services.ec2.model.ExportTask withExportTaskId(java.lang.String)
	WithExportTaskId(a string) *ServicesEc2ModelExportTask

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ExportTask withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelExportTask

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.ExportTask withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelExportTask

	// public void setState(com.amazonaws.services.ec2.model.ExportTaskState)
	SetState(a ServicesEc2ModelExportTaskStateInterface) 

	// public com.amazonaws.services.ec2.model.ExportTask withState(com.amazonaws.services.ec2.model.ExportTaskState)
	WithState(a ServicesEc2ModelExportTaskStateInterface) *ServicesEc2ModelExportTask

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.ExportTask withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelExportTask

	// public void setInstanceExportDetails(com.amazonaws.services.ec2.model.InstanceExportDetails)
	SetInstanceExportDetails(a ServicesEc2ModelInstanceExportDetailsInterface) 

	// public com.amazonaws.services.ec2.model.InstanceExportDetails getInstanceExportDetails()
	GetInstanceExportDetails() *ServicesEc2ModelInstanceExportDetails

	// public com.amazonaws.services.ec2.model.ExportTask withInstanceExportDetails(com.amazonaws.services.ec2.model.InstanceExportDetails)
	WithInstanceExportDetails(a ServicesEc2ModelInstanceExportDetailsInterface) *ServicesEc2ModelExportTask

	// public void setExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3Task)
	SetExportToS3Task(a ServicesEc2ModelExportToS3TaskInterface) 

	// public com.amazonaws.services.ec2.model.ExportToS3Task getExportToS3Task()
	GetExportToS3Task() *ServicesEc2ModelExportToS3Task

	// public com.amazonaws.services.ec2.model.ExportTask withExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3Task)
	WithExportToS3Task(a ServicesEc2ModelExportToS3TaskInterface) *ServicesEc2ModelExportTask

	// public com.amazonaws.services.ec2.model.ExportTask clone()
	Clone() *ServicesEc2ModelExportTask
}

type ServicesEc2ModelExportTask struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ExportTask()
func NewServicesEc2ModelExportTask() (*ServicesEc2ModelExportTask) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ExportTask")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelExportTask{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setExportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) SetExportTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExportTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getExportTaskId()
func (jbobject *ServicesEc2ModelExportTask) GetExportTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExportTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ExportTask withExportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) WithExportTaskId(a string) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportTaskId", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelExportTask) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ExportTask withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) WithDescription(a string) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelExportTask) GetState() string {
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

// public com.amazonaws.services.ec2.model.ExportTask withState(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) WithState2(a string) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.ExportTaskState)
func (jbobject *ServicesEc2ModelExportTask) SetState(a ServicesEc2ModelExportTaskStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportTaskState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ExportTask withState(com.amazonaws.services.ec2.model.ExportTaskState)
func (jbobject *ServicesEc2ModelExportTask) WithState(a ServicesEc2ModelExportTaskStateInterface) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportTaskState"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) SetStatusMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusMessage()
func (jbobject *ServicesEc2ModelExportTask) GetStatusMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ExportTask withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelExportTask) WithStatusMessage(a string) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceExportDetails(com.amazonaws.services.ec2.model.InstanceExportDetails)
func (jbobject *ServicesEc2ModelExportTask) SetInstanceExportDetails(a ServicesEc2ModelInstanceExportDetailsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceExportDetails", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceExportDetails"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceExportDetails getInstanceExportDetails()
func (jbobject *ServicesEc2ModelExportTask) GetInstanceExportDetails() *ServicesEc2ModelInstanceExportDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceExportDetails", "com/amazonaws/services/ec2/model/InstanceExportDetails")
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
	unique_x := &ServicesEc2ModelInstanceExportDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ExportTask withInstanceExportDetails(com.amazonaws.services.ec2.model.InstanceExportDetails)
func (jbobject *ServicesEc2ModelExportTask) WithInstanceExportDetails(a ServicesEc2ModelInstanceExportDetailsInterface) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceExportDetails", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceExportDetails"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3Task)
func (jbobject *ServicesEc2ModelExportTask) SetExportToS3Task(a ServicesEc2ModelExportToS3TaskInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExportToS3Task", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportToS3Task"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ExportToS3Task getExportToS3Task()
func (jbobject *ServicesEc2ModelExportTask) GetExportToS3Task() *ServicesEc2ModelExportToS3Task {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExportToS3Task", "com/amazonaws/services/ec2/model/ExportToS3Task")
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
	unique_x := &ServicesEc2ModelExportToS3Task{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ExportTask withExportToS3Task(com.amazonaws.services.ec2.model.ExportToS3Task)
func (jbobject *ServicesEc2ModelExportTask) WithExportToS3Task(a ServicesEc2ModelExportToS3TaskInterface) *ServicesEc2ModelExportTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportToS3Task", "com/amazonaws/services/ec2/model/ExportTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportToS3Task"))
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
	unique_x := &ServicesEc2ModelExportTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelExportTask) ToString() string {
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
func (jbobject *ServicesEc2ModelExportTask) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelExportTask) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ExportTask clone()
func (jbobject *ServicesEc2ModelExportTask) Clone() *ServicesEc2ModelExportTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ExportTask")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelExportTask) Clone2() (*JavaLangObject, error) {
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


