package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAttachmentStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.AttachmentStatus[] values()
	Values() []*ServicesEc2ModelAttachmentStatus

	// public static com.amazonaws.services.ec2.model.AttachmentStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAttachmentStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.AttachmentStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAttachmentStatus
}

type ServicesEc2ModelAttachmentStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.AttachmentStatus[] values()
func (jbobject *ServicesEc2ModelAttachmentStatus) Values() []*ServicesEc2ModelAttachmentStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AttachmentStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/AttachmentStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/AttachmentStatus")
	dst := new([]*ServicesEc2ModelAttachmentStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.AttachmentStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAttachmentStatus) ValueOf(a string) *ServicesEc2ModelAttachmentStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AttachmentStatus", "valueOf", "com/amazonaws/services/ec2/model/AttachmentStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachmentStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAttachmentStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.AttachmentStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAttachmentStatus) FromValue(a string) *ServicesEc2ModelAttachmentStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AttachmentStatus", "fromValue", "com/amazonaws/services/ec2/model/AttachmentStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachmentStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAttachmentStatus) Attaching() *ServicesEc2ModelAttachmentStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Attaching", "com/amazonaws/services/ec2/model/AttachmentStatus")
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
	unique_x := &ServicesEc2ModelAttachmentStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAttachmentStatus) SetFieldAttaching(val ServicesEc2ModelAttachmentStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Attaching", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAttachmentStatus) Attached() *ServicesEc2ModelAttachmentStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Attached", "com/amazonaws/services/ec2/model/AttachmentStatus")
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
	unique_x := &ServicesEc2ModelAttachmentStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAttachmentStatus) SetFieldAttached(val ServicesEc2ModelAttachmentStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Attached", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAttachmentStatus) Detaching() *ServicesEc2ModelAttachmentStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Detaching", "com/amazonaws/services/ec2/model/AttachmentStatus")
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
	unique_x := &ServicesEc2ModelAttachmentStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAttachmentStatus) SetFieldDetaching(val ServicesEc2ModelAttachmentStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Detaching", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAttachmentStatus) Detached() *ServicesEc2ModelAttachmentStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Detached", "com/amazonaws/services/ec2/model/AttachmentStatus")
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
	unique_x := &ServicesEc2ModelAttachmentStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAttachmentStatus) SetFieldDetached(val ServicesEc2ModelAttachmentStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AttachmentStatus", "Detached", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


