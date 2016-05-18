package ec2

import "github.com/timob/javabind"

type ProtocolInterface interface {

	// public static com.amazonaws.Protocol[] values()
	Values() []*Protocol

	// public static com.amazonaws.Protocol valueOf(java.lang.String)
	ValueOf(a string) *Protocol

	// public java.lang.String toString()
	ToString() string
}

type Protocol struct {
	*javabind.Callable
}

// public static com.amazonaws.Protocol[] values()
func (jbobject *Protocol) Values() []*Protocol {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/Protocol", "values", javabind.ObjectArrayType("com/amazonaws/Protocol"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/Protocol")
	dst := new([]*Protocol)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.Protocol valueOf(java.lang.String)
func (jbobject *Protocol) ValueOf(a string) *Protocol {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/Protocol", "valueOf", "com/amazonaws/Protocol", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &Protocol{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *Protocol) ToString() string {
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

func (jbobject *Protocol) HTTP() *Protocol {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/Protocol", "HTTP", "com/amazonaws/Protocol")
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
	unique_x := &Protocol{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Protocol) SetFieldHTTP(val ProtocolInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/Protocol", "HTTP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *Protocol) HTTPS() *Protocol {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/Protocol", "HTTPS", "com/amazonaws/Protocol")
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
	unique_x := &Protocol{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Protocol) SetFieldHTTPS(val ProtocolInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/Protocol", "HTTPS", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


