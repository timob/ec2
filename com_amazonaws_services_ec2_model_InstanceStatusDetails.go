package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStatusDetailsInterface interface {
	JavaLangObjectInterface

	// public void setName(java.lang.String)
	SetName2(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.InstanceStatusDetails withName(java.lang.String)
	WithName2(a string) *ServicesEc2ModelInstanceStatusDetails

	// public void setName(com.amazonaws.services.ec2.model.StatusName)
	SetName(a ServicesEc2ModelStatusNameInterface) 

	// public com.amazonaws.services.ec2.model.InstanceStatusDetails withName(com.amazonaws.services.ec2.model.StatusName)
	WithName(a ServicesEc2ModelStatusNameInterface) *ServicesEc2ModelInstanceStatusDetails

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.InstanceStatusDetails withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelInstanceStatusDetails

	// public void setStatus(com.amazonaws.services.ec2.model.StatusType)
	SetStatus(a ServicesEc2ModelStatusTypeInterface) 

	// public com.amazonaws.services.ec2.model.InstanceStatusDetails withStatus(com.amazonaws.services.ec2.model.StatusType)
	WithStatus(a ServicesEc2ModelStatusTypeInterface) *ServicesEc2ModelInstanceStatusDetails

	// public void setImpairedSince(java.util.Date)
	SetImpairedSince(a time.Time) 

	// public java.util.Date getImpairedSince()
	GetImpairedSince() time.Time

	// public com.amazonaws.services.ec2.model.InstanceStatusDetails withImpairedSince(java.util.Date)
	WithImpairedSince(a time.Time) *ServicesEc2ModelInstanceStatusDetails

	// public com.amazonaws.services.ec2.model.InstanceStatusDetails clone()
	Clone() *ServicesEc2ModelInstanceStatusDetails
}

type ServicesEc2ModelInstanceStatusDetails struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceStatusDetails()
func NewServicesEc2ModelInstanceStatusDetails() (*ServicesEc2ModelInstanceStatusDetails) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceStatusDetails")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceStatusDetails{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) SetName2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getName()
func (jbobject *ServicesEc2ModelInstanceStatusDetails) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceStatusDetails withName(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) WithName2(a string) *ServicesEc2ModelInstanceStatusDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/InstanceStatusDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(com.amazonaws.services.ec2.model.StatusName)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) SetName(a ServicesEc2ModelStatusNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/StatusName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusDetails withName(com.amazonaws.services.ec2.model.StatusName)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) WithName(a ServicesEc2ModelStatusNameInterface) *ServicesEc2ModelInstanceStatusDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/InstanceStatusDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/StatusName"))
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
	unique_x := &ServicesEc2ModelInstanceStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) SetStatus2(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceStatusDetails) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.InstanceStatusDetails withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) WithStatus2(a string) *ServicesEc2ModelInstanceStatusDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/InstanceStatusDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.StatusType)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) SetStatus(a ServicesEc2ModelStatusTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/StatusType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusDetails withStatus(com.amazonaws.services.ec2.model.StatusType)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) WithStatus(a ServicesEc2ModelStatusTypeInterface) *ServicesEc2ModelInstanceStatusDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/InstanceStatusDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/StatusType"))
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
	unique_x := &ServicesEc2ModelInstanceStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImpairedSince(java.util.Date)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) SetImpairedSince(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImpairedSince", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getImpairedSince()
func (jbobject *ServicesEc2ModelInstanceStatusDetails) GetImpairedSince() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImpairedSince", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.InstanceStatusDetails withImpairedSince(java.util.Date)
func (jbobject *ServicesEc2ModelInstanceStatusDetails) WithImpairedSince(a time.Time) *ServicesEc2ModelInstanceStatusDetails {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImpairedSince", "com/amazonaws/services/ec2/model/InstanceStatusDetails", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelInstanceStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceStatusDetails) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceStatusDetails) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceStatusDetails) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceStatusDetails clone()
func (jbobject *ServicesEc2ModelInstanceStatusDetails) Clone() *ServicesEc2ModelInstanceStatusDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceStatusDetails")
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
	unique_x := &ServicesEc2ModelInstanceStatusDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceStatusDetails) Clone2() (*JavaLangObject, error) {
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


