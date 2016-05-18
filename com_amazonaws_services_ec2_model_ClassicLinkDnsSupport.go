package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelClassicLinkDnsSupportInterface interface {
	JavaLangObjectInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelClassicLinkDnsSupport

	// public void setClassicLinkDnsSupported(java.lang.Boolean)
	SetClassicLinkDnsSupported(a bool) 

	// public java.lang.Boolean getClassicLinkDnsSupported()
	GetClassicLinkDnsSupported() bool

	// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport withClassicLinkDnsSupported(java.lang.Boolean)
	WithClassicLinkDnsSupported(a bool) *ServicesEc2ModelClassicLinkDnsSupport

	// public java.lang.Boolean isClassicLinkDnsSupported()
	IsClassicLinkDnsSupported() bool

	// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport clone()
	Clone() *ServicesEc2ModelClassicLinkDnsSupport
}

type ServicesEc2ModelClassicLinkDnsSupport struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport()
func NewServicesEc2ModelClassicLinkDnsSupport() (*ServicesEc2ModelClassicLinkDnsSupport) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ClassicLinkDnsSupport")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelClassicLinkDnsSupport{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) WithVpcId(a string) *ServicesEc2ModelClassicLinkDnsSupport {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/ClassicLinkDnsSupport", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelClassicLinkDnsSupport{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClassicLinkDnsSupported(java.lang.Boolean)
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) SetClassicLinkDnsSupported(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClassicLinkDnsSupported", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getClassicLinkDnsSupported()
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) GetClassicLinkDnsSupported() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClassicLinkDnsSupported", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport withClassicLinkDnsSupported(java.lang.Boolean)
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) WithClassicLinkDnsSupported(a bool) *ServicesEc2ModelClassicLinkDnsSupport {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClassicLinkDnsSupported", "com/amazonaws/services/ec2/model/ClassicLinkDnsSupport", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelClassicLinkDnsSupport{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isClassicLinkDnsSupported()
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) IsClassicLinkDnsSupported() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isClassicLinkDnsSupported", "java/lang/Boolean")
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
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) ToString() string {
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
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ClassicLinkDnsSupport clone()
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) Clone() *ServicesEc2ModelClassicLinkDnsSupport {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ClassicLinkDnsSupport")
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
	unique_x := &ServicesEc2ModelClassicLinkDnsSupport{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelClassicLinkDnsSupport) Clone2() (*JavaLangObject, error) {
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


