package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeScheduledInstancesResultInterface interface {
	JavaLangObjectInterface

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeScheduledInstancesResult

	// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstance> getScheduledInstanceSet()
	GetScheduledInstanceSet() []*ServicesEc2ModelScheduledInstance

	// public void setScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
	SetScheduledInstanceSet(a []*ServicesEc2ModelScheduledInstance) 

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult withScheduledInstanceSet(com.amazonaws.services.ec2.model.ScheduledInstance...)
	WithScheduledInstanceSet(a ...*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelDescribeScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult withScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
	WithScheduledInstanceSet2(a []*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelDescribeScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult clone()
	Clone() *ServicesEc2ModelDescribeScheduledInstancesResult
}

type ServicesEc2ModelDescribeScheduledInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult()
func NewServicesEc2ModelDescribeScheduledInstancesResult() (*ServicesEc2ModelDescribeScheduledInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeScheduledInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeScheduledInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) WithNextToken(a string) *ServicesEc2ModelDescribeScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeScheduledInstancesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstance> getScheduledInstanceSet()
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) GetScheduledInstanceSet() []*ServicesEc2ModelScheduledInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScheduledInstanceSet", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelScheduledInstance)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) SetScheduledInstanceSet(a []*ServicesEc2ModelScheduledInstance)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScheduledInstanceSet", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult withScheduledInstanceSet(com.amazonaws.services.ec2.model.ScheduledInstance...)
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) WithScheduledInstanceSet(a ...*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelDescribeScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ScheduledInstance")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceSet", "com/amazonaws/services/ec2/model/DescribeScheduledInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstance"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult withScheduledInstanceSet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstance>)
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) WithScheduledInstanceSet2(a []*ServicesEc2ModelScheduledInstance) *ServicesEc2ModelDescribeScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceSet", "com/amazonaws/services/ec2/model/DescribeScheduledInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstancesResult clone()
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) Clone() *ServicesEc2ModelDescribeScheduledInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeScheduledInstancesResult")
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeScheduledInstancesResult) Clone2() (*JavaLangObject, error) {
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


