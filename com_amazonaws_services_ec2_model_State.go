package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStateInterface interface {

	// public static com.amazonaws.services.ec2.model.State[] values()
	Values() []*ServicesEc2ModelState

	// public static com.amazonaws.services.ec2.model.State valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.State fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelState
}

type ServicesEc2ModelState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.State[] values()
func (jbobject *ServicesEc2ModelState) Values() []*ServicesEc2ModelState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/State", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/State"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/State")
	dst := new([]*ServicesEc2ModelState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.State valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelState) ValueOf(a string) *ServicesEc2ModelState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/State", "valueOf", "com/amazonaws/services/ec2/model/State", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.State fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelState) FromValue(a string) *ServicesEc2ModelState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/State", "fromValue", "com/amazonaws/services/ec2/model/State", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelState) Pending() *ServicesEc2ModelState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/State", "Pending", "com/amazonaws/services/ec2/model/State")
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
	unique_x := &ServicesEc2ModelState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelState) SetFieldPending(val ServicesEc2ModelStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/State", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelState) Available() *ServicesEc2ModelState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/State", "Available", "com/amazonaws/services/ec2/model/State")
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
	unique_x := &ServicesEc2ModelState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelState) SetFieldAvailable(val ServicesEc2ModelStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/State", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelState) Deleting() *ServicesEc2ModelState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/State", "Deleting", "com/amazonaws/services/ec2/model/State")
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
	unique_x := &ServicesEc2ModelState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelState) SetFieldDeleting(val ServicesEc2ModelStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/State", "Deleting", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelState) Deleted() *ServicesEc2ModelState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/State", "Deleted", "com/amazonaws/services/ec2/model/State")
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
	unique_x := &ServicesEc2ModelState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelState) SetFieldDeleted(val ServicesEc2ModelStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/State", "Deleted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


