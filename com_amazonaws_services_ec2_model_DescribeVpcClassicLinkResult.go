package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcClassicLinkResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.VpcClassicLink> getVpcs()
	GetVpcs() []*ServicesEc2ModelVpcClassicLink

	// public void setVpcs(java.util.Collection<com.amazonaws.services.ec2.model.VpcClassicLink>)
	SetVpcs(a []*ServicesEc2ModelVpcClassicLink) 

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult withVpcs(com.amazonaws.services.ec2.model.VpcClassicLink...)
	WithVpcs(a ...*ServicesEc2ModelVpcClassicLink) *ServicesEc2ModelDescribeVpcClassicLinkResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult withVpcs(java.util.Collection<com.amazonaws.services.ec2.model.VpcClassicLink>)
	WithVpcs2(a []*ServicesEc2ModelVpcClassicLink) *ServicesEc2ModelDescribeVpcClassicLinkResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult clone()
	Clone() *ServicesEc2ModelDescribeVpcClassicLinkResult
}

type ServicesEc2ModelDescribeVpcClassicLinkResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult()
func NewServicesEc2ModelDescribeVpcClassicLinkResult() (*ServicesEc2ModelDescribeVpcClassicLinkResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcClassicLinkResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcClassicLinkResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.VpcClassicLink> getVpcs()
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) GetVpcs() []*ServicesEc2ModelVpcClassicLink {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpcClassicLink)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpcs(java.util.Collection<com.amazonaws.services.ec2.model.VpcClassicLink>)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) SetVpcs(a []*ServicesEc2ModelVpcClassicLink)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult withVpcs(com.amazonaws.services.ec2.model.VpcClassicLink...)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) WithVpcs(a ...*ServicesEc2ModelVpcClassicLink) *ServicesEc2ModelDescribeVpcClassicLinkResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VpcClassicLink")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcs", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcClassicLink"))
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult withVpcs(java.util.Collection<com.amazonaws.services.ec2.model.VpcClassicLink>)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) WithVpcs2(a []*ServicesEc2ModelVpcClassicLink) *ServicesEc2ModelDescribeVpcClassicLinkResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcs", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) Clone() *ServicesEc2ModelDescribeVpcClassicLinkResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkResult) Clone2() (*JavaLangObject, error) {
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


