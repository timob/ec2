package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotPriceHistoryResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.SpotPrice> getSpotPriceHistory()
	GetSpotPriceHistory() []*ServicesEc2ModelSpotPrice

	// public void setSpotPriceHistory(java.util.Collection<com.amazonaws.services.ec2.model.SpotPrice>)
	SetSpotPriceHistory(a []*ServicesEc2ModelSpotPrice) 

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult withSpotPriceHistory(com.amazonaws.services.ec2.model.SpotPrice...)
	WithSpotPriceHistory(a ...*ServicesEc2ModelSpotPrice) *ServicesEc2ModelDescribeSpotPriceHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult withSpotPriceHistory(java.util.Collection<com.amazonaws.services.ec2.model.SpotPrice>)
	WithSpotPriceHistory2(a []*ServicesEc2ModelSpotPrice) *ServicesEc2ModelDescribeSpotPriceHistoryResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSpotPriceHistoryResult

	// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult clone()
	Clone() *ServicesEc2ModelDescribeSpotPriceHistoryResult
}

type ServicesEc2ModelDescribeSpotPriceHistoryResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult()
func NewServicesEc2ModelDescribeSpotPriceHistoryResult() (*ServicesEc2ModelDescribeSpotPriceHistoryResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.SpotPrice> getSpotPriceHistory()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) GetSpotPriceHistory() []*ServicesEc2ModelSpotPrice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotPriceHistory", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSpotPrice)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSpotPriceHistory(java.util.Collection<com.amazonaws.services.ec2.model.SpotPrice>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) SetSpotPriceHistory(a []*ServicesEc2ModelSpotPrice)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotPriceHistory", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult withSpotPriceHistory(com.amazonaws.services.ec2.model.SpotPrice...)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) WithSpotPriceHistory(a ...*ServicesEc2ModelSpotPrice) *ServicesEc2ModelDescribeSpotPriceHistoryResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/SpotPrice")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotPriceHistory", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotPrice"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult withSpotPriceHistory(java.util.Collection<com.amazonaws.services.ec2.model.SpotPrice>)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) WithSpotPriceHistory2(a []*ServicesEc2ModelSpotPrice) *ServicesEc2ModelDescribeSpotPriceHistoryResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotPriceHistory", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) WithNextToken(a string) *ServicesEc2ModelDescribeSpotPriceHistoryResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotPriceHistoryResult clone()
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) Clone() *ServicesEc2ModelDescribeSpotPriceHistoryResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotPriceHistoryResult")
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
	unique_x := &ServicesEc2ModelDescribeSpotPriceHistoryResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSpotPriceHistoryResult) Clone2() (*JavaLangObject, error) {
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


