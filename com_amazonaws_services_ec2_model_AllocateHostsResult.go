package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAllocateHostsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<java.lang.String> getHostIds()
	GetHostIds() []string

	// public void setHostIds(java.util.Collection<java.lang.String>)
	SetHostIds(a []string) 

	// public com.amazonaws.services.ec2.model.AllocateHostsResult withHostIds(java.lang.String...)
	WithHostIds(a ...string) *ServicesEc2ModelAllocateHostsResult

	// public com.amazonaws.services.ec2.model.AllocateHostsResult withHostIds(java.util.Collection<java.lang.String>)
	WithHostIds2(a []string) *ServicesEc2ModelAllocateHostsResult

	// public com.amazonaws.services.ec2.model.AllocateHostsResult clone()
	Clone() *ServicesEc2ModelAllocateHostsResult
}

type ServicesEc2ModelAllocateHostsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AllocateHostsResult()
func NewServicesEc2ModelAllocateHostsResult() (*ServicesEc2ModelAllocateHostsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AllocateHostsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAllocateHostsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getHostIds()
func (jbobject *ServicesEc2ModelAllocateHostsResult) GetHostIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostIds", "java/util/List")
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

// public void setHostIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelAllocateHostsResult) SetHostIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AllocateHostsResult withHostIds(java.lang.String...)
func (jbobject *ServicesEc2ModelAllocateHostsResult) WithHostIds(a ...string) *ServicesEc2ModelAllocateHostsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostIds", "com/amazonaws/services/ec2/model/AllocateHostsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocateHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AllocateHostsResult withHostIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelAllocateHostsResult) WithHostIds2(a []string) *ServicesEc2ModelAllocateHostsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostIds", "com/amazonaws/services/ec2/model/AllocateHostsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelAllocateHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAllocateHostsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAllocateHostsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAllocateHostsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AllocateHostsResult clone()
func (jbobject *ServicesEc2ModelAllocateHostsResult) Clone() *ServicesEc2ModelAllocateHostsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AllocateHostsResult")
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
	unique_x := &ServicesEc2ModelAllocateHostsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAllocateHostsResult) Clone2() (*JavaLangObject, error) {
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


