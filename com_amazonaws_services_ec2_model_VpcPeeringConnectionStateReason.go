package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpcPeeringConnectionStateReasonInterface interface {
	JavaLangObjectInterface

	// public void setCode(java.lang.String)
	SetCode2(a string) 

	// public java.lang.String getCode()
	GetCode() string

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason withCode(java.lang.String)
	WithCode2(a string) *ServicesEc2ModelVpcPeeringConnectionStateReason

	// public void setCode(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode)
	SetCode(a ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) 

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason withCode(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode)
	WithCode(a ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) *ServicesEc2ModelVpcPeeringConnectionStateReason

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public java.lang.String getMessage()
	GetMessage() string

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason withMessage(java.lang.String)
	WithMessage(a string) *ServicesEc2ModelVpcPeeringConnectionStateReason

	// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason clone()
	Clone() *ServicesEc2ModelVpcPeeringConnectionStateReason
}

type ServicesEc2ModelVpcPeeringConnectionStateReason struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason()
func NewServicesEc2ModelVpcPeeringConnectionStateReason() (*ServicesEc2ModelVpcPeeringConnectionStateReason) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpcPeeringConnectionStateReason{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCode(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) SetCode2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCode()
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) GetCode() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCode", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason withCode(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) WithCode2(a string) *ServicesEc2ModelVpcPeeringConnectionStateReason {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCode(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) SetCode(a ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason withCode(com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) WithCode(a ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) *ServicesEc2ModelVpcPeeringConnectionStateReason {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMessage(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) SetMessage(a string)  {
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
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) GetMessage() string {
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

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason withMessage(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) WithMessage(a string) *ServicesEc2ModelVpcPeeringConnectionStateReason {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMessage", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) ToString() string {
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
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReason clone()
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) Clone() *ServicesEc2ModelVpcPeeringConnectionStateReason {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReason")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReason) Clone2() (*JavaLangObject, error) {
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


