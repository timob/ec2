package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDetachVolumeResultInterface interface {
	JavaLangObjectInterface

	// public void setAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
	SetAttachment(a ServicesEc2ModelVolumeAttachmentInterface) 

	// public com.amazonaws.services.ec2.model.VolumeAttachment getAttachment()
	GetAttachment() *ServicesEc2ModelVolumeAttachment

	// public com.amazonaws.services.ec2.model.DetachVolumeResult withAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
	WithAttachment(a ServicesEc2ModelVolumeAttachmentInterface) *ServicesEc2ModelDetachVolumeResult

	// public com.amazonaws.services.ec2.model.DetachVolumeResult clone()
	Clone() *ServicesEc2ModelDetachVolumeResult
}

type ServicesEc2ModelDetachVolumeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DetachVolumeResult()
func NewServicesEc2ModelDetachVolumeResult() (*ServicesEc2ModelDetachVolumeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DetachVolumeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDetachVolumeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
func (jbobject *ServicesEc2ModelDetachVolumeResult) SetAttachment(a ServicesEc2ModelVolumeAttachmentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeAttachment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeAttachment getAttachment()
func (jbobject *ServicesEc2ModelDetachVolumeResult) GetAttachment() *ServicesEc2ModelVolumeAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachment", "com/amazonaws/services/ec2/model/VolumeAttachment")
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
	unique_x := &ServicesEc2ModelVolumeAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DetachVolumeResult withAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
func (jbobject *ServicesEc2ModelDetachVolumeResult) WithAttachment(a ServicesEc2ModelVolumeAttachmentInterface) *ServicesEc2ModelDetachVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachment", "com/amazonaws/services/ec2/model/DetachVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeAttachment"))
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
	unique_x := &ServicesEc2ModelDetachVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDetachVolumeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDetachVolumeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDetachVolumeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DetachVolumeResult clone()
func (jbobject *ServicesEc2ModelDetachVolumeResult) Clone() *ServicesEc2ModelDetachVolumeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DetachVolumeResult")
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
	unique_x := &ServicesEc2ModelDetachVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDetachVolumeResult) Clone2() (*JavaLangObject, error) {
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


