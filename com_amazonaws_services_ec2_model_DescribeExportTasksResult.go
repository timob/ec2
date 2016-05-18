package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeExportTasksResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ExportTask> getExportTasks()
	GetExportTasks() []*ServicesEc2ModelExportTask

	// public void setExportTasks(java.util.Collection<com.amazonaws.services.ec2.model.ExportTask>)
	SetExportTasks(a []*ServicesEc2ModelExportTask) 

	// public com.amazonaws.services.ec2.model.DescribeExportTasksResult withExportTasks(com.amazonaws.services.ec2.model.ExportTask...)
	WithExportTasks(a ...*ServicesEc2ModelExportTask) *ServicesEc2ModelDescribeExportTasksResult

	// public com.amazonaws.services.ec2.model.DescribeExportTasksResult withExportTasks(java.util.Collection<com.amazonaws.services.ec2.model.ExportTask>)
	WithExportTasks2(a []*ServicesEc2ModelExportTask) *ServicesEc2ModelDescribeExportTasksResult

	// public com.amazonaws.services.ec2.model.DescribeExportTasksResult clone()
	Clone() *ServicesEc2ModelDescribeExportTasksResult
}

type ServicesEc2ModelDescribeExportTasksResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeExportTasksResult()
func NewServicesEc2ModelDescribeExportTasksResult() (*ServicesEc2ModelDescribeExportTasksResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeExportTasksResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeExportTasksResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ExportTask> getExportTasks()
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) GetExportTasks() []*ServicesEc2ModelExportTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExportTasks", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelExportTask)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setExportTasks(java.util.Collection<com.amazonaws.services.ec2.model.ExportTask>)
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) SetExportTasks(a []*ServicesEc2ModelExportTask)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExportTasks", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeExportTasksResult withExportTasks(com.amazonaws.services.ec2.model.ExportTask...)
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) WithExportTasks(a ...*ServicesEc2ModelExportTask) *ServicesEc2ModelDescribeExportTasksResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ExportTask")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportTasks", "com/amazonaws/services/ec2/model/DescribeExportTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ExportTask"))
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
	unique_x := &ServicesEc2ModelDescribeExportTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeExportTasksResult withExportTasks(java.util.Collection<com.amazonaws.services.ec2.model.ExportTask>)
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) WithExportTasks2(a []*ServicesEc2ModelExportTask) *ServicesEc2ModelDescribeExportTasksResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExportTasks", "com/amazonaws/services/ec2/model/DescribeExportTasksResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeExportTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeExportTasksResult clone()
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) Clone() *ServicesEc2ModelDescribeExportTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeExportTasksResult")
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
	unique_x := &ServicesEc2ModelDescribeExportTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeExportTasksResult) Clone2() (*JavaLangObject, error) {
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


