package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateCustomerGatewayResultInterface interface {
	JavaLangObjectInterface

	// public void setCustomerGateway(com.amazonaws.services.ec2.model.CustomerGateway)
	SetCustomerGateway(a ServicesEc2ModelCustomerGatewayInterface) 

	// public com.amazonaws.services.ec2.model.CustomerGateway getCustomerGateway()
	GetCustomerGateway() *ServicesEc2ModelCustomerGateway

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult withCustomerGateway(com.amazonaws.services.ec2.model.CustomerGateway)
	WithCustomerGateway(a ServicesEc2ModelCustomerGatewayInterface) *ServicesEc2ModelCreateCustomerGatewayResult

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult clone()
	Clone() *ServicesEc2ModelCreateCustomerGatewayResult
}

type ServicesEc2ModelCreateCustomerGatewayResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult()
func NewServicesEc2ModelCreateCustomerGatewayResult() (*ServicesEc2ModelCreateCustomerGatewayResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateCustomerGatewayResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateCustomerGatewayResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCustomerGateway(com.amazonaws.services.ec2.model.CustomerGateway)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) SetCustomerGateway(a ServicesEc2ModelCustomerGatewayInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCustomerGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CustomerGateway"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CustomerGateway getCustomerGateway()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) GetCustomerGateway() *ServicesEc2ModelCustomerGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomerGateway", "com/amazonaws/services/ec2/model/CustomerGateway")
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
	unique_x := &ServicesEc2ModelCustomerGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult withCustomerGateway(com.amazonaws.services.ec2.model.CustomerGateway)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) WithCustomerGateway(a ServicesEc2ModelCustomerGatewayInterface) *ServicesEc2ModelCreateCustomerGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGateway", "com/amazonaws/services/ec2/model/CreateCustomerGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CustomerGateway"))
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayResult clone()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) Clone() *ServicesEc2ModelCreateCustomerGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateCustomerGatewayResult")
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateCustomerGatewayResult) Clone2() (*JavaLangObject, error) {
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


