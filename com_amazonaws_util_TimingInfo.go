package ec2

import "github.com/timob/javabind"

type UtilTimingInfoInterface interface {
	JavaLangObjectInterface

	// public static com.amazonaws.util.TimingInfo startTiming()
	StartTiming() *UtilTimingInfo

	// public static com.amazonaws.util.TimingInfo startTimingFullSupport()
	StartTimingFullSupport() *UtilTimingInfo

	// public static com.amazonaws.util.TimingInfo startTimingFullSupport(long)
	StartTimingFullSupport2(a int64) *UtilTimingInfo

	// public static com.amazonaws.util.TimingInfo newTimingInfoFullSupport(long, long)
	NewTimingInfoFullSupport(a int64, b int64) *UtilTimingInfo

	// public static com.amazonaws.util.TimingInfo newTimingInfoFullSupport(long, long, long)
	NewTimingInfoFullSupport2(a int64, b int64, c int64) *UtilTimingInfo

	// public static com.amazonaws.util.TimingInfo unmodifiableTimingInfo(long, java.lang.Long)
	UnmodifiableTimingInfo(a int64, b int64) *UtilTimingInfo

	// public static com.amazonaws.util.TimingInfo unmodifiableTimingInfo(long, long, java.lang.Long)
	UnmodifiableTimingInfo2(a int64, b int64, c int64) *UtilTimingInfo

	// public final long getStartTime()
	GetStartTime() int64

	// public final long getStartEpochTimeMilli()
	GetStartEpochTimeMilli() int64

	// public final java.lang.Long getStartEpochTimeMilliIfKnown()
	GetStartEpochTimeMilliIfKnown() int64

	// public final long getStartTimeNano()
	GetStartTimeNano() int64

	// public final long getEndTime()
	GetEndTime() int64

	// public final long getEndEpochTimeMilli()
	GetEndEpochTimeMilli() int64

	// public final java.lang.Long getEndEpochTimeMilliIfKnown()
	GetEndEpochTimeMilliIfKnown() int64

	// public final long getEndTimeNano()
	GetEndTimeNano() int64

	// public final java.lang.Long getEndTimeNanoIfKnown()
	GetEndTimeNanoIfKnown() int64

	// public final double getTimeTakenMillis()
	GetTimeTakenMillis() float64

	// public final java.lang.Double getTimeTakenMillisIfKnown()
	GetTimeTakenMillisIfKnown() float64

	// public static double durationMilliOf(long, long)
	DurationMilliOf(a int64, b int64) float64

	// public final long getElapsedTimeMillis()
	GetElapsedTimeMillis() int64

	// public final boolean isEndTimeKnown()
	IsEndTimeKnown() bool

	// public final boolean isStartEpochTimeMilliKnown()
	IsStartEpochTimeMilliKnown() bool

	// public void setEndTime(long)
	SetEndTime(a int64) 

	// public void setEndTimeNano(long)
	SetEndTimeNano(a int64) 

	// public com.amazonaws.util.TimingInfo endTiming()
	EndTiming() *UtilTimingInfo

	// public void addSubMeasurement(java.lang.String, com.amazonaws.util.TimingInfo)
	AddSubMeasurement(a string, b UtilTimingInfoInterface) 

	// public com.amazonaws.util.TimingInfo getSubMeasurement(java.lang.String)
	GetSubMeasurement(a string) *UtilTimingInfo

	// public com.amazonaws.util.TimingInfo getSubMeasurement(java.lang.String, int)
	GetSubMeasurement2(a string, b int) *UtilTimingInfo

	// public com.amazonaws.util.TimingInfo getLastSubMeasurement(java.lang.String)
	GetLastSubMeasurement(a string) *UtilTimingInfo

	// public java.util.List<com.amazonaws.util.TimingInfo> getAllSubMeasurements(java.lang.String)
	GetAllSubMeasurements(a string) []*UtilTimingInfo

	// public java.util.Map<java.lang.String, java.util.List<com.amazonaws.util.TimingInfo>> getSubMeasurementsByName()
	GetSubMeasurementsByName() map[string]*[]*UtilTimingInfo

	// public void setCounter(java.lang.String, long)
	SetCounter(a string, b int64) 

	// public void incrementCounter(java.lang.String)
	IncrementCounter(a string) 
}

type UtilTimingInfo struct {
	JavaLangObject
}

