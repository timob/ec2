package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDiskImageInterface interface {
	JavaLangObjectInterface

	// public void setImage(com.amazonaws.services.ec2.model.DiskImageDetail)
	SetImage(a ServicesEc2ModelDiskImageDetailInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageDetail getImage()
	GetImage() *ServicesEc2ModelDiskImageDetail

	// public com.amazonaws.services.ec2.model.DiskImage withImage(com.amazonaws.services.ec2.model.DiskImageDetail)
	WithImage(a ServicesEc2ModelDiskImageDetailInterface) *ServicesEc2ModelDiskImage

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.DiskImage withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelDiskImage

	// public void setVolume(com.amazonaws.services.ec2.model.VolumeDetail)
	SetVolume(a ServicesEc2ModelVolumeDetailInterface) 

	// public com.amazonaws.services.ec2.model.VolumeDetail getVolume()
	GetVolume() *ServicesEc2ModelVolumeDetail

	// public com.amazonaws.services.ec2.model.DiskImage withVolume(com.amazonaws.services.ec2.model.VolumeDetail)
	WithVolume(a ServicesEc2ModelVolumeDetailInterface) *ServicesEc2ModelDiskImage

	// public com.amazonaws.services.ec2.model.DiskImage clone()
	Clone() *ServicesEc2ModelDiskImage
}

type ServicesEc2ModelDiskImage struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DiskImage()
func NewServicesEc2ModelDiskImage() (*ServicesEc2ModelDiskImage) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DiskImage")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDiskImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImage(com.amazonaws.services.ec2.model.DiskImageDetail)
func (jbobject *ServicesEc2ModelDiskImage) SetImage(a ServicesEc2ModelDiskImageDetailInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDetail"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DiskImageDetail getImage()
func (jbobject *ServicesEc2ModelDiskImage) GetImage() *ServicesEc2ModelDiskImageDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImage", "com/amazonaws/services/ec2/model/DiskImageDetail")
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DiskImage withImage(com.amazonaws.services.ec2.model.DiskImageDetail)
func (jbobject *ServicesEc2ModelDiskImage) WithImage(a ServicesEc2ModelDiskImageDetailInterface) *ServicesEc2ModelDiskImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImage", "com/amazonaws/services/ec2/model/DiskImage", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDetail"))
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
	unique_x := &ServicesEc2ModelDiskImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImage) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelDiskImage) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DiskImage withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImage) WithDescription(a string) *ServicesEc2ModelDiskImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/DiskImage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolume(com.amazonaws.services.ec2.model.VolumeDetail)
func (jbobject *ServicesEc2ModelDiskImage) SetVolume(a ServicesEc2ModelVolumeDetailInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolume", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeDetail"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeDetail getVolume()
func (jbobject *ServicesEc2ModelDiskImage) GetVolume() *ServicesEc2ModelVolumeDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolume", "com/amazonaws/services/ec2/model/VolumeDetail")
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
	unique_x := &ServicesEc2ModelVolumeDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DiskImage withVolume(com.amazonaws.services.ec2.model.VolumeDetail)
func (jbobject *ServicesEc2ModelDiskImage) WithVolume(a ServicesEc2ModelVolumeDetailInterface) *ServicesEc2ModelDiskImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolume", "com/amazonaws/services/ec2/model/DiskImage", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeDetail"))
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
	unique_x := &ServicesEc2ModelDiskImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDiskImage) ToString() string {
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
func (jbobject *ServicesEc2ModelDiskImage) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDiskImage) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DiskImage clone()
func (jbobject *ServicesEc2ModelDiskImage) Clone() *ServicesEc2ModelDiskImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DiskImage")
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
	unique_x := &ServicesEc2ModelDiskImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDiskImage) Clone2() (*JavaLangObject, error) {
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


