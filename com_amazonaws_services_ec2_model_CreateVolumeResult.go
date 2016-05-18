package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVolumeResultInterface interface {
	JavaLangObjectInterface

	// public void setVolume(com.amazonaws.services.ec2.model.Volume)
	SetVolume(a ServicesEc2ModelVolumeInterface) 

	// public com.amazonaws.services.ec2.model.Volume getVolume()
	GetVolume() *ServicesEc2ModelVolume

	// public com.amazonaws.services.ec2.model.CreateVolumeResult withVolume(com.amazonaws.services.ec2.model.Volume)
	WithVolume(a ServicesEc2ModelVolumeInterface) *ServicesEc2ModelCreateVolumeResult

	// public com.amazonaws.services.ec2.model.CreateVolumeResult clone()
	Clone() *ServicesEc2ModelCreateVolumeResult
}

type ServicesEc2ModelCreateVolumeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVolumeResult()
func NewServicesEc2ModelCreateVolumeResult() (*ServicesEc2ModelCreateVolumeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVolumeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVolumeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVolume(com.amazonaws.services.ec2.model.Volume)
func (jbobject *ServicesEc2ModelCreateVolumeResult) SetVolume(a ServicesEc2ModelVolumeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolume", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Volume"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Volume getVolume()
func (jbobject *ServicesEc2ModelCreateVolumeResult) GetVolume() *ServicesEc2ModelVolume {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolume", "com/amazonaws/services/ec2/model/Volume")
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
	unique_x := &ServicesEc2ModelVolume{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVolumeResult withVolume(com.amazonaws.services.ec2.model.Volume)
func (jbobject *ServicesEc2ModelCreateVolumeResult) WithVolume(a ServicesEc2ModelVolumeInterface) *ServicesEc2ModelCreateVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolume", "com/amazonaws/services/ec2/model/CreateVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Volume"))
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
	unique_x := &ServicesEc2ModelCreateVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVolumeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVolumeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVolumeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVolumeResult clone()
func (jbobject *ServicesEc2ModelCreateVolumeResult) Clone() *ServicesEc2ModelCreateVolumeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVolumeResult")
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
	unique_x := &ServicesEc2ModelCreateVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVolumeResult) Clone2() (*JavaLangObject, error) {
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


