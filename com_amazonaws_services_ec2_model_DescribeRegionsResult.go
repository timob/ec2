package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeRegionsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.Region> getRegions()
	GetRegions() []*ServicesEc2ModelRegion

	// public void setRegions(java.util.Collection<com.amazonaws.services.ec2.model.Region>)
	SetRegions(a []*ServicesEc2ModelRegion) 

	// public com.amazonaws.services.ec2.model.DescribeRegionsResult withRegions(com.amazonaws.services.ec2.model.Region...)
	WithRegions(a ...*ServicesEc2ModelRegion) *ServicesEc2ModelDescribeRegionsResult

	// public com.amazonaws.services.ec2.model.DescribeRegionsResult withRegions(java.util.Collection<com.amazonaws.services.ec2.model.Region>)
	WithRegions2(a []*ServicesEc2ModelRegion) *ServicesEc2ModelDescribeRegionsResult

	// public com.amazonaws.services.ec2.model.DescribeRegionsResult clone()
	Clone() *ServicesEc2ModelDescribeRegionsResult
}

type ServicesEc2ModelDescribeRegionsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeRegionsResult()
func NewServicesEc2ModelDescribeRegionsResult() (*ServicesEc2ModelDescribeRegionsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeRegionsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeRegionsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.Region> getRegions()
func (jbobject *ServicesEc2ModelDescribeRegionsResult) GetRegions() []*ServicesEc2ModelRegion {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRegions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelRegion)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRegions(java.util.Collection<com.amazonaws.services.ec2.model.Region>)
func (jbobject *ServicesEc2ModelDescribeRegionsResult) SetRegions(a []*ServicesEc2ModelRegion)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRegions", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeRegionsResult withRegions(com.amazonaws.services.ec2.model.Region...)
func (jbobject *ServicesEc2ModelDescribeRegionsResult) WithRegions(a ...*ServicesEc2ModelRegion) *ServicesEc2ModelDescribeRegionsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Region")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRegions", "com/amazonaws/services/ec2/model/DescribeRegionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Region"))
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
	unique_x := &ServicesEc2ModelDescribeRegionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeRegionsResult withRegions(java.util.Collection<com.amazonaws.services.ec2.model.Region>)
func (jbobject *ServicesEc2ModelDescribeRegionsResult) WithRegions2(a []*ServicesEc2ModelRegion) *ServicesEc2ModelDescribeRegionsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRegions", "com/amazonaws/services/ec2/model/DescribeRegionsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeRegionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeRegionsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeRegionsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeRegionsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeRegionsResult clone()
func (jbobject *ServicesEc2ModelDescribeRegionsResult) Clone() *ServicesEc2ModelDescribeRegionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeRegionsResult")
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
	unique_x := &ServicesEc2ModelDescribeRegionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeRegionsResult) Clone2() (*JavaLangObject, error) {
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


