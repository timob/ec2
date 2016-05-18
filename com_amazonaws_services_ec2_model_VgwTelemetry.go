package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelVgwTelemetryInterface interface {
	JavaLangObjectInterface

	// public void setOutsideIpAddress(java.lang.String)
	SetOutsideIpAddress(a string) 

	// public java.lang.String getOutsideIpAddress()
	GetOutsideIpAddress() string

	// public com.amazonaws.services.ec2.model.VgwTelemetry withOutsideIpAddress(java.lang.String)
	WithOutsideIpAddress(a string) *ServicesEc2ModelVgwTelemetry

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.VgwTelemetry withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelVgwTelemetry

	// public void setStatus(com.amazonaws.services.ec2.model.TelemetryStatus)
	SetStatus(a ServicesEc2ModelTelemetryStatusInterface) 

	// public com.amazonaws.services.ec2.model.VgwTelemetry withStatus(com.amazonaws.services.ec2.model.TelemetryStatus)
	WithStatus(a ServicesEc2ModelTelemetryStatusInterface) *ServicesEc2ModelVgwTelemetry

	// public void setLastStatusChange(java.util.Date)
	SetLastStatusChange(a time.Time) 

	// public java.util.Date getLastStatusChange()
	GetLastStatusChange() time.Time

	// public com.amazonaws.services.ec2.model.VgwTelemetry withLastStatusChange(java.util.Date)
	WithLastStatusChange(a time.Time) *ServicesEc2ModelVgwTelemetry

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.VgwTelemetry withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelVgwTelemetry

	// public void setAcceptedRouteCount(java.lang.Integer)
	SetAcceptedRouteCount(a int) 

	// public java.lang.Integer getAcceptedRouteCount()
	GetAcceptedRouteCount() int

	// public com.amazonaws.services.ec2.model.VgwTelemetry withAcceptedRouteCount(java.lang.Integer)
	WithAcceptedRouteCount(a int) *ServicesEc2ModelVgwTelemetry

	// public com.amazonaws.services.ec2.model.VgwTelemetry clone()
	Clone() *ServicesEc2ModelVgwTelemetry
}

type ServicesEc2ModelVgwTelemetry struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VgwTelemetry()
func NewServicesEc2ModelVgwTelemetry() (*ServicesEc2ModelVgwTelemetry) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VgwTelemetry")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVgwTelemetry{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setOutsideIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelVgwTelemetry) SetOutsideIpAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOutsideIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOutsideIpAddress()
func (jbobject *ServicesEc2ModelVgwTelemetry) GetOutsideIpAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOutsideIpAddress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VgwTelemetry withOutsideIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelVgwTelemetry) WithOutsideIpAddress(a string) *ServicesEc2ModelVgwTelemetry {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOutsideIpAddress", "com/amazonaws/services/ec2/model/VgwTelemetry", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelVgwTelemetry) SetStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelVgwTelemetry) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VgwTelemetry withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelVgwTelemetry) WithStatus2(a string) *ServicesEc2ModelVgwTelemetry {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/VgwTelemetry", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.TelemetryStatus)
func (jbobject *ServicesEc2ModelVgwTelemetry) SetStatus(a ServicesEc2ModelTelemetryStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/TelemetryStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VgwTelemetry withStatus(com.amazonaws.services.ec2.model.TelemetryStatus)
func (jbobject *ServicesEc2ModelVgwTelemetry) WithStatus(a ServicesEc2ModelTelemetryStatusInterface) *ServicesEc2ModelVgwTelemetry {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/VgwTelemetry", conv_a.Value().Cast("com/amazonaws/services/ec2/model/TelemetryStatus"))
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLastStatusChange(java.util.Date)
func (jbobject *ServicesEc2ModelVgwTelemetry) SetLastStatusChange(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLastStatusChange", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getLastStatusChange()
func (jbobject *ServicesEc2ModelVgwTelemetry) GetLastStatusChange() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLastStatusChange", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.VgwTelemetry withLastStatusChange(java.util.Date)
func (jbobject *ServicesEc2ModelVgwTelemetry) WithLastStatusChange(a time.Time) *ServicesEc2ModelVgwTelemetry {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLastStatusChange", "com/amazonaws/services/ec2/model/VgwTelemetry", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelVgwTelemetry) SetStatusMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusMessage()
func (jbobject *ServicesEc2ModelVgwTelemetry) GetStatusMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VgwTelemetry withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelVgwTelemetry) WithStatusMessage(a string) *ServicesEc2ModelVgwTelemetry {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/VgwTelemetry", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAcceptedRouteCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelVgwTelemetry) SetAcceptedRouteCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAcceptedRouteCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getAcceptedRouteCount()
func (jbobject *ServicesEc2ModelVgwTelemetry) GetAcceptedRouteCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAcceptedRouteCount", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.VgwTelemetry withAcceptedRouteCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelVgwTelemetry) WithAcceptedRouteCount(a int) *ServicesEc2ModelVgwTelemetry {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAcceptedRouteCount", "com/amazonaws/services/ec2/model/VgwTelemetry", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVgwTelemetry) ToString() string {
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
func (jbobject *ServicesEc2ModelVgwTelemetry) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVgwTelemetry) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VgwTelemetry clone()
func (jbobject *ServicesEc2ModelVgwTelemetry) Clone() *ServicesEc2ModelVgwTelemetry {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VgwTelemetry")
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
	unique_x := &ServicesEc2ModelVgwTelemetry{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVgwTelemetry) Clone2() (*JavaLangObject, error) {
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


