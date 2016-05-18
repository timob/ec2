package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpnConnectionOptionsSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setStaticRoutesOnly(java.lang.Boolean)
	SetStaticRoutesOnly(a bool) 

	// public java.lang.Boolean getStaticRoutesOnly()
	GetStaticRoutesOnly() bool

	// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification withStaticRoutesOnly(java.lang.Boolean)
	WithStaticRoutesOnly(a bool) *ServicesEc2ModelVpnConnectionOptionsSpecification

	// public java.lang.Boolean isStaticRoutesOnly()
	IsStaticRoutesOnly() bool

	// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification clone()
	Clone() *ServicesEc2ModelVpnConnectionOptionsSpecification
}

type ServicesEc2ModelVpnConnectionOptionsSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification()
func NewServicesEc2ModelVpnConnectionOptionsSpecification() (*ServicesEc2ModelVpnConnectionOptionsSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VpnConnectionOptionsSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpnConnectionOptionsSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setStaticRoutesOnly(java.lang.Boolean)
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) SetStaticRoutesOnly(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStaticRoutesOnly", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getStaticRoutesOnly()
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) GetStaticRoutesOnly() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStaticRoutesOnly", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification withStaticRoutesOnly(java.lang.Boolean)
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) WithStaticRoutesOnly(a bool) *ServicesEc2ModelVpnConnectionOptionsSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStaticRoutesOnly", "com/amazonaws/services/ec2/model/VpnConnectionOptionsSpecification", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelVpnConnectionOptionsSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isStaticRoutesOnly()
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) IsStaticRoutesOnly() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isStaticRoutesOnly", "java/lang/Boolean")
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
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification clone()
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) Clone() *ServicesEc2ModelVpnConnectionOptionsSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VpnConnectionOptionsSpecification")
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
	unique_x := &ServicesEc2ModelVpnConnectionOptionsSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpnConnectionOptionsSpecification) Clone2() (*JavaLangObject, error) {
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


