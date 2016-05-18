package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeKeyPairsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.KeyPairInfo> getKeyPairs()
	GetKeyPairs() []*ServicesEc2ModelKeyPairInfo

	// public void setKeyPairs(java.util.Collection<com.amazonaws.services.ec2.model.KeyPairInfo>)
	SetKeyPairs(a []*ServicesEc2ModelKeyPairInfo) 

	// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult withKeyPairs(com.amazonaws.services.ec2.model.KeyPairInfo...)
	WithKeyPairs(a ...*ServicesEc2ModelKeyPairInfo) *ServicesEc2ModelDescribeKeyPairsResult

	// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult withKeyPairs(java.util.Collection<com.amazonaws.services.ec2.model.KeyPairInfo>)
	WithKeyPairs2(a []*ServicesEc2ModelKeyPairInfo) *ServicesEc2ModelDescribeKeyPairsResult

	// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult clone()
	Clone() *ServicesEc2ModelDescribeKeyPairsResult
}

type ServicesEc2ModelDescribeKeyPairsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult()
func NewServicesEc2ModelDescribeKeyPairsResult() (*ServicesEc2ModelDescribeKeyPairsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeKeyPairsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeKeyPairsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.KeyPairInfo> getKeyPairs()
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) GetKeyPairs() []*ServicesEc2ModelKeyPairInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKeyPairs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelKeyPairInfo)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setKeyPairs(java.util.Collection<com.amazonaws.services.ec2.model.KeyPairInfo>)
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) SetKeyPairs(a []*ServicesEc2ModelKeyPairInfo)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKeyPairs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult withKeyPairs(com.amazonaws.services.ec2.model.KeyPairInfo...)
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) WithKeyPairs(a ...*ServicesEc2ModelKeyPairInfo) *ServicesEc2ModelDescribeKeyPairsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/KeyPairInfo")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyPairs", "com/amazonaws/services/ec2/model/DescribeKeyPairsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/KeyPairInfo"))
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
	unique_x := &ServicesEc2ModelDescribeKeyPairsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult withKeyPairs(java.util.Collection<com.amazonaws.services.ec2.model.KeyPairInfo>)
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) WithKeyPairs2(a []*ServicesEc2ModelKeyPairInfo) *ServicesEc2ModelDescribeKeyPairsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKeyPairs", "com/amazonaws/services/ec2/model/DescribeKeyPairsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeKeyPairsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeKeyPairsResult clone()
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) Clone() *ServicesEc2ModelDescribeKeyPairsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeKeyPairsResult")
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
	unique_x := &ServicesEc2ModelDescribeKeyPairsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeKeyPairsResult) Clone2() (*JavaLangObject, error) {
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


