package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ClassicLinkDnsSupport> getVpcs()
	GetVpcs() []*ServicesEc2ModelClassicLinkDnsSupport

	// public void setVpcs(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkDnsSupport>)
	SetVpcs(a []*ServicesEc2ModelClassicLinkDnsSupport) 

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult withVpcs(com.amazonaws.services.ec2.model.ClassicLinkDnsSupport...)
	WithVpcs(a ...*ServicesEc2ModelClassicLinkDnsSupport) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult withVpcs(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkDnsSupport>)
	WithVpcs2(a []*ServicesEc2ModelClassicLinkDnsSupport) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult

	// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult clone()
	Clone() *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult
}

type ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult()
func NewServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult() (*ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ClassicLinkDnsSupport> getVpcs()
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) GetVpcs() []*ServicesEc2ModelClassicLinkDnsSupport {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelClassicLinkDnsSupport)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpcs(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkDnsSupport>)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) SetVpcs(a []*ServicesEc2ModelClassicLinkDnsSupport)  {
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

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult withVpcs(com.amazonaws.services.ec2.model.ClassicLinkDnsSupport...)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) WithVpcs(a ...*ServicesEc2ModelClassicLinkDnsSupport) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ClassicLinkDnsSupport")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcs", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ClassicLinkDnsSupport"))
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult withVpcs(java.util.Collection<com.amazonaws.services.ec2.model.ClassicLinkDnsSupport>)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) WithVpcs2(a []*ServicesEc2ModelClassicLinkDnsSupport) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcs", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) WithNextToken(a string) *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcClassicLinkDnsSupportResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) Clone() *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcClassicLinkDnsSupportResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcClassicLinkDnsSupportResult) Clone2() (*JavaLangObject, error) {
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


