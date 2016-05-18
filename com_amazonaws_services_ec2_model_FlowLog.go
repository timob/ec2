package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelFlowLogInterface interface {
	JavaLangObjectInterface

	// public void setCreationTime(java.util.Date)
	SetCreationTime(a time.Time) 

	// public java.util.Date getCreationTime()
	GetCreationTime() time.Time

	// public com.amazonaws.services.ec2.model.FlowLog withCreationTime(java.util.Date)
	WithCreationTime(a time.Time) *ServicesEc2ModelFlowLog

	// public void setFlowLogId(java.lang.String)
	SetFlowLogId(a string) 

	// public java.lang.String getFlowLogId()
	GetFlowLogId() string

	// public com.amazonaws.services.ec2.model.FlowLog withFlowLogId(java.lang.String)
	WithFlowLogId(a string) *ServicesEc2ModelFlowLog

	// public void setFlowLogStatus(java.lang.String)
	SetFlowLogStatus(a string) 

	// public java.lang.String getFlowLogStatus()
	GetFlowLogStatus() string

	// public com.amazonaws.services.ec2.model.FlowLog withFlowLogStatus(java.lang.String)
	WithFlowLogStatus(a string) *ServicesEc2ModelFlowLog

	// public void setResourceId(java.lang.String)
	SetResourceId(a string) 

	// public java.lang.String getResourceId()
	GetResourceId() string

	// public com.amazonaws.services.ec2.model.FlowLog withResourceId(java.lang.String)
	WithResourceId(a string) *ServicesEc2ModelFlowLog

	// public void setTrafficType(java.lang.String)
	SetTrafficType2(a string) 

	// public java.lang.String getTrafficType()
	GetTrafficType() string

	// public com.amazonaws.services.ec2.model.FlowLog withTrafficType(java.lang.String)
	WithTrafficType2(a string) *ServicesEc2ModelFlowLog

	// public void setTrafficType(com.amazonaws.services.ec2.model.TrafficType)
	SetTrafficType(a ServicesEc2ModelTrafficTypeInterface) 

	// public com.amazonaws.services.ec2.model.FlowLog withTrafficType(com.amazonaws.services.ec2.model.TrafficType)
	WithTrafficType(a ServicesEc2ModelTrafficTypeInterface) *ServicesEc2ModelFlowLog

	// public void setLogGroupName(java.lang.String)
	SetLogGroupName(a string) 

	// public java.lang.String getLogGroupName()
	GetLogGroupName() string

	// public com.amazonaws.services.ec2.model.FlowLog withLogGroupName(java.lang.String)
	WithLogGroupName(a string) *ServicesEc2ModelFlowLog

	// public void setDeliverLogsStatus(java.lang.String)
	SetDeliverLogsStatus(a string) 

	// public java.lang.String getDeliverLogsStatus()
	GetDeliverLogsStatus() string

	// public com.amazonaws.services.ec2.model.FlowLog withDeliverLogsStatus(java.lang.String)
	WithDeliverLogsStatus(a string) *ServicesEc2ModelFlowLog

	// public void setDeliverLogsErrorMessage(java.lang.String)
	SetDeliverLogsErrorMessage(a string) 

	// public java.lang.String getDeliverLogsErrorMessage()
	GetDeliverLogsErrorMessage() string

	// public com.amazonaws.services.ec2.model.FlowLog withDeliverLogsErrorMessage(java.lang.String)
	WithDeliverLogsErrorMessage(a string) *ServicesEc2ModelFlowLog

	// public void setDeliverLogsPermissionArn(java.lang.String)
	SetDeliverLogsPermissionArn(a string) 

	// public java.lang.String getDeliverLogsPermissionArn()
	GetDeliverLogsPermissionArn() string

	// public com.amazonaws.services.ec2.model.FlowLog withDeliverLogsPermissionArn(java.lang.String)
	WithDeliverLogsPermissionArn(a string) *ServicesEc2ModelFlowLog

	// public com.amazonaws.services.ec2.model.FlowLog clone()
	Clone() *ServicesEc2ModelFlowLog
}

