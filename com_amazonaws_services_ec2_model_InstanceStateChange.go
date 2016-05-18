package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStateChangeInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.InstanceStateChange withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelInstanceStateChange

	// public void setCurrentState(com.amazonaws.services.ec2.model.InstanceState)
	SetCurrentState(a ServicesEc2ModelInstanceStateInterface) 

	// public com.amazonaws.services.ec2.model.InstanceState getCurrentState()
	GetCurrentState() *ServicesEc2ModelInstanceState

	// public com.amazonaws.services.ec2.model.InstanceStateChange withCurrentState(com.amazonaws.services.ec2.model.InstanceState)
	WithCurrentState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstanceStateChange

	// public void setPreviousState(com.amazonaws.services.ec2.model.InstanceState)
	SetPreviousState(a ServicesEc2ModelInstanceStateInterface) 

	// public com.amazonaws.services.ec2.model.InstanceState getPreviousState()
	GetPreviousState() *ServicesEc2ModelInstanceState

	// public com.amazonaws.services.ec2.model.InstanceStateChange withPreviousState(com.amazonaws.services.ec2.model.InstanceState)
	WithPreviousState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstanceStateChange

	// public com.amazonaws.services.ec2.model.InstanceStateChange clone()
	Clone() *ServicesEc2ModelInstanceStateChange
}

type ServicesEc2ModelInstanceStateChange struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceStateChange()
func NewServicesEc2ModelInstanceStateChange() (*ServicesEc2ModelInstanceStateChange) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceStateChange")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceStateChange{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStateChange) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelInstanceStateChange) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceStateChange withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStateChange) WithInstanceId(a string) *ServicesEc2ModelInstanceStateChange {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/InstanceStateChange", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStateChange{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrentState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstanceStateChange) SetCurrentState(a ServicesEc2ModelInstanceStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCurrentState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceState getCurrentState()
func (jbobject *ServicesEc2ModelInstanceStateChange) GetCurrentState() *ServicesEc2ModelInstanceState {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCurrentState", "com/amazonaws/services/ec2/model/InstanceState")
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStateChange withCurrentState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstanceStateChange) WithCurrentState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstanceStateChange {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrentState", "com/amazonaws/services/ec2/model/InstanceStateChange", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
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
	unique_x := &ServicesEc2ModelInstanceStateChange{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPreviousState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstanceStateChange) SetPreviousState(a ServicesEc2ModelInstanceStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreviousState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceState getPreviousState()
func (jbobject *ServicesEc2ModelInstanceStateChange) GetPreviousState() *ServicesEc2ModelInstanceState {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPreviousState", "com/amazonaws/services/ec2/model/InstanceState")
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStateChange withPreviousState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstanceStateChange) WithPreviousState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstanceStateChange {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPreviousState", "com/amazonaws/services/ec2/model/InstanceStateChange", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
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
	unique_x := &ServicesEc2ModelInstanceStateChange{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceStateChange) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceStateChange) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceStateChange) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceStateChange clone()
func (jbobject *ServicesEc2ModelInstanceStateChange) Clone() *ServicesEc2ModelInstanceStateChange {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceStateChange")
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
	unique_x := &ServicesEc2ModelInstanceStateChange{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceStateChange) Clone2() (*JavaLangObject, error) {
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


