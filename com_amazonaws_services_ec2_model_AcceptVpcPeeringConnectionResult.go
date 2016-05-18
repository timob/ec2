package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAcceptVpcPeeringConnectionResultInterface interface {
	JavaLangObjectInterface

	// public void setVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
	SetVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnection getVpcPeeringConnection()
	GetVpcPeeringConnection() *ServicesEc2ModelVpcPeeringConnection

	// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult withVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
	WithVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface) *ServicesEc2ModelAcceptVpcPeeringConnectionResult

	// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult clone()
	Clone() *ServicesEc2ModelAcceptVpcPeeringConnectionResult
}

type ServicesEc2ModelAcceptVpcPeeringConnectionResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult()
func NewServicesEc2ModelAcceptVpcPeeringConnectionResult() (*ServicesEc2ModelAcceptVpcPeeringConnectionResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AcceptVpcPeeringConnectionResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAcceptVpcPeeringConnectionResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) SetVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface)  {
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
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) GetVpcPeeringConnection() *ServicesEc2ModelVpcPeeringConnection {
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

// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult withVpcPeeringConnection(com.amazonaws.services.ec2.model.VpcPeeringConnection)
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) WithVpcPeeringConnection(a ServicesEc2ModelVpcPeeringConnectionInterface) *ServicesEc2ModelAcceptVpcPeeringConnectionResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnection", "com/amazonaws/services/ec2/model/AcceptVpcPeeringConnectionResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnection"))
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
	unique_x := &ServicesEc2ModelAcceptVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AcceptVpcPeeringConnectionResult clone()
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) Clone() *ServicesEc2ModelAcceptVpcPeeringConnectionResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AcceptVpcPeeringConnectionResult")
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
	unique_x := &ServicesEc2ModelAcceptVpcPeeringConnectionResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAcceptVpcPeeringConnectionResult) Clone2() (*JavaLangObject, error) {
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


