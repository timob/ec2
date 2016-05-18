package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAvailabilityZoneMessageInterface interface {
	JavaLangObjectInterface

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public java.lang.String getMessage()
	GetMessage() string

	// public com.amazonaws.services.ec2.model.AvailabilityZoneMessage withMessage(java.lang.String)
	WithMessage(a string) *ServicesEc2ModelAvailabilityZoneMessage

	// public com.amazonaws.services.ec2.model.AvailabilityZoneMessage clone()
	Clone() *ServicesEc2ModelAvailabilityZoneMessage
}

type ServicesEc2ModelAvailabilityZoneMessage struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AvailabilityZoneMessage()
func NewServicesEc2ModelAvailabilityZoneMessage() (*ServicesEc2ModelAvailabilityZoneMessage) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AvailabilityZoneMessage")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAvailabilityZoneMessage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setMessage(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) SetMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getMessage()
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) GetMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AvailabilityZoneMessage withMessage(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) WithMessage(a string) *ServicesEc2ModelAvailabilityZoneMessage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMessage", "com/amazonaws/services/ec2/model/AvailabilityZoneMessage", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAvailabilityZoneMessage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) ToString() string {
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
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AvailabilityZoneMessage clone()
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) Clone() *ServicesEc2ModelAvailabilityZoneMessage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AvailabilityZoneMessage")
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
	unique_x := &ServicesEc2ModelAvailabilityZoneMessage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAvailabilityZoneMessage) Clone2() (*JavaLangObject, error) {
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


