package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateSubnetResultInterface interface {
	JavaLangObjectInterface

	// public void setSubnet(com.amazonaws.services.ec2.model.Subnet)
	SetSubnet(a ServicesEc2ModelSubnetInterface) 

	// public com.amazonaws.services.ec2.model.Subnet getSubnet()
	GetSubnet() *ServicesEc2ModelSubnet

	// public com.amazonaws.services.ec2.model.CreateSubnetResult withSubnet(com.amazonaws.services.ec2.model.Subnet)
	WithSubnet(a ServicesEc2ModelSubnetInterface) *ServicesEc2ModelCreateSubnetResult

	// public com.amazonaws.services.ec2.model.CreateSubnetResult clone()
	Clone() *ServicesEc2ModelCreateSubnetResult
}

type ServicesEc2ModelCreateSubnetResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateSubnetResult()
func NewServicesEc2ModelCreateSubnetResult() (*ServicesEc2ModelCreateSubnetResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateSubnetResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateSubnetResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSubnet(com.amazonaws.services.ec2.model.Subnet)
func (jbobject *ServicesEc2ModelCreateSubnetResult) SetSubnet(a ServicesEc2ModelSubnetInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnet", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Subnet"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Subnet getSubnet()
func (jbobject *ServicesEc2ModelCreateSubnetResult) GetSubnet() *ServicesEc2ModelSubnet {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnet", "com/amazonaws/services/ec2/model/Subnet")
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
	unique_x := &ServicesEc2ModelSubnet{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSubnetResult withSubnet(com.amazonaws.services.ec2.model.Subnet)
func (jbobject *ServicesEc2ModelCreateSubnetResult) WithSubnet(a ServicesEc2ModelSubnetInterface) *ServicesEc2ModelCreateSubnetResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnet", "com/amazonaws/services/ec2/model/CreateSubnetResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Subnet"))
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
	unique_x := &ServicesEc2ModelCreateSubnetResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateSubnetResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateSubnetResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateSubnetResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateSubnetResult clone()
func (jbobject *ServicesEc2ModelCreateSubnetResult) Clone() *ServicesEc2ModelCreateSubnetResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateSubnetResult")
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
	unique_x := &ServicesEc2ModelCreateSubnetResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateSubnetResult) Clone2() (*JavaLangObject, error) {
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


