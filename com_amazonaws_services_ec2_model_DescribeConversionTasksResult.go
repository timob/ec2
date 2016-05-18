package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeConversionTasksResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ConversionTask> getConversionTasks()
	GetConversionTasks() []*ServicesEc2ModelConversionTask

	// public void setConversionTasks(java.util.Collection<com.amazonaws.services.ec2.model.ConversionTask>)
	SetConversionTasks(a []*ServicesEc2ModelConversionTask) 

	// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult withConversionTasks(com.amazonaws.services.ec2.model.ConversionTask...)
	WithConversionTasks(a ...*ServicesEc2ModelConversionTask) *ServicesEc2ModelDescribeConversionTasksResult

	// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult withConversionTasks(java.util.Collection<com.amazonaws.services.ec2.model.ConversionTask>)
	WithConversionTasks2(a []*ServicesEc2ModelConversionTask) *ServicesEc2ModelDescribeConversionTasksResult

	// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult clone()
	Clone() *ServicesEc2ModelDescribeConversionTasksResult
}

type ServicesEc2ModelDescribeConversionTasksResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult()
func NewServicesEc2ModelDescribeConversionTasksResult() (*ServicesEc2ModelDescribeConversionTasksResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeConversionTasksResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeConversionTasksResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ConversionTask> getConversionTasks()
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) GetConversionTasks() []*ServicesEc2ModelConversionTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConversionTasks", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelConversionTask)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setConversionTasks(java.util.Collection<com.amazonaws.services.ec2.model.ConversionTask>)
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) SetConversionTasks(a []*ServicesEc2ModelConversionTask)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConversionTasks", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult withConversionTasks(com.amazonaws.services.ec2.model.ConversionTask...)
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) WithConversionTasks(a ...*ServicesEc2ModelConversionTask) *ServicesEc2ModelDescribeConversionTasksResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ConversionTask")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConversionTasks", "com/amazonaws/services/ec2/model/DescribeConversionTasksResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ConversionTask"))
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
	unique_x := &ServicesEc2ModelDescribeConversionTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult withConversionTasks(java.util.Collection<com.amazonaws.services.ec2.model.ConversionTask>)
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) WithConversionTasks2(a []*ServicesEc2ModelConversionTask) *ServicesEc2ModelDescribeConversionTasksResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConversionTasks", "com/amazonaws/services/ec2/model/DescribeConversionTasksResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeConversionTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeConversionTasksResult clone()
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) Clone() *ServicesEc2ModelDescribeConversionTasksResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeConversionTasksResult")
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
	unique_x := &ServicesEc2ModelDescribeConversionTasksResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeConversionTasksResult) Clone2() (*JavaLangObject, error) {
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


