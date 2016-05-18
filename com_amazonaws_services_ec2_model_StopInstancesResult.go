package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStopInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceStateChange> getStoppingInstances()
	GetStoppingInstances() []*ServicesEc2ModelInstanceStateChange

	// public void setStoppingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
	SetStoppingInstances(a []*ServicesEc2ModelInstanceStateChange) 

	// public com.amazonaws.services.ec2.model.StopInstancesResult withStoppingInstances(com.amazonaws.services.ec2.model.InstanceStateChange...)
	WithStoppingInstances(a ...*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStopInstancesResult

	// public com.amazonaws.services.ec2.model.StopInstancesResult withStoppingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
	WithStoppingInstances2(a []*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStopInstancesResult

	// public com.amazonaws.services.ec2.model.StopInstancesResult clone()
	Clone() *ServicesEc2ModelStopInstancesResult
}

type ServicesEc2ModelStopInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.StopInstancesResult()
func NewServicesEc2ModelStopInstancesResult() (*ServicesEc2ModelStopInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/StopInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelStopInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceStateChange> getStoppingInstances()
func (jbobject *ServicesEc2ModelStopInstancesResult) GetStoppingInstances() []*ServicesEc2ModelInstanceStateChange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStoppingInstances", "java/util/List")
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

// public void setStoppingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
func (jbobject *ServicesEc2ModelStopInstancesResult) SetStoppingInstances(a []*ServicesEc2ModelInstanceStateChange)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStoppingInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.StopInstancesResult withStoppingInstances(com.amazonaws.services.ec2.model.InstanceStateChange...)
func (jbobject *ServicesEc2ModelStopInstancesResult) WithStoppingInstances(a ...*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStopInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceStateChange")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStoppingInstances", "com/amazonaws/services/ec2/model/StopInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStateChange"))
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
	unique_x := &ServicesEc2ModelStopInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.StopInstancesResult withStoppingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
func (jbobject *ServicesEc2ModelStopInstancesResult) WithStoppingInstances2(a []*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStopInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStoppingInstances", "com/amazonaws/services/ec2/model/StopInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelStopInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelStopInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelStopInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelStopInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.StopInstancesResult clone()
func (jbobject *ServicesEc2ModelStopInstancesResult) Clone() *ServicesEc2ModelStopInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/StopInstancesResult")
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
	unique_x := &ServicesEc2ModelStopInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelStopInstancesResult) Clone2() (*JavaLangObject, error) {
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


