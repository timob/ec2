package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDomainTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.DomainType[] values()
	Values() []*ServicesEc2ModelDomainType

	// public static com.amazonaws.services.ec2.model.DomainType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelDomainType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.DomainType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelDomainType
}

type ServicesEc2ModelDomainType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.DomainType[] values()
func (jbobject *ServicesEc2ModelDomainType) Values() []*ServicesEc2ModelDomainType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DomainType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/DomainType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/DomainType")
	dst := new([]*ServicesEc2ModelDomainType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.DomainType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelDomainType) ValueOf(a string) *ServicesEc2ModelDomainType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DomainType", "valueOf", "com/amazonaws/services/ec2/model/DomainType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDomainType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDomainType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.DomainType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelDomainType) FromValue(a string) *ServicesEc2ModelDomainType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DomainType", "fromValue", "com/amazonaws/services/ec2/model/DomainType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDomainType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDomainType) Vpc() *ServicesEc2ModelDomainType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DomainType", "Vpc", "com/amazonaws/services/ec2/model/DomainType")
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
	unique_x := &ServicesEc2ModelDomainType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDomainType) SetFieldVpc(val ServicesEc2ModelDomainTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DomainType", "Vpc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelDomainType) Standard() *ServicesEc2ModelDomainType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DomainType", "Standard", "com/amazonaws/services/ec2/model/DomainType")
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
	unique_x := &ServicesEc2ModelDomainType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDomainType) SetFieldStandard(val ServicesEc2ModelDomainTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DomainType", "Standard", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


