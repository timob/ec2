package ec2

import "github.com/timob/javabind"

type EventProgressEventTypeInterface interface {

	// public static com.amazonaws.event.ProgressEventType[] values()
	Values() []*EventProgressEventType

	// public static com.amazonaws.event.ProgressEventType valueOf(java.lang.String)
	ValueOf(a string) *EventProgressEventType

	// public boolean isTransferEvent()
	IsTransferEvent() bool

	// public boolean isRequestCycleEvent()
	IsRequestCycleEvent() bool

	// public boolean isByteCountEvent()
	IsByteCountEvent() bool
}

type EventProgressEventType struct {
	*javabind.Callable
}

// public static com.amazonaws.event.ProgressEventType[] values()
func (jbobject *EventProgressEventType) Values() []*EventProgressEventType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/event/ProgressEventType", "values", javabind.ObjectArrayType("com/amazonaws/event/ProgressEventType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/event/ProgressEventType")
	dst := new([]*EventProgressEventType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.event.ProgressEventType valueOf(java.lang.String)
func (jbobject *EventProgressEventType) ValueOf(a string) *EventProgressEventType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/event/ProgressEventType", "valueOf", "com/amazonaws/event/ProgressEventType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isTransferEvent()
func (jbobject *EventProgressEventType) IsTransferEvent() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTransferEvent", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isRequestCycleEvent()
func (jbobject *EventProgressEventType) IsRequestCycleEvent() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isRequestCycleEvent", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isByteCountEvent()
func (jbobject *EventProgressEventType) IsByteCountEvent() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isByteCountEvent", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *EventProgressEventType) BYTE_TRANSFER_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "BYTE_TRANSFER_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldBYTE_TRANSFER_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "BYTE_TRANSFER_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) REQUEST_CONTENT_LENGTH_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "REQUEST_CONTENT_LENGTH_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldREQUEST_CONTENT_LENGTH_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "REQUEST_CONTENT_LENGTH_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) RESPONSE_CONTENT_LENGTH_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "RESPONSE_CONTENT_LENGTH_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldRESPONSE_CONTENT_LENGTH_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "RESPONSE_CONTENT_LENGTH_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) REQUEST_BYTE_TRANSFER_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "REQUEST_BYTE_TRANSFER_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldREQUEST_BYTE_TRANSFER_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "REQUEST_BYTE_TRANSFER_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) RESPONSE_BYTE_TRANSFER_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "RESPONSE_BYTE_TRANSFER_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldRESPONSE_BYTE_TRANSFER_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "RESPONSE_BYTE_TRANSFER_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) RESPONSE_BYTE_DISCARD_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "RESPONSE_BYTE_DISCARD_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldRESPONSE_BYTE_DISCARD_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "RESPONSE_BYTE_DISCARD_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) CLIENT_REQUEST_STARTED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_STARTED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldCLIENT_REQUEST_STARTED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_STARTED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) HTTP_REQUEST_STARTED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_REQUEST_STARTED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldHTTP_REQUEST_STARTED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_REQUEST_STARTED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) HTTP_REQUEST_COMPLETED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_REQUEST_COMPLETED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldHTTP_REQUEST_COMPLETED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_REQUEST_COMPLETED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) HTTP_REQUEST_CONTENT_RESET_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_REQUEST_CONTENT_RESET_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldHTTP_REQUEST_CONTENT_RESET_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_REQUEST_CONTENT_RESET_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) CLIENT_REQUEST_RETRY_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_RETRY_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldCLIENT_REQUEST_RETRY_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_RETRY_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) HTTP_RESPONSE_STARTED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_RESPONSE_STARTED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldHTTP_RESPONSE_STARTED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_RESPONSE_STARTED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) HTTP_RESPONSE_COMPLETED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_RESPONSE_COMPLETED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldHTTP_RESPONSE_COMPLETED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_RESPONSE_COMPLETED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) HTTP_RESPONSE_CONTENT_RESET_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_RESPONSE_CONTENT_RESET_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldHTTP_RESPONSE_CONTENT_RESET_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "HTTP_RESPONSE_CONTENT_RESET_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) CLIENT_REQUEST_SUCCESS_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_SUCCESS_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldCLIENT_REQUEST_SUCCESS_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_SUCCESS_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) CLIENT_REQUEST_FAILED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_FAILED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldCLIENT_REQUEST_FAILED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "CLIENT_REQUEST_FAILED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_PREPARING_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PREPARING_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_PREPARING_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PREPARING_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_STARTED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_STARTED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_STARTED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_STARTED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_COMPLETED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_COMPLETED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_COMPLETED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_COMPLETED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_FAILED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_FAILED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_FAILED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_FAILED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_CANCELED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_CANCELED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_CANCELED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_CANCELED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_PART_STARTED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PART_STARTED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_PART_STARTED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PART_STARTED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_PART_COMPLETED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PART_COMPLETED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_PART_COMPLETED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PART_COMPLETED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *EventProgressEventType) TRANSFER_PART_FAILED_EVENT() *EventProgressEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PART_FAILED_EVENT", "com/amazonaws/event/ProgressEventType")
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
	unique_x := &EventProgressEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressEventType) SetFieldTRANSFER_PART_FAILED_EVENT(val EventProgressEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEventType", "TRANSFER_PART_FAILED_EVENT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


