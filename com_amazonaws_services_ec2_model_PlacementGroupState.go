package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPlacementGroupStateInterface interface {

	// public static com.amazonaws.services.ec2.model.PlacementGroupState[] values()
	Values() []*ServicesEc2ModelPlacementGroupState

	// public static com.amazonaws.services.ec2.model.PlacementGroupState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelPlacementGroupState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.PlacementGroupState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelPlacementGroupState
}

type ServicesEc2ModelPlacementGroupState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.PlacementGroupState[] values()
func (jbobject *ServicesEc2ModelPlacementGroupState) Values() []*ServicesEc2ModelPlacementGroupState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlacementGroupState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/PlacementGroupState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/PlacementGroupState")
	dst := new([]*ServicesEc2ModelPlacementGroupState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.PlacementGroupState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroupState) ValueOf(a string) *ServicesEc2ModelPlacementGroupState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlacementGroupState", "valueOf", "com/amazonaws/services/ec2/model/PlacementGroupState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementGroupState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPlacementGroupState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.PlacementGroupState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementGroupState) FromValue(a string) *ServicesEc2ModelPlacementGroupState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlacementGroupState", "fromValue", "com/amazonaws/services/ec2/model/PlacementGroupState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementGroupState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementGroupState) Pending() *ServicesEc2ModelPlacementGroupState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Pending", "com/amazonaws/services/ec2/model/PlacementGroupState")
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
	unique_x := &ServicesEc2ModelPlacementGroupState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementGroupState) SetFieldPending(val ServicesEc2ModelPlacementGroupStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelPlacementGroupState) Available() *ServicesEc2ModelPlacementGroupState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Available", "com/amazonaws/services/ec2/model/PlacementGroupState")
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
	unique_x := &ServicesEc2ModelPlacementGroupState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementGroupState) SetFieldAvailable(val ServicesEc2ModelPlacementGroupStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelPlacementGroupState) Deleting() *ServicesEc2ModelPlacementGroupState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Deleting", "com/amazonaws/services/ec2/model/PlacementGroupState")
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
	unique_x := &ServicesEc2ModelPlacementGroupState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementGroupState) SetFieldDeleting(val ServicesEc2ModelPlacementGroupStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Deleting", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelPlacementGroupState) Deleted() *ServicesEc2ModelPlacementGroupState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Deleted", "com/amazonaws/services/ec2/model/PlacementGroupState")
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
	unique_x := &ServicesEc2ModelPlacementGroupState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementGroupState) SetFieldDeleted(val ServicesEc2ModelPlacementGroupStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PlacementGroupState", "Deleted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


