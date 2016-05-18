package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAttachVolumeResultInterface interface {
	JavaLangObjectInterface

	// public void setAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
	SetAttachment(a ServicesEc2ModelVolumeAttachmentInterface) 

	// public com.amazonaws.services.ec2.model.VolumeAttachment getAttachment()
	GetAttachment() *ServicesEc2ModelVolumeAttachment

	// public com.amazonaws.services.ec2.model.AttachVolumeResult withAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
	WithAttachment(a ServicesEc2ModelVolumeAttachmentInterface) *ServicesEc2ModelAttachVolumeResult

	// public com.amazonaws.services.ec2.model.AttachVolumeResult clone()
	Clone() *ServicesEc2ModelAttachVolumeResult
}

type ServicesEc2ModelAttachVolumeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AttachVolumeResult()
func NewServicesEc2ModelAttachVolumeResult() (*ServicesEc2ModelAttachVolumeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AttachVolumeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAttachVolumeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
func (jbobject *ServicesEc2ModelAttachVolumeResult) SetAttachment(a ServicesEc2ModelVolumeAttachmentInterface)  {
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
func (jbobject *ServicesEc2ModelAttachVolumeResult) GetAttachment() *ServicesEc2ModelVolumeAttachment {
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

// public com.amazonaws.services.ec2.model.AttachVolumeResult withAttachment(com.amazonaws.services.ec2.model.VolumeAttachment)
func (jbobject *ServicesEc2ModelAttachVolumeResult) WithAttachment(a ServicesEc2ModelVolumeAttachmentInterface) *ServicesEc2ModelAttachVolumeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachment", "com/amazonaws/services/ec2/model/AttachVolumeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeAttachment"))
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
	unique_x := &ServicesEc2ModelAttachVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAttachVolumeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAttachVolumeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAttachVolumeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AttachVolumeResult clone()
func (jbobject *ServicesEc2ModelAttachVolumeResult) Clone() *ServicesEc2ModelAttachVolumeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AttachVolumeResult")
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
	unique_x := &ServicesEc2ModelAttachVolumeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAttachVolumeResult) Clone2() (*JavaLangObject, error) {
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


