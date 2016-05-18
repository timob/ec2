package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeRouteTablesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.RouteTable> getRouteTables()
	GetRouteTables() []*ServicesEc2ModelRouteTable

	// public void setRouteTables(java.util.Collection<com.amazonaws.services.ec2.model.RouteTable>)
	SetRouteTables(a []*ServicesEc2ModelRouteTable) 

	// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult withRouteTables(com.amazonaws.services.ec2.model.RouteTable...)
	WithRouteTables(a ...*ServicesEc2ModelRouteTable) *ServicesEc2ModelDescribeRouteTablesResult

	// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult withRouteTables(java.util.Collection<com.amazonaws.services.ec2.model.RouteTable>)
	WithRouteTables2(a []*ServicesEc2ModelRouteTable) *ServicesEc2ModelDescribeRouteTablesResult

	// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult clone()
	Clone() *ServicesEc2ModelDescribeRouteTablesResult
}

type ServicesEc2ModelDescribeRouteTablesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult()
func NewServicesEc2ModelDescribeRouteTablesResult() (*ServicesEc2ModelDescribeRouteTablesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeRouteTablesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeRouteTablesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.RouteTable> getRouteTables()
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) GetRouteTables() []*ServicesEc2ModelRouteTable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTables", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelRouteTable)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRouteTables(java.util.Collection<com.amazonaws.services.ec2.model.RouteTable>)
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) SetRouteTables(a []*ServicesEc2ModelRouteTable)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTables", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult withRouteTables(com.amazonaws.services.ec2.model.RouteTable...)
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) WithRouteTables(a ...*ServicesEc2ModelRouteTable) *ServicesEc2ModelDescribeRouteTablesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/RouteTable")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTables", "com/amazonaws/services/ec2/model/DescribeRouteTablesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteTable"))
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
	unique_x := &ServicesEc2ModelDescribeRouteTablesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult withRouteTables(java.util.Collection<com.amazonaws.services.ec2.model.RouteTable>)
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) WithRouteTables2(a []*ServicesEc2ModelRouteTable) *ServicesEc2ModelDescribeRouteTablesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTables", "com/amazonaws/services/ec2/model/DescribeRouteTablesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeRouteTablesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeRouteTablesResult clone()
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) Clone() *ServicesEc2ModelDescribeRouteTablesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeRouteTablesResult")
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
	unique_x := &ServicesEc2ModelDescribeRouteTablesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeRouteTablesResult) Clone2() (*JavaLangObject, error) {
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