// public static com.amazonaws.util.TimingInfo startTiming()
func (jbobject *UtilTimingInfo) StartTiming() *UtilTimingInfo {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "startTiming", "com/amazonaws/util/TimingInfo")
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.util.TimingInfo startTimingFullSupport()
func (jbobject *UtilTimingInfo) StartTimingFullSupport() *UtilTimingInfo {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "startTimingFullSupport", "com/amazonaws/util/TimingInfo")
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.util.TimingInfo startTimingFullSupport(long)
func (jbobject *UtilTimingInfo) StartTimingFullSupport2(a int64) *UtilTimingInfo {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "startTimingFullSupport", "com/amazonaws/util/TimingInfo", a)
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.util.TimingInfo newTimingInfoFullSupport(long, long)
func (jbobject *UtilTimingInfo) NewTimingInfoFullSupport(a int64, b int64) *UtilTimingInfo {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "newTimingInfoFullSupport", "com/amazonaws/util/TimingInfo", a, b)
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.util.TimingInfo newTimingInfoFullSupport(long, long, long)
func (jbobject *UtilTimingInfo) NewTimingInfoFullSupport2(a int64, b int64, c int64) *UtilTimingInfo {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "newTimingInfoFullSupport", "com/amazonaws/util/TimingInfo", a, b, c)
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.util.TimingInfo unmodifiableTimingInfo(long, java.lang.Long)
func (jbobject *UtilTimingInfo) UnmodifiableTimingInfo(a int64, b int64) *UtilTimingInfo {
	conv_b := javabind.NewGoToJavaLong()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "unmodifiableTimingInfo", "com/amazonaws/util/TimingInfo", a, conv_b.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.util.TimingInfo unmodifiableTimingInfo(long, long, java.lang.Long)
func (jbobject *UtilTimingInfo) UnmodifiableTimingInfo2(a int64, b int64, c int64) *UtilTimingInfo {
	conv_c := javabind.NewGoToJavaLong()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "unmodifiableTimingInfo", "com/amazonaws/util/TimingInfo", a, b, conv_c.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public final long getStartTime()
func (jbobject *UtilTimingInfo) GetStartTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final long getStartEpochTimeMilli()
func (jbobject *UtilTimingInfo) GetStartEpochTimeMilli() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartEpochTimeMilli", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final java.lang.Long getStartEpochTimeMilliIfKnown()
func (jbobject *UtilTimingInfo) GetStartEpochTimeMilliIfKnown() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartEpochTimeMilliIfKnown", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final long getStartTimeNano()
func (jbobject *UtilTimingInfo) GetStartTimeNano() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStartTimeNano", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final long getEndTime()
func (jbobject *UtilTimingInfo) GetEndTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final long getEndEpochTimeMilli()
func (jbobject *UtilTimingInfo) GetEndEpochTimeMilli() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndEpochTimeMilli", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final java.lang.Long getEndEpochTimeMilliIfKnown()
func (jbobject *UtilTimingInfo) GetEndEpochTimeMilliIfKnown() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndEpochTimeMilliIfKnown", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final long getEndTimeNano()
func (jbobject *UtilTimingInfo) GetEndTimeNano() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndTimeNano", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final java.lang.Long getEndTimeNanoIfKnown()
func (jbobject *UtilTimingInfo) GetEndTimeNanoIfKnown() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndTimeNanoIfKnown", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final double getTimeTakenMillis()
func (jbobject *UtilTimingInfo) GetTimeTakenMillis() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeTakenMillis", javabind.Double)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public final java.lang.Double getTimeTakenMillisIfKnown()
func (jbobject *UtilTimingInfo) GetTimeTakenMillisIfKnown() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeTakenMillisIfKnown", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static double durationMilliOf(long, long)
func (jbobject *UtilTimingInfo) DurationMilliOf(a int64, b int64) float64 {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/util/TimingInfo", "durationMilliOf", javabind.Double, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(float64)
}

// public final long getElapsedTimeMillis()
func (jbobject *UtilTimingInfo) GetElapsedTimeMillis() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getElapsedTimeMillis", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final boolean isEndTimeKnown()
func (jbobject *UtilTimingInfo) IsEndTimeKnown() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEndTimeKnown", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public final boolean isStartEpochTimeMilliKnown()
func (jbobject *UtilTimingInfo) IsStartEpochTimeMilliKnown() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isStartEpochTimeMilliKnown", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public final java.lang.String toString()
func (jbobject *UtilTimingInfo) ToString() string {
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

// public void setEndTime(long)
func (jbobject *UtilTimingInfo) SetEndTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setEndTimeNano(long)
func (jbobject *UtilTimingInfo) SetEndTimeNano(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndTimeNano", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.amazonaws.util.TimingInfo endTiming()
func (jbobject *UtilTimingInfo) EndTiming() *UtilTimingInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "endTiming", "com/amazonaws/util/TimingInfo")
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public void addSubMeasurement(java.lang.String, com.amazonaws.util.TimingInfo)
func (jbobject *UtilTimingInfo) AddSubMeasurement(a string, b UtilTimingInfoInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSubMeasurement", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/util/TimingInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public com.amazonaws.util.TimingInfo getSubMeasurement(java.lang.String)
func (jbobject *UtilTimingInfo) GetSubMeasurement(a string) *UtilTimingInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubMeasurement", "com/amazonaws/util/TimingInfo", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.util.TimingInfo getSubMeasurement(java.lang.String, int)
func (jbobject *UtilTimingInfo) GetSubMeasurement2(a string, b int) *UtilTimingInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubMeasurement", "com/amazonaws/util/TimingInfo", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.util.TimingInfo getLastSubMeasurement(java.lang.String)
func (jbobject *UtilTimingInfo) GetLastSubMeasurement(a string) *UtilTimingInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLastSubMeasurement", "com/amazonaws/util/TimingInfo", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &UtilTimingInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.util.TimingInfo> getAllSubMeasurements(java.lang.String)
func (jbobject *UtilTimingInfo) GetAllSubMeasurements(a string) []*UtilTimingInfo {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllSubMeasurements", "java/util/List", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*UtilTimingInfo)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.Map<java.lang.String, java.util.List<com.amazonaws.util.TimingInfo>> getSubMeasurementsByName()
func (jbobject *UtilTimingInfo) GetSubMeasurementsByName() map[string]*[]*UtilTimingInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubMeasurementsByName", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoList(javabind.NewJavaToGoCallable()))
	dst := new(map[string]*[]*UtilTimingInfo)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCounter(java.lang.String, long)
func (jbobject *UtilTimingInfo) SetCounter(a string, b int64)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCounter", javabind.Void, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void incrementCounter(java.lang.String)
func (jbobject *UtilTimingInfo) IncrementCounter(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incrementCounter", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


