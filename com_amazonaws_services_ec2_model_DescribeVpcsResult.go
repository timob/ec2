package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.Vpc> getVpcs()
	GetVpcs() []*ServicesEc2ModelVpc

	// public void setVpcs(java.util.Collection<com.amazonaws.services.ec2.model.Vpc>)
	SetVpcs(a []*ServicesEc2ModelVpc) 

	// public com.amazonaws.services.ec2.model.DescribeVpcsResult withVpcs(com.amazonaws.services.ec2.model.Vpc...)
	WithVpcs(a ...*ServicesEc2ModelVpc) *ServicesEc2ModelDescribeVpcsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcsResult withVpcs(java.util.Collection<com.amazonaws.services.ec2.model.Vpc>)
	WithVpcs2(a []*ServicesEc2ModelVpc) *ServicesEc2ModelDescribeVpcsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcsResult clone()
	Clone() *ServicesEc2ModelDescribeVpcsResult
}

type ServicesEc2ModelDescribeVpcsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcsResult()
func NewServicesEc2ModelDescribeVpcsResult() (*ServicesEc2ModelDescribeVpcsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.Vpc> getVpcs()
func (jbobject *ServicesEc2ModelDescribeVpcsResult) GetVpcs() []*ServicesEc2ModelVpc {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpc)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpcs(java.util.Collection<com.amazonaws.services.ec2.model.Vpc>)
func (jbobject *ServicesEc2ModelDescribeVpcsResult) SetVpcs(a []*ServicesEc2ModelVpc)  {
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

// public com.amazonaws.services.ec2.model.DescribeVpcsResult withVpcs(com.amazonaws.services.ec2.model.Vpc...)
func (jbobject *ServicesEc2ModelDescribeVpcsResult) WithVpcs(a ...*ServicesEc2ModelVpc) *ServicesEc2ModelDescribeVpcsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Vpc")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcs", "com/amazonaws/services/ec2/model/DescribeVpcsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Vpc"))
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
	unique_x := &ServicesEc2ModelDescribeVpcsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcsResult withVpcs(java.util.Collection<com.amazonaws.services.ec2.model.Vpc>)
func (jbobject *ServicesEc2ModelDescribeVpcsResult) WithVpcs2(a []*ServicesEc2ModelVpc) *ServicesEc2ModelDescribeVpcsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcs", "com/amazonaws/services/ec2/model/DescribeVpcsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpcsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcsResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcsResult) Clone() *ServicesEc2ModelDescribeVpcsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcsResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcsResult) Clone2() (*JavaLangObject, error) {
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


