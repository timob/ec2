package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotFleetRequestsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.SpotFleetRequestConfig> getSpotFleetRequestConfigs()
	GetSpotFleetRequestConfigs() []*ServicesEc2ModelSpotFleetRequestConfig

	// public void setSpotFleetRequestConfigs(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetRequestConfig>)
	SetSpotFleetRequestConfigs(a []*ServicesEc2ModelSpotFleetRequestConfig) 

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult withSpotFleetRequestConfigs(com.amazonaws.services.ec2.model.SpotFleetRequestConfig...)
	WithSpotFleetRequestConfigs(a ...*ServicesEc2ModelSpotFleetRequestConfig) *ServicesEc2ModelDescribeSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult withSpotFleetRequestConfigs(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetRequestConfig>)
	WithSpotFleetRequestConfigs2(a []*ServicesEc2ModelSpotFleetRequestConfig) *ServicesEc2ModelDescribeSpotFleetRequestsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult clone()
	Clone() *ServicesEc2ModelDescribeSpotFleetRequestsResult
}

type ServicesEc2ModelDescribeSpotFleetRequestsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult()
func NewServicesEc2ModelDescribeSpotFleetRequestsResult() (*ServicesEc2ModelDescribeSpotFleetRequestsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.SpotFleetRequestConfig> getSpotFleetRequestConfigs()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) GetSpotFleetRequestConfigs() []*ServicesEc2ModelSpotFleetRequestConfig {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestConfigs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSpotFleetRequestConfig)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSpotFleetRequestConfigs(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetRequestConfig>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) SetSpotFleetRequestConfigs(a []*ServicesEc2ModelSpotFleetRequestConfig)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestConfigs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult withSpotFleetRequestConfigs(com.amazonaws.services.ec2.model.SpotFleetRequestConfig...)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) WithSpotFleetRequestConfigs(a ...*ServicesEc2ModelSpotFleetRequestConfig) *ServicesEc2ModelDescribeSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/SpotFleetRequestConfig")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestConfigs", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotFleetRequestConfig"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult withSpotFleetRequestConfigs(java.util.Collection<com.amazonaws.services.ec2.model.SpotFleetRequestConfig>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) WithSpotFleetRequestConfigs2(a []*ServicesEc2ModelSpotFleetRequestConfig) *ServicesEc2ModelDescribeSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestConfigs", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetRequestsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetRequestsResult clone()
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) Clone() *ServicesEc2ModelDescribeSpotFleetRequestsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotFleetRequestsResult")
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSpotFleetRequestsResult) Clone2() (*JavaLangObject, error) {
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


