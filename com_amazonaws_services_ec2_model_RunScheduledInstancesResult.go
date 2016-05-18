package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRunScheduledInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<java.lang.String> getInstanceIdSet()
	GetInstanceIdSet() []string

	// public void setInstanceIdSet(java.util.Collection<java.lang.String>)
	SetInstanceIdSet(a []string) 

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult withInstanceIdSet(java.lang.String...)
	WithInstanceIdSet(a ...string) *ServicesEc2ModelRunScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult withInstanceIdSet(java.util.Collection<java.lang.String>)
	WithInstanceIdSet2(a []string) *ServicesEc2ModelRunScheduledInstancesResult

	// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult clone()
	Clone() *ServicesEc2ModelRunScheduledInstancesResult
}

type ServicesEc2ModelRunScheduledInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult()
func NewServicesEc2ModelRunScheduledInstancesResult() (*ServicesEc2ModelRunScheduledInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RunScheduledInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRunScheduledInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getInstanceIdSet()
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) GetInstanceIdSet() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceIdSet", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstanceIdSet(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) SetInstanceIdSet(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceIdSet", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult withInstanceIdSet(java.lang.String...)
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) WithInstanceIdSet(a ...string) *ServicesEc2ModelRunScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceIdSet", "com/amazonaws/services/ec2/model/RunScheduledInstancesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult withInstanceIdSet(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) WithInstanceIdSet2(a []string) *ServicesEc2ModelRunScheduledInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceIdSet", "com/amazonaws/services/ec2/model/RunScheduledInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RunScheduledInstancesResult clone()
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) Clone() *ServicesEc2ModelRunScheduledInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RunScheduledInstancesResult")
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
	unique_x := &ServicesEc2ModelRunScheduledInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRunScheduledInstancesResult) Clone2() (*JavaLangObject, error) {
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


