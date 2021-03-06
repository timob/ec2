package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcPeeringConnectionsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.VpcPeeringConnection> getVpcPeeringConnections()
	GetVpcPeeringConnections() []*ServicesEc2ModelVpcPeeringConnection

	// public void setVpcPeeringConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpcPeeringConnection>)
	SetVpcPeeringConnections(a []*ServicesEc2ModelVpcPeeringConnection) 

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult withVpcPeeringConnections(com.amazonaws.services.ec2.model.VpcPeeringConnection...)
	WithVpcPeeringConnections(a ...*ServicesEc2ModelVpcPeeringConnection) *ServicesEc2ModelDescribeVpcPeeringConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult withVpcPeeringConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpcPeeringConnection>)
	WithVpcPeeringConnections2(a []*ServicesEc2ModelVpcPeeringConnection) *ServicesEc2ModelDescribeVpcPeeringConnectionsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult clone()
	Clone() *ServicesEc2ModelDescribeVpcPeeringConnectionsResult
}

type ServicesEc2ModelDescribeVpcPeeringConnectionsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult()
func NewServicesEc2ModelDescribeVpcPeeringConnectionsResult() (*ServicesEc2ModelDescribeVpcPeeringConnectionsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcPeeringConnectionsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.VpcPeeringConnection> getVpcPeeringConnections()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) GetVpcPeeringConnections() []*ServicesEc2ModelVpcPeeringConnection {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcPeeringConnections", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpcPeeringConnection)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpcPeeringConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpcPeeringConnection>)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) SetVpcPeeringConnections(a []*ServicesEc2ModelVpcPeeringConnection)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcPeeringConnections", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult withVpcPeeringConnections(com.amazonaws.services.ec2.model.VpcPeeringConnection...)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) WithVpcPeeringConnections(a ...*ServicesEc2ModelVpcPeeringConnection) *ServicesEc2ModelDescribeVpcPeeringConnectionsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VpcPeeringConnection")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnections", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult withVpcPeeringConnections(java.util.Collection<com.amazonaws.services.ec2.model.VpcPeeringConnection>)
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) WithVpcPeeringConnections2(a []*ServicesEc2ModelVpcPeeringConnection) *ServicesEc2ModelDescribeVpcPeeringConnectionsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnections", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcPeeringConnectionsResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) Clone() *ServicesEc2ModelDescribeVpcPeeringConnectionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcPeeringConnectionsResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcPeeringConnectionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcPeeringConnectionsResult) Clone2() (*JavaLangObject, error) {
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


