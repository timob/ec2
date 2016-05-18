package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRuleActionInterface interface {

	// public static com.amazonaws.services.ec2.model.RuleAction[] values()
	Values() []*ServicesEc2ModelRuleAction

	// public static com.amazonaws.services.ec2.model.RuleAction valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelRuleAction

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.RuleAction fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelRuleAction
}

type ServicesEc2ModelRuleAction struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.RuleAction[] values()
func (jbobject *ServicesEc2ModelRuleAction) Values() []*ServicesEc2ModelRuleAction {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RuleAction", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/RuleAction"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/RuleAction")
	dst := new([]*ServicesEc2ModelRuleAction)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.RuleAction valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelRuleAction) ValueOf(a string) *ServicesEc2ModelRuleAction {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RuleAction", "valueOf", "com/amazonaws/services/ec2/model/RuleAction", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRuleAction{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRuleAction) ToString() string {
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

// public static com.amazonaws.services.ec2.model.RuleAction fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelRuleAction) FromValue(a string) *ServicesEc2ModelRuleAction {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RuleAction", "fromValue", "com/amazonaws/services/ec2/model/RuleAction", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRuleAction{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRuleAction) Allow() *ServicesEc2ModelRuleAction {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RuleAction", "Allow", "com/amazonaws/services/ec2/model/RuleAction")
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
	unique_x := &ServicesEc2ModelRuleAction{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRuleAction) SetFieldAllow(val ServicesEc2ModelRuleActionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RuleAction", "Allow", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRuleAction) Deny() *ServicesEc2ModelRuleAction {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RuleAction", "Deny", "com/amazonaws/services/ec2/model/RuleAction")
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
	unique_x := &ServicesEc2ModelRuleAction{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRuleAction) SetFieldDeny(val ServicesEc2ModelRuleActionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RuleAction", "Deny", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


