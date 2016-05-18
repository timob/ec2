package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeFlowLogsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.FlowLog> getFlowLogs()
	GetFlowLogs() []*ServicesEc2ModelFlowLog

	// public void setFlowLogs(java.util.Collection<com.amazonaws.services.ec2.model.FlowLog>)
	SetFlowLogs(a []*ServicesEc2ModelFlowLog) 

	// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult withFlowLogs(com.amazonaws.services.ec2.model.FlowLog...)
	WithFlowLogs(a ...*ServicesEc2ModelFlowLog) *ServicesEc2ModelDescribeFlowLogsResult

	// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult withFlowLogs(java.util.Collection<com.amazonaws.services.ec2.model.FlowLog>)
	WithFlowLogs2(a []*ServicesEc2ModelFlowLog) *ServicesEc2ModelDescribeFlowLogsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeFlowLogsResult

	// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult clone()
	Clone() *ServicesEc2ModelDescribeFlowLogsResult
}

type ServicesEc2ModelDescribeFlowLogsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult()
func NewServicesEc2ModelDescribeFlowLogsResult() (*ServicesEc2ModelDescribeFlowLogsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeFlowLogsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeFlowLogsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.FlowLog> getFlowLogs()
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) GetFlowLogs() []*ServicesEc2ModelFlowLog {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFlowLogs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelFlowLog)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setFlowLogs(java.util.Collection<com.amazonaws.services.ec2.model.FlowLog>)
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) SetFlowLogs(a []*ServicesEc2ModelFlowLog)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFlowLogs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult withFlowLogs(com.amazonaws.services.ec2.model.FlowLog...)
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) WithFlowLogs(a ...*ServicesEc2ModelFlowLog) *ServicesEc2ModelDescribeFlowLogsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/FlowLog")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFlowLogs", "com/amazonaws/services/ec2/model/DescribeFlowLogsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/FlowLog"))
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
	unique_x := &ServicesEc2ModelDescribeFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult withFlowLogs(java.util.Collection<com.amazonaws.services.ec2.model.FlowLog>)
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) WithFlowLogs2(a []*ServicesEc2ModelFlowLog) *ServicesEc2ModelDescribeFlowLogsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFlowLogs", "com/amazonaws/services/ec2/model/DescribeFlowLogsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) WithNextToken(a string) *ServicesEc2ModelDescribeFlowLogsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeFlowLogsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeFlowLogsResult clone()
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) Clone() *ServicesEc2ModelDescribeFlowLogsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeFlowLogsResult")
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
	unique_x := &ServicesEc2ModelDescribeFlowLogsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeFlowLogsResult) Clone2() (*JavaLangObject, error) {
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


