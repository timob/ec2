package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPrivateIpAddressSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelPrivateIpAddressSpecification

	// public void setPrimary(java.lang.Boolean)
	SetPrimary(a bool) 

	// public java.lang.Boolean getPrimary()
	GetPrimary() bool

	// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification withPrimary(java.lang.Boolean)
	WithPrimary(a bool) *ServicesEc2ModelPrivateIpAddressSpecification

	// public java.lang.Boolean isPrimary()
	IsPrimary() bool

	// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification clone()
	Clone() *ServicesEc2ModelPrivateIpAddressSpecification
}

type ServicesEc2ModelPrivateIpAddressSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification()
func NewServicesEc2ModelPrivateIpAddressSpecification() (*ServicesEc2ModelPrivateIpAddressSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PrivateIpAddressSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPrivateIpAddressSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) SetPrivateIpAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrivateIpAddress()
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) GetPrivateIpAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) WithPrivateIpAddress(a string) *ServicesEc2ModelPrivateIpAddressSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/PrivateIpAddressSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPrivateIpAddressSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) SetPrimary(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrimary", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getPrimary()
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) GetPrimary() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrimary", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification withPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) WithPrimary(a bool) *ServicesEc2ModelPrivateIpAddressSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrimary", "com/amazonaws/services/ec2/model/PrivateIpAddressSpecification", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelPrivateIpAddressSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isPrimary()
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) IsPrimary() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPrimary", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PrivateIpAddressSpecification clone()
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) Clone() *ServicesEc2ModelPrivateIpAddressSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PrivateIpAddressSpecification")
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
	unique_x := &ServicesEc2ModelPrivateIpAddressSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPrivateIpAddressSpecification) Clone2() (*JavaLangObject, error) {
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


