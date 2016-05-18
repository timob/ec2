package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelIdFormatInterface interface {
	JavaLangObjectInterface

	// public void setResource(java.lang.String)
	SetResource(a string) 

	// public java.lang.String getResource()
	GetResource() string

	// public com.amazonaws.services.ec2.model.IdFormat withResource(java.lang.String)
	WithResource(a string) *ServicesEc2ModelIdFormat

	// public void setUseLongIds(java.lang.Boolean)
	SetUseLongIds(a bool) 

	// public java.lang.Boolean getUseLongIds()
	GetUseLongIds() bool

	// public com.amazonaws.services.ec2.model.IdFormat withUseLongIds(java.lang.Boolean)
	WithUseLongIds(a bool) *ServicesEc2ModelIdFormat

	// public java.lang.Boolean isUseLongIds()
	IsUseLongIds() bool

	// public void setDeadline(java.util.Date)
	SetDeadline(a time.Time) 

	// public java.util.Date getDeadline()
	GetDeadline() time.Time

	// public com.amazonaws.services.ec2.model.IdFormat withDeadline(java.util.Date)
	WithDeadline(a time.Time) *ServicesEc2ModelIdFormat

	// public com.amazonaws.services.ec2.model.IdFormat clone()
	Clone() *ServicesEc2ModelIdFormat
}

type ServicesEc2ModelIdFormat struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.IdFormat()
func NewServicesEc2ModelIdFormat() (*ServicesEc2ModelIdFormat) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/IdFormat")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelIdFormat{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setResource(java.lang.String)
func (jbobject *ServicesEc2ModelIdFormat) SetResource(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResource", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getResource()
func (jbobject *ServicesEc2ModelIdFormat) GetResource() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResource", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.IdFormat withResource(java.lang.String)
func (jbobject *ServicesEc2ModelIdFormat) WithResource(a string) *ServicesEc2ModelIdFormat {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResource", "com/amazonaws/services/ec2/model/IdFormat", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIdFormat{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUseLongIds(java.lang.Boolean)
func (jbobject *ServicesEc2ModelIdFormat) SetUseLongIds(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseLongIds", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getUseLongIds()
func (jbobject *ServicesEc2ModelIdFormat) GetUseLongIds() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUseLongIds", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.IdFormat withUseLongIds(java.lang.Boolean)
func (jbobject *ServicesEc2ModelIdFormat) WithUseLongIds(a bool) *ServicesEc2ModelIdFormat {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUseLongIds", "com/amazonaws/services/ec2/model/IdFormat", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelIdFormat{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isUseLongIds()
func (jbobject *ServicesEc2ModelIdFormat) IsUseLongIds() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isUseLongIds", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDeadline(java.util.Date)
func (jbobject *ServicesEc2ModelIdFormat) SetDeadline(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeadline", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getDeadline()
func (jbobject *ServicesEc2ModelIdFormat) GetDeadline() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeadline", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.IdFormat withDeadline(java.util.Date)
func (jbobject *ServicesEc2ModelIdFormat) WithDeadline(a time.Time) *ServicesEc2ModelIdFormat {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeadline", "com/amazonaws/services/ec2/model/IdFormat", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelIdFormat{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelIdFormat) ToString() string {
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
func (jbobject *ServicesEc2ModelIdFormat) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelIdFormat) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.IdFormat clone()
func (jbobject *ServicesEc2ModelIdFormat) Clone() *ServicesEc2ModelIdFormat {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/IdFormat")
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
	unique_x := &ServicesEc2ModelIdFormat{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelIdFormat) Clone2() (*JavaLangObject, error) {
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


