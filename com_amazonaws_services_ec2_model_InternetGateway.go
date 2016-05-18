package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInternetGatewayInterface interface {
	JavaLangObjectInterface

	// public void setInternetGatewayId(java.lang.String)
	SetInternetGatewayId(a string) 

	// public java.lang.String getInternetGatewayId()
	GetInternetGatewayId() string

	// public com.amazonaws.services.ec2.model.InternetGateway withInternetGatewayId(java.lang.String)
	WithInternetGatewayId(a string) *ServicesEc2ModelInternetGateway

	// public java.util.List<com.amazonaws.services.ec2.model.InternetGatewayAttachment> getAttachments()
	GetAttachments() []*ServicesEc2ModelInternetGatewayAttachment

	// public void setAttachments(java.util.Collection<com.amazonaws.services.ec2.model.InternetGatewayAttachment>)
	SetAttachments(a []*ServicesEc2ModelInternetGatewayAttachment) 

	// public com.amazonaws.services.ec2.model.InternetGateway withAttachments(com.amazonaws.services.ec2.model.InternetGatewayAttachment...)
	WithAttachments(a ...*ServicesEc2ModelInternetGatewayAttachment) *ServicesEc2ModelInternetGateway

	// public com.amazonaws.services.ec2.model.InternetGateway withAttachments(java.util.Collection<com.amazonaws.services.ec2.model.InternetGatewayAttachment>)
	WithAttachments2(a []*ServicesEc2ModelInternetGatewayAttachment) *ServicesEc2ModelInternetGateway

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.InternetGateway withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelInternetGateway

	// public com.amazonaws.services.ec2.model.InternetGateway withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelInternetGateway

	// public com.amazonaws.services.ec2.model.InternetGateway clone()
	Clone() *ServicesEc2ModelInternetGateway
}

type ServicesEc2ModelInternetGateway struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InternetGateway()
func NewServicesEc2ModelInternetGateway() (*ServicesEc2ModelInternetGateway) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InternetGateway")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInternetGateway{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInternetGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelInternetGateway) SetInternetGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInternetGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInternetGatewayId()
func (jbobject *ServicesEc2ModelInternetGateway) GetInternetGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInternetGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InternetGateway withInternetGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelInternetGateway) WithInternetGatewayId(a string) *ServicesEc2ModelInternetGateway {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInternetGatewayId", "com/amazonaws/services/ec2/model/InternetGateway", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InternetGatewayAttachment> getAttachments()
func (jbobject *ServicesEc2ModelInternetGateway) GetAttachments() []*ServicesEc2ModelInternetGatewayAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInternetGatewayAttachment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAttachments(java.util.Collection<com.amazonaws.services.ec2.model.InternetGatewayAttachment>)
func (jbobject *ServicesEc2ModelInternetGateway) SetAttachments(a []*ServicesEc2ModelInternetGatewayAttachment)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachments", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InternetGateway withAttachments(com.amazonaws.services.ec2.model.InternetGatewayAttachment...)
func (jbobject *ServicesEc2ModelInternetGateway) WithAttachments(a ...*ServicesEc2ModelInternetGatewayAttachment) *ServicesEc2ModelInternetGateway {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InternetGatewayAttachment")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachments", "com/amazonaws/services/ec2/model/InternetGateway", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InternetGatewayAttachment"))
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InternetGateway withAttachments(java.util.Collection<com.amazonaws.services.ec2.model.InternetGatewayAttachment>)
func (jbobject *ServicesEc2ModelInternetGateway) WithAttachments2(a []*ServicesEc2ModelInternetGatewayAttachment) *ServicesEc2ModelInternetGateway {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachments", "com/amazonaws/services/ec2/model/InternetGateway", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelInternetGateway) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelInternetGateway) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InternetGateway withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelInternetGateway) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelInternetGateway {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/InternetGateway", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InternetGateway withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelInternetGateway) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelInternetGateway {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/InternetGateway", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInternetGateway) ToString() string {
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
func (jbobject *ServicesEc2ModelInternetGateway) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInternetGateway) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InternetGateway clone()
func (jbobject *ServicesEc2ModelInternetGateway) Clone() *ServicesEc2ModelInternetGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InternetGateway")
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
	unique_x := &ServicesEc2ModelInternetGateway{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInternetGateway) Clone2() (*JavaLangObject, error) {
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


