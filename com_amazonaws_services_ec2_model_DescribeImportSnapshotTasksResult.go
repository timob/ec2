package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeImportSnapshotTasksResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ImportSnapshotTask> getImportSnapshotTasks()
	GetImportSnapshotTasks() []*ServicesEc2ModelImportSnapshotTask

	// public void setImportSnapshotTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportSnapshotTask>)
	SetImportSnapshotTasks(a []*ServicesEc2ModelImportSnapshotTask) 

	// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult withImportSnapshotTasks(com.amazonaws.services.ec2.model.ImportSnapshotTask...)
	WithImportSnapshotTasks(a ...*ServicesEc2ModelImportSnapshotTask) *ServicesEc2ModelDescribeImportSnapshotTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult withImportSnapshotTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportSnapshotTask>)
	WithImportSnapshotTasks2(a []*ServicesEc2ModelImportSnapshotTask) *ServicesEc2ModelDescribeImportSnapshotTasksResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeImportSnapshotTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult clone()
	Clone() *ServicesEc2ModelDescribeImportSnapshotTasksResult
}

type ServicesEc2ModelDescribeImportSnapshotTasksResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult()
func NewServicesEc2ModelDescribeImportSnapshotTasksResult() (*ServicesEc2ModelDescribeImportSnapshotTasksResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ImportSnapshotTask> getImportSnapshotTasks()
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) GetImportSnapshotTasks() []*ServicesEc2ModelImportSnapshotTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportSnapshotTasks", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelImportSnapshotTask)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setImportSnapshotTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportSnapshotTask>)
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) SetImportSnapshotTasks(a []*ServicesEc2ModelImportSnapshotTask)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportSnapshotTasks", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult withImportSnapshotTasks(com.amazonaws.services.ec2.model.ImportSnapshotTask...)
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) WithImportSnapshotTasks(a ...*ServicesEc2ModelImportSnapshotTask) *ServicesEc2ModelDescribeImportSnapshotTasksResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ImportSnapshotTask")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportSnapshotTasks", "com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportSnapshotTask"))
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
	unique_x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult withImportSnapshotTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportSnapshotTask>)
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) WithImportSnapshotTasks2(a []*ServicesEc2ModelImportSnapshotTask) *ServicesEc2ModelDescribeImportSnapshotTasksResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportSnapshotTasks", "com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) SetNextToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNextToken()
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) GetNextToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) WithNextToken(a string) *ServicesEc2ModelDescribeImportSnapshotTasksResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeImportSnapshotTasksResult clone()
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) Clone() *ServicesEc2ModelDescribeImportSnapshotTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeImportSnapshotTasksResult")
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
	unique_x := &ServicesEc2ModelDescribeImportSnapshotTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeImportSnapshotTasksResult) Clone2() (*JavaLangObject, error) {
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


