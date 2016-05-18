package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeReservedInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstances> getReservedInstances()
	GetReservedInstances() []*ServicesEc2ModelReservedInstances

	// public void setReservedInstances(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstances>)
	SetReservedInstances(a []*ServicesEc2ModelReservedInstances) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult withReservedInstances(com.amazonaws.services.ec2.model.ReservedInstances...)
	WithReservedInstances(a ...*ServicesEc2ModelReservedInstances) *ServicesEc2ModelDescribeReservedInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult withReservedInstances(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstances>)
	WithReservedInstances2(a []*ServicesEc2ModelReservedInstances) *ServicesEc2ModelDescribeReservedInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult clone()
	Clone() *ServicesEc2ModelDescribeReservedInstancesResult
}

type ServicesEc2ModelDescribeReservedInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult()
func NewServicesEc2ModelDescribeReservedInstancesResult() (*ServicesEc2ModelDescribeReservedInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeReservedInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeReservedInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstances> getReservedInstances()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) GetReservedInstances() []*ServicesEc2ModelReservedInstances {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstances", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstances)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setReservedInstances(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstances>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) SetReservedInstances(a []*ServicesEc2ModelReservedInstances)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult withReservedInstances(com.amazonaws.services.ec2.model.ReservedInstances...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) WithReservedInstances(a ...*ServicesEc2ModelReservedInstances) *ServicesEc2ModelDescribeReservedInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstances")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstances", "com/amazonaws/services/ec2/model/DescribeReservedInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstances"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult withReservedInstances(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstances>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) WithReservedInstances2(a []*ServicesEc2ModelReservedInstances) *ServicesEc2ModelDescribeReservedInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstances", "com/amazonaws/services/ec2/model/DescribeReservedInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesResult clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) Clone() *ServicesEc2ModelDescribeReservedInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesResult")
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeReservedInstancesResult) Clone2() (*JavaLangObject, error) {
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


