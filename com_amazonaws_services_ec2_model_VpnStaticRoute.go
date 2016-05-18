package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpnStaticRouteInterface interface {
	JavaLangObjectInterface

	// public void setDestinationCidrBlock(java.lang.String)
	SetDestinationCidrBlock(a string) 

	// public java.lang.String getDestinationCidrBlock()
	GetDestinationCidrBlock() string

	// public com.amazonaws.services.ec2.model.VpnStaticRoute withDestinationCidrBlock(java.lang.String)
	WithDestinationCidrBlock(a string) *ServicesEc2ModelVpnStaticRoute

	// public void setSource(java.lang.String)
	SetSource2(a string) 

	// public java.lang.String getSource()
	GetSource() string

	// public com.amazonaws.services.ec2.model.VpnStaticRoute withSource(java.lang.String)
	WithSource2(a string) *ServicesEc2ModelVpnStaticRoute

	// public void setSource(com.amazonaws.services.ec2.model.VpnStaticRouteSource)
	SetSource(a ServicesEc2ModelVpnStaticRouteSourceInterface) 

	// public com.amazonaws.services.ec2.model.VpnStaticRoute withSource(com.amazonaws.services.ec2.model.VpnStaticRouteSource)
	WithSource(a ServicesEc2ModelVpnStaticRouteSourceInterface) *ServicesEc2ModelVpnStaticRoute

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.VpnStaticRoute withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelVpnStaticRoute

	// public void setState(com.amazonaws.services.ec2.model.VpnState)
	SetState(a ServicesEc2ModelVpnStateInterface) 

	// public com.amazonaws.services.ec2.model.VpnStaticRoute withState(com.amazonaws.services.ec2.model.VpnState)
	WithState(a ServicesEc2ModelVpnStateInterface) *ServicesEc2ModelVpnStaticRoute

	// public com.amazonaws.services.ec2.model.VpnStaticRoute clone()
	Clone() *ServicesEc2ModelVpnStaticRoute
}

type ServicesEc2ModelVpnStaticRoute struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VpnStaticRoute()
func NewServicesEc2ModelVpnStaticRoute() (*ServicesEc2ModelVpnStaticRoute) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VpnStaticRoute")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpnStaticRoute{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRoute) SetDestinationCidrBlock(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDestinationCidrBlock", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDestinationCidrBlock()
func (jbobject *ServicesEc2ModelVpnStaticRoute) GetDestinationCidrBlock() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDestinationCidrBlock", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnStaticRoute withDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRoute) WithDestinationCidrBlock(a string) *ServicesEc2ModelVpnStaticRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDestinationCidrBlock", "com/amazonaws/services/ec2/model/VpnStaticRoute", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnStaticRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSource(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRoute) SetSource2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSource", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSource()
func (jbobject *ServicesEc2ModelVpnStaticRoute) GetSource() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSource", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnStaticRoute withSource(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRoute) WithSource2(a string) *ServicesEc2ModelVpnStaticRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSource", "com/amazonaws/services/ec2/model/VpnStaticRoute", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnStaticRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSource(com.amazonaws.services.ec2.model.VpnStaticRouteSource)
func (jbobject *ServicesEc2ModelVpnStaticRoute) SetSource(a ServicesEc2ModelVpnStaticRouteSourceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSource", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnStaticRouteSource"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnStaticRoute withSource(com.amazonaws.services.ec2.model.VpnStaticRouteSource)
func (jbobject *ServicesEc2ModelVpnStaticRoute) WithSource(a ServicesEc2ModelVpnStaticRouteSourceInterface) *ServicesEc2ModelVpnStaticRoute {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSource", "com/amazonaws/services/ec2/model/VpnStaticRoute", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnStaticRouteSource"))
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
	unique_x := &ServicesEc2ModelVpnStaticRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRoute) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelVpnStaticRoute) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnStaticRoute withState(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRoute) WithState2(a string) *ServicesEc2ModelVpnStaticRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/VpnStaticRoute", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnStaticRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.VpnState)
func (jbobject *ServicesEc2ModelVpnStaticRoute) SetState(a ServicesEc2ModelVpnStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnStaticRoute withState(com.amazonaws.services.ec2.model.VpnState)
func (jbobject *ServicesEc2ModelVpnStaticRoute) WithState(a ServicesEc2ModelVpnStateInterface) *ServicesEc2ModelVpnStaticRoute {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/VpnStaticRoute", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnState"))
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
	unique_x := &ServicesEc2ModelVpnStaticRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpnStaticRoute) ToString() string {
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
func (jbobject *ServicesEc2ModelVpnStaticRoute) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpnStaticRoute) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VpnStaticRoute clone()
func (jbobject *ServicesEc2ModelVpnStaticRoute) Clone() *ServicesEc2ModelVpnStaticRoute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VpnStaticRoute")
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
	unique_x := &ServicesEc2ModelVpnStaticRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpnStaticRoute) Clone2() (*JavaLangObject, error) {
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


