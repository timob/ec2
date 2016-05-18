package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeAccountAttributesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.AccountAttribute> getAccountAttributes()
	GetAccountAttributes() []*ServicesEc2ModelAccountAttribute

	// public void setAccountAttributes(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttribute>)
	SetAccountAttributes(a []*ServicesEc2ModelAccountAttribute) 

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult withAccountAttributes(com.amazonaws.services.ec2.model.AccountAttribute...)
	WithAccountAttributes(a ...*ServicesEc2ModelAccountAttribute) *ServicesEc2ModelDescribeAccountAttributesResult

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult withAccountAttributes(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttribute>)
	WithAccountAttributes2(a []*ServicesEc2ModelAccountAttribute) *ServicesEc2ModelDescribeAccountAttributesResult

	// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult clone()
	Clone() *ServicesEc2ModelDescribeAccountAttributesResult
}

type ServicesEc2ModelDescribeAccountAttributesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult()
func NewServicesEc2ModelDescribeAccountAttributesResult() (*ServicesEc2ModelDescribeAccountAttributesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeAccountAttributesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeAccountAttributesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.AccountAttribute> getAccountAttributes()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) GetAccountAttributes() []*ServicesEc2ModelAccountAttribute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAccountAttributes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelAccountAttribute)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAccountAttributes(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttribute>)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) SetAccountAttributes(a []*ServicesEc2ModelAccountAttribute)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAccountAttributes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult withAccountAttributes(com.amazonaws.services.ec2.model.AccountAttribute...)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) WithAccountAttributes(a ...*ServicesEc2ModelAccountAttribute) *ServicesEc2ModelDescribeAccountAttributesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/AccountAttribute")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAccountAttributes", "com/amazonaws/services/ec2/model/DescribeAccountAttributesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AccountAttribute"))
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult withAccountAttributes(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttribute>)
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) WithAccountAttributes2(a []*ServicesEc2ModelAccountAttribute) *ServicesEc2ModelDescribeAccountAttributesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAccountAttributes", "com/amazonaws/services/ec2/model/DescribeAccountAttributesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeAccountAttributesResult clone()
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) Clone() *ServicesEc2ModelDescribeAccountAttributesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeAccountAttributesResult")
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
	unique_x := &ServicesEc2ModelDescribeAccountAttributesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeAccountAttributesResult) Clone2() (*JavaLangObject, error) {
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


