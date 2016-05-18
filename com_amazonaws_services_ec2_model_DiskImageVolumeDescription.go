package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDiskImageVolumeDescriptionInterface interface {
	JavaLangObjectInterface

	// public void setSize(java.lang.Long)
	SetSize(a int64) 

	// public java.lang.Long getSize()
	GetSize() int64

	// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription withSize(java.lang.Long)
	WithSize(a int64) *ServicesEc2ModelDiskImageVolumeDescription

	// public void setId(java.lang.String)
	SetId(a string) 

	// public java.lang.String getId()
	GetId() string

	// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription withId(java.lang.String)
	WithId(a string) *ServicesEc2ModelDiskImageVolumeDescription

	// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription clone()
	Clone() *ServicesEc2ModelDiskImageVolumeDescription
}

type ServicesEc2ModelDiskImageVolumeDescription struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription()
func NewServicesEc2ModelDiskImageVolumeDescription() (*ServicesEc2ModelDiskImageVolumeDescription) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DiskImageVolumeDescription")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDiskImageVolumeDescription{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSize(java.lang.Long)
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) SetSize(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getSize()
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) GetSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription withSize(java.lang.Long)
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) WithSize(a int64) *ServicesEc2ModelDiskImageVolumeDescription {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSize", "com/amazonaws/services/ec2/model/DiskImageVolumeDescription", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelDiskImageVolumeDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setId(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) SetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getId()
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) GetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription withId(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) WithId(a string) *ServicesEc2ModelDiskImageVolumeDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withId", "com/amazonaws/services/ec2/model/DiskImageVolumeDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageVolumeDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) ToString() string {
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
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription clone()
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) Clone() *ServicesEc2ModelDiskImageVolumeDescription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DiskImageVolumeDescription")
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
	unique_x := &ServicesEc2ModelDiskImageVolumeDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDiskImageVolumeDescription) Clone2() (*JavaLangObject, error) {
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


