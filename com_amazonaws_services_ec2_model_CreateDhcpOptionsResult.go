package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateDhcpOptionsResultInterface interface {
	JavaLangObjectInterface

	// public void setDhcpOptions(com.amazonaws.services.ec2.model.DhcpOptions)
	SetDhcpOptions(a ServicesEc2ModelDhcpOptionsInterface) 

	// public com.amazonaws.services.ec2.model.DhcpOptions getDhcpOptions()
	GetDhcpOptions() *ServicesEc2ModelDhcpOptions

	// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult withDhcpOptions(com.amazonaws.services.ec2.model.DhcpOptions)
	WithDhcpOptions(a ServicesEc2ModelDhcpOptionsInterface) *ServicesEc2ModelCreateDhcpOptionsResult

	// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult clone()
	Clone() *ServicesEc2ModelCreateDhcpOptionsResult
}

type ServicesEc2ModelCreateDhcpOptionsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult()
func NewServicesEc2ModelCreateDhcpOptionsResult() (*ServicesEc2ModelCreateDhcpOptionsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateDhcpOptionsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateDhcpOptionsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDhcpOptions(com.amazonaws.services.ec2.model.DhcpOptions)
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) SetDhcpOptions(a ServicesEc2ModelDhcpOptionsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDhcpOptions", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DhcpOptions"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DhcpOptions getDhcpOptions()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) GetDhcpOptions() *ServicesEc2ModelDhcpOptions {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDhcpOptions", "com/amazonaws/services/ec2/model/DhcpOptions")
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult withDhcpOptions(com.amazonaws.services.ec2.model.DhcpOptions)
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) WithDhcpOptions(a ServicesEc2ModelDhcpOptionsInterface) *ServicesEc2ModelCreateDhcpOptionsResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpOptions", "com/amazonaws/services/ec2/model/CreateDhcpOptionsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DhcpOptions"))
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
	unique_x := &ServicesEc2ModelCreateDhcpOptionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateDhcpOptionsResult clone()
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) Clone() *ServicesEc2ModelCreateDhcpOptionsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateDhcpOptionsResult")
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
	unique_x := &ServicesEc2ModelCreateDhcpOptionsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateDhcpOptionsResult) Clone2() (*JavaLangObject, error) {
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


