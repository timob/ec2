package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceBlockDeviceMappingInterface interface {
	JavaLangObjectInterface

	// public void setDeviceName(java.lang.String)
	SetDeviceName(a string) 

	// public java.lang.String getDeviceName()
	GetDeviceName() string

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping withDeviceName(java.lang.String)
	WithDeviceName(a string) *ServicesEc2ModelInstanceBlockDeviceMapping

	// public void setEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDevice)
	SetEbs(a ServicesEc2ModelEbsInstanceBlockDeviceInterface) 

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice getEbs()
	GetEbs() *ServicesEc2ModelEbsInstanceBlockDevice

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping withEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDevice)
	WithEbs(a ServicesEc2ModelEbsInstanceBlockDeviceInterface) *ServicesEc2ModelInstanceBlockDeviceMapping

	// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping clone()
	Clone() *ServicesEc2ModelInstanceBlockDeviceMapping
}

type ServicesEc2ModelInstanceBlockDeviceMapping struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping()
func NewServicesEc2ModelInstanceBlockDeviceMapping() (*ServicesEc2ModelInstanceBlockDeviceMapping) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceBlockDeviceMapping{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) SetDeviceName(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) GetDeviceName() string {
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

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping withDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) WithDeviceName(a string) *ServicesEc2ModelInstanceBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceName", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDevice)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) SetEbs(a ServicesEc2ModelEbsInstanceBlockDeviceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEbs", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EbsInstanceBlockDevice"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice getEbs()
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) GetEbs() *ServicesEc2ModelEbsInstanceBlockDevice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEbs", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice")
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping withEbs(com.amazonaws.services.ec2.model.EbsInstanceBlockDevice)
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) WithEbs(a ServicesEc2ModelEbsInstanceBlockDeviceInterface) *ServicesEc2ModelInstanceBlockDeviceMapping {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbs", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EbsInstanceBlockDevice"))
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceBlockDeviceMapping clone()
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) Clone() *ServicesEc2ModelInstanceBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceBlockDeviceMapping")
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
	unique_x := &ServicesEc2ModelInstanceBlockDeviceMapping{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceBlockDeviceMapping) Clone2() (*JavaLangObject, error) {
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


