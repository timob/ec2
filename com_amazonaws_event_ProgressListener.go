package ec2

import "github.com/timob/javabind"

type EventProgressListenerInterface interface {

	// public abstract void progressChanged(com.amazonaws.event.ProgressEvent)
	ProgressChanged(a EventProgressEventInterface) 
}

type EventProgressListener struct {
	JavaLangObject
}

// public abstract void progressChanged(com.amazonaws.event.ProgressEvent)
func (jbobject *EventProgressListener) ProgressChanged(a EventProgressEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "progressChanged", javabind.Void, conv_a.Value().Cast("com/amazonaws/event/ProgressEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

func (jbobject *EventProgressListener) NOOP() *EventProgressListener {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/event/ProgressListener", "NOOP", "com/amazonaws/event/ProgressListener")
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
	unique_x := &EventProgressListener{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *EventProgressListener) SetFieldNOOP(val EventProgressListenerInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/event/ProgressListener", "NOOP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


