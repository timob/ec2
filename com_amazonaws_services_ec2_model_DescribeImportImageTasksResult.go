package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeImportImageTasksResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ImportImageTask> getImportImageTasks()
	GetImportImageTasks() []*ServicesEc2ModelImportImageTask

	// public void setImportImageTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportImageTask>)
	SetImportImageTasks(a []*ServicesEc2ModelImportImageTask) 

	// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult withImportImageTasks(com.amazonaws.services.ec2.model.ImportImageTask...)
	WithImportImageTasks(a ...*ServicesEc2ModelImportImageTask) *ServicesEc2ModelDescribeImportImageTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult withImportImageTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportImageTask>)
	WithImportImageTasks2(a []*ServicesEc2ModelImportImageTask) *ServicesEc2ModelDescribeImportImageTasksResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeImportImageTasksResult

	// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult clone()
	Clone() *ServicesEc2ModelDescribeImportImageTasksResult
}

type ServicesEc2ModelDescribeImportImageTasksResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult()
func NewServicesEc2ModelDescribeImportImageTasksResult() (*ServicesEc2ModelDescribeImportImageTasksResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeImportImageTasksResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ImportImageTask> getImportImageTasks()
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) GetImportImageTasks() []*ServicesEc2ModelImportImageTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportImageTasks", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelImportImageTask)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setImportImageTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportImageTask>)
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) SetImportImageTasks(a []*ServicesEc2ModelImportImageTask)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportImageTasks", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult withImportImageTasks(com.amazonaws.services.ec2.model.ImportImageTask...)
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) WithImportImageTasks(a ...*ServicesEc2ModelImportImageTask) *ServicesEc2ModelDescribeImportImageTasksResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ImportImageTask")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportImageTasks", "com/amazonaws/services/ec2/model/DescribeImportImageTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportImageTask"))
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
	unique_x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult withImportImageTasks(java.util.Collection<com.amazonaws.services.ec2.model.ImportImageTask>)
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) WithImportImageTasks2(a []*ServicesEc2ModelImportImageTask) *ServicesEc2ModelDescribeImportImageTasksResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportImageTasks", "com/amazonaws/services/ec2/model/DescribeImportImageTasksResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) WithNextToken(a string) *ServicesEc2ModelDescribeImportImageTasksResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeImportImageTasksResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeImportImageTasksResult clone()
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) Clone() *ServicesEc2ModelDescribeImportImageTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeImportImageTasksResult")
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
	unique_x := &ServicesEc2ModelDescribeImportImageTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeImportImageTasksResult) Clone2() (*JavaLangObject, error) {
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


