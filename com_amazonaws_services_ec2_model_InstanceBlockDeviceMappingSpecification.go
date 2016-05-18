package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceBlockDeviceMappingSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setDeviceName(java.lang.String)
	SetDeviceName(a string) 

	// public java.lang.String getDeviceName()
	GetDeviceName() string

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withDeviceName(java.lang.String)
	WithDeviceName(a string) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification

	// public void setEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDeviceSpecification)
	SetEbs(a ServicesEc2ModelEbsInstanceBlockDeviceSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDeviceSpecification getEbs()
	GetEbs() *ServicesEc2ModelEbsInstanceBlockDeviceSpecification

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDeviceSpecification)
	WithEbs(a ServicesEc2ModelEbsInstanceBlockDeviceSpecificationInterface) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification

	// public void setVirtualName(java.lang.String)
	SetVirtualName(a string) 

	// public java.lang.String getVirtualName()
	GetVirtualName() string

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withVirtualName(java.lang.String)
	WithVirtualName(a string) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification

	// public void setNoDevice(java.lang.String)
	SetNoDevice(a string) 

	// public java.lang.String getNoDevice()
	GetNoDevice() string

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withNoDevice(java.lang.String)
	WithNoDevice(a string) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification clone()
	Clone() *ServicesEc2ModelInstanceBlockDeviceMappingSpecification
}

type ServicesEc2ModelInstanceBlockDeviceMappingSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification()
func NewServicesEc2ModelInstanceBlockDeviceMappingSpecification() (*ServicesEc2ModelInstanceBlockDeviceMappingSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceBlockDeviceMappingSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) SetDeviceName(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) GetDeviceName() string {
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

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) WithDeviceName(a string) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceName", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMappingSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDeviceSpecification)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) SetEbs(a ServicesEc2ModelEbsInstanceBlockDeviceSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEbs", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EbsInstanceBlockDeviceSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDeviceSpecification getEbs()
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) GetEbs() *ServicesEc2ModelEbsInstanceBlockDeviceSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEbs", "com/amazonaws/services/ec2/model/EbsInstanceBlockDeviceSpecification")
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDeviceSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDeviceSpecification)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) WithEbs(a ServicesEc2ModelEbsInstanceBlockDeviceSpecificationInterface) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbs", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EbsInstanceBlockDeviceSpecification"))
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMappingSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVirtualName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) SetVirtualName(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) GetVirtualName() string {
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

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withVirtualName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) WithVirtualName(a string) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualName", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMappingSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNoDevice(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) SetNoDevice(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) GetNoDevice() string {
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

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification withNoDevice(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) WithNoDevice(a string) *ServicesEc2ModelInstanceBlockDeviceMappingSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNoDevice", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMappingSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification clone()
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) Clone() *ServicesEc2ModelInstanceBlockDeviceMappingSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification")
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMappingSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMappingSpecification) Clone2() (*JavaLangObject, error) {
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


