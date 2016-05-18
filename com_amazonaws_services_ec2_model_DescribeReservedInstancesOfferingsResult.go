package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeReservedInstancesOfferingsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesOffering> getReservedInstancesOfferings()
	GetReservedInstancesOfferings() []*ServicesEc2ModelReservedInstancesOffering

	// public void setReservedInstancesOfferings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesOffering>)
	SetReservedInstancesOfferings(a []*ServicesEc2ModelReservedInstancesOffering) 

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult withReservedInstancesOfferings(com.amazonaws.services.ec2.model.ReservedInstancesOffering...)
	WithReservedInstancesOfferings(a ...*ServicesEc2ModelReservedInstancesOffering) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult withReservedInstancesOfferings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesOffering>)
	WithReservedInstancesOfferings2(a []*ServicesEc2ModelReservedInstancesOffering) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult

	// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult clone()
	Clone() *ServicesEc2ModelDescribeReservedInstancesOfferingsResult
}

type ServicesEc2ModelDescribeReservedInstancesOfferingsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult()
func NewServicesEc2ModelDescribeReservedInstancesOfferingsResult() (*ServicesEc2ModelDescribeReservedInstancesOfferingsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ReservedInstancesOffering> getReservedInstancesOfferings()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) GetReservedInstancesOfferings() []*ServicesEc2ModelReservedInstancesOffering {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesOfferings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelReservedInstancesOffering)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setReservedInstancesOfferings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesOffering>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) SetReservedInstancesOfferings(a []*ServicesEc2ModelReservedInstancesOffering)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesOfferings", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult withReservedInstancesOfferings(com.amazonaws.services.ec2.model.ReservedInstancesOffering...)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) WithReservedInstancesOfferings(a ...*ServicesEc2ModelReservedInstancesOffering) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ReservedInstancesOffering")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesOfferings", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstancesOffering"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult withReservedInstancesOfferings(java.util.Collection<com.amazonaws.services.ec2.model.ReservedInstancesOffering>)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) WithReservedInstancesOfferings2(a []*ServicesEc2ModelReservedInstancesOffering) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesOfferings", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) WithNextToken(a string) *ServicesEc2ModelDescribeReservedInstancesOfferingsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeReservedInstancesOfferingsResult clone()
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) Clone() *ServicesEc2ModelDescribeReservedInstancesOfferingsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeReservedInstancesOfferingsResult")
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
	unique_x := &ServicesEc2ModelDescribeReservedInstancesOfferingsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeReservedInstancesOfferingsResult) Clone2() (*JavaLangObject, error) {
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


