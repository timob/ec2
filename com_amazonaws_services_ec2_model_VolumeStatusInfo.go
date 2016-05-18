package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusInfoInterface interface {
	JavaLangObjectInterface

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.VolumeStatusInfo withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelVolumeStatusInfo

	// public void setStatus(com.amazonaws.services.ec2.model.VolumeStatusInfoStatus)
	SetStatus(a ServicesEc2ModelVolumeStatusInfoStatusInterface) 

	// public com.amazonaws.services.ec2.model.VolumeStatusInfo withStatus(com.amazonaws.services.ec2.model.VolumeStatusInfoStatus)
	WithStatus(a ServicesEc2ModelVolumeStatusInfoStatusInterface) *ServicesEc2ModelVolumeStatusInfo

	// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusDetails> getDetails()
	GetDetails() []*ServicesEc2ModelVolumeStatusDetails

	// public void setDetails(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusDetails>)
	SetDetails(a []*ServicesEc2ModelVolumeStatusDetails) 

	// public com.amazonaws.services.ec2.model.VolumeStatusInfo withDetails(com.amazonaws.services.ec2.model.VolumeStatusDetails...)
	WithDetails(a ...*ServicesEc2ModelVolumeStatusDetails) *ServicesEc2ModelVolumeStatusInfo

	// public com.amazonaws.services.ec2.model.VolumeStatusInfo withDetails(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusDetails>)
	WithDetails2(a []*ServicesEc2ModelVolumeStatusDetails) *ServicesEc2ModelVolumeStatusInfo

	// public com.amazonaws.services.ec2.model.VolumeStatusInfo clone()
	Clone() *ServicesEc2ModelVolumeStatusInfo
}

type ServicesEc2ModelVolumeStatusInfo struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VolumeStatusInfo()
func NewServicesEc2ModelVolumeStatusInfo() (*ServicesEc2ModelVolumeStatusInfo) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VolumeStatusInfo")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVolumeStatusInfo{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) SetStatus2(a string)  {
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
func (jbobject *ServicesEc2ModelVolumeStatusInfo) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.VolumeStatusInfo withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) WithStatus2(a string) *ServicesEc2ModelVolumeStatusInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/VolumeStatusInfo", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.VolumeStatusInfoStatus)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) SetStatus(a ServicesEc2ModelVolumeStatusInfoStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeStatusInfo withStatus(com.amazonaws.services.ec2.model.VolumeStatusInfoStatus)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) WithStatus(a ServicesEc2ModelVolumeStatusInfoStatusInterface) *ServicesEc2ModelVolumeStatusInfo {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/VolumeStatusInfo", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus"))
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
	unique_x := &ServicesEc2ModelVolumeStatusInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusDetails> getDetails()
func (jbobject *ServicesEc2ModelVolumeStatusInfo) GetDetails() []*ServicesEc2ModelVolumeStatusDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDetails", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVolumeStatusDetails)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDetails(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusDetails>)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) SetDetails(a []*ServicesEc2ModelVolumeStatusDetails)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDetails", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeStatusInfo withDetails(com.amazonaws.services.ec2.model.VolumeStatusDetails...)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) WithDetails(a ...*ServicesEc2ModelVolumeStatusDetails) *ServicesEc2ModelVolumeStatusInfo {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VolumeStatusDetails")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDetails", "com/amazonaws/services/ec2/model/VolumeStatusInfo", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusDetails"))
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
	unique_x := &ServicesEc2ModelVolumeStatusInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VolumeStatusInfo withDetails(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusDetails>)
func (jbobject *ServicesEc2ModelVolumeStatusInfo) WithDetails2(a []*ServicesEc2ModelVolumeStatusDetails) *ServicesEc2ModelVolumeStatusInfo {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDetails", "com/amazonaws/services/ec2/model/VolumeStatusInfo", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVolumeStatusInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusInfo) ToString() string {
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
func (jbobject *ServicesEc2ModelVolumeStatusInfo) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVolumeStatusInfo) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VolumeStatusInfo clone()
func (jbobject *ServicesEc2ModelVolumeStatusInfo) Clone() *ServicesEc2ModelVolumeStatusInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VolumeStatusInfo")
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
	unique_x := &ServicesEc2ModelVolumeStatusInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVolumeStatusInfo) Clone2() (*JavaLangObject, error) {
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


