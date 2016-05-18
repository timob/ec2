package ec2

import "github.com/timob/javabind"

type ResponseMetadataInterface interface {
	JavaLangObjectInterface

	// public java.lang.String getRequestId()
	GetRequestId() string
}

type ResponseMetadata struct {
	JavaLangObject
}

// public com.amazonaws.ResponseMetadata(java.util.Map<java.lang.String, java.lang.String>)
func NewResponseMetadata2(a map[string]string) (*ResponseMetadata) {
	conv_a := javabind.NewGoToJavaMap(javabind.NewGoToJavaString(), javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/ResponseMetadata", conv_a.Value().Cast("java/util/Map"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ResponseMetadata{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.ResponseMetadata(com.amazonaws.ResponseMetadata)
func NewResponseMetadata(a ResponseMetadataInterface) (*ResponseMetadata) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/ResponseMetadata", conv_a.Value().Cast("com/amazonaws/ResponseMetadata"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ResponseMetadata{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getRequestId()
func (jbobject *ResponseMetadata) GetRequestId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestId", "java/lang/String")
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

// public java.lang.String toString()
func (jbobject *ResponseMetadata) ToString() string {
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

func (jbobject *ResponseMetadata) AWS_REQUEST_ID() string {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/ResponseMetadata", "AWS_REQUEST_ID", "java/lang/String")
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

func (jbobject *ResponseMetadata) SetFieldAWS_REQUEST_ID(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/ResponseMetadata", "AWS_REQUEST_ID", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


