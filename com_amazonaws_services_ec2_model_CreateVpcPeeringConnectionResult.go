package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpcPeeringConnectionResultInterface interface {
	JavaLangObjectInterface

	// public void setVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
	SetVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection getVpcPeeringConnection()
	GetVpcPeeringConnection() *ServicesEc2ModelVpcPeeringConnection

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult withVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
	WithVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface) *ServicesEc2ModelCreateVpcPeeringConnectionResult

	// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult clone()
	Clone() *ServicesEc2ModelCreateVpcPeeringConnectionResult
}

type ServicesEc2ModelCreateVpcPeeringConnectionResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult()
func NewServicesEc2ModelCreateVpcPeeringConnectionResult() (*ServicesEc2ModelCreateVpcPeeringConnectionResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpcPeeringConnectionResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) SetVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcPeeringConnection", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcPeeringConnection getVpcPeeringConnection()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) GetVpcPeeringConnection() *ServicesEc2ModelVpcPeeringConnection {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcPeeringConnection", "com/amazonaws/services/ec2/model/VpcPeeringConnection")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult withVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) WithVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface) *ServicesEc2ModelCreateVpcPeeringConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnection", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnection"))
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
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpcPeeringConnectionResult clone()
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) Clone() *ServicesEc2ModelCreateVpcPeeringConnectionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpcPeeringConnectionResult")
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
	unique_x := &ServicesEc2ModelCreateVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVpcPeeringConnectionResult) Clone2() (*JavaLangObject, error) {
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


