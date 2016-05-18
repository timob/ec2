package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelTagDescriptionInterface interface {
	JavaLangObjectInterface

	// public void setResourceId(java.lang.String)
	SetResourceId(a string) 

	// public java.lang.String getResourceId()
	GetResourceId() string

	// public com.amazonaws.services.ec2.model.TagDescription withResourceId(java.lang.String)
	WithResourceId(a string) *ServicesEc2ModelTagDescription

	// public void setResourceType(java.lang.String)
	SetResourceType2(a string) 

	// public java.lang.String getResourceType()
	GetResourceType() string

	// public com.amazonaws.services.ec2.model.TagDescription withResourceType(java.lang.String)
	WithResourceType2(a string) *ServicesEc2ModelTagDescription

	// public void setResourceType(com.amazonaws.services.ec2.model.ResourceType)
	SetResourceType(a ServicesEc2ModelResourceTypeInterface) 

	// public com.amazonaws.services.ec2.model.TagDescription withResourceType(com.amazonaws.services.ec2.model.ResourceType)
	WithResourceType(a ServicesEc2ModelResourceTypeInterface) *ServicesEc2ModelTagDescription

	// public void setKey(java.lang.String)
	SetKey(a string) 

	// public java.lang.String getKey()
	GetKey() string

	// public com.amazonaws.services.ec2.model.TagDescription withKey(java.lang.String)
	WithKey(a string) *ServicesEc2ModelTagDescription

	// public void setValue(java.lang.String)
	SetValue(a string) 

	// public java.lang.String getValue()
	GetValue() string

	// public com.amazonaws.services.ec2.model.TagDescription withValue(java.lang.String)
	WithValue(a string) *ServicesEc2ModelTagDescription

	// public com.amazonaws.services.ec2.model.TagDescription clone()
	Clone() *ServicesEc2ModelTagDescription
}

type ServicesEc2ModelTagDescription struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.TagDescription()
func NewServicesEc2ModelTagDescription() (*ServicesEc2ModelTagDescription) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/TagDescription")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelTagDescription{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setResourceId(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) SetResourceId(a string)  {
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
func (jbobject *ServicesEc2ModelTagDescription) GetResourceId() string {
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

// public com.amazonaws.services.ec2.model.TagDescription withResourceId(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) WithResourceId(a string) *ServicesEc2ModelTagDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceId", "com/amazonaws/services/ec2/model/TagDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTagDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setResourceType(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) SetResourceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getResourceType()
func (jbobject *ServicesEc2ModelTagDescription) GetResourceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResourceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.TagDescription withResourceType(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) WithResourceType2(a string) *ServicesEc2ModelTagDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceType", "com/amazonaws/services/ec2/model/TagDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTagDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setResourceType(com.amazonaws.services.ec2.model.ResourceType)
func (jbobject *ServicesEc2ModelTagDescription) SetResourceType(a ServicesEc2ModelResourceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResourceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.TagDescription withResourceType(com.amazonaws.services.ec2.model.ResourceType)
func (jbobject *ServicesEc2ModelTagDescription) WithResourceType(a ServicesEc2ModelResourceTypeInterface) *ServicesEc2ModelTagDescription {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceType", "com/amazonaws/services/ec2/model/TagDescription", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ResourceType"))
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
	unique_x := &ServicesEc2ModelTagDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKey(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) SetKey(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKey", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKey()
func (jbobject *ServicesEc2ModelTagDescription) GetKey() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKey", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.TagDescription withKey(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) WithKey(a string) *ServicesEc2ModelTagDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKey", "com/amazonaws/services/ec2/model/TagDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTagDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValue(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) SetValue(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getValue()
func (jbobject *ServicesEc2ModelTagDescription) GetValue() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.TagDescription withValue(java.lang.String)
func (jbobject *ServicesEc2ModelTagDescription) WithValue(a string) *ServicesEc2ModelTagDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValue", "com/amazonaws/services/ec2/model/TagDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTagDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelTagDescription) ToString() string {
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
func (jbobject *ServicesEc2ModelTagDescription) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelTagDescription) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.TagDescription clone()
func (jbobject *ServicesEc2ModelTagDescription) Clone() *ServicesEc2ModelTagDescription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/TagDescription")
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
	unique_x := &ServicesEc2ModelTagDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelTagDescription) Clone2() (*JavaLangObject, error) {
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


