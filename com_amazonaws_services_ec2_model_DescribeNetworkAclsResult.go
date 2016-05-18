package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeNetworkAclsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.NetworkAcl> getNetworkAcls()
	GetNetworkAcls() []*ServicesEc2ModelNetworkAcl

	// public void setNetworkAcls(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAcl>)
	SetNetworkAcls(a []*ServicesEc2ModelNetworkAcl) 

	// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult withNetworkAcls(com.amazonaws.services.ec2.model.NetworkAcl...)
	WithNetworkAcls(a ...*ServicesEc2ModelNetworkAcl) *ServicesEc2ModelDescribeNetworkAclsResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult withNetworkAcls(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAcl>)
	WithNetworkAcls2(a []*ServicesEc2ModelNetworkAcl) *ServicesEc2ModelDescribeNetworkAclsResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult clone()
	Clone() *ServicesEc2ModelDescribeNetworkAclsResult
}

type ServicesEc2ModelDescribeNetworkAclsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult()
func NewServicesEc2ModelDescribeNetworkAclsResult() (*ServicesEc2ModelDescribeNetworkAclsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeNetworkAclsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeNetworkAclsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.NetworkAcl> getNetworkAcls()
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) GetNetworkAcls() []*ServicesEc2ModelNetworkAcl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkAcls", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNetworkAcl)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNetworkAcls(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAcl>)
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) SetNetworkAcls(a []*ServicesEc2ModelNetworkAcl)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkAcls", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult withNetworkAcls(com.amazonaws.services.ec2.model.NetworkAcl...)
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) WithNetworkAcls(a ...*ServicesEc2ModelNetworkAcl) *ServicesEc2ModelDescribeNetworkAclsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NetworkAcl")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAcls", "com/amazonaws/services/ec2/model/DescribeNetworkAclsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkAcl"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkAclsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult withNetworkAcls(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAcl>)
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) WithNetworkAcls2(a []*ServicesEc2ModelNetworkAcl) *ServicesEc2ModelDescribeNetworkAclsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAcls", "com/amazonaws/services/ec2/model/DescribeNetworkAclsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkAclsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeNetworkAclsResult clone()
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) Clone() *ServicesEc2ModelDescribeNetworkAclsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeNetworkAclsResult")
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
	unique_x := &ServicesEc2ModelDescribeNetworkAclsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeNetworkAclsResult) Clone2() (*JavaLangObject, error) {
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


