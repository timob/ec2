package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeMovingAddressesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.MovingAddressStatus> getMovingAddressStatuses()
	GetMovingAddressStatuses() []*ServicesEc2ModelMovingAddressStatus

	// public void setMovingAddressStatuses(java.util.Collection<com.amazonaws.services.ec2.model.MovingAddressStatus>)
	SetMovingAddressStatuses(a []*ServicesEc2ModelMovingAddressStatus) 

	// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult withMovingAddressStatuses(com.amazonaws.services.ec2.model.MovingAddressStatus...)
	WithMovingAddressStatuses(a ...*ServicesEc2ModelMovingAddressStatus) *ServicesEc2ModelDescribeMovingAddressesResult

	// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult withMovingAddressStatuses(java.util.Collection<com.amazonaws.services.ec2.model.MovingAddressStatus>)
	WithMovingAddressStatuses2(a []*ServicesEc2ModelMovingAddressStatus) *ServicesEc2ModelDescribeMovingAddressesResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeMovingAddressesResult

	// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult clone()
	Clone() *ServicesEc2ModelDescribeMovingAddressesResult
}

type ServicesEc2ModelDescribeMovingAddressesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult()
func NewServicesEc2ModelDescribeMovingAddressesResult() (*ServicesEc2ModelDescribeMovingAddressesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeMovingAddressesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.MovingAddressStatus> getMovingAddressStatuses()
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) GetMovingAddressStatuses() []*ServicesEc2ModelMovingAddressStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMovingAddressStatuses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelMovingAddressStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setMovingAddressStatuses(java.util.Collection<com.amazonaws.services.ec2.model.MovingAddressStatus>)
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) SetMovingAddressStatuses(a []*ServicesEc2ModelMovingAddressStatus)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMovingAddressStatuses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult withMovingAddressStatuses(com.amazonaws.services.ec2.model.MovingAddressStatus...)
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) WithMovingAddressStatuses(a ...*ServicesEc2ModelMovingAddressStatus) *ServicesEc2ModelDescribeMovingAddressesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/MovingAddressStatus")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMovingAddressStatuses", "com/amazonaws/services/ec2/model/DescribeMovingAddressesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/MovingAddressStatus"))
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
	unique_x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult withMovingAddressStatuses(java.util.Collection<com.amazonaws.services.ec2.model.MovingAddressStatus>)
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) WithMovingAddressStatuses2(a []*ServicesEc2ModelMovingAddressStatus) *ServicesEc2ModelDescribeMovingAddressesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMovingAddressStatuses", "com/amazonaws/services/ec2/model/DescribeMovingAddressesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) WithNextToken(a string) *ServicesEc2ModelDescribeMovingAddressesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeMovingAddressesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeMovingAddressesResult clone()
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) Clone() *ServicesEc2ModelDescribeMovingAddressesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeMovingAddressesResult")
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
	unique_x := &ServicesEc2ModelDescribeMovingAddressesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeMovingAddressesResult) Clone2() (*JavaLangObject, error) {
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


