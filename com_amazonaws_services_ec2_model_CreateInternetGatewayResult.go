package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateInternetGatewayResultInterface interface {
	JavaLangObjectInterface

	// public void setInternetGateway(com.amazonaws.services.ec2.model.InternetGateway)
	SetInternetGateway(a ServicesEc2ModelInternetGatewayInterface) 

	// public com.amazonaws.services.ec2.model.InternetGateway getInternetGateway()
	GetInternetGateway() *ServicesEc2ModelInternetGateway

	// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult withInternetGateway(com.amazonaws.services.ec2.model.InternetGateway)
	WithInternetGateway(a ServicesEc2ModelInternetGatewayInterface) *ServicesEc2ModelCreateInternetGatewayResult

	// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult clone()
	Clone() *ServicesEc2ModelCreateInternetGatewayResult
}

type ServicesEc2ModelCreateInternetGatewayResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult()
func NewServicesEc2ModelCreateInternetGatewayResult() (*ServicesEc2ModelCreateInternetGatewayResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateInternetGatewayResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateInternetGatewayResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInternetGateway(com.amazonaws.services.ec2.model.InternetGateway)
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) SetInternetGateway(a ServicesEc2ModelInternetGatewayInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInternetGateway", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InternetGateway"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InternetGateway getInternetGateway()
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) GetInternetGateway() *ServicesEc2ModelInternetGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInternetGateway", "com/amazonaws/services/ec2/model/InternetGateway")
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult withInternetGateway(com.amazonaws.services.ec2.model.InternetGateway)
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) WithInternetGateway(a ServicesEc2ModelInternetGatewayInterface) *ServicesEc2ModelCreateInternetGatewayResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInternetGateway", "com/amazonaws/services/ec2/model/CreateInternetGatewayResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InternetGateway"))
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
	unique_x := &ServicesEc2ModelCreateInternetGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateInternetGatewayResult clone()
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) Clone() *ServicesEc2ModelCreateInternetGatewayResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateInternetGatewayResult")
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
	unique_x := &ServicesEc2ModelCreateInternetGatewayResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateInternetGatewayResult) Clone2() (*JavaLangObject, error) {
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


