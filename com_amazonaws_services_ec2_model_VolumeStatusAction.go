package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusActionInterface interface {
	JavaLangObjectInterface

	// public void setCode(java.lang.String)
	SetCode(a string) 

	// public java.lang.String getCode()
	GetCode() string

	// public com.amazonaws.services.ec2.model.VolumeStatusAction withCode(java.lang.String)
	WithCode(a string) *ServicesEc2ModelVolumeStatusAction

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.VolumeStatusAction withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelVolumeStatusAction

	// public void setEventType(java.lang.String)
	SetEventType(a string) 

	// public java.lang.String getEventType()
	GetEventType() string

	// public com.amazonaws.services.ec2.model.VolumeStatusAction withEventType(java.lang.String)
	WithEventType(a string) *ServicesEc2ModelVolumeStatusAction

	// public void setEventId(java.lang.String)
	SetEventId(a string) 

	// public java.lang.String getEventId()
	GetEventId() string

	// public com.amazonaws.services.ec2.model.VolumeStatusAction withEventId(java.lang.String)
	WithEventId(a string) *ServicesEc2ModelVolumeStatusAction

	// public com.amazonaws.services.ec2.model.VolumeStatusAction clone()
	Clone() *ServicesEc2ModelVolumeStatusAction
}

type ServicesEc2ModelVolumeStatusAction struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VolumeStatusAction()
func NewServicesEc2ModelVolumeStatusAction() (*ServicesEc2ModelVolumeStatusAction) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VolumeStatusAction")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVolumeStatusAction{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCode(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) SetCode(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCode()
func (jbobject *ServicesEc2ModelVolumeStatusAction) GetCode() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCode", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusAction withCode(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) WithCode(a string) *ServicesEc2ModelVolumeStatusAction {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/VolumeStatusAction", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusAction{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelVolumeStatusAction) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.VolumeStatusAction withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) WithDescription(a string) *ServicesEc2ModelVolumeStatusAction {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/VolumeStatusAction", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusAction{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEventType(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) SetEventType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEventType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getEventType()
func (jbobject *ServicesEc2ModelVolumeStatusAction) GetEventType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusAction withEventType(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) WithEventType(a string) *ServicesEc2ModelVolumeStatusAction {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventType", "com/amazonaws/services/ec2/model/VolumeStatusAction", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusAction{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEventId(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) SetEventId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEventId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getEventId()
func (jbobject *ServicesEc2ModelVolumeStatusAction) GetEventId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusAction withEventId(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusAction) WithEventId(a string) *ServicesEc2ModelVolumeStatusAction {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventId", "com/amazonaws/services/ec2/model/VolumeStatusAction", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusAction{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusAction) ToString() string {
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
func (jbobject *ServicesEc2ModelVolumeStatusAction) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVolumeStatusAction) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VolumeStatusAction clone()
func (jbobject *ServicesEc2ModelVolumeStatusAction) Clone() *ServicesEc2ModelVolumeStatusAction {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VolumeStatusAction")
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
	unique_x := &ServicesEc2ModelVolumeStatusAction{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVolumeStatusAction) Clone2() (*JavaLangObject, error) {
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


