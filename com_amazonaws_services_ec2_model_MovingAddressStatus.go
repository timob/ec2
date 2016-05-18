package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelMovingAddressStatusInterface interface {
	JavaLangObjectInterface

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.MovingAddressStatus withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelMovingAddressStatus

	// public void setMoveStatus(java.lang.String)
	SetMoveStatus2(a string) 

	// public java.lang.String getMoveStatus()
	GetMoveStatus() string

	// public com.amazonaws.services.ec2.model.MovingAddressStatus withMoveStatus(java.lang.String)
	WithMoveStatus2(a string) *ServicesEc2ModelMovingAddressStatus

	// public void setMoveStatus(com.amazonaws.services.ec2.model.MoveStatus)
	SetMoveStatus(a ServicesEc2ModelMoveStatusInterface) 

	// public com.amazonaws.services.ec2.model.MovingAddressStatus withMoveStatus(com.amazonaws.services.ec2.model.MoveStatus)
	WithMoveStatus(a ServicesEc2ModelMoveStatusInterface) *ServicesEc2ModelMovingAddressStatus

	// public com.amazonaws.services.ec2.model.MovingAddressStatus clone()
	Clone() *ServicesEc2ModelMovingAddressStatus
}

type ServicesEc2ModelMovingAddressStatus struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.MovingAddressStatus()
func NewServicesEc2ModelMovingAddressStatus() (*ServicesEc2ModelMovingAddressStatus) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/MovingAddressStatus")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelMovingAddressStatus{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelMovingAddressStatus) SetPublicIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIp()
func (jbobject *ServicesEc2ModelMovingAddressStatus) GetPublicIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.MovingAddressStatus withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelMovingAddressStatus) WithPublicIp(a string) *ServicesEc2ModelMovingAddressStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/MovingAddressStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelMovingAddressStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMoveStatus(java.lang.String)
func (jbobject *ServicesEc2ModelMovingAddressStatus) SetMoveStatus2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMoveStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getMoveStatus()
func (jbobject *ServicesEc2ModelMovingAddressStatus) GetMoveStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMoveStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.MovingAddressStatus withMoveStatus(java.lang.String)
func (jbobject *ServicesEc2ModelMovingAddressStatus) WithMoveStatus2(a string) *ServicesEc2ModelMovingAddressStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMoveStatus", "com/amazonaws/services/ec2/model/MovingAddressStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelMovingAddressStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMoveStatus(com.amazonaws.services.ec2.model.MoveStatus)
func (jbobject *ServicesEc2ModelMovingAddressStatus) SetMoveStatus(a ServicesEc2ModelMoveStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMoveStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/MoveStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.MovingAddressStatus withMoveStatus(com.amazonaws.services.ec2.model.MoveStatus)
func (jbobject *ServicesEc2ModelMovingAddressStatus) WithMoveStatus(a ServicesEc2ModelMoveStatusInterface) *ServicesEc2ModelMovingAddressStatus {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMoveStatus", "com/amazonaws/services/ec2/model/MovingAddressStatus", conv_a.Value().Cast("com/amazonaws/services/ec2/model/MoveStatus"))
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
	unique_x := &ServicesEc2ModelMovingAddressStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelMovingAddressStatus) ToString() string {
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
func (jbobject *ServicesEc2ModelMovingAddressStatus) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelMovingAddressStatus) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.MovingAddressStatus clone()
func (jbobject *ServicesEc2ModelMovingAddressStatus) Clone() *ServicesEc2ModelMovingAddressStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/MovingAddressStatus")
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
	unique_x := &ServicesEc2ModelMovingAddressStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelMovingAddressStatus) Clone2() (*JavaLangObject, error) {
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


