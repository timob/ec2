package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpnConnectionsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.VpnConnection> getVpnConnections()
	GetVpnConnections() []*ServicesEc2ModelVpnConnection

	// public void setVpnConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpnConnection>)
	SetVpnConnections(a []*ServicesEc2ModelVpnConnection) 

	// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult withVpnConnections(com.amazonaws.services.ec2.model.VpnConnection...)
	WithVpnConnections(a ...*ServicesEc2ModelVpnConnection) *ServicesEc2ModelDescribeVpnConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult withVpnConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpnConnection>)
	WithVpnConnections2(a []*ServicesEc2ModelVpnConnection) *ServicesEc2ModelDescribeVpnConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult clone()
	Clone() *ServicesEc2ModelDescribeVpnConnectionsResult
}

type ServicesEc2ModelDescribeVpnConnectionsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult()
func NewServicesEc2ModelDescribeVpnConnectionsResult() (*ServicesEc2ModelDescribeVpnConnectionsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpnConnectionsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpnConnectionsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.VpnConnection> getVpnConnections()
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) GetVpnConnections() []*ServicesEc2ModelVpnConnection {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnConnections", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpnConnection)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpnConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpnConnection>)
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) SetVpnConnections(a []*ServicesEc2ModelVpnConnection)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnConnections", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult withVpnConnections(com.amazonaws.services.ec2.model.VpnConnection...)
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) WithVpnConnections(a ...*ServicesEc2ModelVpnConnection) *ServicesEc2ModelDescribeVpnConnectionsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VpnConnection")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnections", "com/amazonaws/services/ec2/model/DescribeVpnConnectionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnection"))
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
	unique_x := &ServicesEc2ModelDescribeVpnConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult withVpnConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpnConnection>)
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) WithVpnConnections2(a []*ServicesEc2ModelVpnConnection) *ServicesEc2ModelDescribeVpnConnectionsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnections", "com/amazonaws/services/ec2/model/DescribeVpnConnectionsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpnConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpnConnectionsResult clone()
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) Clone() *ServicesEc2ModelDescribeVpnConnectionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpnConnectionsResult")
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
	unique_x := &ServicesEc2ModelDescribeVpnConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpnConnectionsResult) Clone2() (*JavaLangObject, error) {
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


