package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeIdFormatResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.IdFormat> getStatuses()
	GetStatuses() []*ServicesEc2ModelIdFormat

	// public void setStatuses(java.util.Collection<com.amazonaws.services.ec2.model.IdFormat>)
	SetStatuses(a []*ServicesEc2ModelIdFormat) 

	// public com.amazonaws.services.ec2.model.DescribeIdFormatResult withStatuses(com.amazonaws.services.ec2.model.IdFormat...)
	WithStatuses(a ...*ServicesEc2ModelIdFormat) *ServicesEc2ModelDescribeIdFormatResult

	// public com.amazonaws.services.ec2.model.DescribeIdFormatResult withStatuses(java.util.Collection<com.amazonaws.services.ec2.model.IdFormat>)
	WithStatuses2(a []*ServicesEc2ModelIdFormat) *ServicesEc2ModelDescribeIdFormatResult

	// public com.amazonaws.services.ec2.model.DescribeIdFormatResult clone()
	Clone() *ServicesEc2ModelDescribeIdFormatResult
}

type ServicesEc2ModelDescribeIdFormatResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeIdFormatResult()
func NewServicesEc2ModelDescribeIdFormatResult() (*ServicesEc2ModelDescribeIdFormatResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeIdFormatResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeIdFormatResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.IdFormat> getStatuses()
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) GetStatuses() []*ServicesEc2ModelIdFormat {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatuses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelIdFormat)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setStatuses(java.util.Collection<com.amazonaws.services.ec2.model.IdFormat>)
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) SetStatuses(a []*ServicesEc2ModelIdFormat)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatuses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeIdFormatResult withStatuses(com.amazonaws.services.ec2.model.IdFormat...)
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) WithStatuses(a ...*ServicesEc2ModelIdFormat) *ServicesEc2ModelDescribeIdFormatResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/IdFormat")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatuses", "com/amazonaws/services/ec2/model/DescribeIdFormatResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IdFormat"))
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
	unique_x := &ServicesEc2ModelDescribeIdFormatResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeIdFormatResult withStatuses(java.util.Collection<com.amazonaws.services.ec2.model.IdFormat>)
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) WithStatuses2(a []*ServicesEc2ModelIdFormat) *ServicesEc2ModelDescribeIdFormatResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatuses", "com/amazonaws/services/ec2/model/DescribeIdFormatResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeIdFormatResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeIdFormatResult clone()
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) Clone() *ServicesEc2ModelDescribeIdFormatResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeIdFormatResult")
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
	unique_x := &ServicesEc2ModelDescribeIdFormatResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeIdFormatResult) Clone2() (*JavaLangObject, error) {
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


