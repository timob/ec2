package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeReservedInstancesModificationsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesModification> getReservedInstancesModifications()
	GetReservedInstancesModifications() []*ServicesEc2ModelReservedInstancesModification

	// public void setReservedInstancesModifications(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModification>)
	SetReservedInstancesModifications(a []*ServicesEc2ModelReservedInstancesModification) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult withReservedInstancesModifications(com.amazonaws.services.ec2.model.ReservedInstancesModification...)
	WithReservedInstancesModifications(a ...*ServicesEc2ModelReservedInstancesModification) *ServicesEc2ModelDescribeReservedInstancesModificationsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult withReservedInstancesModifications(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModification>)
	WithReservedInstancesModifications2(a []*ServicesEc2ModelReservedInstancesModification) *ServicesEc2ModelDescribeReservedInstancesModificationsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesModificationsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult clone()
	Clone() *ServicesEc2ModelDescribeReservedInstancesModificationsResult
}

type ServicesEc2ModelDescribeReservedInstancesModificationsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult()
func NewServicesEc2ModelDescribeReservedInstancesModificationsResult() (*ServicesEc2ModelDescribeReservedInstancesModificationsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesModification> getReservedInstancesModifications()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) GetReservedInstancesModifications() []*ServicesEc2ModelReservedInstancesModification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesModifications", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstancesModification)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setReservedInstancesModifications(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModification>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) SetReservedInstancesModifications(a []*ServicesEc2ModelReservedInstancesModification)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesModifications", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult withReservedInstancesModifications(com.amazonaws.services.ec2.model.ReservedInstancesModification...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) WithReservedInstancesModifications(a ...*ServicesEc2ModelReservedInstancesModification) *ServicesEc2ModelDescribeReservedInstancesModificationsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesModification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesModifications", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesModification"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult withReservedInstancesModifications(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesModification>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) WithReservedInstancesModifications2(a []*ServicesEc2ModelReservedInstancesModification) *ServicesEc2ModelDescribeReservedInstancesModificationsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesModifications", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesModificationsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesModificationsResult clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) Clone() *ServicesEc2ModelDescribeReservedInstancesModificationsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesModificationsResult")
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesModificationsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeReservedInstancesModificationsResult) Clone2() (*JavaLangObject, error) {
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


