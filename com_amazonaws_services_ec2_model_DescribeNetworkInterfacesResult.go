package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeNetworkInterfacesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.NetworkInterface> getNetworkInterfaces()
	GetNetworkInterfaces() []*ServicesEc2ModelNetworkInterface

	// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterface>)
	SetNetworkInterfaces(a []*ServicesEc2ModelNetworkInterface) 

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult withNetworkInterfaces(com.amazonaws.services.ec2.model.NetworkInterface...)
	WithNetworkInterfaces(a ...*ServicesEc2ModelNetworkInterface) *ServicesEc2ModelDescribeNetworkInterfacesResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterface>)
	WithNetworkInterfaces2(a []*ServicesEc2ModelNetworkInterface) *ServicesEc2ModelDescribeNetworkInterfacesResult

	// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult clone()
	Clone() *ServicesEc2ModelDescribeNetworkInterfacesResult
}

type ServicesEc2ModelDescribeNetworkInterfacesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult()
func NewServicesEc2ModelDescribeNetworkInterfacesResult() (*ServicesEc2ModelDescribeNetworkInterfacesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeNetworkInterfacesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeNetworkInterfacesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.NetworkInterface> getNetworkInterfaces()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) GetNetworkInterfaces() []*ServicesEc2ModelNetworkInterface {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaces", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNetworkInterface)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterface>)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) SetNetworkInterfaces(a []*ServicesEc2ModelNetworkInterface)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterfaces", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult withNetworkInterfaces(com.amazonaws.services.ec2.model.NetworkInterface...)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) WithNetworkInterfaces(a ...*ServicesEc2ModelNetworkInterface) *ServicesEc2ModelDescribeNetworkInterfacesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NetworkInterface")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/DescribeNetworkInterfacesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterface"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfacesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult withNetworkInterfaces(java.util.Collection<com.amazonaws.services.ec2.model.NetworkInterface>)
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) WithNetworkInterfaces2(a []*ServicesEc2ModelNetworkInterface) *ServicesEc2ModelDescribeNetworkInterfacesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaces", "com/amazonaws/services/ec2/model/DescribeNetworkInterfacesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfacesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeNetworkInterfacesResult clone()
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) Clone() *ServicesEc2ModelDescribeNetworkInterfacesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeNetworkInterfacesResult")
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
	unique_x := &ServicesEc2ModelDescribeNetworkInterfacesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeNetworkInterfacesResult) Clone2() (*JavaLangObject, error) {
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


