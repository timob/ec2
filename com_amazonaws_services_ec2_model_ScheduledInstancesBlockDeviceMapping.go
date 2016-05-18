package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstancesBlockDeviceMappingInterface interface {
	JavaLangObjectInterface

	// public void setDeviceName(java.lang.String)
	SetDeviceName(a string) 

	// public java.lang.String getDeviceName()
	GetDeviceName() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withDeviceName(java.lang.String)
	WithDeviceName(a string) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping

	// public void setNoDevice(java.lang.String)
	SetNoDevice(a string) 

	// public java.lang.String getNoDevice()
	GetNoDevice() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withNoDevice(java.lang.String)
	WithNoDevice(a string) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping

	// public void setVirtualName(java.lang.String)
	SetVirtualName(a string) 

	// public java.lang.String getVirtualName()
	GetVirtualName() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withVirtualName(java.lang.String)
	WithVirtualName(a string) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping

	// public void setEbs(com.amazonaws.services.ec2.model.ScheduledInstancesEbs)
	SetEbs(a ServicesEc2ModelScheduledInstancesEbsInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs getEbs()
	GetEbs() *ServicesEc2ModelScheduledInstancesEbs

	// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withEbs(com.amazonaws.services.ec2.model.ScheduledInstancesEbs)
	WithEbs(a ServicesEc2ModelScheduledInstancesEbsInterface) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping

	// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping clone()
	Clone() *ServicesEc2ModelScheduledInstancesBlockDeviceMapping
}

type ServicesEc2ModelScheduledInstancesBlockDeviceMapping struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping()
func NewServicesEc2ModelScheduledInstancesBlockDeviceMapping() (*ServicesEc2ModelScheduledInstancesBlockDeviceMapping) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstancesBlockDeviceMapping{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) SetDeviceName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeviceName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDeviceName()
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) GetDeviceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeviceName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) WithDeviceName(a string) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceName", "com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNoDevice(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) SetNoDevice(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNoDevice", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNoDevice()
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) GetNoDevice() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNoDevice", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withNoDevice(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) WithNoDevice(a string) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNoDevice", "com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVirtualName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) SetVirtualName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVirtualName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVirtualName()
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) GetVirtualName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVirtualName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withVirtualName(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) WithVirtualName(a string) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualName", "com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbs(com.amazonaws.services.ec2.model.ScheduledInstancesEbs)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) SetEbs(a ServicesEc2ModelScheduledInstancesEbsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEbs", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesEbs"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstancesEbs getEbs()
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) GetEbs() *ServicesEc2ModelScheduledInstancesEbs {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEbs", "com/amazonaws/services/ec2/model/ScheduledInstancesEbs")
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
	unique_x := &ServicesEc2ModelScheduledInstancesEbs{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping withEbs(com.amazonaws.services.ec2.model.ScheduledInstancesEbs)
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) WithEbs(a ServicesEc2ModelScheduledInstancesEbsInterface) *ServicesEc2ModelScheduledInstancesBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbs", "com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstancesEbs"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesBlockDeviceMapping clone()
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) Clone() *ServicesEc2ModelScheduledInstancesBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstancesBlockDeviceMapping")
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
	unique_x := &ServicesEc2ModelScheduledInstancesBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstancesBlockDeviceMapping) Clone2() (*JavaLangObject, error) {
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


