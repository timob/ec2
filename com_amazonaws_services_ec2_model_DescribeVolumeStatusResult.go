package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVolumeStatusResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusItem> getVolumeStatuses()
	GetVolumeStatuses() []*ServicesEc2ModelVolumeStatusItem

	// public void setVolumeStatuses(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusItem>)
	SetVolumeStatuses(a []*ServicesEc2ModelVolumeStatusItem) 

	// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult withVolumeStatuses(com.amazonaws.services.ec2.model.VolumeStatusItem...)
	WithVolumeStatuses(a ...*ServicesEc2ModelVolumeStatusItem) *ServicesEc2ModelDescribeVolumeStatusResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult withVolumeStatuses(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusItem>)
	WithVolumeStatuses2(a []*ServicesEc2ModelVolumeStatusItem) *ServicesEc2ModelDescribeVolumeStatusResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeVolumeStatusResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult clone()
	Clone() *ServicesEc2ModelDescribeVolumeStatusResult
}

type ServicesEc2ModelDescribeVolumeStatusResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult()
func NewServicesEc2ModelDescribeVolumeStatusResult() (*ServicesEc2ModelDescribeVolumeStatusResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVolumeStatusResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusItem> getVolumeStatuses()
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) GetVolumeStatuses() []*ServicesEc2ModelVolumeStatusItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeStatuses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVolumeStatusItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVolumeStatuses(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusItem>)
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) SetVolumeStatuses(a []*ServicesEc2ModelVolumeStatusItem)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeStatuses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult withVolumeStatuses(com.amazonaws.services.ec2.model.VolumeStatusItem...)
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) WithVolumeStatuses(a ...*ServicesEc2ModelVolumeStatusItem) *ServicesEc2ModelDescribeVolumeStatusResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VolumeStatusItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeStatuses", "com/amazonaws/services/ec2/model/DescribeVolumeStatusResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusItem"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult withVolumeStatuses(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusItem>)
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) WithVolumeStatuses2(a []*ServicesEc2ModelVolumeStatusItem) *ServicesEc2ModelDescribeVolumeStatusResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeStatuses", "com/amazonaws/services/ec2/model/DescribeVolumeStatusResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) WithNextToken(a string) *ServicesEc2ModelDescribeVolumeStatusResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeVolumeStatusResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVolumeStatusResult clone()
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) Clone() *ServicesEc2ModelDescribeVolumeStatusResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVolumeStatusResult")
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
	unique_x := &ServicesEc2ModelDescribeVolumeStatusResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVolumeStatusResult) Clone2() (*JavaLangObject, error) {
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


