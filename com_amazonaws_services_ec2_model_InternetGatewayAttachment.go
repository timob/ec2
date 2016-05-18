package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInternetGatewayAttachmentInterface interface {
	JavaLangObjectInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.InternetGatewayAttachment withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelInternetGatewayAttachment

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.InternetGatewayAttachment withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelInternetGatewayAttachment

	// public void setState(com.amazonaws.services.ec2.model.AttachmentStatus)
	SetState(a ServicesEc2ModelAttachmentStatusInterface) 

	// public com.amazonaws.services.ec2.model.InternetGatewayAttachment withState(com.amazonaws.services.ec2.model.AttachmentStatus)
	WithState(a ServicesEc2ModelAttachmentStatusInterface) *ServicesEc2ModelInternetGatewayAttachment

	// public com.amazonaws.services.ec2.model.InternetGatewayAttachment clone()
	Clone() *ServicesEc2ModelInternetGatewayAttachment
}

type ServicesEc2ModelInternetGatewayAttachment struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InternetGatewayAttachment()
func NewServicesEc2ModelInternetGatewayAttachment() (*ServicesEc2ModelInternetGatewayAttachment) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InternetGatewayAttachment")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInternetGatewayAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InternetGatewayAttachment withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) WithVpcId(a string) *ServicesEc2ModelInternetGatewayAttachment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/InternetGatewayAttachment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInternetGatewayAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InternetGatewayAttachment withState(java.lang.String)
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) WithState2(a string) *ServicesEc2ModelInternetGatewayAttachment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/InternetGatewayAttachment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInternetGatewayAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.AttachmentStatus)
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) SetState(a ServicesEc2ModelAttachmentStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachmentStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InternetGatewayAttachment withState(com.amazonaws.services.ec2.model.AttachmentStatus)
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) WithState(a ServicesEc2ModelAttachmentStatusInterface) *ServicesEc2ModelInternetGatewayAttachment {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/InternetGatewayAttachment", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachmentStatus"))
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
	unique_x := &ServicesEc2ModelInternetGatewayAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) ToString() string {
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
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InternetGatewayAttachment clone()
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) Clone() *ServicesEc2ModelInternetGatewayAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InternetGatewayAttachment")
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
	unique_x := &ServicesEc2ModelInternetGatewayAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInternetGatewayAttachment) Clone2() (*JavaLangObject, error) {
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


