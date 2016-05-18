package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateNatGatewayResultInterface interface {
	JavaLangObjectInterface

	// public void setNatGateway(com.amazonaws.services.ec2.model.NatGateway)
	SetNatGateway(a ServicesEc2ModelNatGatewayInterface) 

	// public com.amazonaws.services.ec2.model.NatGateway getNatGateway()
	GetNatGateway() *ServicesEc2ModelNatGateway

	// public com.amazonaws.services.ec2.model.CreateNatGatewayResult withNatGateway(com.amazonaws.services.ec2.model.NatGateway)
	WithNatGateway(a ServicesEc2ModelNatGatewayInterface) *ServicesEc2ModelCreateNatGatewayResult

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CreateNatGatewayResult withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCreateNatGatewayResult

	// public com.amazonaws.services.ec2.model.CreateNatGatewayResult clone()
	Clone() *ServicesEc2ModelCreateNatGatewayResult
}

type ServicesEc2ModelCreateNatGatewayResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateNatGatewayResult()
func NewServicesEc2ModelCreateNatGatewayResult() (*ServicesEc2ModelCreateNatGatewayResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateNatGatewayResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateNatGatewayResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNatGateway(com.amazonaws.services.ec2.model.NatGateway)
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) SetNatGateway(a ServicesEc2ModelNatGatewayInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNatGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NatGateway"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NatGateway getNatGateway()
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) GetNatGateway() *ServicesEc2ModelNatGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNatGateway", "com/amazonaws/services/ec2/model/NatGateway")
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
	unique_x := &ServicesEc2ModelNatGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNatGatewayResult withNatGateway(com.amazonaws.services.ec2.model.NatGateway)
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) WithNatGateway(a ServicesEc2ModelNatGatewayInterface) *ServicesEc2ModelCreateNatGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGateway", "com/amazonaws/services/ec2/model/CreateNatGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NatGateway"))
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
	unique_x := &ServicesEc2ModelCreateNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNatGatewayResult withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) WithClientToken(a string) *ServicesEc2ModelCreateNatGatewayResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CreateNatGatewayResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateNatGatewayResult clone()
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) Clone() *ServicesEc2ModelCreateNatGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateNatGatewayResult")
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
	unique_x := &ServicesEc2ModelCreateNatGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateNatGatewayResult) Clone2() (*JavaLangObject, error) {
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


