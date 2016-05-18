package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelMoveStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.MoveStatus[] values()
	Values() []*ServicesEc2ModelMoveStatus

	// public static com.amazonaws.services.ec2.model.MoveStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelMoveStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.MoveStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelMoveStatus
}

type ServicesEc2ModelMoveStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.MoveStatus[] values()
func (jbobject *ServicesEc2ModelMoveStatus) Values() []*ServicesEc2ModelMoveStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/MoveStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/MoveStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/MoveStatus")
	dst := new([]*ServicesEc2ModelMoveStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.MoveStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelMoveStatus) ValueOf(a string) *ServicesEc2ModelMoveStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/MoveStatus", "valueOf", "com/amazonaws/services/ec2/model/MoveStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelMoveStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelMoveStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.MoveStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelMoveStatus) FromValue(a string) *ServicesEc2ModelMoveStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/MoveStatus", "fromValue", "com/amazonaws/services/ec2/model/MoveStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelMoveStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMoveStatus) MovingToVpc() *ServicesEc2ModelMoveStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/MoveStatus", "MovingToVpc", "com/amazonaws/services/ec2/model/MoveStatus")
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
	unique_x := &ServicesEc2ModelMoveStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMoveStatus) SetFieldMovingToVpc(val ServicesEc2ModelMoveStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/MoveStatus", "MovingToVpc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelMoveStatus) RestoringToClassic() *ServicesEc2ModelMoveStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/MoveStatus", "RestoringToClassic", "com/amazonaws/services/ec2/model/MoveStatus")
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
	unique_x := &ServicesEc2ModelMoveStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMoveStatus) SetFieldRestoringToClassic(val ServicesEc2ModelMoveStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/MoveStatus", "RestoringToClassic", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


