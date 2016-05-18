package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSubnetsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.Subnet> getSubnets()
	GetSubnets() []*ServicesEc2ModelSubnet

	// public void setSubnets(java.util.Collection<com.amazonaws.services.ec2.model.Subnet>)
	SetSubnets(a []*ServicesEc2ModelSubnet) 

	// public com.amazonaws.services.ec2.model.DescribeSubnetsResult withSubnets(com.amazonaws.services.ec2.model.Subnet...)
	WithSubnets(a ...*ServicesEc2ModelSubnet) *ServicesEc2ModelDescribeSubnetsResult

	// public com.amazonaws.services.ec2.model.DescribeSubnetsResult withSubnets(java.util.Collection<com.amazonaws.services.ec2.model.Subnet>)
	WithSubnets2(a []*ServicesEc2ModelSubnet) *ServicesEc2ModelDescribeSubnetsResult

	// public com.amazonaws.services.ec2.model.DescribeSubnetsResult clone()
	Clone() *ServicesEc2ModelDescribeSubnetsResult
}

type ServicesEc2ModelDescribeSubnetsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSubnetsResult()
func NewServicesEc2ModelDescribeSubnetsResult() (*ServicesEc2ModelDescribeSubnetsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSubnetsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSubnetsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.Subnet> getSubnets()
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) GetSubnets() []*ServicesEc2ModelSubnet {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnets", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSubnet)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSubnets(java.util.Collection<com.amazonaws.services.ec2.model.Subnet>)
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) SetSubnets(a []*ServicesEc2ModelSubnet)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnets", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSubnetsResult withSubnets(com.amazonaws.services.ec2.model.Subnet...)
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) WithSubnets(a ...*ServicesEc2ModelSubnet) *ServicesEc2ModelDescribeSubnetsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Subnet")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnets", "com/amazonaws/services/ec2/model/DescribeSubnetsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Subnet"))
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
	unique_x := &ServicesEc2ModelDescribeSubnetsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSubnetsResult withSubnets(java.util.Collection<com.amazonaws.services.ec2.model.Subnet>)
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) WithSubnets2(a []*ServicesEc2ModelSubnet) *ServicesEc2ModelDescribeSubnetsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnets", "com/amazonaws/services/ec2/model/DescribeSubnetsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSubnetsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSubnetsResult clone()
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) Clone() *ServicesEc2ModelDescribeSubnetsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSubnetsResult")
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
	unique_x := &ServicesEc2ModelDescribeSubnetsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSubnetsResult) Clone2() (*JavaLangObject, error) {
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


