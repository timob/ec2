package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeTagsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.TagDescription> getTags()
	GetTags() []*ServicesEc2ModelTagDescription

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.TagDescription>)
	SetTags(a []*ServicesEc2ModelTagDescription) 

	// public com.amazonaws.services.ec2.model.DescribeTagsResult withTags(com.amazonaws.services.ec2.model.TagDescription...)
	WithTags(a ...*ServicesEc2ModelTagDescription) *ServicesEc2ModelDescribeTagsResult

	// public com.amazonaws.services.ec2.model.DescribeTagsResult withTags(java.util.Collection<com.amazonaws.services.ec2.model.TagDescription>)
	WithTags2(a []*ServicesEc2ModelTagDescription) *ServicesEc2ModelDescribeTagsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeTagsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeTagsResult

	// public com.amazonaws.services.ec2.model.DescribeTagsResult clone()
	Clone() *ServicesEc2ModelDescribeTagsResult
}

type ServicesEc2ModelDescribeTagsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeTagsResult()
func NewServicesEc2ModelDescribeTagsResult() (*ServicesEc2ModelDescribeTagsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeTagsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeTagsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.TagDescription> getTags()
func (jbobject *ServicesEc2ModelDescribeTagsResult) GetTags() []*ServicesEc2ModelTagDescription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTagDescription)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.TagDescription>)
func (jbobject *ServicesEc2ModelDescribeTagsResult) SetTags(a []*ServicesEc2ModelTagDescription)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeTagsResult withTags(com.amazonaws.services.ec2.model.TagDescription...)
func (jbobject *ServicesEc2ModelDescribeTagsResult) WithTags(a ...*ServicesEc2ModelTagDescription) *ServicesEc2ModelDescribeTagsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/TagDescription")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/DescribeTagsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/TagDescription"))
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
	unique_x := &ServicesEc2ModelDescribeTagsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeTagsResult withTags(java.util.Collection<com.amazonaws.services.ec2.model.TagDescription>)
func (jbobject *ServicesEc2ModelDescribeTagsResult) WithTags2(a []*ServicesEc2ModelTagDescription) *ServicesEc2ModelDescribeTagsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/DescribeTagsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeTagsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeTagsResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeTagsResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeTagsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeTagsResult) WithNextToken(a string) *ServicesEc2ModelDescribeTagsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeTagsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeTagsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeTagsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeTagsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeTagsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeTagsResult clone()
func (jbobject *ServicesEc2ModelDescribeTagsResult) Clone() *ServicesEc2ModelDescribeTagsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeTagsResult")
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
	unique_x := &ServicesEc2ModelDescribeTagsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeTagsResult) Clone2() (*JavaLangObject, error) {
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


