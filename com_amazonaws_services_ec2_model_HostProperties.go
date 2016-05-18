package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelHostPropertiesInterface interface {
	JavaLangObjectInterface

	// public void setSockets(java.lang.Integer)
	SetSockets(a int) 

	// public java.lang.Integer getSockets()
	GetSockets() int

	// public com.amazonaws.services.ec2.model.HostProperties withSockets(java.lang.Integer)
	WithSockets(a int) *ServicesEc2ModelHostProperties

	// public void setCores(java.lang.Integer)
	SetCores(a int) 

	// public java.lang.Integer getCores()
	GetCores() int

	// public com.amazonaws.services.ec2.model.HostProperties withCores(java.lang.Integer)
	WithCores(a int) *ServicesEc2ModelHostProperties

	// public void setTotalVCpus(java.lang.Integer)
	SetTotalVCpus(a int) 

	// public java.lang.Integer getTotalVCpus()
	GetTotalVCpus() int

	// public com.amazonaws.services.ec2.model.HostProperties withTotalVCpus(java.lang.Integer)
	WithTotalVCpus(a int) *ServicesEc2ModelHostProperties

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.HostProperties withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelHostProperties

	// public com.amazonaws.services.ec2.model.HostProperties clone()
	Clone() *ServicesEc2ModelHostProperties
}

type ServicesEc2ModelHostProperties struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.HostProperties()
func NewServicesEc2ModelHostProperties() (*ServicesEc2ModelHostProperties) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/HostProperties")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelHostProperties{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSockets(java.lang.Integer)
func (jbobject *ServicesEc2ModelHostProperties) SetSockets(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSockets", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getSockets()
func (jbobject *ServicesEc2ModelHostProperties) GetSockets() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSockets", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.HostProperties withSockets(java.lang.Integer)
func (jbobject *ServicesEc2ModelHostProperties) WithSockets(a int) *ServicesEc2ModelHostProperties {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSockets", "com/amazonaws/services/ec2/model/HostProperties", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelHostProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCores(java.lang.Integer)
func (jbobject *ServicesEc2ModelHostProperties) SetCores(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCores", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getCores()
func (jbobject *ServicesEc2ModelHostProperties) GetCores() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCores", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.HostProperties withCores(java.lang.Integer)
func (jbobject *ServicesEc2ModelHostProperties) WithCores(a int) *ServicesEc2ModelHostProperties {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCores", "com/amazonaws/services/ec2/model/HostProperties", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelHostProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTotalVCpus(java.lang.Integer)
func (jbobject *ServicesEc2ModelHostProperties) SetTotalVCpus(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTotalVCpus", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getTotalVCpus()
func (jbobject *ServicesEc2ModelHostProperties) GetTotalVCpus() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTotalVCpus", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.HostProperties withTotalVCpus(java.lang.Integer)
func (jbobject *ServicesEc2ModelHostProperties) WithTotalVCpus(a int) *ServicesEc2ModelHostProperties {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTotalVCpus", "com/amazonaws/services/ec2/model/HostProperties", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelHostProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelHostProperties) SetInstanceType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelHostProperties) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.HostProperties withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelHostProperties) WithInstanceType(a string) *ServicesEc2ModelHostProperties {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/HostProperties", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHostProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelHostProperties) ToString() string {
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
func (jbobject *ServicesEc2ModelHostProperties) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelHostProperties) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.HostProperties clone()
func (jbobject *ServicesEc2ModelHostProperties) Clone() *ServicesEc2ModelHostProperties {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/HostProperties")
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
	unique_x := &ServicesEc2ModelHostProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelHostProperties) Clone2() (*JavaLangObject, error) {
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


