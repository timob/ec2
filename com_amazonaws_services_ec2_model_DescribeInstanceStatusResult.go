package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeInstanceStatusResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceStatus> getInstanceStatuses()
	GetInstanceStatuses() []*ServicesEc2ModelInstanceStatus

	// public void setInstanceStatuses(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatus>)
	SetInstanceStatuses(a []*ServicesEc2ModelInstanceStatus) 

	// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult withInstanceStatuses(com.amazonaws.services.ec2.model.InstanceStatus...)
	WithInstanceStatuses(a ...*ServicesEc2ModelInstanceStatus) *ServicesEc2ModelDescribeInstanceStatusResult

	// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult withInstanceStatuses(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatus>)
	WithInstanceStatuses2(a []*ServicesEc2ModelInstanceStatus) *ServicesEc2ModelDescribeInstanceStatusResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeInstanceStatusResult

	// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult clone()
	Clone() *ServicesEc2ModelDescribeInstanceStatusResult
}

type ServicesEc2ModelDescribeInstanceStatusResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult()
func NewServicesEc2ModelDescribeInstanceStatusResult() (*ServicesEc2ModelDescribeInstanceStatusResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeInstanceStatusResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceStatus> getInstanceStatuses()
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) GetInstanceStatuses() []*ServicesEc2ModelInstanceStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceStatuses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstanceStatuses(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatus>)
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) SetInstanceStatuses(a []*ServicesEc2ModelInstanceStatus)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceStatuses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult withInstanceStatuses(com.amazonaws.services.ec2.model.InstanceStatus...)
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) WithInstanceStatuses(a ...*ServicesEc2ModelInstanceStatus) *ServicesEc2ModelDescribeInstanceStatusResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceStatus")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceStatuses", "com/amazonaws/services/ec2/model/DescribeInstanceStatusResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatus"))
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
	unique_x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult withInstanceStatuses(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatus>)
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) WithInstanceStatuses2(a []*ServicesEc2ModelInstanceStatus) *ServicesEc2ModelDescribeInstanceStatusResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceStatuses", "com/amazonaws/services/ec2/model/DescribeInstanceStatusResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) WithNextToken(a string) *ServicesEc2ModelDescribeInstanceStatusResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeInstanceStatusResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeInstanceStatusResult clone()
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) Clone() *ServicesEc2ModelDescribeInstanceStatusResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeInstanceStatusResult")
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
	unique_x := &ServicesEc2ModelDescribeInstanceStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeInstanceStatusResult) Clone2() (*JavaLangObject, error) {
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