type ServicesEc2ModelFlowLog struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.FlowLog()
func NewServicesEc2ModelFlowLog() (*ServicesEc2ModelFlowLog) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/FlowLog")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelFlowLog{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCreationTime(java.util.Date)
func (jbobject *ServicesEc2ModelFlowLog) SetCreationTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreationTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreationTime()
func (jbobject *ServicesEc2ModelFlowLog) GetCreationTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreationTime", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.FlowLog withCreationTime(java.util.Date)
func (jbobject *ServicesEc2ModelFlowLog) WithCreationTime(a time.Time) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreationTime", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFlowLogId(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetFlowLogId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFlowLogId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFlowLogId()
func (jbobject *ServicesEc2ModelFlowLog) GetFlowLogId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFlowLogId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withFlowLogId(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithFlowLogId(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFlowLogId", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFlowLogStatus(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetFlowLogStatus(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFlowLogStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFlowLogStatus()
func (jbobject *ServicesEc2ModelFlowLog) GetFlowLogStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFlowLogStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withFlowLogStatus(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithFlowLogStatus(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFlowLogStatus", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setResourceId(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetResourceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getResourceId()
func (jbobject *ServicesEc2ModelFlowLog) GetResourceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResourceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withResourceId(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithResourceId(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceId", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTrafficType(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetTrafficType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTrafficType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getTrafficType()
func (jbobject *ServicesEc2ModelFlowLog) GetTrafficType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTrafficType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withTrafficType(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithTrafficType2(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTrafficType", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTrafficType(com.amazonaws.services.ec2.model.TrafficType)
func (jbobject *ServicesEc2ModelFlowLog) SetTrafficType(a ServicesEc2ModelTrafficTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTrafficType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/TrafficType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.FlowLog withTrafficType(com.amazonaws.services.ec2.model.TrafficType)
func (jbobject *ServicesEc2ModelFlowLog) WithTrafficType(a ServicesEc2ModelTrafficTypeInterface) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTrafficType", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("com/amazonaws/services/ec2/model/TrafficType"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLogGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetLogGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLogGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getLogGroupName()
func (jbobject *ServicesEc2ModelFlowLog) GetLogGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLogGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withLogGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithLogGroupName(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLogGroupName", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeliverLogsStatus(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetDeliverLogsStatus(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeliverLogsStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDeliverLogsStatus()
func (jbobject *ServicesEc2ModelFlowLog) GetDeliverLogsStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeliverLogsStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withDeliverLogsStatus(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithDeliverLogsStatus(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeliverLogsStatus", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeliverLogsErrorMessage(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetDeliverLogsErrorMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeliverLogsErrorMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDeliverLogsErrorMessage()
func (jbobject *ServicesEc2ModelFlowLog) GetDeliverLogsErrorMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeliverLogsErrorMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withDeliverLogsErrorMessage(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithDeliverLogsErrorMessage(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeliverLogsErrorMessage", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeliverLogsPermissionArn(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) SetDeliverLogsPermissionArn(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeliverLogsPermissionArn", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDeliverLogsPermissionArn()
func (jbobject *ServicesEc2ModelFlowLog) GetDeliverLogsPermissionArn() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeliverLogsPermissionArn", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.FlowLog withDeliverLogsPermissionArn(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLog) WithDeliverLogsPermissionArn(a string) *ServicesEc2ModelFlowLog {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeliverLogsPermissionArn", "com/amazonaws/services/ec2/model/FlowLog", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelFlowLog) ToString() string {
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

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelFlowLog) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelFlowLog) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.FlowLog clone()
func (jbobject *ServicesEc2ModelFlowLog) Clone() *ServicesEc2ModelFlowLog {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/FlowLog")
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
	unique_x := &ServicesEc2ModelFlowLog{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelFlowLog) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


