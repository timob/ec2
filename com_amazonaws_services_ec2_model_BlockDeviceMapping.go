package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelBlockDeviceMappingInterface interface {
	JavaLangObjectInterface

	// public void setVirtualName(java.lang.String)
	SetVirtualName(a string) 

	// public java.lang.String getVirtualName()
	GetVirtualName() string

	// public com.amazonaws.services.ec2.model.BlockDeviceMapping withVirtualName(java.lang.String)
	WithVirtualName(a string) *ServicesEc2ModelBlockDeviceMapping

	// public void setDeviceName(java.lang.String)
	SetDeviceName(a string) 

	// public java.lang.String getDeviceName()
	GetDeviceName() string

	// public com.amazonaws.services.ec2.model.BlockDeviceMapping withDeviceName(java.lang.String)
	WithDeviceName(a string) *ServicesEc2ModelBlockDeviceMapping

	// public void setEbs(com.amazonaws.services.ec2.model.EbsBlockDevice)
	SetEbs(a ServicesEc2ModelEbsBlockDeviceInterface) 

	// public com.amazonaws.services.ec2.model.EbsBlockDevice getEbs()
	GetEbs() *ServicesEc2ModelEbsBlockDevice

	// public com.amazonaws.services.ec2.model.BlockDeviceMapping withEbs(com.amazonaws.services.ec2.model.EbsBlockDevice)
	WithEbs(a ServicesEc2ModelEbsBlockDeviceInterface) *ServicesEc2ModelBlockDeviceMapping

	// public void setNoDevice(java.lang.String)
	SetNoDevice(a string) 

	// public java.lang.String getNoDevice()
	GetNoDevice() string

	// public com.amazonaws.services.ec2.model.BlockDeviceMapping withNoDevice(java.lang.String)
	WithNoDevice(a string) *ServicesEc2ModelBlockDeviceMapping

	// public com.amazonaws.services.ec2.model.BlockDeviceMapping clone()
	Clone() *ServicesEc2ModelBlockDeviceMapping
}

type ServicesEc2ModelBlockDeviceMapping struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.BlockDeviceMapping()
func NewServicesEc2ModelBlockDeviceMapping() (*ServicesEc2ModelBlockDeviceMapping) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/BlockDeviceMapping")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelBlockDeviceMapping{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVirtualName(java.lang.String)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) SetVirtualName(a string)  {
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
func (jbobject *ServicesEc2ModelBlockDeviceMapping) GetVirtualName() string {
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

// public com.amazonaws.services.ec2.model.BlockDeviceMapping withVirtualName(java.lang.String)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) WithVirtualName(a string) *ServicesEc2ModelBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualName", "com/amazonaws/services/ec2/model/BlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) SetDeviceName(a string)  {
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
func (jbobject *ServicesEc2ModelBlockDeviceMapping) GetDeviceName() string {
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

// public com.amazonaws.services.ec2.model.BlockDeviceMapping withDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) WithDeviceName(a string) *ServicesEc2ModelBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceName", "com/amazonaws/services/ec2/model/BlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbs(com.amazonaws.services.ec2.model.EbsBlockDevice)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) SetEbs(a ServicesEc2ModelEbsBlockDeviceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEbs", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EbsBlockDevice"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EbsBlockDevice getEbs()
func (jbobject *ServicesEc2ModelBlockDeviceMapping) GetEbs() *ServicesEc2ModelEbsBlockDevice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEbs", "com/amazonaws/services/ec2/model/EbsBlockDevice")
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.BlockDeviceMapping withEbs(com.amazonaws.services.ec2.model.EbsBlockDevice)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) WithEbs(a ServicesEc2ModelEbsBlockDeviceInterface) *ServicesEc2ModelBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbs", "com/amazonaws/services/ec2/model/BlockDeviceMapping", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EbsBlockDevice"))
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
	unique_x := &ServicesEc2ModelBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNoDevice(java.lang.String)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) SetNoDevice(a string)  {
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
func (jbobject *ServicesEc2ModelBlockDeviceMapping) GetNoDevice() string {
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

// public com.amazonaws.services.ec2.model.BlockDeviceMapping withNoDevice(java.lang.String)
func (jbobject *ServicesEc2ModelBlockDeviceMapping) WithNoDevice(a string) *ServicesEc2ModelBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNoDevice", "com/amazonaws/services/ec2/model/BlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelBlockDeviceMapping) ToString() string {
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
func (jbobject *ServicesEc2ModelBlockDeviceMapping) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelBlockDeviceMapping) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.BlockDeviceMapping clone()
func (jbobject *ServicesEc2ModelBlockDeviceMapping) Clone() *ServicesEc2ModelBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/BlockDeviceMapping")
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
	unique_x := &ServicesEc2ModelBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelBlockDeviceMapping) Clone2() (*JavaLangObject, error) {
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


