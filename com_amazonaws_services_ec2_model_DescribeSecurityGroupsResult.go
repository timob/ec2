package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSecurityGroupsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.SecurityGroup> getSecurityGroups()
	GetSecurityGroups() []*ServicesEc2ModelSecurityGroup

	// public void setSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.SecurityGroup>)
	SetSecurityGroups(a []*ServicesEc2ModelSecurityGroup) 

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult withSecurityGroups(com.amazonaws.services.ec2.model.SecurityGroup...)
	WithSecurityGroups(a ...*ServicesEc2ModelSecurityGroup) *ServicesEc2ModelDescribeSecurityGroupsResult

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult withSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.SecurityGroup>)
	WithSecurityGroups2(a []*ServicesEc2ModelSecurityGroup) *ServicesEc2ModelDescribeSecurityGroupsResult

	// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult clone()
	Clone() *ServicesEc2ModelDescribeSecurityGroupsResult
}

type ServicesEc2ModelDescribeSecurityGroupsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult()
func NewServicesEc2ModelDescribeSecurityGroupsResult() (*ServicesEc2ModelDescribeSecurityGroupsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSecurityGroupsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSecurityGroupsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.SecurityGroup> getSecurityGroups()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) GetSecurityGroups() []*ServicesEc2ModelSecurityGroup {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecurityGroups", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSecurityGroup)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.SecurityGroup>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) SetSecurityGroups(a []*ServicesEc2ModelSecurityGroup)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSecurityGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult withSecurityGroups(com.amazonaws.services.ec2.model.SecurityGroup...)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) WithSecurityGroups(a ...*ServicesEc2ModelSecurityGroup) *ServicesEc2ModelDescribeSecurityGroupsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/SecurityGroup")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SecurityGroup"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult withSecurityGroups(java.util.Collection<com.amazonaws.services.ec2.model.SecurityGroup>)
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) WithSecurityGroups2(a []*ServicesEc2ModelSecurityGroup) *ServicesEc2ModelDescribeSecurityGroupsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecurityGroups", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSecurityGroupsResult clone()
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) Clone() *ServicesEc2ModelDescribeSecurityGroupsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSecurityGroupsResult")
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
	unique_x := &ServicesEc2ModelDescribeSecurityGroupsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSecurityGroupsResult) Clone2() (*JavaLangObject, error) {
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


