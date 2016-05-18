package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeClassicLinkInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ClassicLinkInstance> getInstances()
	GetInstances() []*ServicesEc2ModelClassicLinkInstance

	// public void setInstances(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkInstance>)
	SetInstances(a []*ServicesEc2ModelClassicLinkInstance) 

	// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult withInstances(com.amazonaws.services.ec2.model.ClassicLinkInstance...)
	WithInstances(a ...*ServicesEc2ModelClassicLinkInstance) *ServicesEc2ModelDescribeClassicLinkInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult withInstances(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkInstance>)
	WithInstances2(a []*ServicesEc2ModelClassicLinkInstance) *ServicesEc2ModelDescribeClassicLinkInstancesResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeClassicLinkInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult clone()
	Clone() *ServicesEc2ModelDescribeClassicLinkInstancesResult
}

type ServicesEc2ModelDescribeClassicLinkInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult()
func NewServicesEc2ModelDescribeClassicLinkInstancesResult() (*ServicesEc2ModelDescribeClassicLinkInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ClassicLinkInstance> getInstances()
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) GetInstances() []*ServicesEc2ModelClassicLinkInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstances", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelClassicLinkInstance)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstances(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkInstance>)
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) SetInstances(a []*ServicesEc2ModelClassicLinkInstance)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult withInstances(com.amazonaws.services.ec2.model.ClassicLinkInstance...)
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) WithInstances(a ...*ServicesEc2ModelClassicLinkInstance) *ServicesEc2ModelDescribeClassicLinkInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ClassicLinkInstance")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ClassicLinkInstance"))
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
	unique_x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult withInstances(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkInstance>)
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) WithInstances2(a []*ServicesEc2ModelClassicLinkInstance) *ServicesEc2ModelDescribeClassicLinkInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) WithNextToken(a string) *ServicesEc2ModelDescribeClassicLinkInstancesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeClassicLinkInstancesResult clone()
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) Clone() *ServicesEc2ModelDescribeClassicLinkInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeClassicLinkInstancesResult")
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
	unique_x := &ServicesEc2ModelDescribeClassicLinkInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeClassicLinkInstancesResult) Clone2() (*JavaLangObject, error) {
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


