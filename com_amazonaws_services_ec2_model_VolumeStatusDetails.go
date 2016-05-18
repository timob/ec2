package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusDetailsInterface interface {
	JavaLangObjectInterface

	// public void setName(java.lang.String)
	SetName2(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.VolumeStatusDetails withName(java.lang.String)
	WithName2(a string) *ServicesEc2ModelVolumeStatusDetails

	// public void setName(com.amazonaws.services.ec2.model.VolumeStatusName)
	SetName(a ServicesEc2ModelVolumeStatusNameInterface) 

	// public com.amazonaws.services.ec2.model.VolumeStatusDetails withName(com.amazonaws.services.ec2.model.VolumeStatusName)
	WithName(a ServicesEc2ModelVolumeStatusNameInterface) *ServicesEc2ModelVolumeStatusDetails

	// public void setStatus(java.lang.String)
	SetStatus(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.VolumeStatusDetails withStatus(java.lang.String)
	WithStatus(a string) *ServicesEc2ModelVolumeStatusDetails

	// public com.amazonaws.services.ec2.model.VolumeStatusDetails clone()
	Clone() *ServicesEc2ModelVolumeStatusDetails
}

type ServicesEc2ModelVolumeStatusDetails struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VolumeStatusDetails()
func NewServicesEc2ModelVolumeStatusDetails() (*ServicesEc2ModelVolumeStatusDetails) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VolumeStatusDetails")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVolumeStatusDetails{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusDetails) SetName2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getName()
func (jbobject *ServicesEc2ModelVolumeStatusDetails) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusDetails withName(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusDetails) WithName2(a string) *ServicesEc2ModelVolumeStatusDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/VolumeStatusDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(com.amazonaws.services.ec2.model.VolumeStatusName)
func (jbobject *ServicesEc2ModelVolumeStatusDetails) SetName(a ServicesEc2ModelVolumeStatusNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeStatusDetails withName(com.amazonaws.services.ec2.model.VolumeStatusName)
func (jbobject *ServicesEc2ModelVolumeStatusDetails) WithName(a ServicesEc2ModelVolumeStatusNameInterface) *ServicesEc2ModelVolumeStatusDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/VolumeStatusDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusName"))
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
	unique_x := &ServicesEc2ModelVolumeStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusDetails) SetStatus(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelVolumeStatusDetails) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusDetails withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusDetails) WithStatus(a string) *ServicesEc2ModelVolumeStatusDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/VolumeStatusDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusDetails) ToString() string {
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
func (jbobject *ServicesEc2ModelVolumeStatusDetails) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVolumeStatusDetails) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VolumeStatusDetails clone()
func (jbobject *ServicesEc2ModelVolumeStatusDetails) Clone() *ServicesEc2ModelVolumeStatusDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VolumeStatusDetails")
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
	unique_x := &ServicesEc2ModelVolumeStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVolumeStatusDetails) Clone2() (*JavaLangObject, error) {
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


