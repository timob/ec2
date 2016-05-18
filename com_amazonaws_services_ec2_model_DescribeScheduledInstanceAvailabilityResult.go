package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeScheduledInstanceAvailabilityResultInterface interface {
	JavaLangObjectInterface

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult

	// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstanceAvailability> getScheduledInstanceAvailabilitySet()
	GetScheduledInstanceAvailabilitySet() []*ServicesEc2ModelScheduledInstanceAvailability

	// public void setScheduledInstanceAvailabilitySet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstanceAvailability>)
	SetScheduledInstanceAvailabilitySet(a []*ServicesEc2ModelScheduledInstanceAvailability) 

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult withScheduledInstanceAvailabilitySet(com.amazonaws.services.ec2.model.ScheduledInstanceAvailability...)
	WithScheduledInstanceAvailabilitySet(a ...*ServicesEc2ModelScheduledInstanceAvailability) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult withScheduledInstanceAvailabilitySet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstanceAvailability>)
	WithScheduledInstanceAvailabilitySet2(a []*ServicesEc2ModelScheduledInstanceAvailability) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult

	// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult clone()
	Clone() *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult
}

type ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult()
func NewServicesEc2ModelDescribeScheduledInstanceAvailabilityResult() (*ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) WithNextToken(a string) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ScheduledInstanceAvailability> getScheduledInstanceAvailabilitySet()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) GetScheduledInstanceAvailabilitySet() []*ServicesEc2ModelScheduledInstanceAvailability {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScheduledInstanceAvailabilitySet", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelScheduledInstanceAvailability)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setScheduledInstanceAvailabilitySet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstanceAvailability>)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) SetScheduledInstanceAvailabilitySet(a []*ServicesEc2ModelScheduledInstanceAvailability)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScheduledInstanceAvailabilitySet", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult withScheduledInstanceAvailabilitySet(com.amazonaws.services.ec2.model.ScheduledInstanceAvailability...)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) WithScheduledInstanceAvailabilitySet(a ...*ServicesEc2ModelScheduledInstanceAvailability) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceAvailabilitySet", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstanceAvailability"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult withScheduledInstanceAvailabilitySet(java.util.Collection<com.amazonaws.services.ec2.model.ScheduledInstanceAvailability>)
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) WithScheduledInstanceAvailabilitySet2(a []*ServicesEc2ModelScheduledInstanceAvailability) *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceAvailabilitySet", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeScheduledInstanceAvailabilityResult clone()
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) Clone() *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeScheduledInstanceAvailabilityResult")
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
	unique_x := &ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeScheduledInstanceAvailabilityResult) Clone2() (*JavaLangObject, error) {
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


