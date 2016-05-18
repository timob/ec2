package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAttachNetworkInterfaceResultInterface interface {
	JavaLangObjectInterface

	// public void setAttachmentId(java.lang.String)
	SetAttachmentId(a string) 

	// public java.lang.String getAttachmentId()
	GetAttachmentId() string

	// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult withAttachmentId(java.lang.String)
	WithAttachmentId(a string) *ServicesEc2ModelAttachNetworkInterfaceResult

	// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult clone()
	Clone() *ServicesEc2ModelAttachNetworkInterfaceResult
}

type ServicesEc2ModelAttachNetworkInterfaceResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult()
func NewServicesEc2ModelAttachNetworkInterfaceResult() (*ServicesEc2ModelAttachNetworkInterfaceResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AttachNetworkInterfaceResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAttachNetworkInterfaceResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) SetAttachmentId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachmentId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttachmentId()
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) GetAttachmentId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachmentId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult withAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) WithAttachmentId(a string) *ServicesEc2ModelAttachNetworkInterfaceResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachmentId", "com/amazonaws/services/ec2/model/AttachNetworkInterfaceResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAttachNetworkInterfaceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AttachNetworkInterfaceResult clone()
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) Clone() *ServicesEc2ModelAttachNetworkInterfaceResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AttachNetworkInterfaceResult")
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
	unique_x := &ServicesEc2ModelAttachNetworkInterfaceResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAttachNetworkInterfaceResult) Clone2() (*JavaLangObject, error) {
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


