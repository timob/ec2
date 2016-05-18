package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelOperationTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.OperationType[] values()
	Values() []*ServicesEc2ModelOperationType

	// public static com.amazonaws.services.ec2.model.OperationType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelOperationType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.OperationType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelOperationType
}

type ServicesEc2ModelOperationType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.OperationType[] values()
func (jbobject *ServicesEc2ModelOperationType) Values() []*ServicesEc2ModelOperationType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/OperationType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/OperationType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/OperationType")
	dst := new([]*ServicesEc2ModelOperationType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.OperationType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelOperationType) ValueOf(a string) *ServicesEc2ModelOperationType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/OperationType", "valueOf", "com/amazonaws/services/ec2/model/OperationType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelOperationType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelOperationType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.OperationType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelOperationType) FromValue(a string) *ServicesEc2ModelOperationType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/OperationType", "fromValue", "com/amazonaws/services/ec2/model/OperationType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelOperationType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOperationType) Add() *ServicesEc2ModelOperationType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OperationType", "Add", "com/amazonaws/services/ec2/model/OperationType")
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
	unique_x := &ServicesEc2ModelOperationType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOperationType) SetFieldAdd(val ServicesEc2ModelOperationTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OperationType", "Add", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelOperationType) Remove() *ServicesEc2ModelOperationType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OperationType", "Remove", "com/amazonaws/services/ec2/model/OperationType")
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
	unique_x := &ServicesEc2ModelOperationType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOperationType) SetFieldRemove(val ServicesEc2ModelOperationTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OperationType", "Remove", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


