package ec2

import "github.com/timob/javabind"

type EventProgressEventInterface interface {
	JavaLangObjectInterface

	// public long getBytes()
	GetBytes() int64

	// public long getBytesTransferred()
	GetBytesTransferred() int64

	// public int getEventCode()
	GetEventCode() int

	// public com.amazonaws.event.ProgressEventType getEventType()
	GetEventType() *EventProgressEventType
}

type EventProgressEvent struct {
	JavaLangObject
}

// public com.amazonaws.event.ProgressEvent(long)
func NewEventProgressEvent(a int64) (*EventProgressEvent) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/event/ProgressEvent", a)
	if err != nil {
		panic(err)
	}
	x := &EventProgressEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.event.ProgressEvent(com.amazonaws.event.ProgressEventType)
func NewEventProgressEvent2(a EventProgressEventTypeInterface) (*EventProgressEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/event/ProgressEvent", conv_a.Value().Cast("com/amazonaws/event/ProgressEventType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventProgressEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.event.ProgressEvent(com.amazonaws.event.ProgressEventType, long)
func NewEventProgressEvent3(a EventProgressEventTypeInterface, b int64) (*EventProgressEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/event/ProgressEvent", conv_a.Value().Cast("com/amazonaws/event/ProgressEventType"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &EventProgressEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long getBytes()
func (jbobject *EventProgressEvent) GetBytes() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBytes", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getBytesTransferred()
func (jbobject *EventProgressEvent) GetBytesTransferred() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBytesTransferred", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public int getEventCode()
func (jbobject *EventProgressEvent) GetEventCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.event.ProgressEventType getEventType()
func (jbobject *EventProgressEvent) GetEventType() *EventProgressEventType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventType", "com/amazonaws/event/ProgressEventType")
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

// public java.lang.String toString()
func (jbobject *EventProgressEvent) ToString() string {
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

func (jbobject *EventProgressEvent) PREPARING_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "PREPARING_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldPREPARING_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "PREPARING_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) STARTED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "STARTED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldSTARTED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "STARTED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) COMPLETED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "COMPLETED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldCOMPLETED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "COMPLETED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) FAILED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "FAILED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldFAILED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "FAILED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) CANCELED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "CANCELED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldCANCELED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "CANCELED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) RESET_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "RESET_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldRESET_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "RESET_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) PART_STARTED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "PART_STARTED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldPART_STARTED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "PART_STARTED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) PART_COMPLETED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "PART_COMPLETED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldPART_COMPLETED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "PART_COMPLETED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *EventProgressEvent) PART_FAILED_EVENT_CODE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressEvent", "PART_FAILED_EVENT_CODE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *EventProgressEvent) SetFieldPART_FAILED_EVENT_CODE(val int) {
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressEvent", "PART_FAILED_EVENT_CODE", val)
	if err != nil {
		panic(err)
	}

}


