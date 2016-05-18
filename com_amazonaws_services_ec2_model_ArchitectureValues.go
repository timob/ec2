package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelArchitectureValuesInterface interface {

	// public static com.amazonaws.services.ec2.model.ArchitectureValues[] values()
	Values() []*ServicesEc2ModelArchitectureValues

	// public static com.amazonaws.services.ec2.model.ArchitectureValues valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelArchitectureValues

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ArchitectureValues fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelArchitectureValues
}

type ServicesEc2ModelArchitectureValues struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ArchitectureValues[] values()
func (jbobject *ServicesEc2ModelArchitectureValues) Values() []*ServicesEc2ModelArchitectureValues {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ArchitectureValues", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ArchitectureValues"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ArchitectureValues")
	dst := new([]*ServicesEc2ModelArchitectureValues)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ArchitectureValues valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelArchitectureValues) ValueOf(a string) *ServicesEc2ModelArchitectureValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ArchitectureValues", "valueOf", "com/amazonaws/services/ec2/model/ArchitectureValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelArchitectureValues{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelArchitectureValues) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ArchitectureValues fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelArchitectureValues) FromValue(a string) *ServicesEc2ModelArchitectureValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ArchitectureValues", "fromValue", "com/amazonaws/services/ec2/model/ArchitectureValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelArchitectureValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelArchitectureValues) I386() *ServicesEc2ModelArchitectureValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ArchitectureValues", "I386", "com/amazonaws/services/ec2/model/ArchitectureValues")
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
	unique_x := &ServicesEc2ModelArchitectureValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelArchitectureValues) SetFieldI386(val ServicesEc2ModelArchitectureValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ArchitectureValues", "I386", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelArchitectureValues) X86_64() *ServicesEc2ModelArchitectureValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ArchitectureValues", "X86_64", "com/amazonaws/services/ec2/model/ArchitectureValues")
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
	unique_x := &ServicesEc2ModelArchitectureValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelArchitectureValues) SetFieldX86_64(val ServicesEc2ModelArchitectureValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ArchitectureValues", "X86_64", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


