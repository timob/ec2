package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStateInterface interface {
	JavaLangObjectInterface

	// public void setCode(java.lang.Integer)
	SetCode(a int) 

	// public java.lang.Integer getCode()
	GetCode() int

	// public com.amazonaws.services.ec2.model.InstanceState withCode(java.lang.Integer)
	WithCode(a int) *ServicesEc2ModelInstanceState

	// public void setName(java.lang.String)
	SetName2(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.InstanceState withName(java.lang.String)
	WithName2(a string) *ServicesEc2ModelInstanceState

	// public void setName(com.amazonaws.services.ec2.model.InstanceStateName)
	SetName(a ServicesEc2ModelInstanceStateNameInterface) 

	// public com.amazonaws.services.ec2.model.InstanceState withName(com.amazonaws.services.ec2.model.InstanceStateName)
	WithName(a ServicesEc2ModelInstanceStateNameInterface) *ServicesEc2ModelInstanceState

	// public com.amazonaws.services.ec2.model.InstanceState clone()
	Clone() *ServicesEc2ModelInstanceState
}

type ServicesEc2ModelInstanceState struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceState()
func NewServicesEc2ModelInstanceState() (*ServicesEc2ModelInstanceState) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceState")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceState{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCode(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstanceState) SetCode(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getCode()
func (jbobject *ServicesEc2ModelInstanceState) GetCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCode", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.InstanceState withCode(java.lang.Integer)
func (jbobject *ServicesEc2ModelInstanceState) WithCode(a int) *ServicesEc2ModelInstanceState {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/InstanceState", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceState) SetName2(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceState) GetName() string {
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

// public com.amazonaws.services.ec2.model.InstanceState withName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceState) WithName2(a string) *ServicesEc2ModelInstanceState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/InstanceState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(com.amazonaws.services.ec2.model.InstanceStateName)
func (jbobject *ServicesEc2ModelInstanceState) SetName(a ServicesEc2ModelInstanceStateNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStateName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceState withName(com.amazonaws.services.ec2.model.InstanceStateName)
func (jbobject *ServicesEc2ModelInstanceState) WithName(a ServicesEc2ModelInstanceStateNameInterface) *ServicesEc2ModelInstanceState {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/InstanceState", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStateName"))
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceState) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceState) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceState) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceState clone()
func (jbobject *ServicesEc2ModelInstanceState) Clone() *ServicesEc2ModelInstanceState {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceState")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceState) Clone2() (*JavaLangObject, error) {
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


