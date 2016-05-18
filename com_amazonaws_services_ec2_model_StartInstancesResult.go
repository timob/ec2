package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStartInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceStateChange> getStartingInstances()
	GetStartingInstances() []*ServicesEc2ModelInstanceStateChange

	// public void setStartingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
	SetStartingInstances(a []*ServicesEc2ModelInstanceStateChange) 

	// public com.amazonaws.services.ec2.model.StartInstancesResult withStartingInstances(com.amazonaws.services.ec2.model.InstanceStateChange...)
	WithStartingInstances(a ...*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStartInstancesResult

	// public com.amazonaws.services.ec2.model.StartInstancesResult withStartingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
	WithStartingInstances2(a []*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStartInstancesResult

	// public com.amazonaws.services.ec2.model.StartInstancesResult clone()
	Clone() *ServicesEc2ModelStartInstancesResult
}

type ServicesEc2ModelStartInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.StartInstancesResult()
func NewServicesEc2ModelStartInstancesResult() (*ServicesEc2ModelStartInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/StartInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelStartInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceStateChange> getStartingInstances()
func (jbobject *ServicesEc2ModelStartInstancesResult) GetStartingInstances() []*ServicesEc2ModelInstanceStateChange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartingInstances", "java/util/List")
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

// public void setStartingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
func (jbobject *ServicesEc2ModelStartInstancesResult) SetStartingInstances(a []*ServicesEc2ModelInstanceStateChange)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStartingInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.StartInstancesResult withStartingInstances(com.amazonaws.services.ec2.model.InstanceStateChange...)
func (jbobject *ServicesEc2ModelStartInstancesResult) WithStartingInstances(a ...*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStartInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceStateChange")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStartingInstances", "com/amazonaws/services/ec2/model/StartInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStateChange"))
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
	unique_x := &ServicesEc2ModelStartInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.StartInstancesResult withStartingInstances(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStateChange>)
func (jbobject *ServicesEc2ModelStartInstancesResult) WithStartingInstances2(a []*ServicesEc2ModelInstanceStateChange) *ServicesEc2ModelStartInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStartingInstances", "com/amazonaws/services/ec2/model/StartInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelStartInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelStartInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelStartInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelStartInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.StartInstancesResult clone()
func (jbobject *ServicesEc2ModelStartInstancesResult) Clone() *ServicesEc2ModelStartInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/StartInstancesResult")
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
	unique_x := &ServicesEc2ModelStartInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelStartInstancesResult) Clone2() (*JavaLangObject, error) {
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


