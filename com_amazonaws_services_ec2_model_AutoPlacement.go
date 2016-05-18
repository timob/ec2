package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAutoPlacementInterface interface {

	// public static com.amazonaws.services.ec2.model.AutoPlacement[] values()
	Values() []*ServicesEc2ModelAutoPlacement

	// public static com.amazonaws.services.ec2.model.AutoPlacement valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAutoPlacement

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.AutoPlacement fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAutoPlacement
}

type ServicesEc2ModelAutoPlacement struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.AutoPlacement[] values()
func (jbobject *ServicesEc2ModelAutoPlacement) Values() []*ServicesEc2ModelAutoPlacement {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AutoPlacement", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/AutoPlacement"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/AutoPlacement")
	dst := new([]*ServicesEc2ModelAutoPlacement)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.AutoPlacement valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAutoPlacement) ValueOf(a string) *ServicesEc2ModelAutoPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AutoPlacement", "valueOf", "com/amazonaws/services/ec2/model/AutoPlacement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAutoPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAutoPlacement) ToString() string {
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

// public static com.amazonaws.services.ec2.model.AutoPlacement fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAutoPlacement) FromValue(a string) *ServicesEc2ModelAutoPlacement {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AutoPlacement", "fromValue", "com/amazonaws/services/ec2/model/AutoPlacement", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAutoPlacement{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAutoPlacement) On() *ServicesEc2ModelAutoPlacement {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AutoPlacement", "On", "com/amazonaws/services/ec2/model/AutoPlacement")
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
	unique_x := &ServicesEc2ModelAutoPlacement{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAutoPlacement) SetFieldOn(val ServicesEc2ModelAutoPlacementInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AutoPlacement", "On", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAutoPlacement) Off() *ServicesEc2ModelAutoPlacement {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AutoPlacement", "Off", "com/amazonaws/services/ec2/model/AutoPlacement")
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
	unique_x := &ServicesEc2ModelAutoPlacement{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAutoPlacement) SetFieldOff(val ServicesEc2ModelAutoPlacementInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AutoPlacement", "Off", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


