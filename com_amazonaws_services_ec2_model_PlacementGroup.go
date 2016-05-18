package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPlacementGroupInterface interface {
	JavaLangObjectInterface

	// public void setGroupName(java.lang.String)
	SetGroupName(a string) 

	// public java.lang.String getGroupName()
	GetGroupName() string

	// public com.amazonaws.services.ec2.model.PlacementGroup withGroupName(java.lang.String)
	WithGroupName(a string) *ServicesEc2ModelPlacementGroup

	// public void setStrategy(java.lang.String)
	SetStrategy2(a string) 

	// public java.lang.String getStrategy()
	GetStrategy() string

	// public com.amazonaws.services.ec2.model.PlacementGroup withStrategy(java.lang.String)
	WithStrategy2(a string) *ServicesEc2ModelPlacementGroup

	// public void setStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
	SetStrategy(a ServicesEc2ModelPlacementStrategyInterface) 

	// public com.amazonaws.services.ec2.model.PlacementGroup withStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
	WithStrategy(a ServicesEc2ModelPlacementStrategyInterface) *ServicesEc2ModelPlacementGroup

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.PlacementGroup withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelPlacementGroup

	// public void setState(com.amazonaws.services.ec2.model.PlacementGroupState)
	SetState(a ServicesEc2ModelPlacementGroupStateInterface) 

	// public com.amazonaws.services.ec2.model.PlacementGroup withState(com.amazonaws.services.ec2.model.PlacementGroupState)
	WithState(a ServicesEc2ModelPlacementGroupStateInterface) *ServicesEc2ModelPlacementGroup

	// public com.amazonaws.services.ec2.model.PlacementGroup clone()
	Clone() *ServicesEc2ModelPlacementGroup
}

type ServicesEc2ModelPlacementGroup struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PlacementGroup()
func NewServicesEc2ModelPlacementGroup() (*ServicesEc2ModelPlacementGroup) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PlacementGroup")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPlacementGroup{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.PlacementGroup(java.lang.String)
func NewServicesEc2ModelPlacementGroup2(a string) (*ServicesEc2ModelPlacementGroup) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PlacementGroup", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelPlacementGroup{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroup) SetGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupName()
func (jbobject *ServicesEc2ModelPlacementGroup) GetGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PlacementGroup withGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroup) WithGroupName(a string) *ServicesEc2ModelPlacementGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupName", "com/amazonaws/services/ec2/model/PlacementGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStrategy(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroup) SetStrategy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStrategy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStrategy()
func (jbobject *ServicesEc2ModelPlacementGroup) GetStrategy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStrategy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PlacementGroup withStrategy(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroup) WithStrategy2(a string) *ServicesEc2ModelPlacementGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStrategy", "com/amazonaws/services/ec2/model/PlacementGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
func (jbobject *ServicesEc2ModelPlacementGroup) SetStrategy(a ServicesEc2ModelPlacementStrategyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStrategy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementStrategy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PlacementGroup withStrategy(com.amazonaws.services.ec2.model.PlacementStrategy)
func (jbobject *ServicesEc2ModelPlacementGroup) WithStrategy(a ServicesEc2ModelPlacementStrategyInterface) *ServicesEc2ModelPlacementGroup {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStrategy", "com/amazonaws/services/ec2/model/PlacementGroup", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementStrategy"))
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
	unique_x := &ServicesEc2ModelPlacementGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroup) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelPlacementGroup) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PlacementGroup withState(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroup) WithState2(a string) *ServicesEc2ModelPlacementGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/PlacementGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.PlacementGroupState)
func (jbobject *ServicesEc2ModelPlacementGroup) SetState(a ServicesEc2ModelPlacementGroupStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementGroupState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PlacementGroup withState(com.amazonaws.services.ec2.model.PlacementGroupState)
func (jbobject *ServicesEc2ModelPlacementGroup) WithState(a ServicesEc2ModelPlacementGroupStateInterface) *ServicesEc2ModelPlacementGroup {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/PlacementGroup", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlacementGroupState"))
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
	unique_x := &ServicesEc2ModelPlacementGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPlacementGroup) ToString() string {
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
func (jbobject *ServicesEc2ModelPlacementGroup) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPlacementGroup) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PlacementGroup clone()
func (jbobject *ServicesEc2ModelPlacementGroup) Clone() *ServicesEc2ModelPlacementGroup {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PlacementGroup")
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
	unique_x := &ServicesEc2ModelPlacementGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPlacementGroup) Clone2() (*JavaLangObject, error) {
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


