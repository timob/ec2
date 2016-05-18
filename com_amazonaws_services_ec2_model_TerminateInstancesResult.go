package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelTerminateInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceStateChange> getTerminatingInstances()
	GetTerminatingInstances() []*ServicesEc2ModelInstanceStateChange

	// public void setTerminatingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
	SetTerminatingInstances(a []*ServicesEc2ModelInstanceStateChange) 

	// public com.amazonaws.services.ec2.model.TerminateInstancesResult withTerminatingInstances(com.amazonaws.services.ec2.model.InstanceStateChange...)
	WithTerminatingInstances(a ...*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelTerminateInstancesResult

	// public com.amazonaws.services.ec2.model.TerminateInstancesResult withTerminatingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
	WithTerminatingInstances2(a []*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelTerminateInstancesResult

	// public com.amazonaws.services.ec2.model.TerminateInstancesResult clone()
	Clone() *ServicesEc2ModelTerminateInstancesResult
}

type ServicesEc2ModelTerminateInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.TerminateInstancesResult()
func NewServicesEc2ModelTerminateInstancesResult() (*ServicesEc2ModelTerminateInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/TerminateInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelTerminateInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceStateChange> getTerminatingInstances()
func (jbobject *ServicesEc2ModelTerminateInstancesResult) GetTerminatingInstances() []*ServicesEc2ModelInstanceStateChange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTerminatingInstances", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceStateChange)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTerminatingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
func (jbobject *ServicesEc2ModelTerminateInstancesResult) SetTerminatingInstances(a []*ServicesEc2ModelInstanceStateChange)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTerminatingInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.TerminateInstancesResult withTerminatingInstances(com.amazonaws.services.ec2.model.InstanceStateChange...)
func (jbobject *ServicesEc2ModelTerminateInstancesResult) WithTerminatingInstances(a ...*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelTerminateInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceStateChange")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTerminatingInstances", "com/amazonaws/services/ec2/model/TerminateInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStateChange"))
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
	unique_x := &ServicesEc2ModelTerminateInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.TerminateInstancesResult withTerminatingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
func (jbobject *ServicesEc2ModelTerminateInstancesResult) WithTerminatingInstances2(a []*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelTerminateInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTerminatingInstances", "com/amazonaws/services/ec2/model/TerminateInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelTerminateInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelTerminateInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelTerminateInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelTerminateInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.TerminateInstancesResult clone()
func (jbobject *ServicesEc2ModelTerminateInstancesResult) Clone() *ServicesEc2ModelTerminateInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/TerminateInstancesResult")
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
	unique_x := &ServicesEc2ModelTerminateInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelTerminateInstancesResult) Clone2() (*JavaLangObject, error) {
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


