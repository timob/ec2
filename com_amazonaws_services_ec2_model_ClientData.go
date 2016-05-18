package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelClientDataInterface interface {
	JavaLangObjectInterface

	// public void setUploadStart(java.util.Date)
	SetUploadStart(a time.Time) 

	// public java.util.Date getUploadStart()
	GetUploadStart() time.Time

	// public com.amazonaws.services.ec2.model.ClientData withUploadStart(java.util.Date)
	WithUploadStart(a time.Time) *ServicesEc2ModelClientData

	// public void setUploadEnd(java.util.Date)
	SetUploadEnd(a time.Time) 

	// public java.util.Date getUploadEnd()
	GetUploadEnd() time.Time

	// public com.amazonaws.services.ec2.model.ClientData withUploadEnd(java.util.Date)
	WithUploadEnd(a time.Time) *ServicesEc2ModelClientData

	// public void setUploadSize(java.lang.Double)
	SetUploadSize(a float64) 

	// public java.lang.Double getUploadSize()
	GetUploadSize() float64

	// public com.amazonaws.services.ec2.model.ClientData withUploadSize(java.lang.Double)
	WithUploadSize(a float64) *ServicesEc2ModelClientData

	// public void setComment(java.lang.String)
	SetComment(a string) 

	// public java.lang.String getComment()
	GetComment() string

	// public com.amazonaws.services.ec2.model.ClientData withComment(java.lang.String)
	WithComment(a string) *ServicesEc2ModelClientData

	// public com.amazonaws.services.ec2.model.ClientData clone()
	Clone() *ServicesEc2ModelClientData
}

type ServicesEc2ModelClientData struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ClientData()
func NewServicesEc2ModelClientData() (*ServicesEc2ModelClientData) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ClientData")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelClientData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setUploadStart(java.util.Date)
func (jbobject *ServicesEc2ModelClientData) SetUploadStart(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUploadStart", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getUploadStart()
func (jbobject *ServicesEc2ModelClientData) GetUploadStart() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUploadStart", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ClientData withUploadStart(java.util.Date)
func (jbobject *ServicesEc2ModelClientData) WithUploadStart(a time.Time) *ServicesEc2ModelClientData {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUploadStart", "com/amazonaws/services/ec2/model/ClientData", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelClientData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUploadEnd(java.util.Date)
func (jbobject *ServicesEc2ModelClientData) SetUploadEnd(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUploadEnd", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getUploadEnd()
func (jbobject *ServicesEc2ModelClientData) GetUploadEnd() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUploadEnd", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ClientData withUploadEnd(java.util.Date)
func (jbobject *ServicesEc2ModelClientData) WithUploadEnd(a time.Time) *ServicesEc2ModelClientData {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUploadEnd", "com/amazonaws/services/ec2/model/ClientData", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelClientData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUploadSize(java.lang.Double)
func (jbobject *ServicesEc2ModelClientData) SetUploadSize(a float64)  {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUploadSize", javabind.Void, conv_a.Value().Cast("java/lang/Double"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Double getUploadSize()
func (jbobject *ServicesEc2ModelClientData) GetUploadSize() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUploadSize", "java/lang/Double")
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

// public com.amazonaws.services.ec2.model.ClientData withUploadSize(java.lang.Double)
func (jbobject *ServicesEc2ModelClientData) WithUploadSize(a float64) *ServicesEc2ModelClientData {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUploadSize", "com/amazonaws/services/ec2/model/ClientData", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelClientData{}
	unique_x.Callable = dst
	return unique_x
}

// public void setComment(java.lang.String)
func (jbobject *ServicesEc2ModelClientData) SetComment(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setComment", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getComment()
func (jbobject *ServicesEc2ModelClientData) GetComment() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getComment", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ClientData withComment(java.lang.String)
func (jbobject *ServicesEc2ModelClientData) WithComment(a string) *ServicesEc2ModelClientData {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withComment", "com/amazonaws/services/ec2/model/ClientData", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelClientData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelClientData) ToString() string {
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
func (jbobject *ServicesEc2ModelClientData) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelClientData) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ClientData clone()
func (jbobject *ServicesEc2ModelClientData) Clone() *ServicesEc2ModelClientData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ClientData")
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
	unique_x := &ServicesEc2ModelClientData{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelClientData) Clone2() (*JavaLangObject, error) {
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


