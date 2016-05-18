package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotFleetInstancesResultInterface interface {
	JavaLangObjectInterface

	// public void setSpotFleetRequestId(java.lang.String)
	SetSpotFleetRequestId(a string) 

	// public java.lang.String getSpotFleetRequestId()
	GetSpotFleetRequestId() string

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withSpotFleetRequestId(java.lang.String)
	WithSpotFleetRequestId(a string) *ServicesEc2ModelDescribeSpotFleetInstancesResult

	// public java.util.List<com.amazonaws.services.ec2.model.ActiveInstance> getActiveInstances()
	GetActiveInstances() []*ServicesEc2ModelActiveInstance

	// public void setActiveInstances(java.util.Collection<com.amazonaws.services.ec2.model.ActiveInstance>)
	SetActiveInstances(a []*ServicesEc2ModelActiveInstance) 

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withActiveInstances(com.amazonaws.services.ec2.model.ActiveInstance...)
	WithActiveInstances(a ...*ServicesEc2ModelActiveInstance) *ServicesEc2ModelDescribeSpotFleetInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withActiveInstances(java.util.Collection<com.amazonaws.services.ec2.model.ActiveInstance>)
	WithActiveInstances2(a []*ServicesEc2ModelActiveInstance) *ServicesEc2ModelDescribeSpotFleetInstancesResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetInstancesResult

	// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult clone()
	Clone() *ServicesEc2ModelDescribeSpotFleetInstancesResult
}

type ServicesEc2ModelDescribeSpotFleetInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult()
func NewServicesEc2ModelDescribeSpotFleetInstancesResult() (*ServicesEc2ModelDescribeSpotFleetInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) SetSpotFleetRequestId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSpotFleetRequestId()
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) GetSpotFleetRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withSpotFleetRequestId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) WithSpotFleetRequestId(a string) *ServicesEc2ModelDescribeSpotFleetInstancesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestId", "com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ActiveInstance> getActiveInstances()
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) GetActiveInstances() []*ServicesEc2ModelActiveInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActiveInstances", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelActiveInstance)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setActiveInstances(java.util.Collection<com.amazonaws.services.ec2.model.ActiveInstance>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) SetActiveInstances(a []*ServicesEc2ModelActiveInstance)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setActiveInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withActiveInstances(com.amazonaws.services.ec2.model.ActiveInstance...)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) WithActiveInstances(a ...*ServicesEc2ModelActiveInstance) *ServicesEc2ModelDescribeSpotFleetInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ActiveInstance")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withActiveInstances", "com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ActiveInstance"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withActiveInstances(java.util.Collection<com.amazonaws.services.ec2.model.ActiveInstance>)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) WithActiveInstances2(a []*ServicesEc2ModelActiveInstance) *ServicesEc2ModelDescribeSpotFleetInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withActiveInstances", "com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) WithNextToken(a string) *ServicesEc2ModelDescribeSpotFleetInstancesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotFleetInstancesResult clone()
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) Clone() *ServicesEc2ModelDescribeSpotFleetInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotFleetInstancesResult")
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
	unique_x := &ServicesEc2ModelDescribeSpotFleetInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSpotFleetInstancesResult) Clone2() (*JavaLangObject, error) {
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


